const inputFields = document.querySelectorAll('.passcode-input input[type="number"]');
const passcodeForm = document.getElementById("passcodeForm");

// Move focus to next input field on input
inputFields.forEach((input, index) => {
    input.addEventListener('input', (event) => {
        if (event.data !== null) {
            input.value = event.data;
            // Submit form on last input
            let filledOut = true;
            for (const input of inputFields.values()) {
                if (input.value === '') {
                    filledOut = false;
                }
            }
            if (filledOut) passcodeForm.submit();
            if (index < inputFields.length - 1) {
                inputFields[index + 1].focus();
            }
        }
    });
});

// Move focus to previous input field on backspace/delete
inputFields.forEach((input, index) => {
    input.addEventListener('keydown', (event) => {
        if (event.key === 'Backspace' || event.key === 'Delete') {
            if (input.value === '') {
                if (index > 0) {
                    inputFields[index - 1].focus();
                }
            }
        } else if (event.key === 'ArrowLeft') {
            if (index > 0) {
                inputFields[index - 1].focus();
            }
        } else if (event.key === 'ArrowRight') {
            if (index < inputFields.length - 1) {
                inputFields[index + 1].focus();
            }
        }
    });
});