<script lang="ts">
  import { userClient } from '../../clients';
  import type {
    CreateUserRequest,
    CreateUserResponse
  } from '../../gen/proto/user_pb';
  import { goto } from '$app/navigation';
  import AuthForm from '../../components/AuthForm.svelte';
  import { promiseHandler } from '../../utils';
  import type { ConnectError } from '@bufbuild/connect';

  let signupPayload: Partial<CreateUserRequest> = {};

  const handleSignup = async (e: SubmitEvent) => {
    const formData = new FormData(e.target as HTMLFormElement);
    for (const field of formData) {
      const [key, value] = field;
      const fieldName = key as keyof Pick<
        CreateUserRequest,
        'email' | 'password' | 'username'
      >;
      const fieldValue = value as string;
      signupPayload[fieldName] = fieldValue;
    }
    const [res, err] = await promiseHandler<
      CreateUserResponse,
      ConnectError | null
    >(userClient.createUser(signupPayload));
    if (err) {
      // do something with error
    }
    if (res !== null && res.user) goto('/login');
  };
</script>

<svelte:head>
  <title>Signup</title>
  <meta name="description" content="Sign up for connect chat" />
</svelte:head>

<AuthForm mode="signup" handleSubmit={handleSignup} />
