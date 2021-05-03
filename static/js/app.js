// base.page
function Pronpt() {
    let toast = function(c) {
        console.log(c)
        const {
            title = "",
            icon = "success",
            position = "top-end"
        } = c;
        console.log(title)

        const Toast = Swal.mixin({
            toast: true,
            title,
            position,
            icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })

        Toast.fire({
            // icon: 'success',
            // title: 'Signed in successfully'
        })



    }
    let success = function(c) {
        const {
            title = "",
            text = "",
            icon = "success",
            position = "top-end",
            footer = "",
        } = c;
        Swal.fire({
            title,
            icon,
            text,
            footer,
        })
    }
    let error = function(c) {
        const {
            title = "",
            text = "",
            icon = "error",
            position = "top-end",
            footer = "",
        } = c;
        Swal.fire({
            title,
            icon,
            text,
            footer,
        })
    }

    async function custom(c) {
        const {
            icon="",
            html = "",
            title = "",
            showConfirmButton = true,
        } = c;

        const { value: result } = await Swal.fire({
            icon,
            title,
            html,
                // '<input id="swal-input1" class="swal2-input">' +
                // '<input id="swal-input2" class="swal2-input">',
            backdrop: true,
            focusConfirm: false,
            showCancelButton: true,
            showConfirmButton,
            willOpen: () => {
                if (c.willOpen !== undefined) 
                c.willOpen();
            },
            preConfirm: () => {
                return [
                    document.getElementById('start_date').value,
                    document.getElementById('end_date').value
                ]
            },
            didOpen: () => {
                if (c.didOpen !== undefined)
                    c.didOpen();
            },
        })

        if (result) {
            if (result.dismiss !== Swal.DismissReason.cancel){
                if (result.value !== "") {
                    if (c.callback !== undefined){
                        c.callback(result);

                    }
                } else {
                    c.callback(false);
                }

            } else {
                console.log("cancel")
                c.callback(false);
            }
        }
    }

    return {
        toast,
        success,
        error,
        custom,
    }
}

// generals.page majors.page
function CheckAvailability(roomID, csrfToken) {
    document.getElementById("check-availability-button").addEventListener("click",function(){
        // notify("This is my message", "warning")
        // notifyModal("title", "<em>hello, world</em","success", "My Text for the button")
        // attention.success({title: "Hello, world",footer:"footer"}); // msg => title
        let title = "Choose your dates"
        let html =`
        <form 
            id="check-availability-form"
            action=""
            method="POST"
            novalidate
            class="needs-validation">
            
            <div class="row">
                <div class="col">
                    <div class="row" id="reservation-dates-modal">
                        <div class="col">
                            <div class="mb-3">
                                <input disabled required type="text" class="form-control" id="start_date" name="start_date" placeholder="Arrival">
                            </div>
                        </div>
                        <div class="col">
                            <div class="mb-3">
                                <input disabled required type="text" class="form-control" id="end_date" name="end_date" placeholder="Departure">
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </form>
        `
        attention.custom({
            title,
            html,
            willOpen: () => {
                const elem = document.getElementById("reservation-dates-modal");
                const rangepicker = new DateRangePicker(elem, {
                    format: "yyyy-mm-dd",
                    showOnFocus: true,
                    minDate: new Date(),
                });

            },
            didOpen: () => {
                document.getElementById('start_date').removeAttribute('disabled');
                document.getElementById('end_date').removeAttribute('disabled');
            },
            callback: function(result) {

                let form = document.getElementById("check-availability-form");
                let formData = new FormData(form);
                formData.append("csrf_token", csrfToken);
                formData.append("room_id", roomID);

                fetch('/search-availability-json', {
                    method: "post",
                    body: formData,
                })
                    .then(response => response.json())
                    .then(data => {
                        if (data.ok) {
                            attention.custom({
                                icon: 'success',
                                showConfirmButton: false,
                                html:  '<p>Room is available!</p>'
                                    + '<p><a href="/book-room'
                                    + '?id='
                                    + data.room_id
                                    + '&s='
                                    + data.start_date
                                    + '&e='
                                    + data.end_date
                                    + '" class="btn btn-primary">'
                                    + 'Book now!</a></p>',
                            })
                        } else {
                            attention.error({
                                title: "No availability",
                        });
                    }
                })

            },
        });
    })
}
