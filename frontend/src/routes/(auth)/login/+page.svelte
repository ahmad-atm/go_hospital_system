<script>
    import {goto} from "$app/navigation"
    var email =  "";
    var password = "";

    function loginUser() { 
        fetch("http://localhost:8000/login", {
            method:"POST",
            headers:({"Content-Type":"application/json"}),
            body:JSON.stringify({
                email,password
            })
        })
        .then(res=> res.json())
        .then(data=> {
            console.log(data)
            localStorage.setItem("hospital_token", data)
            goto("/home")
        })
        .catch(err => console.log(err))
    }    
</script>

<div class="flex items-center justify-center h-screen w-screen bg-gray-50">
    <!-- <div class="w-[50%] flex flex-col space-y-4 items-center justify-center bg-blue-800 border-4 rounded-xl border-white">
        <p class="flex text-white font-semibold text-2xl">Doctor's Appointment Project</p> -->
        <!-- <div class="border-8 py-12 px-20 rounded-full border-white hover:py-12 hover hover:px-20"> -->
            <!-- <p class="text-9xl text-white font-extrabold flex items-center justify-center mb-4">+</p> -->
        <!-- </div> -->
    <!-- </div> -->

    <div class="w-[50%] flex items-center justify-center">
        <form on:submit|preventDefault={loginUser}>
            
            <div class="gap-y-7 flex flex-col bg-white shadow-lg shadow-gray-300 p-10 rounded-2xl">
               <p class="flex justify-center items-center font-semibold text-gray-800">Login Form</p>
                <div class="flex flex-col gap-y-5">
                    <div class="flex space-x-4 items-center px-5">
                        <label for="email" class="text-gray-700">Email:</label>
                        <input type="email" name="email" id="email" bind:value={email} 
                            placeholder="Email" class="p-2 mb-2 border-b border-gray-800 outline-none text-gray-800" 
                            autocomplete="name" required />                                
                    </div>
                    <div class="flex space-x-4 items-center">
                        <label for="password" class="text-gray-700">Password:</label>
                        <input type="password" name="password" id="password" 
                            autocomplete="current-password" bind:value={password} 
                            placeholder="Password" class="p-2 mb-2 border-b border-gray-800 outline-none text-gray-800"
                            required/>
                    </div>
                 </div>
                 <button type="submit" class="bg-blue-800 p-3 rounded-lg text-gray-200 hover:bg-blue-400">Login</button>
                 <a href="/register_patient" class="flex items-center justify-center text-blue-800">Sign Up</a>
            </div>
        </form>
    </div>
</div>



<style>
    
</style>