<script>
    import {goto} from "$app/navigation"
    import LabNav from "../../../../components/labNav.svelte";


    let firstName = "";
    let lastName = "";
    let surname = "";
    let address = "";
    let contact = "";
    let files = []

    function registerLabProfile(){ 
        var formData = new FormData()
        formData.append("firstName", firstName)
        formData.append("lastName", lastName)
        formData.append("surname", surname)
        formData.append("address", address)
        formData.append("contact", contact)
        formData.append("profileImage", files[0])
        fetch("http://localhost:8000/lab_profile_register", {
            method:"POST",
            headers:{
                "Accept" :"*/*",
                Authorization : `Bearer ${localStorage.getItem('hospital_token')}`
            },
            body:formData
        })
        .then(res=> res.json())
        .then(data=> {
            console.log(data)            
            goto("/lab_home")
        })
        .catch(err => console.log(err))
    }     
</script>

<LabNav />
<form on:submit|preventDefault={registerLabProfile}>
    <div class="gap-y-7 flex flex-col items-center justify-center min-h-screen">
        <div class="flex flex-col justify-center pb-5 px-10 gap-y-5 min-w-[30%] shadow-lg shadow-gray-400 rounded-lg">                         
            <p class="flex items-center justify-center pt-5 border-b-2 border-rose-300 text-lg text-gray-700">
                Lab Admin Profile
            </p>
            <div class="flex space-x-3 items-center">
                <label for="firstName">First Name: *</label>
                <input type="text" name="firstName" id="firstName" required placeholder="First Name" 
                    bind:value={firstName} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
            </div>
            <div class="flex space-x-3 items-center">
                <label for="surname">Surname: *</label>
                <input type="text" name="surname" id="surname" required placeholder="Surname"
                     bind:value={surname} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
            </div>            
            <div class="flex space-x-3 items-center">
                <label for="lastName">Last Name:</label>
                <input type="text" name="lastName" id="lastName" placeholder="Last Name" 
                bind:value={lastName} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
            </div>
            <div class="flex space-x-3 items-center">
                <label for="contact">Contact: *</label>
                <input type="contact" name="contact" id="contact" required  
                    bind:value={contact} placeholder="Contact" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
            </div>
            <div class="flex space-x-3 items-center">
                <label for="address">Address: *</label>
                <input type="address" name="address" id="address" required  
                    bind:value={address} placeholder="Address" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-rose-600"  />                                
            </div>
            <div class="flex flex-col">
                <label for="image">*Profile Image:</label>                
                <input type="file" accept="image/png, image/jpg, image/jpeg" required 
                    name="image1" id="image1" bind:files />
            </div>
            <button type="submit" class="bg-rose-600 p-3 rounded-lg text-gray-200">Create Lab Profile</button>
         </div>
    </div>
</form>