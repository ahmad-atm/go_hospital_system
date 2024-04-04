<script>
    import { onMount } from "svelte";
    import LabNav from "../../../../../components/labNav.svelte";
    import {page} from "$app/stores"
    import {goto} from "$app/navigation"

    let title = "";
    let description = "";
    let price = "";
    let testId = $page.params.testId
    let test = {};
    console.log(testId)

    onMount(()=>{
        fetch(`http://localhost:8000/get_test/${testId}`,{
          headers:{
              "content-type": "application/json",
              Authorization: `Bearer ${localStorage.getItem("hospital_token")}`
          }
      })
      .then(res => res.json())
      .then(data => {test=data; console.log(test)})
      .catch(err => console.log(err))
    })

    function updateTest(){
        fetch(`http://localhost:8000/update_test`,{
            method:"POST",
            headers:{
                  "content-type": "application/json",
                Authorization: `Bearer ${localStorage.getItem("hospital_token")}`              
            },
            body:JSON.stringify({
                ID:testId, 
                title:test.title, 
                description:test.description, 
                price:test.price
            })
      })
      .then(res => res.json())
      .then(data => {test=data; goto("/tests")})
      .catch(err => console.log(err))
    }
</script>

<div>
    <LabNav />

    <form on:submit|preventDefault={updateTest}>
        <div class="gap-y-7 mb-10 flex flex-col items-center justify-center min-h-screen">
            <div class="flex flex-col justify-center pb-5 px-10 gap-y-5 min-w-[30%] shadow-lg shadow-gray-400 rounded-lg">                         
                <p class="flex items-center justify-center pt-5 border-b-2 border-rose-300 text-lg text-gray-700">
                    Test
                </p>
                <div class="flex space-x-3 items-center">
                    <label for="title">Test Name: *</label>
                    <input type="text" name="title" id="title" required placeholder="Test Name" 
                        bind:value={test['title']} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
                </div>
                <div class="flex space-x-3 items-center">
                    <label for="description">Test Description: *</label>
                    <input type="text" name="description" id="description" required placeholder="Test Description"
                         bind:value={test['description']} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
                </div>            
                <div class="flex space-x-3 items-center">
                    <label for="price">Price:</label>
                    <input type="text" name="price" id="price" placeholder="Test Price" 
                    bind:value={test['price']} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
                </div>
                <button type="submit" class="bg-rose-600 p-3 rounded-lg text-gray-200">Update Test</button>
             </div>
        </div>
    </form>    
</div>