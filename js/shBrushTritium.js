/**
 * SyntaxHighlighter
 * http://alexgorbatchev.com/SyntaxHighlighter
 *
 * SyntaxHighlighter is donationware. If you are using it, please donate.
 * http://alexgorbatchev.com/SyntaxHighlighter/donate.html
 *
 * @version
 * 3.0.83 (July 02 2010)
 * 
 * @copyright
 * Copyright (C) 2004-2010 Alex Gorbatchev.
 *
 * @license
 * Dual licensed under the MIT and GPL licenses.
 */
(function()
{
	// CommonJS
	typeof(require) != 'undefined' ? SyntaxHighlighter = require('shCore').SyntaxHighlighter : null;

	function Brush()
	{
		// Contributed by Kevin Liou

		this.regexList = [
			{ regex: SyntaxHighlighter.regexLib.singleLinePerlComments,	css: 'comments' },	// one line comments
			{ regex: SyntaxHighlighter.regexLib.singleLineCComments,		css: 'comments' },	// one line comments
			{ regex: SyntaxHighlighter.regexLib.multiLineCComments, 		css: 'comments' },	// multi line comments
			{ regex: /"([^\\"]|\\.)*"/mg,																css: 'string' },		// double quoted strings
			{ regex: /'([^\\']|\\.)*'/mg,																css: 'string' },		// single quoted strings
			{ regex: /\/([^\\\/\n]|\\.)*\/\B/g,														css: 'regex' },			// regular expressions
			{ regex: /[\w-:]*:/g,																				css: 'symbol' },		// attribute symbols
			{ regex: /[$%](?=\w)/g,																			css: 'variable' },	// variables
			{ regex: /@\w+/g,																						css: 'keyword' }		// keywords
			];
	};

	Brush.prototype	= new SyntaxHighlighter.Highlighter();
	Brush.aliases	= ['tritium', 'ts'];

	SyntaxHighlighter.brushes.Tritium = Brush;

	// CommonJS
	typeof(exports) != 'undefined' ? exports.Brush = Brush : null;
})();
