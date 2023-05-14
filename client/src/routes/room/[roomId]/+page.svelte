<script lang="ts">
  import { page } from '$app/stores';
  import { onMount } from 'svelte';
  import { getCookie } from 'svelte-cookie';
  import { roomClient } from '../../../clients';

  onMount(async () => {
    try {
      const token = getCookie('token');
      // fetch people in room
      const res = await roomClient.getRoomMembers(
        { roomId: $page.params.roomId },
        {
          headers: {
            Authorization: `Bearer ${token}`
          }
        }
      );
      console.log('res', res);
    } catch (err) {
      // do something with error
      console.log(err);
    }
  });
</script>

<div>room page</div>
