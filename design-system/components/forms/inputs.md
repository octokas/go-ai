# Form Inputs

## Text Input
```html
<div class="form-field">
  <label class="form-label" for="text-input">Label</label>
  <div class="input-wrapper">
    <input
      type="text"
      id="text-input"
      class="form-input"
      placeholder="Enter text..."
    >
    <div class="input-icon">
      <svg><!-- Optional icon SVG --></svg>
    </div>
  </div>
  <span class="form-hint">Helper text goes here</span>
  <span class="form-error">Error message</span>
</div>
```

```scss
.form-field {
  --input-height: 40px;
  --input-padding: var(--space-3);
  --input-radius: var(--radius-md);

  display: flex;
  flex-direction: column;
  gap: var(--space-2);

  &.error {
    .form-input {
      border-color: var(--color-error);
      &:focus {
        box-shadow: 0 0 0 2px var(--color-error-light);
      }
    }
  }
}

.form-input {
  height: var(--input-height);
  padding: var(--input-padding);
  border: 1px solid var(--color-border);
  border-radius: var(--input-radius);
  background: var(--color-input-bg);
  font-family: var(--font-primary);
  font-size: var(--text-base);
  transition: all var(--duration-fast) var(--ease-default);

  &:hover {
    border-color: var(--color-border-hover);
  }

  &:focus {
    outline: none;
    border-color: var(--color-primary);
    box-shadow: 0 0 0 2px var(--color-primary-light);
  }
}
