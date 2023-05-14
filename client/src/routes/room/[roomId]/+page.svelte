<script lang="ts">
  import { page } from '$app/stores';
  import { onMount } from 'svelte';
  console.log('age', $page.params);

  onMount(async () => {
    // fetch people in room
  });
</script>

<div>room page</div>
