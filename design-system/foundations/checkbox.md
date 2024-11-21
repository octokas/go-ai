# Checkbox

```html
<label class="checkbox-wrapper">
  <input type="checkbox" class="checkbox-input">
  <span class="checkbox-control">
    <svg class="checkbox-check" viewBox="0 0 14 14">
      <path d="M11.914 2.914l-5.5 5.5-2.828-2.828"/>
    </svg>
  </span>
  <span class="checkbox-label">Checkbox Label</span>
</label>
```

```scss
.checkbox-wrapper {
  --checkbox-size: 20px;

  display: inline-flex;
  align-items: center;
  gap: var(--space-2);

  .checkbox-input {
    position: absolute;
    opacity: 0;

    &:checked + .checkbox-control {
      background: var(--color-primary);
      border-color: var(--color-primary);

      .checkbox-check {
        opacity: 1;
        transform: scale(1);
      }
    }
  }

  .checkbox-control {
    width: var(--checkbox-size);
    height: var(--checkbox-size);
    border: 2px solid var(--color-border);
    border-radius: var(--radius-sm);
    transition: all var(--duration-fast) var(--ease-spring);

    .checkbox-check {
      stroke: white;
      stroke-width: 2;
      fill: none;
      opacity: 0;
      transform: scale(0.8);
      transition: all var(--duration-fast) var(--ease-bounce);
    }
  }
}
