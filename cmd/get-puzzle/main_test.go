package main

import (
	"strings"
	"testing"
)

func TestFindPreCode(t *testing.T) {
	got := strings.TrimSpace(findPreCode(htmlExample))
	if got != strings.TrimSpace(want) {
		t.Errorf("findPreCode:\n%v\nwant:\n%v", got, want)
	}
}

var want = `
less than: < greater than: >
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg


acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf

be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |
fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |
fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |
cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |
efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |
gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |
gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |
cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |
ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |
gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |
fgae cfgab fg bagce


acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf

 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc
`

var htmlExample = `<!DOCTYPE html>
<html lang="en-us">
...
<p>Each digit of a seven-segment display is rendered by turning on or off any of seven segments named <code>a</code> through <code>g</code>:</p>
<pre><code>less than: &lt; greater than: &gt;
  0:      1:      2:      3:      4:
 <em>aaaa</em>    ....    <em>aaaa    aaaa</em>    ....
<em>b    c</em>  .    <em>c</em>  .    <em>c</em>  .    <em>c  b    c</em>
<em>b    c</em>  .    <em>c</em>  .    <em>c</em>  .    <em>c  b    c</em>
 ....    ....    <em>dddd    dddd    dddd</em>
<em>e    f</em>  .    <em>f  e</em>    .  .    <em>f</em>  .    <em>f</em>
<em>e    f</em>  .    <em>f  e</em>    .  .    <em>f</em>  .    <em>f</em>
 <em>gggg</em>    ....    <em>gggg    gggg</em>    ....

  5:      6:      7:      8:      9:
 <em>aaaa    aaaa    aaaa    aaaa    aaaa</em>
<em>b</em>    .  <em>b</em>    .  .    <em>c  b    c  b    c</em>
<em>b</em>    .  <em>b</em>    .  .    <em>c  b    c  b    c</em>
 <em>dddd    dddd</em>    ....    <em>dddd    dddd</em>
.    <em>f  e    f</em>  .    <em>f  e    f</em>  .    <em>f</em>
.    <em>f  e    f</em>  .    <em>f  e    f</em>  .    <em>f</em>
 <em>gggg    gggg</em>    ....    <em>gggg    gggg</em>
</code></pre>
<p>So, to render a <code>1</code>, only segments <code>c</code> and <code>f</code> would be turned on; the rest would be off. To render a <code>7</code>, only segments <code>a</code>, <code>c</code>, and <code>f</code> would be turned on.</p>
...
<p>For example, here is what you might see in a single entry in your notes:</p>
<pre><code>acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf</code></pre>
<p>(The entry is wrapped here to two lines so it fits; in your notes, it will all be on a single line.)</p>
...
<p>Using this information, you should be able to work out which combination of signal wires corresponds to each of the ten digits. Then, you can decode the four digit output value. Unfortunately, in the above example, all of the digits in the output value (<code>cdfeb fcadb cdfeb cdbaf</code>) use five segments and are more difficult to deduce.</p>
<p>For now, <em>focus on the easy digits</em>. Consider this larger example:</p>
<pre><code>be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |
<em>fdgacbe</em> cefdb cefbgd <em>gcbe</em>
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |
fcgedb <em>cgb</em> <em>dgebacf</em> <em>gc</em>
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |
<em>cg</em> <em>cg</em> fdcagb <em>cbg</em>
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |
efabcd cedba gadfec <em>cb</em>
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |
<em>gecf</em> <em>egdcabf</em> <em>bgf</em> bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |
<em>gebdcfa</em> <em>ecba</em> <em>ca</em> <em>fadegcb</em>
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |
<em>cefg</em> dcbef <em>fcge</em> <em>gbcadfe</em>
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |
<em>ed</em> bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |
<em>gbdfcae</em> <em>bgc</em> <em>cg</em> <em>cgb</em>
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |
<em>fgae</em> cfgab <em>fg</em> bagce
</code></pre>
<p>Because the digits <code>1</code>, <code>4</code>, <code>7</code>, and <code>8</code> each use a unique number of segments, you should be able to tell which combinations of signals correspond to those digits. Counting <em>only digits in the output values</em> (the part after <code>|</code> on each line), in the above example, there are <code><em>26</em></code> instances of digits that use a unique number of segments (highlighted above).</p>
<p><em>In the output values, how many times do digits <code>1</code>, <code>4</code>, <code>7</code>, or <code>8</code> appear?</em></p>
</article>
<p>Your puzzle answer was <code>512</code>.</p><article class="day-desc"><h2 id="part2">--- Part Two ---</h2><p>Through a little deduction, you should now be able to determine the remaining digits. Consider again the first example above:</p>
<pre><code>acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf</code></pre>
<p>After some careful analysis, the mapping between signal wires and segments only make sense in the following configuration:</p>
<pre><code> dddd
e    a
e    a
 ffff
g    b
g    b
 cccc
</code></pre>
<p>So, the unique signal patterns would correspond to the following digits:</p>
<ul>
<li><code>acedgfb</code>: <code>8</code></li>
<li><code>cdfbe</code>: <code>5</code></li>
<li><code>gcdfa</code>: <code>2</code></li>
<li><code>fbcad</code>: <code>3</code></li>
<li><code>dab</code>: <code>7</code></li>
<li><code>cefabd</code>: <code>9</code></li>
<li><code>cdfgeb</code>: <code>6</code></li>
<li><code>eafb</code>: <code>4</code></li>
<li><code>cagedb</code>: <code>0</code></li>
<li><code>ab</code>: <code>1</code></li>
</ul>
<p>Then, the four digits of the output value can be decoded:</p>
<ul>
<li><code>cdfeb</code>: <code><em>5</em></code></li>
<li><code>fcadb</code>: <code><em>3</em></code></li>
<li><code>cdfeb</code>: <code><em>5</em></code></li>
<li><code>cdbaf</code>: <code><em>3</em></code></li>
</ul>
<p>Therefore, the output value for this entry is <code><em>5353</em></code>.</p>
<p>Following this same process for each entry in the second, larger example above, the output value of each entry can be determined:</p>
<ul>
<li><code>fdgacbe cefdb cefbgd gcbe</code>: <code>8394</code></li>
<li><code>fcgedb cgb dgebacf gc</code>: <code>9781</code></li>
<li><code>cg cg fdcagb cbg</code>: <code>1197</code></li>
<li><code>efabcd cedba gadfec cb</code>: <code>9361</code></li>
<li><code>gecf egdcabf bgf bfgea</code>: <code>4873</code></li>
<li><code>gebdcfa ecba ca fadegcb</code>: <code>8418</code></li>
<li><code>cefg dcbef fcge gbcadfe</code>: <code>4548</code></li>
<li><code>ed bcgafe cdgba cbgef</code>: <code>1625</code></li>
<li><code>gbdfcae bgc cg cgb</code>: <code>8717</code></li>
<li><code>fgae cfgab fg bagce</code>: <code>4315</code></li>
</ul>
<p>Adding all of the output values in this larger example produces <code><em>61229</em></code>.</p>
<p>For each entry, determine all of the wire/segment connections and decode the four-digit output values. <em>What do you get if you add up all of the output values?</em></p>
</article>
<p>Your puzzle answer was <code>1091165</code>.</p><p class="day-success">Both parts of this puzzle are complete! They provide two gold stars: **</p>
<p>At this point, you should <a href="/2021">return to your Advent calendar</a> and try another puzzle.</p>
<p>If you still want to see it, you can <a href="8/input" target="_blank">get your puzzle input</a>.</p>
...
</body>
</html>`
