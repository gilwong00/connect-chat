<script lang="ts">
  import type { AuthMode } from '../@types';

  export let mode: AuthMode;
  export let handleSubmit = (_: SubmitEvent) => {};

  $: height = mode === 'login' ? '160px' : '220px';
</script>

<div class="auth-container">
  <h3>Let's Connect!</h3>
  <form
    class="auth-form-container"
    style="height: {height}"
    on:submit|preventDefault={handleSubmit}
  >
    {#if mode === 'signup'}
      <input name="username" type="text" placeholder="Username" />
    {:else if mode === 'login'}
      <input name="username" type="text" placeholder="Username or Email" />
    {/if}
    {#if mode === 'signup'}
      <input name="email" type="text" placeholder="Email" />
    {/if}
    <input name="password" type="password" placeholder="Password" />
    {#if mode === 'signup'}
      <button>Sign up</button>
    {:else}
      <button>auth</button>
    {/if}
  </form>
  {#if mode === 'login'}
    <a class="auth-link" href="signup">Don't have an account? Sign up</a>
  {:else}
    <a class="auth-link" href="login">Already have an account? auth</a>
  {/if}
</div>

<style lang="scss">
  .auth-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: auto 0;

    @media (max-width: 500px) {
      margin: initial;
    }
  }

  .auth-form-container {
    display: flex;
    flex-direction: column;
    justify-content: space-between;

    > input {
      padding: 10px;
      width: 350px;
      border: 1px solid lightgrey;
      border-radius: 4px;
    }

    > button {
      padding: 12px 10px;
      background: purple;
      color: #fff;
      border-radius: 4px;
      border: none;
      cursor: pointer;
    }
  }

  .auth-link {
    margin-top: 10px;
  }
</style>
