<script lang="ts">
  import { userClient } from '../../clients';
  import type { LoginRequest } from '../../gen/proto/user_pb';

  type LoginFormFields = 'username' | 'email' | 'password';

  let errors: Record<LoginFormFields, string> = {};
  let loginPayload: Partial<LoginRequest> = {};

  const handleLogin = async (e: SubmitEvent) => {
    const formData = new FormData(e.target as HTMLFormElement);
    for (const field of formData) {
      const [key, value] = field;
      const fieldName = key as string;
      const fieldValue = value as string;

      if (fieldName === 'username') {
        if (fieldValue.includes('@')) {
          loginPayload.email = fieldValue;
        } else {
          loginPayload.username = fieldValue;
        }
      }
      if (fieldName === 'password') {
        loginPayload.password = fieldValue;
      }
    }

    const res = await userClient.login(loginPayload);
    console.log('res', res);
  };
</script>

<svelte:head>
  <title>Login</title>
  <meta name="description" content="Login for connect chat" />
</svelte:head>

<div class="login-container">
  <h3>Let's Connect!</h3>
  <form class="login-form-container" on:submit|preventDefault={handleLogin}>
    <input name="username" type="text" placeholder="Username or Email" />
    <input name="password" type="password" placeholder="Password" />
    <button>Login</button>
  </form>
  <a class="signup-link" href="signup">Don't have an account? Sign up</a>
</div>

<style lang="scss">
  .login-container {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: auto 0;

    @media (max-width: 500px) {
      margin: initial;
    }
  }

  .login-form-container {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 180px;

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

  .signup-link {
    margin-top: 10px;
  }
</style>
