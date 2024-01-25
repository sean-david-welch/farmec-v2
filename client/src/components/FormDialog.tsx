import utils from '../styles/Utils.module.css';

import { useRef, useEffect } from 'react';

interface Props {
    children: React.ReactNode;
    visible: boolean;
    onClose: () => void;
}

const FormDialog = ({ children, visible, onClose }: Props) => {
    const dialogRef = useRef<HTMLDialogElement | null>(null);

    useEffect(() => {
        if (visible) {
            dialogRef.current?.showModal();
            dialogRef.current?.classList.add(utils.dialogVisible);
        } else {
            dialogRef.current?.close();
            dialogRef.current?.classList.remove(utils.dialogVisible);
        }
    }, [visible]);

    const closeDialog = () => {
        onClose();
    };

    return (
        <dialog ref={dialogRef} className={utils.dialog}>
            {children}
            <button className={utils.dialogBtn} formMethod="dialog" onClick={closeDialog}>
                <img src="/icons/x.svg" alt="x-icon" />
            </button>
        </dialog>
    );
};

export default FormDialog;
