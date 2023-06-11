<script lang="ts">
  import { userClient } from '../../clients';
  import { setCookie } from 'svelte-cookie';
  import type { LoginRequest, LoginResponse } from '../../gen/proto/user_pb';
  import { userStore } from '../../store';
  import { goto } from '$app/navigation';
  import AuthForm from '../../components/AuthForm.svelte';
  import type { ConnectError } from '@bufbuild/connect';
  import { promiseHandler } from '../../utils';

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

    // const res = await userClient.login(loginPayload);
    const [res, err] = await promiseHandler<LoginResponse, ConnectError | null>(
      userClient.login(loginPayload)
    );

    if (err !== null) {
      // some error. maybe user a toaster
    }

    if (res !== null) {
      setCookie('token', res.accessToken, 1, true);
      if (res.user) userStore.setUser(res.user);
      // redirect to home
      goto('/');
    } else {
      // show some failure
    }
  };
</script>

<svelte:head>
  <title>Login</title>
  <meta name="description" content="Login for connect chat" />
</svelte:head>

<AuthForm mode="login" handleSubmit={handleLogin} />
