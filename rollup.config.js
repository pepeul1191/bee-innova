import { spawn } from 'child_process';
import svelte from 'rollup-plugin-svelte';
import commonjs from '@rollup/plugin-commonjs';
import terser from '@rollup/plugin-terser';
import resolve from '@rollup/plugin-node-resolve';
import css from 'rollup-plugin-css-only';
import copy from 'rollup-plugin-copy';

const production = !process.env.ROLLUP_WATCH;

const Admin = {
	input: 'src/entries/admin.js',
	output: {
		sourcemap: true,
		format: 'iife',
		name: 'admin',
		file: production ? 'static/dist/admin.min.js' : 'static/dist/admin.js',
	},
	plugins: [
		svelte({
			compilerOptions: {
				dev: !production
			}
		}),
		css({ output: production ?  'admin.min.css' : 'admin.css' }),
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

export default [Vendor, Web, Admin];