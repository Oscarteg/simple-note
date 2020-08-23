<script>
  import Button from "./Button.svelte";

  let isCollapsed = false;

  export let header;
  export let body;

  function toggle() {
    isCollapsed = !isCollapsed;
  }
</script>

<div class="flex flex-col">
  <div class="flex justify-between">
    <input
      class="flex-1 placeholder-gray-400 bg-transparent text-xl
      focus:outline-none"
      type="text"
      placeholder="heading"
      bind:value={header} />
    <div class="button-bar">
      <Button onClick={toggle}>C</Button>
      <Button>R</Button>
    </div>
  </div>

  {#if !isCollapsed}
    <div class="flex-1">
      <textarea
        placeholder="write..."
        bind:value={body}
        class="bg-transparent placeholder-gray-400 focus:outline-none
        resize-none w-full" />
    </div>
  {/if}

</div>
