# Select Dropdown

```html
<div class="form-field">
  <label class="form-label" for="select">Select Option</label>
  <div class="select-wrapper">
    <select class="form-select" id="select">
      <option value="">Choose an option</option>
      <option value="1">Option 1</option>
      <option value="2">Option 2</option>
    </select>
    <svg class="select-chevron"><!-- Chevron SVG --></svg>
  </div>
</div>
```

```scss
.select-wrapper {
  position: relative;

  .select-chevron {
    position: absolute;
    right: var(--space-3);
    top: 50%;
    transform: translateY(-50%);
    pointer-events: none;
    transition: transform var(--duration-fast) var(--ease-spring);
  }

  &:focus-within .select-chevron {
    transform: translateY(-50%) rotate(180deg);
  }
}

.form-select {
  appearance: none;
  width: 100%;
  padding-right: var(--space-8); // Space for chevron
  cursor: pointer;
}
