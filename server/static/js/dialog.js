const initDialog = id => {
    const dialog = document.getElementById(id);
    const closeBtn = dialog.querySelector('.dialog-btn');

    function openDialog(id) {
        document.getElementById(id).showModal();
    }

    function closeDialog(id) {
        document.getElementById(id).close();
    }

    closeBtn.addEventListener('click', () => closeDialog(id));

    dialog.addEventListener('click', (e) => {
        if (e.target === dialog) {
            dialog.close();
        }
    });
};

// Initialize all dialogs on page load
document.addEventListener('DOMContentLoaded', () => {
    const dialogs = document.querySelectorAll('dialog');
    dialogs.forEach(dialog => initDialog(dialog.id));
});