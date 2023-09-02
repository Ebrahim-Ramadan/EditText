'use client'
import React, { useState, useEffect } from 'react';

const Tiptap = () => {
  const [editableContent, setEditableContent] = useState('<strong>Hello World!</strong>');

  const handleContentChange = (e) => {
    setEditableContent(e.target.innerHTML);
  };

  useEffect(() => {
    const handleKeyPress = (e) => {
      // Check if the user pressed Alt+D
      if (e.altKey && e.key === 'g') {
        // Enable editing
        document.querySelector('.editable-content').setAttribute('contentEditable', 'true');
      }
    };

    const handleBlur = () => {
      // Disable editing when the element loses focus
      document.querySelector('.editable-content').setAttribute('contentEditable', 'false');
    };

    // Add event listeners
    document.addEventListener('keydown', handleKeyPress);
    document.querySelector('.editable-content').addEventListener('blur', handleBlur);

    return () => {
      // Remove event listeners when the component unmounts
      document.removeEventListener('keydown', handleKeyPress);
      document.querySelector('.editable-content').removeEventListener('blur', handleBlur);
    };
  }, []);

  return (
    <>
      <div
        className="editable-content"
        contentEditable={false} // Initially disabled
        dangerouslySetInnerHTML={{ __html: editableContent }}
        onBlur={handleContentChange}
      />
    </>
  );
}

export default Tiptap;     













