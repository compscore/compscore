import { useState, useEffect } from 'react';
import { darkTheme } from './Dark';
import { lightTheme } from './Light';
import { Theme } from '@mui/material';

export function useSystemTheme() {
    const [theme, setTheme] = useState<Theme>(darkTheme);
    
    useEffect(() => {
        const dark = window.matchMedia('(prefers-color-scheme: dark)');
        const listener = (e: MediaQueryListEvent) => {
            setTheme(e.matches ? darkTheme : lightTheme);
        };
        dark.addEventListener('change', listener);
        return () => dark.removeEventListener('change', listener);
    }, []);

    return theme;
  }
