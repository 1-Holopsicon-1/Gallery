<script>
    import {link} from "./consts.js";

    function autocompl(array) {
        let list = document.createElement('ul');
        for (let i = 0; i < array.length; i++) {
            let item = document.createElement('li');
            item.appendChild(document.createTextNode(array[i]));
            item.addEventListener("click", function (){
                let inputField = document.getElementById("tagInput")
                let tags = inputField.value.trim().split(' ')
                tags.splice(-1)
                tags.push(item.textContent)
                inputField.value = tags.join(' ')
                let resList = document.getElementById('resultList')
                resList.innerHTML = ''
            })
            list.appendChild(item);
        }
        return list;
    }

    async function getTags(tagInput) {
        console.log(JSON.stringify(tagInput) + " tagInput")
        const res = await fetch(link + "tag/name",
            {
                method: "POST",
                body: JSON.stringify(tagInput)
            });
        console.log(res)
        const data = await res.json();
        if (res.ok) {
            return data;
        } else {
            throw new Error(data);
        }
    }

    async function getPostsByTag(tagInput){
        const res = await fetch(link + "tag/" + tagInput);
        const data = await res.json();
        if (res.ok) {
            return data;
        } else {
            throw new Error(data);
        }
    }

    export async function searchTags() {
        let tags = document.getElementById("tagInput").value
        if (tags == "") {
            document.getElementById("resultList").replaceChildren();
            return
        }
        console.log(tags)
        let tagsSend = tags.split(' ').pop()
        console.log(tagsSend)
        let tagsInCome = await getTags(tagsSend)
        let res = Object.values(tagsInCome)
        let resList = document.getElementById('resultList')
        resList.innerHTML = ''
        resList.appendChild(autocompl(res))
    }
    function filterImages(){
        let tags = document.getElementById("tagInput").value.split(' ')
        let tagsArray = tags.join('+')
        return getPostsByTag(tagsArray)
    }
    export function filterImagesByEnter(e){
        let promise = filterImages()

    }
    export function filterImagesByButton(e){
        let promise = filterImages()
    }
</script>

<form method="post" on:submit|preventDefault={filterImagesByEnter} class="">
    <input type="text" id="tagInput" class="tagInput" on:input={(e) => searchTags()}>
    <button class="sendButton" type="submit" >Send</button>
    <div id="resultList" class="resultList"></div>
</form>

