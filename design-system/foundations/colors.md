# Color System

Our color system harmoniously blends macOS's clarity with Monokai Pro's depth, creating a unique experience that shifts dramatically between light and dark modes.

## Primary Colors

### Deep Ocean
Our primary brand color, transitioning from Apple's clarity to Monokai's richness.
```scss
static let deepOcean = Color(
  light: "#0A84FF", //apple inspired blue
  dark: "#FF6188", //monokai pro inspired magenta
  accessibility: {
    contrast: 4.5,
    role: "Primary Action"
  }
)
```
### Graphite
Our primary neutral color, used for text and UI elements.
```scss
static let graphite = Color(
  light: "#1D1D1F", // apple inspired dark gray
  dark: "#2D2A2E", // monokai pro inspired dark background
  accessibility: {
    contrast: 7.0,
    role: "Text"
  }
)
```

## Semantic Colors

### Dark Mode (Monokai Pro-inspired)
- Primary: `#FF6188` (magenta)
- Secondary: `#FC9867` (orange)
- Success: `#A9DC76` (green)
- Warning: `#FFD866` (yellow)
- Info: `#78DCE8` (blue)
- Purple: `#AB9DF2` (purple)

### Light Mode (Apple-inspired)
- Primary: `#0A84FF` (blue)
- Secondary: `#5856D6` (purple)
- Success: `#32D74B` (green)
- Warning: `#FF9F0A` (orange)
- Info: `#007AFF` (blue)
- Accent: `#BF5AF2` (purple)

## Background Scales

### Dark Mode
```scss
static let darkMode = {
    background: "#2D2A2E",    // Base
    surface: "#403E41",       // Elevated
    elevated: "#5B595C",      // Floating
    overlay: "#727072"        // Overlay
}
```

### Light Mode
```scss
static let lightMode = {
    background: "#FFFFFF",    // Base
    surface: "#F5F5F7",      // Elevated
    elevated: "#E5E5E5",     // Floating
    overlay: "#99999A"       // Overlay
}
```

## Usage Principles

### Task Priority Colors
```scss
static let priorityColors = {
    critical: {
      light: "#FF3B30",
      dark: "#FF6188"
    },
    high: {
      light: "#FF9F0A",
      dark: "#FC9867"
    },
    medium: {
      light: "#32D74B",
      dark: "#A9DC76"
    },
    low: {
      light: "#64D2FF",
      dark: "#78DCE8"
    }
}
```

### Status Indicators
```scss
static let statusColors = {
    success: {
      light: "#32D74B",
      dark: "#A9DC76",
      text: {
        light: "#1D1D1F",
        dark: "#2D2A2E"
      }
    },
    warning: {
      light: "#FF9F0A",
      dark: "#FC9867",
      text: {
        light: "#1D1D1F",
        dark: "#2D2A2E"
      }
    },
    error: {
      light: "#FF3B30",
      dark: "#FF6188",
      text: {
        light: "#FFFFFF",
        dark: "#2D2A2E"
      }
    }
}
```
