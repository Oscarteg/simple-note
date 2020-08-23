<script lang="ts">
  import { slide } from "svelte/transition";
  import Block from "./Block.svelte";
  import Button from "./Button.svelte";

  let blocks = [];

  function addBlock() {
    blocks = [
      ...blocks,
      {
        header: new Date().toLocaleDateString(),
        body: "",
        id: Math.random()
          .toString(36)
          .substr(2, 9)
      }
    ];
  }

  function toggleAll() {}
</script>

<main class="container prose py-3 max-w-64">
  <div>
    <div class="flex justify-between items-center">
      <div>
        <h1>Svelte TODO</h1>
        <p class="text-opacity-75">Saved 42 min ago, 147b</p>
      </div>

      <div class="space-x-3">
        <Button onClick={toggleAll}>CA</Button>
        <Button onClick={addBlock}>A</Button>
      </div>
    </div>
    {#each blocks as block (block.id)}
      <div transition:slide>
        <Block header={block.header} body={block.body} />
      </div>
    {:else}
      <div>Nothing here!</div>
    {/each}
  </div>
</main>
