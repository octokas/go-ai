# Advanced Form Components

## Textarea
```html
<div class="form-field">
  <label class="form-label" for="textarea">Message</label>
  <div class="textarea-wrapper">
    <textarea
      id="textarea"
      class="form-textarea"
      rows="4"
      placeholder="Enter your message..."
    ></textarea>
    <div class="textarea-counter">0/500</div>
  </div>
</div>
```

```scss
.form-textarea {
  --textarea-min-height: 120px;

  min-height: var(--textarea-min-height);
  width: 100%;
  padding: var(--space-3);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
  resize: vertical;
  font-family: var(--font-primary);
  line-height: var(--leading-relaxed);
  transition: all var(--duration-fast) var(--ease-default);

  &:focus {
    outline: none;
    border-color: var(--color-primary);
    box-shadow: 0 0 0 2px var(--color-primary-light);
  }
}

.textarea-counter {
  font-size: var(--text-sm);
  color: var(--color-text-muted);
  text-align: right;
  margin-top: var(--space-1);
}
```

## Range Slider
```html
<div class="form-field">
  <label class="form-label">
    <span class="label-text">Volume</span>
    <span class="label-value">50%</span>
  </label>
  <div class="range-wrapper">
    <input
      type="range"
      class="form-range"
      min="0"
      max="100"
      value="50"
    >
    <div class="range-track">
      <div class="range-progress"></div>
      <div class="range-markers">
        <span class="marker" style="left: 0%">0</span>
        <span class="marker" style="left: 50%">50</span>
        <span class="marker" style="left: 100%">100</span>
      </div>
    </div>
  </div>
</div>
```

```scss
.form-range {
  --range-height: 6px;
  --thumb-size: 20px;

  appearance: none;
  width: 100%;
  height: var(--range-height);
  background: transparent;
  position: relative;
  z-index: 1;

  &::-webkit-slider-thumb {
    appearance: none;
    width: var(--thumb-size);
    height: var(--thumb-size);
    border-radius: 50%;
    background: var(--color-primary);
    border: 2px solid var(--color-surface);
    box-shadow: var(--shadow-sm);
    cursor: pointer;
    transition: transform var(--duration-fast) var(--ease-spring);

    &:hover {
      transform: scale(1.1);
    }

    &:active {
      transform: scale(0.95);
    }
  }
}

.range-track {
  position: absolute;
  left: 0;
  right: 0;
  top: 50%;
  height: var(--range-height);
  background: var(--color-border);
  border-radius: var(--radius-full);
  transform: translateY(-50%);
}

.range-progress {
  height: 100%;
  background: var(--color-primary);
  border-radius: inherit;
  width: var(--progress-width);
}
```

## Radio Group
```html
<div class="form-field">
  <label class="form-label">Options</label>
  <div class="radio-group" role="radiogroup">
    <label class="radio-label">
      <input type="radio" name="options" class="radio-input">
      <span class="radio-control"></span>
      <span class="radio-text">Option 1</span>
    </label>
    <!-- More radio options -->
  </div>
</div>
```

```scss
.radio-group {
  display: flex;
  flex-direction: column;
  gap: var(--space-3);
}

.radio-label {
  --radio-size: 20px;

  display: flex;
  align-items: center;
  gap: var(--space-3);
  cursor: pointer;

  &:hover .radio-control {
    border-color: var(--color-primary);
  }
}

.radio-input {
  position: absolute;
  opacity: 0;

  &:checked + .radio-control {
    border-color: var(--color-primary);

    &::after {
      transform: scale(1);
    }
  }

  &:focus + .radio-control {
    box-shadow: 0 0 0 2px var(--color-primary-light);
  }
}

.radio-control {
  width: var(--radio-size);
  height: var(--radio-size);
  border: 2px solid var(--color-border);
  border-radius: 50%;
  position: relative;
  transition: all var(--duration-fast) var(--ease-default);

  &::after {
    content: '';
    position: absolute;
    inset: 4px;
    border-radius: 50%;
    background: var(--color-primary);
    transform: scale(0);
    transition: transform var(--duration-fast) var(--ease-spring);
  }
}
```
