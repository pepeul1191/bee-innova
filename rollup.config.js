import { spawn } from 'child_process';
import svelte from 'rollup-plugin-svelte';
import commonjs from '@rollup/plugin-commonjs';
import terser from '@rollup/plugin-terser';
import resolve from '@rollup/plugin-node-resolve';
import css from 'rollup-plugin-css-only';
import copy from 'rollup-plugin-copy';

const production = !process.env.ROLLUP_WATCH;

const App = {
	input: 'src/entries/app.js',
	output: {
		sourcemap: true,
		format: 'iife',
		name: 'app',
		file: production ? 'static/dist/app.min.js' : 'static/dist/app.js',
	},
	plugins: [
		svelte({
			compilerOptions: {
				dev: !production
			}
		}),
		css({ output: production ?  'app.min.css' : 'app.css' }),
		resolve({
			browser: true,
			dedupe: ['svelte'],
			exportConditions: ['svelte']
		}),
		commonjs(),
		production && terser()
	],
	watch: {
		clearScreen: false
	}
};

const Error = {
	input: 'src/entries/error.js',
	output: {
		sourcemap: true,
		format: 'iife',
		name: 'error',
		file: production ? 'static/dist/error.min.js' : 'static/dist/error.js',
	},
	plugins: [
		svelte({
			compilerOptions: {
				dev: !production
			}
		}),
		css({ output: production ?  'error.min.css' : 'error.css' }),
		resolve({
			browser: true,
			dedupe: ['svelte'],
			exportConditions: ['svelte']
		}),
		commonjs(),
		production && terser()
	],
	watch: {
		clearScreen: false
	}
};

const Access = {
	input: 'src/entries/access.js',
	output: {
		sourcemap: true,
		format: 'iife',
		name: 'access',
		file: production ? 'static/dist/access.min.js' : 'static/dist/access.js',
	},
	plugins: [
		svelte({
			compilerOptions: {
				dev: !production
			}
		}),
		css({ output: production ?  'access.min.css' : 'access.css' }),
		resolve({
			browser: true,
			dedupe: ['svelte'],
			exportConditions: ['svelte']
		}),
		commonjs(),
		production && terser()
	],
	watch: {
		clearScreen: false
	}
};

const Vendor = {
	input: 'src/entries/vendor.js',
	output: {
			sourcemap: true,
			format: 'iife',
			name: 'vendor',
			file: production ? 'static/dist/vendor.min.js' : 'static/dist/vendor.js',
	},
	plugins: [
		svelte({
			compilerOptions: {
				dev: !production
			}
		}),
	
		css({
			output: 'vendor.min.css', // siempre este nombre
			minify: true              // siempre minificado
		}),
	
		resolve({
			browser: true,
			dedupe: ['svelte'],
			exportConditions: ['svelte']
		}),
	
		commonjs(),
		production && terser(),
	
		copy({
			hook: 'writeBundle',
			targets: [
				{
					src: 'node_modules/font-awesome/fonts/*',
					dest: 'static/fonts/'
				}
			]
		})
	],
	watch: {
			clearScreen: false
	}
};

const Web = {
  input: 'src/entries/web.js',
  output: {
    sourcemap: true,
    format: 'iife',
    name: 'web',
    file: production ? 'static/dist/web.min.js' : 'static/dist/web.js',
    globals: {
      'axios': 'axios',
      'bootstrap': 'bootstrap'
    }
  },
  plugins: [
    svelte({
      compilerOptions: {
        dev: !production
      }
    }),

    css({
      output: 'web.min.css',
      minify: true
    }),

    resolve({
      browser: true,
      dedupe: ['svelte'],
      exportConditions: ['svelte']
    }),

    commonjs(),

    production && terser(),

    copy({
      hook: 'writeBundle',
      targets: [
        {
          src: 'node_modules/font-awesome/fonts/*',
          dest: 'static/fonts/'
        }
      ]
    })
  ],
  watch: {
    clearScreen: false
  }
};

export default [App, Error, Access, Vendor, Web];