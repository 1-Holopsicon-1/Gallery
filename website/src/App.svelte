<script>
    import SearchBar from "./SearchBar.svelte";
    import Image from "./Image.svelte";
    import ImagesOnPage from "./ImagesOnPage.svelte";
</script>

<main>
    <SearchBar/>
    <Image/>
    <ImagesOnPage/>
</main>

<style>
    :global(body) {
        background-color: black;
        color: white;
    }
    :global(ul) {
        list-style-type: none;
    }
</style>