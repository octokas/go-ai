# Mega Menu

```html
<nav class="mega-menu" aria-label="Main Navigation">
  <div class="mega-menu-container">
    <ul class="mega-menu-primary">
      <li class="mega-menu-item">
        <button class="mega-menu-trigger">
          Products
          <svg class="mega-menu-chevron"><!-- Chevron SVG --></svg>
        </button>

        <div class="mega-menu-content">
          <div class="mega-menu-grid">
            <div class="mega-menu-section">
              <h3 class="mega-menu-title">Featured</h3>
              <ul class="mega-menu-features">
                <li class="mega-menu-feature">
                  <svg class="feature-icon"><!-- Icon SVG --></svg>
                  <div class="feature-content">
                    <h4>Feature Title</h4>
                    <p>Feature description</p>
                  </div>
                </li>
              </ul>
            </div>

            <div class="mega-menu-section">
              <h3 class="mega-menu-title">Categories</h3>
              <ul class="mega-menu-list">
                <li><a href="#">Category 1</a></li>
                <!-- More categories -->
              </ul>
            </div>
          </div>
        </div>
      </li>
    </ul>
  </div>
</nav>
```

```scss
.mega-menu {
  --mega-menu-height: 64px;
  --mega-content-width: 100%;
  --mega-max-width: 1200px;

  position: relative;
  height: var(--mega-menu-height);
  border-bottom: 1px solid var(--color-border);

  &-container {
    max-width: var(--mega-max-width);
    margin: 0 auto;
    padding: 0 var(--space-4);
  }

  &-content {
    position: absolute;
    left: 0;
    right: 0;
    top: 100%;
    background: var(--color-surface);
    box-shadow: var(--shadow-lg);
    opacity: 0;
    transform: translateY(-8px);
    pointer-events: none;
    transition: all var(--duration-normal) var(--ease-spring);

    &.active {
      opacity: 1;
      transform: translateY(0);
      pointer-events: auto;
    }
  }

  &-grid {
    display: grid;
    grid-template-columns: 2fr 1fr 1fr;
    gap: var(--space-8);
    padding: var(--space-8);
  }
}
