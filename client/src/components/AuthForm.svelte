<script lang="ts">
  import type { AuthMode } from '../@types';
  import { A, Button, Card, Label, Input } from 'flowbite-svelte';

  export let mode: AuthMode;
  export let handleSubmit = (_: SubmitEvent) => {};
</script>

<div class="auth-container">
  <Card>
    <form
      class="flex flex-col space-y-6"
      on:submit|preventDefault={handleSubmit}
      style="width: 350px"
    >
      <h3 class="text-xl font-medium text-gray-900 dark:text-white text-center">
        Let's Connect!
      </h3>
      {#if mode === 'signup'}
        <Label class="space-y-2">
          <span>Username</span>
          <Input type="text" name="text" placeholder="Username" required />
        </Label>
      {:else if mode === 'login'}
        <Label class="space-y-2">
          <span>Username</span>
          <Input
            type="text"
            name="text"
            placeholder="Username or Email"
            required
          />
        </Label>
      {/if}
      {#if mode === 'signup'}
        <Label class="space-y-2">
          <span>Email</span>
          <Input
            type="email"
            name="email"
            placeholder="someone@someone.com"
            required
          />
        </Label>
      {/if}
      <Label class="space-y-2">
        <span>Password</span>
        <Input
          type="password"
          name="password"
          placeholder="Password"
          required
        />
      </Label>
      {#if mode === 'signup'}
        <Button color="purple">Sign up</Button>
      {:else}
        <Button color="purple">Login</Button>
      {/if}
    </form>
    {#if mode === 'login'}
      <a class="auth-link" href="signup">Don't have an account? Sign up</a>
    {:else}
      <a class="auth-link" href="login">Already have an account? Login</a>
    {/if}
  </Card>
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

  .auth-link {
    margin-top: 15px;
    margin-left: auto;
    margin-right: auto;
  }
</style>
