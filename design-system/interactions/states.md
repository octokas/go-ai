# Interaction States System

## Button States
```scss
.interactive-element {
  --hover-scale: 1.02;
  --active-scale: 0.98;
  --focus-ring-color: var(--color-primary-light);
  --focus-ring-size: 4px;
  --disabled-opacity: 0.5;

  // Base State
  position: relative;
  transition: all var(--duration-fast) var(--ease-spring);

  // Hover State
  &:hover:not(:disabled) {
    transform: scale(var(--hover-scale));
    background: var(--color-hover);
  }

  // Active/Pressed State
  &:active:not(:disabled) {
    transform: scale(var(--active-scale));
  }

  // Focus State
  &:focus-visible {
    outline: none;
    box-shadow: 0 0 0 var(--focus-ring-size) var(--focus-ring-color);
  }

  // Disabled State
  &:disabled {
    opacity: var(--disabled-opacity);
    cursor: not-allowed;
    pointer-events: none;
  }
}

// Interactive States Mixin
@mixin interactive-states($config: ()) {
  $defaults: (
    'hover-transform': scale(1.02),
    'active-transform': scale(0.98),
    'focus-ring-color': var(--color-primary-light),
    'disabled-opacity': 0.5
  );

  $config: map-merge($defaults, $config);

  &:hover:not(:disabled) {
    transform: map-get($config, 'hover-transform');
  }

  &:active:not(:disabled) {
    transform: map-get($config, 'active-transform');
  }

  // ... other states
}
```

## Form Input States
```scss
.form-control {
  --input-border: var(--color-border);
  --input-focus: var(--color-primary);
  --input-error: var(--color-error);
  --input-success: var(--color-success);

  // Default State
  border: 1px solid var(--input-border);
  transition: all var(--duration-fast) var(--ease-default);

  // Focus State
  &:focus {
    border-color: var(--input-focus);
    box-shadow: 0 0 0 3px rgba(var(--input-focus-rgb), 0.2);
  }

  // Error State
  &.error {
    border-color: var(--input-error);
    &:focus {
      box-shadow: 0 0 0 3px rgba(var(--input-error-rgb), 0.2);
    }
  }

  // Success State
  &.success {
    border-color: var(--input-success);
    &:focus {
      box-shadow: 0 0 0 3px rgba(var(--input-success-rgb), 0.2);
    }
  }
}
```

## Drag and Drop States
```scss
.draggable {
  --drag-scale: 1.05;
  --drag-shadow: var(--shadow-lg);
  --drag-opacity: 0.9;

  &.dragging {
    transform: scale(var(--drag-scale));
    box-shadow: var(--drag-shadow);
    opacity: var(--drag-opacity);
    cursor: grabbing;
  }

  .drag-handle {
    cursor: grab;
    &:active {
      cursor: grabbing;
    }
  }
}

.droppable {
  &.drag-over {
    background: var(--color-surface-hover);
    border-color: var(--color-primary);
  }
}
```

Now, let's create our animation library:

````markdown:design-system/animations/library.md
# Animation Library

## Timing Functions
```scss
:root {
  // Natural Movement
  --ease-default: cubic-bezier(0.4, 0, 0.2, 1);
  --ease-in: cubic-bezier(0.4, 0, 1, 1);
  --ease-out: cubic-bezier(0, 0, 0.2, 1);
  --ease-in-out: cubic-bezier(0.4, 0, 0.2, 1);

  // Elastic & Bounce
  --ease-elastic: cubic-bezier(0.68, -0.55, 0.265, 1.55);
  --ease-bounce: cubic-bezier(0.87, -0.41, 0.19, 1.44);
  --ease-spring: cubic-bezier(0.43, 0.195, 0.02, 1.2);

  // Duration Scale
  --duration-instant: 100ms;
  --duration-fast: 200ms;
  --duration-normal: 300ms;
  --duration-slow: 400ms;
  --duration-slower: 500ms;
}
```

## Core Animations
```scss
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes scaleIn {
  from {
    opacity: 0;
    transform: scale(0);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}
```

## Animation Library
```scss
@keyframes fadeIn {
  from { opacity: 0; }
  to { opacity: 1; }
}

@keyframes slideIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes scaleIn {
  from {
    opacity: 0;
    transform: scale(0);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}
```
