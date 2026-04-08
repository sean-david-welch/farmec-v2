(function () {
    function buildPartHtml(index) {
        return `
            <div class="part-entry" data-index="${index}">
                <div class="part-entry-header">
                    <h2 class="main-heading">Part ${index + 1}</h2>
                    <button type="button" class="part-remove-btn" onclick="removeWarrantyPart(this)">
                        <i class="fa-solid fa-trash"></i>
                    </button>
                </div>
                <div class="form-row">
                    <div>
                        <label for="part_number_${index}">Part Number</label>
                        <input type="text" name="part_number_${index}" id="part_number_${index}" placeholder="Enter part number">
                    </div>
                    <div>
                        <label for="quantity_needed_${index}">Quantity Needed</label>
                        <input type="number" name="quantity_needed_${index}" id="quantity_needed_${index}" placeholder="Enter quantity needed">
                    </div>
                </div>
                <div class="form-row">
                    <div>
                        <label for="invoice_number_${index}">Invoice Number</label>
                        <input type="text" name="invoice_number_${index}" id="invoice_number_${index}" placeholder="Enter invoice number">
                    </div>
                    <div>
                        <label for="part_description_${index}">Part Description</label>
                        <input type="text" name="part_description_${index}" id="part_description_${index}" placeholder="Enter description">
                    </div>
                </div>
            </div>`;
    }

    function renumberParts() {
        const entries = document.querySelectorAll('#parts-container .part-entry');
        entries.forEach((entry, i) => {
            entry.dataset.index = i;
            entry.querySelector('h2').textContent = `Part ${i + 1}`;
            entry.querySelector('[name^="part_number_"]').name = `part_number_${i}`;
            entry.querySelector('[name^="part_number_"]').id = `part_number_${i}`;
            entry.querySelector('[name^="quantity_needed_"]').name = `quantity_needed_${i}`;
            entry.querySelector('[name^="quantity_needed_"]').id = `quantity_needed_${i}`;
            entry.querySelector('[name^="invoice_number_"]').name = `invoice_number_${i}`;
            entry.querySelector('[name^="invoice_number_"]').id = `invoice_number_${i}`;
            entry.querySelector('[name^="part_description_"]').name = `part_description_${i}`;
            entry.querySelector('[name^="part_description_"]').id = `part_description_${i}`;
            const removeBtn = entry.querySelector('.part-remove-btn');
            if (removeBtn) removeBtn.style.display = entries.length === 1 ? 'none' : '';
        });
        document.getElementById('part_count').value = entries.length;
    }

    window.addWarrantyPart = function () {
        const container = document.getElementById('parts-container');
        const index = container.querySelectorAll('.part-entry').length;
        container.insertAdjacentHTML('beforeend', buildPartHtml(index));
        renumberParts();
    };

    window.removeWarrantyPart = function (btn) {
        btn.closest('.part-entry').remove();
        renumberParts();
    };

    renumberParts();
})();
