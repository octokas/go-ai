# Footer Menu

```html
<footer class="footer-menu">
  <div class="footer-container">
    <div class="footer-grid">
      <div class="footer-section">
        <h4 class="footer-title">Product</h4>
        <ul class="footer-list">
          <li><a href="#" class="footer-link">Features</a></li>
          <!-- More links -->
        </ul>
      </div>
      <!-- More sections -->
    </div>

    <div class="footer-bottom">
      <div class="footer-brand">
        <img src="logo.svg" alt="Logo" class="footer-logo">
      </div>
      <div class="footer-social">
        <!-- Social links -->
      </div>
      <div class="footer-legal">
        <!-- Legal links -->
      </div>
    </div>
  </div>
</footer>
```

```scss
.footer-menu {
  --footer-spacing: var(--space-8);

  padding: var(--footer-spacing) 0;
  background: var(--color-surface);
  border-top: 1px solid var(--color-border);

  &-container {
    max-width: var(--container-xl);
    margin: 0 auto;
    padding: 0 var(--space-4);
  }

  &-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: var(--space-8);
    margin-bottom: var(--footer-spacing);
  }

  &-bottom {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding-top: var(--space-6);
    border-top: 1px solid var(--color-border);
  }
}
