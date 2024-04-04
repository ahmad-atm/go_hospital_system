<script>
    import {goto} from "$app/navigation"
    import PatNav from "../../../../components/patNav.svelte";
    import PatientTab from "../../../../components/patient_tab.svelte";

    let firstName = "";
    let lastName = "";
    let surname = "";
    let contact = "";
    let address = "";
    let gender = "";
    let dob = "";
    let bloodType = "";
    let files = []

    function patientProfileRegister(){  
        var formData = new FormData();
        formData.append("firstName", firstName);
        formData.append("lastName", lastName);
        formData.append("surname", surname);
        formData.append("contact", contact);
        formData.append("gender", gender);
        formData.append("dob", dob);
        formData.append("address", address);
        formData.append("bloodType", bloodType);
        formData.append("contact", contact);
        formData.append("profileImage",  files[0]);         

        fetch("http://localhost:8000/patient_profile_register", {
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
            goto("/patient_home")
        })
        .catch(err => console.log(err))
    }     
</script>

<div class="bg-gray-200 h-screen grid grid-cols-12">
    <div class="col-span-1">
        <PatNav />
    </div>
    <div class="col-span-11 flex items-center justify-center">
        <form on:submit|preventDefault={patientProfileRegister} enctype="multipart/form-data">
            <div class="gap-y-7 bg-white flex flex-col items-center justify-center min-h-screen">
                <div class="flex flex-col justify-center items-center pb-5 px-10 gap-y-5 w-[50%] rounded-lg">                         
                    <p class="flex items-center justify-center pt-5 border-b-2 border-blue-300 text-lg text-gray-700">
                        Patient Profile
                    </p>   
                    <div class="flex space-x-4">
                        <div class="flex space-x-3 items-center">
                            <label for="firstName" class="text-md text-gray-700">*First Name:</label>
                            <input type="text" name="firstName" id="firstName" required placeholder="First Name" 
                                bind:value={firstName} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-blue-600"  />                               
                        </div>
                        <div class="flex space-x-3 items-center">
                            <label for="surname" class="text-md text-gray-700">*Surname:</label>
                            <input type="text" name="surname" id="surname" required placeholder="Surname" bind:value={surname} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-blue-600"  />
                        </div>                         
                    </div>         

                    <div class="flex space-x-4">
                        <div class="flex space-x-3 items-center">
                            <label for="lastName" class="text-md text-gray-700">Last Name:</label>
                            <input type="text" name="lastName" id="lastName" placeholder="Last Name" bind:value={lastName} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-blue-600"  />                               
                        </div>

                        <div class="flex space-x-3 items-center">
                            <label for="contact" class="text-md text-gray-700">*Contact:</label>
                            <input type="text" name="contact" id="contact" required  bind:value={contact} placeholder="Contact" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-blue-600"  />                       
                        </div>                    
                    </div>
                    <div class="flex space-x-4">
                        <div class="flex space-x-3 items-center">
                            <label for="address" class="text-md text-gray-700">*Address:</label>
                            <input type="text" name="address" id="address" required  bind:value={address} placeholder="Address" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-blue-600"  />                       
                        </div>

                        <div class="flex space-x-3 items-center">
                            <label for="gender" class="text-md text-gray-700">*Gender:</label>
                            <input type="text" name="gender" id="gender" required  bind:value={gender} placeholder="Gender" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-blue-600"  />                       
                        </div>                      
                    </div>
                    <div class="flex space-x-4">
                        <div class="flex space-x-3 items-center">
                            <label for="dob" class="text-md text-gray-700">*Date of Birth:</label>
                            <input type="text" name="dob" id="dob" required  bind:value={dob} placeholder="Date of Birth" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-blue-600"  />                       
                        </div>  
                
                        <div class="flex space-x-3 items-center">
                            <label for="bloodType" class="text-md text-gray-700">*Blood Type:</label>
                            <input type="text" name="bloodType" id="bloodType" required  bind:value={bloodType} placeholder="Blood Type" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-blue-600"  />                       
                        </div>                      
                    </div>
  
            
 

                    <div class="flex flex-col">
                        <label for="image">*Profile Image:</label>                
                        <input type="file" accept="image/png, image/jpg, image/jpeg" required 
                        name="image1" id="image1" bind:files>
                    </div>
                </div>
                <button type="submit" class="bg-blue-800 hover:bg-blue-600 p-3 rounded-lg text-gray-200">Create Patient Profile</button>
            </div>
        </form>
    </div>
    <!-- <div class="col-span-3">
        <PatientTab />
    </div> -->
</div>