# Typography System

Our typography system combines Apple's SF Pro with MonoLisa for code, creating a premium native feel while maintaining excellent readability.

## Primary Typefaces

### SF Pro
Our primary typeface for UI elements and content.
```scss
--font-primary: "SF Pro Text", -apple-system, BlinkMacSystemFont, "Montserrat", sans-serif;
--font-display: "SF Pro Display", -apple-system, BlinkMacSystemFont, "Montserrat", sans-serif;
```

### MonoLisa

Our monospace typeface for code and numerical data.
```scss
--font-mono: "MonoLisa", "SF Mono", "SFMono-Regular", Courier, ui-monospace, monospace;
```

## Type Scale
Using a 1.25 modular scale, we create a harmonious typographic system that scales beautifully.

### Font Sizes
```scss
--text-xs: 0.75rem; //12px
--text-sm: 0.875rem; //14px
--text-base: 1rem; //16px
--text-lg: 1.25rem; //20px
--text-xl: 1.5rem; //24px
--text-2xl: 1.875rem; //30px
--text-3xl: 2.25rem; //36px
--text-4xl: 3rem; //48px
```

### Font Weights
```scss
--font-light: 300;
--font-regular: 400;
--font-medium: 500;
--font-semibold: 600;
--font-bold: 700;
```

### Line Heights
```scss
--leading-none: 1rem; //16px
--leading-tight: 1.25rem; //20px
--leading-normal: 1.5rem; //24px
--leading-relaxed: 1.75rem; //28px
--leading-loose: 2rem; //32px
```

### Letter Spacing
```scss
--tracking-tighter: -0.02em; //-0.32px
--tracking-tight: -0.01em; //-0.16px
--tracking-normal: 0em; //0px
--tracking-wide: 0.01em; //0.16px
--tracking-wider: 0.02em; //0.32px
--tracking-widest: 0.04em; //0.64px
```
