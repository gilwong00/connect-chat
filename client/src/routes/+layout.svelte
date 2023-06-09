<script>
  import '../app.css';
  import Header from './Header.svelte';
  import {
    Navbar,
    NavBrand,
    NavLi,
    NavUl,
    NavHamburger
  } from 'flowbite-svelte';

  import './styles.css';
</script>

<div class="app">
  <!-- <Header /> -->
  <Navbar let:hidden let:toggle>
    <NavBrand href="/home">
      <span
        class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
      >
        Connect Chat
      </span>
    </NavBrand>
    <NavHamburger on:click={toggle} />
    <NavUl {hidden}>
      <NavLi href="/login">Login</NavLi>
      <NavLi href="/signup">Sign up</NavLi>
      <!-- <NavLi href="/contact">Contact</NavLi> -->
    </NavUl>
  </Navbar>
  <main>
    <slot />
  </main>
</div>

<style>
  .app {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
  }

  main {
    flex: 1;
    display: flex;
    flex-direction: column;
    padding: 1rem;
    width: 100%;
    max-width: 64rem;
    margin: 0 auto;
    box-sizing: border-box;
  }
</style>
