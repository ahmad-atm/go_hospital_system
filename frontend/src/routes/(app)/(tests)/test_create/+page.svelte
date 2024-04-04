<script>
    import LabNav from "../../../../components/labNav.svelte";
    import {goto} from "$app/navigation"

    let title = ""
    let price = ""
    let description = ""

    function createTest(){
        fetch("http://localhost:8000/create_test", {
            method:"POST",
            headers:{
                "Content-type":"application/json",
                Authorization:`Bearer ${localStorage.getItem("hospital_token")}`
            },
            body:JSON.stringify({
                title, price, description
            })
        })
        .then(res => res.json())
        .then(data => {
                console.log(data);
                goto("/tests")})
        .catch(err => console.log(err))
    }
</script>

<LabNav />
<form on:submit|preventDefault={createTest}>
    <div class="gap-y-7 flex flex-col items-center justify-center py-20">
        <div class="flex flex-col justify-center pb-5 px-10 gap-y-5 min-w-[30%] shadow-lg shadow-gray-400 rounded-lg">                         
            <p class="flex items-center justify-center pt-5 border-b-2 border-rose-300 text-lg text-gray-700">
               Add Test
            </p>
            <div class="flex space-x-3 items-center">
                <label for="title">Test Title: *</label>
                <input type="text" name="title" id="title" required placeholder="Test Title" 
                    bind:value={title} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
            </div>
            <div class="flex space-x-3 items-center">
                <label for="description">Test Description: *</label>
                <input type="text" name="description" id="description" required placeholder="Test Description" 
                    bind:value={description} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
            </div>            
            <div class="flex space-x-3 items-center">
                <label for="price">Price: *</label>
                <input type="text" name="price" id="price" required placeholder="price"
                     bind:value={price} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
            </div>            
            <button type="submit" class="bg-rose-600 p-3 rounded-lg text-gray-200">New Test</button>
         </div>
    </div>
</form>