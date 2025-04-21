package main

import (
	"log"
	"net/http"
	"os"
    "path/filepath" // Use filepath for OS-agnostic path joining
	"path" // Use path for URL path cleaning
	"strings"

	"github.com/gin-gonic/gin"
)

// Define the directory where static files (including markdown) are located.
// Relative to the executable's working directory.
const bookDir = "book"

func main() {
	// Check if the book directory exists
    if _, err := os.Stat(bookDir); os.IsNotExist(err) {
        log.Fatalf("Fatal: Static file directory '%s' not found relative to the working directory. Make sure it exists.", bookDir)
    } else if err != nil {
        log.Fatalf("Fatal: Error checking static file directory '%s': %v", bookDir, err)
    } else {
        absPath, _ := filepath.Abs(bookDir) // Get absolute path for logging
		log.Printf("Serving files from directory: %s (absolute: %s)", bookDir, absPath)
	}


	router := gin.Default()
	gin.SetMode(gin.ReleaseMode) // Production mode

	// --- 1. Middleware: Prioritize serving .md files as HTML ---
	router.Use(serveMarkdownAsHTML(bookDir))
	log.Println("Markdown serving middleware registered.")

	// --- 2. Static file server: Handle remaining requests ---
	// Serve files from the 'bookDir' directory for requests starting with "/"
	// This will handle CSS, JS, images, and index.html IF the markdown middleware doesn't handle it.
	log.Printf("Registering Static file server for '/' based on directory '%s'", bookDir)
	router.Static("/", bookDir) // Serve files directly from the filesystem


	// --- Server configuration ---
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888" // Default port
		log.Printf("PORT environment variable not set, using default port %s", port)
	} else {
		log.Printf("Using port from environment variable: %s", port)
	}

	log.Printf("Starting server on :%s", port)
	log.Printf(" - .md files (or paths resolving to .md) served as text/html from '%s'.", bookDir)
	log.Printf(" - Other files served directly from '%s' (including index.html).", bookDir)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("Fatal: Failed to start server: %v", err)
	}
}

// serveMarkdownAsHTML middleware reads .md files from the filesystem
func serveMarkdownAsHTML(baseDir string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request path, clean it
		requestedPath := path.Clean(c.Request.URL.Path)
        requestedPath = strings.TrimPrefix(requestedPath, "/")

		effectivePath := requestedPath
        isIndexRequest := (requestedPath == "" || requestedPath == ".")

        if isIndexRequest {
            effectivePath = "index.md"
            log.Printf("[MD Handler] Root path request ('%s'), checking for '%s'", c.Request.URL.Path, filepath.Join(baseDir, effectivePath))
        } else if !strings.HasSuffix(strings.ToLower(effectivePath), ".md") {
            mdPath := effectivePath + ".md"
            fullMdPath := filepath.Join(baseDir, mdPath)
             // Check if the corresponding .md file exists on the filesystem
            if _, statErr := os.Stat(fullMdPath); statErr == nil {
                log.Printf("[MD Handler] Request '/%s' doesn't end with .md, but found corresponding '%s'. Using it.", requestedPath, fullMdPath)
				effectivePath = mdPath // Update to use the .md path
            } else if !os.IsNotExist(statErr) {
                log.Printf("[MD Handler] Warning: Error checking for %s: %v", fullMdPath, statErr)
            } else {
                log.Printf("[MD Handler] Request '/%s' doesn't end with .md, and corresponding '%s' not found. Passing to static handler.", requestedPath, fullMdPath)
            }
        } else {
             log.Printf("[MD Handler] Request path '%s' ends with .md. Checking file '%s'.", effectivePath, filepath.Join(baseDir, effectivePath))
        }


		// Check if the effective path targets a .md file and try to read it
		if strings.HasSuffix(strings.ToLower(effectivePath), ".md") {
            // Construct the full filesystem path
            fullFilePath := filepath.Join(baseDir, effectivePath)
            log.Printf("[MD Handler] Attempting to read potential MD file from filesystem: %s", fullFilePath)

            // Read the file directly from the filesystem
            fileContent, readErr := os.ReadFile(fullFilePath)

            if readErr == nil {
                // Successfully read the .md file
                log.Printf("[MD Handler] Success: Found and read MD file '%s'. Serving as text/html.", fullFilePath)
				c.Data(http.StatusOK, "text/html; charset=utf-8", fileContent)
				c.Abort() // IMPORTANT: Stop processing, don't let Static handler take over
                return
            }

			// Handle file reading errors
			if os.IsNotExist(readErr) {
                 if isIndexRequest {
					log.Printf("[MD Handler] Info: '%s' not found for root request. Passing to Static handler (will try index.html).", fullFilePath)
				} else {
					log.Printf("[MD Handler] Info: MD file '%s' not found. Passing request for '/%s' to Static handler.", fullFilePath, requestedPath)
				}
            } else {
				log.Printf("[MD Handler] Error: Failed to read MD file '%s': %v. Aborting with 500.", fullFilePath, readErr)
				c.String(http.StatusInternalServerError, "Internal Server Error reading markdown file.")
                c.Abort() // Internal error, stop processing
				return
			}
		}

		// If we reached here, no .md file was served. Let the request continue.
		log.Printf("[MD Handler] Path '/%s' (effective check path '%s') did not result in serving an MD file. Passing to next handler (Static).", requestedPath, effectivePath)
		c.Next() // Continue to the next handler (router.Static)
	}
}
