import { onMounted, onBeforeUnmount } from 'vue'

export const useClickOutside = (elementRef, callback) => {
    if (!elementRef) return;

    let listener = (e) => {
        //Click inside elementRef
        if (e.target == elementRef.value || e.composedPath().includes(elementRef.value)) {
            return;
        }
        //Click outside
        if (typeof callback == 'function') {
            callback();
        }
    };

    onMounted(() => {
        window.addEventListener('click', listener);
    });
    onBeforeUnmount(() => {
        window.removeEventListener('event', listener);
    });

    return {
        listener,
    };
};