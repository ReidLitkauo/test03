import subprocess

################################################################################
# Configuration

# List of all static pages
# Please do not include leading or trailing slashes
static_pages = [
	'_index'
]

# Location of all static pages relative to working directory for script
# Needs trailing slash
static_location = 'src/rsc/'

# Location of all compiled pages relative to working directory for script
# Needs trailing slash
compiled_location = 'out/rsc/'

################################################################################
# Compilation functions

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # 
# Template HTML page
def html_template (ghead, gstyle, gscript, gbody, lhead, lstyle, lscript, lbody):
# Fill in a default HTML page with all sorts of goodies.
	return f"""
		<!DOCTYPE html>
		<html>
			<head>
				{ghead}
				{lhead}
				<style>{gstyle}</style>
				<style>{lstyle}</style>
				<script>{gscript}</script>
				<script>{lscript}</script>
			</head><body>
				{gbody}
				{lbody}
			</body>
		</html>
	"""

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # 
# Compile Pug files
def compile_pug (rsc_name):
# Accepts the name of the resource to compile
# Returns a tuple with the head in [0] and body in [1]
# Supports emojis and other UTF8 chars in source files thanks to .decode()

	#   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   
	# The Pug compilation script to run within Node
	def pug_js_script (filename):
		return f"const pug = require('pug'); const compiled = pug.compileFile('{filename}'); console.log(compiled());"
	
	#   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   
	# The command to call Node and capture its input
	def run_node (script):
		return subprocess.check_output(f"node -e \"{script}\"", shell=True).decode('utf8')

	#   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   #   
	# Call Node twice, once for head and once for body
	return ( run_node(pug_js_script(rsc_name + '.head.pug')), run_node(pug_js_script(rsc_name + '.body.pug')) )

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # 
# Compile CoffeeScript files
def compile_coffee (rsc_name):
	return subprocess.check_output(f"coffee -cp {rsc_name}.coffee", shell=True).decode('utf8')

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # 
# Compile Sass files
def compile_sass (rsc_name):
	return subprocess.check_output(f"sass --indented {rsc_name}.sass", shell=True).decode('utf8')

################################################################################
# Run compilation

# Compile global resources
ghead, gbody = compile_pug(static_location)
gscript = compile_coffee(static_location)
gstyle = compile_sass(static_location)

# Generate HTML files based on template
for page in static_pages:

	# Compile local resources
	lhead, lbody = compile_pug(static_location + page + '/')
	lscript = compile_coffee(static_location + page + '/')
	lstyle = compile_sass(static_location + page + '/')

	# Paste it all together and write to appropriate file
	with open(compiled_location + page + '.html', "w") as file:
		file.write(html_template(ghead, gstyle, gscript, gbody, lhead, lstyle, lscript, lbody))

