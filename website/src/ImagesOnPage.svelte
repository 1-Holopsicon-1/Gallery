<script>
    import Image, {fullSize} from "./Image.svelte";

    import {link} from "./consts.js";
    import SearchBar from "./SearchBar.svelte";

    async function getPosts() {
        const res = await fetch(link + "post/all");
        const data = await res.json();
        if (res.ok) {
            return data;
        } else {
            throw new Error(data);
        }
    }

    let promise = getPosts();

    export function loadAll() {
        promise = getPosts();
    }
</script>

<div on:load={loadAll()} id="imagesList" class="posts">
    {#await promise}
        <img src="../public/loading-gif.gif" alt="loading">
    {:then posts}
        {#each posts as post}
            <div class="post" on:click={fullSize({post})}>
                <p>{post["id"]}</p>
                <img id={post["id"]} src={post["url"]} alt="not png">
                <p>{post["tags"].text}</p>
            </div>
        {/each}
        {:catch error}
        <p style="color: red">{error.message}</p>
    {/await}
</div>
<style>
    .posts {
        display: grid;
        gap: 5rem;
        grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    }

    .post {
        display: grid;
        place-items: center;
        max-width: 1fr;

    }

    img {
        width: 100%;
        max-height: 15rem;
        object-fit: contain;
    }
</style>