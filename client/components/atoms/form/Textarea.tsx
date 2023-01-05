type TProps = {
  label: string;
  content: string;
  changeContent: (content: string) => void;
  required?: boolean;
};
export const Textarea = ({
  label,
  content,
  changeContent,
  required,
}: TProps) => {
  return (
    <div className="grid grid-cols-1 my-5">
      <label htmlFor={label} className="py-1">
        {label}
        {required && (
          <span className="w-6 h-6 ml-1 rounded-full text-accent">※</span>
        )}
      </label>
      <textarea
        id={label}
        className="py-1 px-2 rounded-sm max-h-32 bg-origin border-2 border-white"
        onChange={(e) => changeContent(e.target.value)}
        value={content}
        required={required}
      />
    </div>
  );
};
