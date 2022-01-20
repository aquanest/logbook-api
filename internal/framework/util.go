package framework

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/labstack/echo/v4"
)

// printStartupBanner prints boot banner when starts server.
func (s *Server) printStartupBanner() {
	// Generate by https://devops.datenkollektiv.de/banner.txt/index.html: doom
	banner := `
   _                _                    _
  | |              | |                  | |
  | |  ___    __ _ | |__    ___    ___  | | __
  | | / _ \  / _' || '_ \  / _ \  / _ \ | |/ /
  | || (_) || (_| || |_) || (_) || (_) ||   <
  |_| \___/  \__, ||_.__/  \___/  \___/ |_|\_\
              __/ |
             |___/
------------------------------------------------
  `
	fmt.Println(banner)
}

// printBootMessage prints boot message when booted ther server.
func (s *Server) printBootMessage(addr *string, port *int) {
	fmt.Println("✅ listening on http://" + *addr + ":" + strconv.Itoa(*port) + " ...\n")
}

// PrintHaltMessage prints halt message when shutdown server.
func (s *Server) printHaltMessage(err *error) {
	fmt.Printf("❌ %s\n\n", *err)
	fmt.Printf("Shutting down the server.\n")
}

// printRoutes prints implemented routes when starts server.
func (s *Server) printRoutes(routes []*echo.Route) {
	header := table.Row{"Method", "Path"}
	var rows []table.Row

	for _, route := range routes {
		if route.Path == "/*" {
			continue
		}
		rows = append(rows, table.Row{
			route.Method,
			route.Path,
		})
	}

	w := table.NewWriter()
	w.SetOutputMirror(os.Stdout)
	w.AppendHeader(header)
	w.AppendRows(rows)
	w.AppendSeparator()
	w.Render()
}
