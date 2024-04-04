<script>
    import {goto} from "$app/navigation"
    import DocNav from "../../../../components/docNav.svelte"

    let firstName = "";
    let lastName = "";
    let surname = "";
    let address = "";
    let contact = "";
    let dob = "";
    let gender = "";
    let bloodType = "";
    let department = "";
    let specialty = "";
    let medicalLicenceNumber = ""
    let files = []

    function registerDoctorProfile(){ 
        var formData = new FormData()
        formData.append("firstName", firstName)
        formData.append("lastName", lastName)
        formData.append("surname", surname)
        formData.append("address", address)
        formData.append("contact", contact)
        formData.append("dob", dob)
        formData.append("gender", gender)
        formData.append("bloodType", bloodType)
        formData.append("department", department)
        formData.append("specialty", specialty)
        formData.append("medicalLicenceNumber", medicalLicenceNumber)
        formData.append("profileImage", files[0])     

        fetch("http://localhost:8000/doctor_profile_register", {
            method:"POST",
            headers:{
                "accept":"*/*",
                Authorization : `Bearer ${localStorage.getItem('hospital_token')}`
            },
            body:formData
        })
        .then(res=> res.json())
        .then(data=> {
            console.log(data)            
            goto("/doctor_home")
        })
        .catch(err => console.log(err))
    }     
</script>

<div class="bg-gray-200 h-screen grid grid-cols-12">
    <div class="col-span-1">
        <DocNav />
    </div>
    <div class="col-span-11 flex flex-col items-center justify-center">
        <form on:submit|preventDefault={registerDoctorProfile}>
            <div class="gap-y-7 flex flex-col items-center justify-center min-h-screen bg-white">
                <div class="flex flex-col justify-center pb-5 px-10 gap-y-5 min-w-[30%] rounded-lg">                         
                    <p class="flex items-center justify-center pt-5 border-b-2 border-teal-300 text-lg text-gray-700">
                        Doctor Profile
                    </p>   
                    <div class="flex space-x-4">                                          
                        <div class="flex space-x-3 items-center">
                            <label for="firstName">*First Name:</label>
                            <input type="text" name="firstName" id="firstName" required placeholder="First Name" bind:value={firstName} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />                                
                        </div>
                        <div class="flex space-x-3 items-center">
                            <label for="surname">*Surname:</label>
                            <input type="text" name="surname" id="surname" required placeholder="Surname" bind:value={surname} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />
                        </div> 
                    </div>
                    <div class="flex space-x-4">                                
                        <div class="flex space-x-3 items-center">
                            <label for="lastName">Last Name:</label>
                            <input type="text" name="lastName" id="lastName" placeholder="Last Name" bind:value={lastName} class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />                                
                        </div>
                        <div class="flex space-x-3 items-center">
                            <label for="contact">*Contact:</label>
                            <input type="contact" name="contact" id="contact" required  bind:value={contact} placeholder="Contact" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />                        
                        </div>
                    </div>
                    <div class="flex space-x-4">                     
                        <div class="flex space-x-3 items-center">
                            <label for="address">*Address:</label>
                            <input type="address" name="address" id="address" required  bind:value={address} placeholder="Address" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />                        
                        </div>

                        <div class="flex space-x-3 items-center">
                            <label for="dob">Date Of Birth:</label>
                            <input type="dob" name="dob" id="dob" required  bind:value={dob} placeholder="DOB" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />                        
                        </div>
                    </div>
                    <div class="flex space-x-4">                     
                        <div class="flex space-x-3 items-center">
                            <label for="gender">Gender:</label>
                            <input type="gender" name="gender" id="gender" required  bind:value={gender} placeholder="Gender" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />                        
                        </div>
                        <div class="flex space-x-3 items-center">
                            <label for="bloodType">Blood Type:</label>
                            <input type="bloodType" name="bloodType" id="bloodType" required  bind:value={bloodType} placeholder="Blood Type" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />                        
                        </div>           
                    </div>
                    <div class="flex space-x-4">                     
                        <div class="flex space-x-3 items-center">
                            <label for="department">Department:</label>
                            <input type="department" name="department" id="department" required  bind:value={department} placeholder="" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />                        
                        </div>

                        <div class="flex space-x-3 items-center">
                            <label for="specialty">Specialty:</label>
                            <input type="specialty" name="specialty" id="specialty" required  bind:value={specialty} placeholder="Specialty" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />                        
                        </div>
                    </div>
					<div class="flex space-x-4">                     
                        <div class="flex space-x-3 items-center">
                            <label for="medicalLicenceNumber">Medical Licence Number:</label>
                            <input type="medicalLicenceNumber" name="medicalLicenceNumber" id="medicalLicenceNumber" required  bind:value={medicalLicenceNumber} placeholder="" class="p-2 mb-2 border-b border-gray-600 outline-none focus:border-teal-600"  />                        
                        </div>    

                        <div class="flex flex-col">
                            <label for="image">*Profile Image:</label>                
                            <input type="file" accept="image/png, image/jpg, image/jpeg" required 
                                name="image1" id="image1" bind:files />
                        </div>
                    </div>
                    <button type="submit" class="bg-teal-700 hover:bg-teal-600 p-3 rounded-md text-gray-200">Create Doctor Profile</button>

                </div>
            </div>
        </form>
    </div>
</div>