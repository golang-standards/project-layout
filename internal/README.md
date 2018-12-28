# `/internal`

Private application and library code. This is the code you don't want others importing in their applications or libraries.

You can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. It's not required, but it's nice to have visual clues showing the intended package use. Your actual application code can go in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).
