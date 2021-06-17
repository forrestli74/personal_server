declare module "*.ejs" {
	const value: (locals: any) => string;
	export = value;
}