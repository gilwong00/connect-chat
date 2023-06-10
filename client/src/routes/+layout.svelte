<script>
  import '../app.css';
  import { browser } from '$app/environment';
  import {
    Navbar,
    NavBrand,
    NavLi,
    NavUl,
    NavHamburger
  } from 'flowbite-svelte';
  import { getCookie, deleteCookie } from 'svelte-cookie';
  import { goto } from '$app/navigation';

  $: authToken = browser ? getCookie('token') : '';

  const handleLogout = () => {
    deleteCookie('token');
    goto('/login');
  };
</script>

<div class="app">
  <Navbar let:hidden let:toggle rounded color="form">
    <NavBrand href="/">
      <span
        class="self-center whitespace-nowrap text-xl font-semibold dark:text-white"
      >
        Connect Chat
      </span>
    </NavBrand>
    <NavHamburger on:click={toggle} />
    <NavUl {hidden}>
      {#if authToken.length === 0}
        <NavLi href="/login">Login</NavLi>
        <NavLi href="/signup">Sign up</NavLi>
      {:else}
        <NavLi on:click={handleLogout}>Logout</NavLi>
      {/if}
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
