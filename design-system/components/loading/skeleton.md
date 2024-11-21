# Skeleton Loading

```html
<div class="skeleton-wrapper">
  <div class="skeleton-header">
    <div class="skeleton avatar"></div>
    <div class="skeleton text"></div>
  </div>
  <div class="skeleton-content">
    <div class="skeleton text"></div>
    <div class="skeleton text"></div>
    <div class="skeleton text short"></div>
  </div>
</div>
```

```scss
.skeleton {
  --skeleton-color: var(--color-skeleton);
  --skeleton-shine: var(--color-skeleton-shine);

  background: var(--skeleton-color);
  border-radius: var(--radius-md);
  overflow: hidden;
  position: relative;

  &::after {
    content: '';
    position: absolute;
    inset: 0;
    transform: translateX(-100%);
    background-image: linear-gradient(
      90deg,
      transparent,
      var(--skeleton-shine),
      transparent
    );
    animation: skeleton-shine 1.5s infinite;
  }

  &.avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
  }

  &.text {
    height: 1em;
    margin: var(--space-2) 0;

    &.short {
      width: 60%;
    }
  }
}

@keyframes skeleton-shine {
  100% { transform: translateX(100%); }
}
