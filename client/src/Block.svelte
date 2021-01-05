<script>
  import { onMount } from "svelte";
  import Button from "./Button.svelte";

  let isCollapsed = false;

  export let header = "";
  export let body = "";
  export let ref = null;
  export let remove = () => {};
  export let onBlur = () => {};
  export let onKeyUp = () => {};

  function toggle() {
    isCollapsed = !isCollapsed;
  }

  onMount(() => {
    ref.focus();
  });
</script>

<div class="flex flex-col">
  <div class="flex justify-between">
    <input
      on:blur={onBlur}
      on:keyup={onKeyUp}
      class="flex-1 placeholder-gray-400 bg-transparent text-xl
      focus:outline-none"
      type="text"
      placeholder="heading"
      bind:value={header} />
    <div class="button-bar">
      <Button onClick={toggle}>C</Button>
      <Button onClick={remove}>R</Button>
    </div>
  </div>

  {#if !isCollapsed}
    <div class="flex-1">
      <p
        on:blur={onBlur}
        on:keyup={onKeyUp}
        contenteditable
        bind:this={ref}
        placeholder="write..."
        class="bg-transparent placeholder-gray-400 focus:outline-none
        resize-none w-full mt-0">
        {body}
      </p>
    </div>
  {/if}

</div>
