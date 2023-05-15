<script lang="ts">
  import { browser } from '$app/environment';
  import { getCookie } from 'svelte-cookie';

  const authToken = browser ? getCookie('token') : null;
</script>

<h1>Create Room</h1>
