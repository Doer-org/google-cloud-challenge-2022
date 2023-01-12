type TProps<T> = {
  type: React.HTMLInputTypeAttribute;
  label: string;
  content: string | number;
  changeContent: (content: T) => void;
  required?: boolean;
};
export const Input = <T,>({
  type,
  label,
  content,
  changeContent,
  required,
}: TProps<T>) => {
  return (
    <div className="grid grid-cols-1 my-5">
      <label htmlFor={label} className="py-1">
        {label}
        {required && (
          <span className="w-6 h-6 ml-1 rounded-full text-accent">※</span>
        )}
      </label>
      <input
        type={type}
        className="py-1 px-2 rounded-sm bg-origin border-2 border-white"
        id={label}
        onChange={(e) => changeContent(e.target.value as T)}
        value={content}
        required={required}
      />
    </div>
  );
};
