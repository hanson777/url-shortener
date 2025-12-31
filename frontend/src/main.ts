import { mount } from 'svelte'
import './app.css'
import AppSvelte from './App.svelte'

const app = mount(AppSvelte, {
  target: document.getElementById('app')!,
})

export default app
