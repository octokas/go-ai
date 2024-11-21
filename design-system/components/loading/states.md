# Loading States

## Spinner
```html
<div class="spinner" role="status">
  <svg viewBox="0 0 24 24">
    <circle class="spinner-track" cx="12" cy="12" r="10" />
    <circle class="spinner-head" cx="12" cy="12" r="10" />
  </svg>
</div>
```

```scss
.spinner {
  --spinner-size: 24px;
  --spinner-width: 2px;
  --spinner-color: var(--color-primary);

  width: var(--spinner-size);
  height: var(--spinner-size);

  svg {
    animation: spinner-rotate 2s linear infinite;
  }

  circle {
    fill: none;
    stroke-width: var(--spinner-width);
    stroke-linecap: round;

    &.spinner-track {
      stroke: var(--color-border);
    }

    &.spinner-head {
      stroke: var(--spinner-color);
      stroke-dasharray: 60;
      stroke-dashoffset: 60;
      animation: spinner-dash 1.5s ease-in-out infinite;
    }
  }
}

@keyframes spinner-rotate {
  100% { transform: rotate(360deg); }
}

@keyframes spinner-dash {
  0% {
    stroke-dasharray: 1, 150;
    stroke-dashoffset: 0;
  }
  50% {
    stroke-dasharray: 90, 150;
    stroke-dashoffset: -35;
  }
  100% {
    stroke-dasharray: 90, 150;
    stroke-dashoffset: -124;
  }
}
