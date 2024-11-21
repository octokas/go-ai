# Loading Patterns

## 1. Skeleton Loaders
```html
<div class="skeleton-table">
  <div class="skeleton-header">
    <div class="skeleton-cell" style="width: 30%"></div>
    <div class="skeleton-cell" style="width: 20%"></div>
    <div class="skeleton-cell" style="width: 25%"></div>
    <div class="skeleton-cell" style="width: 25%"></div>
  </div>

  <div class="skeleton-body">
    <div class="skeleton-row">
      <div class="skeleton-cell">
        <div class="skeleton-content">
          <div class="skeleton-avatar"></div>
          <div class="skeleton-text"></div>
        </div>
      </div>
      <div class="skeleton-cell">
        <div class="skeleton-badge"></div>
      </div>
      <div class="skeleton-cell">
        <div class="skeleton-text short"></div>
      </div>
      <div class="skeleton-cell">
        <div class="skeleton-actions">
          <div class="skeleton-button"></div>
          <div class="skeleton-button"></div>
        </div>
      </div>
    </div>
    <!-- Repeat skeleton-row -->
  </div>
</div>
```

```scss
.skeleton-table {
  --skeleton-bg: var(--color-skeleton);
  --skeleton-shine: var(--color-skeleton-shine);
  --animation-duration: 1.5s;

  @keyframes shimmer {
    0% {
      transform: translateX(-100%);
    }
    100% {
      transform: translateX(100%);
    }
  }

  .skeleton-cell {
    position: relative;
    overflow: hidden;
    background: var(--skeleton-bg);
    border-radius: var(--radius-md);

    &::after {
      content: '';
      position: absolute;
      inset: 0;
      background: linear-gradient(
        90deg,
        transparent,
        var(--skeleton-shine),
        transparent
      );
      animation: shimmer var(--animation-duration) infinite;
    }
  }

  .skeleton-avatar {
    width: 40px;
    height: 40px;
    border-radius: 50%;
  }

  .skeleton-text {
    height: 16px;
    width: 100%;

    &.short {
      width: 60%;
    }
  }
}
```

## 2. Progressive Loading
```html
<div class="progressive-table">
  <div class="table-overlay">
    <div class="loading-state">
      <svg class="progress-ring">
        <circle class="progress-ring-circle" />
      </svg>
      <div class="loading-text">
        Loading rows 1-50
        <span class="loading-count">25%</span>
      </div>
    </div>
  </div>

  <table class="table loading">
    <!-- Partially loaded content -->
  </table>
</div>
```

```scss
.progressive-table {
  --progress-size: 48px;
  --progress-stroke: 4px;
  --progress-color: var(--color-primary);

  .table-overlay {
    position: absolute;
    inset: 0;
    background: rgba(255, 255, 255, 0.8);
    backdrop-filter: blur(2px);
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .progress-ring {
    width: var(--progress-size);
    height: var(--progress-size);
    transform: rotate(-90deg);

    &-circle {
      stroke: var(--progress-color);
      stroke-width: var(--progress-stroke);
      fill: transparent;
      stroke-dasharray: var(--circumference);
      stroke-dashoffset: var(--offset);
      transition: stroke-dashoffset 0.35s;
    }
  }
}
```

## 3. Staggered Loading
```html
<div class="staggered-content">
  <div class="stagger-item" style="--stagger-delay: 0">
    <!-- Content -->
  </div>
  <div class="stagger-item" style="--stagger-delay: 1">
    <!-- Content -->
  </div>
  <!-- More items -->
</div>
```

```scss
.staggered-content {
  --stagger-base-delay: 50ms;

  .stagger-item {
    opacity: 0;
    transform: translateY(10px);
    animation: staggerIn 0.5s ease-out forwards;
    animation-delay: calc(var(--stagger-delay) * var(--stagger-base-delay));
  }

  @keyframes staggerIn {
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
}
```

## 4. Inline Loading States
```html
<button class="btn-loading">
  <span class="btn-text">Save Changes</span>
  <span class="btn-loader">
    <svg class="spinner"><!-- Spinner SVG --></svg>
  </span>
</button>

<div class="cell-loading">
  <div class="pulse-loader">
    <span></span>
    <span></span>
    <span></span>
  </div>
</div>
```

```scss
.btn-loading {
  position: relative;

  .btn-loader {
    position: absolute;
    inset: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    background: inherit;
    opacity: 0;
    transition: opacity 0.2s;
  }

  &.loading {
    .btn-text { opacity: 0; }
    .btn-loader { opacity: 1; }
  }
}

.pulse-loader {
  display: flex;
  gap: 4px;

  span {
    width: 4px;
    height: 4px;
    border-radius: 50%;
    background: currentColor;
    animation: pulse 1s ease-in-out infinite;

    &:nth-child(2) { animation-delay: 0.2s; }
    &:nth-child(3) { animation-delay: 0.4s; }
  }
}
