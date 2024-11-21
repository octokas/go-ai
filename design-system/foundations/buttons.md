# Interactive Elements

## Buttons

### Base Button Styles
```scss
.button {
  --button-height: 32px;
  --button-padding: var(--space-3) var(--space-4);
  --button-radius: var(--radius-md);

  position: relative;
  height: var(--button-height);
  padding: var(--button-padding);
  border-radius: var(--button-radius);
  font-family: var(--font-primary);
  font-weight: var(--font-medium);
  font-size: var(--text-sm);
  transition: all var(--duration-fast) var(--ease-spring);

  &:focus-visible {
    outline: 2px solid var(--color-focus-ring);
    outline-offset: 2px;
  }
}
```

### Primary Button
```html
<button class="button primary">
  <span class="button-icon">
    <svg><!-- Icon SVG --></svg>
  </span>
  <span class="button-label">Primary Action</span>
  <span class="button-indicator"></span>
</button>
```

```scss
.button.primary {
  background: var(--color-primary);
  color: var(--color-primary-text);

  @media (prefers-color-scheme: dark) {
    background: var(--color-primary-dark);
    color: var(--color-primary-text-dark);
  }

  &:hover {
    transform: translateY(-1px);
    box-shadow: 0 2px 8px var(--color-primary-shadow);
  }

  &:active {
    transform: translateY(1px);
    box-shadow: none;
  }

  &.loading {
    .button-indicator {
      animation: button-loading 1s infinite;
    }
  }
}
```

### Ghost Button
```html
<button class="button ghost">
  <span class="button-label">Ghost Action</span>
</button>
```

```scss
.button.ghost {
  background: transparent;
  color: var(--color-text);
  border: 1px solid var(--color-border);

  &:hover {
    background: var(--color-hover);
    border-color: var(--color-border-hover);
  }
}
```

### Icon Button
```html
<button class="button icon-only" aria-label="Action description">
  <svg><!-- Icon SVG --></svg>
</button>
```

```scss
.button.icon-only {
  --button-size: 36px;
  width: var(--button-size);
  height: var(--button-size);
  padding: var(--space-2);
  border-radius: 50%;

  svg {
    width: 20px;
    height: 20px;
  }
}
```

### Secondary
```html
<button class="button secondary">
  <span class="button-label secondary">Secondary</span>
</button>
```

```scss
.button.secondary {
  background-color: var(--color-secondary);
}

.button-label.secondary {
  color: var(--color-secondary-text);
}
```

### Tertiary
```html
<button class="button tertiary">
  <span class="button-label tertiary">Tertiary</span>
</button>
```

```scss
.button.tertiary {
  background-color: var(--color-tertiary);
}
```
