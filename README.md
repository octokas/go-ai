# Design System Documentation

## Foundations

### Color System
We blend Apple's clean aesthetic with Monokai Pro's rich color palette:

```scss
// Light Mode (Apple-inspired)
--color-background: #ffffff        // Primary background
--color-surface: #f5f5f7          // Secondary background
--color-primary: #0071e3          // Primary actions
--color-text: #1d1d1f             // Primary text
--color-text-secondary: #86868b    // Secondary text

// Dark Mode (Monokai Pro-inspired)
--color-background-dark: #2d2a2e   // Primary background
--color-surface-dark: #363537      // Secondary background
--color-primary-dark: #78dce8      // Primary actions
--color-text-dark: #fcfcfa        // Primary text
--color-accent: {
  green: #a9dc76,                 // Success states
  yellow: #ffd866,                // Warning states
  orange: #fc9867,                // Alert states
  purple: #ab9df2,                // Info states
  red: #ff6188                    // Error states
}
```

### Typography
Following Apple's SF Pro family:

```scss
// Font Families
--font-sf-pro: "SF Pro Text"
--font-sf-pro-display: "SF Pro Display"
--font-sf-mono: "SF Mono"

// Type Scale (Apple-standard)
--text-xs: 11px    // Labels, badges
--text-sm: 13px    // Secondary text
--text-base: 15px  // Body text
--text-lg: 17px    // Emphasized body
--text-xl: 20px    // Subheadings
--text-2xl: 24px   // Section headers
--text-3xl: 28px   // Page titles
--text-4xl: 34px   // Hero text
```

### Motion
Apple-inspired spring animations:

```scss
// Timing Functions
--ease-smooth: cubic-bezier(0.2, 0, 0.38, 0.9)
--ease-bounce: cubic-bezier(0.68, -0.55, 0.265, 1.55)

// Duration Scale
--duration-instant: 100ms  // Micro-interactions
--duration-fast: 150ms    // Button states
--duration-moderate: 200ms // Component transitions
--duration-slow: 300ms    // Page transitions
```

## Component Naming Conventions

### Structure
- Component wrapper: `[component-name]`
- Component elements: `[component-name]-[element]`
- Component modifiers: `[component-name]--[modifier]`
- State classes: `is-[state]` or `has-[state]`

Example:
```html
<div class="card">
  <div class="card-header">
    <h3 class="card-title">Title</h3>
  </div>
  <div class="card-content">
    <!-- Content -->
  </div>
</div>
```

### State Classes
```scss
.is-active      // Current/selected state
.is-disabled    // Disabled state
.is-loading     // Loading state
.is-error       // Error state
.has-error      // Contains an error
.has-icon       // Contains an icon
```

## File Structure
```
design-system/
├── foundations/
│   ├── colors/
│   │   ├── _variables.scss
│   │   ├── _light.scss
│   │   └── _dark.scss
│   ├── typography/
│   │   └── _scale.scss
│   └── motion/
│       └── _spring.scss
├── components/
│   ├── core/
│   │   ├── _buttons.scss
│   │   └── _cards.scss
│   └── patterns/
│       ├── _navigation.scss
│       └── _forms.scss
└── themes/
    ├── light.scss
    └── dark.scss
```
