// Dafny program main.dfy compiled into JavaScript
// Copyright by the contributors to the Dafny Project
// SPDX-License-Identifier: MIT

const BigNumber = require('bignumber.js');
BigNumber.config({ MODULO_MODE: BigNumber.EUCLID })
let _dafny = (function() {
  let $module = {};
  $module.areEqual = function(a, b) {
    if (typeof a === 'string' && b instanceof _dafny.Seq) {
      // Seq.equals(string) works as expected,
      // and the catch-all else block handles that direction.
      // But the opposite direction doesn't work; handle it here.
      return b.equals(a);
    } else if (typeof a === 'number' && BigNumber.isBigNumber(b)) {
      // This conditional would be correct even without the `typeof a` part,
      // but in most cases it's probably faster to short-circuit on a `typeof`
      // than to call `isBigNumber`. (But it remains to properly test this.)
      return b.isEqualTo(a);
    } else if (typeof a !== 'object' || a === null || b === null) {
      return a === b;
    } else if (BigNumber.isBigNumber(a)) {
      return a.isEqualTo(b);
    } else if (a._tname !== undefined || (Array.isArray(a) && a.constructor.name == "Array")) {
      return a === b;  // pointer equality
    } else {
      return a.equals(b);  // value-type equality
    }
  }
  $module.toString = function(a) {
    if (a === null) {
      return "null";
    } else if (typeof a === "number") {
      return a.toFixed();
    } else if (BigNumber.isBigNumber(a)) {
      return a.toFixed();
    } else if (a._tname !== undefined) {
      return a._tname;
    } else {
      return a.toString();
    }
  }
  $module.escapeCharacter = function(cp) {
    let s = String.fromCodePoint(cp.value)
    switch (s) {
      case '\n': return "\\n";
      case '\r': return "\\r";
      case '\t': return "\\t";
      case '\0': return "\\0";
      case '\'': return "\\'";
      case '\"': return "\\\"";
      case '\\': return "\\\\";
      default: return s;
    };
  }
  $module.NewObject = function() {
    return { _tname: "object" };
  }
  $module.InstanceOfTrait = function(obj, trait) {
    return obj._parentTraits !== undefined && obj._parentTraits().includes(trait);
  }
  $module.Rtd_bool = class {
    static get Default() { return false; }
  }
  $module.Rtd_char = class {
    static get Default() { return 'D'; }  // See CharType.DefaultValue in Dafny source code
  }
  $module.Rtd_codepoint = class {
    static get Default() { return new _dafny.CodePoint('D'.codePointAt(0)); }
  }
  $module.Rtd_int = class {
    static get Default() { return BigNumber(0); }
  }
  $module.Rtd_number = class {
    static get Default() { return 0; }
  }
  $module.Rtd_ref = class {
    static get Default() { return null; }
  }
  $module.Rtd_array = class {
    static get Default() { return []; }
  }
  $module.ZERO = new BigNumber(0);
  $module.ONE = new BigNumber(1);
  $module.NUMBER_LIMIT = new BigNumber(0x20).multipliedBy(0x1000000000000);  // 2^53
  $module.Tuple = class Tuple extends Array {
    constructor(...elems) {
      super(...elems);
    }
    toString() {
      return "(" + arrayElementsToString(this) + ")";
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static Default(...values) {
      return Tuple.of(...values);
    }
    static Rtd(...rtdArgs) {
      return {
        Default: Tuple.from(rtdArgs, rtd => rtd.Default)
      };
    }
  }
  $module.Set = class Set extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Set.Empty;
    }
    toString() {
      return "{" + arrayElementsToString(this) + "}";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Set();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new Set();
      for (let k of elmts) {
        s.add(k);
      }
      return s;
    }
    contains(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i], k)) {
          return true;
        }
      }
      return false;
    }
    add(k) {  // mutates the Set; use only during construction
      if (!this.contains(k)) {
        this.push(k);
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        if (!other.contains(e)) {
          return false;
        }
      }
      return true;
    }
    get Elements() {
      return this;
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = Set.of(...this);
        for (let k of that) {
          s.add(k);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new Set();
        for (let k of this) {
          if (that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length == 0 || that.length == 0) {
        return this;
      } else {
        let s = new Set();
        for (let k of this) {
          if (!that.contains(k)) {
            s.push(k);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      for (let k of this) {
        if (that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsSubsetOf(that) {
      if (that.length < this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      if (that.length <= this.length) {
        return false;
      }
      for (let k of this) {
        if (!that.contains(k)) {
          return false;
        }
      }
      return true;
    }
    get AllSubsets() {
      return this.AllSubsets_();
    }
    *AllSubsets_() {
      // Start by putting all set elements into a list, but don't include null
      let elmts = Array.of(...this);
      let n = elmts.length;
      let which = new Array(n);
      which.fill(false);
      let a = [];
      while (true) {
        yield Set.of(...a);
        // "add 1" to "which", as if doing a carry chain.  For every digit changed, change the membership of the corresponding element in "a".
        let i = 0;
        for (; i < n && which[i]; i++) {
          which[i] = false;
          // remove elmts[i] from a
          for (let j = 0; j < a.length; j++) {
            if (_dafny.areEqual(a[j], elmts[i])) {
              // move the last element of a into slot j
              a[j] = a[-1];
              a.pop();
              break;
            }
          }
        }
        if (i === n) {
          // we have cycled through all the subsets
          break;
        }
        which[i] = true;
        a.push(elmts[i]);
      }
    }
  }
  $module.MultiSet = class MultiSet extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return MultiSet.Empty;
    }
    toString() {
      let s = "multiset{";
      let sep = "";
      for (let e of this) {
        let [k, n] = e;
        let ks = _dafny.toString(k);
        while (!n.isZero()) {
          n = n.minus(1);
          s += sep + ks;
          sep = ", ";
        }
      }
      s += "}";
      return s;
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new MultiSet();
      }
      return this._empty;
    }
    static fromElements(...elmts) {
      let s = new MultiSet();
      for (let e of elmts) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    static FromArray(arr) {
      let s = new MultiSet();
      for (let e of arr) {
        s.add(e, _dafny.ONE);
      }
      return s;
    }
    cardinality() {
      let c = _dafny.ZERO;
      for (let e of this) {
        let [k, n] = e;
        c = c.plus(n);
      }
      return c;
    }
    clone() {
      let s = new MultiSet();
      for (let e of this) {
        let [k, n] = e;
        s.push([k, n]);  // make sure to create a new array [k, n] here
      }
      return s;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return _dafny.ZERO;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return !this.get(k).isZero();
    }
    add(k, n) {
      let i = this.findIndex(k);
      if (i === this.length) {
        this.push([k, n]);
      } else {
        let m = this[i][1];
        this[i] = [k, m.plus(n)];
      }
    }
    update(k, n) {
      let i = this.findIndex(k);
      if (i < this.length && this[i][1].isEqualTo(n)) {
        return this;
      } else if (i === this.length && n.isZero()) {
        return this;
      } else if (i === this.length) {
        let m = this.slice();
        m.push([k, n]);
        return m;
      } else {
        let m = this.slice();
        m[i] = [k, n];
        return m;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      for (let e of this) {
        let [k, n] = e;
        let m = other.get(k);
        if (!n.isEqualTo(m)) {
          return false;
        }
      }
      return this.cardinality().isEqualTo(other.cardinality());
    }
    get Elements() {
      return this.Elements_();
    }
    *Elements_() {
      for (let i = 0; i < this.length; i++) {
        let [k, n] = this[i];
        while (!n.isZero()) {
          yield k;
          n = n.minus(1);
        }
      }
    }
    get UniqueElements() {
      return this.UniqueElements_();
    }
    *UniqueElements_() {
      for (let e of this) {
        let [k, n] = e;
        if (!n.isZero()) {
          yield k;
        }
      }
    }
    Union(that) {
      if (this.length === 0) {
        return that;
      } else if (that.length === 0) {
        return this;
      } else {
        let s = this.clone();
        for (let e of that) {
          let [k, n] = e;
          s.add(k, n);
        }
        return s;
      }
    }
    Intersect(that) {
      if (this.length === 0) {
        return this;
      } else if (that.length === 0) {
        return that;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let m = that.get(k);
          if (!m.isZero()) {
            s.push([k, m.isLessThan(n) ? m : n]);
          }
        }
        return s;
      }
    }
    Difference(that) {
      if (this.length === 0 || that.length === 0) {
        return this;
      } else {
        let s = new MultiSet();
        for (let e of this) {
          let [k, n] = e;
          let d = n.minus(that.get(k));
          if (d.isGreaterThan(0)) {
            s.push([k, d]);
          }
        }
        return s;
      }
    }
    IsDisjointFrom(that) {
      let intersection = this.Intersect(that);
      return intersection.cardinality().isZero();
    }
    IsSubsetOf(that) {
      for (let e of this) {
        let [k, n] = e;
        let m = that.get(k);
        if (!n.isLessThanOrEqualTo(m)) {
          return false;
        }
      }
      return true;
    }
    IsProperSubsetOf(that) {
      return this.IsSubsetOf(that) && this.cardinality().isLessThan(that.cardinality());
    }
  }
  $module.CodePoint = class CodePoint {
    constructor(value) {
      this.value = value
    }
    equals(other) {
      if (this === other) {
        return true;
      }
      return this.value === other.value
    }
    isLessThan(other) {
      return this.value < other.value
    }
    isLessThanOrEqual(other) {
      return this.value <= other.value
    }
    toString() {
      return "'" + $module.escapeCharacter(this) + "'";
    }
    static isCodePoint(i) {
      return (
        (_dafny.ZERO.isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0xD800))) ||
        (new BigNumber(0xE000).isLessThanOrEqualTo(i) && i.isLessThan(new BigNumber(0x11_0000))))
    }
  }
  $module.Seq = class Seq extends Array {
    constructor(...elems) {
      super(...elems);
    }
    static get Default() {
      return Seq.of();
    }
    static Create(n, init) {
      return Seq.from({length: n}, (_, i) => init(new BigNumber(i)));
    }
    static UnicodeFromString(s) {
      return new Seq(...([...s].map(c => new _dafny.CodePoint(c.codePointAt(0)))))
    }
    toString() {
      return "[" + arrayElementsToString(this) + "]";
    }
    toVerbatimString(asLiteral) {
      if (asLiteral) {
        return '"' + this.map(c => _dafny.escapeCharacter(c)).join("") + '"';
      } else {
        return this.map(c => String.fromCodePoint(c.value)).join("");
      }
    }
    static update(s, i, v) {
      if (typeof s === "string") {
        let p = s.slice(0, i);
        let q = s.slice(i.toNumber() + 1);
        return p.concat(v, q);
      } else {
        let t = s.slice();
        t[i] = v;
        return t;
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let i = 0; i < this.length; i++) {
        if (!_dafny.areEqual(this[i], other[i])) {
          return false;
        }
      }
      return true;
    }
    static contains(s, k) {
      if (typeof s === "string") {
        return s.includes(k);
      } else {
        for (let x of s) {
          if (_dafny.areEqual(x, k)) {
            return true;
          }
        }
        return false;
      }
    }
    get Elements() {
      return this;
    }
    get UniqueElements() {
      return _dafny.Set.fromElements(...this);
    }
    static Concat(a, b) {
      if (typeof a === "string" || typeof b === "string") {
        // string concatenation, so make sure both operands are strings before concatenating
        if (typeof a !== "string") {
          // a must be a Seq
          a = a.join("");
        }
        if (typeof b !== "string") {
          // b must be a Seq
          b = b.join("");
        }
        return a + b;
      } else {
        // ordinary concatenation
        let r = Seq.of(...a);
        r.push(...b);
        return r;
      }
    }
    static JoinIfPossible(x) {
      try { return x.join(""); } catch(_error) { return x; }
    }
    static IsPrefixOf(a, b) {
      if (b.length < a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
    static IsProperPrefixOf(a, b) {
      if (b.length <= a.length) {
        return false;
      }
      for (let i = 0; i < a.length; i++) {
        if (!_dafny.areEqual(a[i], b[i])) {
          return false;
        }
      }
      return true;
    }
  }
  $module.Map = class Map extends Array {
    constructor() {
      super();
    }
    static get Default() {
      return Map.of();
    }
    toString() {
      return "map[" + this.map(maplet => _dafny.toString(maplet[0]) + " := " + _dafny.toString(maplet[1])).join(", ") + "]";
    }
    static get Empty() {
      if (this._empty === undefined) {
        this._empty = new Map();
      }
      return this._empty;
    }
    findIndex(k) {
      for (let i = 0; i < this.length; i++) {
        if (_dafny.areEqual(this[i][0], k)) {
          return i;
        }
      }
      return this.length;
    }
    get(k) {
      let i = this.findIndex(k);
      if (i === this.length) {
        return undefined;
      } else {
        return this[i][1];
      }
    }
    contains(k) {
      return this.findIndex(k) < this.length;
    }
    update(k, v) {
      let m = this.slice();
      m.updateUnsafe(k, v);
      return m;
    }
    // Similar to update, but make the modification in-place.
    // Meant to be used in the map constructor.
    updateUnsafe(k, v) {
      let m = this;
      let i = m.findIndex(k);
      m[i] = [k, v];
      return m;
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.length !== other.length) {
        return false;
      }
      for (let e of this) {
        let [k, v] = e;
        let w = other.get(k);
        if (w === undefined || !_dafny.areEqual(v, w)) {
          return false;
        }
      }
      return true;
    }
    get Keys() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(k);
      }
      return s;
    }
    get Values() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.add(v);
      }
      return s;
    }
    get Items() {
      let s = new _dafny.Set();
      for (let e of this) {
        let [k, v] = e;
        s.push(_dafny.Tuple.of(k, v));
      }
      return s;
    }
    Merge(that) {
      let m = that.slice();
      for (let e of this) {
        let [k, v] = e;
        let i = m.findIndex(k);
        if (i == m.length) {
          m[i] = [k, v];
        }
      }
      return m;
    }
    Subtract(keys) {
      if (this.length === 0 || keys.length === 0) {
        return this;
      }
      let m = new Map();
      for (let e of this) {
        let [k, v] = e;
        if (!keys.contains(k)) {
          m[m.length] = e;
        }
      }
      return m;
    }
  }
  $module.newArray = function(initValue, ...dims) {
    return { dims: dims, elmts: buildArray(initValue, ...dims) };
  }
  $module.BigOrdinal = class BigOrdinal {
    static get Default() {
      return _dafny.ZERO;
    }
    static IsLimit(ord) {
      return ord.isZero();
    }
    static IsSucc(ord) {
      return ord.isGreaterThan(0);
    }
    static Offset(ord) {
      return ord;
    }
    static IsNat(ord) {
      return true;  // at run time, every ORDINAL is a natural number
    }
  }
  $module.BigRational = class BigRational {
    static get ZERO() {
      if (this._zero === undefined) {
        this._zero = new BigRational(_dafny.ZERO);
      }
      return this._zero;
    }
    constructor (n, d) {
      // requires d === undefined || 1 <= d
      this.num = n;
      this.den = d === undefined ? _dafny.ONE : d;
      // invariant 1 <= den || (num == 0 && den == 0)
    }
    static get Default() {
      return _dafny.BigRational.ZERO;
    }
    // We need to deal with the special case `num == 0 && den == 0`, because
    // that's what C#'s default struct constructor will produce for BigRational. :(
    // To deal with it, we ignore `den` when `num` is 0.
    toString() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num.toFixed() + ".0";
      }
      let answer = this.dividesAPowerOf10(this.den);
      if (answer !== undefined) {
        let n = this.num.multipliedBy(answer[0]);
        let log10 = answer[1];
        let sign, digits;
        if (this.num.isLessThan(0)) {
          sign = "-"; digits = n.negated().toFixed();
        } else {
          sign = ""; digits = n.toFixed();
        }
        if (log10 < digits.length) {
          let digitCount = digits.length - log10;
          return sign + digits.slice(0, digitCount) + "." + digits.slice(digitCount);
        } else {
          return sign + "0." + "0".repeat(log10 - digits.length) + digits;
        }
      } else {
        return "(" + this.num.toFixed() + ".0 / " + this.den.toFixed() + ".0)";
      }
    }
    isPowerOf10(x) {
      if (x.isZero()) {
        return undefined;
      }
      let log10 = 0;
      while (true) {  // invariant: x != 0 && x * 10^log10 == old(x)
        if (x.isEqualTo(1)) {
          return log10;
        } else if (x.mod(10).isZero()) {
          log10++;
          x = x.dividedToIntegerBy(10);
        } else {
          return undefined;
        }
      }
    }
    dividesAPowerOf10(i) {
      let factor = _dafny.ONE;
      let log10 = 0;
      if (i.isLessThanOrEqualTo(_dafny.ZERO)) {
        return undefined;
      }

      // invariant: 1 <= i && i * 10^log10 == factor * old(i)
      while (i.mod(10).isZero()) {
        i = i.dividedToIntegerBy(10);
       log10++;
      }

      while (i.mod(5).isZero()) {
        i = i.dividedToIntegerBy(5);
        factor = factor.multipliedBy(2);
        log10++;
      }
      while (i.mod(2).isZero()) {
        i = i.dividedToIntegerBy(2);
        factor = factor.multipliedBy(5);
        log10++;
      }

      if (i.isEqualTo(_dafny.ONE)) {
        return [factor, log10];
      } else {
        return undefined;
      }
    }
    toBigNumber() {
      if (this.num.isZero() || this.den.isEqualTo(1)) {
        return this.num;
      } else if (this.num.isGreaterThan(0)) {
        return this.num.dividedToIntegerBy(this.den);
      } else {
        return this.num.minus(this.den).plus(1).dividedToIntegerBy(this.den);
      }
    }
    isInteger() {
      return this.equals(new _dafny.BigRational(this.toBigNumber(), _dafny.ONE));
    }
    // Returns values such that aa/dd == a and bb/dd == b.
    normalize(b) {
      let a = this;
      let aa, bb, dd;
      if (a.num.isZero()) {
        aa = a.num;
        bb = b.num;
        dd = b.den;
      } else if (b.num.isZero()) {
        aa = a.num;
        dd = a.den;
        bb = b.num;
      } else {
        let gcd = BigNumberGcd(a.den, b.den);
        let xx = a.den.dividedToIntegerBy(gcd);
        let yy = b.den.dividedToIntegerBy(gcd);
        // We now have a == a.num / (xx * gcd) and b == b.num / (yy * gcd).
        aa = a.num.multipliedBy(yy);
        bb = b.num.multipliedBy(xx);
        dd = a.den.multipliedBy(yy);
      }
      return [aa, bb, dd];
    }
    compareTo(that) {
      // simple things first
      let asign = this.num.isZero() ? 0 : this.num.isLessThan(0) ? -1 : 1;
      let bsign = that.num.isZero() ? 0 : that.num.isLessThan(0) ? -1 : 1;
      if (asign < 0 && 0 <= bsign) {
        return -1;
      } else if (asign <= 0 && 0 < bsign) {
        return -1;
      } else if (bsign < 0 && 0 <= asign) {
        return 1;
      } else if (bsign <= 0 && 0 < asign) {
        return 1;
      }
      let [aa, bb, dd] = this.normalize(that);
      if (aa.isLessThan(bb)) {
        return -1;
      } else if (aa.isEqualTo(bb)){
        return 0;
      } else {
        return 1;
      }
    }
    equals(that) {
      return this.compareTo(that) === 0;
    }
    isLessThan(that) {
      return this.compareTo(that) < 0;
    }
    isAtMost(that) {
      return this.compareTo(that) <= 0;
    }
    plus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.plus(bb), dd);
    }
    minus(b) {
      let [aa, bb, dd] = this.normalize(b);
      return new BigRational(aa.minus(bb), dd);
    }
    negated() {
      return new BigRational(this.num.negated(), this.den);
    }
    multipliedBy(b) {
      return new BigRational(this.num.multipliedBy(b.num), this.den.multipliedBy(b.den));
    }
    dividedBy(b) {
      let a = this;
      // Compute the reciprocal of b
      let bReciprocal;
      if (b.num.isGreaterThan(0)) {
        bReciprocal = new BigRational(b.den, b.num);
      } else {
        // this is the case b.num < 0
        bReciprocal = new BigRational(b.den.negated(), b.num.negated());
      }
      return a.multipliedBy(bReciprocal);
    }
  }
  $module.EuclideanDivisionNumber = function(a, b) {
    if (0 <= a) {
      if (0 <= b) {
        // +a +b: a/b
        return Math.floor(a / b);
      } else {
        // +a -b: -(a/(-b))
        return -Math.floor(a / -b);
      }
    } else {
      if (0 <= b) {
        // -a +b: -((-a-1)/b) - 1
        return -Math.floor((-a-1) / b) - 1;
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return Math.floor((-a-1) / -b) + 1;
      }
    }
  }
  $module.EuclideanDivision = function(a, b) {
    if (a.isGreaterThanOrEqualTo(0)) {
      if (b.isGreaterThanOrEqualTo(0)) {
        // +a +b: a/b
        return a.dividedToIntegerBy(b);
      } else {
        // +a -b: -(a/(-b))
        return a.dividedToIntegerBy(b.negated()).negated();
      }
    } else {
      if (b.isGreaterThanOrEqualTo(0)) {
        // -a +b: -((-a-1)/b) - 1
        return a.negated().minus(1).dividedToIntegerBy(b).negated().minus(1);
      } else {
        // -a -b: ((-a-1)/(-b)) + 1
        return a.negated().minus(1).dividedToIntegerBy(b.negated()).plus(1);
      }
    }
  }
  $module.EuclideanModuloNumber = function(a, b) {
    let bp = Math.abs(b);
    if (0 <= a) {
      // +a: a % bp
      return a % bp;
    } else {
      // c = ((-a) % bp)
      // -a: bp - c if c > 0
      // -a: 0 if c == 0
      let c = (-a) % bp;
      return c === 0 ? c : bp - c;
    }
  }
  $module.ShiftLeft = function(b, n) {
    return b.multipliedBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.ShiftRight = function(b, n) {
    return b.dividedToIntegerBy(new BigNumber(2).exponentiatedBy(n));
  }
  $module.RotateLeft = function(b, n, w) {  // truncate(b << n) | (b >> (w - n))
    let x = _dafny.ShiftLeft(b, n).mod(new BigNumber(2).exponentiatedBy(w));
    let y = _dafny.ShiftRight(b, w - n);
    return x.plus(y);
  }
  $module.RotateRight = function(b, n, w) {  // (b >> n) | truncate(b << (w - n))
    let x = _dafny.ShiftRight(b, n);
    let y = _dafny.ShiftLeft(b, w - n).mod(new BigNumber(2).exponentiatedBy(w));;
    return x.plus(y);
  }
  $module.BitwiseAnd = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 & b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    return r;
  }
  $module.BitwiseOr = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 | b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseXor = function(a, b) {
    let r = _dafny.ZERO;
    const m = _dafny.NUMBER_LIMIT;  // 2^53
    let h = _dafny.ONE;
    while (!a.isZero() && !b.isZero()) {
      let a0 = a.mod(m);
      let b0 = b.mod(m);
      r = r.plus(h.multipliedBy(a0 ^ b0));
      a = a.dividedToIntegerBy(m);
      b = b.dividedToIntegerBy(m);
      h = h.multipliedBy(m);
    }
    r = r.plus(h.multipliedBy(a | b));
    return r;
  }
  $module.BitwiseNot = function(a, bits) {
    let r = _dafny.ZERO;
    let h = _dafny.ONE;
    for (let i = 0; i < bits; i++) {
      let bit = a.mod(2);
      if (bit.isZero()) {
        r = r.plus(h);
      }
      a = a.dividedToIntegerBy(2);
      h = h.multipliedBy(2);
    }
    return r;
  }
  $module.Quantifier = function(vals, frall, pred) {
    for (let u of vals) {
      if (pred(u) !== frall) { return !frall; }
    }
    return frall;
  }
  $module.PlusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) + b.charCodeAt(0));
  }
  $module.UnicodePlusChar = function(a, b) {
    return new _dafny.CodePoint(a.value + b.value);
  }
  $module.MinusChar = function(a, b) {
    return String.fromCharCode(a.charCodeAt(0) - b.charCodeAt(0));
  }
  $module.UnicodeMinusChar = function(a, b) {
    return new _dafny.CodePoint(a.value - b.value);
  }
  $module.AllBooleans = function*() {
    yield false;
    yield true;
  }
  $module.AllChars = function*() {
    for (let i = 0; i < 0x10000; i++) {
      yield String.fromCharCode(i);
    }
  }
  $module.AllUnicodeChars = function*() {
    for (let i = 0; i < 0xD800; i++) {
      yield new _dafny.CodePoint(i);
    }
    for (let i = 0xE0000; i < 0x110000; i++) {
      yield new _dafny.CodePoint(i);
    }
  }
  $module.AllIntegers = function*() {
    yield _dafny.ZERO;
    for (let j = _dafny.ONE;; j = j.plus(1)) {
      yield j;
      yield j.negated();
    }
  }
  $module.IntegerRange = function*(lo, hi) {
    if (lo === null) {
      while (true) {
        hi = hi.minus(1);
        yield hi;
      }
    } else if (hi === null) {
      while (true) {
        yield lo;
        lo = lo.plus(1);
      }
    } else {
      while (lo.isLessThan(hi)) {
        yield lo;
        lo = lo.plus(1);
      }
    }
  }
  $module.SingleValue = function*(v) {
    yield v;
  }
  $module.HaltException = class HaltException extends Error {
    constructor(message) {
      super(message)
    }
  }
  $module.HandleHaltExceptions = function(f) {
    try {
      f()
    } catch (e) {
      if (e instanceof _dafny.HaltException) {
        process.stdout.write("[Program halted] " + e.message + "\n")
        process.exitCode = 1
      } else {
        throw e
      }
    }
  }
  $module.FromMainArguments = function(args) {
    var a = [...args];
    a.splice(0, 2, args[0] + " " + args[1]);
    return a;
  }
  $module.UnicodeFromMainArguments = function(args) {
    return $module.FromMainArguments(args).map(_dafny.Seq.UnicodeFromString);
  }
  return $module;

  // What follows are routines private to the Dafny runtime
  function buildArray(initValue, ...dims) {
    if (dims.length === 0) {
      return initValue;
    } else {
      let a = Array(dims[0].toNumber());
      let b = Array.from(a, (x) => buildArray(initValue, ...dims.slice(1)));
      return b;
    }
  }
  function arrayElementsToString(a) {
    // like `a.join(", ")`, but calling _dafny.toString(x) on every element x instead of x.toString()
    let s = "";
    let sep = "";
    for (let x of a) {
      s += sep + _dafny.toString(x);
      sep = ", ";
    }
    return s;
  }
  function BigNumberGcd(a, b){  // gcd of two non-negative BigNumber's
    while (true) {
      if (a.isZero()) {
        return b;
      } else if (b.isZero()) {
        return a;
      }
      if (a.isLessThan(b)) {
        b = b.modulo(a);
      } else {
        a = a.modulo(b);
      }
    }
  }
})();
// Dafny program systemModulePopulator.dfy compiled into JavaScript
let _System = (function() {
  let $module = {};

  $module.nat = class nat {
    constructor () {
    }
    static get Default() {
      return _dafny.ZERO;
    }
    static _Is(__source) {
      let _0_x = (__source);
      return (_dafny.ZERO).isLessThanOrEqualTo(_0_x);
    }
  };

  return $module;
})(); // end of module _System
let _module = (function() {
  let $module = {};

  $module.__default = class __default {
    constructor () {
      this._tname = "_module._default";
    }
    _parentTraits() {
      return [];
    }
    static abs(x) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return (new BigNumber(-1)).multipliedBy(x);
      } else {
        return x;
      }
    };
    static safeIndex(x, length) {
      if ((x).isLessThan(_dafny.ZERO)) {
        return _dafny.ZERO;
      } else if ((length).isLessThanOrEqualTo(x)) {
        return (x).mod(length);
      } else {
        return x;
      }
    };
    static safeDivisionInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return _dafny.EuclideanDivision(x1, x2);
      }
    };
    static safeModuloInt(x1, x2) {
      if ((x2).isEqualTo(_dafny.ZERO)) {
        return x1;
      } else {
        return (x1).mod(x2);
      }
    };
    static fm1(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge((_dafny.Map.Empty.slice().updateUnsafe((_module.D12.create_DC29(true, new BigNumber(620), false, new BigNumber((_dafny.Seq.of(false)).length), false)).dtor_cf52,!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true)));
    };
    static fm2(p0, p1, globalState) {
      return ((_module.D9.create_DC22(new BigNumber((function () {
  let _coll0 = new _dafny.Map();
  for (const _compr_0 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(192)), function (_0_i0) {
    return new _dafny.CodePoint('d'.codePointAt(0));
  }),new BigNumber(-91))).Keys.Elements) {
    let _1_v0 = _compr_0;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(192)), function (_0_i0) {
      return new _dafny.CodePoint('d'.codePointAt(0));
    }),new BigNumber(-91))).contains(_1_v0)) {
      _coll0.push([_1_v0,new BigNumber(865)]);
    }
  }
  return _coll0;
}()).length))).dtor_cf38).isLessThanOrEqualTo(new BigNumber(103));
    };
    static fm3(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(true)),new _dafny.CodePoint('y'.codePointAt(0)))).Merge(function () {
        let _coll1 = new _dafny.Map();
        for (const _compr_1 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(895)), function (_2_i0) {
          return _dafny.Seq.of(false);
        })).Elements) {
          let _3_v0 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(895)), function (_2_i0) {
            return _dafny.Seq.of(false);
          }), _3_v0)) {
            _coll1.push([_3_v0,new _dafny.CodePoint('e'.codePointAt(0))]);
          }
        }
        return _coll1;
      }())).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true),new _dafny.CodePoint('m'.codePointAt(0))));
    };
    static fm4(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(271),!_dafny.areEqual(_dafny.Seq.UnicodeFromString("s"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-254)), function (_4_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      })));
    };
    static fm10(p0, globalState) {
      return (_dafny.Set.fromElements(new BigNumber(810))).Intersect(((true) ? (function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of _dafny.IntegerRange(new BigNumber(192), new BigNumber(644))) {
          let _5_v0 = _compr_2;
          if (((new BigNumber(192)).isLessThanOrEqualTo(_5_v0)) && ((_5_v0).isLessThan(new BigNumber(644)))) {
            _coll2.add((_5_v0).minus(new BigNumber(-269)));
          }
        }
        return _coll2;
      }()) : (_dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(-526)))));
    };
    static fm13(p0, p1, globalState) {
      return _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, true)).length)),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(716)), function (_6_i0) {
        return new _dafny.CodePoint('d'.codePointAt(0));
      })).length))).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(38)), function (_7_i1) {
        return new _dafny.CodePoint('b'.codePointAt(0));
      })).length),new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(904)), function (_8_i2) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length);
      })).length))).length)));
    };
    static fm16(p0, p1, globalState) {
      return _dafny.Seq.UnicodeFromString("l");
    };
    static fm17(globalState) {
      return !(!(!(((true) ? (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-780)), function (_9_i0) {
        return _dafny.Seq.of(true);
      }), _dafny.Seq.of(_dafny.Seq.of(false, false)))) : ((_dafny.Set.fromElements(false)).IsProperSubsetOf(_dafny.Set.fromElements(false)))))));
    };
    static fm18(globalState) {
      return (_dafny.MultiSet.fromElements(_dafny.Set.fromElements(!(!(true))))).Intersect((_dafny.MultiSet.fromElements(_dafny.Set.fromElements(true), _dafny.Set.fromElements(false, true))).Difference((_module.D26.create_DC67(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(false, !(false), true), _dafny.Set.fromElements(true, false)))).dtor_cf111));
    };
    static fm19(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(false, true, false, false, !(false)), _dafny.Seq.Concat(_dafny.Seq.of(true, false, false, false, false), _dafny.Seq.of(false, false)));
    };
    static fm20(p0, globalState) {
      return (new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(291))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(!(true), true))).length)))).length)).plus((_dafny.ZERO).minus(new BigNumber(((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("qeutyx"), _dafny.Seq.UnicodeFromString("dpui"))).Difference((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("ts"), _dafny.Seq.UnicodeFromString("jkwix"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-889)), function (_10_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      }))))).cardinality())));
    };
    static fm21(globalState) {
      if ((false) === (true)) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(817),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(746),false));
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(494),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("lno")).length),true));
      }
    };
    static fm22(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), function (_11_i0) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length), new BigNumber(905)),new BigNumber(-703));
    };
    static fm23(p0, globalState) {
      if (!(new BigNumber((_dafny.Seq.UnicodeFromString("hlbbvk")).length)).isEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),false)).length)))) {
        return _dafny.Seq.Concat(_dafny.Seq.of(false, !(true), false, !(!(true)), true), _dafny.Seq.of(!(true)));
      } else if (true) {
        return _dafny.Seq.of(true, !(false), false);
      } else {
        return _dafny.Seq.of(true);
      }
    };
    static fm24(p0, globalState) {
      return _dafny.MultiSet.fromElements(new BigNumber(-72), new BigNumber(-329));
    };
    static fm25(p0, p1, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("vy"), _dafny.Seq.UnicodeFromString("y")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("mbgtm"), _dafny.Seq.UnicodeFromString("pacqqiw")));
    };
    static fm26(p0, p1, p2, p3, globalState) {
      return (_module.D11.create_DC24(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("le"),new _dafny.CodePoint('s'.codePointAt(0))))).dtor_cf40;
    };
    static fm29(p0, globalState) {
      if (false) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('t'.codePointAt(0));
      }
    };
    static fm30(p0, p1, p2, globalState) {
      return true;
    };
    static fm31(p0, p1, p2, globalState) {
      return _dafny.Seq.UnicodeFromString("roi");
    };
    static fm32(p0, globalState) {
      return new BigNumber((function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of _dafny.IntegerRange(new BigNumber(89), new BigNumber(359))) {
            let _12_v1 = _compr_4;
            if (((new BigNumber(89)).isLessThanOrEqualTo(_12_v1)) && ((_12_v1).isLessThan(new BigNumber(359)))) {
              _coll4.push([(_12_v1).minus(new BigNumber(952)),false]);
            }
          }
          return _coll4;
        }()).length),new BigNumber((_dafny.Seq.of(false)).length))).Keys.Elements) {
          let _13_v0 = _compr_3;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((function () {
            let _coll5 = new _dafny.Map();
            for (const _compr_5 of _dafny.IntegerRange(new BigNumber(89), new BigNumber(359))) {
              let _12_v1 = _compr_5;
              if (((new BigNumber(89)).isLessThanOrEqualTo(_12_v1)) && ((_12_v1).isLessThan(new BigNumber(359)))) {
                _coll5.push([(_12_v1).minus(new BigNumber(952)),false]);
              }
            }
            return _coll5;
          }()).length),new BigNumber((_dafny.Seq.of(false)).length))).contains(_13_v0)) {
            _coll3.push([(_13_v0).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("rqy")).length))),new BigNumber(326)]);
          }
        }
        return _coll3;
      }()).length);
    };
    static fm33(p0, p1, globalState) {
      return _module.D6.create_DC14(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("f"), _dafny.Seq.UnicodeFromString("csnarsqlx")));
    };
    static fm34(p0, p1, p2, p3, globalState) {
      if (!((new BigNumber((_dafny.Seq.of(new BigNumber(550), new BigNumber(706))).length)).isLessThan(new BigNumber((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('i'.codePointAt(0)))).length))).length)))) {
        return _module.D12.create_DC28(true, new BigNumber(-395));
      } else {
        return _module.D12.create_DC28(!(false), new BigNumber((_dafny.Seq.of(true)).length));
      }
    };
    static fm35(p0, p1, globalState) {
      return _dafny.Seq.of(new BigNumber(325));
    };
    static fm36(p0, p1, p2, p3, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber(-65))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new BigNumber((_dafny.Seq.UnicodeFromString("xaxw")).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(!(false)),new BigNumber((_dafny.Seq.UnicodeFromString("utuis")).length))).length)));
    };
    static fm37(p0, globalState) {
      return _module.D7.create_DC19(new BigNumber(236));
    };
    static fm38(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(!(false)), _dafny.Seq.of(true, true)), _dafny.Seq.of(true, false));
    };
    static fm39(p0, globalState) {
      return ((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of _dafny.IntegerRange(new BigNumber(888), new BigNumber(779))) {
          let _14_v0 = _compr_6;
          if (((new BigNumber(888)).isLessThanOrEqualTo(_14_v0)) && ((_14_v0).isLessThan(new BigNumber(779)))) {
            _coll6.push([_module.__default.safeModuloInt(_14_v0, new BigNumber(778)),!(!(false))]);
          }
        }
        return _coll6;
      }()).Merge((_module.D28.create_DC71(function () {
  let _coll7 = new _dafny.Map();
  for (const _compr_7 of (_dafny.Seq.of(new BigNumber(376))).Elements) {
    let _15_v1 = _compr_7;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber(376)), _15_v1)) {
      _coll7.push([(_15_v1).minus(new BigNumber((_dafny.Seq.of(true)).length)),false]);
    }
  }
  return _coll7;
}())).dtor_cf116)).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(884),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality())),false)));
    };
    static fm40(p0, p1, p2, p3, globalState) {
      if (true) {
        return _module.D12.create_DC27(_dafny.Set.fromElements(true));
      } else {
        return _module.D12.create_DC27(_dafny.Set.fromElements(false, true));
      }
    };
    static fm43(p0, p1, p2, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(997),!(false))).length)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("vwotwnwen")).length))),_dafny.Seq.UnicodeFromString("amno"));
    };
    static fm44(p0, globalState) {
      return new BigNumber(862);
    };
    static fm45(p0, p1, p2, globalState) {
      return _module.D14.create_DC35((new BigNumber(656)).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true, true, true)).length))).length)), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(65)), function (_16_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
}), _dafny.Seq.Create(_module.__default.abs(new BigNumber(584)), function (_17_i1) {
  return new _dafny.CodePoint('j'.codePointAt(0));
})), new BigNumber((_dafny.MultiSet.fromElements(true, !(true))).cardinality()), (_dafny.MultiSet.fromElements(!(true))).IsSubsetOf(_dafny.MultiSet.fromElements(!(true))), (new BigNumber((_dafny.Seq.UnicodeFromString("edklwnr")).length)).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.of(false, true)).length)));
    };
    static fm46(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("x"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(493)), function (_18_i0) {
        return new _dafny.CodePoint('m'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("sdfbv"));
    };
    static fm47(p0, p1, globalState) {
      if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0)))).cardinality())), new BigNumber(-789))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(false,!(true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,false));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(!(false),false);
      }
    };
    static fm48(p0, p1, p2, p3, globalState) {
      return (_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), function (_19_i0) {
        return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(false), false, false),new BigNumber(-383))).length);
      }))).Union((_dafny.MultiSet.fromElements(new BigNumber(817))).Intersect(_dafny.MultiSet.fromElements(new BigNumber((function () {
        let _coll8 = new _dafny.Map();
        for (const _compr_8 of _dafny.IntegerRange(new BigNumber(958), new BigNumber(817))) {
          let _20_v0 = _compr_8;
          if (((new BigNumber(958)).isLessThanOrEqualTo(_20_v0)) && ((_20_v0).isLessThan(new BigNumber(817)))) {
            _coll8.push([_module.__default.safeDivisionInt(_20_v0, new BigNumber(713)),new BigNumber(855)]);
          }
        }
        return _coll8;
      }()).length), new BigNumber((_dafny.Seq.UnicodeFromString("kjajll")).length))));
    };
    static fm49(p0, globalState) {
      return _module.D0.create_DC1(((false) ? (new BigNumber(814)) : (new BigNumber((_dafny.Seq.of(_dafny.MultiSet.fromElements(false), _dafny.MultiSet.fromElements(!(false)))).length))));
    };
    static fm50(p0, p1, p2, globalState) {
      return function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(970), new BigNumber(345))) {
          let _21_v0 = _compr_9;
          if (((new BigNumber(970)).isLessThanOrEqualTo(_21_v0)) && ((_21_v0).isLessThan(new BigNumber(345)))) {
            _coll9.push([_module.__default.safeModuloInt(_21_v0, new BigNumber(169)),!(false)]);
          }
        }
        return _coll9;
      }();
    };
    static fm51(p0, p1, p2, globalState) {
      return new _dafny.CodePoint('k'.codePointAt(0));
    };
    static fm52(p0, p1, p2, globalState) {
      if (((false) ? (true) : (true))) {
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),(_dafny.ZERO).minus(new BigNumber((function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),(_module.D20.create_DC53(false)).dtor_cf94)).Keys.Elements) {
            let _22_v0 = _compr_10;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),(_module.D20.create_DC53(false)).dtor_cf94)).contains(_22_v0)) {
              _coll10.push([_22_v0,new BigNumber(935)]);
            }
          }
          return _coll10;
        }()).length)));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),new BigNumber((_dafny.Seq.of(new BigNumber(910))).length));
      }
    };
    static fm53(p0, globalState) {
      return _dafny.Set.fromElements((new BigNumber(981)).isLessThanOrEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("souhprl")).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("ywax")).length))).length),false)).length), new BigNumber(426), new BigNumber((_dafny.Seq.UnicodeFromString("liy")).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, false)).length), new BigNumber(-525))).length))),false)).length)), (!(!(false))) === (true), !(!(false) || (true)), true);
    };
    static fm54(p0, p1, globalState) {
      let _source0 = _module.D11.create_DC25(new _dafny.CodePoint('t'.codePointAt(0)), new BigNumber(75), true);
      if (_source0.is_DC25) {
        let _23___mcc_h0 = (_source0).cf41;
        let _24___mcc_h1 = (_source0).cf42;
        let _25___mcc_h2 = (_source0).cf43;
        let _26_cf43 = _25___mcc_h2;
        let _27_cf42 = _24___mcc_h1;
        let _28_cf41 = _23___mcc_h0;
        if (_26_cf43) {
          return _dafny.Map.Empty.slice().updateUnsafe(_28_cf41,_27_cf42);
        } else {
          return function () {
            let _coll11 = new _dafny.Map();
            for (const _compr_11 of (_dafny.Seq.of(_28_cf41)).Elements) {
              let _29_v0 = _compr_11;
              if (_dafny.Seq.contains(_dafny.Seq.of(_28_cf41), _29_v0)) {
                _coll11.push([_29_v0,_27_cf42]);
              }
            }
            return _coll11;
          }();
        }
      } else if (_source0.is_DC24) {
        let _30___mcc_h3 = (_source0).cf40;
        let _31_cf40 = _30___mcc_h3;
        if (false) {
          return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('t'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality()));
        } else {
          return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('q'.codePointAt(0)),new BigNumber(997));
        }
      } else {
        let _32___mcc_h4 = (_source0).cf44;
        let _33_cf44 = _32___mcc_h4;
        return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,!(true))).length));
      }
    };
    static fm55(p0, globalState) {
      return (_dafny.Set.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(false), _dafny.Seq.of(false, false, true, true), _dafny.Seq.of(true, false, true), _dafny.Seq.of(!(true), true, false, false, true))).Difference((_dafny.Set.fromElements(_dafny.Seq.of(false), _dafny.Seq.of(!(false), false))).Union(_dafny.Set.fromElements(_dafny.Seq.of(false, false, false))));
    };
    static fm56(p0, globalState) {
      return ((function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(false)),function () {
          let _coll13 = new _dafny.Set();
          for (const _compr_13 of (_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))).Elements) {
            let _34_v1 = _compr_13;
            if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))), _34_v1)) {
              _coll13.add(_34_v1);
            }
          }
          return _coll13;
        }())).Keys.Elements) {
          let _35_v0 = _compr_12;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(!(false)),function () {
            let _coll14 = new _dafny.Set();
            for (const _compr_14 of (_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)))).Elements) {
              let _36_v1 = _compr_14;
              if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)), new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('y'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0))), _36_v1)) {
                _coll14.add(_36_v1);
              }
            }
            return _coll14;
          }())).contains(_35_v0)) {
            _coll12.push([_35_v0,new BigNumber(421)]);
          }
        }
        return _coll12;
      }()).Merge(function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of (_dafny.Set.fromElements(_dafny.Seq.of(!(false)))).Elements) {
          let _37_v2 = _compr_15;
          if ((_dafny.Set.fromElements(_dafny.Seq.of(!(false)))).contains(_37_v2)) {
            _coll15.push([_37_v2,new BigNumber(753)]);
          }
        }
        return _coll15;
      }())).Merge(function () {
        let _coll16 = new _dafny.Map();
        for (const _compr_16 of (_dafny.Set.fromElements(_dafny.Seq.of(false, false))).Elements) {
          let _38_v3 = _compr_16;
          if ((_dafny.Set.fromElements(_dafny.Seq.of(false, false))).contains(_38_v3)) {
            _coll16.push([_38_v3,new BigNumber(-829)]);
          }
        }
        return _coll16;
      }());
    };
    static fm57(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(!(false),new _dafny.CodePoint('g'.codePointAt(0)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('u'.codePointAt(0))));
    };
    static fm58(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0)));
    };
    static fm59(globalState) {
      return _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("wsdwwpuv"));
    };
    static fm60(p0, p1, globalState) {
      return function () {
        let _coll17 = new _dafny.Map();
        for (const _compr_17 of (((!(true)) ? (_dafny.Set.fromElements(_module.D11.create_DC25(new _dafny.CodePoint('o'.codePointAt(0)), new BigNumber(-847), true), _module.D11.create_DC25(new _dafny.CodePoint('d'.codePointAt(0)), new BigNumber(-798), true))) : (_dafny.Set.fromElements(_module.D11.create_DC25(new _dafny.CodePoint('h'.codePointAt(0)), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), function (_39_i0) {
  return new BigNumber((_dafny.Seq.of(true)).length);
})).length))).length)), true))))).Elements) {
          let _40_v0 = _compr_17;
          if ((((!(true)) ? (_dafny.Set.fromElements(_module.D11.create_DC25(new _dafny.CodePoint('o'.codePointAt(0)), new BigNumber(-847), true), _module.D11.create_DC25(new _dafny.CodePoint('d'.codePointAt(0)), new BigNumber(-798), true))) : (_dafny.Set.fromElements(_module.D11.create_DC25(new _dafny.CodePoint('h'.codePointAt(0)), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(500)), function (_39_i0) {
  return new BigNumber((_dafny.Seq.of(true)).length);
})).length))).length)), true))))).contains(_40_v0)) {
            _coll17.push([_40_v0,new BigNumber(-489)]);
          }
        }
        return _coll17;
      }();
    };
    static fm61(p0, p1, globalState) {
      if (!(true)) {
        return _module.D9.create_DC21(_dafny.MultiSet.fromElements(false));
      } else {
        return _module.D9.create_DC21(_dafny.MultiSet.fromElements(true, true));
      }
    };
    static fm62(p0, p1, p2, p3, globalState) {
      return _module.D15.create_DC39(_dafny.Seq.UnicodeFromString("liockxaf"), _module.D9.create_DC21(_dafny.MultiSet.FromArray(_dafny.Seq.of(false))));
    };
    static fm63(p0, p1, globalState) {
      return ((_dafny.MultiSet.fromElements(true)).Intersect(_dafny.MultiSet.fromElements(!(true), true))).Intersect((_dafny.MultiSet.fromElements(true, !(!(false)))).Union(_dafny.MultiSet.fromElements(true)));
    };
    static fm64(globalState) {
      return _module.D4.create_DC10(_module.__default.safeModuloInt(new BigNumber(499), new BigNumber(664)), ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber(845), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(-856))).length),true)).length))).cardinality()))).length))).minus(new BigNumber(575)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ga"), _dafny.Seq.UnicodeFromString("bywhwnw"))).length));
    };
    static fm65(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(704))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.of(true, true, true)).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("mdjiex")).length))));
    };
    static fm66(p0, p1, globalState) {
      return _module.D16.create_DC42(_module.__default.safeModuloInt(new BigNumber(149), new BigNumber((_dafny.Set.fromElements(new BigNumber(-16))).length)), (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(true, false), _dafny.Seq.of(true, false))).length))), _module.__default.safeModuloInt(new BigNumber(318), new BigNumber(-834)));
    };
    static fm67(p0, globalState) {
      return _module.D11.create_DC25(new _dafny.CodePoint('a'.codePointAt(0)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(691)), function (_41_i0) {
  return new BigNumber(127);
}), _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length), new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber((_dafny.Seq.UnicodeFromString("vhdvmehq")).length)))).length)), true);
    };
    static fm68(p0, p1, globalState) {
      return _module.D14.create_DC34(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-61),_dafny.Seq.UnicodeFromString("vj")));
    };
    static fm69(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(((!((_module.D20.create_DC54(_dafny.Seq.UnicodeFromString("irhb"), new BigNumber((_dafny.Seq.of(!(true))).length), true)).dtor_cf97)) ? ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("wfcecnt")).length))) : (new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("cicqjbm")).length))).length)))),_module.D15.create_DC37(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(484)), function (_42_i0) {
  return new BigNumber(313);
}), _dafny.Seq.of(new BigNumber(-802)))));
    };
    static fm70(p0, globalState) {
      let _source1 = _module.D16.create_DC43(_dafny.Seq.UnicodeFromString("eqqftomtx"), new _dafny.CodePoint('v'.codePointAt(0)), (_dafny.ZERO).minus(new BigNumber((function () {
  let _coll18 = new _dafny.Map();
  for (const _compr_18 of _dafny.IntegerRange(new BigNumber(979), new BigNumber(837))) {
    let _43_v0 = _compr_18;
    if (((new BigNumber(979)).isLessThanOrEqualTo(_43_v0)) && ((_43_v0).isLessThan(new BigNumber(837)))) {
      _coll18.push([(_43_v0).minus(new BigNumber(-243)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('g'.codePointAt(0)))).length))).length)]);
    }
  }
  return _coll18;
}()).length)));
      if (_source1.is_DC42) {
        let _44___mcc_h0 = (_source1).cf76;
        let _45___mcc_h1 = (_source1).cf77;
        let _46___mcc_h2 = (_source1).cf78;
        let _47_cf78 = _46___mcc_h2;
        let _48_cf77 = _45___mcc_h1;
        let _49_cf76 = _44___mcc_h0;
        return function () {
          let _coll19 = new _dafny.Set();
          for (const _compr_19 of (_dafny.Seq.of(function () {
            let _coll20 = new _dafny.Map();
            for (const _compr_20 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), function (_50_i0) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            })).Elements) {
              let _51_v1 = _compr_20;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), function (_50_i0) {
                return new _dafny.CodePoint('d'.codePointAt(0));
              }), _51_v1)) {
                _coll20.push([_51_v1,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-369)), function (_52_i1) {
                  return _dafny.Seq.UnicodeFromString("ttkmplyt");
                })).length)]);
              }
            }
            return _coll20;
          }())).Elements) {
            let _53_v2 = _compr_19;
            if (_dafny.Seq.contains(_dafny.Seq.of(function () {
              let _coll21 = new _dafny.Map();
              for (const _compr_21 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), function (_50_i0) {
                return new _dafny.CodePoint('d'.codePointAt(0));
              })).Elements) {
                let _51_v1 = _compr_21;
                if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), function (_50_i0) {
                  return new _dafny.CodePoint('d'.codePointAt(0));
                }), _51_v1)) {
                  _coll21.push([_51_v1,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-369)), function (_52_i1) {
                    return _dafny.Seq.UnicodeFromString("ttkmplyt");
                  })).length)]);
                }
              }
              return _coll21;
            }()), _53_v2)) {
              _coll19.add(_53_v2);
            }
          }
          return _coll19;
        }();
      } else if (_source1.is_DC43) {
        let _54___mcc_h3 = (_source1).cf79;
        let _55___mcc_h4 = (_source1).cf80;
        let _56___mcc_h5 = (_source1).cf81;
        let _57_cf81 = _56___mcc_h5;
        let _58_cf80 = _55___mcc_h4;
        let _59_cf79 = _54___mcc_h3;
        return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_58_cf80,new BigNumber((_59_cf79).length)))).Intersect(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(_58_cf80,_57_cf81)));
      } else if (_source1.is_DC44) {
        let _60___mcc_h6 = (_source1).cf82;
        let _61_cf82 = _60___mcc_h6;
        return (function () {
          let _coll22 = new _dafny.Set();
          for (const _compr_22 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),_61_cf82), function () {
            let _coll23 = new _dafny.Map();
            for (const _compr_23 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).Elements) {
              let _62_v3 = _compr_23;
              if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).contains(_62_v3)) {
                _coll23.push([_62_v3,_61_cf82]);
              }
            }
            return _coll23;
          }())).Elements) {
            let _63_v4 = _compr_22;
            if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),_61_cf82), function () {
              let _coll24 = new _dafny.Map();
              for (const _compr_24 of (_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).Elements) {
                let _62_v3 = _compr_24;
                if ((_dafny.MultiSet.fromElements(new _dafny.CodePoint('u'.codePointAt(0)))).contains(_62_v3)) {
                  _coll24.push([_62_v3,_61_cf82]);
                }
              }
              return _coll24;
            }())).contains(_63_v4)) {
              _coll22.add(_63_v4);
            }
          }
          return _coll22;
        }()).Intersect(function () {
          let _coll25 = new _dafny.Set();
          for (const _compr_25 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),new BigNumber(815)))).Elements) {
            let _64_v5 = _compr_25;
            if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('a'.codePointAt(0)),new BigNumber(815)))).contains(_64_v5)) {
              _coll25.add(_64_v5);
            }
          }
          return _coll25;
        }());
      } else {
        let _65___mcc_h7 = (_source1).cf75;
        let _66_cf75 = _65___mcc_h7;
        return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('k'.codePointAt(0)),new BigNumber((_dafny.Seq.UnicodeFromString("imekcx")).length)), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),new BigNumber(-246)), function () {
          let _coll26 = new _dafny.Map();
          for (const _compr_26 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),false)).Keys.Elements) {
            let _67_v6 = _compr_26;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),false)).contains(_67_v6)) {
              _coll26.push([_67_v6,(_dafny.ZERO).minus(new BigNumber(-297))]);
            }
          }
          return _coll26;
        }())).Union(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('d'.codePointAt(0)),new BigNumber(85)), function () {
          let _coll27 = new _dafny.Map();
          for (const _compr_27 of (_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0)))).Elements) {
            let _68_v7 = _compr_27;
            if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('q'.codePointAt(0))), _68_v7)) {
              _coll27.push([_68_v7,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-771)), function (_69_i2) {
                return new _dafny.CodePoint('d'.codePointAt(0));
              })).length))).length)]);
            }
          }
          return _coll27;
        }()));
      }
    };
    static fm71(p0, globalState) {
      return _dafny.Seq.of(_dafny.Seq.of(false, !(false)), _dafny.Seq.of(!(true)), _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true)), _dafny.Seq.Concat(_dafny.Seq.of(false, true, true), _dafny.Seq.of(!(true))));
    };
    static fm72(globalState) {
      return _dafny.Seq.of(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(395)), function (_70_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("usta")));
    };
    static Main(__noArgsParameter) {
      let _71_globalState;
      let _nw0 = new _module.GlobalState();
      _nw0.__ctor();
      _71_globalState = _nw0;
      let _72_v0;
      _72_v0 = new BigNumber(936);
      let _73_v1;
      let _nw1 = new _module.C4();
      _nw1.__ctor();
      _73_v1 = _nw1;
      let _74_v3;
      _74_v3 = _dafny.Seq.UnicodeFromString("dtn");
      let _75_v4;
      _75_v4 = _dafny.MultiSet.fromElements(_74_v3, _74_v3);
      let _76_v5;
      let _init0 = function (_77_i0) {
        return _dafny.Set.fromElements(true, true);
      };
      let _nw2 = Array((new BigNumber(5)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw2.length); _i0_0++) {
        _nw2[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _76_v5 = _nw2;
      let _78_v6;
      _78_v6 = _module.D0.create_DC0(_76_v5, _72_v0);
      let _79_v7;
      _79_v7 = _dafny.Seq.of(_72_v0);
      let _80_v8;
      _80_v8 = _module.D0.create_DC1(_72_v0);
      let _81_v9;
      let _nw3 = new _module.C8();
      _nw3.__ctor(_module.__default.safeDivisionInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_72_v0)).length),_73_v1)).length), new BigNumber((function () {
        let _coll28 = new _dafny.Map();
        for (const _compr_28 of (_75_v4).Elements) {
          let _82_v2 = _compr_28;
          if ((_75_v4).contains(_82_v2)) {
            _coll28.push([_82_v2,new BigNumber(-676)]);
          }
        }
        return _coll28;
      }()).length)), _78_v6, _dafny.Seq.Concat(_79_v7, _79_v7), _80_v8);
      _81_v9 = _nw3;
      let _83_v10;
      _83_v10 = false;
      let _pat_let_tv0 = _72_v0;
      let _84_v11;
      let _85_v12;
      let _86_v13;
      let _out0;
      let _out1;
      let _out2;
      let _outcollector0 = (_81_v9).m3(new BigNumber(490), function (_pat_let0_0) {
        return function (_87_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_88_dt__update_hcf1_h0) {
              return _module.D0.create_DC0((_87_dt__update__tmp_h0).dtor_cf0, _88_dt__update_hcf1_h0);
            }(_pat_let1_0);
          }(_pat_let_tv0);
        }(_pat_let0_0);
      }(_module.D0.create_DC0(_76_v5, new BigNumber((_74_v3).length))), _83_v10, _71_globalState);
      _out0 = _outcollector0[0];
      _out1 = _outcollector0[1];
      _out2 = _outcollector0[2];
      _84_v11 = _out0;
      _85_v12 = _out1;
      _86_v13 = _out2;
      let _89_v14;
      _89_v14 = _dafny.Map.Empty.slice().updateUnsafe(_83_v10,new BigNumber(548));
      let _90_v15;
      _90_v15 = _dafny.Map.Empty.slice().updateUnsafe(_89_v14,_74_v3);
      let _91_v16;
      _91_v16 = _module.D18.create_DC47(_90_v15);
      let _pat_let_tv1 = _91_v16;
      let _source2 = ((_83_v10) ? (function (_pat_let2_0) {
        return function (_92_dt__update__tmp_h1) {
          return function (_pat_let3_0) {
            return function (_93_dt__update_hcf85_h0) {
              return _module.D18.create_DC47(_93_dt__update_hcf85_h0);
            }(_pat_let3_0);
          }((_pat_let_tv1).dtor_cf85);
        }(_pat_let2_0);
      }(_91_v16)) : (_91_v16));
      if (_source2.is_DC48) {
        if (_83_v10) {
          let _94_v17;
          let _nw4 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _94_v17 = _nw4;
          _94_v17 = _94_v17;
          let _95_v18;
          let _nw5 = Array((new BigNumber(14)).toNumber()).fill(false);
          _95_v18 = _nw5;
          let _index0 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_95_v18).length));
          (_95_v18)[_index0] = _83_v10;
          let _96_v19;
          _96_v19 = _dafny.MultiSet.fromElements(_84_v11);
          let _pat_let_tv2 = _81_v9;
          let _97_v20;
          let _nw6 = new _module.C7();
          _nw6.__ctor(_96_v19, function (_pat_let4_0) {
            return function (_98_dt__update__tmp_h2) {
              return function (_pat_let5_0) {
                return function (_99_dt__update_hcf2_h0) {
                  return _module.D0.create_DC1(_99_dt__update_hcf2_h0);
                }(_pat_let5_0);
              }((_pat_let_tv2).f2);
            }(_pat_let4_0);
          }(_80_v8), _dafny.Seq.Concat(_79_v7, _79_v7), _module.D0.create_DC1(_72_v0));
          _97_v20 = _nw6;
          let _100_v21;
          let _nw7 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _100_v21 = _nw7;
          let _index1 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_100_v21).length));
          (_100_v21)[_index1] = new BigNumber((_module.__default.fm70(_83_v10, _71_globalState)).length);
          let _101_v22;
          _101_v22 = new _dafny.CodePoint('e'.codePointAt(0));
          let _102_v23;
          _102_v23 = _module.D16.create_DC43(_dafny.Seq.Create(_module.__default.abs(new BigNumber(403)), function (_103_i1) {
  return new _dafny.CodePoint('c'.codePointAt(0));
}), _101_v22, _84_v11);
          let _104_v24;
          let _nw8 = new _module.C0();
          _nw8.__ctor((_102_v23).dtor_cf79);
          _104_v24 = _nw8;
          let _105_v25;
          _105_v25 = _dafny.Map.Empty.slice().updateUnsafe(_104_v24,_83_v10);
          let _index2 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_95_v18).length));
          let _index3 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_100_v21).length));
          let _rhs0 = _83_v10;
          let _rhs1 = (((((_105_v25).contains(_104_v24)) ? ((_105_v25).get(_104_v24)) : (_83_v10))) ? (_97_v20) : (_97_v20));
          let _rhs2 = _95_v18;
          let _rhs3 = (_81_v9).f2;
          let _rhs4 = ((_module.__default.fm44(true, _71_globalState)).isLessThan(_86_v13)) && (_83_v10);
          let _lhs0 = _95_v18;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_95_v18).length));
          let _lhs2 = _100_v21;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(721), new BigNumber((_100_v21).length));
          _lhs0[_lhs1] = _rhs0;
          _97_v20 = _rhs1;
          _95_v18 = _rhs2;
          _lhs2[_lhs3] = _rhs3;
          _83_v10 = _rhs4;
          let _index4 = _module.__default.safeIndex(new BigNumber(213), new BigNumber((_95_v18).length));
          (_95_v18)[_index4] = !(!(_89_v14).contains(_83_v10)) || (true);
          let _106_v26;
          _106_v26 = _dafny.Map.Empty.slice().updateUnsafe((_95_v18)[_module.__default.safeIndex(new BigNumber(213), new BigNumber((_95_v18).length))],_module.__default.fm17(_71_globalState));
          let _107_v27;
          _107_v27 = _dafny.Seq.of((_95_v18)[_module.__default.safeIndex(new BigNumber(213), new BigNumber((_95_v18).length))], false, _83_v10, _83_v10, (_95_v18)[_module.__default.safeIndex(new BigNumber(213), new BigNumber((_95_v18).length))]);
          let _108_v28;
          _108_v28 = _dafny.MultiSet.fromElements((_107_v27)[_module.__default.safeIndex(_86_v13, new BigNumber((_107_v27).length))]);
          let _109_v29;
          _109_v29 = _dafny.Map.Empty.slice().updateUnsafe((_81_v9).f2,new BigNumber(((_dafny.MultiSet.fromElements(!((((_106_v26).contains(_83_v10)) ? ((_106_v26).get(_83_v10)) : ((_95_v18)[_module.__default.safeIndex(new BigNumber(213), new BigNumber((_95_v18).length))]))))).Difference(_108_v28)).cardinality()));
          _109_v29 = (_109_v29).update(_86_v13, (_81_v9).f2);
          let _110_v30;
          let _111_v31;
          let _out3;
          let _out4;
          let _outcollector1 = (_81_v9).m1((((_109_v29).contains((_100_v21)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_100_v21).length))])) ? ((_109_v29).get((_100_v21)[_module.__default.safeIndex(new BigNumber(721), new BigNumber((_100_v21).length))])) : ((_102_v23).dtor_cf81)), _71_globalState);
          _out3 = _outcollector1[0];
          _out4 = _outcollector1[1];
          _110_v30 = _out3;
          _111_v31 = _out4;
        } else {
          let _112_v32;
          _112_v32 = new _dafny.CodePoint('c'.codePointAt(0));
          let _113_v33;
          _113_v33 = _dafny.Map.Empty.slice().updateUnsafe(_112_v32,(_dafny.ZERO).minus(_module.__default.fm32((_81_v9).f2, _71_globalState)));
          let _114_v34;
          let _nw9 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
          _114_v34 = _nw9;
          let _115_v35;
          let _nw10 = new _module.C5();
          _nw10.__ctor(_114_v34, _79_v7, _80_v8);
          _115_v35 = _nw10;
          let _116_v36;
          let _nw11 = new _module.C0();
          _nw11.__ctor(_74_v3);
          _116_v36 = _nw11;
          let _117_v37;
          _117_v37 = _dafny.Map.Empty.slice().updateUnsafe(_115_v35,_116_v36);
          let _118_v38;
          let _nw12 = Array((new BigNumber(29)).toNumber());
          _nw12[(_dafny.ZERO).toNumber()] = false;
          _nw12[(_dafny.ONE).toNumber()] = _83_v10;
          _nw12[(new BigNumber(2)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(3)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(4)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(5)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(6)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(7)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(8)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(9)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(10)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(11)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(12)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(13)).toNumber()] = true;
          _nw12[(new BigNumber(14)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(15)).toNumber()] = !(_83_v10);
          _nw12[(new BigNumber(16)).toNumber()] = _module.__default.fm17(_71_globalState);
          _nw12[(new BigNumber(17)).toNumber()] = true;
          _nw12[(new BigNumber(18)).toNumber()] = true;
          _nw12[(new BigNumber(19)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(20)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(21)).toNumber()] = false;
          _nw12[(new BigNumber(22)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(23)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(24)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(25)).toNumber()] = true;
          _nw12[(new BigNumber(26)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(27)).toNumber()] = _83_v10;
          _nw12[(new BigNumber(28)).toNumber()] = _83_v10;
          _118_v38 = _nw12;
          let _119_v39;
          let _120_v40;
          let _out5;
          let _out6;
          let _outcollector2 = (_73_v1).m10(((_81_v9).f2).plus((_81_v9).f2), new BigNumber(((_113_v33).update(_module.__default.fm51(_83_v10, (_81_v9).f2, new BigNumber((_117_v37).length), _71_globalState), _72_v0)).length), _83_v10, _118_v38, _71_globalState);
          _out5 = _outcollector2[0];
          _out6 = _outcollector2[1];
          _119_v39 = _out5;
          _120_v40 = _out6;
          let _121_v41;
          _121_v41 = _dafny.Seq.of(_83_v10);
          _121_v41 = _dafny.Seq.of((_119_v39) === (_83_v10), _119_v39, false, _83_v10);
          _84_v11 = new BigNumber(-509);
          let _122_v42;
          _122_v42 = _dafny.Seq.of(_118_v38, _118_v38, _118_v38);
          _83_v10 = _dafny.Seq.IsPrefixOf(_dafny.Seq.of(_118_v38, _118_v38), _122_v42);
          let _123_v43;
          _123_v43 = _dafny.Map.Empty.slice().updateUnsafe(_83_v10,false);
          let _124_v44;
          let _125_v45;
          let _126_v46;
          let _out7;
          let _out8;
          let _out9;
          let _outcollector3 = (_81_v9).m3((_dafny.ZERO).minus((_dafny.ZERO).minus(_86_v13)), _78_v6, (((_123_v43).contains(false)) ? ((_123_v43).get(false)) : (!(_83_v10))), _71_globalState);
          _out7 = _outcollector3[0];
          _out8 = _outcollector3[1];
          _out9 = _outcollector3[2];
          _124_v44 = _out7;
          _125_v45 = _out8;
          _126_v46 = _out9;
        }
        let _127_v47;
        _127_v47 = _dafny.Seq.of(!(_83_v10), _83_v10);
        let _128_v48;
        _128_v48 = _dafny.Set.fromElements(new _dafny.CodePoint('f'.codePointAt(0)));
        let _129_v49;
        _129_v49 = _dafny.Map.Empty.slice().updateUnsafe(_127_v47,_dafny.Seq.update(_127_v47, _module.__default.safeIndex(new BigNumber((_128_v48).length), new BigNumber((_127_v47).length)), _83_v10));
        _129_v49 = (_dafny.Map.Empty.slice().updateUnsafe(_127_v47,_127_v47)).Merge((_129_v49).Merge(_129_v49));
        _83_v10 = _83_v10;
        let _130_v50;
        let _nw13 = new _module.C4();
        _nw13.__ctor();
        _130_v50 = _nw13;
        let _131_v51;
        _131_v51 = _module.D20.create_DC53(_83_v10);
        let _132_v52;
        _132_v52 = _dafny.Map.Empty.slice().updateUnsafe(_130_v50,(_131_v51).dtor_cf94);
        let _133_v53;
        _133_v53 = _130_v50;
        _132_v52 = (_132_v52).update((_133_v53), true);
      } else if (_source2.is_DC49) {
        let _134___mcc_h0 = (_source2).cf86;
        let _135_cf86 = _134___mcc_h0;
        if (_83_v10) {
          let _136_v54;
          _136_v54 = new _dafny.CodePoint('n'.codePointAt(0));
          let _137_v55;
          _137_v55 = _dafny.Map.Empty.slice().updateUnsafe(_136_v54,_80_v8);
          let _138_v56;
          _138_v56 = _dafny.Map.Empty.slice().updateUnsafe(true,_module.__default.fm71((_81_v9).fm5((_81_v9).f2, _83_v10, _137_v55, _83_v10, _71_globalState), _71_globalState));
          let _139_v57;
          _139_v57 = _dafny.Seq.of(_83_v10, _83_v10, _83_v10);
          let _140_v58;
          _140_v58 = _dafny.MultiSet.fromElements(_83_v10, _83_v10);
          let _141_v59;
          _141_v59 = _dafny.Seq.of(_dafny.Seq.of(_83_v10, !(_83_v10)));
          _84_v11 = new BigNumber(((((_138_v56).contains((_140_v58).IsProperSubsetOf(_dafny.MultiSet.FromArray(_139_v57)))) ? ((_138_v56).get((_140_v58).IsProperSubsetOf(_dafny.MultiSet.FromArray(_139_v57)))) : (_dafny.Seq.Concat(_141_v59, _141_v59)))).length);
          let _142_v60;
          _142_v60 = _dafny.Set.fromElements(_83_v10);
          _89_v14 = (_89_v14).update((_142_v60).IsProperSubsetOf(_142_v60), (_81_v9).f2);
          let _143_v62;
          _143_v62 = _dafny.MultiSet.fromElements(new BigNumber((function () {
            let _coll29 = new _dafny.Set();
            for (const _compr_29 of (_79_v7).Elements) {
              let _144_v61 = _compr_29;
              if (_dafny.Seq.contains(_79_v7, _144_v61)) {
                _coll29.add((_144_v61).minus(new BigNumber(185)));
              }
            }
            return _coll29;
          }()).length), new BigNumber(12));
          let _145_v63;
          let _nw14 = new _module.C7();
          _nw14.__ctor((((_143_v62).update((_81_v9).f2, _module.__default.abs(_72_v0))).update(_72_v0, _module.__default.abs(_86_v13))).update(_72_v0, _module.__default.abs(_72_v0)), ((_83_v10) ? (_module.__default.fm49(_136_v54, _71_globalState)) : (_80_v8)), _79_v7, _80_v8);
          _145_v63 = _nw14;
          let _146_v64;
          let _nw15 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
          _146_v64 = _nw15;
          let _147_v65;
          let _nw16 = new _module.C1();
          _nw16.__ctor(_84_v11, _146_v64, _dafny.Seq.update(_79_v7, _module.__default.safeIndex((_81_v9).f2, new BigNumber((_79_v7).length)), _72_v0), _module.D0.create_DC1(_72_v0));
          _147_v65 = _nw16;
          let _148_v66;
          _148_v66 = _dafny.Map.Empty.slice().updateUnsafe(_147_v65,_136_v54);
          _148_v66 = (_148_v66).update(_147_v65, _136_v54);
          let _149_v67;
          let _nw17 = new _module.C8();
          _nw17.__ctor(_147_v65.f9, _78_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(169)), ((_150_v9) => function (_151_i2) {
            return (_150_v9).f2;
          })(_81_v9)), _module.D0.create_DC1(new BigNumber((_139_v57).length)));
          _149_v67 = _nw17;
          let _152_v68;
          _152_v68 = _dafny.Map.Empty.slice().updateUnsafe(_149_v67,(_81_v9).f2);
          let _153_v69;
          let _nw18 = Array((new BigNumber(24)).toNumber()).fill(false);
          _153_v69 = _nw18;
          let _index5 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_153_v69).length));
          (_153_v69)[_index5] = _83_v10;
          let _index6 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_153_v69).length));
          let _rhs5 = (_152_v68).Merge(_152_v68);
          let _rhs6 = _83_v10;
          let _lhs4 = _153_v69;
          let _lhs5 = _module.__default.safeIndex(new BigNumber(414), new BigNumber((_153_v69).length));
          _152_v68 = _rhs5;
          _lhs4[_lhs5] = _rhs6;
        } else {
          let _154_v70;
          let _nw19 = Array((_dafny.ONE).toNumber()).fill(_module.D17.Default());
          _154_v70 = _nw19;
          let _155_v71;
          _155_v71 = _dafny.MultiSet.fromElements(_86_v13, (_81_v9).f2);
          let _156_v72;
          _156_v72 = _module.D17.create_DC45(_dafny.MultiSet.fromElements(_155_v71, _155_v71));
          let _index7 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_154_v70).length));
          (_154_v70)[_index7] = _156_v72;
          let _index8 = _module.__default.safeIndex(new BigNumber(395), new BigNumber((_154_v70).length));
          (_154_v70)[_index8] = _156_v72;
          let _157_v73;
          _157_v73 = _73_v1;
          let _158_v74;
          let _nw20 = Array((new BigNumber(28)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = _73_v1;
          _nw20[(_dafny.ONE).toNumber()] = _73_v1;
          _nw20[(new BigNumber(2)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(3)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(4)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(5)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(6)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(7)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(8)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(9)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(10)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(11)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(12)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(13)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(14)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(15)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(16)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(17)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(18)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(19)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(20)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(21)).toNumber()] = (_157_v73);
          _nw20[(new BigNumber(22)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(23)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(24)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(25)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(26)).toNumber()] = _73_v1;
          _nw20[(new BigNumber(27)).toNumber()] = _73_v1;
          _158_v74 = _nw20;
          let _159_v75;
          _159_v75 = _dafny.Seq.of(_158_v74, _158_v74, _158_v74, _158_v74);
          let _rhs7 = (_81_v9).f2;
          let _rhs8 = (_81_v9).fm7(_86_v13, _71_globalState);
          let _rhs9 = ((_83_v10) ? (_dafny.Seq.Concat(_dafny.Seq.of(_158_v74, _158_v74), _dafny.Seq.of(_158_v74))) : (_159_v75));
          _84_v11 = _rhs7;
          _74_v3 = _rhs8;
          _159_v75 = _rhs9;
          let _160_v76;
          let _nw21 = Array((new BigNumber(21)).toNumber()).fill([]);
          _160_v76 = _nw21;
          let _161_v77;
          _161_v77 = _module.D4.create_DC10((_dafny.ZERO).minus((_84_v11).plus((_81_v9).f2)), _86_v13, new BigNumber((_74_v3).length));
          let _162_v78;
          _162_v78 = _dafny.Seq.of(_160_v76, _160_v76, _160_v76, _160_v76);
          let _163_v79;
          _163_v79 = _dafny.Map.Empty.slice().updateUnsafe(_83_v10,(_162_v78)[_module.__default.safeIndex(_module.__default.fm44(_83_v10, _71_globalState), new BigNumber((_162_v78).length))]);
          let _rhs10 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(335)), function (_164_i3) {
            return new _dafny.CodePoint('d'.codePointAt(0));
          }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(235)), function (_165_i4) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          }))).length);
          let _rhs11 = (((_163_v79).contains(true)) ? ((_163_v79).get(true)) : (_160_v76));
          let _rhs12 = _161_v77;
          let _rhs13 = ((_83_v10) ? ((_81_v9).f2) : (_module.__default.fm32((_dafny.ZERO).minus(_72_v0), _71_globalState)));
          _84_v11 = _rhs10;
          _160_v76 = _rhs11;
          _161_v77 = _rhs12;
          _84_v11 = _rhs13;
          _84_v11 = ((_72_v0).plus(new BigNumber(-618))).minus((_dafny.ZERO).minus(new BigNumber((_74_v3).length)));
          _84_v11 = (((_75_v4).contains(_74_v3)) ? ((_75_v4).get(_74_v3)) : (_84_v11));
        }
        _83_v10 = _83_v10;
        let _166_v80;
        _166_v80 = _module.D7.create_DC18();
        _166_v80 = _166_v80;
        _83_v10 = (_83_v10) === ((_84_v11).isEqualTo((_81_v9).f2));
      } else if (_source2.is_DC50) {
        let _167___mcc_h1 = (_source2).cf87;
        let _168___mcc_h2 = (_source2).cf88;
        let _169___mcc_h3 = (_source2).cf89;
        let _170___mcc_h4 = (_source2).cf90;
        let _171___mcc_h5 = (_source2).cf91;
        let _172_cf91 = _171___mcc_h5;
        let _173_cf90 = _170___mcc_h4;
        let _174_cf89 = _169___mcc_h3;
        let _175_cf88 = _168___mcc_h2;
        let _176_cf87 = _167___mcc_h1;
        let _177_v81;
        let _nw22 = Array((new BigNumber(7)).toNumber()).fill(_dafny.Map.Empty);
        _177_v81 = _nw22;
        let _178_v82;
        _178_v82 = _dafny.Seq.of(_79_v7);
        let _179_v83;
        let _nw23 = new _module.C1();
        _nw23.__ctor((new BigNumber(-276)).plus((_81_v9).f2), ((_174_cf89) ? (_177_v81) : (_177_v81)), _dafny.Seq.of(_72_v0, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Set.fromElements((_178_v82)[_module.__default.safeIndex((_81_v9).f2, new BigNumber((_178_v82).length))])).length))).cardinality())), _80_v8);
        _179_v83 = _nw23;
        _179_v83 = _179_v83;
        _174_cf89 = !(!(_83_v10));
        let _180_v84;
        let _181_v85;
        let _out10;
        let _out11;
        let _outcollector4 = (_81_v9).m1(_86_v13, _71_globalState);
        _out10 = _outcollector4[0];
        _out11 = _outcollector4[1];
        _180_v84 = _out10;
        _181_v85 = _out11;
        _84_v11 = (_81_v9).f2;
      } else {
        let _182___mcc_h6 = (_source2).cf85;
        let _183_cf85 = _182___mcc_h6;
        _83_v10 = _83_v10;
        _83_v10 = !(!(new BigNumber(780)).isEqualTo(((_83_v10) ? (_84_v11) : (new BigNumber(770)))));
        _84_v11 = (_86_v13).multipliedBy(_module.__default.safeDivisionInt(_84_v11, new BigNumber(707)));
        let _184_v86;
        _184_v86 = _dafny.Set.fromElements(_72_v0);
        let _185_v87;
        _185_v87 = _dafny.Set.fromElements(_184_v86);
        let _186_v88;
        _186_v88 = _dafny.Map.Empty.slice().updateUnsafe(_83_v10,(_module.D23.create_DC57(_185_v87)).dtor_cf100);
        _186_v88 = (_186_v88).update(_83_v10, _185_v87);
      }
      let _187_v89;
      _187_v89 = _module.D20.create_DC53(!(_83_v10));
      if ((_187_v89).dtor_cf94) {
        _84_v11 = (_81_v9).f2;
        let _188_v90;
        let _init1 = ((_189_v9) => function (_190_i5) {
          return _module.__default.safeModuloInt(_190_i5, (_189_v9).f2);
        })(_81_v9);
        let _nw24 = Array((new BigNumber(23)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw24.length); _i0_1++) {
          _nw24[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _188_v90 = _nw24;
        let _index9 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_188_v90).length));
        (_188_v90)[_index9] = new BigNumber((_dafny.Seq.Concat(_74_v3, _dafny.Seq.UnicodeFromString("bnkl"))).length);
        let _191_v91;
        _191_v91 = _dafny.Set.fromElements(_83_v10);
        let _index10 = _module.__default.safeIndex(new BigNumber(898), new BigNumber((_188_v90).length));
        (_188_v90)[_index10] = new BigNumber(((_dafny.Set.fromElements(_83_v10)).Difference(_191_v91)).length);
        let _192_v92;
        let _nw25 = new _module.C2();
        _nw25.__ctor();
        _192_v92 = _nw25;
        let _193_v93;
        _193_v93 = _dafny.MultiSet.fromElements(new BigNumber(394));
        let _194_v94;
        _194_v94 = _192_v92;
        let _195_v95;
        _195_v95 = _dafny.Map.Empty.slice().updateUnsafe(!((_193_v93).equals(_193_v93)),(_194_v94));
        let _196_v96;
        _196_v96 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('m'.codePointAt(0)),_80_v8);
        let _197_v97;
        _197_v97 = _dafny.Map.Empty.slice().updateUnsafe(!(!(_83_v10)),(_192_v92).fm5(_72_v0, _83_v10, _196_v96, _83_v10, _71_globalState));
        _192_v92 = (((_195_v95).contains((((_197_v97).contains(_83_v10)) ? ((_197_v97).get(_83_v10)) : (_83_v10)))) ? ((_195_v95).get((((_197_v97).contains(_83_v10)) ? ((_197_v97).get(_83_v10)) : (_83_v10)))) : (_192_v92));
        let _198_v98;
        let _nw26 = Array((new BigNumber(24)).toNumber()).fill([]);
        _198_v98 = _nw26;
        let _199_v99;
        let _init2 = ((_200_v9, _201_v14, _202_v0) => function (_203_i6) {
          return _module.D16.create_DC42((_200_v9).f2, new BigNumber((_201_v14).length), _202_v0);
        })(_81_v9, _89_v14, _72_v0);
        let _nw27 = Array((new BigNumber(15)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw27.length); _i0_2++) {
          _nw27[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _199_v99 = _nw27;
        let _index11 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_198_v98).length));
        (_198_v98)[_index11] = _199_v99;
        let _204_v100;
        _204_v100 = _module.D24.create_DC59(_199_v99);
        let _index12 = _module.__default.safeIndex(new BigNumber(548), new BigNumber((_198_v98).length));
        (_198_v98)[_index12] = (_204_v100).dtor_cf101;
        let _205_v101;
        _205_v101 = new _dafny.CodePoint('a'.codePointAt(0));
        _205_v101 = _205_v101;
      } else {
        let _206_v102;
        let _nw28 = Array((new BigNumber(19)).toNumber()).fill(false);
        _206_v102 = _nw28;
        let _index13 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length));
        (_206_v102)[_index13] = _83_v10;
        let _207_v103;
        _207_v103 = _module.D12.create_DC28(_83_v10, new BigNumber((_74_v3).length));
        let _index14 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length));
        (_206_v102)[_index14] = (_module.__default.safeModuloInt((_207_v103).dtor_cf47, new BigNumber(860))).isEqualTo(new BigNumber(765));
        _83_v10 = !(_83_v10) || (_83_v10);
        let _208_v105;
        _208_v105 = _dafny.Set.fromElements((_dafny.ZERO).minus(_72_v0), _84_v11, new BigNumber((function () {
          let _coll30 = new _dafny.Map();
          for (const _compr_30 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(163)), ((_209_v12) => function (_210_i7) {
            return _209_v12;
          })(_85_v12))).Elements) {
            let _211_v104 = _compr_30;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(163)), ((_212_v12) => function (_210_i7) {
              return _212_v12;
            })(_85_v12)), _211_v104)) {
              _coll30.push([_211_v104,_86_v13]);
            }
          }
          return _coll30;
        }()).length), _72_v0);
        let _213_v106;
        _213_v106 = _dafny.Seq.of((_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))]);
        let _214_v107;
        _214_v107 = new _dafny.CodePoint('m'.codePointAt(0));
        let _215_v108;
        _215_v108 = _dafny.Map.Empty.slice().updateUnsafe((_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))],_214_v107);
        let _216_v109;
        _216_v109 = _dafny.MultiSet.fromElements((_81_v9).f2, new BigNumber((_215_v108).length), _84_v11);
        let _nw29 = Array((new BigNumber(29)).toNumber());
        _nw29[(_dafny.ZERO).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(_dafny.ONE).toNumber()] = _83_v10;
        _nw29[(new BigNumber(2)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(3)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(4)).toNumber()] = !((_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))]);
        _nw29[(new BigNumber(5)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(6)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(7)).toNumber()] = _83_v10;
        _nw29[(new BigNumber(8)).toNumber()] = true;
        _nw29[(new BigNumber(9)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(10)).toNumber()] = ((_83_v10) ? ((_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))]) : (false));
        _nw29[(new BigNumber(11)).toNumber()] = _83_v10;
        _nw29[(new BigNumber(12)).toNumber()] = _83_v10;
        _nw29[(new BigNumber(13)).toNumber()] = (_83_v10) === (_83_v10);
        _nw29[(new BigNumber(14)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(15)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(16)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(17)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(18)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(19)).toNumber()] = (_72_v0).isLessThanOrEqualTo(_86_v13);
        _nw29[(new BigNumber(20)).toNumber()] = (_208_v105).IsSubsetOf(_dafny.Set.fromElements(_72_v0));
        _nw29[(new BigNumber(21)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(22)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(23)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_74_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-976)), function (_217_i8) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        }));
        _nw29[(new BigNumber(24)).toNumber()] = (_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))];
        _nw29[(new BigNumber(25)).toNumber()] = !((_213_v106)[_module.__default.safeIndex(new BigNumber((_216_v109).cardinality()), new BigNumber((_213_v106).length))]);
        _nw29[(new BigNumber(26)).toNumber()] = ((_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))]) === (true);
        _nw29[(new BigNumber(27)).toNumber()] = _83_v10;
        _nw29[(new BigNumber(28)).toNumber()] = ((_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))]) && (_83_v10);
        _206_v102 = _nw29;
        let _218_v110;
        _218_v110 = _dafny.Map.Empty.slice().updateUnsafe(_83_v10,((_83_v10) ? ((_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))]) : ((_206_v102)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_206_v102).length))])));
        _218_v110 = (_218_v110).update((_72_v0).isEqualTo(_86_v13), _83_v10);
        let _219_v111;
        _219_v111 = _dafny.Map.Empty.slice().updateUnsafe(_86_v13,(_81_v9).fm6(_72_v0, new BigNumber((_79_v7).length), _71_globalState));
        _72_v0 = (((_219_v111).contains(_84_v11)) ? ((_219_v111).get(_84_v11)) : (_72_v0));
      }
      let _220_v112;
      _220_v112 = _dafny.Map.Empty.slice().updateUnsafe(_72_v0,(_81_v9).f2);
      _72_v0 = ((_81_v9).f2).multipliedBy((_dafny.ZERO).minus(new BigNumber((_220_v112).length)));
      let _221_i9;
      _221_i9 = _dafny.ZERO;
      L0: {
        let _pat_let_tv3 = _83_v10;
        let _pat_let_tv4 = _83_v10;
        let _pat_let_tv5 = _83_v10;
        while (!(function (_source3) {
          if (_source3.is_DC48) {
            return _pat_let_tv3;
          } else if (_source3.is_DC49) {
            let _227___mcc_h7 = (_source3).cf86;
            let _228_cf86 = _227___mcc_h7;
            return _pat_let_tv4;
          } else if (_source3.is_DC50) {
            let _229___mcc_h8 = (_source3).cf87;
            let _230___mcc_h9 = (_source3).cf88;
            let _231___mcc_h10 = (_source3).cf89;
            let _232___mcc_h11 = (_source3).cf90;
            let _233___mcc_h12 = (_source3).cf91;
            let _234_cf91 = _233___mcc_h12;
            let _235_cf90 = _232___mcc_h11;
            let _236_cf89 = _231___mcc_h10;
            let _237_cf88 = _230___mcc_h9;
            let _238_cf87 = _229___mcc_h8;
            return _238_cf87;
          } else {
            let _239___mcc_h13 = (_source3).cf85;
            let _240_cf85 = _239___mcc_h13;
            return _pat_let_tv5;
          }
        }(_91_v16))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_221_i9)) {
              break L0;
            }
            _221_i9 = (_221_i9).plus(_dafny.ONE);
            let _222_v113;
            _222_v113 = _dafny.Map.Empty.slice().updateUnsafe(_83_v10,_83_v10);
            let _rhs14 = _module.__default.fm47(((_81_v9).f2).multipliedBy(_72_v0), ((_81_v9).f2).isLessThanOrEqualTo(_72_v0), _71_globalState);
            let _rhs15 = (_72_v0).isLessThan(_84_v11);
            _222_v113 = _rhs14;
            _83_v10 = _rhs15;
            let _223_v114;
            let _nw30 = Array((new BigNumber(5)).toNumber()).fill(false);
            _223_v114 = _nw30;
            let _index15 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_223_v114).length));
            (_223_v114)[_index15] = _83_v10;
            let _224_v115;
            let _nw31 = new _module.C2();
            _nw31.__ctor();
            _224_v115 = _nw31;
            let _index16 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_223_v114).length));
            let _rhs16 = _73_v1;
            let _rhs17 = !(!(_83_v10));
            let _rhs18 = _223_v114;
            let _rhs19 = (_86_v13).plus(new BigNumber((_dafny.Seq.of(_224_v115)).length));
            let _lhs6 = _223_v114;
            let _lhs7 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_223_v114).length));
            _73_v1 = _rhs16;
            _lhs6[_lhs7] = _rhs17;
            _223_v114 = _rhs18;
            _72_v0 = _rhs19;
            let _225_v116;
            let _nw32 = new _module.C3();
            _nw32.__ctor();
            _225_v116 = _nw32;
            let _226_v117;
            _226_v117 = _dafny.Seq.of(_225_v116);
            _225_v116 = ((_83_v10) ? (_225_v116) : ((_226_v117)[_module.__default.safeIndex((_81_v9).f2, new BigNumber((_226_v117).length))]));
            _83_v10 = (_223_v114)[_module.__default.safeIndex(new BigNumber(792), new BigNumber((_223_v114).length))];
          }
        }
      }
      let _241_v118;
      let _nw33 = new _module.C4();
      _nw33.__ctor();
      _241_v118 = _nw33;
      let _242_v119;
      let _nw34 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
      _242_v119 = _nw34;
      let _243_v120;
      _243_v120 = _dafny.Map.Empty.slice().updateUnsafe(_242_v119,_242_v119);
      let _244_v121;
      _244_v121 = _dafny.Map.Empty.slice().updateUnsafe(_83_v10,_dafny.Map.Empty.slice().updateUnsafe(_242_v119,_242_v119));
      let _245_v122;
      let _nw35 = Array((new BigNumber(16)).toNumber());
      _nw35[(_dafny.ZERO).toNumber()] = ((_83_v10) ? (_243_v120) : (_243_v120));
      _nw35[(_dafny.ONE).toNumber()] = (_243_v120).Merge(_243_v120);
      _nw35[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_242_v119,_242_v119)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_242_v119,_242_v119));
      _nw35[(new BigNumber(3)).toNumber()] = (_243_v120).Merge(_243_v120);
      _nw35[(new BigNumber(4)).toNumber()] = (_243_v120).Merge(_dafny.Map.Empty.slice().updateUnsafe(_242_v119,_242_v119));
      _nw35[(new BigNumber(5)).toNumber()] = (_243_v120).update(_242_v119, _242_v119);
      _nw35[(new BigNumber(6)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_242_v119,_242_v119)).Merge(_243_v120);
      _nw35[(new BigNumber(7)).toNumber()] = (((_244_v121).contains(_83_v10)) ? ((_244_v121).get(_83_v10)) : (_243_v120));
      _nw35[(new BigNumber(8)).toNumber()] = _243_v120;
      _nw35[(new BigNumber(9)).toNumber()] = _243_v120;
      _nw35[(new BigNumber(10)).toNumber()] = (_243_v120).update(_242_v119, _242_v119);
      _nw35[(new BigNumber(11)).toNumber()] = _243_v120;
      _nw35[(new BigNumber(12)).toNumber()] = _243_v120;
      _nw35[(new BigNumber(13)).toNumber()] = _243_v120;
      _nw35[(new BigNumber(14)).toNumber()] = _243_v120;
      _nw35[(new BigNumber(15)).toNumber()] = _243_v120;
      _245_v122 = _nw35;
      let _index17 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_245_v122).length));
      (_245_v122)[_index17] = _243_v120;
      let _index18 = _module.__default.safeIndex(new BigNumber(343), new BigNumber((_245_v122).length));
      (_245_v122)[_index18] = _243_v120;
      _83_v10 = false;
      let _246_v123;
      _246_v123 = _dafny.Set.fromElements(false, _83_v10, true, _83_v10, true);
      let _247_v124;
      _247_v124 = _dafny.Map.Empty.slice().updateUnsafe(_83_v10,_246_v123);
      let _248_v125;
      _248_v125 = _dafny.Map.Empty.slice().updateUnsafe(_72_v0,_247_v124);
      let _249_v126;
      _249_v126 = _dafny.Seq.of(false, _83_v10, false);
      let _250_v127;
      _250_v127 = _dafny.Seq.of(_247_v124, _247_v124, _247_v124);
      let _251_v128;
      let _nw36 = Array((new BigNumber(28)).toNumber());
      _nw36[(_dafny.ZERO).toNumber()] = (((_248_v125).contains(new BigNumber((_249_v126).length))) ? ((_248_v125).get(new BigNumber((_249_v126).length))) : (_247_v124));
      _nw36[(_dafny.ONE).toNumber()] = _247_v124;
      _nw36[(new BigNumber(2)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(3)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(4)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(5)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(6)).toNumber()] = (_247_v124).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_246_v123));
      _nw36[(new BigNumber(7)).toNumber()] = (_250_v127)[_module.__default.safeIndex(_72_v0, new BigNumber((_250_v127).length))];
      _nw36[(new BigNumber(8)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(9)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(10)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(true,_246_v123)).Merge(_247_v124);
      _nw36[(new BigNumber(11)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(12)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(13)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(14)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(15)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(16)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(17)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(18)).toNumber()] = ((_83_v10) ? (_247_v124) : (_247_v124));
      _nw36[(new BigNumber(19)).toNumber()] = (_247_v124).Merge(_247_v124);
      _nw36[(new BigNumber(20)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(21)).toNumber()] = (((_248_v125).contains((_81_v9).f2)) ? ((_248_v125).get((_81_v9).f2)) : (_247_v124));
      _nw36[(new BigNumber(22)).toNumber()] = (_247_v124).update(true, _246_v123);
      _nw36[(new BigNumber(23)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(24)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(25)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(26)).toNumber()] = _247_v124;
      _nw36[(new BigNumber(27)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(false,_246_v123)).Merge(_247_v124);
      _251_v128 = _nw36;
      let _index19 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_251_v128).length));
      (_251_v128)[_index19] = _247_v124;
      let _index20 = _module.__default.safeIndex(new BigNumber(312), new BigNumber((_251_v128).length));
      (_251_v128)[_index20] = _247_v124;
      _83_v10 = _83_v10;
      let _252_v129;
      _252_v129 = new _dafny.CodePoint('a'.codePointAt(0));
      let _253_v130;
      let _nw37 = Array((new BigNumber(3)).toNumber());
      _nw37[(_dafny.ZERO).toNumber()] = _252_v129;
      _nw37[(_dafny.ONE).toNumber()] = _252_v129;
      _nw37[(new BigNumber(2)).toNumber()] = _252_v129;
      _253_v130 = _nw37;
      let _254_v131;
      _254_v131 = _dafny.Seq.of(_253_v130);
      let _255_v132;
      _255_v132 = _dafny.Map.Empty.slice().updateUnsafe((_187_v89).dtor_cf94,(_254_v131)[_module.__default.safeIndex(_84_v11, new BigNumber((_254_v131).length))]);
      _255_v132 = (_255_v132).update(_83_v10, _253_v130);
      if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("vdgmtsaje"), _dafny.Seq.UnicodeFromString("lhgie"))).equals(_75_v4)) {
        let _256_v133;
        let _nw38 = new _module.C0();
        _nw38.__ctor(_74_v3);
        _256_v133 = _nw38;
        let _257_v134;
        _257_v134 = _dafny.Map.Empty.slice().updateUnsafe(_256_v133,_83_v10);
        _83_v10 = (((((_257_v134).contains(_256_v133)) ? ((_257_v134).get(_256_v133)) : (_83_v10))) ? ((((_249_v126)[_module.__default.safeIndex(_86_v13, new BigNumber((_249_v126).length))]) ? (_83_v10) : (_83_v10))) : (_83_v10));
        let _258_v135;
        _258_v135 = _module.D24.create_DC62((_81_v9).f2);
        let _pat_let_tv6 = _84_v11;
        let _259_v136;
        _259_v136 = _dafny.Map.Empty.slice().updateUnsafe(function (_pat_let6_0) {
          return function (_260_dt__update__tmp_h3) {
            return function (_pat_let7_0) {
              return function (_261_dt__update_hcf104_h0) {
                return _module.D24.create_DC62(_261_dt__update_hcf104_h0);
              }(_pat_let7_0);
            }(_pat_let_tv6);
          }(_pat_let6_0);
        }(_258_v135),_242_v119);
        let _262_v137;
        let _nw39 = Array((new BigNumber(27)).toNumber());
        _nw39[(_dafny.ZERO).toNumber()] = _242_v119;
        _nw39[(_dafny.ONE).toNumber()] = _242_v119;
        _nw39[(new BigNumber(2)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(3)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(4)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(5)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(6)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(7)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(8)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(9)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(10)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(11)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(12)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(13)).toNumber()] = (((_259_v136).contains(_258_v135)) ? ((_259_v136).get(_258_v135)) : (_242_v119));
        _nw39[(new BigNumber(14)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(15)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(16)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(17)).toNumber()] = ((true) ? (_242_v119) : (_242_v119));
        _nw39[(new BigNumber(18)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(19)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(20)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(21)).toNumber()] = ((_83_v10) ? (_242_v119) : (_242_v119));
        _nw39[(new BigNumber(22)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(23)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(24)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(25)).toNumber()] = _242_v119;
        _nw39[(new BigNumber(26)).toNumber()] = _242_v119;
        _262_v137 = _nw39;
        let _index21 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_262_v137).length));
        (_262_v137)[_index21] = _242_v119;
        let _index22 = _module.__default.safeIndex(new BigNumber(439), new BigNumber((_262_v137).length));
        (_262_v137)[_index22] = _242_v119;
        let _263_v138;
        _263_v138 = _module.D15.create_DC38(_252_v129);
        let _source4 = _263_v138;
        if (_source4.is_DC38) {
          let _264___mcc_h14 = (_source4).cf71;
          let _265_cf71 = _264___mcc_h14;
          let _266_v140;
          _266_v140 = _dafny.Set.fromElements(new BigNumber((_74_v3).length), _84_v11);
          let _267_v141;
          let _nw40 = new _module.C8();
          _nw40.__ctor(new BigNumber((((_83_v10) ? (function () {
            let _coll31 = new _dafny.Set();
            for (const _compr_31 of _dafny.IntegerRange(new BigNumber(109), new BigNumber(874))) {
              let _268_v139 = _compr_31;
              if (((new BigNumber(109)).isLessThanOrEqualTo(_268_v139)) && ((_268_v139).isLessThan(new BigNumber(874)))) {
                _coll31.add((_268_v139).plus((_81_v9).f2));
              }
            }
            return _coll31;
          }()) : (_266_v140))).length), _module.D0.create_DC0(_76_v5, (_81_v9).f2), _dafny.Seq.Concat(_79_v7, _79_v7), _80_v8);
          _267_v141 = _nw40;
          let _nw41 = new _module.C2();
          _nw41.__ctor();
          _267_v141 = _nw41;
          let _269_v142;
          let _out12;
          _out12 = (_81_v9).m2(_84_v11, _71_globalState);
          _269_v142 = _out12;
          let _270_v143;
          _270_v143 = _module.D16.create_DC42((_81_v9).f2, (_81_v9).f2, _86_v13);
          let _271_v145;
          _271_v145 = _dafny.MultiSet.fromElements(_83_v10);
          let _272_v146;
          _272_v146 = _dafny.Seq.of(_271_v145, _271_v145, _271_v145);
          let _273_v147;
          _273_v147 = _dafny.Seq.of(_256_v133.f8, _256_v133.f8, (_256_v133).fm15(_dafny.Seq.of(_86_v13, new BigNumber((function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of (_dafny.MultiSet.FromArray(_272_v146)).Elements) {
              let _274_v144 = _compr_32;
              if ((_dafny.MultiSet.FromArray(_272_v146)).contains(_274_v144)) {
                _coll32.push([_274_v144,_265_cf71]);
              }
            }
            return _coll32;
          }()).length), (_81_v9).f2, _84_v11), _83_v10, _71_globalState), _256_v133.f8, _dafny.Seq.UnicodeFromString("s"));
          let _275_v148;
          let _nw42 = Array((new BigNumber(7)).toNumber());
          _nw42[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(224)), ((_276_v3, _277_v10) => function (_278_i10) {
            return (_module.D15.create_DC39(_276_v3, _module.D9.create_DC21(_dafny.MultiSet.fromElements(false, _277_v10)))).dtor_cf72;
          })(_74_v3, _83_v10)), _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-827)), ((_279_v133) => function (_280_i11) {
            return _279_v133.f8;
          })(_256_v133)), _module.__default.safeIndex((_270_v143).dtor_cf77, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-827)), ((_281_v133) => function (_282_i11) {
            return _281_v133.f8;
          })(_256_v133))).length)), _74_v3));
          _nw42[(_dafny.ONE).toNumber()] = _273_v147;
          _nw42[(new BigNumber(2)).toNumber()] = _273_v147;
          _nw42[(new BigNumber(3)).toNumber()] = _module.__default.fm72(_71_globalState);
          _nw42[(new BigNumber(4)).toNumber()] = _273_v147;
          _nw42[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_74_v3, _256_v133.f8, _dafny.Seq.UnicodeFromString("ea"), _256_v133.f8, _256_v133.f8);
          _nw42[(new BigNumber(6)).toNumber()] = _dafny.Seq.update(_273_v147, _module.__default.safeIndex(new BigNumber(549), new BigNumber((_273_v147).length)), _256_v133.f8);
          _275_v148 = _nw42;
          let _index23 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_275_v148).length));
          (_275_v148)[_index23] = _dafny.Seq.update(_273_v147, _module.__default.safeIndex(_86_v13, new BigNumber((_273_v147).length)), _74_v3);
          let _index24 = _module.__default.safeIndex(new BigNumber(350), new BigNumber((_275_v148).length));
          (_275_v148)[_index24] = _dafny.Seq.Concat(_273_v147, _dafny.Seq.Concat(_273_v147, _273_v147));
          _265_cf71 = _265_cf71;
        } else if (_source4.is_DC39) {
          let _283___mcc_h15 = (_source4).cf72;
          let _284___mcc_h16 = (_source4).cf73;
          let _285_cf73 = _284___mcc_h16;
          let _286_cf72 = _283___mcc_h15;
          let _287_v149;
          _287_v149 = _dafny.Map.Empty.slice().updateUnsafe(_75_v4,((_81_v9).f2).isLessThanOrEqualTo(new BigNumber((_220_v112).length)));
          _287_v149 = (_287_v149).update(_dafny.MultiSet.fromElements(_256_v133.f8, _256_v133.f8, _256_v133.f8), _83_v10);
          let _288_v150;
          let _nw43 = new _module.C8();
          _nw43.__ctor((_81_v9).f2, _78_v6, _79_v7, _80_v8);
          _288_v150 = _nw43;
          let _289_v151;
          _289_v151 = _dafny.MultiSet.fromElements(_288_v150);
          let _290_v152;
          _290_v152 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm32((_dafny.ZERO).minus((((_289_v151).contains(_288_v150)) ? ((_289_v151).get(_288_v150)) : ((_81_v9).f2))), _71_globalState),_256_v133.f8);
          _290_v152 = (_290_v152).update(_72_v0, _dafny.Seq.Concat(_256_v133.f8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(915)), ((_291_v129) => function (_292_i12) {
            return _291_v129;
          })(_252_v129))));
          let _293_v153;
          let _nw44 = new _module.C9();
          _nw44.__ctor();
          _293_v153 = _nw44;
          let _294_v154;
          _294_v154 = (_262_v137)[_module.__default.safeIndex(new BigNumber(439), new BigNumber((_262_v137).length))];
          let _295_v155;
          _295_v155 = _dafny.MultiSet.fromElements(_242_v119, _294_v154, _294_v154);
          _295_v155 = _295_v155;
        } else if (_source4.is_DC37) {
          let _296___mcc_h17 = (_source4).cf70;
          let _297_cf70 = _296___mcc_h17;
          let _298_v156;
          let _299_v157;
          let _300_v158;
          let _out13;
          let _out14;
          let _out15;
          let _outcollector5 = (_73_v1).m9(_71_globalState);
          _out13 = _outcollector5[0];
          _out14 = _outcollector5[1];
          _out15 = _outcollector5[2];
          _298_v156 = _out13;
          _299_v157 = _out14;
          _300_v158 = _out15;
          _84_v11 = ((_81_v9).f2).multipliedBy(_module.__default.fm20(_300_v158, _71_globalState));
          let _301_v159;
          _301_v159 = _dafny.Set.fromElements((_81_v9).f2);
          let _302_v160;
          _302_v160 = _dafny.MultiSet.fromElements(_300_v158);
          let _303_v161;
          let _nw45 = Array((new BigNumber(22)).toNumber());
          _nw45[(_dafny.ZERO).toNumber()] = _300_v158;
          _nw45[(_dafny.ONE).toNumber()] = _dafny.Seq.IsProperPrefixOf(_256_v133.f8, _dafny.Seq.UnicodeFromString("jtsgpyod"));
          _nw45[(new BigNumber(2)).toNumber()] = _300_v158;
          _nw45[(new BigNumber(3)).toNumber()] = (_301_v159).equals(_301_v159);
          _nw45[(new BigNumber(4)).toNumber()] = _83_v10;
          _nw45[(new BigNumber(5)).toNumber()] = _300_v158;
          _nw45[(new BigNumber(6)).toNumber()] = !(!(_302_v160).contains(!(_300_v158)));
          _nw45[(new BigNumber(7)).toNumber()] = _83_v10;
          _nw45[(new BigNumber(8)).toNumber()] = _300_v158;
          _nw45[(new BigNumber(9)).toNumber()] = _300_v158;
          _nw45[(new BigNumber(10)).toNumber()] = _83_v10;
          _nw45[(new BigNumber(11)).toNumber()] = _83_v10;
          _nw45[(new BigNumber(12)).toNumber()] = !(_300_v158) || (_300_v158);
          _nw45[(new BigNumber(13)).toNumber()] = _83_v10;
          _nw45[(new BigNumber(14)).toNumber()] = _83_v10;
          _nw45[(new BigNumber(15)).toNumber()] = (_249_v126)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(_84_v11)).cardinality()), new BigNumber((_249_v126).length))];
          _nw45[(new BigNumber(16)).toNumber()] = !(_300_v158);
          _nw45[(new BigNumber(17)).toNumber()] = _300_v158;
          _nw45[(new BigNumber(18)).toNumber()] = _83_v10;
          _nw45[(new BigNumber(19)).toNumber()] = _dafny.Seq.contains(_256_v133.f8, _252_v129);
          _nw45[(new BigNumber(20)).toNumber()] = _300_v158;
          _nw45[(new BigNumber(21)).toNumber()] = !(_83_v10) || (_83_v10);
          _303_v161 = _nw45;
          let _index25 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_303_v161).length));
          (_303_v161)[_index25] = !(_83_v10) || (_83_v10);
          let _index26 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_303_v161).length));
          let _rhs20 = (_86_v13).isLessThanOrEqualTo((_84_v11).plus((_81_v9).f2));
          let _rhs21 = (_81_v9).f2;
          let _rhs22 = !(_83_v10);
          let _lhs8 = _303_v161;
          let _lhs9 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_303_v161).length));
          _lhs8[_lhs9] = _rhs20;
          _84_v11 = _rhs21;
          _300_v158 = _rhs22;
          let _304_v162;
          _304_v162 = _dafny.MultiSet.fromElements(_86_v13);
          let _index27 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_303_v161).length));
          let _index28 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_303_v161).length));
          let _index29 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_303_v161).length));
          let _rhs23 = _249_v126;
          let _rhs24 = true;
          let _rhs25 = (_304_v162).IsSubsetOf((_304_v162).Union(_module.__default.fm24(_83_v10, _71_globalState)));
          let _rhs26 = (_303_v161)[_module.__default.safeIndex(new BigNumber(491), new BigNumber((_303_v161).length))];
          let _lhs10 = _303_v161;
          let _lhs11 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_303_v161).length));
          let _lhs12 = _303_v161;
          let _lhs13 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_303_v161).length));
          let _lhs14 = _303_v161;
          let _lhs15 = _module.__default.safeIndex(new BigNumber(491), new BigNumber((_303_v161).length));
          _249_v126 = _rhs23;
          _lhs10[_lhs11] = _rhs24;
          _lhs12[_lhs13] = _rhs25;
          _lhs14[_lhs15] = _rhs26;
        } else {
          let _305___mcc_h18 = (_source4).cf74;
          let _306_cf74 = _305___mcc_h18;
          _86_v13 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_241_v118)).length), (_81_v9).f2), (_81_v9).f2);
          let _307_v163;
          let _nw46 = new _module.C4();
          _nw46.__ctor();
          _307_v163 = _nw46;
          let _308_v164;
          let _out16;
          _out16 = (_241_v118).m2((_81_v9).f2, _71_globalState);
          _308_v164 = _out16;
          let _309_v165;
          _309_v165 = _dafny.Map.Empty.slice().updateUnsafe((_81_v9).f2,_83_v10);
          let _pat_let_tv7 = _81_v9;
          let _310_v166;
          _310_v166 = _dafny.Map.Empty.slice().updateUnsafe(_252_v129,function (_pat_let8_0) {
            return function (_311_dt__update__tmp_h4) {
              return function (_pat_let9_0) {
                return function (_312_dt__update_hcf2_h1) {
                  return _module.D0.create_DC1(_312_dt__update_hcf2_h1);
                }(_pat_let9_0);
              }((_pat_let_tv7).f2);
            }(_pat_let8_0);
          }(_80_v8));
          _83_v10 = (_81_v9).fm5(_module.__default.safeDivisionInt(new BigNumber(200), _84_v11), (new BigNumber(993)).isLessThan(new BigNumber((_309_v165).length)), _310_v166, _83_v10, _71_globalState);
        }
        _83_v10 = _dafny.areEqual(_dafny.Seq.UnicodeFromString("sdjic"), _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("q"), _256_v133.f8), _module.__default.safeIndex(new BigNumber((_249_v126).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("q"), _256_v133.f8)).length)), _252_v129));
        let _313_v167;
        _313_v167 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_81_v9).f2, (_81_v9).f2),_242_v119);
        _313_v167 = _313_v167;
      } else {
        let _314_v168;
        let _nw47 = Array((new BigNumber(14)).toNumber());
        _nw47[(_dafny.ZERO).toNumber()] = _242_v119;
        _nw47[(_dafny.ONE).toNumber()] = _242_v119;
        _nw47[(new BigNumber(2)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(3)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(4)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(5)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(6)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(7)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(8)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(9)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(10)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(11)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(12)).toNumber()] = _242_v119;
        _nw47[(new BigNumber(13)).toNumber()] = _242_v119;
        _314_v168 = _nw47;
        let _315_v169;
        _315_v169 = _dafny.MultiSet.fromElements(_83_v10, _83_v10, _83_v10);
        let _316_v170;
        _316_v170 = _dafny.Map.Empty.slice().updateUnsafe((((_315_v169).contains(_83_v10)) ? ((_315_v169).get(_83_v10)) : (new BigNumber(202))),_dafny.MultiSet.FromArray(_249_v126));
        let _317_v171;
        _317_v171 = _module.D25.create_DC63(_316_v170);
        let _318_v172;
        _318_v172 = _dafny.Map.Empty.slice().updateUnsafe(_314_v168,((_317_v171).dtor_cf105).update(new BigNumber(879), (((_316_v170).contains(_72_v0)) ? ((_316_v170).get(_72_v0)) : (_315_v169))));
        _318_v172 = (_318_v172).update(_314_v168, _316_v170);
        let _319_v173;
        let _nw48 = new _module.C4();
        _nw48.__ctor();
        _319_v173 = _nw48;
        let _320_v174;
        _320_v174 = _module.D24.create_DC61(_84_v11, _dafny.MultiSet.fromElements(_242_v119));
        _320_v174 = _320_v174;
        let _321_v175;
        let _nw49 = new _module.C5();
        _nw49.__ctor(_242_v119, _dafny.Seq.of(_86_v13), _80_v8);
        _321_v175 = _nw49;
        _83_v10 = !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("xg"), (_74_v3)[_module.__default.safeIndex(new BigNumber((_74_v3).length), new BigNumber((_74_v3).length))]);
      }
      let _322_i13;
      _322_i13 = _dafny.ZERO;
      L1: {
        while (((_83_v10) ? (_83_v10) : (_83_v10))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_322_i13)) {
              break L1;
            }
            _322_i13 = (_322_i13).plus(_dafny.ONE);
            let _323_v176;
            _323_v176 = _module.D16.create_DC44(_84_v11);
            let _324_v177;
            _324_v177 = _dafny.Set.fromElements(_module.D16.create_DC44((_81_v9).f2), _323_v176, _323_v176, _323_v176);
            let _325_v178;
            _325_v178 = _module.D11.create_DC25(_252_v129, (_81_v9).f2, _83_v10);
            let _rhs27 = ((_324_v177).Intersect(_324_v177)).IsDisjointFrom(_324_v177);
            let _rhs28 = (_325_v178).dtor_cf43;
            let _rhs29 = new BigNumber(559);
            let _rhs30 = _83_v10;
            _83_v10 = _rhs27;
            _83_v10 = _rhs28;
            _72_v0 = _rhs29;
            _83_v10 = _rhs30;
            let _326_v179;
            _326_v179 = _dafny.Map.Empty.slice().updateUnsafe(_249_v126,(_81_v9).f2);
            _326_v179 = (_326_v179).update(_249_v126, _86_v13);
            _84_v11 = ((_83_v10) ? (_84_v11) : ((_dafny.ZERO).minus((_81_v9).f2)));
            let _327_v180;
            let _328_v181;
            let _out17;
            let _out18;
            let _outcollector6 = (_241_v118).m1((_72_v0).plus(_84_v11), _71_globalState);
            _out17 = _outcollector6[0];
            _out18 = _outcollector6[1];
            _327_v180 = _out17;
            _328_v181 = _out18;
          }
        }
      }
      _86_v13 = new BigNumber(670);
      _74_v3 = _74_v3;
      process.stdout.write(_dafny.toString(_72_v0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_74_v3).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_75_v4).equals(_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("dtn"), _dafny.Seq.UnicodeFromString("dtn")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_v5)[_dafny.ZERO]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_v5)[_dafny.ONE]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_v5)[new BigNumber(2)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_v5)[new BigNumber(3)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_76_v5)[new BigNumber(4)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_78_v6).dtor_cf0)[_dafny.ZERO]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_78_v6).dtor_cf0)[_dafny.ONE]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_78_v6).dtor_cf0)[new BigNumber(2)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_78_v6).dtor_cf0)[new BigNumber(3)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_78_v6).dtor_cf0)[new BigNumber(4)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_78_v6).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_79_v7, _dafny.Seq.of(new BigNumber(936)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_80_v8).dtor_cf2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_81_v9).f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_81_v9).f3).dtor_cf0)[_dafny.ZERO]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_81_v9).f3).dtor_cf0)[_dafny.ONE]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_81_v9).f3).dtor_cf0)[new BigNumber(2)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_81_v9).f3).dtor_cf0)[new BigNumber(3)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((((_81_v9).f3).dtor_cf0)[new BigNumber(4)]).equals(_dafny.Set.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_81_v9).f3).dtor_cf1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_83_v10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_84_v11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_85_v12).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(936),new _dafny.CodePoint('i'.codePointAt(0))).updateUnsafe(new BigNumber(3),new _dafny.CodePoint('i'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_86_v13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_89_v14).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(548)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_90_v15).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(548)),_dafny.Seq.UnicodeFromString("dtn")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_91_v16).dtor_cf85).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(548)),_dafny.Seq.UnicodeFromString("dtn")))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_187_v89).dtor_cf94));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v112).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(936),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_221_i9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_243_v120).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_244_v121).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[_dafny.ZERO]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[_dafny.ONE]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(2)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(3)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(4)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(5)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(6)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(7)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(8)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(9)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(10)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(11)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(12)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(13)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(14)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_245_v122)[new BigNumber(15)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_246_v123).equals(_dafny.Set.fromElements(false, true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_247_v124).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_248_v125).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(341),_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_249_v126, _dafny.Seq.of(false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_250_v127, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)), _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)), _dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(false, true)).updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(17)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(18)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(19)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(20)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(21)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(22)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)).updateUnsafe(true,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(23)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(24)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(25)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(26)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_251_v128)[new BigNumber(27)]).equals(_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Set.fromElements(false, true)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_252_v129));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v130)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v130)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_253_v130)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_254_v131).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_255_v132).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_322_i13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0, cf1) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      $dt.cf1 = cf1;
      return $dt;
    }
    static create_DC1(cf2) {
      let $dt = new D0(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get is_DC1() { return this.$tag === 1; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ", " + _dafny.toString(this.cf1) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf0 === other.cf0 && _dafny.areEqual(this.cf1, other.cf1);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf2, other.cf2);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC0([], _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC2(cf3) {
      let $dt = new D1(0);
      $dt.cf3 = cf3;
      return $dt;
    }
    get is_DC2() { return this.$tag === 0; }
    get dtor_cf3() { return this.cf3; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC2" + "(" + _dafny.toString(this.cf3) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf3 === other.cf3;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return false;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D1.Default();
        }
      };
    }
  }

  $module.D2 = class D2 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC4(cf5, cf6, cf7) {
      let $dt = new D2(0);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      return $dt;
    }
    static create_DC5(cf8, cf9, cf10) {
      let $dt = new D2(1);
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      return $dt;
    }
    static create_DC3(cf4) {
      let $dt = new D2(2);
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC6(cf11) {
      let $dt = new D2(3);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC4() { return this.$tag === 0; }
    get is_DC5() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC6() { return this.$tag === 3; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC4" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC5" + "(" + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ")";
      } else if (this.$tag === 2) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 3) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf11) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6 && this.cf7 === other.cf7;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9) && this.cf10 === other.cf10;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC4(_dafny.ZERO, false, []);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC7(cf12) {
      let $dt = new D3(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC7() { return this.$tag === 0; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC7" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf12, other.cf12);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Map.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D3.Default();
        }
      };
    }
  }

  $module.D4 = class D4 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC9(cf14, cf15, cf16, cf17, cf18) {
      let $dt = new D4(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC10(cf19, cf20, cf21) {
      let $dt = new D4(1);
      $dt.cf19 = cf19;
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC11(cf22, cf23, cf24, cf25, cf26) {
      let $dt = new D4(2);
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC8(cf13) {
      let $dt = new D4(3);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC12(cf27) {
      let $dt = new D4(4);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC9() { return this.$tag === 0; }
    get is_DC10() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get is_DC8() { return this.$tag === 3; }
    get is_DC12() { return this.$tag === 4; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC9" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf19) + ", " + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf22) + ", " + _dafny.toString(this.cf23) + ", " + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ", " + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 3) {
        return "D4.DC8" + "(" + _dafny.toString(this.cf13) + ")";
      } else if (this.$tag === 4) {
        return "D4.DC12" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && _dafny.areEqual(this.cf15, other.cf15) && _dafny.areEqual(this.cf16, other.cf16) && _dafny.areEqual(this.cf17, other.cf17) && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19) && _dafny.areEqual(this.cf20, other.cf20) && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22) && this.cf23 === other.cf23 && this.cf24 === other.cf24 && this.cf25 === other.cf25 && this.cf26 === other.cf26;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC9(_dafny.MultiSet.Empty, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, _dafny.Set.Empty, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC13(cf28) {
      let $dt = new D5(0);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC13() { return this.$tag === 0; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D5.Default();
        }
      };
    }
  }

  $module.D6 = class D6 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC15(cf30, cf31, cf32) {
      let $dt = new D6(0);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC14(cf29) {
      let $dt = new D6(1);
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC16(cf33) {
      let $dt = new D6(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf33() { return this.cf33; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC14" + "(" + this.cf29.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf33) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf30, other.cf30) && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf33, other.cf33);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC15(_dafny.ZERO, _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D6.Default();
        }
      };
    }
  }

  $module.D7 = class D7 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC18() {
      let $dt = new D7(0);
      return $dt;
    }
    static create_DC19(cf35) {
      let $dt = new D7(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC17(cf34) {
      let $dt = new D7(2);
      $dt.cf34 = cf34;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf34() { return this.cf34; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18";
      } else if (this.$tag === 1) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf34) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf34, other.cf34);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC18();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D7.Default();
        }
      };
    }
  }

  $module.D8 = class D8 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC20(cf36) {
      let $dt = new D8(0);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC20() { return this.$tag === 0; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Set.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC22(cf38) {
      let $dt = new D9(0);
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC21(cf37) {
      let $dt = new D9(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get is_DC21() { return this.$tag === 1; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf37() { return this.cf37; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC21" + "(" + _dafny.toString(this.cf37) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf37, other.cf37);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC22(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf39) {
      let $dt = new D10(0);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf39 === other.cf39;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return [];
    }
    static Rtd() {
      return class {
        static get Default() {
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC25(cf41, cf42, cf43) {
      let $dt = new D11(0);
      $dt.cf41 = cf41;
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC24(cf40) {
      let $dt = new D11(1);
      $dt.cf40 = cf40;
      return $dt;
    }
    static create_DC26(cf44) {
      let $dt = new D11(2);
      $dt.cf44 = cf44;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC26() { return this.$tag === 2; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf44() { return this.cf44; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf41) + ", " + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf40) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf44) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf41, other.cf41) && _dafny.areEqual(this.cf42, other.cf42) && this.cf43 === other.cf43;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf40, other.cf40);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf44, other.cf44);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC25(new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D11.Default();
        }
      };
    }
  }

  $module.D12 = class D12 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC28(cf46, cf47) {
      let $dt = new D12(0);
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      return $dt;
    }
    static create_DC29(cf48, cf49, cf50, cf51, cf52) {
      let $dt = new D12(1);
      $dt.cf48 = cf48;
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      $dt.cf52 = cf52;
      return $dt;
    }
    static create_DC27(cf45) {
      let $dt = new D12(2);
      $dt.cf45 = cf45;
      return $dt;
    }
    get is_DC28() { return this.$tag === 0; }
    get is_DC29() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf45() { return this.cf45; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC29" + "(" + _dafny.toString(this.cf48) + ", " + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ", " + _dafny.toString(this.cf52) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC27" + "(" + _dafny.toString(this.cf45) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf48 === other.cf48 && _dafny.areEqual(this.cf49, other.cf49) && this.cf50 === other.cf50 && _dafny.areEqual(this.cf51, other.cf51) && this.cf52 === other.cf52;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf45, other.cf45);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC28(false, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D12.Default();
        }
      };
    }
  }

  $module.D13 = class D13 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC31(cf54, cf55, cf56) {
      let $dt = new D13(0);
      $dt.cf54 = cf54;
      $dt.cf55 = cf55;
      $dt.cf56 = cf56;
      return $dt;
    }
    static create_DC32(cf57, cf58, cf59, cf60) {
      let $dt = new D13(1);
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC30(cf53) {
      let $dt = new D13(2);
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC33(cf61) {
      let $dt = new D13(3);
      $dt.cf61 = cf61;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get is_DC32() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get is_DC33() { return this.$tag === 3; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf55() { return this.cf55; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf61() { return this.cf61; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC31" + "(" + _dafny.toString(this.cf54) + ", " + _dafny.toString(this.cf55) + ", " + _dafny.toString(this.cf56) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC32" + "(" + this.cf57.toVerbatimString(true) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf61) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf54 === other.cf54 && _dafny.areEqual(this.cf55, other.cf55) && _dafny.areEqual(this.cf56, other.cf56);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf57, other.cf57) && this.cf58 === other.cf58 && _dafny.areEqual(this.cf59, other.cf59) && this.cf60 === other.cf60;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf53 === other.cf53;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf61, other.cf61);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC31(false, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D13.Default();
        }
      };
    }
  }

  $module.D14 = class D14 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC35(cf63, cf64, cf65, cf66, cf67) {
      let $dt = new D14(0);
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC36(cf68, cf69) {
      let $dt = new D14(1);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC34(cf62) {
      let $dt = new D14(2);
      $dt.cf62 = cf62;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC34() { return this.$tag === 2; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf62() { return this.cf62; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC35" + "(" + _dafny.toString(this.cf63) + ", " + this.cf64.toVerbatimString(true) + ", " + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 2) {
        return "D14.DC34" + "(" + _dafny.toString(this.cf62) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf63, other.cf63) && _dafny.areEqual(this.cf64, other.cf64) && _dafny.areEqual(this.cf65, other.cf65) && this.cf66 === other.cf66 && this.cf67 === other.cf67;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf68, other.cf68) && _dafny.areEqual(this.cf69, other.cf69);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf62, other.cf62);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC35(_dafny.ZERO, _dafny.Seq.UnicodeFromString(""), _dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D14.Default();
        }
      };
    }
  }

  $module.D15 = class D15 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC38(cf71) {
      let $dt = new D15(0);
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC39(cf72, cf73) {
      let $dt = new D15(1);
      $dt.cf72 = cf72;
      $dt.cf73 = cf73;
      return $dt;
    }
    static create_DC37(cf70) {
      let $dt = new D15(2);
      $dt.cf70 = cf70;
      return $dt;
    }
    static create_DC40(cf74) {
      let $dt = new D15(3);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC38() { return this.$tag === 0; }
    get is_DC39() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get is_DC40() { return this.$tag === 3; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf72() { return this.cf72; }
    get dtor_cf73() { return this.cf73; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC38" + "(" + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC39" + "(" + this.cf72.toVerbatimString(true) + ", " + _dafny.toString(this.cf73) + ")";
      } else if (this.$tag === 2) {
        return "D15.DC37" + "(" + _dafny.toString(this.cf70) + ")";
      } else if (this.$tag === 3) {
        return "D15.DC40" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf72, other.cf72) && _dafny.areEqual(this.cf73, other.cf73);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf70, other.cf70);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf74, other.cf74);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC38(new _dafny.CodePoint('D'.codePointAt(0)));
    }
    static Rtd() {
      return class {
        static get Default() {
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC42(cf76, cf77, cf78) {
      let $dt = new D16(0);
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC43(cf79, cf80, cf81) {
      let $dt = new D16(1);
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC44(cf82) {
      let $dt = new D16(2);
      $dt.cf82 = cf82;
      return $dt;
    }
    static create_DC41(cf75) {
      let $dt = new D16(3);
      $dt.cf75 = cf75;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get is_DC44() { return this.$tag === 2; }
    get is_DC41() { return this.$tag === 3; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf75() { return this.cf75; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC42" + "(" + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC43" + "(" + this.cf79.toVerbatimString(true) + ", " + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC44" + "(" + _dafny.toString(this.cf82) + ")";
      } else if (this.$tag === 3) {
        return "D16.DC41" + "(" + _dafny.toString(this.cf75) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf76, other.cf76) && _dafny.areEqual(this.cf77, other.cf77) && _dafny.areEqual(this.cf78, other.cf78);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf79, other.cf79) && _dafny.areEqual(this.cf80, other.cf80) && _dafny.areEqual(this.cf81, other.cf81);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf82, other.cf82);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf75 === other.cf75;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC42(_dafny.ZERO, _dafny.ZERO, _dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
        }
      };
    }
  }

  $module.D17 = class D17 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC46(cf84) {
      let $dt = new D17(0);
      $dt.cf84 = cf84;
      return $dt;
    }
    static create_DC45(cf83) {
      let $dt = new D17(1);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC46" + "(" + _dafny.toString(this.cf84) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC45" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf84, other.cf84);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf83, other.cf83);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC46(_dafny.ZERO);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D17.Default();
        }
      };
    }
  }

  $module.D18 = class D18 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC48() {
      let $dt = new D18(0);
      return $dt;
    }
    static create_DC49(cf86) {
      let $dt = new D18(1);
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC50(cf87, cf88, cf89, cf90, cf91) {
      let $dt = new D18(2);
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      return $dt;
    }
    static create_DC47(cf85) {
      let $dt = new D18(3);
      $dt.cf85 = cf85;
      return $dt;
    }
    get is_DC48() { return this.$tag === 0; }
    get is_DC49() { return this.$tag === 1; }
    get is_DC50() { return this.$tag === 2; }
    get is_DC47() { return this.$tag === 3; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf85() { return this.cf85; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC48";
      } else if (this.$tag === 1) {
        return "D18.DC49" + "(" + _dafny.toString(this.cf86) + ")";
      } else if (this.$tag === 2) {
        return "D18.DC50" + "(" + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ", " + this.cf91.toVerbatimString(true) + ")";
      } else if (this.$tag === 3) {
        return "D18.DC47" + "(" + _dafny.toString(this.cf85) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf86, other.cf86);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf87 === other.cf87 && this.cf88 === other.cf88 && this.cf89 === other.cf89 && this.cf90 === other.cf90 && _dafny.areEqual(this.cf91, other.cf91);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf85, other.cf85);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC48();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC51(cf92) {
      let $dt = new D19(0);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC51" + "(" + _dafny.toString(this.cf92) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf92 === other.cf92;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D19.Default();
        }
      };
    }
  }

  $module.D20 = class D20 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC53(cf94) {
      let $dt = new D20(0);
      $dt.cf94 = cf94;
      return $dt;
    }
    static create_DC54(cf95, cf96, cf97) {
      let $dt = new D20(1);
      $dt.cf95 = cf95;
      $dt.cf96 = cf96;
      $dt.cf97 = cf97;
      return $dt;
    }
    static create_DC52(cf93) {
      let $dt = new D20(2);
      $dt.cf93 = cf93;
      return $dt;
    }
    get is_DC53() { return this.$tag === 0; }
    get is_DC54() { return this.$tag === 1; }
    get is_DC52() { return this.$tag === 2; }
    get dtor_cf94() { return this.cf94; }
    get dtor_cf95() { return this.cf95; }
    get dtor_cf96() { return this.cf96; }
    get dtor_cf97() { return this.cf97; }
    get dtor_cf93() { return this.cf93; }
    toString() {
      if (this.$tag === 0) {
        return "D20.DC53" + "(" + _dafny.toString(this.cf94) + ")";
      } else if (this.$tag === 1) {
        return "D20.DC54" + "(" + this.cf95.toVerbatimString(true) + ", " + _dafny.toString(this.cf96) + ", " + _dafny.toString(this.cf97) + ")";
      } else if (this.$tag === 2) {
        return "D20.DC52" + "(" + _dafny.toString(this.cf93) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf94 === other.cf94;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf95, other.cf95) && _dafny.areEqual(this.cf96, other.cf96) && this.cf97 === other.cf97;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf93 === other.cf93;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D20.create_DC53(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D20.Default();
        }
      };
    }
  }

  $module.D21 = class D21 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC55(cf98) {
      let $dt = new D21(0);
      $dt.cf98 = cf98;
      return $dt;
    }
    get is_DC55() { return this.$tag === 0; }
    get dtor_cf98() { return this.cf98; }
    toString() {
      if (this.$tag === 0) {
        return "D21.DC55" + "(" + _dafny.toString(this.cf98) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf98 === other.cf98;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D21.Default();
        }
      };
    }
  }

  $module.D22 = class D22 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC56(cf99) {
      let $dt = new D22(0);
      $dt.cf99 = cf99;
      return $dt;
    }
    get is_DC56() { return this.$tag === 0; }
    get dtor_cf99() { return this.cf99; }
    toString() {
      if (this.$tag === 0) {
        return "D22.DC56" + "(" + _dafny.toString(this.cf99) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf99 === other.cf99;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return null;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D22.Default();
        }
      };
    }
  }

  $module.D23 = class D23 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC58() {
      let $dt = new D23(0);
      return $dt;
    }
    static create_DC57(cf100) {
      let $dt = new D23(1);
      $dt.cf100 = cf100;
      return $dt;
    }
    get is_DC58() { return this.$tag === 0; }
    get is_DC57() { return this.$tag === 1; }
    get dtor_cf100() { return this.cf100; }
    toString() {
      if (this.$tag === 0) {
        return "D23.DC58";
      } else if (this.$tag === 1) {
        return "D23.DC57" + "(" + _dafny.toString(this.cf100) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf100, other.cf100);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D23.create_DC58();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D23.Default();
        }
      };
    }
  }

  $module.D24 = class D24 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC60() {
      let $dt = new D24(0);
      return $dt;
    }
    static create_DC61(cf102, cf103) {
      let $dt = new D24(1);
      $dt.cf102 = cf102;
      $dt.cf103 = cf103;
      return $dt;
    }
    static create_DC62(cf104) {
      let $dt = new D24(2);
      $dt.cf104 = cf104;
      return $dt;
    }
    static create_DC59(cf101) {
      let $dt = new D24(3);
      $dt.cf101 = cf101;
      return $dt;
    }
    get is_DC60() { return this.$tag === 0; }
    get is_DC61() { return this.$tag === 1; }
    get is_DC62() { return this.$tag === 2; }
    get is_DC59() { return this.$tag === 3; }
    get dtor_cf102() { return this.cf102; }
    get dtor_cf103() { return this.cf103; }
    get dtor_cf104() { return this.cf104; }
    get dtor_cf101() { return this.cf101; }
    toString() {
      if (this.$tag === 0) {
        return "D24.DC60";
      } else if (this.$tag === 1) {
        return "D24.DC61" + "(" + _dafny.toString(this.cf102) + ", " + _dafny.toString(this.cf103) + ")";
      } else if (this.$tag === 2) {
        return "D24.DC62" + "(" + _dafny.toString(this.cf104) + ")";
      } else if (this.$tag === 3) {
        return "D24.DC59" + "(" + _dafny.toString(this.cf101) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf102, other.cf102) && _dafny.areEqual(this.cf103, other.cf103);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf104, other.cf104);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf101 === other.cf101;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D24.create_DC60();
    }
    static Rtd() {
      return class {
        static get Default() {
          return D24.Default();
        }
      };
    }
  }

  $module.D25 = class D25 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC64(cf106) {
      let $dt = new D25(0);
      $dt.cf106 = cf106;
      return $dt;
    }
    static create_DC65(cf107, cf108, cf109) {
      let $dt = new D25(1);
      $dt.cf107 = cf107;
      $dt.cf108 = cf108;
      $dt.cf109 = cf109;
      return $dt;
    }
    static create_DC63(cf105) {
      let $dt = new D25(2);
      $dt.cf105 = cf105;
      return $dt;
    }
    static create_DC66(cf110) {
      let $dt = new D25(3);
      $dt.cf110 = cf110;
      return $dt;
    }
    get is_DC64() { return this.$tag === 0; }
    get is_DC65() { return this.$tag === 1; }
    get is_DC63() { return this.$tag === 2; }
    get is_DC66() { return this.$tag === 3; }
    get dtor_cf106() { return this.cf106; }
    get dtor_cf107() { return this.cf107; }
    get dtor_cf108() { return this.cf108; }
    get dtor_cf109() { return this.cf109; }
    get dtor_cf105() { return this.cf105; }
    get dtor_cf110() { return this.cf110; }
    toString() {
      if (this.$tag === 0) {
        return "D25.DC64" + "(" + _dafny.toString(this.cf106) + ")";
      } else if (this.$tag === 1) {
        return "D25.DC65" + "(" + _dafny.toString(this.cf107) + ", " + _dafny.toString(this.cf108) + ", " + _dafny.toString(this.cf109) + ")";
      } else if (this.$tag === 2) {
        return "D25.DC63" + "(" + _dafny.toString(this.cf105) + ")";
      } else if (this.$tag === 3) {
        return "D25.DC66" + "(" + _dafny.toString(this.cf110) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf106, other.cf106);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf107, other.cf107) && this.cf108 === other.cf108 && _dafny.areEqual(this.cf109, other.cf109);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf105, other.cf105);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf110, other.cf110);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D25.create_DC64(_dafny.MultiSet.Empty);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D25.Default();
        }
      };
    }
  }

  $module.D26 = class D26 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC68(cf112) {
      let $dt = new D26(0);
      $dt.cf112 = cf112;
      return $dt;
    }
    static create_DC69(cf113, cf114) {
      let $dt = new D26(1);
      $dt.cf113 = cf113;
      $dt.cf114 = cf114;
      return $dt;
    }
    static create_DC67(cf111) {
      let $dt = new D26(2);
      $dt.cf111 = cf111;
      return $dt;
    }
    get is_DC68() { return this.$tag === 0; }
    get is_DC69() { return this.$tag === 1; }
    get is_DC67() { return this.$tag === 2; }
    get dtor_cf112() { return this.cf112; }
    get dtor_cf113() { return this.cf113; }
    get dtor_cf114() { return this.cf114; }
    get dtor_cf111() { return this.cf111; }
    toString() {
      if (this.$tag === 0) {
        return "D26.DC68" + "(" + _dafny.toString(this.cf112) + ")";
      } else if (this.$tag === 1) {
        return "D26.DC69" + "(" + this.cf113.toVerbatimString(true) + ", " + _dafny.toString(this.cf114) + ")";
      } else if (this.$tag === 2) {
        return "D26.DC67" + "(" + _dafny.toString(this.cf111) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf112 === other.cf112;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf113, other.cf113) && this.cf114 === other.cf114;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf111, other.cf111);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D26.create_DC68(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D26.Default();
        }
      };
    }
  }

  $module.D27 = class D27 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC70(cf115) {
      let $dt = new D27(0);
      $dt.cf115 = cf115;
      return $dt;
    }
    get is_DC70() { return this.$tag === 0; }
    get dtor_cf115() { return this.cf115; }
    toString() {
      if (this.$tag === 0) {
        return "D27.DC70" + "(" + _dafny.toString(this.cf115) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf115, other.cf115);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.MultiSet.Empty;
    }
    static Rtd() {
      return class {
        static get Default() {
          return D27.Default();
        }
      };
    }
  }

  $module.D28 = class D28 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC72(cf117, cf118, cf119, cf120) {
      let $dt = new D28(0);
      $dt.cf117 = cf117;
      $dt.cf118 = cf118;
      $dt.cf119 = cf119;
      $dt.cf120 = cf120;
      return $dt;
    }
    static create_DC71(cf116) {
      let $dt = new D28(1);
      $dt.cf116 = cf116;
      return $dt;
    }
    get is_DC72() { return this.$tag === 0; }
    get is_DC71() { return this.$tag === 1; }
    get dtor_cf117() { return this.cf117; }
    get dtor_cf118() { return this.cf118; }
    get dtor_cf119() { return this.cf119; }
    get dtor_cf120() { return this.cf120; }
    get dtor_cf116() { return this.cf116; }
    toString() {
      if (this.$tag === 0) {
        return "D28.DC72" + "(" + _dafny.toString(this.cf117) + ", " + _dafny.toString(this.cf118) + ", " + _dafny.toString(this.cf119) + ", " + _dafny.toString(this.cf120) + ")";
      } else if (this.$tag === 1) {
        return "D28.DC71" + "(" + _dafny.toString(this.cf116) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf117, other.cf117) && _dafny.areEqual(this.cf118, other.cf118) && this.cf119 === other.cf119 && this.cf120 === other.cf120;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf116, other.cf116);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D28.create_DC72(_dafny.ZERO, _dafny.ZERO, false, false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D28.Default();
        }
      };
    }
  }

  $module.T0 = class T0 {
  };

  $module.T1 = class T1 {
  };

  $module.GlobalState = class GlobalState {
    constructor () {
      this._tname = "_module.GlobalState";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f8 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f8) {
      let _this = this;
      (_this).f8 = f8;
      return;
    }
    fm14(globalState) {
      let _this = this;
      return new _dafny.CodePoint('w'.codePointAt(0));
    };
    fm15(p0, p1, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(5)), function (_329_i0) {
        return new _dafny.CodePoint('u'.codePointAt(0));
      }), _this.f8);
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f1 = _module.D0.Default();
      this._f0 = _dafny.Seq.of();
      this.f9 = _dafny.ZERO;
      this._f10 = [];
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    set f1(value) {
      let _this = this;
      _this._f1 = value;
    };
    get f0() {
      let _this = this;
      return _this._f0;
    };
    __ctor(f9, f10, f0, f1) {
      let _this = this;
      (_this).f9 = f9;
      (_this)._f10 = f10;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      return;
    }
    fm6(p0, p1, globalState) {
      let _this = this;
      return new BigNumber((_dafny.MultiSet.fromElements(_module.__default.safeDivisionInt(new BigNumber(603), _this.f9), _this.f9, _this.f9, new BigNumber(((function () {
        let _coll33 = new _dafny.Set();
        for (const _compr_33 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("drwqyvqsw")).length))).Elements) {
          let _330_v0 = _compr_33;
          if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("drwqyvqsw")).length))).contains(_330_v0)) {
            _coll33.add((_330_v0).plus(new BigNumber(613)));
          }
        }
        return _coll33;
      }()).Difference(_dafny.Set.fromElements((_dafny.ZERO).minus(_this.f9), _this.f9, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(true))).cardinality())))).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(797)), function (_331_i0) {
        return new _dafny.CodePoint('o'.codePointAt(0));
      }), _dafny.Seq.of(new _dafny.CodePoint('h'.codePointAt(0)), new _dafny.CodePoint('c'.codePointAt(0))))).length))).cardinality());
    };
    fm27(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(!(!(false))), _dafny.Seq.of(true)), _dafny.Seq.Concat(_dafny.Seq.of(false), _dafny.Seq.of(false))));
    };
    fm28(p0, p1, globalState) {
      let _this = this;
      return _this.f9;
    };
    m13(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      (_this).f9 = _this.f9;
      r0 = p1;
      if (p1) {
        let _332_v0;
        _332_v0 = _dafny.Seq.UnicodeFromString("dmxvlol");
        let _333_v1;
        let _nw50 = new _module.C0();
        _nw50.__ctor(_332_v0);
        _333_v1 = _nw50;
        let _334_v2;
        let _nw51 = new _module.C0();
        _nw51.__ctor(_dafny.Seq.Concat(_332_v0, _333_v1.f8));
        _334_v2 = _nw51;
        let _rhs31 = true;
        let _rhs32 = p1;
        let _rhs33 = _module.__default.safeDivisionInt((_this.f9).minus(_this.f9), (_this.f9).minus(_this.f9));
        let _lhs16 = _this;
        r0 = _rhs31;
        r0 = _rhs32;
        _lhs16.f9 = _rhs33;
        let _335_v3;
        _335_v3 = _dafny.MultiSet.fromElements(p1, p1);
        let _336_v4;
        _336_v4 = _dafny.Seq.of(p1);
        let _337_v5;
        _337_v5 = _module.D7.create_DC17(_336_v4);
        let _338_v6;
        _338_v6 = _dafny.MultiSet.fromElements(_dafny.areEqual(_332_v0, _334_v2.f8), p1, p1, !((_dafny.MultiSet.FromArray((_337_v5).dtor_cf34)).IsProperSubsetOf(_335_v3)));
        _335_v3 = ((_dafny.MultiSet.fromElements(p1)).Union(_dafny.MultiSet.FromArray(_dafny.Seq.of(p1, p1)))).Difference(_335_v3);
        let _339_v7;
        let _nw52 = new _module.C0();
        _nw52.__ctor(_dafny.Seq.Concat(_332_v0, _333_v1.f8));
        _339_v7 = _nw52;
      } else {
        let _340_v8;
        let _nw53 = Array((new BigNumber(29)).toNumber()).fill(false);
        _340_v8 = _nw53;
        let _index30 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_340_v8).length));
        (_340_v8)[_index30] = p1;
        let _index31 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_340_v8).length));
        (_340_v8)[_index31] = p1;
        let _341_v9;
        let _nw54 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _341_v9 = _nw54;
        let _index32 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_341_v9).length));
        (_341_v9)[_index32] = _this.f9;
        let _index33 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_341_v9).length));
        (_341_v9)[_index33] = _module.__default.safeModuloInt(_this.f9, _this.f9);
        _340_v8 = _340_v8;
        let _342_v10;
        let _init3 = ((_343_p0, _344_p1) => function (_345_i0) {
          return (_module.D11.create_DC25(_343_p0, new BigNumber((_dafny.MultiSet.fromElements(_module.D6.create_DC14(_dafny.Seq.UnicodeFromString("pqjlcv")), _module.D6.create_DC14(_dafny.Seq.UnicodeFromString("lcnl")))).cardinality()), _344_p1)).dtor_cf41;
        })(p0, p1);
        let _nw55 = Array((new BigNumber(2)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw55.length); _i0_3++) {
          _nw55[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _342_v10 = _nw55;
        let _index34 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_342_v10).length));
        (_342_v10)[_index34] = _module.__default.fm29((_340_v8)[_module.__default.safeIndex(new BigNumber(517), new BigNumber((_340_v8).length))], globalState);
        let _index35 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_342_v10).length));
        let _rhs34 = p0;
        let _rhs35 = (_340_v8)[_module.__default.safeIndex(new BigNumber(517), new BigNumber((_340_v8).length))];
        let _lhs17 = _342_v10;
        let _lhs18 = _module.__default.safeIndex(new BigNumber(801), new BigNumber((_342_v10).length));
        _lhs17[_lhs18] = _rhs34;
        r0 = _rhs35;
        let _index36 = _module.__default.safeIndex(new BigNumber(609), new BigNumber((_341_v9).length));
        (_341_v9)[_index36] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((new BigNumber(-65)).plus(_this.f9), ((_341_v9)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_341_v9).length))]).multipliedBy((_341_v9)[_module.__default.safeIndex(new BigNumber(609), new BigNumber((_341_v9).length))])));
      }
      let _346_v11;
      let _init4 = ((_347_p0) => function (_348_i2) {
        return _347_p0;
      })(p0);
      let _nw56 = Array((new BigNumber(23)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw56.length); _i0_4++) {
        _nw56[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _346_v11 = _nw56;
      let _349_v12;
      _349_v12 = _dafny.Seq.of(_346_v11, _346_v11, _346_v11);
      let _hi0 = new BigNumber((_349_v12).length);
      for (let _350_i1 = _this.f9; _350_i1.isLessThan(_hi0); _350_i1 = _350_i1.plus(_dafny.ONE)) {
        let _351_v13;
        let _nw57 = Array((new BigNumber(13)).toNumber()).fill([]);
        _351_v13 = _nw57;
        let _352_v14;
        _352_v14 = _module.D2.create_DC4(_this.f9, p1, _351_v13);
        let _353_v15;
        _353_v15 = _dafny.MultiSet.fromElements((_this.f9).multipliedBy((_352_v14).dtor_cf5));
        _353_v15 = _353_v15;
        let _354_v16;
        let _init5 = ((_355_p1) => function (_356_i3) {
          return (_355_p1) && (_355_p1);
        })(p1);
        let _nw58 = Array((new BigNumber(14)).toNumber());
        for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw58.length); _i0_5++) {
          _nw58[_i0_5] = _init5(new BigNumber(_i0_5));
        }
        _354_v16 = _nw58;
        let _index37 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_354_v16).length));
        (_354_v16)[_index37] = (p1) && (p1);
        let _index38 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_354_v16).length));
        (_354_v16)[_index38] = p1;
        let _357_v17;
        let _nw59 = Array((new BigNumber(17)).toNumber()).fill(_module.D6.Default());
        _357_v17 = _nw59;
        let _358_v18;
        _358_v18 = _dafny.Seq.of(p1);
        let _359_v19;
        _359_v19 = _module.D6.create_DC15(_350_i1, _this.f9, _module.__default.fm30(new BigNumber((_358_v18).length), (_354_v16)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_354_v16).length))], (_354_v16)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_354_v16).length))], globalState));
        let _index39 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_357_v17).length));
        (_357_v17)[_index39] = _359_v19;
        let _360_v20;
        _360_v20 = _dafny.Set.fromElements(_this.f9);
        let _361_v21;
        let _nw60 = Array((new BigNumber(4)).toNumber());
        _nw60[(_dafny.ZERO).toNumber()] = _dafny.Set.fromElements(_this.f9);
        _nw60[(_dafny.ONE).toNumber()] = (_360_v20).Difference(_360_v20);
        _nw60[(new BigNumber(2)).toNumber()] = _360_v20;
        _nw60[(new BigNumber(3)).toNumber()] = _360_v20;
        _361_v21 = _nw60;
        let _362_v22;
        let _init6 = ((_363_i1) => function (_364_i4) {
          return _module.__default.safeDivisionInt(_364_i4, _363_i1);
        })(_350_i1);
        let _nw61 = Array((new BigNumber(6)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw61.length); _i0_6++) {
          _nw61[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _362_v22 = _nw61;
        let _index40 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_362_v22).length));
        (_362_v22)[_index40] = _module.__default.safeModuloInt(_this.f9, _this.f9);
        let _index41 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_357_v17).length));
        let _index42 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_354_v16).length));
        let _index43 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_362_v22).length));
        let _rhs36 = _359_v19;
        let _rhs37 = _361_v21;
        let _rhs38 = (_354_v16)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_354_v16).length))];
        let _rhs39 = ((p1) ? ((_dafny.ZERO).minus(_350_i1)) : ((_this.f9).minus(_this.f9)));
        let _lhs19 = _357_v17;
        let _lhs20 = _module.__default.safeIndex(new BigNumber(194), new BigNumber((_357_v17).length));
        let _lhs21 = _354_v16;
        let _lhs22 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_354_v16).length));
        let _lhs23 = _362_v22;
        let _lhs24 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_362_v22).length));
        _lhs19[_lhs20] = _rhs36;
        _361_v21 = _rhs37;
        _lhs21[_lhs22] = _rhs38;
        _lhs23[_lhs24] = _rhs39;
        let _365_v23;
        _365_v23 = _dafny.Seq.UnicodeFromString("rarjd");
        let _366_v24;
        _366_v24 = _module.D6.create_DC14(_dafny.Seq.Concat(_365_v23, _dafny.Seq.UnicodeFromString("jmypydeh")));
        let _pat_let_tv8 = _365_v23;
        let _index44 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_362_v22).length));
        let _rhs40 = function (_pat_let10_0) {
          return function (_367_dt__update__tmp_h0) {
            return function (_pat_let11_0) {
              return function (_368_dt__update_hcf29_h0) {
                return _module.D6.create_DC14(_368_dt__update_hcf29_h0);
              }(_pat_let11_0);
            }(_pat_let_tv8);
          }(_pat_let10_0);
        }(_module.D6.create_DC14(_365_v23));
        let _rhs41 = _dafny.Seq.of((_354_v16)[_module.__default.safeIndex(new BigNumber(887), new BigNumber((_354_v16).length))], p1, p1);
        let _rhs42 = _this.f9;
        let _lhs25 = _362_v22;
        let _lhs26 = _module.__default.safeIndex(new BigNumber(256), new BigNumber((_362_v22).length));
        _366_v24 = _rhs40;
        _358_v18 = _rhs41;
        _lhs25[_lhs26] = _rhs42;
      }
      let _hi1 = _this.f9;
      for (let _369_i5 = _this.f9; _369_i5.isLessThan(_hi1); _369_i5 = _369_i5.plus(_dafny.ONE)) {
        let _370_v25;
        let _nw62 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Seq.of());
        _370_v25 = _nw62;
        let _371_v26;
        _371_v26 = _dafny.Seq.of(p1, p1);
        let _index45 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_370_v25).length));
        (_370_v25)[_index45] = _dafny.Seq.Concat(_371_v26, _dafny.Seq.update(_371_v26, _module.__default.safeIndex(_this.f9, new BigNumber((_371_v26).length)), p1));
        let _372_v27;
        _372_v27 = _module.D7.create_DC17(_371_v26);
        let _373_v28;
        _373_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,_371_v26);
        let _index46 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_370_v25).length));
        (_370_v25)[_index46] = _dafny.Seq.update((((p1) ? (_372_v27) : (_module.D7.create_DC17((((_373_v28).contains(p1)) ? ((_373_v28).get(p1)) : (_371_v26)))))).dtor_cf34, _module.__default.safeIndex(_this.f9, new BigNumber(((((p1) ? (_372_v27) : (_module.D7.create_DC17((((_373_v28).contains(p1)) ? ((_373_v28).get(p1)) : (_371_v26)))))).dtor_cf34).length)), (_module.D12.create_DC29(p1, _369_i5, p1, new BigNumber(897), p1)).dtor_cf50);
        let _374_v29;
        _374_v29 = _dafny.Seq.UnicodeFromString("nyg");
        _374_v29 = _dafny.Seq.Concat(_374_v29, _374_v29);
        (_this).f9 = _module.__default.safeDivisionInt(_this.f9, _this.f9);
        let _375_v30;
        _375_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_374_v29).length),p1);
        let _376_v31;
        let _nw63 = Array((new BigNumber(12)).toNumber()).fill(false);
        _376_v31 = _nw63;
        let _377_v32;
        _377_v32 = _dafny.MultiSet.fromElements(_376_v31, _376_v31);
        let _378_v33;
        _378_v33 = _dafny.MultiSet.fromElements(p0);
        let _379_v34;
        _379_v34 = _dafny.MultiSet.fromElements(_369_i5, _369_i5, new BigNumber((_378_v33).cardinality()));
        let _380_v35;
        _380_v35 = _dafny.Set.fromElements(_dafny.MultiSet.FromArray(_dafny.Seq.of(_369_i5)), _379_v34, _dafny.MultiSet.FromArray((_this).f0), _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm31(_this.f9, p1, _this.f9, globalState)).length)));
        let _381_v36;
        let _nw64 = Array((new BigNumber(22)).toNumber());
        _nw64[(_dafny.ZERO).toNumber()] = ((_dafny.ZERO).minus(_369_i5)).minus(_369_i5);
        _nw64[(_dafny.ONE).toNumber()] = new BigNumber((_375_v30).length);
        _nw64[(new BigNumber(2)).toNumber()] = (_this.f9).plus(_369_i5);
        _nw64[(new BigNumber(3)).toNumber()] = _369_i5;
        _nw64[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus(_369_i5), _369_i5);
        _nw64[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(_this.f9, _369_i5);
        _nw64[(new BigNumber(6)).toNumber()] = _369_i5;
        _nw64[(new BigNumber(7)).toNumber()] = (_this.f9).multipliedBy(_this.f9);
        _nw64[(new BigNumber(8)).toNumber()] = new BigNumber(((_this).f0).length);
        _nw64[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((_this).fm28(new BigNumber(-80), p1, globalState));
        _nw64[(new BigNumber(10)).toNumber()] = (((_377_v32).contains(_376_v31)) ? ((_377_v32).get(_376_v31)) : (new BigNumber((_380_v35).length)));
        _nw64[(new BigNumber(11)).toNumber()] = _this.f9;
        _nw64[(new BigNumber(12)).toNumber()] = _369_i5;
        _nw64[(new BigNumber(13)).toNumber()] = _369_i5;
        _nw64[(new BigNumber(14)).toNumber()] = _this.f9;
        _nw64[(new BigNumber(15)).toNumber()] = new BigNumber(((_this).f0).length);
        _nw64[(new BigNumber(16)).toNumber()] = _369_i5;
        _nw64[(new BigNumber(17)).toNumber()] = (_369_i5).multipliedBy((_dafny.ZERO).minus(new BigNumber((_module.__default.fm31(_this.f9, p1, (_dafny.ZERO).minus(new BigNumber((_374_v29).length)), globalState)).length)));
        _nw64[(new BigNumber(18)).toNumber()] = (_this.f9).plus(new BigNumber(345));
        _nw64[(new BigNumber(19)).toNumber()] = _369_i5;
        _nw64[(new BigNumber(20)).toNumber()] = _this.f9;
        _nw64[(new BigNumber(21)).toNumber()] = _this.f9;
        _381_v36 = _nw64;
        let _index47 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_381_v36).length));
        (_381_v36)[_index47] = new BigNumber(882);
        let _index48 = _module.__default.safeIndex(new BigNumber(48), new BigNumber((_381_v36).length));
        (_381_v36)[_index48] = (((_377_v32).contains(_376_v31)) ? ((_377_v32).get(_376_v31)) : (_this.f9));
      }
      let _382_v37;
      _382_v37 = _module.D12.create_DC29(p1, (_this).fm28(_this.f9, !(!(p1)), globalState), p1, _this.f9, !(p1));
      let _383_v38;
      _383_v38 = _dafny.Map.Empty.slice().updateUnsafe((_382_v37).dtor_cf49,_this.f9);
      let _384_v39;
      _384_v39 = _dafny.Map.Empty.slice().updateUnsafe(_this.f9,p1);
      let _385_v40;
      _385_v40 = _module.D11.create_DC25(p0, new BigNumber(390), (((_384_v39).contains(_this.f9)) ? ((_384_v39).get(_this.f9)) : (p1)));
      let _386_v41;
      _386_v41 = _dafny.Seq.UnicodeFromString("bhh");
      let _hi2 = new BigNumber((_386_v41).length);
      for (let _387_i6 = (_dafny.ZERO).minus((((_383_v38).contains((_385_v40).dtor_cf42)) ? ((_383_v38).get((_385_v40).dtor_cf42)) : (new BigNumber((_dafny.MultiSet.fromElements(p0)).cardinality())))); _387_i6.isLessThan(_hi2); _387_i6 = _387_i6.plus(_dafny.ONE)) {
        r0 = ((p1) ? (!(p1)) : (p1));
        let _388_v42;
        _388_v42 = _dafny.MultiSet.fromElements(_387_i6, _387_i6, (_this).fm6(_this.f9, _this.f9, globalState), _387_i6, _this.f9);
        (_this).f9 = _module.__default.safeModuloInt((new BigNumber((_388_v42).cardinality())).plus(_387_i6), ((_this).fm6(_this.f9, _this.f9, globalState)).plus(_this.f9));
        let _389_v43;
        _389_v43 = new _dafny.CodePoint('q'.codePointAt(0));
        _389_v43 = _389_v43;
        let _390_v44;
        let _init7 = function (_391_i7) {
          return _dafny.Seq.UnicodeFromString("htrgdatxk");
        };
        let _nw65 = Array((new BigNumber(25)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw65.length); _i0_7++) {
          _nw65[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _390_v44 = _nw65;
        let _index49 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_390_v44).length));
        (_390_v44)[_index49] = _386_v41;
        let _index50 = _module.__default.safeIndex(new BigNumber(115), new BigNumber((_390_v44).length));
        (_390_v44)[_index50] = ((p1) ? (_386_v41) : (_386_v41));
      }
      r0 = ((_dafny.Seq.IsPrefixOf(_386_v41, _dafny.Seq.UnicodeFromString("ly"))) ? (!_dafny.Seq.contains(_module.__default.fm31(_this.f9, p1, new BigNumber((_dafny.MultiSet.fromElements(p0, p0, p0)).cardinality()), globalState), p0)) : (p1));
      return r0;
    }
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return false;
    };
    fm41(p0, p1, globalState) {
      let _this = this;
      return _dafny.MultiSet.fromElements(false);
    };
    fm42(p0, globalState) {
      let _this = this;
      return (new BigNumber((function () {
        let _coll34 = new _dafny.Map();
        for (const _compr_34 of _dafny.IntegerRange(new BigNumber(-558), new BigNumber(560))) {
          let _392_v0 = _compr_34;
          if (((new BigNumber(-558)).isLessThanOrEqualTo(_392_v0)) && ((_392_v0).isLessThan(new BigNumber(560)))) {
            _coll34.push([_module.__default.safeDivisionInt(_392_v0, new BigNumber((_dafny.Seq.of(false)).length)),_dafny.Seq.of(_module.D2.create_DC5(new BigNumber(-832), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-101)), function (_393_i0) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new _dafny.CodePoint('p'.codePointAt(0)))).length);
}), false), _module.D2.create_DC5(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("tam")).length),new _dafny.CodePoint('k'.codePointAt(0)))).length), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, !(false), !(true))).length), new BigNumber((_dafny.Set.fromElements(true)).length)), true))]);
          }
        }
        return _coll34;
      }()).length)).isLessThan(new BigNumber(281));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _394_v0;
      _394_v0 = _dafny.MultiSet.fromElements(p0, p0);
      let _395_v1;
      _395_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_394_v0).cardinality()),p0);
      let _396_v2;
      let _nw66 = Array((_dafny.ONE).toNumber()).fill(_dafny.Map.Empty);
      _396_v2 = _nw66;
      let _397_v3;
      _397_v3 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("kpmro")).length));
      let _398_v4;
      let _nw67 = new _module.C1();
      _nw67.__ctor(p0, _396_v2, _397_v3, _module.D0.create_DC1(p0));
      _398_v4 = _nw67;
      let _399_v5;
      _399_v5 = _dafny.Map.Empty.slice().updateUnsafe(_398_v4,p0);
      let _400_v6;
      let _nw68 = new _module.C0();
      _nw68.__ctor(_dafny.Seq.UnicodeFromString("us"));
      _400_v6 = _nw68;
      let _401_v7;
      _401_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,_400_v6);
      let _402_v8;
      _402_v8 = _dafny.Seq.of(_395_v1, _dafny.Map.Empty.slice().updateUnsafe(p0,p0), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_399_v5).length),new BigNumber((_401_v7).length)));
      let _403_v9;
      _403_v9 = new BigNumber(156);
      let _rhs43 = _dafny.Seq.Concat(_402_v8, _402_v8);
      let _rhs44 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-425)), ((_404_v9) => function (_405_i0) {
        return _404_v9;
      })(_403_v9))).length);
      let _rhs45 = _dafny.Seq.UnicodeFromString("fvuiwig");
      let _lhs27 = _400_v6;
      _402_v8 = _rhs43;
      _403_v9 = _rhs44;
      _lhs27.f8 = _rhs45;
      let _406_v10;
      let _init8 = function (_407_i2) {
        return (_dafny.Map.Empty.slice().updateUnsafe(!(false),_dafny.MultiSet.fromElements(false))).contains(false);
      };
      let _nw69 = Array((new BigNumber(17)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw69.length); _i0_8++) {
        _nw69[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _406_v10 = _nw69;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_406_v10).length))) {
        let _408_i1 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_408_i1)) && ((_408_i1).isLessThan(new BigNumber((_406_v10).length))))) {
          (_406_v10)[(_408_i1)] = (_403_v9).isEqualTo(((_398_v4).f0)[_module.__default.safeIndex(_403_v9, new BigNumber(((_398_v4).f0).length))]);
        }
      }
      let _409_v11;
      _409_v11 = true;
      let _410_v12;
      _410_v12 = _dafny.Map.Empty.slice().updateUnsafe(_409_v11,_403_v9);
      let _411_v13;
      _411_v13 = _dafny.Seq.of(_400_v6.f8, _400_v6.f8, _400_v6.f8);
      let _412_v14;
      _412_v14 = _dafny.Map.Empty.slice().updateUnsafe(_410_v12,(_411_v13)[_module.__default.safeIndex(new BigNumber(359), new BigNumber((_411_v13).length))]);
      let _413_v16;
      _413_v16 = _dafny.Seq.of(_410_v12, _410_v12);
      let _414_v17;
      let _out19;
      _out19 = (_this).m14(_409_v11, _409_v11, (_412_v14).Merge(function () {
        let _coll35 = new _dafny.Map();
        for (const _compr_35 of (_413_v16).Elements) {
          let _415_v15 = _compr_35;
          if (_dafny.Seq.contains(_413_v16, _415_v15)) {
            _coll35.push([_415_v15,_400_v6.f8]);
          }
        }
        return _coll35;
      }()), p0, globalState);
      _414_v17 = _out19;
      r0 = _409_v11;
      let _416_v18;
      _416_v18 = new _dafny.CodePoint('x'.codePointAt(0));
      let _417_i3;
      _417_i3 = _dafny.ZERO;
      L2: {
        while ((function () {
          let _coll37 = new _dafny.Set();
          for (const _compr_37 of (_dafny.Map.Empty.slice().updateUnsafe(_403_v9,_416_v18)).Keys.Elements) {
            let _427_v19 = _compr_37;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_403_v9,_416_v18)).contains(_427_v19)) {
              _coll37.add(_module.__default.safeModuloInt(_427_v19, (_dafny.ZERO).minus(new BigNumber(-404))));
            }
          }
          return _coll37;
        }()).IsSubsetOf(_dafny.Set.fromElements(p0))) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_417_i3)) {
              break L2;
            }
            _417_i3 = (_417_i3).plus(_dafny.ONE);
            let _418_v20;
            _418_v20 = _module.D6.create_DC15((_398_v4).fm6(_414_v17, _414_v17, globalState), p0, _409_v11);
            let _419_v21;
            _419_v21 = _module.D6.create_DC16(_418_v20);
            let _420_v22;
            _420_v22 = _module.D6.create_DC16(_418_v20);
            _420_v22 = _420_v22;
            r1 = _409_v11;
            let _421_v24;
            _421_v24 = _dafny.Map.Empty.slice().updateUnsafe(_400_v6.f8,_416_v18);
            let _422_v25;
            _422_v25 = _module.D11.create_DC24(_421_v24);
            let _423_v26;
            _423_v26 = _module.D11.create_DC26(_422_v25);
            let _424_v27;
            _424_v27 = _dafny.Seq.of(_409_v11, _409_v11);
            let _425_v28;
            _425_v28 = _module.D14.create_DC34(_module.__default.fm43(_423_v26, _409_v11, new BigNumber((_424_v27).length), globalState));
            _409_v11 = !((function () {
              let _coll36 = new _dafny.Map();
              for (const _compr_36 of _dafny.IntegerRange(new BigNumber(442), new BigNumber(-928))) {
                let _426_v23 = _compr_36;
                if (((new BigNumber(442)).isLessThanOrEqualTo(_426_v23)) && ((_426_v23).isLessThan(new BigNumber(-928)))) {
                  _coll36.push([(_426_v23).plus(_414_v17),_dafny.Seq.UnicodeFromString("qkbtfupqy")]);
                }
              }
              return _coll36;
            }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,_400_v6.f8))).equals((_425_v28).dtor_cf62);
            let _index51 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_406_v10).length));
            (_406_v10)[_index51] = _409_v11;
            let _index52 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_406_v10).length));
            (_406_v10)[_index52] = _409_v11;
          }
        }
      }
      r1 = _409_v11;
      r0 = _409_v11;
      r1 = _409_v11;
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _428_v0;
      _428_v0 = _dafny.Seq.UnicodeFromString("jgwejo");
      _428_v0 = _dafny.Seq.Concat(_428_v0, _428_v0);
      let _429_i0;
      _429_i0 = _dafny.ZERO;
      L3: {
        while (true) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_429_i0)) {
              break L3;
            }
            _429_i0 = (_429_i0).plus(_dafny.ONE);
            let _430_v1;
            let _nw70 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
            _430_v1 = _nw70;
            let _index53 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_430_v1).length));
            (_430_v1)[_index53] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(54)), function (_431_i1) {
              return new _dafny.CodePoint('i'.codePointAt(0));
            });
            let _index54 = _module.__default.safeIndex(new BigNumber(932), new BigNumber((_430_v1).length));
            (_430_v1)[_index54] = _dafny.Seq.UnicodeFromString("pekucbt");
            let _432_v2;
            _432_v2 = true;
            let _433_v3;
            _433_v3 = _module.D0.create_DC1(p0);
            let _434_v4;
            _434_v4 = _dafny.Seq.of(_432_v2);
            let _435_v5;
            _435_v5 = _dafny.Map.Empty.slice().updateUnsafe(_433_v3,new BigNumber((_434_v4).length));
            let _436_v6;
            _436_v6 = _dafny.Seq.of(_435_v5);
            let _437_v7;
            _437_v7 = _dafny.Seq.of(p0, p0, p0);
            let _438_v8;
            _438_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_437_v7).length),_432_v2);
            let _439_v9;
            _439_v9 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_436_v6)).cardinality())),((((_438_v8).contains(new BigNumber(432))) ? ((_438_v8).get(new BigNumber(432))) : (_432_v2))) && (_432_v2));
            let _440_v10;
            _440_v10 = _dafny.Set.fromElements(p0);
            _432_v2 = (((_438_v8).contains(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_440_v10).length),false)).Merge(_438_v8)).length))) ? ((_438_v8).get(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_440_v10).length),false)).Merge(_438_v8)).length))) : ((_432_v2) === (_432_v2)));
            let _441_v11;
            let _nw71 = Array((new BigNumber(2)).toNumber()).fill(false);
            _441_v11 = _nw71;
            let _442_v12;
            _442_v12 = _dafny.Set.fromElements(_441_v11);
            let _443_v13;
            _443_v13 = _442_v12;
            let _444_v14;
            _444_v14 = _dafny.Map.Empty.slice().updateUnsafe(_432_v2,_443_v13);
            _444_v14 = (_444_v14).update(_432_v2, _442_v12);
            let _445_v15;
            let _init9 = ((_446_p0) => function (_447_i2) {
              return (_447_i2).plus(_446_p0);
            })(p0);
            let _nw72 = Array((new BigNumber(24)).toNumber());
            for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw72.length); _i0_9++) {
              _nw72[_i0_9] = _init9(new BigNumber(_i0_9));
            }
            _445_v15 = _nw72;
            let _index55 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_445_v15).length));
            (_445_v15)[_index55] = p0;
            let _448_v16;
            let _nw73 = Array((new BigNumber(10)).toNumber()).fill(_dafny.Map.Empty);
            _448_v16 = _nw73;
            let _index56 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_448_v16).length));
            (_448_v16)[_index56] = _dafny.Map.Empty.slice().updateUnsafe(_432_v2,_432_v2);
            let _449_v17;
            _449_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_430_v1)[_module.__default.safeIndex(new BigNumber(932), new BigNumber((_430_v1).length))]);
            let _450_v18;
            _450_v18 = new _dafny.CodePoint('i'.codePointAt(0));
            let _451_v19;
            _451_v19 = _dafny.Map.Empty.slice().updateUnsafe(!((_434_v4)[_module.__default.safeIndex(p0, new BigNumber((_434_v4).length))]),_432_v2);
            let _index57 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_445_v15).length));
            let _index58 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_448_v16).length));
            let _rhs46 = _module.__default.fm44(!((p0).isLessThan(new BigNumber((_dafny.Set.fromElements(_432_v2)).length))), globalState);
            let _rhs47 = _dafny.Seq.Concat(_dafny.Seq.update(((_432_v2) ? ((((_449_v17).contains(p0)) ? ((_449_v17).get(p0)) : (_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(120)), function (_452_i3) {
              return new _dafny.CodePoint('s'.codePointAt(0));
            }), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(120)), function (_453_i3) {
              return new _dafny.CodePoint('s'.codePointAt(0));
            })).length)), _450_v18)))) : (_dafny.Seq.UnicodeFromString("fhkvhgy"))), _module.__default.safeIndex(p0, new BigNumber((((_432_v2) ? ((((_449_v17).contains(p0)) ? ((_449_v17).get(p0)) : (_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(120)), function (_454_i3) {
              return new _dafny.CodePoint('s'.codePointAt(0));
            }), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(120)), function (_455_i3) {
              return new _dafny.CodePoint('s'.codePointAt(0));
            })).length)), _450_v18)))) : (_dafny.Seq.UnicodeFromString("fhkvhgy")))).length)), _450_v18), _428_v0);
            let _rhs48 = _432_v2;
            let _rhs49 = _451_v19;
            let _rhs50 = _437_v7;
            let _lhs28 = _445_v15;
            let _lhs29 = _module.__default.safeIndex(new BigNumber(249), new BigNumber((_445_v15).length));
            let _lhs30 = _448_v16;
            let _lhs31 = _module.__default.safeIndex(new BigNumber(948), new BigNumber((_448_v16).length));
            _lhs28[_lhs29] = _rhs46;
            _428_v0 = _rhs47;
            _432_v2 = _rhs48;
            _lhs30[_lhs31] = _rhs49;
            _437_v7 = _rhs50;
          }
        }
      }
      let _456_v20;
      _456_v20 = true;
      let _457_v21;
      _457_v21 = _dafny.Seq.of(p0);
      let _source5 = _module.__default.fm45(_456_v20, _456_v20, _module.__default.safeModuloInt(new BigNumber((_457_v21).length), p0), globalState);
      if (_source5.is_DC35) {
        let _458___mcc_h0 = (_source5).cf63;
        let _459___mcc_h1 = (_source5).cf64;
        let _460___mcc_h2 = (_source5).cf65;
        let _461___mcc_h3 = (_source5).cf66;
        let _462___mcc_h4 = (_source5).cf67;
        let _463_cf67 = _462___mcc_h4;
        let _464_cf66 = _461___mcc_h3;
        let _465_cf65 = _460___mcc_h2;
        let _466_cf64 = _459___mcc_h1;
        let _467_cf63 = _458___mcc_h0;
        let _468_v22;
        _468_v22 = _dafny.Map.Empty.slice().updateUnsafe(_467_cf63,_465_cf65);
        let _469_v23;
        _469_v23 = _module.D12.create_DC29(_456_v20, (((_468_v22).contains(p0)) ? ((_468_v22).get(p0)) : (_465_cf65)), _456_v20, (_dafny.ZERO).minus((_465_cf65).plus(_467_cf63)), _464_cf66);
        let _source6 = _469_v23;
        if (_source6.is_DC28) {
          let _470___mcc_h8 = (_source6).cf46;
          let _471___mcc_h9 = (_source6).cf47;
          let _472_cf47 = _471___mcc_h9;
          let _473_cf46 = _470___mcc_h8;
          _472_cf47 = (p0).plus(_472_cf47);
          let _474_v24;
          _474_v24 = _dafny.Seq.of(((_457_v21)[_module.__default.safeIndex(_467_cf63, new BigNumber((_457_v21).length))]).isLessThanOrEqualTo(_472_cf47));
          _474_v24 = _474_v24;
          _467_cf63 = _465_cf65;
          _465_cf65 = _472_cf47;
        } else if (_source6.is_DC29) {
          let _475___mcc_h10 = (_source6).cf48;
          let _476___mcc_h11 = (_source6).cf49;
          let _477___mcc_h12 = (_source6).cf50;
          let _478___mcc_h13 = (_source6).cf51;
          let _479___mcc_h14 = (_source6).cf52;
          let _480_cf52 = _479___mcc_h14;
          let _481_cf51 = _478___mcc_h13;
          let _482_cf50 = _477___mcc_h12;
          let _483_cf49 = _476___mcc_h11;
          let _484_cf48 = _475___mcc_h10;
          let _485_v25;
          let _nw74 = Array((new BigNumber(7)).toNumber()).fill(_dafny.ZERO);
          _485_v25 = _nw74;
          let _index59 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_485_v25).length));
          (_485_v25)[_index59] = (_465_cf65).minus(p0);
          let _index60 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_485_v25).length));
          (_485_v25)[_index60] = _483_cf49;
          let _index61 = _module.__default.safeIndex(new BigNumber(982), new BigNumber((_485_v25).length));
          (_485_v25)[_index61] = (new BigNumber((_428_v0).length)).minus(_467_cf63);
          let _486_v26;
          _486_v26 = _module.D6.create_DC14(_428_v0);
          let _487_v27;
          _487_v27 = new _dafny.CodePoint('d'.codePointAt(0));
          let _488_v28;
          _488_v28 = _dafny.Seq.of(_466_cf64);
          let _489_v29;
          let _nw75 = Array((new BigNumber(28)).toNumber());
          _nw75[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update((_486_v26).dtor_cf29, _module.__default.safeIndex(_module.__default.fm44(_484_cf48, globalState), new BigNumber(((_486_v26).dtor_cf29).length)), _487_v27), _428_v0);
          _nw75[(_dafny.ONE).toNumber()] = _428_v0;
          _nw75[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_466_cf64, _module.__default.fm46(true, new BigNumber((_dafny.Seq.of(_482_cf50)).length), _484_cf48, globalState));
          _nw75[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("y");
          _nw75[(new BigNumber(4)).toNumber()] = _428_v0;
          _nw75[(new BigNumber(5)).toNumber()] = _466_cf64;
          _nw75[(new BigNumber(6)).toNumber()] = _dafny.Seq.UnicodeFromString("rc");
          _nw75[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_466_cf64, _428_v0);
          _nw75[(new BigNumber(8)).toNumber()] = _428_v0;
          _nw75[(new BigNumber(9)).toNumber()] = _466_cf64;
          _nw75[(new BigNumber(10)).toNumber()] = _dafny.Seq.UnicodeFromString("ovqriexd");
          _nw75[(new BigNumber(11)).toNumber()] = _module.__default.fm46(true, _module.__default.fm44(_484_cf48, globalState), _482_cf50, globalState);
          _nw75[(new BigNumber(12)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(134)), ((_490_v27) => function (_491_i4) {
            return _490_v27;
          })(_487_v27));
          _nw75[(new BigNumber(13)).toNumber()] = _428_v0;
          _nw75[(new BigNumber(14)).toNumber()] = (_486_v26).dtor_cf29;
          _nw75[(new BigNumber(15)).toNumber()] = _428_v0;
          _nw75[(new BigNumber(16)).toNumber()] = _466_cf64;
          _nw75[(new BigNumber(17)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(552)), function (_492_i5) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          });
          _nw75[(new BigNumber(18)).toNumber()] = _428_v0;
          _nw75[(new BigNumber(19)).toNumber()] = _428_v0;
          _nw75[(new BigNumber(20)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-733)), ((_493_v27) => function (_494_i6) {
            return _493_v27;
          })(_487_v27));
          _nw75[(new BigNumber(21)).toNumber()] = _dafny.Seq.Concat(_466_cf64, (_488_v28)[_module.__default.safeIndex(_465_cf65, new BigNumber((_488_v28).length))]);
          _nw75[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("crcuft"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(599)), ((_495_v27) => function (_496_i7) {
            return _495_v27;
          })(_487_v27)));
          _nw75[(new BigNumber(23)).toNumber()] = (_module.D6.create_DC14(_466_cf64)).dtor_cf29;
          _nw75[(new BigNumber(24)).toNumber()] = _428_v0;
          _nw75[(new BigNumber(25)).toNumber()] = _428_v0;
          _nw75[(new BigNumber(26)).toNumber()] = _dafny.Seq.UnicodeFromString("nys");
          _nw75[(new BigNumber(27)).toNumber()] = _dafny.Seq.UnicodeFromString("jijhsve");
          _489_v29 = _nw75;
          let _index62 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_489_v29).length));
          (_489_v29)[_index62] = _466_cf64;
          let _497_v30;
          _497_v30 = _dafny.MultiSet.fromElements(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("sfisdg"), _module.__default.safeIndex(_467_cf63, new BigNumber((_dafny.Seq.UnicodeFromString("sfisdg")).length)), _487_v27), _466_cf64);
          let _498_v31;
          _498_v31 = _dafny.Map.Empty.slice().updateUnsafe(_464_cf66,_497_v30);
          let _499_v32;
          _499_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,_484_cf48);
          let _index63 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_489_v29).length));
          let _rhs51 = _485_v25;
          let _rhs52 = _428_v0;
          let _rhs53 = (_dafny.ZERO).minus(_465_cf65);
          let _rhs54 = _module.__default.safeModuloInt(new BigNumber(-480), new BigNumber(((((_498_v31).contains((((_499_v32).contains((_485_v25)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_485_v25).length))])) ? ((_499_v32).get((_485_v25)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_485_v25).length))])) : (_484_cf48)))) ? ((_498_v31).get((((_499_v32).contains((_485_v25)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_485_v25).length))])) ? ((_499_v32).get((_485_v25)[_module.__default.safeIndex(new BigNumber(982), new BigNumber((_485_v25).length))])) : (_484_cf48)))) : (_497_v30))).cardinality()));
          let _lhs32 = _489_v29;
          let _lhs33 = _module.__default.safeIndex(new BigNumber(595), new BigNumber((_489_v29).length));
          _485_v25 = _rhs51;
          _lhs32[_lhs33] = _rhs52;
          _467_cf63 = _rhs53;
          _467_cf63 = _rhs54;
          _463_cf67 = _484_cf48;
        } else {
          let _500___mcc_h15 = (_source6).cf45;
          let _501_cf45 = _500___mcc_h15;
          let _502_v33;
          let _init10 = function (_503_i8) {
            return (_503_i8).minus(new BigNumber((_dafny.Seq.of(true)).length));
          };
          let _nw76 = Array((new BigNumber(29)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw76.length); _i0_10++) {
            _nw76[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _502_v33 = _nw76;
          let _504_v34;
          _504_v34 = _502_v33;
          _504_v34 = _504_v34;
          let _index64 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_502_v33).length));
          (_502_v33)[_index64] = _467_cf63;
          let _index65 = _module.__default.safeIndex(new BigNumber(921), new BigNumber((_502_v33).length));
          (_502_v33)[_index65] = (_dafny.ZERO).minus(_465_cf65);
          let _505_v35;
          _505_v35 = new _dafny.CodePoint('h'.codePointAt(0));
          let _506_v36;
          _506_v36 = _module.D0.create_DC1(new BigNumber((_466_cf64).length));
          let _507_v37;
          _507_v37 = _dafny.Map.Empty.slice().updateUnsafe(_505_v35,_506_v36);
          let _508_v38;
          _508_v38 = _dafny.Map.Empty.slice().updateUnsafe(!(_464_cf66),_463_cf67);
          let _509_v39;
          let _nw77 = Array((new BigNumber(28)).toNumber());
          _nw77[(_dafny.ZERO).toNumber()] = _456_v20;
          _nw77[(_dafny.ONE).toNumber()] = _464_cf66;
          _nw77[(new BigNumber(2)).toNumber()] = false;
          _nw77[(new BigNumber(3)).toNumber()] = _463_cf67;
          _nw77[(new BigNumber(4)).toNumber()] = _464_cf66;
          _nw77[(new BigNumber(5)).toNumber()] = _463_cf67;
          _nw77[(new BigNumber(6)).toNumber()] = _463_cf67;
          _nw77[(new BigNumber(7)).toNumber()] = _463_cf67;
          _nw77[(new BigNumber(8)).toNumber()] = _464_cf66;
          _nw77[(new BigNumber(9)).toNumber()] = _464_cf66;
          _nw77[(new BigNumber(10)).toNumber()] = _463_cf67;
          _nw77[(new BigNumber(11)).toNumber()] = (_465_cf65).isLessThan((_dafny.ZERO).minus(_467_cf63));
          _nw77[(new BigNumber(12)).toNumber()] = !(!(!(_456_v20)));
          _nw77[(new BigNumber(13)).toNumber()] = (_this).fm42(_464_cf66, globalState);
          _nw77[(new BigNumber(14)).toNumber()] = _456_v20;
          _nw77[(new BigNumber(15)).toNumber()] = _464_cf66;
          _nw77[(new BigNumber(16)).toNumber()] = (_this).fm42(_464_cf66, globalState);
          _nw77[(new BigNumber(17)).toNumber()] = _463_cf67;
          _nw77[(new BigNumber(18)).toNumber()] = (_465_cf65).isLessThanOrEqualTo(_465_cf65);
          _nw77[(new BigNumber(19)).toNumber()] = _463_cf67;
          _nw77[(new BigNumber(20)).toNumber()] = _463_cf67;
          _nw77[(new BigNumber(21)).toNumber()] = !(_464_cf66);
          _nw77[(new BigNumber(22)).toNumber()] = (_this).fm5((_502_v33)[_module.__default.safeIndex(new BigNumber(921), new BigNumber((_502_v33).length))], _464_cf66, _507_v37, (_this).fm5(p0, true, _dafny.Map.Empty.slice().updateUnsafe(_505_v35,_506_v36), _464_cf66, globalState), globalState);
          _nw77[(new BigNumber(23)).toNumber()] = false;
          _nw77[(new BigNumber(24)).toNumber()] = (_464_cf66) === (_463_cf67);
          _nw77[(new BigNumber(25)).toNumber()] = (((_508_v38).contains((_this).fm42(_456_v20, globalState))) ? ((_508_v38).get((_this).fm42(_456_v20, globalState))) : (_456_v20));
          _nw77[(new BigNumber(26)).toNumber()] = _456_v20;
          _nw77[(new BigNumber(27)).toNumber()] = _464_cf66;
          _509_v39 = _nw77;
          let _510_v40;
          _510_v40 = _module.D2.create_DC5(_465_cf65, _457_v21, _456_v20);
          let _index66 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_509_v39).length));
          (_509_v39)[_index66] = ((_464_cf66) ? ((_510_v40).dtor_cf10) : (_463_cf67));
          let _index67 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_509_v39).length));
          (_509_v39)[_index67] = (_463_cf67) || (_464_cf66);
          _502_v33 = _502_v33;
        }
        let _511_v41;
        let _init11 = ((_512_v20) => function (_513_i9) {
          return _512_v20;
        })(_456_v20);
        let _nw78 = Array((new BigNumber(27)).toNumber());
        for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw78.length); _i0_11++) {
          _nw78[_i0_11] = _init11(new BigNumber(_i0_11));
        }
        _511_v41 = _nw78;
        let _514_v42;
        _514_v42 = _dafny.Seq.of(true, false);
        let _index68 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length));
        (_511_v41)[_index68] = (_514_v42)[_module.__default.safeIndex(p0, new BigNumber((_514_v42).length))];
        let _515_v43;
        let _init12 = ((_516_p0) => function (_517_i10) {
          return (_517_i10).multipliedBy(_516_p0);
        })(p0);
        let _nw79 = Array((new BigNumber(5)).toNumber());
        for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw79.length); _i0_12++) {
          _nw79[_i0_12] = _init12(new BigNumber(_i0_12));
        }
        _515_v43 = _nw79;
        let _index69 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length));
        (_515_v43)[_index69] = _465_cf65;
        let _518_v44;
        _518_v44 = new _dafny.CodePoint('t'.codePointAt(0));
        let _519_v45;
        _519_v45 = _module.D0.create_DC1(p0);
        let _520_v46;
        _520_v46 = _dafny.Map.Empty.slice().updateUnsafe(_518_v44,_519_v45);
        let _index70 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length));
        let _index71 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length));
        let _rhs55 = _dafny.areEqual(_428_v0, _428_v0);
        let _rhs56 = _module.__default.fm44((_this).fm5(_465_cf65, _464_cf66, _520_v46, _464_cf66, globalState), globalState);
        let _lhs34 = _511_v41;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length));
        let _lhs36 = _515_v43;
        let _lhs37 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length));
        _lhs34[_lhs35] = _rhs55;
        _lhs36[_lhs37] = _rhs56;
        if (_456_v20) {
          let _521_v47;
          _521_v47 = _dafny.Map.Empty.slice().updateUnsafe(p0,_464_cf66);
          let _522_v48;
          _522_v48 = _dafny.Set.fromElements((_521_v47).update((_515_v43)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length))], (_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))]), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm47(new BigNumber((_428_v0).length), (_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))], globalState)).length),(_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))]));
          let _523_v49;
          _523_v49 = _module.D13.create_DC31(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_456_v20,_465_cf65)).length), (_515_v43)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length))]);
          let _524_v50;
          _524_v50 = _dafny.MultiSet.fromElements(_467_cf63, new BigNumber((_522_v48).length), _module.__default.safeModuloInt(new BigNumber((_514_v42).length), (_523_v49).dtor_cf56));
          let _525_v51;
          _525_v51 = _module.D7.create_DC18();
          _524_v50 = (_dafny.MultiSet.fromElements((_dafny.ZERO).minus(_module.__default.fm44(false, globalState)), _465_cf65)).Union(_module.__default.fm48(_525_v51, (_515_v43)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length))], _module.__default.fm44(true, globalState), new BigNumber(-50), globalState));
          let _526_v52;
          _526_v52 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_518_v44,!((_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))])),_465_cf65);
          _526_v52 = (((_456_v20) ? (_526_v52) : (_526_v52))).Merge(_526_v52);
          let _527_v53;
          let _nw80 = Array((new BigNumber(24)).toNumber()).fill(_dafny.Map.Empty);
          _527_v53 = _nw80;
          let _528_v54;
          let _nw81 = new _module.C1();
          _nw81.__ctor(_467_cf63, _527_v53, _dafny.Seq.Concat(_457_v21, _457_v21), _519_v45);
          _528_v54 = _nw81;
          let _529_v55;
          _529_v55 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm46((_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))], new BigNumber((_524_v50).cardinality()), false, globalState),(_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))]);
          let _530_v56;
          _530_v56 = _dafny.Map.Empty.slice().updateUnsafe(_457_v21,_466_cf64);
          let _531_v57;
          _531_v57 = _dafny.Seq.of(_428_v0, (((_530_v56).contains(_457_v21)) ? ((_530_v56).get(_457_v21)) : (_dafny.Seq.UnicodeFromString("mmjnwlr"))), _466_cf64, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("l"), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.UnicodeFromString("l")).length)), _518_v44));
          let _532_v58;
          _532_v58 = _dafny.Set.fromElements((_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))], _456_v20, true);
          let _index72 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length));
          (_511_v41)[_index72] = (((_529_v55).contains((_531_v57)[_module.__default.safeIndex(new BigNumber(558), new BigNumber((_531_v57).length))])) ? ((_529_v55).get((_531_v57)[_module.__default.safeIndex(new BigNumber(558), new BigNumber((_531_v57).length))])) : ((_532_v58).IsProperSubsetOf(_dafny.Set.fromElements((_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))]))));
          let _533_v59;
          _533_v59 = _dafny.Map.Empty.slice().updateUnsafe(_518_v44,_528_v54.f9);
          let _534_v60;
          _534_v60 = _dafny.Map.Empty.slice().updateUnsafe((_533_v59).Merge(_533_v59),_518_v44);
          _534_v60 = _534_v60;
        } else {
          let _535_v61;
          let _nw82 = new _module.C0();
          _nw82.__ctor(_dafny.Seq.UnicodeFromString("acttpcc"));
          _535_v61 = _nw82;
          let _536_v62;
          let _nw83 = Array((_dafny.ONE).toNumber());
          _nw83[(_dafny.ZERO).toNumber()] = _514_v42;
          _536_v62 = _nw83;
          let _index73 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_536_v62).length));
          (_536_v62)[_index73] = _514_v42;
          let _index74 = _module.__default.safeIndex(new BigNumber(392), new BigNumber((_536_v62).length));
          (_536_v62)[_index74] = _514_v42;
          _456_v20 = _463_cf67;
          _465_cf65 = p0;
          let _537_v63;
          let _nw84 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
          _537_v63 = _nw84;
          let _538_v64;
          let _nw85 = new _module.C1();
          _nw85.__ctor((_515_v43)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length))], _537_v63, _dafny.Seq.update(_457_v21, _module.__default.safeIndex(_465_cf65, new BigNumber((_457_v21).length)), _465_cf65), _519_v45);
          _538_v64 = _nw85;
        }
        if (_463_cf67) {
          let _539_v65;
          _539_v65 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.IsPrefixOf(_428_v0, _dafny.Seq.of(_518_v44)));
          _539_v65 = (_539_v65).update((_515_v43)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length))], true);
          _456_v20 = !((_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))]);
          let _540_v66;
          _540_v66 = _dafny.Set.fromElements(_457_v21);
          let _541_v67;
          _541_v67 = _module.D15.create_DC37(_540_v66);
          let _index75 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length));
          (_511_v41)[_index75] = (_540_v66).IsSubsetOf((_541_v67).dtor_cf70);
          let _index76 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length));
          (_511_v41)[_index76] = _464_cf66;
          let _index77 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length));
          (_515_v43)[_index77] = _module.__default.safeDivisionInt(p0, (_515_v43)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length))]);
        } else {
          let _index78 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length));
          (_515_v43)[_index78] = _465_cf65;
          let _542_v68;
          _542_v68 = _dafny.Set.fromElements((_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))], (_this).fm5(p0, true, _dafny.Map.Empty.slice().updateUnsafe((_428_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_467_cf63,(_511_v41)[_module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length))])).length), new BigNumber((_428_v0).length))],_519_v45), false, globalState));
          let _543_v69;
          _543_v69 = _dafny.Map.Empty.slice().updateUnsafe(_542_v68,_456_v20);
          _465_cf65 = ((_515_v43)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_515_v43).length))]).minus(new BigNumber((_543_v69).length));
          _463_cf67 = _463_cf67;
          let _544_v70;
          let _nw86 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
          _544_v70 = _nw86;
          let _pat_let_tv9 = _467_cf63;
          let _545_v71;
          let _nw87 = new _module.C1();
          _nw87.__ctor(new BigNumber(-834), _544_v70, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-92)), ((_546_v20) => function (_547_i11) {
            return (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_546_v20)).cardinality()));
          })(_456_v20)), function (_pat_let12_0) {
            return function (_548_dt__update__tmp_h1) {
              return function (_pat_let13_0) {
                return function (_549_dt__update_hcf2_h0) {
                  return _module.D0.create_DC1(_549_dt__update_hcf2_h0);
                }(_pat_let13_0);
              }(_pat_let_tv9);
            }(_pat_let12_0);
          }(_519_v45));
          _545_v71 = _nw87;
          let _index79 = _module.__default.safeIndex(new BigNumber(737), new BigNumber((_511_v41).length));
          (_511_v41)[_index79] = false;
        }
      } else if (_source5.is_DC36) {
        let _550___mcc_h5 = (_source5).cf68;
        let _551___mcc_h6 = (_source5).cf69;
        let _552_cf69 = _551___mcc_h6;
        let _553_cf68 = _550___mcc_h5;
        let _554_v72;
        let _init13 = ((_555_p0) => function (_556_i12) {
          return (_556_i12).minus(_555_p0);
        })(p0);
        let _nw88 = Array((new BigNumber(3)).toNumber());
        for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw88.length); _i0_13++) {
          _nw88[_i0_13] = _init13(new BigNumber(_i0_13));
        }
        _554_v72 = _nw88;
        let _557_v73;
        _557_v73 = _dafny.Map.Empty.slice().updateUnsafe(true,_554_v72);
        let _558_v74;
        _558_v74 = _module.D12.create_DC28((_557_v73).equals((_557_v73).update(_456_v20, _554_v72)), p0);
        _558_v74 = _558_v74;
        let _559_v75;
        _559_v75 = _dafny.Map.Empty.slice().updateUnsafe(!(false),!(_456_v20) || (_456_v20));
        _559_v75 = (_559_v75).update(!(_552_cf69).contains(_456_v20), _456_v20);
        let _560_v76;
        _560_v76 = _dafny.Map.Empty.slice().updateUnsafe(p0,_456_v20);
        let _561_v77;
        _561_v77 = _dafny.Map.Empty.slice().updateUnsafe(_456_v20,_560_v76);
        _552_cf69 = (_552_cf69).update(!(_561_v77).contains(_456_v20), p0);
        if (_456_v20) {
          let _562_v78;
          _562_v78 = new BigNumber(633);
          _562_v78 = p0;
          let _563_v79;
          _563_v79 = _dafny.Set.fromElements(true, _456_v20);
          let _564_v80;
          _564_v80 = _dafny.MultiSet.fromElements(_563_v79);
          let _565_v81;
          _565_v81 = _dafny.Map.Empty.slice().updateUnsafe(_456_v20,_564_v80);
          _564_v80 = ((_564_v80).update(_563_v79, _module.__default.abs(_562_v78))).Union((((_565_v81).contains(_456_v20)) ? ((_565_v81).get(_456_v20)) : (_dafny.MultiSet.fromElements(_563_v79))));
          _557_v73 = (_557_v73).update((((_559_v75).contains(_456_v20)) ? ((_559_v75).get(_456_v20)) : (_456_v20)), _554_v72);
          let _566_v82;
          _566_v82 = _dafny.Seq.of(!(_456_v20), _456_v20, _456_v20);
          let _567_v83;
          _567_v83 = _dafny.Seq.of(_566_v82, _566_v82);
          let _568_v84;
          _568_v84 = _dafny.Map.Empty.slice().updateUnsafe((_566_v82)[_module.__default.safeIndex(_562_v78, new BigNumber((_566_v82).length))],(_567_v83)[_module.__default.safeIndex((_457_v21)[_module.__default.safeIndex(_562_v78, new BigNumber((_457_v21).length))], new BigNumber((_567_v83).length))]);
          _568_v84 = _568_v84;
          let _569_v85;
          let _nw89 = new _module.C0();
          _nw89.__ctor(_428_v0);
          _569_v85 = _nw89;
        } else {
          let _570_v86;
          let _nw90 = Array((new BigNumber(19)).toNumber()).fill(false);
          _570_v86 = _nw90;
          let _index80 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_570_v86).length));
          (_570_v86)[_index80] = _456_v20;
          let _index81 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_570_v86).length));
          let _rhs57 = _457_v21;
          let _rhs58 = !(_456_v20) || (_456_v20);
          let _lhs38 = _570_v86;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_570_v86).length));
          _457_v21 = _rhs57;
          _lhs38[_lhs39] = _rhs58;
          let _571_v87;
          _571_v87 = _module.D6.create_DC15(p0, _module.__default.fm44((_570_v86)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_570_v86).length))], globalState), !(_456_v20));
          _571_v87 = _571_v87;
          let _index82 = _module.__default.safeIndex(new BigNumber(133), new BigNumber((_570_v86).length));
          (_570_v86)[_index82] = _dafny.Seq.IsPrefixOf(_428_v0, _428_v0);
          let _572_v88;
          _572_v88 = _dafny.Map.Empty.slice().updateUnsafe(_552_cf69,_428_v0);
          let _573_v89;
          let _out20;
          _out20 = (_this).m14(false, ((_570_v86)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_570_v86).length))]) && ((_570_v86)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_570_v86).length))]), (_572_v88).Merge(_dafny.Map.Empty.slice().updateUnsafe(_552_cf69,_428_v0)), (_dafny.ZERO).minus(new BigNumber((_428_v0).length)), globalState);
          _573_v89 = _out20;
          let _574_v90;
          _574_v90 = new _dafny.CodePoint('s'.codePointAt(0));
          let _575_v91;
          _575_v91 = _dafny.Map.Empty.slice().updateUnsafe(false,_574_v90);
          _575_v91 = (_575_v91).update((_570_v86)[_module.__default.safeIndex(new BigNumber(133), new BigNumber((_570_v86).length))], _574_v90);
        }
      } else {
        let _576___mcc_h7 = (_source5).cf62;
        let _577_cf62 = _576___mcc_h7;
        let _578_v92;
        _578_v92 = _module.D12.create_DC28(!(_456_v20) || (_456_v20), p0);
        _578_v92 = _578_v92;
        _456_v20 = _456_v20;
        let _579_v93;
        let _init14 = ((_580_v20) => function (_581_i13) {
          return (_dafny.Set.fromElements(_580_v20, _580_v20)).Difference(_dafny.Set.fromElements(true));
        })(_456_v20);
        let _nw91 = Array((new BigNumber(9)).toNumber());
        for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw91.length); _i0_14++) {
          _nw91[_i0_14] = _init14(new BigNumber(_i0_14));
        }
        _579_v93 = _nw91;
        let _582_v94;
        _582_v94 = _dafny.Set.fromElements(_456_v20);
        let _index83 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_579_v93).length));
        (_579_v93)[_index83] = _582_v94;
        let _index84 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_579_v93).length));
        (_579_v93)[_index84] = _582_v94;
        let _583_v95;
        _583_v95 = _456_v20;
        if (!((_583_v95))) {
          let _584_v96;
          _584_v96 = new _dafny.CodePoint('c'.codePointAt(0));
          let _585_v97;
          _585_v97 = _dafny.Map.Empty.slice().updateUnsafe((_module.__default.fm44(_456_v20, globalState)).multipliedBy(p0),p0);
          let _rhs59 = ((_456_v20) ? (_584_v96) : (_584_v96));
          let _rhs60 = (_456_v20) && (_456_v20);
          let _rhs61 = _585_v97;
          let _rhs62 = (_dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus(p0))).equals((_585_v97).Merge(_585_v97));
          _584_v96 = _rhs59;
          _456_v20 = _rhs60;
          _585_v97 = _rhs61;
          _456_v20 = _rhs62;
          let _586_v98;
          _586_v98 = _dafny.Map.Empty.slice().updateUnsafe((p0).plus(p0),((_579_v93)[_module.__default.safeIndex(new BigNumber(60), new BigNumber((_579_v93).length))]).contains(_456_v20));
          _586_v98 = (_586_v98).update(p0, true);
          _456_v20 = _456_v20;
          let _587_v99;
          _587_v99 = _dafny.MultiSet.fromElements(true, _456_v20, false);
          let _588_v100;
          _588_v100 = _dafny.Map.Empty.slice().updateUnsafe(p0,_587_v99);
          let _589_v101;
          let _nw92 = Array((new BigNumber(19)).toNumber());
          _nw92[(_dafny.ZERO).toNumber()] = _587_v99;
          _nw92[(_dafny.ONE).toNumber()] = _587_v99;
          _nw92[(new BigNumber(2)).toNumber()] = (_587_v99).Union(_587_v99);
          _nw92[(new BigNumber(3)).toNumber()] = _dafny.MultiSet.fromElements(_456_v20);
          _nw92[(new BigNumber(4)).toNumber()] = _587_v99;
          _nw92[(new BigNumber(5)).toNumber()] = (_587_v99).Union(_587_v99);
          _nw92[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(_456_v20, _456_v20);
          _nw92[(new BigNumber(7)).toNumber()] = _587_v99;
          _nw92[(new BigNumber(8)).toNumber()] = _587_v99;
          _nw92[(new BigNumber(9)).toNumber()] = _587_v99;
          _nw92[(new BigNumber(10)).toNumber()] = _587_v99;
          _nw92[(new BigNumber(11)).toNumber()] = _587_v99;
          _nw92[(new BigNumber(12)).toNumber()] = (((_588_v100).contains(p0)) ? ((_588_v100).get(p0)) : (_587_v99));
          _nw92[(new BigNumber(13)).toNumber()] = _587_v99;
          _nw92[(new BigNumber(14)).toNumber()] = _587_v99;
          _nw92[(new BigNumber(15)).toNumber()] = _587_v99;
          _nw92[(new BigNumber(16)).toNumber()] = (_dafny.MultiSet.fromElements(_456_v20, _456_v20)).update(_456_v20, _module.__default.abs(p0));
          _nw92[(new BigNumber(17)).toNumber()] = _587_v99;
          _nw92[(new BigNumber(18)).toNumber()] = _587_v99;
          _589_v101 = _nw92;
          let _590_v102;
          _590_v102 = _dafny.Seq.of(_dafny.MultiSet.fromElements((_578_v92).dtor_cf46));
          let _index85 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_589_v101).length));
          (_589_v101)[_index85] = (_590_v102)[_module.__default.safeIndex(p0, new BigNumber((_590_v102).length))];
          let _index86 = _module.__default.safeIndex(new BigNumber(741), new BigNumber((_589_v101).length));
          (_589_v101)[_index86] = _587_v99;
          _456_v20 = (p0).isLessThanOrEqualTo(p0);
        } else {
          let _591_v103;
          let _nw93 = Array((new BigNumber(24)).toNumber()).fill(false);
          _591_v103 = _nw93;
          let _592_v104;
          _592_v104 = _module.D16.create_DC41(_591_v103);
          _591_v103 = (_592_v104).dtor_cf75;
          let _593_v105;
          _593_v105 = new BigNumber(795);
          let _594_v106;
          _594_v106 = _dafny.Map.Empty.slice().updateUnsafe(_593_v105,_456_v20);
          let _595_v107;
          _595_v107 = _module.D0.create_DC1(_593_v105);
          _593_v105 = _module.__default.safeModuloInt((p0).multipliedBy(_module.__default.fm44((((_594_v106).contains(p0)) ? ((_594_v106).get(p0)) : (_456_v20)), globalState)), ((_595_v107).dtor_cf2).minus(new BigNumber((_dafny.Seq.of(_456_v20)).length)));
          _456_v20 = !(_456_v20);
          _456_v20 = !(p0).isEqualTo((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_428_v0).length), p0)));
          let _596_v108;
          let _nw94 = Array((new BigNumber(13)).toNumber()).fill(_module.D0.Default());
          _596_v108 = _nw94;
          let _index87 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_596_v108).length));
          (_596_v108)[_index87] = _595_v107;
          let _597_v109;
          _597_v109 = new _dafny.CodePoint('l'.codePointAt(0));
          let _index88 = _module.__default.safeIndex(new BigNumber(631), new BigNumber((_596_v108).length));
          (_596_v108)[_index88] = _module.__default.fm49(_597_v109, globalState);
        }
      }
      if (!(_456_v20)) {
        let _598_v110;
        _598_v110 = _dafny.MultiSet.fromElements(_456_v20);
        let _599_v111;
        _599_v111 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_598_v110).cardinality()),_456_v20);
        let _600_v112;
        _600_v112 = _dafny.Seq.of(_599_v111);
        let _601_v113;
        let _nw95 = Array((new BigNumber(5)).toNumber());
        _nw95[(_dafny.ZERO).toNumber()] = _599_v111;
        _nw95[(_dafny.ONE).toNumber()] = _599_v111;
        _nw95[(new BigNumber(2)).toNumber()] = _module.__default.fm50(_456_v20, _456_v20, p0, globalState);
        _nw95[(new BigNumber(3)).toNumber()] = (_600_v112)[_module.__default.safeIndex(p0, new BigNumber((_600_v112).length))];
        _nw95[(new BigNumber(4)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p0,_456_v20)).update(new BigNumber(-789), _456_v20);
        _601_v113 = _nw95;
        let _602_v114;
        _602_v114 = _dafny.Map.Empty.slice().updateUnsafe(_456_v20,_601_v113);
        let _603_v115;
        let _nw96 = new _module.C0();
        _nw96.__ctor(_428_v0);
        _603_v115 = _nw96;
        let _604_v116;
        _604_v116 = _dafny.Seq.of(_603_v115);
        _602_v114 = (_602_v114).update((p0).isLessThanOrEqualTo(new BigNumber((_604_v116).length)), _601_v113);
        let _605_v117;
        let _nw97 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
        _605_v117 = _nw97;
        let _index89 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_605_v117).length));
        (_605_v117)[_index89] = _module.__default.fm44(!(_456_v20), globalState);
        let _index90 = _module.__default.safeIndex(new BigNumber(99), new BigNumber((_605_v117).length));
        (_605_v117)[_index90] = p0;
        _457_v21 = _dafny.Seq.Concat(_457_v21, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-305)), ((_606_p0) => function (_607_i14) {
          return _606_p0;
        })(p0)));
        _456_v20 = _456_v20;
        let _608_v118;
        _608_v118 = _dafny.Set.fromElements(p0, new BigNumber(151), (_605_v117)[_module.__default.safeIndex(new BigNumber(99), new BigNumber((_605_v117).length))]);
        _608_v118 = (_608_v118).Intersect(_608_v118);
      } else {
        _456_v20 = _456_v20;
        let _609_v119;
        let _nw98 = Array((new BigNumber(7)).toNumber()).fill([]);
        _609_v119 = _nw98;
        let _610_v120;
        _610_v120 = _dafny.Set.fromElements(_456_v20);
        let _611_v121;
        _611_v121 = _dafny.Seq.of(_610_v120, _610_v120);
        let _612_v122;
        _612_v122 = _module.D12.create_DC27(_610_v120);
        let _613_v123;
        let _nw99 = Array((new BigNumber(18)).toNumber());
        _nw99[(_dafny.ZERO).toNumber()] = _610_v120;
        _nw99[(_dafny.ONE).toNumber()] = _610_v120;
        _nw99[(new BigNumber(2)).toNumber()] = (_611_v121)[_module.__default.safeIndex(p0, new BigNumber((_611_v121).length))];
        _nw99[(new BigNumber(3)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(4)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(5)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(6)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(7)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(8)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(9)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(10)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(11)).toNumber()] = (_612_v122).dtor_cf45;
        _nw99[(new BigNumber(12)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(13)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(14)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(15)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(16)).toNumber()] = _610_v120;
        _nw99[(new BigNumber(17)).toNumber()] = _610_v120;
        _613_v123 = _nw99;
        let _614_v124;
        _614_v124 = _module.D0.create_DC0(_613_v123, new BigNumber(596));
        let _615_v125;
        _615_v125 = _module.D15.create_DC38(_module.__default.fm51(_456_v20, p0, p0, globalState));
        let _616_v126;
        _616_v126 = _dafny.Map.Empty.slice().updateUnsafe(_614_v124,(_615_v125).dtor_cf71);
        let _617_v127;
        _617_v127 = _dafny.Seq.of(_456_v20);
        let _618_v128;
        _618_v128 = _dafny.Seq.of(_457_v21);
        let _619_v129;
        let _nw100 = Array((new BigNumber(23)).toNumber());
        _nw100[(_dafny.ZERO).toNumber()] = p0;
        _nw100[(_dafny.ONE).toNumber()] = p0;
        _nw100[(new BigNumber(2)).toNumber()] = new BigNumber(550);
        _nw100[(new BigNumber(3)).toNumber()] = new BigNumber((_428_v0).length);
        _nw100[(new BigNumber(4)).toNumber()] = _module.__default.fm44(_456_v20, globalState);
        _nw100[(new BigNumber(5)).toNumber()] = p0;
        _nw100[(new BigNumber(6)).toNumber()] = p0;
        _nw100[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
        _nw100[(new BigNumber(8)).toNumber()] = p0;
        _nw100[(new BigNumber(9)).toNumber()] = p0;
        _nw100[(new BigNumber(10)).toNumber()] = p0;
        _nw100[(new BigNumber(11)).toNumber()] = new BigNumber((_428_v0).length);
        _nw100[(new BigNumber(12)).toNumber()] = p0;
        _nw100[(new BigNumber(13)).toNumber()] = new BigNumber((_616_v126).length);
        _nw100[(new BigNumber(14)).toNumber()] = new BigNumber((_617_v127).length);
        _nw100[(new BigNumber(15)).toNumber()] = new BigNumber((_618_v128).length);
        _nw100[(new BigNumber(16)).toNumber()] = p0;
        _nw100[(new BigNumber(17)).toNumber()] = p0;
        _nw100[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw100[(new BigNumber(19)).toNumber()] = new BigNumber(259);
        _nw100[(new BigNumber(20)).toNumber()] = p0;
        _nw100[(new BigNumber(21)).toNumber()] = p0;
        _nw100[(new BigNumber(22)).toNumber()] = new BigNumber(-31);
        _619_v129 = _nw100;
        let _index91 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_609_v119).length));
        (_609_v119)[_index91] = _619_v129;
        let _620_v130;
        let _nw101 = Array((new BigNumber(14)).toNumber()).fill(false);
        _620_v130 = _nw101;
        let _index92 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_620_v130).length));
        (_620_v130)[_index92] = _456_v20;
        let _index93 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_619_v129).length));
        (_619_v129)[_index93] = _module.__default.fm44(_456_v20, globalState);
        let _index94 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_619_v129).length));
        (_619_v129)[_index94] = p0;
        let _index95 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_609_v119).length));
        let _index96 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_620_v130).length));
        let _index97 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_619_v129).length));
        let _index98 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_619_v129).length));
        let _rhs63 = _619_v129;
        let _rhs64 = _456_v20;
        let _rhs65 = p0;
        let _rhs66 = _module.__default.safeModuloInt(p0, (_457_v21)[_module.__default.safeIndex(p0, new BigNumber((_457_v21).length))]);
        let _lhs40 = _609_v119;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_609_v119).length));
        let _lhs42 = _620_v130;
        let _lhs43 = _module.__default.safeIndex(new BigNumber(703), new BigNumber((_620_v130).length));
        let _lhs44 = _619_v129;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_619_v129).length));
        let _lhs46 = _619_v129;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_619_v129).length));
        _lhs40[_lhs41] = _rhs63;
        _lhs42[_lhs43] = _rhs64;
        _lhs44[_lhs45] = _rhs65;
        _lhs46[_lhs47] = _rhs66;
        _456_v20 = (_this).fm42(_456_v20, globalState);
        let _621_v131;
        _621_v131 = _dafny.Set.fromElements(_620_v130);
        let _622_v132;
        _622_v132 = _621_v131;
        let _623_v133;
        _623_v133 = (_622_v132);
        let _source7 = _623_v133;
        let _624___mcc_h16 = _source7;
        let _625_cf36 = _624___mcc_h16;
        let _626_v134;
        _626_v134 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(518)), function (_627_i15) {
          return new _dafny.CodePoint('f'.codePointAt(0));
        }),(_dafny.ZERO).minus(p0));
        let _628_v135;
        _628_v135 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_626_v134).length),(_620_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_620_v130).length))]);
        let _629_v137;
        _629_v137 = new _dafny.CodePoint('f'.codePointAt(0));
        _456_v20 = (_this).fm5(new BigNumber((_628_v135).length), !(_456_v20) || (_456_v20), function () {
          let _coll38 = new _dafny.Map();
          for (const _compr_38 of (_dafny.MultiSet.fromElements(_629_v137, _629_v137, _629_v137, _629_v137, _629_v137)).Elements) {
            let _630_v136 = _compr_38;
            if ((_dafny.MultiSet.fromElements(_629_v137, _629_v137, _629_v137, _629_v137, _629_v137)).contains(_630_v136)) {
              _coll38.push([_630_v136,_module.D0.create_DC1(p0)]);
            }
          }
          return _coll38;
        }(), _456_v20, globalState);
        let _index99 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_613_v123).length));
        (_613_v123)[_index99] = _610_v120;
        let _index100 = _module.__default.safeIndex(new BigNumber(607), new BigNumber((_613_v123).length));
        (_613_v123)[_index100] = _610_v120;
        let _631_v139;
        let _init15 = ((_632_v130, _633_v129, _634_p0) => function (_635_i16) {
          return (_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_632_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_632_v130).length))],(_633_v129)[_module.__default.safeIndex(new BigNumber(483), new BigNumber((_633_v129).length))]))).Union(_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_632_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_632_v130).length))],new BigNumber((function () {
            let _coll39 = new _dafny.Set();
            for (const _compr_39 of (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_632_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_632_v130).length))],(_632_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_632_v130).length))])).length), _634_p0))).Elements) {
              let _636_v138 = _compr_39;
              if ((_dafny.Set.fromElements(_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_632_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_632_v130).length))],(_632_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_632_v130).length))])).length), _634_p0))).contains(_636_v138)) {
                _coll39.add(_636_v138);
              }
            }
            return _coll39;
          }()).length))));
        })(_620_v130, _619_v129, p0);
        let _nw102 = Array((new BigNumber(6)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw102.length); _i0_15++) {
          _nw102[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _631_v139 = _nw102;
        let _637_v140;
        _637_v140 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_619_v129)[_module.__default.safeIndex(new BigNumber(483), new BigNumber((_619_v129).length))]);
        let _638_v141;
        _638_v141 = _dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(!((_620_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_620_v130).length))]),new BigNumber((_637_v140).length)));
        let _index101 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_631_v139).length));
        (_631_v139)[_index101] = _638_v141;
        let _index102 = _module.__default.safeIndex(new BigNumber(64), new BigNumber((_631_v139).length));
        (_631_v139)[_index102] = _638_v141;
        let _639_v142;
        _639_v142 = _module.D13.create_DC32(_428_v0, (_620_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_620_v130).length))], (_619_v129)[_module.__default.safeIndex(new BigNumber(483), new BigNumber((_619_v129).length))], (_609_v119)[_module.__default.safeIndex(new BigNumber(170), new BigNumber((_609_v119).length))]);
        let _640_v143;
        _640_v143 = _dafny.Map.Empty.slice().updateUnsafe(!((_639_v142).dtor_cf58),p0);
        let _641_v144;
        _641_v144 = _dafny.MultiSet.fromElements(_623_v133);
        let _642_v145;
        _642_v145 = _dafny.Map.Empty.slice().updateUnsafe(_640_v143,(((_641_v144).contains(_dafny.Set.fromElements(_620_v130, _620_v130))) ? ((_641_v144).get(_dafny.Set.fromElements(_620_v130, _620_v130))) : (new BigNumber((_428_v0).length))));
        let _643_v146;
        _643_v146 = _module.D0.create_DC1((_619_v129)[_module.__default.safeIndex(new BigNumber(483), new BigNumber((_619_v129).length))]);
        let _644_v147;
        _644_v147 = _dafny.Map.Empty.slice().updateUnsafe(_629_v137,_643_v146);
        _642_v145 = (((_456_v20) && ((_this).fm5(new BigNumber(630), _456_v20, _644_v147, (_620_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_620_v130).length))], globalState))) ? (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(_456_v20,p0),(((_640_v143).contains(true)) ? ((_640_v143).get(true)) : (p0)))) : (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_620_v130)[_module.__default.safeIndex(new BigNumber(703), new BigNumber((_620_v130).length))],new BigNumber(-451)),p0)));
        let _645_v148;
        let _init16 = function (_646_i17) {
          return new _dafny.CodePoint('s'.codePointAt(0));
        };
        let _nw103 = Array((new BigNumber(19)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw103.length); _i0_16++) {
          _nw103[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _645_v148 = _nw103;
        let _647_v149;
        _647_v149 = _dafny.Seq.of(_645_v148);
        _647_v149 = _647_v149;
      }
      let _648_v150;
      _648_v150 = _dafny.Map.Empty.slice().updateUnsafe(_428_v0,_456_v20);
      _648_v150 = (_648_v150).update(_dafny.Seq.Concat(_428_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(680)), function (_649_i18) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      })), _456_v20);
      if (!(false)) {
        let _650_v151;
        let _nw104 = Array((new BigNumber(15)).toNumber()).fill(_dafny.ZERO);
        _650_v151 = _nw104;
        let _651_v152;
        _651_v152 = _650_v151;
        _651_v152 = _651_v152;
        let _652_v153;
        _652_v153 = new BigNumber(404);
        _652_v153 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-90)), function (_653_i19) {
          return new _dafny.CodePoint('b'.codePointAt(0));
        }), _428_v0)).length);
        let _654_v154;
        _654_v154 = new _dafny.CodePoint('g'.codePointAt(0));
        let _655_v155;
        _655_v155 = _dafny.Map.Empty.slice().updateUnsafe(_654_v154,p0);
        _655_v155 = (_655_v155).update(_654_v154, _652_v153);
        _654_v154 = _654_v154;
        _652_v153 = (_dafny.ZERO).minus(p0);
      } else {
        let _656_v156;
        _656_v156 = _module.D4.create_DC10(new BigNumber((_428_v0).length), (_dafny.ZERO).minus(p0), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_456_v20,false)).length));
        let _657_v157;
        _657_v157 = new _dafny.CodePoint('a'.codePointAt(0));
        let _658_v158;
        _658_v158 = _dafny.Map.Empty.slice().updateUnsafe(_456_v20,_657_v157);
        let _659_v159;
        _659_v159 = _dafny.Set.fromElements(new BigNumber((_658_v158).length));
        let _660_v160;
        let _nw105 = Array((new BigNumber(24)).toNumber());
        _nw105[(_dafny.ZERO).toNumber()] = _module.__default.fm44(_456_v20, globalState);
        _nw105[(_dafny.ONE).toNumber()] = p0;
        _nw105[(new BigNumber(2)).toNumber()] = new BigNumber(714);
        _nw105[(new BigNumber(3)).toNumber()] = p0;
        _nw105[(new BigNumber(4)).toNumber()] = p0;
        _nw105[(new BigNumber(5)).toNumber()] = p0;
        _nw105[(new BigNumber(6)).toNumber()] = (_656_v156).dtor_cf20;
        _nw105[(new BigNumber(7)).toNumber()] = p0;
        _nw105[(new BigNumber(8)).toNumber()] = p0;
        _nw105[(new BigNumber(9)).toNumber()] = p0;
        _nw105[(new BigNumber(10)).toNumber()] = p0;
        _nw105[(new BigNumber(11)).toNumber()] = p0;
        _nw105[(new BigNumber(12)).toNumber()] = p0;
        _nw105[(new BigNumber(13)).toNumber()] = new BigNumber((_659_v159).length);
        _nw105[(new BigNumber(14)).toNumber()] = p0;
        _nw105[(new BigNumber(15)).toNumber()] = p0;
        _nw105[(new BigNumber(16)).toNumber()] = p0;
        _nw105[(new BigNumber(17)).toNumber()] = p0;
        _nw105[(new BigNumber(18)).toNumber()] = new BigNumber((_428_v0).length);
        _nw105[(new BigNumber(19)).toNumber()] = p0;
        _nw105[(new BigNumber(20)).toNumber()] = p0;
        _nw105[(new BigNumber(21)).toNumber()] = p0;
        _nw105[(new BigNumber(22)).toNumber()] = p0;
        _nw105[(new BigNumber(23)).toNumber()] = p0;
        _660_v160 = _nw105;
        let _661_v161;
        let _nw106 = Array((new BigNumber(12)).toNumber());
        _nw106[(_dafny.ZERO).toNumber()] = _660_v160;
        _nw106[(_dafny.ONE).toNumber()] = _660_v160;
        _nw106[(new BigNumber(2)).toNumber()] = _660_v160;
        _nw106[(new BigNumber(3)).toNumber()] = _660_v160;
        _nw106[(new BigNumber(4)).toNumber()] = _660_v160;
        _nw106[(new BigNumber(5)).toNumber()] = _660_v160;
        _nw106[(new BigNumber(6)).toNumber()] = _660_v160;
        _nw106[(new BigNumber(7)).toNumber()] = _660_v160;
        _nw106[(new BigNumber(8)).toNumber()] = _660_v160;
        _nw106[(new BigNumber(9)).toNumber()] = _660_v160;
        _nw106[(new BigNumber(10)).toNumber()] = _660_v160;
        _nw106[(new BigNumber(11)).toNumber()] = _660_v160;
        _661_v161 = _nw106;
        let _662_v162;
        _662_v162 = _module.D2.create_DC4(p0, _456_v20, _661_v161);
        let _source8 = _662_v162;
        if (_source8.is_DC4) {
          let _663___mcc_h17 = (_source8).cf5;
          let _664___mcc_h18 = (_source8).cf6;
          let _665___mcc_h19 = (_source8).cf7;
          let _666_cf7 = _665___mcc_h19;
          let _667_cf6 = _664___mcc_h18;
          let _668_cf5 = _663___mcc_h17;
          _456_v20 = false;
          let _669_v163;
          let _nw107 = Array((new BigNumber(16)).toNumber()).fill(_dafny.Map.Empty);
          _669_v163 = _nw107;
          let _670_v164;
          _670_v164 = _dafny.MultiSet.fromElements(_456_v20);
          let _671_v165;
          _671_v165 = _module.D9.create_DC21(_670_v164);
          let _672_v166;
          _672_v166 = _module.D15.create_DC39(_428_v0, _671_v165);
          let _673_v167;
          _673_v167 = _dafny.Seq.of(_428_v0, (_672_v166).dtor_cf72);
          let _674_v168;
          let _nw108 = new _module.C1();
          _nw108.__ctor(p0, _669_v163, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(378)), ((_675_v21) => function (_676_i20) {
            return new BigNumber((_675_v21).length);
          })(_457_v21)), _457_v21), _module.D0.create_DC1(new BigNumber((_673_v167).length)));
          _674_v168 = _nw108;
          _660_v160 = _660_v160;
          _428_v0 = _428_v0;
        } else if (_source8.is_DC5) {
          let _677___mcc_h20 = (_source8).cf8;
          let _678___mcc_h21 = (_source8).cf9;
          let _679___mcc_h22 = (_source8).cf10;
          let _680_cf10 = _679___mcc_h22;
          let _681_cf9 = _678___mcc_h21;
          let _682_cf8 = _677___mcc_h20;
          let _index103 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_660_v160).length));
          (_660_v160)[_index103] = new BigNumber(711);
          let _683_v169;
          _683_v169 = _dafny.MultiSet.fromElements(true, _456_v20);
          let _index104 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_660_v160).length));
          (_660_v160)[_index104] = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(_682_cf8, new BigNumber((_428_v0).length)), (((_683_v169).contains(_456_v20)) ? ((_683_v169).get(_456_v20)) : (p0)));
          let _684_v170;
          _684_v170 = _dafny.Seq.of(_428_v0, _428_v0);
          let _685_v171;
          _685_v171 = _dafny.Map.Empty.slice().updateUnsafe(_684_v170,p0);
          let _index105 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_660_v160).length));
          (_660_v160)[_index105] = (((_685_v171).contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-291)), ((_688_v0) => function (_689_i21) {
            return _688_v0;
          })(_428_v0)))) ? ((_685_v171).get(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-291)), ((_686_v0) => function (_687_i21) {
            return _686_v0;
          })(_428_v0)))) : (((_dafny.ZERO).minus(_682_cf8)).minus(_682_cf8)));
          let _690_v172;
          let _nw109 = Array((new BigNumber(10)).toNumber()).fill(false);
          _690_v172 = _nw109;
          let _index106 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_690_v172).length));
          (_690_v172)[_index106] = (_680_cf10) === (true);
          let _index107 = _module.__default.safeIndex(new BigNumber(535), new BigNumber((_690_v172).length));
          (_690_v172)[_index107] = _456_v20;
          let _index108 = _module.__default.safeIndex(new BigNumber(240), new BigNumber((_660_v160).length));
          (_660_v160)[_index108] = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_660_v160)[_module.__default.safeIndex(new BigNumber(240), new BigNumber((_660_v160).length))], new BigNumber(-912)), (new BigNumber((_428_v0).length)).multipliedBy(p0));
        } else if (_source8.is_DC3) {
          let _691___mcc_h23 = (_source8).cf4;
          let _692_cf4 = _691___mcc_h23;
          let _693_v173;
          _693_v173 = new BigNumber(602);
          let _694_v174;
          _694_v174 = _module.D15.create_DC38(((_456_v20) ? (_module.__default.fm51(_456_v20, _693_v173, _693_v173, globalState)) : (_657_v157)));
          let _695_v175;
          _695_v175 = _dafny.Map.Empty.slice().updateUnsafe(_456_v20,_457_v21);
          let _rhs67 = new BigNumber((((_456_v20) ? ((((_695_v175).contains(_456_v20)) ? ((_695_v175).get(_456_v20)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(10)), ((_696_p0) => function (_697_i22) {
            return _696_p0;
          })(p0))))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-820)), ((_698_v173) => function (_699_i23) {
            return _698_v173;
          })(_693_v173))))).length);
          let _rhs68 = (_692_cf4).Intersect(_692_cf4);
          let _rhs69 = _694_v174;
          let _rhs70 = _456_v20;
          _693_v173 = _rhs67;
          _692_cf4 = _rhs68;
          _694_v174 = _rhs69;
          _456_v20 = _rhs70;
          _693_v173 = ((p0).plus((_dafny.ZERO).minus(p0))).plus(_693_v173);
          let _700_v176;
          let _nw110 = Array((new BigNumber(18)).toNumber()).fill(false);
          _700_v176 = _nw110;
          let _701_v177;
          _701_v177 = _dafny.MultiSet.fromElements(_456_v20);
          let _index109 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_700_v176).length));
          (_700_v176)[_index109] = (_701_v177).IsProperSubsetOf(_dafny.MultiSet.fromElements(_456_v20, _456_v20));
          let _index110 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_700_v176).length));
          (_700_v176)[_index110] = _456_v20;
          let _702_v178;
          _702_v178 = _dafny.Map.Empty.slice().updateUnsafe(p0,_456_v20);
          _702_v178 = _702_v178;
        } else {
          let _703___mcc_h24 = (_source8).cf11;
          let _704_cf11 = _703___mcc_h24;
          let _705_v179;
          _705_v179 = _dafny.Seq.of(_456_v20, _456_v20);
          let _706_v180;
          _706_v180 = _dafny.Map.Empty.slice().updateUnsafe(_657_v157,new BigNumber(68));
          let _707_v181;
          _707_v181 = _module.D15.create_DC38(_657_v157);
          let _708_v182;
          _708_v182 = _dafny.MultiSet.fromElements(p0, new BigNumber(317));
          let _709_v183;
          let _nw111 = Array((new BigNumber(9)).toNumber());
          _nw111[(_dafny.ZERO).toNumber()] = _706_v180;
          _nw111[(_dafny.ONE).toNumber()] = _module.__default.fm52(true, _707_v181, _456_v20, globalState);
          _nw111[(new BigNumber(2)).toNumber()] = _706_v180;
          _nw111[(new BigNumber(3)).toNumber()] = _706_v180;
          _nw111[(new BigNumber(4)).toNumber()] = _706_v180;
          _nw111[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_657_v157,_module.__default.fm44(_456_v20, globalState));
          _nw111[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_657_v157,(((_708_v182).contains(p0)) ? ((_708_v182).get(p0)) : (p0)));
          _nw111[(new BigNumber(7)).toNumber()] = _706_v180;
          _nw111[(new BigNumber(8)).toNumber()] = _706_v180;
          _709_v183 = _nw111;
          let _710_v184;
          let _nw112 = new _module.C1();
          _nw112.__ctor(new BigNumber((_705_v179).length), _709_v183, _457_v21, _module.D0.create_DC1(p0));
          _710_v184 = _nw112;
          let _711_v185;
          _711_v185 = _dafny.Map.Empty.slice().updateUnsafe(_710_v184,new _dafny.CodePoint('b'.codePointAt(0)));
          _711_v185 = (_711_v185).update(_710_v184, new _dafny.CodePoint('e'.codePointAt(0)));
          let _index111 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_660_v160).length));
          (_660_v160)[_index111] = _module.__default.safeModuloInt(p0, new BigNumber((_705_v179).length));
          let _index112 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_660_v160).length));
          (_660_v160)[_index112] = (new BigNumber(602)).minus(new BigNumber(98));
          let _712_v186;
          let _nw113 = new _module.C0();
          _nw113.__ctor(_428_v0);
          _712_v186 = _nw113;
          let _713_v187;
          let _nw114 = Array((new BigNumber(20)).toNumber()).fill(false);
          _713_v187 = _nw114;
          let _index113 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_713_v187).length));
          (_713_v187)[_index113] = ((_456_v20) ? (_456_v20) : (_456_v20));
          let _714_v188;
          _714_v188 = _dafny.MultiSet.fromElements(_705_v179);
          let _index114 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_660_v160).length));
          let _index115 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_713_v187).length));
          let _rhs71 = new BigNumber((_708_v182).cardinality());
          let _rhs72 = _456_v20;
          let _rhs73 = _714_v188;
          let _lhs48 = _660_v160;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(353), new BigNumber((_660_v160).length));
          let _lhs50 = _713_v187;
          let _lhs51 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_713_v187).length));
          _lhs48[_lhs49] = _rhs71;
          _lhs50[_lhs51] = _rhs72;
          _714_v188 = _rhs73;
        }
        let _index116 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length));
        (_660_v160)[_index116] = p0;
        let _715_v189;
        _715_v189 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
        let _716_v190;
        _716_v190 = _dafny.Seq.of(_715_v189, _715_v189, _715_v189, _715_v189);
        let _index117 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length));
        (_660_v160)[_index117] = (p0).plus(new BigNumber((_module.__default.fm53((_716_v190)[_module.__default.safeIndex(p0, new BigNumber((_716_v190).length))], globalState)).length));
        let _717_v191;
        _717_v191 = _dafny.Map.Empty.slice().updateUnsafe(p0,_662_v162);
        let _718_v192;
        _718_v192 = _dafny.Map.Empty.slice().updateUnsafe(_456_v20,_717_v191);
        let _719_v193;
        _719_v193 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))]);
        let _index118 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length));
        let _index119 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length));
        let _index120 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length));
        let _rhs74 = (_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))];
        let _rhs75 = ((!(_715_v189).equals(_715_v189)) ? ((new BigNumber(52)).plus(new BigNumber(499))) : (p0));
        let _rhs76 = ((_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))]).isLessThan((_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))]);
        let _rhs77 = !(_717_v191).equals((((_718_v192).contains(_456_v20)) ? ((_718_v192).get(_456_v20)) : (_dafny.Map.Empty.slice().updateUnsafe((_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))],_module.D2.create_DC4(new BigNumber(438), _456_v20, _661_v161)))));
        let _rhs78 = (p0).multipliedBy(_module.__default.safeDivisionInt((((_719_v193).contains(new BigNumber((_dafny.MultiSet.FromArray(_457_v21)).cardinality()))) ? ((_719_v193).get(new BigNumber((_dafny.MultiSet.FromArray(_457_v21)).cardinality()))) : (p0)), (_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))]));
        let _lhs52 = _660_v160;
        let _lhs53 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length));
        let _lhs54 = _660_v160;
        let _lhs55 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length));
        let _lhs56 = _660_v160;
        let _lhs57 = _module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length));
        _lhs52[_lhs53] = _rhs74;
        _lhs54[_lhs55] = _rhs75;
        _456_v20 = _rhs76;
        _456_v20 = _rhs77;
        _lhs56[_lhs57] = _rhs78;
        let _720_v194;
        _720_v194 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm44(!(_456_v20), globalState),_456_v20);
        let _721_v195;
        _721_v195 = _dafny.Set.fromElements(_657_v157);
        _456_v20 = (((_720_v194).contains(p0)) ? ((_720_v194).get(p0)) : (!((_dafny.Set.fromElements(_657_v157, _657_v157, _657_v157, _657_v157)).IsProperSubsetOf(_721_v195))));
        let _722_v196;
        _722_v196 = _dafny.Seq.of(_456_v20);
        let _723_v197;
        _723_v197 = _dafny.Set.fromElements(_456_v20);
        let _724_v198;
        _724_v198 = _dafny.MultiSet.fromElements(_module.__default.fm44(_456_v20, globalState));
        let _725_v199;
        _725_v199 = _dafny.Seq.of(_module.D13.create_DC32(_428_v0, _456_v20, (_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))], _660_v160));
        let _726_v200;
        let _nw115 = Array((new BigNumber(20)).toNumber());
        _nw115[(_dafny.ZERO).toNumber()] = ((_456_v20) ? (_456_v20) : (_456_v20));
        _nw115[(_dafny.ONE).toNumber()] = true;
        _nw115[(new BigNumber(2)).toNumber()] = (_722_v196)[_module.__default.safeIndex((_dafny.ZERO).minus((_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))]), new BigNumber((_722_v196).length))];
        _nw115[(new BigNumber(3)).toNumber()] = !(_456_v20);
        _nw115[(new BigNumber(4)).toNumber()] = _456_v20;
        _nw115[(new BigNumber(5)).toNumber()] = false;
        _nw115[(new BigNumber(6)).toNumber()] = (_723_v197).IsDisjointFrom(_dafny.Set.fromElements(_456_v20, _456_v20));
        _nw115[(new BigNumber(7)).toNumber()] = false;
        _nw115[(new BigNumber(8)).toNumber()] = _456_v20;
        _nw115[(new BigNumber(9)).toNumber()] = _456_v20;
        _nw115[(new BigNumber(10)).toNumber()] = (_dafny.MultiSet.fromElements((((_715_v189).contains(_456_v20)) ? ((_715_v189).get(_456_v20)) : (p0)), new BigNumber((_725_v199).length))).IsProperSubsetOf(_724_v198);
        _nw115[(new BigNumber(11)).toNumber()] = _456_v20;
        _nw115[(new BigNumber(12)).toNumber()] = ((_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))]).isLessThanOrEqualTo(p0);
        _nw115[(new BigNumber(13)).toNumber()] = (((_720_v194).contains((_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))])) ? ((_720_v194).get((_660_v160)[_module.__default.safeIndex(new BigNumber(653), new BigNumber((_660_v160).length))])) : (_456_v20));
        _nw115[(new BigNumber(14)).toNumber()] = (_this).fm42(_456_v20, globalState);
        _nw115[(new BigNumber(15)).toNumber()] = ((_456_v20) ? (_456_v20) : (_456_v20));
        _nw115[(new BigNumber(16)).toNumber()] = ((_456_v20) ? (_456_v20) : (_456_v20));
        _nw115[(new BigNumber(17)).toNumber()] = !(_456_v20);
        _nw115[(new BigNumber(18)).toNumber()] = _456_v20;
        _nw115[(new BigNumber(19)).toNumber()] = _dafny.Seq.contains(_428_v0, _657_v157);
        _726_v200 = _nw115;
        let _index121 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_726_v200).length));
        (_726_v200)[_index121] = _456_v20;
        let _index122 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_726_v200).length));
        (_726_v200)[_index122] = (p0).isEqualTo(p0);
      }
      let _727_v201;
      _727_v201 = _module.D4.create_DC10(p0, p0, p0);
      let _pat_let_tv10 = _456_v20;
      let _pat_let_tv11 = _457_v21;
      let _pat_let_tv12 = _456_v20;
      let _pat_let_tv13 = _457_v21;
      let _pat_let_tv14 = _456_v20;
      let _pat_let_tv15 = _456_v20;
      let _pat_let_tv16 = _456_v20;
      let _pat_let_tv17 = _456_v20;
      let _pat_let_tv18 = _456_v20;
      let _pat_let_tv19 = p0;
      let _pat_let_tv20 = _456_v20;
      let _pat_let_tv21 = _456_v20;
      let _pat_let_tv22 = p0;
      let _pat_let_tv23 = _456_v20;
      let _pat_let_tv24 = p0;
      let _pat_let_tv25 = p0;
      r0 = function (_source9) {
        if (_source9.is_DC9) {
          let _728___mcc_h25 = (_source9).cf14;
          let _729___mcc_h26 = (_source9).cf15;
          let _730___mcc_h27 = (_source9).cf16;
          let _731___mcc_h28 = (_source9).cf17;
          let _732___mcc_h29 = (_source9).cf18;
          let _733_cf18 = _732___mcc_h29;
          let _734_cf17 = _731___mcc_h28;
          let _735_cf16 = _730___mcc_h27;
          let _736_cf15 = _729___mcc_h26;
          let _737_cf14 = _728___mcc_h25;
          return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv10,_dafny.MultiSet.FromArray(_pat_let_tv11));
        } else if (_source9.is_DC10) {
          let _738___mcc_h30 = (_source9).cf19;
          let _739___mcc_h31 = (_source9).cf20;
          let _740___mcc_h32 = (_source9).cf21;
          let _741_cf21 = _740___mcc_h32;
          let _742_cf20 = _739___mcc_h31;
          let _743_cf19 = _738___mcc_h30;
          return _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv12,_dafny.MultiSet.FromArray((_module.D2.create_DC5(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(842)), function (_744_i24) {
  return new _dafny.CodePoint('c'.codePointAt(0));
})).length), _pat_let_tv13, _pat_let_tv14)).dtor_cf9));
        } else if (_source9.is_DC11) {
          let _745___mcc_h33 = (_source9).cf22;
          let _746___mcc_h34 = (_source9).cf23;
          let _747___mcc_h35 = (_source9).cf24;
          let _748___mcc_h36 = (_source9).cf25;
          let _749___mcc_h37 = (_source9).cf26;
          let _750_cf26 = _749___mcc_h37;
          let _751_cf25 = _748___mcc_h36;
          let _752_cf24 = _747___mcc_h35;
          let _753_cf23 = _746___mcc_h34;
          let _754_cf22 = _745___mcc_h33;
          return _dafny.Map.Empty.slice().updateUnsafe(_750_cf26,_dafny.MultiSet.fromElements(_754_cf22));
        } else if (_source9.is_DC8) {
          let _755___mcc_h38 = (_source9).cf13;
          let _756_cf13 = _755___mcc_h38;
          return (_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv15,_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(true)).length),_pat_let_tv16),_pat_let_tv17)).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv18,_dafny.MultiSet.fromElements(_pat_let_tv19, new BigNumber((_dafny.Seq.of(_pat_let_tv20, _pat_let_tv21)).length))));
        } else {
          let _757___mcc_h39 = (_source9).cf27;
          let _758_cf27 = _757___mcc_h39;
          return (_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.MultiSet.fromElements(_pat_let_tv22))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv23,_dafny.MultiSet.fromElements(_pat_let_tv24)));
        }
      }(((_456_v20) ? (function (_pat_let14_0) {
        return function (_759_dt__update__tmp_h4) {
          return function (_pat_let15_0) {
            return function (_760_dt__update_hcf21_h0) {
              return _module.D4.create_DC10((_759_dt__update__tmp_h4).dtor_cf19, (_759_dt__update__tmp_h4).dtor_cf20, _760_dt__update_hcf21_h0);
            }(_pat_let15_0);
          }(_pat_let_tv25);
        }(_pat_let14_0);
      }(_727_v201)) : (_727_v201)));
      return r0;
    }
    m14(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      r0 = _module.__default.fm44(p1, globalState);
      let _761_v0;
      _761_v0 = _dafny.Seq.UnicodeFromString("wnm");
      let _762_v1;
      _762_v1 = _dafny.Map.Empty.slice().updateUnsafe(p3,true);
      let _hi3 = (new BigNumber(293)).plus(new BigNumber((_762_v1).length));
      for (let _763_i0 = (_dafny.ZERO).minus(new BigNumber((_761_v0).length)); _763_i0.isLessThan(_hi3); _763_i0 = _763_i0.plus(_dafny.ONE)) {
        let _764_v2;
        let _nw116 = Array((new BigNumber(8)).toNumber()).fill(false);
        _764_v2 = _nw116;
        let _index123 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_764_v2).length));
        (_764_v2)[_index123] = p1;
        let _index124 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_764_v2).length));
        (_764_v2)[_index124] = (p1) || (((p0) ? (false) : (p1)));
        let _765_v3;
        _765_v3 = _dafny.Set.fromElements(p0);
        let _766_v4;
        let _nw117 = Array((new BigNumber(18)).toNumber());
        _nw117[(_dafny.ZERO).toNumber()] = _765_v3;
        _nw117[(_dafny.ONE).toNumber()] = _765_v3;
        _nw117[(new BigNumber(2)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(3)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(4)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(5)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(6)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(7)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(8)).toNumber()] = _dafny.Set.fromElements((_764_v2)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_764_v2).length))], p0);
        _nw117[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements(p0);
        _nw117[(new BigNumber(10)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(11)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(12)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(13)).toNumber()] = _dafny.Set.fromElements((_764_v2)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_764_v2).length))]);
        _nw117[(new BigNumber(14)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(15)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(16)).toNumber()] = _765_v3;
        _nw117[(new BigNumber(17)).toNumber()] = _765_v3;
        _766_v4 = _nw117;
        let _767_v5;
        _767_v5 = _module.D0.create_DC0(_766_v4, new BigNumber(2));
        let _768_v6;
        _768_v6 = _dafny.Seq.of(_767_v5, _767_v5);
        let _769_v7;
        _769_v7 = _module.D4.create_DC8(_768_v6);
        _769_v7 = _769_v7;
        let _770_v8;
        _770_v8 = new _dafny.CodePoint('n'.codePointAt(0));
        let _771_v9;
        _771_v9 = _dafny.Map.Empty.slice().updateUnsafe(_770_v8,p3);
        let _772_v11;
        _772_v11 = ((p1) ? (_771_v9) : (function () {
          let _coll40 = new _dafny.Map();
          for (const _compr_40 of (_761_v0).Elements) {
            let _773_v10 = _compr_40;
            if (_dafny.Seq.contains(_761_v0, _773_v10)) {
              _coll40.push([_773_v10,p3]);
            }
          }
          return _coll40;
        }()));
        _772_v11 = _module.__default.fm54((_764_v2)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_764_v2).length))], _dafny.Seq.of(p1, (_764_v2)[_module.__default.safeIndex(new BigNumber(332), new BigNumber((_764_v2).length))]), globalState);
        r0 = (_dafny.ZERO).minus(p3);
      }
      let _774_v12;
      _774_v12 = false;
      _774_v12 = p0;
      let _hi4 = ((p1) ? (p3) : (p3));
      for (let _775_i1 = p3; _775_i1.isLessThan(_hi4); _775_i1 = _775_i1.plus(_dafny.ONE)) {
        let _776_v13;
        _776_v13 = new _dafny.CodePoint('g'.codePointAt(0));
        _776_v13 = ((_774_v12) ? (_776_v13) : (_776_v13));
        _774_v12 = _dafny.Seq.IsPrefixOf(_761_v0, _761_v0);
        _761_v0 = ((p1) ? (_dafny.Seq.UnicodeFromString("w")) : (_761_v0));
        let _777_v14;
        let _nw118 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
        _777_v14 = _nw118;
        let _index125 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_777_v14).length));
        (_777_v14)[_index125] = p3;
        let _778_v15;
        _778_v15 = _dafny.Seq.of(false);
        let _index126 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_777_v14).length));
        let _rhs79 = (new BigNumber(((_762_v1).Merge(_762_v1)).length)).minus(_775_i1);
        let _rhs80 = _775_i1;
        let _rhs81 = _dafny.Seq.Concat(_778_v15, _dafny.Seq.of((_module.D14.create_DC35(p3, _dafny.Seq.UnicodeFromString("sia"), _775_i1, _774_v12, p0)).dtor_cf67, p0, p1, _774_v12));
        let _lhs58 = _777_v14;
        let _lhs59 = _module.__default.safeIndex(new BigNumber(783), new BigNumber((_777_v14).length));
        _lhs58[_lhs59] = _rhs79;
        r0 = _rhs80;
        _778_v15 = _rhs81;
      }
      let _779_v16;
      _779_v16 = _dafny.Set.fromElements(p3, p3);
      let _780_v17;
      _780_v17 = _module.D6.create_DC15(new BigNumber((_779_v16).length), p3, p0);
      let _781_v18;
      _781_v18 = _module.D6.create_DC16(_780_v17);
      let _782_v19;
      _782_v19 = _dafny.MultiSet.fromElements(_781_v18, _module.D6.create_DC16((_module.D6.create_DC16(_780_v17)).dtor_cf33));
      _774_v12 = !(!_dafny.Seq.contains(((!(p1)) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(204)), ((_783_p3) => function (_784_i2) {
        return _783_p3;
      })(p3))) : (_dafny.Seq.of(new BigNumber((_782_v19).cardinality()), p3, p3, p3))), ((p0) ? (p3) : (p3))));
      _761_v0 = _761_v0;
      r0 = p3;
      return r0;
    }
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(!(_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("hi")).length), new BigNumber((_dafny.Seq.of(new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('f'.codePointAt(0)))).length))).length)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("fqgtkbdg")).length), new BigNumber((_dafny.Set.fromElements(new BigNumber(-285), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('w'.codePointAt(0)),new _dafny.CodePoint('k'.codePointAt(0)))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-394),new BigNumber(772))).length))).length))), (new BigNumber(622)).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber(-779)))))));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _785_v0;
      _785_v0 = false;
      if (_785_v0) {
        let _786_v1;
        let _787_v2;
        let _788_v3;
        let _789_v4;
        let _out21;
        let _out22;
        let _out23;
        let _out24;
        let _outcollector7 = (_this).m12(!(_785_v0) || (_785_v0), globalState);
        _out21 = _outcollector7[0];
        _out22 = _outcollector7[1];
        _out23 = _outcollector7[2];
        _out24 = _outcollector7[3];
        _786_v1 = _out21;
        _787_v2 = _out22;
        _788_v3 = _out23;
        _789_v4 = _out24;
        if (_785_v0) {
          let _790_v5;
          let _nw119 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
          _790_v5 = _nw119;
          let _791_v6;
          _791_v6 = new _dafny.CodePoint('g'.codePointAt(0));
          let _792_v7;
          _792_v7 = _module.D0.create_DC1(p0);
          let _793_v8;
          _793_v8 = _dafny.Seq.UnicodeFromString("rnnc");
          let _794_v9;
          _794_v9 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm5(p0, false, _dafny.Map.Empty.slice().updateUnsafe(_791_v6,_792_v7), _789_v4, globalState),_793_v8);
          let _795_v10;
          _795_v10 = _dafny.Set.fromElements(p0, _788_v3);
          let _796_v11;
          _796_v11 = _dafny.MultiSet.fromElements(p0);
          let _797_v12;
          _797_v12 = _dafny.Seq.of(new BigNumber((_794_v9).length), new BigNumber((_795_v10).length), new BigNumber((_796_v11).cardinality()), _module.__default.fm32(p0, globalState), _788_v3);
          let _798_v13;
          let _nw120 = new _module.C1();
          _nw120.__ctor((_dafny.ZERO).minus(new BigNumber(-771)), _790_v5, _797_v12, _792_v7);
          _798_v13 = _nw120;
          let _799_v14;
          _799_v14 = _dafny.Seq.of(_798_v13);
          _799_v14 = _799_v14;
          let _800_v15;
          let _init17 = ((_801_v0) => function (_802_i0) {
            return _module.__default.safeModuloInt(_802_i0, new BigNumber((_dafny.Seq.of(!(_801_v0))).length));
          })(_785_v0);
          let _nw121 = Array((new BigNumber(9)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw121.length); _i0_17++) {
            _nw121[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _800_v15 = _nw121;
          let _index127 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_800_v15).length));
          (_800_v15)[_index127] = _787_v2;
          let _index128 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_800_v15).length));
          (_800_v15)[_index128] = _787_v2;
          let _803_v16;
          _803_v16 = _dafny.Set.fromElements(_789_v4);
          let _804_v17;
          _804_v17 = _module.D12.create_DC27(_803_v16);
          let _805_v18;
          _805_v18 = _dafny.Map.Empty.slice().updateUnsafe(_789_v4,_803_v16);
          _804_v17 = _module.D12.create_DC27((((_805_v18).contains(_785_v0)) ? ((_805_v18).get(_785_v0)) : (_dafny.Set.fromElements(_785_v0))));
          _800_v15 = _800_v15;
          let _806_v19;
          _806_v19 = _module.D6.create_DC14(_793_v8);
          let _807_v20;
          _807_v20 = _dafny.MultiSet.fromElements(_786_v1);
          let _808_v21;
          _808_v21 = _dafny.Map.Empty.slice().updateUnsafe(_788_v3,(((_807_v20).contains(_789_v4)) ? ((_807_v20).get(_789_v4)) : (new BigNumber(262))));
          let _809_v22;
          let _nw122 = Array((new BigNumber(12)).toNumber());
          _nw122[(_dafny.ZERO).toNumber()] = _806_v19;
          _nw122[(_dafny.ONE).toNumber()] = _806_v19;
          _nw122[(new BigNumber(2)).toNumber()] = function (_pat_let16_0) {
            return function (_811_dt__update__tmp_h0) {
              return function (_pat_let17_0) {
                return function (_813_dt__update_hcf29_h0) {
                  return _module.D6.create_DC14(_813_dt__update_hcf29_h0);
                }(_pat_let17_0);
              }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(287)), function (_812_i2) {
                return new _dafny.CodePoint('i'.codePointAt(0));
              }));
            }(_pat_let16_0);
          }(_module.__default.fm33((((_808_v21).contains(_788_v3)) ? ((_808_v21).get(_788_v3)) : (p0)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(911)), function (_810_i1) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          }), globalState));
          _nw122[(new BigNumber(3)).toNumber()] = _806_v19;
          _nw122[(new BigNumber(4)).toNumber()] = _806_v19;
          _nw122[(new BigNumber(5)).toNumber()] = _806_v19;
          _nw122[(new BigNumber(6)).toNumber()] = _806_v19;
          _nw122[(new BigNumber(7)).toNumber()] = _806_v19;
          _nw122[(new BigNumber(8)).toNumber()] = _module.D6.create_DC14(_793_v8);
          _nw122[(new BigNumber(9)).toNumber()] = _806_v19;
          _nw122[(new BigNumber(10)).toNumber()] = _module.D6.create_DC14(_793_v8);
          _nw122[(new BigNumber(11)).toNumber()] = _module.__default.fm33((_800_v15)[_module.__default.safeIndex(new BigNumber(189), new BigNumber((_800_v15).length))], _dafny.Seq.Create(_module.__default.abs(new BigNumber(190)), ((_814_v6) => function (_815_i3) {
            return _814_v6;
          })(_791_v6)), globalState);
          _809_v22 = _nw122;
          let _pat_let_tv26 = _793_v8;
          let _index129 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_809_v22).length));
          (_809_v22)[_index129] = function (_pat_let18_0) {
            return function (_816_dt__update__tmp_h1) {
              return function (_pat_let19_0) {
                return function (_817_dt__update_hcf29_h1) {
                  return _module.D6.create_DC14(_817_dt__update_hcf29_h1);
                }(_pat_let19_0);
              }(_pat_let_tv26);
            }(_pat_let18_0);
          }(_806_v19);
          let _index130 = _module.__default.safeIndex(new BigNumber(402), new BigNumber((_809_v22).length));
          (_809_v22)[_index130] = _806_v19;
        } else {
          let _818_v23;
          let _nw123 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Map.Empty);
          _818_v23 = _nw123;
          _818_v23 = _818_v23;
          _788_v3 = (_dafny.ZERO).minus(p0);
          r1 = _789_v4;
          let _819_v24;
          let _nw124 = Array((new BigNumber(13)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _819_v24 = _nw124;
          let _820_v25;
          _820_v25 = _dafny.Seq.UnicodeFromString("gvm");
          let _index131 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_819_v24).length));
          (_819_v24)[_index131] = _820_v25;
          let _index132 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_819_v24).length));
          (_819_v24)[_index132] = _820_v25;
          let _821_v26;
          _821_v26 = _dafny.Map.Empty.slice().updateUnsafe((_789_v4) === (_785_v0),_785_v0);
          _821_v26 = (_821_v26).Merge(_821_v26);
        }
        let _822_v27;
        _822_v27 = _dafny.Seq.UnicodeFromString("opafnutak");
        let _823_v28;
        _823_v28 = _dafny.Seq.of(_822_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(620)), function (_824_i4) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        }), _822_v27, _822_v27, _822_v27);
        let _825_v29;
        let _nw125 = new _module.C0();
        _nw125.__ctor((_823_v28)[_module.__default.safeIndex(new BigNumber(171), new BigNumber((_823_v28).length))]);
        _825_v29 = _nw125;
        if (true) {
          let _826_v30;
          _826_v30 = _dafny.Seq.of(_789_v4);
          let _827_v31;
          _827_v31 = _dafny.Seq.of(_826_v30, _826_v30);
          let _828_v32;
          _828_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,_788_v3);
          let _829_v33;
          _829_v33 = _dafny.Seq.of(new BigNumber((_822_v27).length), new BigNumber(553), (((_828_v32).contains(new BigNumber(-361))) ? ((_828_v32).get(new BigNumber(-361))) : (p0)), (_dafny.ZERO).minus(_788_v3), _787_v2);
          let _830_v34;
          _830_v34 = _dafny.Set.fromElements(p0);
          let _831_v35;
          let _nw126 = Array((new BigNumber(28)).toNumber());
          _nw126[(_dafny.ZERO).toNumber()] = _788_v3;
          _nw126[(_dafny.ONE).toNumber()] = _787_v2;
          _nw126[(new BigNumber(2)).toNumber()] = _788_v3;
          _nw126[(new BigNumber(3)).toNumber()] = _787_v2;
          _nw126[(new BigNumber(4)).toNumber()] = _787_v2;
          _nw126[(new BigNumber(5)).toNumber()] = p0;
          _nw126[(new BigNumber(6)).toNumber()] = p0;
          _nw126[(new BigNumber(7)).toNumber()] = p0;
          _nw126[(new BigNumber(8)).toNumber()] = _787_v2;
          _nw126[(new BigNumber(9)).toNumber()] = p0;
          _nw126[(new BigNumber(10)).toNumber()] = p0;
          _nw126[(new BigNumber(11)).toNumber()] = _787_v2;
          _nw126[(new BigNumber(12)).toNumber()] = _787_v2;
          _nw126[(new BigNumber(13)).toNumber()] = new BigNumber(356);
          _nw126[(new BigNumber(14)).toNumber()] = _787_v2;
          _nw126[(new BigNumber(15)).toNumber()] = _788_v3;
          _nw126[(new BigNumber(16)).toNumber()] = new BigNumber(-723);
          _nw126[(new BigNumber(17)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray((_827_v31)[_module.__default.safeIndex(_788_v3, new BigNumber((_827_v31).length))])).cardinality());
          _nw126[(new BigNumber(18)).toNumber()] = (_829_v33)[_module.__default.safeIndex(_788_v3, new BigNumber((_829_v33).length))];
          _nw126[(new BigNumber(19)).toNumber()] = _788_v3;
          _nw126[(new BigNumber(20)).toNumber()] = p0;
          _nw126[(new BigNumber(21)).toNumber()] = _788_v3;
          _nw126[(new BigNumber(22)).toNumber()] = _787_v2;
          _nw126[(new BigNumber(23)).toNumber()] = new BigNumber((_830_v34).length);
          _nw126[(new BigNumber(24)).toNumber()] = _788_v3;
          _nw126[(new BigNumber(25)).toNumber()] = p0;
          _nw126[(new BigNumber(26)).toNumber()] = _787_v2;
          _nw126[(new BigNumber(27)).toNumber()] = p0;
          _831_v35 = _nw126;
          let _832_v36;
          _832_v36 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(true)).length),_831_v35);
          let _833_v37;
          _833_v37 = _module.D12.create_DC28(_785_v0, _787_v2);
          let _834_v39;
          let _init18 = ((_835_p0) => function (_836_i5) {
            return function () {
              let _coll41 = new _dafny.Map();
              for (const _compr_41 of (_dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))).Elements) {
                let _837_v38 = _compr_41;
                if ((_dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0)))).contains(_837_v38)) {
                  _coll41.push([_837_v38,_835_p0]);
                }
              }
              return _coll41;
            }();
          })(p0);
          let _nw127 = Array((new BigNumber(28)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw127.length); _i0_18++) {
            _nw127[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _834_v39 = _nw127;
          let _838_v40;
          _838_v40 = _module.D0.create_DC1(p0);
          let _839_v41;
          let _nw128 = new _module.C1();
          _nw128.__ctor(p0, _834_v39, _829_v33, _838_v40);
          _839_v41 = _nw128;
          let _840_v42;
          _840_v42 = _dafny.Map.Empty.slice().updateUnsafe(_839_v41,false);
          let _841_v43;
          _841_v43 = _dafny.Set.fromElements(!(_786_v1));
          let _842_v44;
          let _nw129 = Array((new BigNumber(20)).toNumber());
          _nw129[(_dafny.ZERO).toNumber()] = _789_v4;
          _nw129[(_dafny.ONE).toNumber()] = (_788_v3).isLessThan(_788_v3);
          _nw129[(new BigNumber(2)).toNumber()] = _786_v1;
          _nw129[(new BigNumber(3)).toNumber()] = (_787_v2).isLessThan(new BigNumber((_832_v36).length));
          _nw129[(new BigNumber(4)).toNumber()] = false;
          _nw129[(new BigNumber(5)).toNumber()] = !(_785_v0);
          _nw129[(new BigNumber(6)).toNumber()] = !(_786_v1);
          _nw129[(new BigNumber(7)).toNumber()] = _786_v1;
          _nw129[(new BigNumber(8)).toNumber()] = ((_789_v4) ? ((_833_v37).dtor_cf46) : (_786_v1));
          _nw129[(new BigNumber(9)).toNumber()] = _786_v1;
          _nw129[(new BigNumber(10)).toNumber()] = !(_dafny.Map.Empty.slice().updateUnsafe(_840_v42,_841_v43)).contains(_840_v42);
          _nw129[(new BigNumber(11)).toNumber()] = _786_v1;
          _nw129[(new BigNumber(12)).toNumber()] = _789_v4;
          _nw129[(new BigNumber(13)).toNumber()] = !(false) || (_789_v4);
          _nw129[(new BigNumber(14)).toNumber()] = _786_v1;
          _nw129[(new BigNumber(15)).toNumber()] = (_839_v41.f9).isLessThanOrEqualTo(_788_v3);
          _nw129[(new BigNumber(16)).toNumber()] = false;
          _nw129[(new BigNumber(17)).toNumber()] = (_830_v34).IsSubsetOf(_830_v34);
          _nw129[(new BigNumber(18)).toNumber()] = !(((_789_v4) ? (_786_v1) : (_785_v0)));
          _nw129[(new BigNumber(19)).toNumber()] = _786_v1;
          _842_v44 = _nw129;
          let _index133 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_842_v44).length));
          (_842_v44)[_index133] = _789_v4;
          let _index134 = _module.__default.safeIndex(new BigNumber(637), new BigNumber((_842_v44).length));
          (_842_v44)[_index134] = _785_v0;
          _788_v3 = _787_v2;
          let _843_v45;
          _843_v45 = _dafny.Map.Empty.slice().updateUnsafe(_785_v0,(_839_v41.f9).plus((_module.__default.fm34(_786_v1, _785_v0, _module.__default.fm30(_839_v41.f9, (_842_v44)[_module.__default.safeIndex(new BigNumber(637), new BigNumber((_842_v44).length))], (_842_v44)[_module.__default.safeIndex(new BigNumber(637), new BigNumber((_842_v44).length))], globalState), new BigNumber(369), globalState)).dtor_cf47));
          _788_v3 = new BigNumber((_843_v45).length);
          (_839_v41).f9 = new BigNumber((_825_v29.f8).length);
          let _844_v46;
          let _nw130 = new _module.C0();
          _nw130.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-8)), ((_845_v2, _846_v30) => function (_847_i6) {
            return (_module.D11.create_DC25(new _dafny.CodePoint('w'.codePointAt(0)), _845_v2, (_846_v30)[_module.__default.safeIndex(new BigNumber(593), new BigNumber((_846_v30).length))])).dtor_cf41;
          })(_787_v2, _826_v30)));
          _844_v46 = _nw130;
        } else {
          let _848_v47;
          _848_v47 = _dafny.MultiSet.fromElements(_786_v1);
          let _849_v48;
          _849_v48 = _dafny.MultiSet.fromElements(new BigNumber((_848_v47).cardinality()));
          let _850_v49;
          let _init19 = ((_851_v1) => function (_852_i7) {
            return _851_v1;
          })(_786_v1);
          let _nw131 = Array((new BigNumber(16)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw131.length); _i0_19++) {
            _nw131[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _850_v49 = _nw131;
          let _853_v50;
          _853_v50 = _module.D12.create_DC28(_786_v1, new BigNumber(524));
          let _854_v51;
          _854_v51 = new _dafny.CodePoint('y'.codePointAt(0));
          let _855_v52;
          _855_v52 = _dafny.Map.Empty.slice().updateUnsafe(_785_v0,new _dafny.CodePoint('c'.codePointAt(0)));
          let _856_v53;
          _856_v53 = _module.D0.create_DC1(new BigNumber((_855_v52).length));
          let _857_v54;
          _857_v54 = _dafny.Map.Empty.slice().updateUnsafe(_854_v51,_856_v53);
          let _858_v55;
          _858_v55 = _dafny.Map.Empty.slice().updateUnsafe(_786_v1,_module.__default.fm30((_dafny.ZERO).minus((_dafny.ZERO).minus(p0)), (_this).fm5(_787_v2, _789_v4, _857_v54, _785_v0, globalState), _789_v4, globalState));
          let _859_v56;
          _859_v56 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_787_v2),_786_v1);
          let _860_v57;
          _860_v57 = _dafny.Map.Empty.slice().updateUnsafe(_785_v0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_788_v3,_786_v1)).length));
          let _861_v58;
          let _nw132 = Array((new BigNumber(22)).toNumber());
          _nw132[(_dafny.ZERO).toNumber()] = _787_v2;
          _nw132[(_dafny.ONE).toNumber()] = _787_v2;
          _nw132[(new BigNumber(2)).toNumber()] = _788_v3;
          _nw132[(new BigNumber(3)).toNumber()] = _787_v2;
          _nw132[(new BigNumber(4)).toNumber()] = (_module.__default.fm32(p0, globalState)).plus(p0);
          _nw132[(new BigNumber(5)).toNumber()] = _module.__default.safeDivisionInt(_787_v2, (_dafny.ZERO).minus(_787_v2));
          _nw132[(new BigNumber(6)).toNumber()] = _787_v2;
          _nw132[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw132[(new BigNumber(8)).toNumber()] = _788_v3;
          _nw132[(new BigNumber(9)).toNumber()] = new BigNumber((_849_v48).cardinality());
          _nw132[(new BigNumber(10)).toNumber()] = p0;
          _nw132[(new BigNumber(11)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_850_v49, _850_v49, _850_v49)).cardinality());
          _nw132[(new BigNumber(12)).toNumber()] = p0;
          _nw132[(new BigNumber(13)).toNumber()] = (_853_v50).dtor_cf47;
          _nw132[(new BigNumber(14)).toNumber()] = new BigNumber((_858_v55).length);
          _nw132[(new BigNumber(15)).toNumber()] = (_787_v2).plus(new BigNumber((_859_v56).length));
          _nw132[(new BigNumber(16)).toNumber()] = p0;
          _nw132[(new BigNumber(17)).toNumber()] = _787_v2;
          _nw132[(new BigNumber(18)).toNumber()] = new BigNumber((_860_v57).length);
          _nw132[(new BigNumber(19)).toNumber()] = p0;
          _nw132[(new BigNumber(20)).toNumber()] = (_dafny.ZERO).minus(p0);
          _nw132[(new BigNumber(21)).toNumber()] = (_787_v2).minus(_787_v2);
          _861_v58 = _nw132;
          let _index135 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_861_v58).length));
          (_861_v58)[_index135] = ((!(_786_v1)) ? (_788_v3) : (_787_v2));
          let _index136 = _module.__default.safeIndex(new BigNumber(289), new BigNumber((_861_v58).length));
          (_861_v58)[_index136] = (_module.__default.fm32(new BigNumber(174), globalState)).plus(p0);
          let _862_v59;
          let _nw133 = Array((new BigNumber(8)).toNumber()).fill(_dafny.MultiSet.Empty);
          _862_v59 = _nw133;
          let _863_v60;
          _863_v60 = _dafny.MultiSet.fromElements(_854_v51);
          let _rhs82 = p0;
          let _rhs83 = _862_v59;
          let _rhs84 = _863_v60;
          _788_v3 = _rhs82;
          _862_v59 = _rhs83;
          _863_v60 = _rhs84;
          let _864_v61;
          _864_v61 = _dafny.Seq.of(_786_v1, _789_v4, _785_v0);
          _864_v61 = _dafny.Seq.update(((_786_v1) ? (_864_v61) : (_864_v61)), _module.__default.safeIndex(p0, new BigNumber((((_786_v1) ? (_864_v61) : (_864_v61))).length)), ((_861_v58)[_module.__default.safeIndex(new BigNumber(289), new BigNumber((_861_v58).length))]).isLessThanOrEqualTo(new BigNumber((_859_v56).length)));
          let _865_v62;
          _865_v62 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(689),_854_v51);
          let _866_v63;
          _866_v63 = _dafny.Map.Empty.slice().updateUnsafe(_865_v62,_787_v2);
          _866_v63 = (_866_v63).update(_865_v62, (_dafny.ZERO).minus(_787_v2));
          _822_v27 = _822_v27;
        }
        let _867_v64;
        let _nw134 = new _module.C0();
        _nw134.__ctor(_822_v27);
        _867_v64 = _nw134;
      } else {
        let _868_v65;
        _868_v65 = _dafny.Seq.of(_785_v0);
        let _869_v66;
        _869_v66 = _dafny.MultiSet.fromElements(_785_v0, true, _785_v0, _785_v0, (_868_v65)[_module.__default.safeIndex(p0, new BigNumber((_868_v65).length))]);
        let _870_v67;
        _870_v67 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,p0)).length),_module.__default.fm35(_785_v0, false, globalState));
        let _871_v69;
        _871_v69 = _dafny.Seq.of(p0, p0, p0);
        _785_v0 = ((((_869_v66).contains(!(_785_v0))) ? ((_869_v66).get(!(_785_v0))) : (p0))).isLessThanOrEqualTo(new BigNumber(((_870_v67).Merge(function () {
          let _coll42 = new _dafny.Map();
          for (const _compr_42 of (_871_v69).Elements) {
            let _872_v68 = _compr_42;
            if (_dafny.Seq.contains(_871_v69, _872_v68)) {
              _coll42.push([(_872_v68).minus((_dafny.ZERO).minus(p0)),_871_v69]);
            }
          }
          return _coll42;
        }())).length));
        let _873_v70;
        _873_v70 = _dafny.Seq.UnicodeFromString("is");
        let _874_v71;
        _874_v71 = new _dafny.CodePoint('v'.codePointAt(0));
        let _875_v72;
        _875_v72 = _dafny.Seq.of((_873_v70)[_module.__default.safeIndex(p0, new BigNumber((_873_v70).length))], (_873_v70)[_module.__default.safeIndex(p0, new BigNumber((_873_v70).length))], _874_v71, _874_v71);
        let _876_v73;
        _876_v73 = _dafny.Set.fromElements(_785_v0);
        let _877_v74;
        _877_v74 = _module.D11.create_DC25((_873_v70)[_module.__default.safeIndex(new BigNumber((_876_v73).length), new BigNumber((_873_v70).length))], p0, _785_v0);
        let _878_v75;
        _878_v75 = _dafny.Map.Empty.slice().updateUnsafe(_785_v0,p0);
        let _879_v76;
        _879_v76 = _module.D12.create_DC29(_785_v0, p0, false, (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_878_v75).length))), _785_v0);
        let _source10 = (((_785_v0) === ((_877_v74).dtor_cf43)) ? (_879_v76) : (_879_v76));
        if (_source10.is_DC28) {
          let _880___mcc_h0 = (_source10).cf46;
          let _881___mcc_h1 = (_source10).cf47;
          let _882_cf47 = _881___mcc_h1;
          let _883_cf46 = _880___mcc_h0;
          let _884_v77;
          _884_v77 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_874_v71);
          let _885_v78;
          let _nw135 = Array((new BigNumber(22)).toNumber());
          _nw135[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('c'.codePointAt(0));
          _nw135[(_dafny.ONE).toNumber()] = _874_v71;
          _nw135[(new BigNumber(2)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(3)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(4)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(5)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(6)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(7)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(8)).toNumber()] = new _dafny.CodePoint('i'.codePointAt(0));
          _nw135[(new BigNumber(9)).toNumber()] = (((_884_v77).contains(new BigNumber(509))) ? ((_884_v77).get(new BigNumber(509))) : (_874_v71));
          _nw135[(new BigNumber(10)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(11)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(12)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(13)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(14)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(15)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(16)).toNumber()] = new _dafny.CodePoint('h'.codePointAt(0));
          _nw135[(new BigNumber(17)).toNumber()] = new _dafny.CodePoint('b'.codePointAt(0));
          _nw135[(new BigNumber(18)).toNumber()] = _module.__default.fm29(_883_cf46, globalState);
          _nw135[(new BigNumber(19)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(20)).toNumber()] = _874_v71;
          _nw135[(new BigNumber(21)).toNumber()] = _874_v71;
          _885_v78 = _nw135;
          let _886_v79;
          _886_v79 = _dafny.Map.Empty.slice().updateUnsafe(_885_v78,_876_v73);
          _886_v79 = _886_v79;
          let _887_v80;
          let _nw136 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
          _887_v80 = _nw136;
          let _888_v81;
          _888_v81 = _dafny.MultiSet.fromElements(new BigNumber(-804));
          let _889_v82;
          let _nw137 = new _module.C0();
          _nw137.__ctor(_875_v72);
          _889_v82 = _nw137;
          let _890_v83;
          _890_v83 = _dafny.Set.fromElements(_889_v82, _889_v82, _889_v82, _889_v82);
          let _891_v84;
          _891_v84 = _module.D4.create_DC9(_888_v81, _874_v71, _882_cf47, _890_v83, new BigNumber(712));
          let _892_v85;
          _892_v85 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Map.Empty.slice().updateUnsafe(true,_module.D4.create_DC9(_888_v81, _874_v71, new BigNumber(207), _890_v83, p0))).update(_785_v0, _891_v84),_dafny.MultiSet.FromArray(_871_v69));
          let _index137 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_887_v80).length));
          (_887_v80)[_index137] = (_892_v85).Merge(_892_v85);
          let _index138 = _module.__default.safeIndex(new BigNumber(682), new BigNumber((_887_v80).length));
          (_887_v80)[_index138] = _892_v85;
          let _893_v86;
          let _nw138 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Set.Empty);
          _893_v86 = _nw138;
          let _894_v87;
          _894_v87 = _module.D0.create_DC0(_893_v86, _882_cf47);
          let _895_v88;
          _895_v88 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm32(p0, globalState),_785_v0);
          let _896_v89;
          _896_v89 = _dafny.Map.Empty.slice().updateUnsafe(_874_v71,new BigNumber((_895_v88).length));
          let _897_v91;
          let _nw139 = Array((new BigNumber(17)).toNumber());
          _nw139[(_dafny.ZERO).toNumber()] = _896_v89;
          _nw139[(_dafny.ONE).toNumber()] = (_896_v89).update(_874_v71, new BigNumber((_875_v72).length));
          _nw139[(new BigNumber(2)).toNumber()] = _module.__default.fm36(_874_v71, (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_878_v75,_874_v71)).length)), _module.__default.fm32(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_785_v0,_882_cf47)).length), globalState), (_module.D12.create_DC27(_876_v73)).dtor_cf45, globalState);
          _nw139[(new BigNumber(3)).toNumber()] = _896_v89;
          _nw139[(new BigNumber(4)).toNumber()] = _896_v89;
          _nw139[(new BigNumber(5)).toNumber()] = _896_v89;
          _nw139[(new BigNumber(6)).toNumber()] = _896_v89;
          _nw139[(new BigNumber(7)).toNumber()] = _896_v89;
          _nw139[(new BigNumber(8)).toNumber()] = _896_v89;
          _nw139[(new BigNumber(9)).toNumber()] = function () {
            let _coll43 = new _dafny.Map();
            for (const _compr_43 of (_873_v70).Elements) {
              let _898_v90 = _compr_43;
              if (_dafny.Seq.contains(_873_v70, _898_v90)) {
                _coll43.push([_898_v90,_882_cf47]);
              }
            }
            return _coll43;
          }();
          _nw139[(new BigNumber(10)).toNumber()] = _896_v89;
          _nw139[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_874_v71,new BigNumber(774));
          _nw139[(new BigNumber(12)).toNumber()] = _896_v89;
          _nw139[(new BigNumber(13)).toNumber()] = _896_v89;
          _nw139[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_874_v71,p0);
          _nw139[(new BigNumber(15)).toNumber()] = _896_v89;
          _nw139[(new BigNumber(16)).toNumber()] = (_896_v89).update(_874_v71, p0);
          _897_v91 = _nw139;
          let _899_v92;
          _899_v92 = _module.D0.create_DC1(new BigNumber((_895_v88).length));
          let _900_v93;
          let _nw140 = new _module.C1();
          _nw140.__ctor((_894_v87).dtor_cf1, _897_v91, _dafny.Seq.of(p0, (_dafny.ZERO).minus(p0)), _899_v92);
          _900_v93 = _nw140;
          let _901_v94;
          let _nw141 = Array((new BigNumber(22)).toNumber());
          _nw141[(_dafny.ZERO).toNumber()] = _785_v0;
          _nw141[(_dafny.ONE).toNumber()] = _785_v0;
          _nw141[(new BigNumber(2)).toNumber()] = _883_cf46;
          _nw141[(new BigNumber(3)).toNumber()] = _785_v0;
          _nw141[(new BigNumber(4)).toNumber()] = true;
          _nw141[(new BigNumber(5)).toNumber()] = _785_v0;
          _nw141[(new BigNumber(6)).toNumber()] = _785_v0;
          _nw141[(new BigNumber(7)).toNumber()] = false;
          _nw141[(new BigNumber(8)).toNumber()] = _785_v0;
          _nw141[(new BigNumber(9)).toNumber()] = _883_cf46;
          _nw141[(new BigNumber(10)).toNumber()] = _883_cf46;
          _nw141[(new BigNumber(11)).toNumber()] = !(_883_cf46);
          _nw141[(new BigNumber(12)).toNumber()] = _883_cf46;
          _nw141[(new BigNumber(13)).toNumber()] = _883_cf46;
          _nw141[(new BigNumber(14)).toNumber()] = _785_v0;
          _nw141[(new BigNumber(15)).toNumber()] = !(_883_cf46);
          _nw141[(new BigNumber(16)).toNumber()] = _785_v0;
          _nw141[(new BigNumber(17)).toNumber()] = _883_cf46;
          _nw141[(new BigNumber(18)).toNumber()] = _785_v0;
          _nw141[(new BigNumber(19)).toNumber()] = _785_v0;
          _nw141[(new BigNumber(20)).toNumber()] = false;
          _nw141[(new BigNumber(21)).toNumber()] = _785_v0;
          _901_v94 = _nw141;
          let _902_v95;
          let _903_v96;
          let _904_v97;
          let _905_v98;
          let _out25;
          let _out26;
          let _out27;
          let _out28;
          let _outcollector8 = (_this).m11(_874_v71, _901_v94, _869_v66, _module.__default.safeDivisionInt(_900_v93.f9, _900_v93.f9), globalState);
          _out25 = _outcollector8[0];
          _out26 = _outcollector8[1];
          _out27 = _outcollector8[2];
          _out28 = _outcollector8[3];
          _902_v95 = _out25;
          _903_v96 = _out26;
          _904_v97 = _out27;
          _905_v98 = _out28;
        } else if (_source10.is_DC29) {
          let _906___mcc_h2 = (_source10).cf48;
          let _907___mcc_h3 = (_source10).cf49;
          let _908___mcc_h4 = (_source10).cf50;
          let _909___mcc_h5 = (_source10).cf51;
          let _910___mcc_h6 = (_source10).cf52;
          let _911_cf52 = _910___mcc_h6;
          let _912_cf51 = _909___mcc_h5;
          let _913_cf50 = _908___mcc_h4;
          let _914_cf49 = _907___mcc_h3;
          let _915_cf48 = _906___mcc_h2;
          let _916_v99;
          let _nw142 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Map.Empty);
          _916_v99 = _nw142;
          let _917_v100;
          _917_v100 = _module.D0.create_DC1(_914_cf49);
          let _918_v101;
          let _nw143 = new _module.C1();
          _nw143.__ctor(_914_cf49, _916_v99, _871_v69, _917_v100);
          _918_v101 = _nw143;
          _912_cf51 = new BigNumber(852);
          _914_cf49 = ((_913_cf50) ? (new BigNumber(501)) : (_module.__default.safeDivisionInt(_912_cf51, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(436)), ((_919_v71) => function (_920_i8) {
            return _919_v71;
          })(_874_v71))).length))));
          let _921_v102;
          let _nw144 = new _module.C0();
          _nw144.__ctor(_dafny.Seq.update(_dafny.Seq.Concat(_873_v70, _875_v72), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_873_v70, _875_v72)).length)), new _dafny.CodePoint('p'.codePointAt(0))));
          _921_v102 = _nw144;
        } else {
          let _922___mcc_h7 = (_source10).cf45;
          let _923_cf45 = _922___mcc_h7;
          r1 = (p0).isLessThan(new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("oggyllotw"), _module.__default.safeIndex(new BigNumber((_878_v75).length), new BigNumber((_dafny.Seq.UnicodeFromString("oggyllotw")).length)), _874_v71)).length));
          let _924_v103;
          _924_v103 = new BigNumber(345);
          _924_v103 = p0;
          _874_v71 = _874_v71;
          _924_v103 = _module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(_785_v0, _785_v0, _785_v0, _785_v0)).length), new BigNumber(469));
        }
        if (false) {
          let _925_v104;
          let _nw145 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _925_v104 = _nw145;
          let _926_v105;
          _926_v105 = _dafny.Map.Empty.slice().updateUnsafe(p0,_785_v0);
          let _index139 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_925_v104).length));
          (_925_v104)[_index139] = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.UnicodeFromString("uxuolybl")).length), new BigNumber((_926_v105).length));
          let _927_v106;
          _927_v106 = _dafny.MultiSet.fromElements(_871_v69);
          let _928_v107;
          _928_v107 = _module.D7.create_DC19(new BigNumber((_927_v106).cardinality()));
          let _index140 = _module.__default.safeIndex(new BigNumber(318), new BigNumber((_925_v104).length));
          (_925_v104)[_index140] = _module.__default.safeDivisionInt((p0).plus(new BigNumber((function () {
            let _coll44 = new _dafny.Set();
            for (const _compr_44 of (_dafny.Seq.of(_module.D7.create_DC19((_dafny.ZERO).minus(p0)), _module.__default.fm37(true, globalState), _928_v107, _928_v107, _928_v107)).Elements) {
              let _929_v108 = _compr_44;
              if (_dafny.Seq.contains(_dafny.Seq.of(_module.D7.create_DC19((_dafny.ZERO).minus(p0)), _module.__default.fm37(true, globalState), _928_v107, _928_v107, _928_v107), _929_v108)) {
                _coll44.add(_929_v108);
              }
            }
            return _coll44;
          }()).length)), new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_785_v0,p0)).Merge(_878_v75)).length));
          let _930_v109;
          _930_v109 = _dafny.MultiSet.fromElements(p0);
          let _931_v110;
          let _nw146 = new _module.C0();
          _nw146.__ctor(_875_v72);
          _931_v110 = _nw146;
          let _932_v111;
          _932_v111 = _dafny.Set.fromElements(_931_v110);
          let _933_v112;
          _933_v112 = _module.D4.create_DC9(_dafny.MultiSet.fromElements(p0), (_875_v72)[_module.__default.safeIndex(p0, new BigNumber((_875_v72).length))], new BigNumber(748), _932_v111, new BigNumber(83));
          let _934_v113;
          let _nw147 = Array((new BigNumber(17)).toNumber());
          _nw147[(_dafny.ZERO).toNumber()] = (((_926_v105).contains(new BigNumber(-785))) ? ((_926_v105).get(new BigNumber(-785))) : (_785_v0));
          _nw147[(_dafny.ONE).toNumber()] = _785_v0;
          _nw147[(new BigNumber(2)).toNumber()] = _785_v0;
          _nw147[(new BigNumber(3)).toNumber()] = !_dafny.Seq.contains(_873_v70, _874_v71);
          _nw147[(new BigNumber(4)).toNumber()] = (_930_v109).IsDisjointFrom((_933_v112).dtor_cf14);
          _nw147[(new BigNumber(5)).toNumber()] = _785_v0;
          _nw147[(new BigNumber(6)).toNumber()] = _785_v0;
          _nw147[(new BigNumber(7)).toNumber()] = (_785_v0) === (!(false));
          _nw147[(new BigNumber(8)).toNumber()] = _785_v0;
          _nw147[(new BigNumber(9)).toNumber()] = (_785_v0) && (_785_v0);
          _nw147[(new BigNumber(10)).toNumber()] = _785_v0;
          _nw147[(new BigNumber(11)).toNumber()] = _785_v0;
          _nw147[(new BigNumber(12)).toNumber()] = _785_v0;
          _nw147[(new BigNumber(13)).toNumber()] = false;
          _nw147[(new BigNumber(14)).toNumber()] = _785_v0;
          _nw147[(new BigNumber(15)).toNumber()] = _785_v0;
          _nw147[(new BigNumber(16)).toNumber()] = _785_v0;
          _934_v113 = _nw147;
          let _nw148 = Array((new BigNumber(11)).toNumber()).fill(false);
          _934_v113 = _nw148;
          let _935_v114;
          let _nw149 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
          _935_v114 = _nw149;
          let _936_v115;
          _936_v115 = _dafny.Seq.of(_875_v72, _875_v72, _875_v72);
          let _index141 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_935_v114).length));
          (_935_v114)[_index141] = _936_v115;
          let _index142 = _module.__default.safeIndex(new BigNumber(171), new BigNumber((_935_v114).length));
          (_935_v114)[_index142] = _936_v115;
          (_931_v110).f8 = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kmpvievok"), _875_v72), _931_v110.f8);
          let _937_v116;
          _937_v116 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_931_v110).fm15(_871_v69, (_868_v65)[_module.__default.safeIndex(new BigNumber(980), new BigNumber((_868_v65).length))], globalState)).length),_dafny.Map.Empty.slice().updateUnsafe((((_878_v75).contains(!(_785_v0))) ? ((_878_v75).get(!(_785_v0))) : (p0)),_875_v72));
          let _938_v117;
          _938_v117 = _dafny.Map.Empty.slice().updateUnsafe((_925_v104)[_module.__default.safeIndex(new BigNumber(318), new BigNumber((_925_v104).length))],_931_v110.f8);
          _937_v116 = (_937_v116).update(p0, _938_v117);
        } else {
          _873_v70 = _873_v70;
          _785_v0 = _785_v0;
          _869_v66 = _869_v66;
          r1 = _dafny.Seq.IsProperPrefixOf(_module.__default.fm35(_785_v0, _785_v0, globalState), _871_v69);
          r0 = _785_v0;
        }
        let _939_v118;
        _939_v118 = new BigNumber(128);
        _939_v118 = p0;
        r1 = _785_v0;
      }
      let _940_v119;
      _940_v119 = new BigNumber(-72);
      _940_v119 = _module.__default.safeModuloInt(_940_v119, p0);
      let _941_v120;
      _941_v120 = _dafny.Seq.of(p0, _940_v119, _940_v119, p0, p0);
      let _942_v121;
      _942_v121 = _dafny.MultiSet.fromElements(p0, new BigNumber(348));
      let _943_v122;
      let _nw150 = new _module.C0();
      _nw150.__ctor(_dafny.Seq.UnicodeFromString("jctntdft"));
      _943_v122 = _nw150;
      let _944_v123;
      _944_v123 = _dafny.Set.fromElements(_943_v122);
      let _hi5 = _940_v119;
      for (let _945_i9 = (_941_v120)[_module.__default.safeIndex((_module.D4.create_DC9(_942_v121, new _dafny.CodePoint('n'.codePointAt(0)), p0, _944_v123, p0)).dtor_cf16, new BigNumber((_941_v120).length))]; _945_i9.isLessThan(_hi5); _945_i9 = _945_i9.plus(_dafny.ONE)) {
        let _946_v124;
        let _nw151 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _946_v124 = _nw151;
        let _index143 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_946_v124).length));
        (_946_v124)[_index143] = _940_v119;
        let _947_v125;
        let _nw152 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.of());
        _947_v125 = _nw152;
        let _948_v126;
        _948_v126 = _dafny.Map.Empty.slice().updateUnsafe(_945_i9,p0);
        let _949_v127;
        _949_v127 = _module.D13.create_DC30(_947_v125);
        let _pat_let_tv27 = _947_v125;
        let _index144 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_946_v124).length));
        let _rhs85 = (((_948_v126).contains((_dafny.ZERO).minus(_945_i9))) ? ((_948_v126).get((_dafny.ZERO).minus(_945_i9))) : ((((_942_v121).contains(_945_i9)) ? ((_942_v121).get(_945_i9)) : ((_dafny.ZERO).minus(_940_v119)))));
        let _rhs86 = p0;
        let _rhs87 = (function (_pat_let20_0) {
          return function (_950_dt__update__tmp_h2) {
            return function (_pat_let21_0) {
              return function (_951_dt__update_hcf53_h0) {
                return _module.D13.create_DC30(_951_dt__update_hcf53_h0);
              }(_pat_let21_0);
            }(_pat_let_tv27);
          }(_pat_let20_0);
        }(_949_v127)).dtor_cf53;
        let _lhs60 = _946_v124;
        let _lhs61 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_946_v124).length));
        _940_v119 = _rhs85;
        _lhs60[_lhs61] = _rhs86;
        _947_v125 = _rhs87;
        r1 = !(((!(_785_v0) || (_785_v0)) ? (_785_v0) : (_785_v0)));
        (_943_v122).f8 = _943_v122.f8;
        let _952_v128;
        _952_v128 = _dafny.Seq.of(_941_v120, _941_v120);
        if ((((new BigNumber(((_952_v128)[_module.__default.safeIndex(_940_v119, new BigNumber((_952_v128).length))]).length)).isLessThanOrEqualTo(new BigNumber((_941_v120).length))) ? (!(new BigNumber(-986)).isEqualTo(new BigNumber(792))) : (_785_v0))) {
          let _index145 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_946_v124).length));
          (_946_v124)[_index145] = (_module.__default.safeDivisionInt(_940_v119, (_946_v124)[_module.__default.safeIndex(new BigNumber(200), new BigNumber((_946_v124).length))])).multipliedBy(new BigNumber((_dafny.Seq.of(_785_v0, _785_v0)).length));
          let _953_v129;
          _953_v129 = _dafny.Map.Empty.slice().updateUnsafe(_785_v0,_946_v124);
          _953_v129 = (_953_v129).update(_785_v0, _946_v124);
          (_943_v122).f8 = _dafny.Seq.Concat((_943_v122).fm15(_941_v120, !(_785_v0), globalState), _943_v122.f8);
          _949_v127 = _949_v127;
          let _index146 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_946_v124).length));
          let _rhs88 = _942_v121;
          let _rhs89 = _943_v122.f8;
          let _rhs90 = _945_i9;
          let _rhs91 = ((_785_v0) ? (_dafny.Seq.Concat(_941_v120, _941_v120)) : (_941_v120));
          let _lhs62 = _943_v122;
          let _lhs63 = _946_v124;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_946_v124).length));
          _942_v121 = _rhs88;
          _lhs62.f8 = _rhs89;
          _lhs63[_lhs64] = _rhs90;
          _941_v120 = _rhs91;
        } else {
          (_943_v122).f8 = _943_v122.f8;
          let _954_v130;
          _954_v130 = _dafny.Seq.of(_943_v122.f8);
          let _955_v131;
          _955_v131 = _dafny.Seq.of(!(_785_v0));
          let _rhs92 = _785_v0;
          let _rhs93 = _785_v0;
          let _rhs94 = (_954_v130)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_955_v131, _module.__default.safeIndex(_945_i9, new BigNumber((_955_v131).length)), _785_v0)).length), new BigNumber((_954_v130).length))];
          let _lhs65 = _943_v122;
          _785_v0 = _rhs92;
          _785_v0 = _rhs93;
          _lhs65.f8 = _rhs94;
          let _956_v132;
          let _nw153 = Array((new BigNumber(9)).toNumber()).fill(false);
          _956_v132 = _nw153;
          let _957_v133;
          let _nw154 = Array((new BigNumber(22)).toNumber()).fill([]);
          _957_v133 = _nw154;
          let _958_v134;
          _958_v134 = _module.D2.create_DC4(new BigNumber(289), _785_v0, _957_v133);
          let _959_v135;
          _959_v135 = new _dafny.CodePoint('c'.codePointAt(0));
          let _960_v136;
          _960_v136 = _dafny.MultiSet.fromElements(_785_v0);
          let _961_v137;
          _961_v137 = _module.D0.create_DC1(new BigNumber((_960_v136).cardinality()));
          let _index147 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_956_v132).length));
          (_956_v132)[_index147] = (_this).fm5((_958_v134).dtor_cf5, _785_v0, _dafny.Map.Empty.slice().updateUnsafe(_959_v135,_961_v137), _785_v0, globalState);
          let _index148 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_956_v132).length));
          (_956_v132)[_index148] = ((_785_v0) ? (((_960_v136).update(_785_v0, _module.__default.abs((_946_v124)[_module.__default.safeIndex(new BigNumber(200), new BigNumber((_946_v124).length))]))).IsProperSubsetOf(_960_v136)) : (((_785_v0) ? (!(!(_785_v0))) : (_785_v0))));
          let _index149 = _module.__default.safeIndex(new BigNumber(200), new BigNumber((_946_v124).length));
          (_946_v124)[_index149] = (_946_v124)[_module.__default.safeIndex(new BigNumber(200), new BigNumber((_946_v124).length))];
          let _962_v138;
          let _nw155 = new _module.C0();
          _nw155.__ctor(_943_v122.f8);
          _962_v138 = _nw155;
        }
      }
      let _963_v139;
      let _nw156 = new _module.C0();
      _nw156.__ctor(_943_v122.f8);
      _963_v139 = _nw156;
      let _964_v140;
      _964_v140 = _dafny.Map.Empty.slice().updateUnsafe(_785_v0,_940_v119);
      let _hi6 = ((_785_v0) ? (_940_v119) : (new BigNumber((_964_v140).length)));
      for (let _965_i10 = new BigNumber((_964_v140).length); _965_i10.isLessThan(_hi6); _965_i10 = _965_i10.plus(_dafny.ONE)) {
        r1 = _785_v0;
        r1 = true;
        (_963_v139).f8 = _963_v139.f8;
        let _966_v141;
        let _nw157 = Array((new BigNumber(25)).toNumber()).fill(false);
        _966_v141 = _nw157;
        let _index150 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_966_v141).length));
        (_966_v141)[_index150] = !(_785_v0);
        let _index151 = _module.__default.safeIndex(new BigNumber(704), new BigNumber((_966_v141).length));
        (_966_v141)[_index151] = _785_v0;
      }
      let _967_v142;
      let _init20 = ((_968_v139, _969_p0) => function (_970_i11) {
        return (_970_i11).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_968_v139.f8,_969_p0)).length));
      })(_963_v139, p0);
      let _nw158 = Array((new BigNumber(19)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw158.length); _i0_20++) {
        _nw158[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _967_v142 = _nw158;
      let _index152 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_967_v142).length));
      (_967_v142)[_index152] = _940_v119;
      let _index153 = _module.__default.safeIndex(new BigNumber(605), new BigNumber((_967_v142).length));
      (_967_v142)[_index153] = _940_v119;
      r0 = _785_v0;
      r1 = _785_v0;
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _971_v0;
      _971_v0 = true;
      let _972_v1;
      _972_v1 = _dafny.MultiSet.fromElements(_971_v0, _971_v0);
      let _973_v2;
      _973_v2 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_module.D9.create_DC21(_972_v1)).dtor_cf37);
      _973_v2 = (_973_v2).update(p0, _972_v1);
      let _974_v3;
      _974_v3 = _module.D7.create_DC18();
      let _975_v4;
      let _nw159 = Array((new BigNumber(4)).toNumber()).fill(_dafny.Set.Empty);
      _975_v4 = _nw159;
      let _976_v5;
      let _nw160 = Array((new BigNumber(19)).toNumber());
      _nw160[(_dafny.ZERO).toNumber()] = _971_v0;
      _nw160[(_dafny.ONE).toNumber()] = _971_v0;
      _nw160[(new BigNumber(2)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(3)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(4)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(5)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(6)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(7)).toNumber()] = false;
      _nw160[(new BigNumber(8)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(9)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(10)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(11)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(12)).toNumber()] = _module.__default.fm30(p0, _971_v0, _971_v0, globalState);
      _nw160[(new BigNumber(13)).toNumber()] = true;
      _nw160[(new BigNumber(14)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(15)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(16)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(17)).toNumber()] = _971_v0;
      _nw160[(new BigNumber(18)).toNumber()] = _971_v0;
      _976_v5 = _nw160;
      let _977_v6;
      let _nw161 = Array((new BigNumber(10)).toNumber());
      _nw161[(_dafny.ZERO).toNumber()] = _976_v5;
      _nw161[(_dafny.ONE).toNumber()] = _976_v5;
      _nw161[(new BigNumber(2)).toNumber()] = _976_v5;
      _nw161[(new BigNumber(3)).toNumber()] = _976_v5;
      _nw161[(new BigNumber(4)).toNumber()] = _976_v5;
      _nw161[(new BigNumber(5)).toNumber()] = _976_v5;
      _nw161[(new BigNumber(6)).toNumber()] = _976_v5;
      _nw161[(new BigNumber(7)).toNumber()] = _976_v5;
      _nw161[(new BigNumber(8)).toNumber()] = _976_v5;
      _nw161[(new BigNumber(9)).toNumber()] = _976_v5;
      _977_v6 = _nw161;
      let _rhs95 = _974_v3;
      let _rhs96 = _975_v4;
      let _rhs97 = _977_v6;
      _974_v3 = _rhs95;
      _975_v4 = _rhs96;
      _977_v6 = _rhs97;
      let _hi7 = p0;
      for (let _978_i0 = _module.__default.safeDivisionInt(p0, new BigNumber(-664)); _978_i0.isLessThan(_hi7); _978_i0 = _978_i0.plus(_dafny.ONE)) {
        let _979_v7;
        _979_v7 = new BigNumber(373);
        let _980_v8;
        _980_v8 = _dafny.Seq.UnicodeFromString("fbkmmu");
        let _981_v9;
        _981_v9 = new _dafny.CodePoint('a'.codePointAt(0));
        _979_v7 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-424)), function (_982_i1) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        }), _dafny.Seq.update(_980_v8, _module.__default.safeIndex(new BigNumber((_980_v8).length), new BigNumber((_980_v8).length)), _981_v9))).length);
        _981_v9 = _981_v9;
        let _983_v10;
        let _nw162 = new _module.C0();
        _nw162.__ctor(_dafny.Seq.Concat(_980_v8, _980_v8));
        _983_v10 = _nw162;
        let _984_v11;
        let _init21 = ((_985_v7) => function (_986_i2) {
          return (_986_i2).plus(_985_v7);
        })(_979_v7);
        let _nw163 = Array((new BigNumber(16)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw163.length); _i0_21++) {
          _nw163[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _984_v11 = _nw163;
        let _index154 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_984_v11).length));
        (_984_v11)[_index154] = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(625)), ((_987_p0) => function (_988_i3) {
          return _987_p0;
        })(p0))).length)).plus(_978_i0);
        let _index155 = _module.__default.safeIndex(new BigNumber(616), new BigNumber((_984_v11).length));
        (_984_v11)[_index155] = ((_971_v0) ? (_979_v7) : (_module.__default.safeDivisionInt(_978_i0, _978_i0)));
      }
      let _989_v12;
      let _nw164 = Array((new BigNumber(14)).toNumber()).fill(_module.D2.Default());
      _989_v12 = _nw164;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_989_v12).length))) {
        let _990_i4 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_990_i4)) && ((_990_i4).isLessThan(new BigNumber((_989_v12).length))))) {
          (_989_v12)[(_990_i4)] = _module.D2.create_DC5(p0, _dafny.Seq.of(p0, p0), _971_v0);
        }
      }
      let _991_v13;
      let _992_v14;
      let _993_v15;
      let _994_v16;
      let _out29;
      let _out30;
      let _out31;
      let _out32;
      let _outcollector9 = (_this).m12(_971_v0, globalState);
      _out29 = _outcollector9[0];
      _out30 = _outcollector9[1];
      _out31 = _outcollector9[2];
      _out32 = _outcollector9[3];
      _991_v13 = _out29;
      _992_v14 = _out30;
      _993_v15 = _out31;
      _994_v16 = _out32;
      let _995_v17;
      _995_v17 = _dafny.Map.Empty.slice().updateUnsafe(_994_v16,false);
      let _hi8 = _module.__default.safeDivisionInt(p0, _993_v15);
      for (let _996_i5 = new BigNumber(((_995_v17).Merge(_995_v17)).length); _996_i5.isLessThan(_hi8); _996_i5 = _996_i5.plus(_dafny.ONE)) {
        let _997_v18;
        _997_v18 = _module.D9.create_DC22((_993_v15).minus(_992_v14));
        let _source11 = _997_v18;
        if (_source11.is_DC22) {
          let _998___mcc_h0 = (_source11).cf38;
          let _999_cf38 = _998___mcc_h0;
          let _1000_v19;
          _1000_v19 = _dafny.Seq.of(_994_v16);
          let _1001_v20;
          _1001_v20 = _dafny.MultiSet.fromElements(_1000_v19, _dafny.Seq.Concat(_dafny.Seq.of(_971_v0), _1000_v19), _1000_v19);
          _1001_v20 = ((_1001_v20).Union(_1001_v20)).Union(_1001_v20);
          _999_cf38 = (p0).plus((_992_v14).plus(_993_v15));
          let _1002_v21;
          let _nw165 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
          _1002_v21 = _nw165;
          let _index156 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_1002_v21).length));
          (_1002_v21)[_index156] = _dafny.Seq.Concat(_module.__default.fm38(_994_v16, _992_v14, _994_v16, globalState), _1000_v19);
          let _index157 = _module.__default.safeIndex(new BigNumber(675), new BigNumber((_1002_v21).length));
          (_1002_v21)[_index157] = _1000_v19;
          let _1003_v22;
          _1003_v22 = _module.D2.create_DC5(_999_cf38, _dafny.Seq.Create(_module.__default.abs(new BigNumber(24)), ((_1004_v14) => function (_1005_i6) {
  return (_dafny.ZERO).minus(_1004_v14);
})(_992_v14)), _994_v16);
          let _1006_v23;
          _1006_v23 = _module.D2.create_DC6(_1003_v22);
          let _1007_v24;
          _1007_v24 = _dafny.Map.Empty.slice().updateUnsafe(_991_v13,p0);
          let _1008_v25;
          _1008_v25 = _dafny.Seq.of(_1007_v24);
          let _1009_v26;
          _1009_v26 = _dafny.MultiSet.fromElements(_995_v17, _dafny.Map.Empty.slice().updateUnsafe(true,_991_v13), (_995_v17).update(false, _971_v0), _995_v17, _995_v17);
          let _1010_v27;
          _1010_v27 = _dafny.Seq.of(_996_i5, new BigNumber((_dafny.Seq.update(_module.__default.fm38(_991_v13, new BigNumber(((_1008_v25)[_module.__default.safeIndex((((_1009_v26).contains(_995_v17)) ? ((_1009_v26).get(_995_v17)) : ((_dafny.ZERO).minus(_module.__default.fm32(_993_v15, globalState)))), new BigNumber((_1008_v25).length))]).length), _994_v16, globalState), _module.__default.safeIndex(_999_cf38, new BigNumber((_module.__default.fm38(_991_v13, new BigNumber(((_1008_v25)[_module.__default.safeIndex((((_1009_v26).contains(_995_v17)) ? ((_1009_v26).get(_995_v17)) : ((_dafny.ZERO).minus(_module.__default.fm32(_993_v15, globalState)))), new BigNumber((_1008_v25).length))]).length), _994_v16, globalState)).length)), _971_v0)).length), _993_v15, _module.__default.fm32(new BigNumber(((_1002_v21)[_module.__default.safeIndex(new BigNumber(675), new BigNumber((_1002_v21).length))]).length), globalState), _999_cf38);
          let _1011_v28;
          _1011_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1006_v23,!_dafny.areEqual(_1010_v27, _1010_v27));
          _971_v0 = (((_1011_v28).contains(_1006_v23)) ? ((_1011_v28).get(_1006_v23)) : (_994_v16));
        } else {
          let _1012___mcc_h1 = (_source11).cf37;
          let _1013_cf37 = _1012___mcc_h1;
          let _1014_v29;
          _1014_v29 = _dafny.Map.Empty.slice().updateUnsafe(_992_v14,_971_v0);
          let _1015_v30;
          _1015_v30 = _dafny.Seq.UnicodeFromString("cmpc");
          let _1016_v31;
          _1016_v31 = _dafny.Set.fromElements(_994_v16);
          let _1017_v32;
          _1017_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1015_v30).length),_1016_v31);
          let _1018_v33;
          _1018_v33 = _module.D12.create_DC27((((_1017_v32).contains(_993_v15)) ? ((_1017_v32).get(_993_v15)) : (_1016_v31)));
          let _1019_v34;
          _1019_v34 = _dafny.Map.Empty.slice().updateUnsafe((_1014_v29).update(_996_i5, _971_v0),_1018_v33);
          let _pat_let_tv28 = _1016_v31;
          _1019_v34 = (_1019_v34).update(_1014_v29, function (_pat_let22_0) {
            return function (_1020_dt__update__tmp_h0) {
              return function (_pat_let23_0) {
                return function (_1021_dt__update_hcf45_h0) {
                  return _module.D12.create_DC27(_1021_dt__update_hcf45_h0);
                }(_pat_let23_0);
              }(_pat_let_tv28);
            }(_pat_let22_0);
          }(_1018_v33));
          let _1022_v35;
          _1022_v35 = _dafny.Set.fromElements(_976_v5);
          let _1023_v36;
          _1023_v36 = _1022_v35;
          let _1024_v37;
          let _nw166 = Array((new BigNumber(22)).toNumber());
          _nw166[(_dafny.ZERO).toNumber()] = _1023_v36;
          _nw166[(_dafny.ONE).toNumber()] = _1022_v35;
          _nw166[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_976_v5);
          _nw166[(new BigNumber(3)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(4)).toNumber()] = _1022_v35;
          _nw166[(new BigNumber(5)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(6)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(7)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(8)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(9)).toNumber()] = _1022_v35;
          _nw166[(new BigNumber(10)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(11)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(12)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(13)).toNumber()] = ((!(_971_v0)) ? (_1023_v36) : (_1022_v35));
          _nw166[(new BigNumber(14)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(15)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(16)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(17)).toNumber()] = _1022_v35;
          _nw166[(new BigNumber(18)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(19)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(20)).toNumber()] = _1023_v36;
          _nw166[(new BigNumber(21)).toNumber()] = _1023_v36;
          _1024_v37 = _nw166;
          let _index158 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_1024_v37).length));
          (_1024_v37)[_index158] = _1023_v36;
          let _index159 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_1024_v37).length));
          (_1024_v37)[_index159] = _1023_v36;
          let _1025_v38;
          _1025_v38 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1026_v39;
          _1026_v39 = _dafny.Seq.of(_994_v16);
          let _1027_v40;
          _1027_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1025_v38,_993_v15);
          let _1028_v42;
          _1028_v42 = _dafny.Set.fromElements(_1025_v38, _1025_v38, _1025_v38);
          let _1029_v43;
          _1029_v43 = _dafny.Seq.of(_1027_v40, _1027_v40);
          let _1030_v44;
          let _nw167 = Array((new BigNumber(25)).toNumber());
          _nw167[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1025_v38,new BigNumber((_1026_v39).length));
          _nw167[(_dafny.ONE).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(2)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(3)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(4)).toNumber()] = function () {
            let _coll45 = new _dafny.Map();
            for (const _compr_45 of (_1028_v42).Elements) {
              let _1031_v41 = _compr_45;
              if ((_1028_v42).contains(_1031_v41)) {
                _coll45.push([_1031_v41,new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-584)), ((_1032_v38) => function (_1033_i7) {
                  return _1032_v38;
                })(_1025_v38))).length)]);
              }
            }
            return _coll45;
          }();
          _nw167[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm29(_971_v0, globalState),_993_v15);
          _nw167[(new BigNumber(6)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1025_v38,_module.__default.fm32(_993_v15, globalState));
          _nw167[(new BigNumber(7)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(8)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(9)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(10)).toNumber()] = (_1029_v43)[_module.__default.safeIndex(p0, new BigNumber((_1029_v43).length))];
          _nw167[(new BigNumber(11)).toNumber()] = (_1029_v43)[_module.__default.safeIndex(_992_v14, new BigNumber((_1029_v43).length))];
          _nw167[(new BigNumber(12)).toNumber()] = (_1027_v40);
          _nw167[(new BigNumber(13)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(14)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(15)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(16)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1025_v38,_996_i5);
          _nw167[(new BigNumber(18)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(19)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(20)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(21)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(22)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(23)).toNumber()] = _1027_v40;
          _nw167[(new BigNumber(24)).toNumber()] = _1027_v40;
          _1030_v44 = _nw167;
          let _1034_v45;
          _1034_v45 = _dafny.Seq.of(_992_v14);
          let _1035_v46;
          _1035_v46 = _module.D0.create_DC1(p0);
          let _1036_v47;
          let _nw168 = new _module.C1();
          _nw168.__ctor((_996_i5).plus(_996_i5), _1030_v44, _1034_v45, _1035_v46);
          _1036_v47 = _nw168;
          let _index160 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_976_v5).length));
          (_976_v5)[_index160] = _991_v13;
          let _index161 = _module.__default.safeIndex(new BigNumber(582), new BigNumber((_976_v5).length));
          (_976_v5)[_index161] = (_module.__default.fm39(_991_v13, globalState)).contains(_993_v15);
        }
        _993_v15 = _996_i5;
        let _1037_v48;
        _1037_v48 = new _dafny.CodePoint('t'.codePointAt(0));
        let _1038_v50;
        _1038_v50 = _module.D7.create_DC19(new BigNumber((function () {
  let _coll46 = new _dafny.Set();
  for (const _compr_46 of _dafny.IntegerRange(new BigNumber(195), new BigNumber(383))) {
    let _1039_v49 = _compr_46;
    if (((new BigNumber(195)).isLessThanOrEqualTo(_1039_v49)) && ((_1039_v49).isLessThan(new BigNumber(383)))) {
      _coll46.add(_module.__default.safeModuloInt(_1039_v49, _996_i5));
    }
  }
  return _coll46;
}()).length));
        let _source12 = _module.__default.fm40(_991_v13, _1037_v48, _994_v16, _1038_v50, globalState);
        if (_source12.is_DC28) {
          let _1040___mcc_h2 = (_source12).cf46;
          let _1041___mcc_h3 = (_source12).cf47;
          let _1042_cf47 = _1041___mcc_h3;
          let _1043_cf46 = _1040___mcc_h2;
          _993_v15 = (_dafny.ZERO).minus(_996_i5);
          let _1044_v51;
          let _nw169 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
          _1044_v51 = _nw169;
          let _index162 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1044_v51).length));
          (_1044_v51)[_index162] = p0;
          let _1045_v52;
          _1045_v52 = _dafny.Map.Empty.slice().updateUnsafe(_996_i5,_1044_v51);
          let _index163 = _module.__default.safeIndex(new BigNumber(625), new BigNumber((_1044_v51).length));
          (_1044_v51)[_index163] = new BigNumber((_dafny.Seq.of(new BigNumber((_1045_v52).length))).length);
          let _1046_v53;
          _1046_v53 = _dafny.Seq.UnicodeFromString("ruavycy");
          let _1047_v54;
          _1047_v54 = _module.D13.create_DC31(_1043_cf46, new BigNumber((_1046_v53).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-504)), ((_1048_v48) => function (_1049_i8) {
  return _1048_v48;
})(_1037_v48))).length));
          let _1050_v55;
          let _nw170 = Array((_dafny.ONE).toNumber());
          _nw170[(_dafny.ZERO).toNumber()] = function (_pat_let24_0) {
            return function (_1051_dt__update__tmp_h4) {
              return function (_pat_let25_0) {
                return function (_1052_dt__update_hcf56_h0) {
                  return _module.D13.create_DC31((_1051_dt__update__tmp_h4).dtor_cf54, (_1051_dt__update__tmp_h4).dtor_cf55, _1052_dt__update_hcf56_h0);
                }(_pat_let25_0);
              }(_996_i5);
            }(_pat_let24_0);
          }(_1047_v54);
          _1050_v55 = _nw170;
          let _1053_v56;
          _1053_v56 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm30(_996_i5, _991_v13, true, globalState),_1050_v55);
          let _1054_v57;
          let _nw171 = Array((new BigNumber(3)).toNumber());
          _nw171[(_dafny.ZERO).toNumber()] = _1050_v55;
          _nw171[(_dafny.ONE).toNumber()] = (((_1053_v56).contains(!(_1043_cf46))) ? ((_1053_v56).get(!(_1043_cf46))) : (_1050_v55));
          _nw171[(new BigNumber(2)).toNumber()] = _1050_v55;
          _1054_v57 = _nw171;
          let _index164 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1054_v57).length));
          (_1054_v57)[_index164] = _1050_v55;
          let _index165 = _module.__default.safeIndex(new BigNumber(332), new BigNumber((_1054_v57).length));
          (_1054_v57)[_index165] = _1050_v55;
          let _1055_v58;
          _1055_v58 = _dafny.Seq.of(_1042_cf47);
          let _1056_v59;
          _1056_v59 = _dafny.Map.Empty.slice().updateUnsafe(_971_v0,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(977)), ((_1057_v15) => function (_1058_i9) {
            return (_dafny.ZERO).minus(_1057_v15);
          })(_993_v15))).length)));
          let _rhs98 = _994_v16;
          let _rhs99 = _1046_v53;
          let _rhs100 = (((new BigNumber(226)).isEqualTo(_992_v14)) ? (_dafny.Seq.of((_1047_v54).dtor_cf55)) : (_dafny.Seq.of(_996_i5, _1042_cf47, _module.__default.fm32(_996_i5, globalState), new BigNumber((_1056_v59).length))));
          let _rhs101 = _994_v16;
          _1043_cf46 = _rhs98;
          _1046_v53 = _rhs99;
          _1055_v58 = _rhs100;
          _991_v13 = _rhs101;
        } else if (_source12.is_DC29) {
          let _1059___mcc_h4 = (_source12).cf48;
          let _1060___mcc_h5 = (_source12).cf49;
          let _1061___mcc_h6 = (_source12).cf50;
          let _1062___mcc_h7 = (_source12).cf51;
          let _1063___mcc_h8 = (_source12).cf52;
          let _1064_cf52 = _1063___mcc_h8;
          let _1065_cf51 = _1062___mcc_h7;
          let _1066_cf50 = _1061___mcc_h6;
          let _1067_cf49 = _1060___mcc_h5;
          let _1068_cf48 = _1059___mcc_h4;
          let _1069_v60;
          _1069_v60 = _dafny.Seq.of(_1068_cf48, _1064_cf52);
          let _index166 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_976_v5).length));
          (_976_v5)[_index166] = (_1069_v60)[_module.__default.safeIndex(_996_i5, new BigNumber((_1069_v60).length))];
          let _index167 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_976_v5).length));
          (_976_v5)[_index167] = (_994_v16) || (!(true) || (_991_v13));
          let _1070_v61;
          _1070_v61 = _dafny.Seq.UnicodeFromString("wnk");
          let _1071_v62;
          _1071_v62 = _module.D6.create_DC14(_1070_v61);
          _1071_v62 = _1071_v62;
          let _1072_v63;
          _1072_v63 = _dafny.Map.Empty.slice().updateUnsafe(_1037_v48,_1065_cf51);
          let _1073_v65;
          _1073_v65 = _dafny.Seq.of((_1072_v63).update(_1037_v48, p0), _dafny.Map.Empty.slice().updateUnsafe(_1037_v48,_992_v14));
          let _1074_v66;
          let _nw172 = Array((new BigNumber(18)).toNumber());
          _nw172[(_dafny.ZERO).toNumber()] = _1072_v63;
          _nw172[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1037_v48,(_dafny.ZERO).minus(new BigNumber((_972_v1).cardinality())));
          _nw172[(new BigNumber(2)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(3)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(4)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(5)).toNumber()] = function () {
            let _coll47 = new _dafny.Map();
            for (const _compr_47 of ((_1072_v63).update(_1037_v48, _1065_cf51)).Keys.Elements) {
              let _1075_v64 = _compr_47;
              if (((_1072_v63).update(_1037_v48, _1065_cf51)).contains(_1075_v64)) {
                _coll47.push([_1075_v64,new BigNumber(813)]);
              }
            }
            return _coll47;
          }();
          _nw172[(new BigNumber(6)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(7)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(8)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(9)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(10)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(11)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(12)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(13)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(14)).toNumber()] = (_1072_v63).update(new _dafny.CodePoint('k'.codePointAt(0)), _996_i5);
          _nw172[(new BigNumber(15)).toNumber()] = _1072_v63;
          _nw172[(new BigNumber(16)).toNumber()] = (_1073_v65)[_module.__default.safeIndex(_993_v15, new BigNumber((_1073_v65).length))];
          _nw172[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1037_v48,(((_972_v1).contains(_1066_cf50)) ? ((_972_v1).get(_1066_cf50)) : (_992_v14)));
          _1074_v66 = _nw172;
          let _1076_v67;
          _1076_v67 = _dafny.Seq.of(p0);
          let _1077_v68;
          let _nw173 = new _module.C1();
          _nw173.__ctor((_dafny.ZERO).minus(_1065_cf51), ((_994_v16) ? (_1074_v66) : (_1074_v66)), _dafny.Seq.Concat(_dafny.Seq.update(_1076_v67, _module.__default.safeIndex(new BigNumber(551), new BigNumber((_1076_v67).length)), _996_i5), _1076_v67), _module.D0.create_DC1(_module.__default.fm32(new BigNumber(-714), globalState)));
          _1077_v68 = _nw173;
          let _1078_v69;
          _1078_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_972_v1).Difference(_972_v1)).cardinality()),new BigNumber(111));
          _1078_v69 = (_1078_v69).update(p0, (_dafny.ZERO).minus((_996_i5).minus(_1065_cf51)));
        } else {
          let _1079___mcc_h9 = (_source12).cf45;
          let _1080_cf45 = _1079___mcc_h9;
          let _1081_v70;
          let _nw174 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
          _1081_v70 = _nw174;
          let _1082_v71;
          let _nw175 = new _module.C2();
          _nw175.__ctor();
          _1082_v71 = _nw175;
          let _1083_v72;
          _1083_v72 = _dafny.Map.Empty.slice().updateUnsafe(_1082_v71,true);
          let _1084_v73;
          _1084_v73 = _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(474)), ((_1085_v48) => function (_1086_i10) {
            return _1085_v48;
          })(_1037_v48))).length), new BigNumber((_1083_v72).length));
          let _1087_v74;
          _1087_v74 = _module.D0.create_DC1(_993_v15);
          let _1088_v75;
          let _nw176 = new _module.C1();
          _nw176.__ctor(_993_v15, _1081_v70, _1084_v73, _1087_v74);
          _1088_v75 = _nw176;
          let _1089_v76;
          _1089_v76 = _dafny.Seq.UnicodeFromString("dkjicxbn");
          let _1090_v77;
          let _nw177 = new _module.C0();
          _nw177.__ctor(_dafny.Seq.Concat(_1089_v76, _1089_v76));
          _1090_v77 = _nw177;
          let _1091_v78;
          let _init22 = ((_1092_v75) => function (_1093_i11) {
            return _module.__default.safeDivisionInt(_1093_i11, new BigNumber((_dafny.Seq.of(_dafny.Seq.of(_1092_v75.f9))).length));
          })(_1088_v75);
          let _nw178 = Array((new BigNumber(18)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw178.length); _i0_22++) {
            _nw178[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _1091_v78 = _nw178;
          let _index168 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1091_v78).length));
          (_1091_v78)[_index168] = new BigNumber((_1090_v77.f8).length);
          let _index169 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_1091_v78).length));
          (_1091_v78)[_index169] = _996_i5;
          let _1094_v79;
          _1094_v79 = _dafny.Set.fromElements(_996_i5);
          _1094_v79 = _1094_v79;
        }
        _1037_v48 = _1037_v48;
      }
      let _1095_v80;
      _1095_v80 = _dafny.Seq.UnicodeFromString("ohddunlw");
      let _1096_v81;
      _1096_v81 = _dafny.Set.fromElements(_994_v16, _971_v0);
      let _1097_v82;
      _1097_v82 = _dafny.Seq.of(_993_v15, _993_v15, new BigNumber((_1096_v81).length));
      let _1098_v83;
      _1098_v83 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_1095_v80, _1095_v80),_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_1097_v82, _1097_v82)));
      r0 = _1098_v83;
      return r0;
    }
    m11(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.Set.Empty;
      let _1099_v0;
      _1099_v0 = true;
      let _index170 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length));
      (p1)[_index170] = _1099_v0;
      let _1100_v1;
      _1100_v1 = _dafny.Seq.of(p3);
      let _1101_v2;
      _1101_v2 = _dafny.Seq.UnicodeFromString("h");
      let _1102_v3;
      _1102_v3 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(p3));
      let _1103_v4;
      _1103_v4 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1101_v2).length),_1102_v3);
      let _1104_v5;
      let _nw179 = Array((new BigNumber(3)).toNumber());
      _nw179[(_dafny.ZERO).toNumber()] = p3;
      _nw179[(_dafny.ONE).toNumber()] = p3;
      _nw179[(new BigNumber(2)).toNumber()] = p3;
      _1104_v5 = _nw179;
      let _1105_v6;
      _1105_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1104_v5,p3);
      let _index171 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length));
      let _rhs102 = true;
      let _rhs103 = !(((((_1103_v4).contains(new BigNumber(-931))) ? ((_1103_v4).get(new BigNumber(-931))) : (_1102_v3))).IsDisjointFrom(_1102_v3));
      let _rhs104 = _dafny.Seq.of(new BigNumber((_1105_v6).length));
      let _rhs105 = _1099_v0;
      let _lhs66 = p1;
      let _lhs67 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length));
      _lhs66[_lhs67] = _rhs102;
      _1099_v0 = _rhs103;
      _1100_v1 = _rhs104;
      r1 = _rhs105;
      if (_1099_v0) {
        _1099_v0 = _1099_v0;
        _1101_v2 = _1101_v2;
        let _index172 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length));
        (p1)[_index172] = (p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))];
        let _1106_v7;
        _1106_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1099_v0,((_1099_v0) ? (_dafny.Seq.of(p0, p0)) : (_1101_v2)));
        let _1107_v8;
        _1107_v8 = new BigNumber(-951);
        let _1108_v9;
        let _init23 = ((_1109_p0) => function (_1110_i0) {
          return _1109_p0;
        })(p0);
        let _nw180 = Array((new BigNumber(11)).toNumber());
        for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw180.length); _i0_23++) {
          _nw180[_i0_23] = _init23(new BigNumber(_i0_23));
        }
        _1108_v9 = _nw180;
        let _index173 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1108_v9).length));
        (_1108_v9)[_index173] = new _dafny.CodePoint('v'.codePointAt(0));
        let _index174 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1108_v9).length));
        let _rhs106 = (_1106_v7).Merge(_1106_v7);
        let _rhs107 = _module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(803)), ((_1111_p0) => function (_1112_i1) {
          return _1111_p0;
        })(p0))).length), (new BigNumber(491)).minus(_1107_v8));
        let _rhs108 = p0;
        let _rhs109 = _1107_v8;
        let _lhs68 = _1108_v9;
        let _lhs69 = _module.__default.safeIndex(new BigNumber(623), new BigNumber((_1108_v9).length));
        _1106_v7 = _rhs106;
        _1107_v8 = _rhs107;
        _lhs68[_lhs69] = _rhs108;
        _1107_v8 = _rhs109;
        r1 = (_1107_v8).isEqualTo(_1107_v8);
      } else {
        let _1113_v10;
        _1113_v10 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1114_v11;
        _1114_v11 = new BigNumber(388);
        let _index175 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1104_v5).length));
        (_1104_v5)[_index175] = (p3).multipliedBy(_module.__default.fm32(_1114_v11, globalState));
        let _1115_v12;
        _1115_v12 = _module.D7.create_DC18();
        let _1116_v13;
        _1116_v13 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))],_1115_v12);
        let _1117_v14;
        _1117_v14 = _dafny.Seq.of(_1116_v13, _1116_v13);
        let _1118_v15;
        _1118_v15 = _module.D13.create_DC31((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))], _1114_v11, p3);
        let _1119_v16;
        _1119_v16 = _dafny.MultiSet.fromElements(_module.__default.fm44(true, globalState), _1114_v11, _1114_v11, _module.__default.fm44(_1099_v0, globalState));
        let _1120_v17;
        _1120_v17 = _dafny.MultiSet.fromElements(_1119_v16, _1119_v16, (_1119_v16).Intersect(_dafny.MultiSet.FromArray(_1100_v1)));
        let _1121_v18;
        _1121_v18 = _dafny.Seq.of(_1099_v0, _1099_v0);
        let _index176 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1104_v5).length));
        let _rhs110 = _module.__default.fm51(_1099_v0, p3, _module.__default.safeDivisionInt(new BigNumber((_1101_v2).length), (_1100_v1)[_module.__default.safeIndex(new BigNumber((_1117_v14).length), new BigNumber((_1100_v1).length))]), globalState);
        let _rhs111 = (((_1099_v0) ? (_1118_v15) : (_module.D13.create_DC31(_1099_v0, p3, p3)))).dtor_cf54;
        let _rhs112 = (((_1120_v17).contains(((_1119_v16).update(_1114_v11, _module.__default.abs(new BigNumber((_1121_v18).length)))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(816))))) ? ((_1120_v17).get(((_1119_v16).update(_1114_v11, _module.__default.abs(new BigNumber((_1121_v18).length)))).Intersect(_dafny.MultiSet.fromElements(new BigNumber(816))))) : (_1114_v11));
        let _rhs113 = !((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))]);
        let _rhs114 = new BigNumber((_dafny.Seq.UnicodeFromString("ghsv")).length);
        let _lhs70 = _1104_v5;
        let _lhs71 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1104_v5).length));
        _1113_v10 = _rhs110;
        r2 = _rhs111;
        _1114_v11 = _rhs112;
        r2 = _rhs113;
        _lhs70[_lhs71] = _rhs114;
        if (((_dafny.ZERO).minus((_1104_v5)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1104_v5).length))])).isLessThanOrEqualTo((_dafny.ZERO).minus(new BigNumber(-737)))) {
          let _1122_v19;
          _1122_v19 = _dafny.Set.fromElements(_1113_v10, p0, p0);
          let _1123_v20;
          _1123_v20 = _dafny.Set.fromElements(_1121_v18, _dafny.Seq.of(_1099_v0));
          let _1124_v21;
          _1124_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1122_v19).length),(_1123_v20).IsSubsetOf(_module.__default.fm55(p3, globalState)));
          let _1125_v22;
          _1125_v22 = _dafny.Set.fromElements(new BigNumber(830));
          _1124_v21 = (_1124_v21).update(new BigNumber(((_dafny.Set.fromElements((_1104_v5)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1104_v5).length))], (_dafny.ZERO).minus(_1114_v11), (_1104_v5)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1104_v5).length))])).Difference(_1125_v22)).length), (_module.D11.create_DC25(_1113_v10, _1114_v11, false)).dtor_cf43);
          let _1126_v23;
          let _nw181 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _1126_v23 = _nw181;
          _1126_v23 = _1126_v23;
          let _1127_v24;
          let _nw182 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Seq.of());
          _1127_v24 = _nw182;
          let _index177 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_1127_v24).length));
          (_1127_v24)[_index177] = _1100_v1;
          let _index178 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_1127_v24).length));
          (_1127_v24)[_index178] = _dafny.Seq.Concat(_1100_v1, _1100_v1);
          let _1128_v25;
          _1128_v25 = _dafny.MultiSet.fromElements(true);
          _1128_v25 = (_1128_v25).update((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))], _module.__default.abs(_1114_v11));
          let _1129_v26;
          _1129_v26 = _dafny.Seq.of(p1, p1, p1, (((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))]) ? (p1) : (p1)), p1);
          _1129_v26 = _1129_v26;
        } else {
          let _index179 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length));
          (p1)[_index179] = _dafny.Seq.IsProperPrefixOf(_1101_v2, _1101_v2);
          let _1130_v27;
          _1130_v27 = _dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))],(p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))]);
          _1130_v27 = (_1130_v27).update(true, !_dafny.areEqual(_dafny.Seq.of(_1099_v0), _1121_v18));
          r1 = _1099_v0;
          let _1131_v28;
          _1131_v28 = _dafny.Set.fromElements(_1099_v0);
          let _1132_v29;
          _1132_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1099_v0,p3);
          let _index180 = _module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length));
          (p1)[_index180] = (_1131_v28).IsDisjointFrom(_module.__default.fm53(_1132_v29, globalState));
          let _1133_v30;
          _1133_v30 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1099_v0);
          let _rhs115 = ((new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1114_v11,_1099_v0)).length)).multipliedBy(new BigNumber((_1133_v30).length))).isLessThanOrEqualTo((_dafny.ZERO).minus(_1114_v11));
          let _rhs116 = _1101_v2;
          let _rhs117 = (_module.__default.fm44((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))], globalState)).multipliedBy(p3);
          r1 = _rhs115;
          _1101_v2 = _rhs116;
          _1114_v11 = _rhs117;
        }
        let _index181 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1104_v5).length));
        (_1104_v5)[_index181] = (_1104_v5)[_module.__default.safeIndex(new BigNumber(98), new BigNumber((_1104_v5).length))];
        let _1134_v31;
        _1134_v31 = _module.D12.create_DC29(false, p3, (p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))], _1114_v11, (p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))]);
        let _index182 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1104_v5).length));
        (_1104_v5)[_index182] = (_1134_v31).dtor_cf49;
        let _1135_v32;
        _1135_v32 = _module.D12.create_DC28((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))], (_dafny.ZERO).minus(_1114_v11));
        let _1136_v33;
        _1136_v33 = _dafny.Set.fromElements((_dafny.ZERO).minus(p3), _1114_v11);
        let _pat_let_tv29 = _1136_v33;
        let _1137_v34;
        _1137_v34 = _dafny.MultiSet.fromElements(_module.D12.create_DC28((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))], _1114_v11), _1135_v32, _1135_v32, function (_pat_let26_0) {
          return function (_1138_dt__update__tmp_h0) {
            return function (_pat_let27_0) {
              return function (_1139_dt__update_hcf47_h0) {
                return _module.D12.create_DC28((_1138_dt__update__tmp_h0).dtor_cf46, _1139_dt__update_hcf47_h0);
              }(_pat_let27_0);
            }(new BigNumber((_pat_let_tv29).length));
          }(_pat_let26_0);
        }(_1135_v32), _1135_v32);
        let _index183 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1104_v5).length));
        (_1104_v5)[_index183] = new BigNumber((_1137_v34).cardinality());
      }
      let _index184 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_1104_v5).length));
      (_1104_v5)[_index184] = new BigNumber((_1101_v2).length);
      let _index185 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_1104_v5).length));
      (_1104_v5)[_index185] = p3;
      _1099_v0 = !(!(!(_1099_v0)) || ((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))])) || (_1099_v0);
      let _1140_v35;
      let _1141_v36;
      let _1142_v37;
      let _1143_v38;
      let _out33;
      let _out34;
      let _out35;
      let _out36;
      let _outcollector10 = (_this).m12((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))], globalState);
      _out33 = _outcollector10[0];
      _out34 = _outcollector10[1];
      _out35 = _outcollector10[2];
      _out36 = _outcollector10[3];
      _1140_v35 = _out33;
      _1141_v36 = _out34;
      _1142_v37 = _out35;
      _1143_v38 = _out36;
      let _index186 = _module.__default.safeIndex(new BigNumber(913), new BigNumber((_1104_v5).length));
      (_1104_v5)[_index186] = _1141_v36;
      let _1144_v39;
      _1144_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1143_v38,(p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))]);
      let _1145_v40;
      _1145_v40 = _module.D16.create_DC42(new BigNumber(173), _1141_v36, _1141_v36);
      let _1146_v41;
      _1146_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm30(_1142_v37, (p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))], _1140_v35, globalState),(_1104_v5)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_1104_v5).length))]);
      let _pat_let_tv30 = _1146_v41;
      r0 = (_dafny.Map.Empty.slice().updateUnsafe((_1104_v5)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_1104_v5).length))],_1144_v39)).update((function (_pat_let28_0) {
        return function (_1147_dt__update__tmp_h1) {
          return function (_pat_let29_0) {
            return function (_1148_dt__update_hcf76_h0) {
              return _module.D16.create_DC42(_1148_dt__update_hcf76_h0, (_1147_dt__update__tmp_h1).dtor_cf77, (_1147_dt__update__tmp_h1).dtor_cf78);
            }(_pat_let29_0);
          }(new BigNumber((_pat_let_tv30).length));
        }(_pat_let28_0);
      }(_1145_v40)).dtor_cf78, _1144_v39);
      r1 = (p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))];
      let _1149_v42;
      _1149_v42 = _module.D0.create_DC1(_1141_v36);
      let _1150_v43;
      _1150_v43 = _dafny.Seq.of(_1140_v35, (p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))]);
      let _pat_let_tv31 = _1104_v5;
      let _pat_let_tv32 = _1104_v5;
      r2 = (_this).fm5(new BigNumber((_dafny.MultiSet.fromElements((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))])).cardinality()), _1140_v35, (_dafny.Map.Empty.slice().updateUnsafe(p0,function (_pat_let30_0) {
        return function (_1151_dt__update__tmp_h2) {
          return function (_pat_let31_0) {
            return function (_1152_dt__update_hcf2_h0) {
              return _module.D0.create_DC1(_1152_dt__update_hcf2_h0);
            }(_pat_let31_0);
          }((_pat_let_tv32)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_pat_let_tv31).length))]);
        }(_pat_let30_0);
      }(_1149_v42))).update(p0, _module.D0.create_DC1(new BigNumber((_1150_v43).length))), true, globalState);
      let _1153_v44;
      _1153_v44 = _dafny.Set.fromElements(_1146_v41, _1146_v41, _1146_v41, _1146_v41);
      r3 = (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((p1)[_module.__default.safeIndex(new BigNumber(659), new BigNumber((p1).length))],(_1104_v5)[_module.__default.safeIndex(new BigNumber(913), new BigNumber((_1104_v5).length))]))).Intersect(_1153_v44);
      return [r0, r1, r2, r3];
    }
    m12(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = false;
      let _1154_v0;
      _1154_v0 = new BigNumber(886);
      let _1155_v1;
      _1155_v1 = _dafny.Seq.of(_1154_v0);
      let _rhs118 = false;
      let _rhs119 = _1154_v0;
      let _rhs120 = _dafny.Seq.update(_dafny.Seq.of(_1154_v0), _module.__default.safeIndex(_1154_v0, new BigNumber((_dafny.Seq.of(_1154_v0)).length)), new BigNumber(667));
      r3 = _rhs118;
      r1 = _rhs119;
      _1155_v1 = _rhs120;
      let _hi9 = _1154_v0;
      for (let _1156_i0 = _1154_v0; _1156_i0.isLessThan(_hi9); _1156_i0 = _1156_i0.plus(_dafny.ONE)) {
        let _1157_v2;
        _1157_v2 = _dafny.Seq.UnicodeFromString("sxccix");
        if (!(_1154_v0).isEqualTo(new BigNumber((_1157_v2).length))) {
          let _rhs121 = !(p0);
          let _rhs122 = p0;
          r3 = _rhs121;
          r3 = _rhs122;
          let _1158_v3;
          let _nw183 = Array((new BigNumber(8)).toNumber()).fill([]);
          _1158_v3 = _nw183;
          let _1159_v4;
          _1159_v4 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),p0);
          let _1160_v5;
          _1160_v5 = _dafny.Set.fromElements(new BigNumber((_1159_v4).length), new BigNumber(415));
          let _1161_v7;
          _1161_v7 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1156_i0);
          let _1162_v8;
          _1162_v8 = _module.D14.create_DC36(function () {
  let _coll48 = new _dafny.Map();
  for (const _compr_48 of _dafny.IntegerRange(new BigNumber(381), new BigNumber(529))) {
    let _1163_v6 = _compr_48;
    if (((new BigNumber(381)).isLessThanOrEqualTo(_1163_v6)) && ((_1163_v6).isLessThan(new BigNumber(529)))) {
      _coll48.push([(_1163_v6).minus(new BigNumber((_dafny.Seq.of(false)).length)),_1156_i0]);
    }
  }
  return _coll48;
}(), _1161_v7);
          let _1164_v9;
          _1164_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1157_v2).length),_1156_i0);
          let _1165_v10;
          let _nw184 = Array((new BigNumber(6)).toNumber());
          _nw184[(_dafny.ZERO).toNumber()] = _1162_v8;
          _nw184[(_dafny.ONE).toNumber()] = _module.D14.create_DC36(_1164_v9, _1161_v7);
          _nw184[(new BigNumber(2)).toNumber()] = _1162_v8;
          _nw184[(new BigNumber(3)).toNumber()] = _1162_v8;
          _nw184[(new BigNumber(4)).toNumber()] = _1162_v8;
          _nw184[(new BigNumber(5)).toNumber()] = _1162_v8;
          _1165_v10 = _nw184;
          let _1166_v11;
          _1166_v11 = _dafny.Seq.of(_1165_v10, _1165_v10);
          let _1167_v14;
          let _nw185 = Array((new BigNumber(25)).toNumber());
          _nw185[(_dafny.ZERO).toNumber()] = _1160_v5;
          _nw185[(_dafny.ONE).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(2)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(3)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(4)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(5)).toNumber()] = _dafny.Set.fromElements(new BigNumber((_1166_v11).length), _1156_i0);
          _nw185[(new BigNumber(6)).toNumber()] = _dafny.Set.fromElements((_dafny.ZERO).minus(_1156_i0), _1156_i0);
          _nw185[(new BigNumber(7)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(8)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(9)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(10)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements(_module.__default.fm44(p0, globalState));
          _nw185[(new BigNumber(12)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(13)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(14)).toNumber()] = function () {
            let _coll49 = new _dafny.Set();
            for (const _compr_49 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber(289))) {
              let _1168_v12 = _compr_49;
              if (((_dafny.ZERO).isLessThanOrEqualTo(_1168_v12)) && ((_1168_v12).isLessThan(new BigNumber(289)))) {
                _coll49.add((_1168_v12).plus(new BigNumber((_1155_v1).length)));
              }
            }
            return _coll49;
          }();
          _nw185[(new BigNumber(15)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(16)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(17)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(18)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(19)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(20)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(21)).toNumber()] = _dafny.Set.fromElements(_1154_v0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-591)), function (_1169_i1) {
            return new _dafny.CodePoint('h'.codePointAt(0));
          })).length), _1156_i0, _1156_i0);
          _nw185[(new BigNumber(22)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(23)).toNumber()] = _1160_v5;
          _nw185[(new BigNumber(24)).toNumber()] = function () {
            let _coll50 = new _dafny.Set();
            for (const _compr_50 of _dafny.IntegerRange(new BigNumber(578), new BigNumber(846))) {
              let _1170_v13 = _compr_50;
              if (((new BigNumber(578)).isLessThanOrEqualTo(_1170_v13)) && ((_1170_v13).isLessThan(new BigNumber(846)))) {
                _coll50.add((_1170_v13).minus(new BigNumber((_1159_v4).length)));
              }
            }
            return _coll50;
          }();
          _1167_v14 = _nw185;
          let _index187 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_1158_v3).length));
          (_1158_v3)[_index187] = _1167_v14;
          let _index188 = _module.__default.safeIndex(new BigNumber(733), new BigNumber((_1158_v3).length));
          (_1158_v3)[_index188] = _1167_v14;
          let _init24 = ((_1171_v5) => function (_1172_i2) {
            return _1171_v5;
          })(_1160_v5);
          let _nw186 = Array((new BigNumber(17)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw186.length); _i0_24++) {
            _nw186[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _1167_v14 = _nw186;
          let _1173_v15;
          _1173_v15 = new _dafny.CodePoint('f'.codePointAt(0));
          let _1174_v16;
          let _nw187 = new _module.C0();
          _nw187.__ctor(_dafny.Seq.Concat(_1157_v2, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(790)), function (_1175_i3) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          }), _module.__default.safeIndex(new BigNumber((_1157_v2).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(790)), function (_1176_i3) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          })).length)), _1173_v15)));
          _1174_v16 = _nw187;
          let _1177_v17;
          _1177_v17 = _module.D0.create_DC1(new BigNumber((_1161_v7).length));
          let _1178_v18;
          _1178_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1173_v15,_1177_v17);
          _1159_v4 = (_1159_v4).update((_this).fm5(_1156_i0, p0, _1178_v18, p0, globalState), !_dafny.areEqual(_1157_v2, _1157_v2));
        } else {
          let _1179_v19;
          _1179_v19 = _module.D7.create_DC18();
          _1179_v19 = _1179_v19;
          r3 = p0;
          let _1180_v20;
          _1180_v20 = _dafny.MultiSet.fromElements(true, p0);
          let _1181_v21;
          _1181_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1156_i0,p0);
          let _1182_v22;
          _1182_v22 = _dafny.Seq.of(p0);
          let _1183_v23;
          _1183_v23 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_1182_v22),p0);
          let _1184_v24;
          _1184_v24 = _dafny.Seq.of(p0, _module.__default.fm30(_1156_i0, (((_1181_v21).contains(new BigNumber((_1183_v23).length))) ? ((_1181_v21).get(new BigNumber((_1183_v23).length))) : (p0)), true, globalState), p0, p0);
          let _1185_v25;
          _1185_v25 = _module.D4.create_DC11(_1154_v0, p0, p0, p0, p0);
          r3 = !(((p0) ? ((_1180_v20).IsDisjointFrom(_dafny.MultiSet.FromArray(_1182_v22))) : ((_1185_v25).dtor_cf26)));
          let _1186_v26;
          _1186_v26 = new _dafny.CodePoint('k'.codePointAt(0));
          _1186_v26 = ((p0) ? (_1186_v26) : ((_1157_v2)[_module.__default.safeIndex(_1154_v0, new BigNumber((_1157_v2).length))]));
          let _1187_v27;
          let _nw188 = Array((new BigNumber(2)).toNumber());
          _nw188[(_dafny.ZERO).toNumber()] = _1154_v0;
          _nw188[(_dafny.ONE).toNumber()] = _1154_v0;
          _1187_v27 = _nw188;
          let _1188_v28;
          _1188_v28 = _dafny.Set.fromElements(_1187_v27);
          let _1189_v29;
          _1189_v29 = _dafny.MultiSet.fromElements(_1156_i0, new BigNumber((_1188_v28).length));
          let _1190_v30;
          _1190_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1189_v29,p0);
          r3 = ((((_1190_v30).contains(_1189_v29)) ? ((_1190_v30).get(_1189_v29)) : (true))) === (p0);
        }
        r0 = p0;
        let _1191_v31;
        _1191_v31 = new _dafny.CodePoint('a'.codePointAt(0));
        let _1192_v32;
        _1192_v32 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1156_i0);
        let _1193_v33;
        _1193_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1191_v31,new BigNumber((_1192_v32).length));
        let _1194_v34;
        _1194_v34 = _1193_v33;
        let _1195_v35;
        _1195_v35 = _module.D11.create_DC25(_1191_v31, new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1194_v34,_1154_v0)).update(_1194_v34, _1156_i0)).length), p0);
        let _source13 = _1195_v35;
        if (_source13.is_DC25) {
          let _1196___mcc_h0 = (_source13).cf41;
          let _1197___mcc_h1 = (_source13).cf42;
          let _1198___mcc_h2 = (_source13).cf43;
          let _1199_cf43 = _1198___mcc_h2;
          let _1200_cf42 = _1197___mcc_h1;
          let _1201_cf41 = _1196___mcc_h0;
          let _1202_v36;
          _1202_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1154_v0,_1156_i0);
          let _1203_v37;
          _1203_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1156_i0,p0);
          let _1204_v38;
          _1204_v38 = _dafny.Seq.of(_1199_cf43);
          let _1205_v39;
          _1205_v39 = _dafny.Set.fromElements(_1204_v38);
          let _1206_v40;
          let _nw189 = Array((new BigNumber(28)).toNumber());
          _nw189[(_dafny.ZERO).toNumber()] = _1156_i0;
          _nw189[(_dafny.ONE).toNumber()] = _1154_v0;
          _nw189[(new BigNumber(2)).toNumber()] = _1156_i0;
          _nw189[(new BigNumber(3)).toNumber()] = (_1154_v0).multipliedBy((((_1202_v36).contains(_1156_i0)) ? ((_1202_v36).get(_1156_i0)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(413)), ((_1207_v31) => function (_1208_i4) {
            return _1207_v31;
          })(_1191_v31))).length))));
          _nw189[(new BigNumber(4)).toNumber()] = _1156_i0;
          _nw189[(new BigNumber(5)).toNumber()] = _module.__default.fm32(_1156_i0, globalState);
          _nw189[(new BigNumber(6)).toNumber()] = (new BigNumber((_1157_v2).length)).plus(_1154_v0);
          _nw189[(new BigNumber(7)).toNumber()] = _1156_i0;
          _nw189[(new BigNumber(8)).toNumber()] = _1156_i0;
          _nw189[(new BigNumber(9)).toNumber()] = ((_1155_v1)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length), new BigNumber((_1155_v1).length))]).minus(new BigNumber((_1203_v37).length));
          _nw189[(new BigNumber(10)).toNumber()] = new BigNumber(((_1205_v39).Intersect(_1205_v39)).length);
          _nw189[(new BigNumber(11)).toNumber()] = _1200_cf42;
          _nw189[(new BigNumber(12)).toNumber()] = _1156_i0;
          _nw189[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(793)), ((_1209_cf41) => function (_1210_i5) {
            return _1209_cf41;
          })(_1201_cf41))).length);
          _nw189[(new BigNumber(14)).toNumber()] = _1156_i0;
          _nw189[(new BigNumber(15)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(487), _1200_cf42);
          _nw189[(new BigNumber(16)).toNumber()] = (_1156_i0).multipliedBy(new BigNumber((_1157_v2).length));
          _nw189[(new BigNumber(17)).toNumber()] = _1200_cf42;
          _nw189[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus((_1200_cf42).plus(_1156_i0));
          _nw189[(new BigNumber(19)).toNumber()] = _1200_cf42;
          _nw189[(new BigNumber(20)).toNumber()] = _1200_cf42;
          _nw189[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus(_1154_v0);
          _nw189[(new BigNumber(22)).toNumber()] = (_1200_cf42).multipliedBy(new BigNumber((_1157_v2).length));
          _nw189[(new BigNumber(23)).toNumber()] = _1156_i0;
          _nw189[(new BigNumber(24)).toNumber()] = ((p0) ? (_1156_i0) : (new BigNumber(940)));
          _nw189[(new BigNumber(25)).toNumber()] = _module.__default.safeDivisionInt(_1154_v0, _1156_i0);
          _nw189[(new BigNumber(26)).toNumber()] = (new BigNumber(727)).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(_1201_cf41)).length)));
          _nw189[(new BigNumber(27)).toNumber()] = _1156_i0;
          _1206_v40 = _nw189;
          let _index189 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_1206_v40).length));
          (_1206_v40)[_index189] = _module.__default.safeModuloInt(_module.__default.fm32(_1154_v0, globalState), new BigNumber(824));
          let _1211_v41;
          _1211_v41 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1206_v40);
          let _1212_v42;
          _1212_v42 = _dafny.Map.Empty.slice().updateUnsafe(_1211_v41,_1200_cf42);
          let _index190 = _module.__default.safeIndex(new BigNumber(518), new BigNumber((_1206_v40).length));
          (_1206_v40)[_index190] = (((_1212_v42).contains(_dafny.Map.Empty.slice().updateUnsafe(p0,_1206_v40))) ? ((_1212_v42).get(_dafny.Map.Empty.slice().updateUnsafe(p0,_1206_v40))) : (_1154_v0));
          let _1213_v43;
          let _init25 = ((_1214_cf43) => function (_1215_i6) {
            return _1214_cf43;
          })(_1199_cf43);
          let _nw190 = Array((new BigNumber(6)).toNumber());
          for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw190.length); _i0_25++) {
            _nw190[_i0_25] = _init25(new BigNumber(_i0_25));
          }
          _1213_v43 = _nw190;
          let _1216_v44;
          _1216_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1202_v36).length),_1213_v43);
          let _1217_v45;
          _1217_v45 = _dafny.Seq.of((((_1216_v44).contains((_1206_v40)[_module.__default.safeIndex(new BigNumber(518), new BigNumber((_1206_v40).length))])) ? ((_1216_v44).get((_1206_v40)[_module.__default.safeIndex(new BigNumber(518), new BigNumber((_1206_v40).length))])) : (_1213_v43)));
          let _1218_v46;
          _1218_v46 = _dafny.MultiSet.fromElements(_1213_v43);
          let _1219_v47;
          _1219_v47 = _module.D16.create_DC41((_1217_v45)[_module.__default.safeIndex((((_1218_v46).contains(_1213_v43)) ? ((_1218_v46).get(_1213_v43)) : (new BigNumber((_1204_v38).length))), new BigNumber((_1217_v45).length))]);
          _1219_v47 = _1219_v47;
          let _1220_v48;
          _1220_v48 = _dafny.Set.fromElements(_1200_cf42, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(956)), function (_1221_i7) {
            return new _dafny.CodePoint('s'.codePointAt(0));
          })).length));
          let _rhs123 = _1220_v48;
          let _rhs124 = _module.__default.safeModuloInt(new BigNumber(411), _1154_v0);
          let _rhs125 = (_1154_v0).multipliedBy(_1200_cf42);
          let _rhs126 = ((p0) ? (false) : (_1199_cf43));
          let _rhs127 = ((!(p0)) ? ((_1206_v40)[_module.__default.safeIndex(new BigNumber(518), new BigNumber((_1206_v40).length))]) : (_module.__default.safeDivisionInt(_1200_cf42, _1156_i0)));
          _1220_v48 = _rhs123;
          r1 = _rhs124;
          r1 = _rhs125;
          r3 = _rhs126;
          _1154_v0 = _rhs127;
          let _1222_v49;
          _1222_v49 = _dafny.Seq.of(_1204_v38);
          let _1223_v50;
          _1223_v50 = _module.D0.create_DC1((_1206_v40)[_module.__default.safeIndex(new BigNumber(518), new BigNumber((_1206_v40).length))]);
          let _1224_v51;
          _1224_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1201_cf41,_1223_v50);
          r3 = (_this).fm5((_dafny.ZERO).minus((_1156_i0).plus(new BigNumber((_1222_v49).length))), p0, _1224_v51, p0, globalState);
        } else if (_source13.is_DC24) {
          let _1225___mcc_h3 = (_source13).cf40;
          let _1226_cf40 = _1225___mcc_h3;
          r3 = p0;
          let _1227_v52;
          _1227_v52 = _dafny.Seq.of(p0, p0);
          let _1228_v53;
          _1228_v53 = _module.D0.create_DC1(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1191_v31,_1156_i0)).length));
          let _1229_v54;
          _1229_v54 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(p0, (_1227_v52)[_module.__default.safeIndex(_1156_i0, new BigNumber((_1227_v52).length))], (_this).fm5(_1156_i0, p0, _dafny.Map.Empty.slice().updateUnsafe(_1191_v31,_1228_v53), p0, globalState), p0, p0),_1154_v0);
          let _1230_v55;
          _1230_v55 = _dafny.Set.fromElements(p0);
          let _1231_v56;
          _1231_v56 = _dafny.Seq.of(_dafny.Set.fromElements(p0, p0));
          _1229_v54 = ((((_1231_v56)[_module.__default.safeIndex(_1156_i0, new BigNumber((_1231_v56).length))]).IsProperSubsetOf(_1230_v55)) ? (_module.__default.fm56(false, globalState)) : (((p0) ? (_1229_v54) : (_1229_v54))));
          let _1232_v57;
          let _init26 = ((_1233_p0) => function (_1234_i8) {
            return _1233_p0;
          })(p0);
          let _nw191 = Array((new BigNumber(15)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw191.length); _i0_26++) {
            _nw191[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _1232_v57 = _nw191;
          let _1235_v58;
          _1235_v58 = _dafny.MultiSet.fromElements(p0);
          let _index191 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_1232_v57).length));
          (_1232_v57)[_index191] = !((_1235_v58).IsProperSubsetOf(_1235_v58));
          let _index192 = _module.__default.safeIndex(new BigNumber(820), new BigNumber((_1232_v57).length));
          (_1232_v57)[_index192] = !((_1227_v52)[_module.__default.safeIndex((_1154_v0).plus(new BigNumber((_module.__default.fm57(p0, p0, p0, globalState)).length)), new BigNumber((_1227_v52).length))]);
          let _1236_v59;
          let _nw192 = Array((new BigNumber(21)).toNumber()).fill([]);
          _1236_v59 = _nw192;
          _1236_v59 = _1236_v59;
        } else {
          let _1237___mcc_h4 = (_source13).cf44;
          let _1238_cf44 = _1237___mcc_h4;
          let _1239_v60;
          _1239_v60 = _dafny.Map.Empty.slice().updateUnsafe(_1156_i0,(_1154_v0).isEqualTo(_1156_i0));
          _1239_v60 = _1239_v60;
          let _1240_v61;
          let _nw193 = new _module.C2();
          _nw193.__ctor();
          _1240_v61 = _nw193;
          let _1241_v62;
          _1241_v62 = _dafny.MultiSet.fromElements(p0);
          r0 = ((_1241_v62).IsDisjointFrom(_1241_v62)) || (p0);
          let _1242_v63;
          let _init27 = ((_1243_p0) => function (_1244_i9) {
            return !(_1243_p0);
          })(p0);
          let _nw194 = Array((new BigNumber(14)).toNumber());
          for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw194.length); _i0_27++) {
            _nw194[_i0_27] = _init27(new BigNumber(_i0_27));
          }
          _1242_v63 = _nw194;
          let _1245_v64;
          _1245_v64 = _dafny.Map.Empty.slice().updateUnsafe(_1154_v0,_1242_v63);
          let _1246_v65;
          _1246_v65 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_1157_v2, _1157_v2),_1245_v64);
          _1246_v65 = (_1246_v65).update(_1157_v2, _1245_v64);
        }
        r3 = p0;
      }
      let _1247_v66;
      _1247_v66 = _dafny.Seq.of(p0);
      let _1248_v67;
      _1248_v67 = _dafny.Set.fromElements(new BigNumber(980));
      let _1249_v69;
      let _nw195 = Array((new BigNumber(8)).toNumber());
      _nw195[(_dafny.ZERO).toNumber()] = p0;
      _nw195[(_dafny.ONE).toNumber()] = p0;
      _nw195[(new BigNumber(2)).toNumber()] = p0;
      _nw195[(new BigNumber(3)).toNumber()] = true;
      _nw195[(new BigNumber(4)).toNumber()] = false;
      _nw195[(new BigNumber(5)).toNumber()] = p0;
      _nw195[(new BigNumber(6)).toNumber()] = p0;
      _nw195[(new BigNumber(7)).toNumber()] = p0;
      _1249_v69 = _nw195;
      let _1250_v70;
      _1250_v70 = _dafny.MultiSet.fromElements(_1249_v69);
      let _1251_v71;
      _1251_v71 = _module.D4.create_DC11(_1154_v0, p0, p0, p0, !(p0));
      let _rhs128 = !(((_1248_v67).Intersect(function () {
        let _coll51 = new _dafny.Set();
        for (const _compr_51 of _dafny.IntegerRange(new BigNumber(445), new BigNumber(117))) {
          let _1252_v68 = _compr_51;
          if (((new BigNumber(445)).isLessThanOrEqualTo(_1252_v68)) && ((_1252_v68).isLessThan(new BigNumber(117)))) {
            _coll51.add((_1252_v68).plus(new BigNumber((_dafny.Seq.UnicodeFromString("yjfgxccy")).length)));
          }
        }
        return _coll51;
      }())).IsSubsetOf(_1248_v67));
      let _rhs129 = p0;
      let _rhs130 = !(_1250_v70).equals((_1250_v70).Intersect(_1250_v70));
      let _rhs131 = (!(p0) || ((_1251_v71).dtor_cf23)) && (p0);
      let _rhs132 = _module.__default.fm38(p0, _1154_v0, p0, globalState);
      r0 = _rhs128;
      r0 = _rhs129;
      r3 = _rhs130;
      r3 = _rhs131;
      _1247_v66 = _rhs132;
      let _1253_v72;
      _1253_v72 = p0;
      let _1254_v73;
      _1254_v73 = _dafny.MultiSet.fromElements(p0);
      let _1255_v74;
      _1255_v74 = _dafny.Map.Empty.slice().updateUnsafe(_1253_v72,_1254_v73);
      _1255_v74 = (_1255_v74).update(_1253_v72, (_dafny.MultiSet.fromElements(p0, p0)).Intersect(_1254_v73));
      let _1256_v75;
      _1256_v75 = _dafny.Map.Empty.slice().updateUnsafe((_1154_v0).isEqualTo(_1154_v0),new BigNumber((_1247_v66).length));
      _1256_v75 = _1256_v75;
      let _1257_v76;
      let _init28 = ((_1258_v0) => function (_1259_i10) {
        return _module.__default.safeModuloInt(_1259_i10, _1258_v0);
      })(_1154_v0);
      let _nw196 = Array((new BigNumber(13)).toNumber());
      for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw196.length); _i0_28++) {
        _nw196[_i0_28] = _init28(new BigNumber(_i0_28));
      }
      _1257_v76 = _nw196;
      _1257_v76 = _1257_v76;
      let _1260_v77;
      _1260_v77 = _dafny.Seq.UnicodeFromString("qtblq");
      let _1261_v78;
      _1261_v78 = _module.D9.create_DC21(_1254_v73);
      let _1262_v79;
      _1262_v79 = _module.D15.create_DC39(_1260_v77, _1261_v78);
      let _pat_let_tv33 = p0;
      let _pat_let_tv34 = p0;
      let _pat_let_tv35 = p0;
      let _pat_let_tv36 = p0;
      let _pat_let_tv37 = _1261_v78;
      r0 = function (_source14) {
        if (_source14.is_DC38) {
          let _1263___mcc_h5 = (_source14).cf71;
          let _1264_cf71 = _1263___mcc_h5;
          return _pat_let_tv33;
        } else if (_source14.is_DC39) {
          let _1265___mcc_h6 = (_source14).cf72;
          let _1266___mcc_h7 = (_source14).cf73;
          let _1267_cf73 = _1266___mcc_h7;
          let _1268_cf72 = _1265___mcc_h6;
          return _pat_let_tv34;
        } else if (_source14.is_DC37) {
          let _1269___mcc_h8 = (_source14).cf70;
          let _1270_cf70 = _1269___mcc_h8;
          return _pat_let_tv35;
        } else {
          let _1271___mcc_h9 = (_source14).cf74;
          let _1272_cf74 = _1271___mcc_h9;
          return _pat_let_tv36;
        }
      }(function (_pat_let32_0) {
        return function (_1273_dt__update__tmp_h0) {
          return function (_pat_let33_0) {
            return function (_1274_dt__update_hcf73_h0) {
              return _module.D15.create_DC39((_1273_dt__update__tmp_h0).dtor_cf72, _1274_dt__update_hcf73_h0);
            }(_pat_let33_0);
          }(_pat_let_tv37);
        }(_pat_let32_0);
      }(_1262_v79));
      r1 = _1154_v0;
      r2 = _1154_v0;
      r3 = p0;
      return [r0, r1, r2, r3];
    }
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
    }
    _parentTraits() {
      return [_module.T0];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(!(false));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _1275_v0;
      let _init29 = function (_1276_i0) {
        return true;
      };
      let _nw197 = Array((new BigNumber(24)).toNumber());
      for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw197.length); _i0_29++) {
        _nw197[_i0_29] = _init29(new BigNumber(_i0_29));
      }
      _1275_v0 = _nw197;
      let _1277_v1;
      _1277_v1 = true;
      let _index193 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
      (_1275_v0)[_index193] = _1277_v1;
      let _index194 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
      (_1275_v0)[_index194] = true;
      if (_1277_v1) {
        let _1278_v2;
        _1278_v2 = _dafny.Set.fromElements((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
        let _1279_v3;
        let _nw198 = Array((new BigNumber(14)).toNumber());
        _nw198[(_dafny.ZERO).toNumber()] = _1278_v2;
        _nw198[(_dafny.ONE).toNumber()] = _1278_v2;
        _nw198[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], true, _1277_v1);
        _nw198[(new BigNumber(3)).toNumber()] = _1278_v2;
        _nw198[(new BigNumber(4)).toNumber()] = _1278_v2;
        _nw198[(new BigNumber(5)).toNumber()] = _1278_v2;
        _nw198[(new BigNumber(6)).toNumber()] = _1278_v2;
        _nw198[(new BigNumber(7)).toNumber()] = _1278_v2;
        _nw198[(new BigNumber(8)).toNumber()] = _1278_v2;
        _nw198[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements(_1277_v1);
        _nw198[(new BigNumber(10)).toNumber()] = _1278_v2;
        _nw198[(new BigNumber(11)).toNumber()] = _dafny.Set.fromElements((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
        _nw198[(new BigNumber(12)).toNumber()] = _1278_v2;
        _nw198[(new BigNumber(13)).toNumber()] = _1278_v2;
        _1279_v3 = _nw198;
        let _1280_v4;
        _1280_v4 = _module.D0.create_DC0(_1279_v3, p0);
        let _1281_v5;
        _1281_v5 = _dafny.Map.Empty.slice().updateUnsafe((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))],_1280_v4);
        _1281_v5 = (_1281_v5).update(_1277_v1, _module.D0.create_DC0(_1279_v3, _module.__default.fm20(!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]), globalState)));
        let _index195 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        (_1275_v0)[_index195] = !(_module.__default.fm20(_1277_v1, globalState)).isEqualTo((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))],_1277_v1)).length)));
        let _1282_v6;
        _1282_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
        let _1283_v7;
        _1283_v7 = _dafny.Map.Empty.slice().updateUnsafe(false,_1282_v6);
        _1283_v7 = (_1283_v7).update((true) === (!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])), (_1282_v6).Merge(_module.__default.fm21(globalState)));
        let _index196 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        (_1275_v0)[_index196] = _1277_v1;
        let _1284_v8;
        let _nw199 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _1284_v8 = _nw199;
        let _1285_v9;
        _1285_v9 = _dafny.MultiSet.fromElements(_1284_v8, _1284_v8);
        let _index197 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        (_1275_v0)[_index197] = ((!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])) ? ((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) : (!((p0).isLessThanOrEqualTo(new BigNumber((_1285_v9).cardinality())))));
      } else {
        let _1286_v10;
        let _nw200 = Array((new BigNumber(23)).toNumber()).fill(_dafny.Seq.of());
        _1286_v10 = _nw200;
        let _init30 = ((_1287_p0) => function (_1288_i1) {
          return _dafny.Seq.of(_1287_p0, _1287_p0);
        })(p0);
        let _nw201 = Array((new BigNumber(8)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw201.length); _i0_30++) {
          _nw201[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1286_v10 = _nw201;
        let _1289_v11;
        let _nw202 = Array((new BigNumber(5)).toNumber()).fill([]);
        _1289_v11 = _nw202;
        let _1290_v12;
        _1290_v12 = _module.D2.create_DC4(p0, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], _1289_v11);
        let _index198 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        let _rhs133 = _1277_v1;
        let _rhs134 = (_1290_v12).dtor_cf6;
        let _lhs72 = _1275_v0;
        let _lhs73 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        _lhs72[_lhs73] = _rhs133;
        _1277_v1 = _rhs134;
        let _1291_v13;
        _1291_v13 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('t'.codePointAt(0)));
        let _1292_v14;
        _1292_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1291_v13,p0);
        _1292_v14 = (_1292_v14).update((_1291_v13).Difference(_1291_v13), new BigNumber(474));
        let _1293_v15;
        let _nw203 = new _module.C0();
        _nw203.__ctor(_dafny.Seq.UnicodeFromString("o"));
        _1293_v15 = _nw203;
        let _1294_v16;
        _1294_v16 = _dafny.Set.fromElements((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
        let _1295_v17;
        _1295_v17 = _dafny.Map.Empty.slice().updateUnsafe(_1275_v0,_1294_v16);
        let _1296_v18;
        let _nw204 = Array((new BigNumber(16)).toNumber());
        _nw204[(_dafny.ZERO).toNumber()] = (_1294_v16).Difference(_1294_v16);
        _nw204[(_dafny.ONE).toNumber()] = _1294_v16;
        _nw204[(new BigNumber(2)).toNumber()] = (_dafny.Set.fromElements(true)).Intersect(_1294_v16);
        _nw204[(new BigNumber(3)).toNumber()] = _1294_v16;
        _nw204[(new BigNumber(4)).toNumber()] = _1294_v16;
        _nw204[(new BigNumber(5)).toNumber()] = _1294_v16;
        _nw204[(new BigNumber(6)).toNumber()] = _1294_v16;
        _nw204[(new BigNumber(7)).toNumber()] = _1294_v16;
        _nw204[(new BigNumber(8)).toNumber()] = _1294_v16;
        _nw204[(new BigNumber(9)).toNumber()] = _dafny.Set.fromElements((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], true);
        _nw204[(new BigNumber(10)).toNumber()] = _1294_v16;
        _nw204[(new BigNumber(11)).toNumber()] = (_1294_v16).Difference(_1294_v16);
        _nw204[(new BigNumber(12)).toNumber()] = _1294_v16;
        _nw204[(new BigNumber(13)).toNumber()] = ((true) ? (_dafny.Set.fromElements(true, _1277_v1, _1277_v1, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])) : ((((_1295_v17).contains(_1275_v0)) ? ((_1295_v17).get(_1275_v0)) : (_1294_v16))));
        _nw204[(new BigNumber(14)).toNumber()] = (_1294_v16).Intersect(_1294_v16);
        _nw204[(new BigNumber(15)).toNumber()] = _1294_v16;
        _1296_v18 = _nw204;
        _1296_v18 = _1296_v18;
      }
      let _1297_v19;
      let _init31 = ((_1298_v1, _1299_v0) => function (_1300_i2) {
        return _dafny.Set.fromElements(_1298_v1, _1298_v1, (_1299_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1299_v0).length))], _1298_v1);
      })(_1277_v1, _1275_v0);
      let _nw205 = Array((new BigNumber(26)).toNumber());
      for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw205.length); _i0_31++) {
        _nw205[_i0_31] = _init31(new BigNumber(_i0_31));
      }
      _1297_v19 = _nw205;
      let _1301_v20;
      _1301_v20 = _module.D0.create_DC0(_1297_v19, p0);
      let _1302_v21;
      _1302_v21 = _dafny.Seq.of(_1301_v20);
      let _1303_v22;
      _1303_v22 = _module.D4.create_DC8(_1302_v21);
      if (_dafny.Seq.IsPrefixOf((_1303_v22).dtor_cf13, _1302_v21)) {
        let _1304_v23;
        _1304_v23 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(689),_1277_v1);
        let _1305_v24;
        _1305_v24 = _dafny.Seq.UnicodeFromString("wuqe");
        let _1306_v25;
        _1306_v25 = _dafny.Seq.of(new BigNumber((_1305_v24).length));
        let _1307_v26;
        _1307_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1277_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1306_v25).length),p0)).length));
        let _1308_v27;
        _1308_v27 = _dafny.MultiSet.fromElements(_1277_v1);
        let _1309_v28;
        _1309_v28 = _dafny.MultiSet.fromElements(p0);
        let _1310_v29;
        _1310_v29 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),new BigNumber(159));
        let _1311_v30;
        _1311_v30 = _dafny.Seq.of(_1310_v29);
        let _1312_v31;
        _1312_v31 = _dafny.MultiSet.fromElements(_1275_v0);
        let _1313_v32;
        let _nw206 = Array((new BigNumber(22)).toNumber());
        _nw206[(_dafny.ZERO).toNumber()] = p0;
        _nw206[(_dafny.ONE).toNumber()] = new BigNumber((_1304_v23).length);
        _nw206[(new BigNumber(2)).toNumber()] = p0;
        _nw206[(new BigNumber(3)).toNumber()] = new BigNumber((_1307_v26).length);
        _nw206[(new BigNumber(4)).toNumber()] = new BigNumber((_1305_v24).length);
        _nw206[(new BigNumber(5)).toNumber()] = p0;
        _nw206[(new BigNumber(6)).toNumber()] = p0;
        _nw206[(new BigNumber(7)).toNumber()] = new BigNumber((_1305_v24).length);
        _nw206[(new BigNumber(8)).toNumber()] = new BigNumber((_1308_v27).cardinality());
        _nw206[(new BigNumber(9)).toNumber()] = p0;
        _nw206[(new BigNumber(10)).toNumber()] = (_1306_v25)[_module.__default.safeIndex(p0, new BigNumber((_1306_v25).length))];
        _nw206[(new BigNumber(11)).toNumber()] = new BigNumber(-17);
        _nw206[(new BigNumber(12)).toNumber()] = (_1306_v25)[_module.__default.safeIndex(p0, new BigNumber((_1306_v25).length))];
        _nw206[(new BigNumber(13)).toNumber()] = p0;
        _nw206[(new BigNumber(14)).toNumber()] = new BigNumber((_1309_v28).cardinality());
        _nw206[(new BigNumber(15)).toNumber()] = p0;
        _nw206[(new BigNumber(16)).toNumber()] = new BigNumber(((_1311_v30)[_module.__default.safeIndex(p0, new BigNumber((_1311_v30).length))]).length);
        _nw206[(new BigNumber(17)).toNumber()] = (((_1312_v31).contains(_1275_v0)) ? ((_1312_v31).get(_1275_v0)) : (p0));
        _nw206[(new BigNumber(18)).toNumber()] = p0;
        _nw206[(new BigNumber(19)).toNumber()] = p0;
        _nw206[(new BigNumber(20)).toNumber()] = new BigNumber((_module.__default.fm22(false, p0, new BigNumber((_1308_v27).cardinality()), globalState)).length);
        _nw206[(new BigNumber(21)).toNumber()] = p0;
        _1313_v32 = _nw206;
        let _1314_v33;
        _1314_v33 = _dafny.Seq.of(_1313_v32, _1313_v32, _1313_v32, _1313_v32, _1313_v32);
        let _1315_v34;
        _1315_v34 = _dafny.Map.Empty.slice().updateUnsafe((_1314_v33)[_module.__default.safeIndex(p0, new BigNumber((_1314_v33).length))],p0);
        _1315_v34 = (_1315_v34).update(((_1277_v1) ? (_1313_v32) : (_1313_v32)), p0);
        let _1316_v35;
        let _nw207 = new _module.C0();
        _nw207.__ctor(_1305_v24);
        _1316_v35 = _nw207;
        let _1317_v36;
        let _nw208 = Array((new BigNumber(3)).toNumber());
        _nw208[(_dafny.ZERO).toNumber()] = _1314_v33;
        _nw208[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1314_v33, _1314_v33);
        _nw208[(new BigNumber(2)).toNumber()] = _1314_v33;
        _1317_v36 = _nw208;
        let _index199 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1317_v36).length));
        (_1317_v36)[_index199] = _dafny.Seq.of(_1313_v32);
        let _index200 = _module.__default.safeIndex(new BigNumber(233), new BigNumber((_1317_v36).length));
        (_1317_v36)[_index200] = _dafny.Seq.of(_1313_v32);
        let _1318_v37;
        _1318_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1277_v1,(_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
        let _source15 = _module.D4.create_DC11((new BigNumber((_1306_v25).length)).plus(p0), (!(false)) || ((((_1318_v37).contains(true)) ? ((_1318_v37).get(true)) : (false))), _1277_v1, (new BigNumber(-703)).isLessThanOrEqualTo(p0), _1277_v1);
        if (_source15.is_DC9) {
          let _1319___mcc_h0 = (_source15).cf14;
          let _1320___mcc_h1 = (_source15).cf15;
          let _1321___mcc_h2 = (_source15).cf16;
          let _1322___mcc_h3 = (_source15).cf17;
          let _1323___mcc_h4 = (_source15).cf18;
          let _1324_cf18 = _1323___mcc_h4;
          let _1325_cf17 = _1322___mcc_h3;
          let _1326_cf16 = _1321___mcc_h2;
          let _1327_cf15 = _1320___mcc_h1;
          let _1328_cf14 = _1319___mcc_h0;
          _1324_cf18 = ((_1326_cf16).plus((_dafny.ZERO).minus(p0))).plus(((false) ? (_1326_cf16) : (_1324_cf18)));
          let _1329_v38;
          let _nw209 = Array((new BigNumber(9)).toNumber());
          _1329_v38 = _nw209;
          let _index201 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1329_v38).length));
          (_1329_v38)[_index201] = _1316_v35;
          let _index202 = _module.__default.safeIndex(new BigNumber(476), new BigNumber((_1329_v38).length));
          (_1329_v38)[_index202] = _1316_v35;
          let _1330_v39;
          _1330_v39 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),_module.D0.create_DC1(new BigNumber(486)));
          r1 = (_this).fm5(_1326_cf16, _1277_v1, (_1330_v39).Merge(_1330_v39), (((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) ? ((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) : (!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]))), globalState);
          let _index203 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
          (_1275_v0)[_index203] = (((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) ? (!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])) : ((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]));
        } else if (_source15.is_DC10) {
          let _1331___mcc_h5 = (_source15).cf19;
          let _1332___mcc_h6 = (_source15).cf20;
          let _1333___mcc_h7 = (_source15).cf21;
          let _1334_cf21 = _1333___mcc_h7;
          let _1335_cf20 = _1332___mcc_h6;
          let _1336_cf19 = _1331___mcc_h5;
          let _1337_v40;
          _1337_v40 = new _dafny.CodePoint('k'.codePointAt(0));
          _1307_v26 = (_1307_v26).update(!_dafny.Seq.contains(_1316_v35.f8, _1337_v40), (_1336_cf19).plus(new BigNumber((_1306_v25).length)));
          let _1338_v41;
          _1338_v41 = _dafny.Set.fromElements(_1305_v24, _dafny.Seq.of(_1337_v40, _1337_v40));
          let _1339_v42;
          _1339_v42 = _module.D2.create_DC3(_1338_v41);
          let _1340_v43;
          _1340_v43 = _module.D2.create_DC6(_1339_v42);
          let _1341_v44;
          _1341_v44 = _dafny.Map.Empty.slice().updateUnsafe(_1340_v43,_1316_v35.f8);
          _1341_v44 = _1341_v44;
          let _1342_v45;
          _1342_v45 = _module.D0.create_DC1(new BigNumber((_1308_v27).cardinality()));
          let _1343_v46;
          _1343_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1337_v40,_1342_v45);
          _1334_cf21 = _module.__default.safeModuloInt(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_1277_v1, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], (((_1318_v37).contains((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])) ? ((_1318_v37).get((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])) : (_1277_v1)))).length), _1335_cf20), (((_1307_v26).contains((_this).fm5(_1334_cf21, _1277_v1, _1343_v46, _1277_v1, globalState))) ? ((_1307_v26).get((_this).fm5(_1334_cf21, _1277_v1, _1343_v46, _1277_v1, globalState))) : (new BigNumber((_dafny.Set.fromElements((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])).length))));
          _1307_v26 = ((_1307_v26).Merge(_1307_v26)).Merge((_dafny.Map.Empty.slice().updateUnsafe((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))],new BigNumber(-533))).Merge(_1307_v26));
        } else if (_source15.is_DC11) {
          let _1344___mcc_h8 = (_source15).cf22;
          let _1345___mcc_h9 = (_source15).cf23;
          let _1346___mcc_h10 = (_source15).cf24;
          let _1347___mcc_h11 = (_source15).cf25;
          let _1348___mcc_h12 = (_source15).cf26;
          let _1349_cf26 = _1348___mcc_h12;
          let _1350_cf25 = _1347___mcc_h11;
          let _1351_cf24 = _1346___mcc_h10;
          let _1352_cf23 = _1345___mcc_h9;
          let _1353_cf22 = _1344___mcc_h8;
          let _1354_v47;
          _1354_v47 = _dafny.Seq.of(_1351_cf24);
          let _1355_v48;
          _1355_v48 = _dafny.Seq.of(_1354_v47, _module.__default.fm23(_1277_v1, globalState), _dafny.Seq.update(_1354_v47, _module.__default.safeIndex(_1353_cf22, new BigNumber((_1354_v47).length)), _1350_cf25));
          _1355_v48 = _1355_v48;
          let _1356_v49;
          _1356_v49 = _dafny.Seq.of(_1275_v0);
          let _1357_v50;
          let _1358_v51;
          let _out37;
          let _out38;
          let _outcollector11 = (_this).m10(_1353_cf22, _1353_cf22, !(!(_1277_v1) || ((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])), (_1356_v49)[_module.__default.safeIndex(p0, new BigNumber((_1356_v49).length))], globalState);
          _out37 = _outcollector11[0];
          _out38 = _outcollector11[1];
          _1357_v50 = _out37;
          _1358_v51 = _out38;
          _1353_cf22 = _1353_cf22;
          let _1359_v52;
          let _nw210 = Array((new BigNumber(10)).toNumber()).fill([]);
          _1359_v52 = _nw210;
          _1359_v52 = _1359_v52;
        } else if (_source15.is_DC8) {
          let _1360___mcc_h13 = (_source15).cf13;
          let _1361_cf13 = _1360___mcc_h13;
          let _1362_v53;
          _1362_v53 = new BigNumber(-892);
          _1362_v53 = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(266)), function (_1363_i3) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          })).length)));
          let _1364_v54;
          _1364_v54 = new _dafny.CodePoint('f'.codePointAt(0));
          _1364_v54 = ((_1277_v1) ? (_1364_v54) : (_1364_v54));
          let _1365_v55;
          _1365_v55 = _module.D2.create_DC6(_module.D2.create_DC5((_dafny.ZERO).minus(_1362_v53), _1306_v25, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]));
          let _1366_v56;
          _1366_v56 = _dafny.Seq.of(_1365_v55);
          let _1367_v57;
          _1367_v57 = _dafny.Seq.contains(_1366_v56, _1365_v55);
          let _1368_v58;
          _1368_v58 = _dafny.Map.Empty.slice().updateUnsafe(_1364_v54,_1275_v0);
          let _1369_v59;
          _1369_v59 = _dafny.Set.fromElements(_1316_v35, _1316_v35);
          let _1370_v60;
          _1370_v60 = _module.D4.create_DC9((_module.__default.fm24((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], globalState)).update((_dafny.ZERO).minus((_dafny.ZERO).minus(_1362_v53)), _module.__default.abs(_1362_v53)), new _dafny.CodePoint('d'.codePointAt(0)), (new BigNumber((_dafny.Set.fromElements(_1277_v1)).length)).plus((_dafny.ZERO).minus(new BigNumber((_1304_v23).length))), _1369_v59, (_1362_v53).minus(_1362_v53));
          let _1371_v61;
          _1371_v61 = _module.D0.create_DC1(p0);
          let _1372_v62;
          _1372_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1364_v54,_1371_v61);
          let _index204 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
          let _rhs135 = _1367_v57;
          let _rhs136 = (true) || ((_this).fm5(_1362_v53, true, _1372_v62, !(!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])), globalState));
          let _rhs137 = _1368_v58;
          let _rhs138 = _1370_v60;
          let _lhs74 = _1275_v0;
          let _lhs75 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
          _1367_v57 = _rhs135;
          _lhs74[_lhs75] = _rhs136;
          _1368_v58 = _rhs137;
          _1370_v60 = _rhs138;
          let _1373_v63;
          let _nw211 = Array((new BigNumber(15)).toNumber()).fill([]);
          _1373_v63 = _nw211;
          let _1374_v64;
          let _nw212 = Array((new BigNumber(17)).toNumber());
          _nw212[(_dafny.ZERO).toNumber()] = _1313_v32;
          _nw212[(_dafny.ONE).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(2)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(3)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(4)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(5)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(6)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(7)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(8)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(9)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(10)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(11)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(12)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(13)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(14)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(15)).toNumber()] = _1313_v32;
          _nw212[(new BigNumber(16)).toNumber()] = _1313_v32;
          _1374_v64 = _nw212;
          let _index205 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_1373_v63).length));
          (_1373_v63)[_index205] = _1374_v64;
          let _index206 = _module.__default.safeIndex(new BigNumber(432), new BigNumber((_1373_v63).length));
          (_1373_v63)[_index206] = _1374_v64;
        } else {
          let _1375___mcc_h14 = (_source15).cf27;
          let _1376_cf27 = _1375___mcc_h14;
          _1277_v1 = ((_1277_v1) ? (((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) || ((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])) : ((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]));
          _1313_v32 = _1313_v32;
          let _index207 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
          (_1275_v0)[_index207] = _1277_v1;
          let _1377_v65;
          _1377_v65 = new BigNumber(226);
          let _1378_v66;
          let _nw213 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
          _1378_v66 = _nw213;
          let _1379_v67;
          _1379_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1277_v1,_1275_v0);
          let _index208 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_1378_v66).length));
          (_1378_v66)[_index208] = _1379_v67;
          let _1380_v68;
          let _init32 = ((_1381_p0) => function (_1382_i4) {
            return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),_1381_p0);
          })(p0);
          let _nw214 = Array((new BigNumber(2)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw214.length); _i0_32++) {
            _nw214[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1380_v68 = _nw214;
          let _1383_v69;
          _1383_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1377_v65,_1380_v68);
          let _1384_v70;
          _1384_v70 = _1380_v68;
          let _1385_v71;
          let _nw215 = Array((new BigNumber(26)).toNumber());
          _nw215[(_dafny.ZERO).toNumber()] = _1380_v68;
          _nw215[(_dafny.ONE).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(2)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(3)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(4)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(5)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(6)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(7)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(8)).toNumber()] = (((_1383_v69).contains(new BigNumber((_1306_v25).length))) ? ((_1383_v69).get(new BigNumber((_1306_v25).length))) : (_1380_v68));
          _nw215[(new BigNumber(9)).toNumber()] = (_1384_v70);
          _nw215[(new BigNumber(10)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(11)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(12)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(13)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(14)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(15)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(16)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(17)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(18)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(19)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(20)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(21)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(22)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(23)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(24)).toNumber()] = _1380_v68;
          _nw215[(new BigNumber(25)).toNumber()] = _1380_v68;
          _1385_v71 = _nw215;
          let _index209 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1385_v71).length));
          (_1385_v71)[_index209] = _1380_v68;
          let _1386_v72;
          _1386_v72 = new _dafny.CodePoint('p'.codePointAt(0));
          let _1387_v73;
          _1387_v73 = _dafny.Set.fromElements(_1277_v1, _1277_v1, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
          let _1388_v74;
          _1388_v74 = _dafny.Seq.of(_1387_v73);
          let _1389_v75;
          _1389_v75 = _dafny.Seq.of(_1379_v67);
          let _index210 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_1378_v66).length));
          let _index211 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1385_v71).length));
          let _rhs139 = !(false);
          let _rhs140 = new BigNumber((_1388_v74).length);
          let _rhs141 = (_1379_v67).Merge((_1389_v75)[_module.__default.safeIndex(_1377_v65, new BigNumber((_1389_v75).length))]);
          let _rhs142 = _1380_v68;
          let _rhs143 = new _dafny.CodePoint('d'.codePointAt(0));
          let _lhs76 = _1378_v66;
          let _lhs77 = _module.__default.safeIndex(new BigNumber(390), new BigNumber((_1378_v66).length));
          let _lhs78 = _1385_v71;
          let _lhs79 = _module.__default.safeIndex(_dafny.ZERO, new BigNumber((_1385_v71).length));
          r0 = _rhs139;
          _1377_v65 = _rhs140;
          _lhs76[_lhs77] = _rhs141;
          _lhs78[_lhs79] = _rhs142;
          _1386_v72 = _rhs143;
        }
        let _1390_v76;
        let _1391_v77;
        let _1392_v78;
        let _out39;
        let _out40;
        let _out41;
        let _outcollector12 = (_this).m9(globalState);
        _out39 = _outcollector12[0];
        _out40 = _outcollector12[1];
        _out41 = _outcollector12[2];
        _1390_v76 = _out39;
        _1391_v77 = _out40;
        _1392_v78 = _out41;
      } else {
        let _1393_v79;
        let _1394_v80;
        let _1395_v81;
        let _out42;
        let _out43;
        let _out44;
        let _outcollector13 = (_this).m9(globalState);
        _out42 = _outcollector13[0];
        _out43 = _outcollector13[1];
        _out44 = _outcollector13[2];
        _1393_v79 = _out42;
        _1394_v80 = _out43;
        _1395_v81 = _out44;
        let _1396_v83;
        _1396_v83 = _dafny.Set.fromElements(new BigNumber(-990), p0);
        let _1397_v84;
        _1397_v84 = _dafny.Set.fromElements(new BigNumber((_1396_v83).length), p0);
        let _1398_v85;
        _1398_v85 = _dafny.MultiSet.fromElements(p0, new BigNumber((_1396_v83).length));
        let _1399_v86;
        _1399_v86 = new _dafny.CodePoint('a'.codePointAt(0));
        let _1400_v87;
        _1400_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1398_v85).cardinality()),_1399_v86);
        let _1401_v88;
        _1401_v88 = _dafny.Seq.of((function () {
          let _coll52 = new _dafny.Map();
          for (const _compr_52 of _dafny.IntegerRange(new BigNumber(-592), new BigNumber(863))) {
            let _1402_v82 = _compr_52;
            if (((new BigNumber(-592)).isLessThanOrEqualTo(_1402_v82)) && ((_1402_v82).isLessThan(new BigNumber(863)))) {
              _coll52.push([(_1402_v82).minus(new BigNumber((_dafny.Seq.UnicodeFromString("dmetdhseu")).length)),new _dafny.CodePoint('y'.codePointAt(0))]);
            }
          }
          return _coll52;
        }()).Merge(_1400_v87));
        _1401_v88 = _1401_v88;
        let _1403_v89;
        _1403_v89 = new BigNumber(768);
        let _1404_v90;
        _1404_v90 = _dafny.Seq.of(!(_1277_v1));
        let _1405_v91;
        _1405_v91 = _dafny.Seq.of(_1404_v90);
        let _1406_v92;
        _1406_v92 = _dafny.Seq.of(new BigNumber(732), new BigNumber((_dafny.Seq.Concat(_1404_v90, (_1405_v91)[_module.__default.safeIndex(p0, new BigNumber((_1405_v91).length))])).length), p0, p0, _1403_v89);
        let _1407_v93;
        _1407_v93 = _dafny.Map.Empty.slice().updateUnsafe(_1395_v81,p0);
        let _1408_v94;
        _1408_v94 = _dafny.Seq.UnicodeFromString("wvfkivcdy");
        let _rhs144 = new BigNumber(665);
        let _rhs145 = (((_1407_v93).contains(_1277_v1)) ? ((_1407_v93).get(_1277_v1)) : (_1403_v89));
        let _rhs146 = !(_dafny.Seq.IsPrefixOf(_1408_v94, _1408_v94)) || ((new BigNumber(890)).isEqualTo((((_1398_v85).contains(p0)) ? ((_1398_v85).get(p0)) : (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p0,p0)).length)))));
        let _rhs147 = _1406_v92;
        _1403_v89 = _rhs144;
        _1403_v89 = _rhs145;
        _1395_v81 = _rhs146;
        _1406_v92 = _rhs147;
        let _1409_v95;
        _1409_v95 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(967)), ((_1410_v86) => function (_1411_i5) {
          return _1410_v86;
        })(_1399_v86)), _dafny.Seq.update((_module.D6.create_DC14(_dafny.Seq.UnicodeFromString("erond"))).dtor_cf29, _module.__default.safeIndex(p0, new BigNumber(((_module.D6.create_DC14(_dafny.Seq.UnicodeFromString("erond"))).dtor_cf29).length)), new _dafny.CodePoint('f'.codePointAt(0))), _dafny.Seq.Create(_module.__default.abs(new BigNumber(523)), ((_1412_v86) => function (_1413_i6) {
          return _1412_v86;
        })(_1399_v86)));
        let _1414_v96;
        let _nw216 = new _module.C0();
        _nw216.__ctor(_dafny.Seq.Concat(_1408_v94, (_1409_v95)[_module.__default.safeIndex(_1403_v89, new BigNumber((_1409_v95).length))]));
        _1414_v96 = _nw216;
        let _1415_v97;
        _1415_v97 = _dafny.Map.Empty.slice().updateUnsafe((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))],_1395_v81);
        let _1416_v98;
        _1416_v98 = _module.D2.create_DC5(new BigNumber((_1415_v97).length), _1406_v92, _1395_v81);
        _1277_v1 = ((_1395_v81) ? ((_1416_v98).dtor_cf10) : (!((_1404_v90)[_module.__default.safeIndex(_1403_v89, new BigNumber((_1404_v90).length))])));
      }
      let _1417_v99;
      _1417_v99 = _dafny.Set.fromElements(p0, p0, p0);
      let _1418_v100;
      _1418_v100 = _dafny.Map.Empty.slice().updateUnsafe((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))],new BigNumber((_1417_v99).length));
      let _1419_v101;
      _1419_v101 = new _dafny.CodePoint('h'.codePointAt(0));
      let _1420_v102;
      _1420_v102 = _module.D0.create_DC1(p0);
      let _1421_v103;
      _1421_v103 = _dafny.Map.Empty.slice().updateUnsafe(_1419_v101,_1420_v102);
      if (!((_this).fm5(new BigNumber((_1418_v100).length), _1277_v1, _1421_v103, _1277_v1, globalState)) || (!(p0).isEqualTo(p0))) {
        let _1422_v104;
        let _nw217 = Array((new BigNumber(22)).toNumber()).fill(_dafny.ZERO);
        _1422_v104 = _nw217;
        let _1423_v105;
        _1423_v105 = _dafny.Seq.UnicodeFromString("awvnycoec");
        let _index212 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length));
        (_1422_v104)[_index212] = new BigNumber((_1423_v105).length);
        let _1424_v106;
        _1424_v106 = new BigNumber(125);
        let _index213 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1422_v104).length));
        (_1422_v104)[_index213] = p0;
        let _index214 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length));
        let _index215 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1422_v104).length));
        let _rhs148 = _1424_v106;
        let _rhs149 = (_1424_v106).minus(_1424_v106);
        let _rhs150 = p0;
        let _lhs80 = _1422_v104;
        let _lhs81 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length));
        let _lhs82 = _1422_v104;
        let _lhs83 = _module.__default.safeIndex(new BigNumber(155), new BigNumber((_1422_v104).length));
        _lhs80[_lhs81] = _rhs148;
        _1424_v106 = _rhs149;
        _lhs82[_lhs83] = _rhs150;
        let _1425_v107;
        _1425_v107 = _dafny.MultiSet.fromElements(_1275_v0, _1275_v0);
        let _1426_v108;
        _1426_v108 = _module.D4.create_DC10(p0, _1424_v106, _1424_v106);
        let _index216 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length));
        let _rhs151 = _module.__default.safeModuloInt(new BigNumber((((_dafny.MultiSet.fromElements(_1275_v0)).update(_1275_v0, _module.__default.abs(p0))).Union(_1425_v107)).cardinality()), _1424_v106);
        let _rhs152 = _module.__default.fm25(_1426_v108, _1423_v105, globalState);
        let _lhs84 = _1422_v104;
        let _lhs85 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length));
        _lhs84[_lhs85] = _rhs151;
        _1423_v105 = _rhs152;
        let _1427_v109;
        _1427_v109 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1422_v104)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length))]);
        let _1428_v110;
        let _nw218 = Array((new BigNumber(14)).toNumber()).fill([]);
        _1428_v110 = _nw218;
        let _1429_v111;
        _1429_v111 = _module.D2.create_DC4((((_1427_v109).contains(p0)) ? ((_1427_v109).get(p0)) : (new BigNumber(-269))), false, _1428_v110);
        if ((_module.__default.fm24((_1429_v111).dtor_cf6, globalState)).IsSubsetOf(_dafny.MultiSet.fromElements(p0, p0))) {
          let _1430_v112;
          _1430_v112 = _dafny.Seq.of((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
          let _1431_v113;
          _1431_v113 = _module.D7.create_DC17(_1430_v112);
          let _1432_v114;
          let _nw219 = Array((new BigNumber(20)).toNumber());
          _nw219[(_dafny.ZERO).toNumber()] = _dafny.Seq.of((_1429_v111).dtor_cf6, false);
          _nw219[(_dafny.ONE).toNumber()] = _1430_v112;
          _nw219[(new BigNumber(2)).toNumber()] = _1430_v112;
          _nw219[(new BigNumber(3)).toNumber()] = _1430_v112;
          _nw219[(new BigNumber(4)).toNumber()] = _1430_v112;
          _nw219[(new BigNumber(5)).toNumber()] = _dafny.Seq.of((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
          _nw219[(new BigNumber(6)).toNumber()] = _1430_v112;
          _nw219[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], _1277_v1, _1277_v1, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]), (_1431_v113).dtor_cf34);
          _nw219[(new BigNumber(8)).toNumber()] = _module.__default.fm23((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], globalState);
          _nw219[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_1430_v112, _dafny.Seq.of(_1277_v1));
          _nw219[(new BigNumber(10)).toNumber()] = _1430_v112;
          _nw219[(new BigNumber(11)).toNumber()] = _1430_v112;
          _nw219[(new BigNumber(12)).toNumber()] = _1430_v112;
          _nw219[(new BigNumber(13)).toNumber()] = _1430_v112;
          _nw219[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat(_1430_v112, _1430_v112);
          _nw219[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(true, _1277_v1, false);
          _nw219[(new BigNumber(16)).toNumber()] = _dafny.Seq.update(_1430_v112, _module.__default.safeIndex(new BigNumber(82), new BigNumber((_1430_v112).length)), (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
          _nw219[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(!(!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])), (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
          _nw219[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_1277_v1, _1277_v1, _1277_v1), _1430_v112);
          _nw219[(new BigNumber(19)).toNumber()] = _dafny.Seq.Concat(_1430_v112, _1430_v112);
          _1432_v114 = _nw219;
          let _index217 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_1432_v114).length));
          (_1432_v114)[_index217] = _dafny.Seq.Concat(_1430_v112, _1430_v112);
          let _1433_v115;
          _1433_v115 = _dafny.Seq.of((_1422_v104)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length))], new BigNumber((_1430_v112).length), _1424_v106);
          let _1434_v116;
          _1434_v116 = _dafny.Map.Empty.slice().updateUnsafe(_1433_v115,(_1422_v104)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length))]);
          let _index218 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
          let _index219 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length));
          let _index220 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_1432_v114).length));
          let _rhs153 = !((_1424_v106).isEqualTo(_module.__default.fm20(_1277_v1, globalState)));
          let _rhs154 = (((_1434_v116).contains(_1433_v115)) ? ((_1434_v116).get(_1433_v115)) : ((_1422_v104)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length))]));
          let _rhs155 = _dafny.Seq.Concat(_dafny.Seq.of(_1277_v1, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]), _dafny.Seq.of(false));
          let _lhs86 = _1275_v0;
          let _lhs87 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
          let _lhs88 = _1422_v104;
          let _lhs89 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length));
          let _lhs90 = _1432_v114;
          let _lhs91 = _module.__default.safeIndex(new BigNumber(198), new BigNumber((_1432_v114).length));
          _lhs86[_lhs87] = _rhs153;
          _lhs88[_lhs89] = _rhs154;
          _lhs90[_lhs91] = _rhs155;
          let _1435_v117;
          let _nw220 = new _module.C0();
          _nw220.__ctor(_1423_v105);
          _1435_v117 = _nw220;
          let _1436_v118;
          _1436_v118 = _dafny.Map.Empty.slice().updateUnsafe(_1435_v117,!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]));
          _1436_v118 = _1436_v118;
          let _1437_v119;
          let _nw221 = new _module.C0();
          _nw221.__ctor(_dafny.Seq.Concat(_module.__default.fm25(_1426_v108, _1423_v105, globalState), _1423_v105));
          _1437_v119 = _nw221;
          let _1438_v120;
          _1438_v120 = _dafny.Set.fromElements(_1277_v1, !((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]));
          let _1439_v121;
          _1439_v121 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_1438_v120).length), new BigNumber(767)));
          _1439_v121 = _dafny.Seq.update(_dafny.Seq.Concat(_1439_v121, _1439_v121), _module.__default.safeIndex(_1424_v106, new BigNumber((_dafny.Seq.Concat(_1439_v121, _1439_v121)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(790)), ((_1440_v109) => function (_1441_i7) {
            return new BigNumber((_1440_v109).length);
          })(_1427_v109)));
          let _1442_v122;
          _1442_v122 = _dafny.Map.Empty.slice().updateUnsafe(_1419_v101,_module.__default.fm20((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], globalState));
          _1442_v122 = (_1442_v122).update(_1419_v101, p0);
        } else {
          let _1443_v123;
          _1443_v123 = _dafny.Map.Empty.slice().updateUnsafe(_1277_v1,_1275_v0);
          let _1444_v124;
          _1444_v124 = _dafny.Set.fromElements(_1275_v0, (((_1443_v123).contains(!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]))) ? ((_1443_v123).get(!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]))) : (_1275_v0)));
          let _1445_v125;
          let _nw222 = Array((new BigNumber(4)).toNumber());
          _nw222[(_dafny.ZERO).toNumber()] = _1444_v124;
          _nw222[(_dafny.ONE).toNumber()] = _1444_v124;
          _nw222[(new BigNumber(2)).toNumber()] = _dafny.Set.fromElements(_1275_v0, _1275_v0);
          _nw222[(new BigNumber(3)).toNumber()] = (_1444_v124).Union(_dafny.Set.fromElements(_1275_v0));
          _1445_v125 = _nw222;
          let _1446_v126;
          _1446_v126 = _1444_v124;
          let _index221 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_1445_v125).length));
          (_1445_v125)[_index221] = ((_1277_v1) ? ((_1446_v126)) : (_1444_v124));
          let _1447_v127;
          _1447_v127 = _dafny.Map.Empty.slice().updateUnsafe(false,(_1444_v124).Union(_1444_v124));
          let _1448_v128;
          _1448_v128 = _dafny.Map.Empty.slice().updateUnsafe(_1277_v1,_1277_v1);
          let _index222 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_1445_v125).length));
          (_1445_v125)[_index222] = (((_1447_v127).contains(((_1277_v1) ? ((((_1448_v128).contains(_1277_v1)) ? ((_1448_v128).get(_1277_v1)) : (false))) : (!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]))))) ? ((_1447_v127).get(((_1277_v1) ? ((((_1448_v128).contains(_1277_v1)) ? ((_1448_v128).get(_1277_v1)) : (false))) : (!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]))))) : ((_1444_v124).Difference(_1444_v124)));
          let _1449_v129;
          let _nw223 = new _module.C0();
          _nw223.__ctor(_dafny.Seq.Create(_module.__default.abs(new BigNumber(445)), ((_1450_v101) => function (_1451_i8) {
            return _1450_v101;
          })(_1419_v101)));
          _1449_v129 = _nw223;
          let _index223 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length));
          (_1422_v104)[_index223] = new BigNumber((_dafny.MultiSet.fromElements(((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) && ((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]), (_this).fm5(_1424_v106, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], _1421_v103, false, globalState))).cardinality());
          let _1452_v130;
          _1452_v130 = _dafny.MultiSet.fromElements(_1424_v106, _1424_v106);
          let _1453_v131;
          _1453_v131 = _module.D4.create_DC9(_1452_v130, _1419_v101, _1424_v106, _dafny.Set.fromElements(_1449_v129), _1424_v106);
          let _1454_v132;
          let _nw224 = new _module.C0();
          _nw224.__ctor(_dafny.Seq.Concat(_1423_v105, _dafny.Seq.update(_1449_v129.f8, _module.__default.safeIndex((_1422_v104)[_module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length))], new BigNumber((_1449_v129.f8).length)), (_1453_v131).dtor_cf15)));
          _1454_v132 = _nw224;
          let _index224 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length));
          (_1422_v104)[_index224] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]),p0)).length);
        }
        let _index225 = _module.__default.safeIndex(new BigNumber(758), new BigNumber((_1422_v104).length));
        (_1422_v104)[_index225] = _1424_v106;
        _1424_v106 = p0;
      } else {
        r0 = !(_1277_v1);
        _1277_v1 = (((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) ? ((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) : (!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])));
        let _1455_v133;
        _1455_v133 = new BigNumber(58);
        _1455_v133 = p0;
        let _1456_v134;
        _1456_v134 = _dafny.Seq.UnicodeFromString("gv");
        let _1457_v135;
        let _nw225 = new _module.C0();
        _nw225.__ctor(_dafny.Seq.Concat(_1456_v134, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-868)), ((_1458_v101) => function (_1459_i9) {
          return _1458_v101;
        })(_1419_v101))));
        _1457_v135 = _nw225;
        let _index226 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        (_1275_v0)[_index226] = _1277_v1;
      }
      let _1460_v136;
      _1460_v136 = _dafny.Seq.UnicodeFromString("ng");
      let _1461_v137;
      _1461_v137 = _dafny.Seq.of(_1277_v1, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
      let _1462_v138;
      _1462_v138 = _dafny.Seq.of(_1461_v137, _1461_v137);
      let _1463_v139;
      _1463_v139 = _dafny.Seq.of(_1461_v137, (_1462_v138)[_module.__default.safeIndex(p0, new BigNumber((_1462_v138).length))]);
      let _1464_v140;
      _1464_v140 = _module.D6.create_DC15(p0, new BigNumber((_1460_v136).length), _dafny.areEqual(_dafny.Seq.of(_module.__default.fm23(_1277_v1, globalState), _1461_v137), _1462_v138));
      let _source16 = _1464_v140;
      if (_source16.is_DC15) {
        let _1465___mcc_h15 = (_source16).cf30;
        let _1466___mcc_h16 = (_source16).cf31;
        let _1467___mcc_h17 = (_source16).cf32;
        let _1468_cf32 = _1467___mcc_h17;
        let _1469_cf31 = _1466___mcc_h16;
        let _1470_cf30 = _1465___mcc_h15;
        _1420_v102 = _module.D0.create_DC1((_1469_cf31).minus(new BigNumber(894)));
        _1461_v137 = _1461_v137;
        let _index227 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        (_1275_v0)[_index227] = _1468_cf32;
        let _1471_v141;
        let _nw226 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
        _1471_v141 = _nw226;
        let _index228 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_1471_v141).length));
        (_1471_v141)[_index228] = (_dafny.ZERO).minus(_1470_cf30);
        let _index229 = _module.__default.safeIndex(new BigNumber(541), new BigNumber((_1471_v141).length));
        (_1471_v141)[_index229] = new BigNumber(335);
      } else if (_source16.is_DC14) {
        let _1472___mcc_h18 = (_source16).cf29;
        let _1473_cf29 = _1472___mcc_h18;
        let _1474_v142;
        _1474_v142 = new BigNumber(308);
        _1474_v142 = (_dafny.ZERO).minus(p0);
        let _1475_v143;
        let _nw227 = new _module.C0();
        _nw227.__ctor(_1473_cf29);
        _1475_v143 = _nw227;
        let _index230 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        let _rhs156 = _1475_v143;
        let _rhs157 = _1474_v142;
        let _rhs158 = (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))];
        let _lhs92 = _1275_v0;
        let _lhs93 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        _1475_v143 = _rhs156;
        _1474_v142 = _rhs157;
        _lhs92[_lhs93] = _rhs158;
        let _rhs159 = _1464_v140;
        let _rhs160 = (_1474_v142).plus(new BigNumber((_1418_v100).length));
        let _rhs161 = _1475_v143;
        _1464_v140 = _rhs159;
        _1474_v142 = _rhs160;
        _1475_v143 = _rhs161;
        let _1476_v144;
        let _nw228 = Array((new BigNumber(8)).toNumber());
        _1476_v144 = _nw228;
        let _1477_v145;
        _1477_v145 = _dafny.MultiSet.fromElements(_1277_v1);
        let _1478_v146;
        _1478_v146 = _module.D9.create_DC21(_1477_v145);
        let _pat_let_tv38 = _1461_v137;
        let _1479_v147;
        let _nw229 = Array((new BigNumber(23)).toNumber());
        _nw229[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_1277_v1,p0)).length);
        _nw229[(_dafny.ONE).toNumber()] = (_1420_v102).dtor_cf2;
        _nw229[(new BigNumber(2)).toNumber()] = p0;
        _nw229[(new BigNumber(3)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw229[(new BigNumber(4)).toNumber()] = p0;
        _nw229[(new BigNumber(5)).toNumber()] = _1474_v142;
        _nw229[(new BigNumber(6)).toNumber()] = p0;
        _nw229[(new BigNumber(7)).toNumber()] = p0;
        _nw229[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw229[(new BigNumber(9)).toNumber()] = p0;
        _nw229[(new BigNumber(10)).toNumber()] = _1474_v142;
        _nw229[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((function (_pat_let34_0) {
          return function (_1480_dt__update__tmp_h0) {
            return function (_pat_let35_0) {
              return function (_1481_dt__update_hcf37_h0) {
                return _module.D9.create_DC21(_1481_dt__update_hcf37_h0);
              }(_pat_let35_0);
            }(_dafny.MultiSet.FromArray(_pat_let_tv38));
          }(_pat_let34_0);
        }(_1478_v146)).dtor_cf37).cardinality()));
        _nw229[(new BigNumber(12)).toNumber()] = _1474_v142;
        _nw229[(new BigNumber(13)).toNumber()] = p0;
        _nw229[(new BigNumber(14)).toNumber()] = _1474_v142;
        _nw229[(new BigNumber(15)).toNumber()] = _1474_v142;
        _nw229[(new BigNumber(16)).toNumber()] = p0;
        _nw229[(new BigNumber(17)).toNumber()] = _1474_v142;
        _nw229[(new BigNumber(18)).toNumber()] = p0;
        _nw229[(new BigNumber(19)).toNumber()] = new BigNumber(70);
        _nw229[(new BigNumber(20)).toNumber()] = p0;
        _nw229[(new BigNumber(21)).toNumber()] = p0;
        _nw229[(new BigNumber(22)).toNumber()] = _1474_v142;
        _1479_v147 = _nw229;
        let _1482_v148;
        let _nw230 = Array((new BigNumber(8)).toNumber());
        _nw230[(_dafny.ZERO).toNumber()] = _1479_v147;
        _nw230[(_dafny.ONE).toNumber()] = _1479_v147;
        _nw230[(new BigNumber(2)).toNumber()] = _1479_v147;
        _nw230[(new BigNumber(3)).toNumber()] = _1479_v147;
        _nw230[(new BigNumber(4)).toNumber()] = _1479_v147;
        _nw230[(new BigNumber(5)).toNumber()] = _1479_v147;
        _nw230[(new BigNumber(6)).toNumber()] = _1479_v147;
        _nw230[(new BigNumber(7)).toNumber()] = _1479_v147;
        _1482_v148 = _nw230;
        let _1483_v149;
        _1483_v149 = _dafny.Map.Empty.slice().updateUnsafe(_1476_v144,_1482_v148);
        let _1484_v150;
        _1484_v150 = _dafny.Seq.of(_1476_v144);
        let _1485_v151;
        _1485_v151 = _dafny.Seq.of((((_1483_v149).contains((_1484_v150)[_module.__default.safeIndex(_1474_v142, new BigNumber((_1484_v150).length))])) ? ((_1483_v149).get((_1484_v150)[_module.__default.safeIndex(_1474_v142, new BigNumber((_1484_v150).length))])) : (_1482_v148)));
        let _1486_v152;
        let _nw231 = Array((new BigNumber(28)).toNumber());
        _nw231[(_dafny.ZERO).toNumber()] = (_1485_v151)[_module.__default.safeIndex(p0, new BigNumber((_1485_v151).length))];
        _nw231[(_dafny.ONE).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(2)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(3)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(4)).toNumber()] = (((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) ? (_1482_v148) : (_1482_v148));
        _nw231[(new BigNumber(5)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(6)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(7)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(8)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(9)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(10)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(11)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(12)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(13)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(14)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(15)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(16)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(17)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(18)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(19)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(20)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(21)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(22)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(23)).toNumber()] = ((true) ? (_1482_v148) : (_1482_v148));
        _nw231[(new BigNumber(24)).toNumber()] = ((_1277_v1) ? (_1482_v148) : (_1482_v148));
        _nw231[(new BigNumber(25)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(26)).toNumber()] = _1482_v148;
        _nw231[(new BigNumber(27)).toNumber()] = _1482_v148;
        _1486_v152 = _nw231;
        let _index231 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_1486_v152).length));
        (_1486_v152)[_index231] = _1482_v148;
        let _index232 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1479_v147).length));
        (_1479_v147)[_index232] = new BigNumber(((_1418_v100).Merge(_1418_v100)).length);
        let _1487_v153;
        _1487_v153 = _dafny.MultiSet.fromElements(_module.__default.fm20(true, globalState));
        let _1488_v154;
        _1488_v154 = _dafny.MultiSet.fromElements(_1487_v153, _1487_v153);
        let _index233 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_1486_v152).length));
        let _index234 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        let _index235 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1479_v147).length));
        let _rhs162 = _1482_v148;
        let _rhs163 = (_dafny.MultiSet.fromElements(_1487_v153, _1487_v153, _module.__default.fm24((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], globalState))).IsSubsetOf(_1488_v154);
        let _rhs164 = (((_this).fm5(p0, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], _1421_v103, _1277_v1, globalState)) ? (new BigNumber(502)) : (p0));
        let _rhs165 = _1474_v142;
        let _rhs166 = _1277_v1;
        let _lhs94 = _1486_v152;
        let _lhs95 = _module.__default.safeIndex(new BigNumber(766), new BigNumber((_1486_v152).length));
        let _lhs96 = _1275_v0;
        let _lhs97 = _module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length));
        let _lhs98 = _1479_v147;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(334), new BigNumber((_1479_v147).length));
        _lhs94[_lhs95] = _rhs162;
        _lhs96[_lhs97] = _rhs163;
        _1474_v142 = _rhs164;
        _lhs98[_lhs99] = _rhs165;
        r0 = _rhs166;
      } else {
        let _1489___mcc_h19 = (_source16).cf33;
        let _1490_cf33 = _1489___mcc_h19;
        let _1491_v155;
        _1491_v155 = _dafny.Map.Empty.slice().updateUnsafe(p0,!((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]));
        let _1492_v156;
        _1492_v156 = _dafny.Seq.of(p0, p0, p0, new BigNumber((_1491_v155).length));
        let _1493_v157;
        _1493_v157 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_1492_v156),_module.__default.fm20((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], globalState));
        let _1494_v158;
        _1494_v158 = _dafny.Set.fromElements((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], (_1461_v137)[_module.__default.safeIndex(new BigNumber(994), new BigNumber((_1461_v137).length))], _1277_v1);
        let _1495_v159;
        let _nw232 = Array((new BigNumber(29)).toNumber());
        _nw232[(_dafny.ZERO).toNumber()] = p0;
        _nw232[(_dafny.ONE).toNumber()] = (p0).multipliedBy(p0);
        _nw232[(new BigNumber(2)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw232[(new BigNumber(3)).toNumber()] = p0;
        _nw232[(new BigNumber(4)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
        _nw232[(new BigNumber(5)).toNumber()] = p0;
        _nw232[(new BigNumber(6)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw232[(new BigNumber(7)).toNumber()] = _module.__default.fm20(_1277_v1, globalState);
        _nw232[(new BigNumber(8)).toNumber()] = p0;
        _nw232[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.of(_1418_v100, _1418_v100)).length);
        _nw232[(new BigNumber(10)).toNumber()] = _module.__default.fm20(!(false), globalState);
        _nw232[(new BigNumber(11)).toNumber()] = p0;
        _nw232[(new BigNumber(12)).toNumber()] = p0;
        _nw232[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(_1277_v1, false, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))])).length), new BigNumber((_1493_v157).length));
        _nw232[(new BigNumber(14)).toNumber()] = p0;
        _nw232[(new BigNumber(15)).toNumber()] = new BigNumber(((_1494_v158).Union(_1494_v158)).length);
        _nw232[(new BigNumber(16)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1461_v137).length), p0);
        _nw232[(new BigNumber(17)).toNumber()] = p0;
        _nw232[(new BigNumber(18)).toNumber()] = p0;
        _nw232[(new BigNumber(19)).toNumber()] = (new BigNumber(199)).multipliedBy(p0);
        _nw232[(new BigNumber(20)).toNumber()] = ((_dafny.ZERO).minus(p0)).plus(p0);
        _nw232[(new BigNumber(21)).toNumber()] = p0;
        _nw232[(new BigNumber(22)).toNumber()] = p0;
        _nw232[(new BigNumber(23)).toNumber()] = (((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) ? (new BigNumber((_1461_v137).length)) : (p0));
        _nw232[(new BigNumber(24)).toNumber()] = new BigNumber(-208);
        _nw232[(new BigNumber(25)).toNumber()] = p0;
        _nw232[(new BigNumber(26)).toNumber()] = _module.__default.fm20((_this).fm5(p0, _1277_v1, _1421_v103, _1277_v1, globalState), globalState);
        _nw232[(new BigNumber(27)).toNumber()] = new BigNumber(-406);
        _nw232[(new BigNumber(28)).toNumber()] = p0;
        _1495_v159 = _nw232;
        let _index236 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1495_v159).length));
        (_1495_v159)[_index236] = (_dafny.ZERO).minus(p0);
        let _index237 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_1495_v159).length));
        (_1495_v159)[_index237] = p0;
        let _1496_v160;
        _1496_v160 = _dafny.Seq.of(_1460_v136);
        let _index238 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1495_v159).length));
        let _index239 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_1495_v159).length));
        let _rhs167 = (new BigNumber((_dafny.Seq.Concat((_1496_v160)[_module.__default.safeIndex(p0, new BigNumber((_1496_v160).length))], _dafny.Seq.UnicodeFromString("wqfk"))).length)).plus(p0);
        let _rhs168 = p0;
        let _lhs100 = _1495_v159;
        let _lhs101 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1495_v159).length));
        let _lhs102 = _1495_v159;
        let _lhs103 = _module.__default.safeIndex(new BigNumber(112), new BigNumber((_1495_v159).length));
        _lhs100[_lhs101] = _rhs167;
        _lhs102[_lhs103] = _rhs168;
        let _1497_v161;
        _1497_v161 = _dafny.MultiSet.fromElements((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], _1277_v1);
        let _1498_v162;
        _1498_v162 = _dafny.Map.Empty.slice().updateUnsafe(_1460_v136,_1497_v161);
        let _1499_v163;
        _1499_v163 = _dafny.MultiSet.fromElements(_1498_v162, _1498_v162);
        let _index240 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1495_v159).length));
        (_1495_v159)[_index240] = new BigNumber((_dafny.Seq.of(new BigNumber((_1497_v161).cardinality()), ((((_1499_v163).contains(_dafny.Map.Empty.slice().updateUnsafe(_1460_v136,_1497_v161))) ? ((_1499_v163).get(_dafny.Map.Empty.slice().updateUnsafe(_1460_v136,_1497_v161))) : (p0))).plus((_1495_v159)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_1495_v159).length))]), p0)).length);
        let _index241 = _module.__default.safeIndex(new BigNumber(691), new BigNumber((_1495_v159).length));
        (_1495_v159)[_index241] = new BigNumber(468);
        let _1500_v164;
        _1500_v164 = _dafny.Map.Empty.slice().updateUnsafe(_1277_v1,_1460_v136);
        let _1501_v165;
        let _1502_v166;
        let _out45;
        let _out46;
        let _outcollector14 = (_this).m10((new BigNumber((_1500_v164).length)).plus(p0), (_1495_v159)[_module.__default.safeIndex(new BigNumber(691), new BigNumber((_1495_v159).length))], !((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]), _1275_v0, globalState);
        _out45 = _outcollector14[0];
        _out46 = _outcollector14[1];
        _1501_v165 = _out45;
        _1502_v166 = _out46;
      }
      let _1503_v167;
      let _init33 = function (_1504_i10) {
        return _dafny.MultiSet.fromElements(true);
      };
      let _nw233 = Array((new BigNumber(20)).toNumber());
      for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw233.length); _i0_33++) {
        _nw233[_i0_33] = _init33(new BigNumber(_i0_33));
      }
      _1503_v167 = _nw233;
      let _nw234 = Array((new BigNumber(20)).toNumber()).fill(_dafny.MultiSet.Empty);
      _1503_v167 = _nw234;
      r0 = !((_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]) || ((_this).fm5(p0, (_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))], _1421_v103, false, globalState));
      let _1505_v168;
      _1505_v168 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1275_v0)[_module.__default.safeIndex(new BigNumber(663), new BigNumber((_1275_v0).length))]);
      r1 = (_1505_v168).equals(_1505_v168);
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _1506_v0;
      _1506_v0 = false;
      let _1507_v1;
      _1507_v1 = _module.D6.create_DC14(_dafny.Seq.Create(_module.__default.abs(new BigNumber(491)), function (_1508_i0) {
  return new _dafny.CodePoint('g'.codePointAt(0));
}));
      let _1509_v2;
      let _nw235 = new _module.C0();
      _nw235.__ctor(((_1506_v0) ? ((_1507_v1).dtor_cf29) : (_dafny.Seq.UnicodeFromString("gpmds"))));
      _1509_v2 = _nw235;
      let _1510_v3;
      _1510_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(751)), ((_1511_p0) => function (_1512_i1) {
        return _1511_p0;
      })(p0))).length),_1506_v0);
      let _1513_v4;
      _1513_v4 = _dafny.MultiSet.fromElements(_1506_v0);
      let _1514_v5;
      _1514_v5 = _dafny.MultiSet.fromElements(p0, p0, p0);
      let _1515_v6;
      _1515_v6 = _dafny.Set.fromElements(p0, p0);
      let _1516_v7;
      _1516_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1514_v5,new BigNumber((_1515_v6).length));
      let _1517_v8;
      _1517_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1516_v7).length),new BigNumber(828));
      let _1518_v9;
      _1518_v9 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_1517_v8).length));
      let _1519_v10;
      let _nw236 = Array((new BigNumber(20)).toNumber());
      _nw236[(_dafny.ZERO).toNumber()] = p0;
      _nw236[(_dafny.ONE).toNumber()] = new BigNumber((_1509_v2.f8).length);
      _nw236[(new BigNumber(2)).toNumber()] = new BigNumber(928);
      _nw236[(new BigNumber(3)).toNumber()] = p0;
      _nw236[(new BigNumber(4)).toNumber()] = p0;
      _nw236[(new BigNumber(5)).toNumber()] = p0;
      _nw236[(new BigNumber(6)).toNumber()] = p0;
      _nw236[(new BigNumber(7)).toNumber()] = p0;
      _nw236[(new BigNumber(8)).toNumber()] = p0;
      _nw236[(new BigNumber(9)).toNumber()] = new BigNumber((_1513_v4).cardinality());
      _nw236[(new BigNumber(10)).toNumber()] = p0;
      _nw236[(new BigNumber(11)).toNumber()] = p0;
      _nw236[(new BigNumber(12)).toNumber()] = new BigNumber((_1518_v9).length);
      _nw236[(new BigNumber(13)).toNumber()] = p0;
      _nw236[(new BigNumber(14)).toNumber()] = p0;
      _nw236[(new BigNumber(15)).toNumber()] = p0;
      _nw236[(new BigNumber(16)).toNumber()] = p0;
      _nw236[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw236[(new BigNumber(18)).toNumber()] = p0;
      _nw236[(new BigNumber(19)).toNumber()] = p0;
      _1519_v10 = _nw236;
      let _1520_v11;
      _1520_v11 = _dafny.Map.Empty.slice().updateUnsafe(_1519_v10,p0);
      let _1521_v12;
      _1521_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1520_v11,_1519_v10);
      let _1522_v13;
      let _nw237 = Array((new BigNumber(15)).toNumber());
      _nw237[(_dafny.ZERO).toNumber()] = _1519_v10;
      _nw237[(_dafny.ONE).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(2)).toNumber()] = (((_1521_v12).contains(_1520_v11)) ? ((_1521_v12).get(_1520_v11)) : (_1519_v10));
      _nw237[(new BigNumber(3)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(4)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(5)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(6)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(7)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(8)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(9)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(10)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(11)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(12)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(13)).toNumber()] = _1519_v10;
      _nw237[(new BigNumber(14)).toNumber()] = _1519_v10;
      _1522_v13 = _nw237;
      let _1523_v14;
      _1523_v14 = _module.D2.create_DC4(p0, _1506_v0, _1522_v13);
      let _1524_v15;
      _1524_v15 = _dafny.Seq.of((((_1518_v9).contains(_1506_v0)) ? ((_1518_v9).get(_1506_v0)) : (p0)), new BigNumber((_dafny.Seq.UnicodeFromString("uotirlywr")).length));
      let _1525_v16;
      _1525_v16 = _1519_v10;
      let _1526_v17;
      _1526_v17 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1525_v16));
      let _1527_v18;
      _1527_v18 = _dafny.Seq.of(_1509_v2);
      let _1528_v19;
      let _nw238 = Array((new BigNumber(19)).toNumber());
      _nw238[(_dafny.ZERO).toNumber()] = new BigNumber((_1509_v2.f8).length);
      _nw238[(_dafny.ONE).toNumber()] = p0;
      _nw238[(new BigNumber(2)).toNumber()] = new BigNumber((_1510_v3).length);
      _nw238[(new BigNumber(3)).toNumber()] = (p0).plus(p0);
      _nw238[(new BigNumber(4)).toNumber()] = (p0).plus(p0);
      _nw238[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(p0);
      _nw238[(new BigNumber(6)).toNumber()] = p0;
      _nw238[(new BigNumber(7)).toNumber()] = (p0).multipliedBy(p0);
      _nw238[(new BigNumber(8)).toNumber()] = new BigNumber((_dafny.Seq.of(_module.__default.fm20(_1506_v0, globalState))).length);
      _nw238[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus((new BigNumber(85)).plus(new BigNumber((_1509_v2.f8).length)));
      _nw238[(new BigNumber(10)).toNumber()] = p0;
      _nw238[(new BigNumber(11)).toNumber()] = p0;
      _nw238[(new BigNumber(12)).toNumber()] = (_1523_v14).dtor_cf5;
      _nw238[(new BigNumber(13)).toNumber()] = p0;
      _nw238[(new BigNumber(14)).toNumber()] = (_1524_v15)[_module.__default.safeIndex(p0, new BigNumber((_1524_v15).length))];
      _nw238[(new BigNumber(15)).toNumber()] = p0;
      _nw238[(new BigNumber(16)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1526_v17).length));
      _nw238[(new BigNumber(17)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1527_v18).length), p0);
      _nw238[(new BigNumber(18)).toNumber()] = p0;
      _1528_v19 = _nw238;
      let _1529_v20;
      _1529_v20 = _dafny.MultiSet.fromElements((_1507_v1).dtor_cf29, _1509_v2.f8);
      let _nw239 = Array((new BigNumber(12)).toNumber());
      _nw239[(_dafny.ZERO).toNumber()] = (_1523_v14).dtor_cf5;
      _nw239[(_dafny.ONE).toNumber()] = (_1524_v15)[_module.__default.safeIndex(p0, new BigNumber((_1524_v15).length))];
      _nw239[(new BigNumber(2)).toNumber()] = p0;
      _nw239[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt((_dafny.ZERO).minus(p0), p0);
      _nw239[(new BigNumber(4)).toNumber()] = p0;
      _nw239[(new BigNumber(5)).toNumber()] = (p0).plus(new BigNumber((_1529_v20).cardinality()));
      _nw239[(new BigNumber(6)).toNumber()] = p0;
      _nw239[(new BigNumber(7)).toNumber()] = (p0).plus(new BigNumber(75));
      _nw239[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((p0).plus(p0));
      _nw239[(new BigNumber(9)).toNumber()] = (p0).plus(p0);
      _nw239[(new BigNumber(10)).toNumber()] = p0;
      _nw239[(new BigNumber(11)).toNumber()] = (p0).minus(p0);
      _1528_v19 = _nw239;
      let _index242 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length));
      (_1519_v10)[_index242] = _module.__default.safeModuloInt(new BigNumber((_1509_v2.f8).length), p0);
      let _index243 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length));
      (_1519_v10)[_index243] = p0;
      let _1530_i2;
      _1530_i2 = _dafny.ZERO;
      L4: {
        while (!(false)) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1530_i2)) {
              break L4;
            }
            _1530_i2 = (_1530_i2).plus(_dafny.ONE);
            let _1531_v21;
            _1531_v21 = _module.D4.create_DC10(p0, p0, (_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))]);
            let _1532_v22;
            _1532_v22 = new _dafny.CodePoint('t'.codePointAt(0));
            let _1533_v23;
            _1533_v23 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm25(_1531_v21, _1509_v2.f8, globalState),_1532_v22);
            let _1534_v24;
            _1534_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1532_v22,_1533_v23);
            let _1535_v25;
            _1535_v25 = _dafny.Seq.of(_1513_v4);
            let _1536_v27;
            _1536_v27 = _dafny.Set.fromElements(_1509_v2.f8);
            let _1537_v28;
            _1537_v28 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_module.D2.create_DC3(_1536_v27));
            let _1538_v29;
            _1538_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1509_v2.f8,_1537_v28);
            let _1539_v30;
            _1539_v30 = _dafny.Set.fromElements(_1506_v0);
            let _1540_v31;
            _1540_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(_1506_v0));
            let _1541_v32;
            _1541_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1509_v2,p0);
            let _1542_v33;
            _1542_v33 = _dafny.Seq.of(_1506_v0, _1506_v0);
            let _1543_v34;
            _1543_v34 = _module.D0.create_DC1((_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))]);
            let _1544_v37;
            _1544_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1509_v2.f8,_1506_v0);
            let _1545_v40;
            _1545_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm25(_1531_v21, _1509_v2.f8, globalState),(_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))]);
            let _1546_v41;
            _1546_v41 = _module.D12.create_DC27(_1539_v30);
            let _1547_v42;
            _1547_v42 = _module.D11.create_DC24(_module.__default.fm26((_1546_v41).dtor_cf45, _1542_v33, _1543_v34, _1532_v22, globalState));
            let _1548_v43;
            let _nw240 = Array((new BigNumber(29)).toNumber());
            _nw240[(_dafny.ZERO).toNumber()] = _1533_v23;
            _nw240[(_dafny.ONE).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1509_v2.f8,_1532_v22);
            _nw240[(new BigNumber(2)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(3)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(4)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(5)).toNumber()] = (_1533_v23).update(_1509_v2.f8, _1532_v22);
            _nw240[(new BigNumber(6)).toNumber()] = (((_1534_v24).contains(_1532_v22)) ? ((_1534_v24).get(_1532_v22)) : (_1533_v23));
            _nw240[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_1509_v2).fm15(_dafny.Seq.of(new BigNumber((_1535_v25).length)), _1506_v0, globalState),new _dafny.CodePoint('s'.codePointAt(0)));
            _nw240[(new BigNumber(8)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(9)).toNumber()] = function () {
              let _coll53 = new _dafny.Map();
              for (const _compr_53 of (_1538_v29).Keys.Elements) {
                let _1549_v26 = _compr_53;
                if ((_1538_v29).contains(_1549_v26)) {
                  _coll53.push([_1549_v26,new _dafny.CodePoint('t'.codePointAt(0))]);
                }
              }
              return _coll53;
            }();
            _nw240[(new BigNumber(10)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(11)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(12)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(13)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(14)).toNumber()] = _module.__default.fm26(_1539_v30, (((_1540_v31).contains((((_1541_v32).contains(_1509_v2)) ? ((_1541_v32).get(_1509_v2)) : ((_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))])))) ? ((_1540_v31).get((((_1541_v32).contains(_1509_v2)) ? ((_1541_v32).get(_1509_v2)) : ((_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))])))) : (_1542_v33)), _1543_v34, _1532_v22, globalState);
            _nw240[(new BigNumber(15)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_1509_v2.f8,(_1509_v2).fm14(globalState))).Merge((function () {
              let _coll54 = new _dafny.Map();
              for (const _compr_54 of (_1529_v20).Elements) {
                let _1550_v35 = _compr_54;
                if ((_1529_v20).contains(_1550_v35)) {
                  _coll54.push([_1550_v35,_1532_v22]);
                }
              }
              return _coll54;
            }()).update(_dafny.Seq.UnicodeFromString("u"), _1532_v22));
            _nw240[(new BigNumber(16)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(17)).toNumber()] = (_1533_v23).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1509_v2.f8,(_1509_v2.f8)[_module.__default.safeIndex((_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))], new BigNumber((_1509_v2.f8).length))]));
            _nw240[(new BigNumber(18)).toNumber()] = function () {
              let _coll55 = new _dafny.Map();
              for (const _compr_55 of (_1544_v37).Keys.Elements) {
                let _1551_v36 = _compr_55;
                if ((_1544_v37).contains(_1551_v36)) {
                  _coll55.push([_1551_v36,_1532_v22]);
                }
              }
              return _coll55;
            }();
            _nw240[(new BigNumber(19)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(20)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(21)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(22)).toNumber()] = function () {
              let _coll56 = new _dafny.Map();
              for (const _compr_56 of (function () {
                let _coll57 = new _dafny.Map();
                for (const _compr_57 of (_1545_v40).Keys.Elements) {
                  let _1552_v39 = _compr_57;
                  if ((_1545_v40).contains(_1552_v39)) {
                    _coll57.push([_1552_v39,_1506_v0]);
                  }
                }
                return _coll57;
              }()).Keys.Elements) {
                let _1553_v38 = _compr_56;
                if ((function () {
                  let _coll58 = new _dafny.Map();
                  for (const _compr_58 of (_1545_v40).Keys.Elements) {
                    let _1552_v39 = _compr_58;
                    if ((_1545_v40).contains(_1552_v39)) {
                      _coll58.push([_1552_v39,_1506_v0]);
                    }
                  }
                  return _coll58;
                }()).contains(_1553_v38)) {
                  _coll56.push([_1553_v38,_1532_v22]);
                }
              }
              return _coll56;
            }();
            _nw240[(new BigNumber(23)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(24)).toNumber()] = _1533_v23;
            _nw240[(new BigNumber(25)).toNumber()] = (_1547_v42).dtor_cf40;
            _nw240[(new BigNumber(26)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(890)), ((_1554_v22) => function (_1555_i3) {
              return _1554_v22;
            })(_1532_v22)),_1532_v22);
            _nw240[(new BigNumber(27)).toNumber()] = (_1533_v23).Merge(_1533_v23);
            _nw240[(new BigNumber(28)).toNumber()] = _1533_v23;
            _1548_v43 = _nw240;
            let _index244 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_1548_v43).length));
            (_1548_v43)[_index244] = _1533_v23;
            let _index245 = _module.__default.safeIndex(new BigNumber(622), new BigNumber((_1548_v43).length));
            (_1548_v43)[_index245] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(136)), ((_1556_v22) => function (_1557_i4) {
              return _1556_v22;
            })(_1532_v22)),_1532_v22);
            let _1558_v44;
            let _nw241 = Array((new BigNumber(21)).toNumber()).fill(_dafny.Map.Empty);
            _1558_v44 = _nw241;
            let _1559_v45;
            _1559_v45 = _dafny.Map.Empty.slice().updateUnsafe(_1532_v22,new BigNumber((_1515_v6).length));
            let _1560_v46;
            _1560_v46 = _dafny.MultiSet.fromElements(_1559_v45, _1559_v45);
            let _1561_v47;
            let _nw242 = new _module.C2();
            _nw242.__ctor();
            _1561_v47 = _nw242;
            let _1562_v48;
            _1562_v48 = _dafny.Map.Empty.slice().updateUnsafe(true,_1561_v47);
            let _1563_v49;
            _1563_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1560_v46,_1562_v48);
            let _index246 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_1558_v44).length));
            (_1558_v44)[_index246] = ((_1506_v0) ? (_1563_v49) : (_1563_v49));
            let _1564_v50;
            _1564_v50 = _dafny.Seq.of(_1509_v2.f8, _1509_v2.f8);
            let _1565_v51;
            _1565_v51 = _dafny.Map.Empty.slice().updateUnsafe(_1506_v0,_1562_v48);
            let _index247 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_1558_v44).length));
            let _index248 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length));
            let _rhs169 = _1506_v0;
            let _rhs170 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.UnicodeFromString("xone"), (_1564_v50)[_module.__default.safeIndex(p0, new BigNumber((_1564_v50).length))]);
            let _rhs171 = (((_1506_v0) ? (_1563_v49) : (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_1559_v45),_1562_v48)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1560_v46,(((_1565_v51).contains(_1506_v0)) ? ((_1565_v51).get(_1506_v0)) : (_1562_v48))));
            let _rhs172 = (_dafny.ZERO).minus(((!(_1510_v3).equals(_1510_v3)) ? ((p0).plus(_module.__default.fm20(_1506_v0, globalState))) : (((_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))]).plus(p0))));
            let _lhs104 = _1558_v44;
            let _lhs105 = _module.__default.safeIndex(new BigNumber(735), new BigNumber((_1558_v44).length));
            let _lhs106 = _1519_v10;
            let _lhs107 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length));
            _1506_v0 = _rhs169;
            _1506_v0 = _rhs170;
            _lhs104[_lhs105] = _rhs171;
            _lhs106[_lhs107] = _rhs172;
            let _1566_v52;
            _1566_v52 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('c'.codePointAt(0)),_1543_v34);
            let _1567_v53;
            _1567_v53 = _module.D2.create_DC5((_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))], _1524_v15, false);
            if ((_this).fm5((_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))], !(((_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))]).isLessThanOrEqualTo(p0)), _1566_v52, (_1567_v53).dtor_cf10, globalState)) {
              let _1568_v54;
              let _init34 = function (_1569_i5) {
                return true;
              };
              let _nw243 = Array((new BigNumber(12)).toNumber());
              for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw243.length); _i0_34++) {
                _nw243[_i0_34] = _init34(new BigNumber(_i0_34));
              }
              _1568_v54 = _nw243;
              _1568_v54 = _1568_v54;
              let _index249 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_1522_v13).length));
              (_1522_v13)[_index249] = _1528_v19;
              let _index250 = _module.__default.safeIndex(new BigNumber(975), new BigNumber((_1522_v13).length));
              (_1522_v13)[_index250] = _1528_v19;
              let _1570_v55;
              let _1571_v56;
              let _out47;
              let _out48;
              let _outcollector15 = (_1561_v47).m1((_dafny.ZERO).minus(_module.__default.safeModuloInt(p0, new BigNumber(5))), globalState);
              _out47 = _outcollector15[0];
              _out48 = _outcollector15[1];
              _1570_v55 = _out47;
              _1571_v56 = _out48;
              let _1572_v57;
              _1572_v57 = _dafny.Map.Empty.slice().updateUnsafe(_1568_v54,true);
              _1571_v56 = (_1572_v57).equals(_1572_v57);
              let _index251 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length));
              (_1519_v10)[_index251] = (_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))];
            } else {
              let _1573_v58;
              _1573_v58 = _module.D7.create_DC17(_1542_v33);
              let _index252 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length));
              (_1519_v10)[_index252] = new BigNumber(((_1573_v58).dtor_cf34).length);
              (_1509_v2).f8 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(324)), ((_1574_v22) => function (_1575_i6) {
                return _1574_v22;
              })(_1532_v22));
              let _1576_v59;
              _1576_v59 = _module.D16.create_DC44((_1519_v10)[_module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length))]);
              let _pat_let_tv39 = _1524_v15;
              let _pat_let_tv40 = _1518_v9;
              let _pat_let_tv41 = _1524_v15;
              _1576_v59 = function (_pat_let36_0) {
                return function (_1577_dt__update__tmp_h0) {
                  return function (_pat_let37_0) {
                    return function (_1578_dt__update_hcf82_h0) {
                      return _module.D16.create_DC44(_1578_dt__update_hcf82_h0);
                    }(_pat_let37_0);
                  }((_pat_let_tv39)[_module.__default.safeIndex(new BigNumber((_pat_let_tv40).length), new BigNumber((_pat_let_tv41).length))]);
                }(_pat_let36_0);
              }(_1576_v59);
              _1544_v37 = (_1544_v37).update(_1509_v2.f8, _1506_v0);
              _1514_v5 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(311)), function (_1579_i7) {
                return new BigNumber(-76);
              }), _1524_v15))).Intersect((_1514_v5).Union(_1514_v5));
            }
            let _1580_v60;
            _1580_v60 = _module.D9.create_DC21((_module.D9.create_DC21(_1513_v4)).dtor_cf37);
            let _1581_v61;
            _1581_v61 = _dafny.Seq.of(_1580_v60, _1580_v60, _1580_v60);
            _1581_v61 = _dafny.Seq.Concat(_1581_v61, _1581_v61);
          }
        }
      }
      let _1582_v62;
      let _nw244 = new _module.C2();
      _nw244.__ctor();
      _1582_v62 = _nw244;
      let _index253 = _module.__default.safeIndex(new BigNumber(655), new BigNumber((_1519_v10).length));
      (_1519_v10)[_index253] = p0;
      let _1583_v63;
      _1583_v63 = _dafny.Map.Empty.slice().updateUnsafe(!(_1506_v0),_1514_v5);
      r0 = _1583_v63;
      return r0;
    }
    m9(globalState) {
      let _this = this;
      let r0 = [];
      let r1 = [];
      let r2 = false;
      let _1584_v0;
      _1584_v0 = new _dafny.CodePoint('x'.codePointAt(0));
      let _1585_v1;
      _1585_v1 = _dafny.Seq.UnicodeFromString("jvypfkdr");
      let _1586_v2;
      _1586_v2 = false;
      let _1587_v3;
      _1587_v3 = new BigNumber(-189);
      let _1588_v4;
      _1588_v4 = _dafny.Seq.of(_1586_v2);
      let _1589_v5;
      _1589_v5 = _dafny.MultiSet.fromElements(true, _1586_v2);
      let _1590_v6;
      _1590_v6 = _module.D0.create_DC1(_1587_v3);
      let _1591_v7;
      _1591_v7 = _dafny.Map.Empty.slice().updateUnsafe(_1584_v0,_1590_v6);
      let _1592_v8;
      _1592_v8 = _dafny.Set.fromElements(_1586_v2);
      let _1593_v9;
      let _nw245 = Array((new BigNumber(28)).toNumber());
      _nw245[(_dafny.ZERO).toNumber()] = _dafny.Seq.contains(_1585_v1, _1584_v0);
      _nw245[(_dafny.ONE).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(2)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(343)), ((_1594_v0) => function (_1595_i0) {
        return _1594_v0;
      })(_1584_v0)), _1584_v0);
      _nw245[(new BigNumber(3)).toNumber()] = _dafny.areEqual(_1585_v1, _module.__default.fm31(_1587_v3, _1586_v2, new BigNumber((_1588_v4).length), globalState));
      _nw245[(new BigNumber(4)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(5)).toNumber()] = (_1589_v5).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).fm5(new BigNumber((_1585_v1).length), _1586_v2, _1591_v7, _1586_v2, globalState))));
      _nw245[(new BigNumber(6)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(7)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(8)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(9)).toNumber()] = (_1587_v3).isLessThanOrEqualTo(_1587_v3);
      _nw245[(new BigNumber(10)).toNumber()] = (_1586_v2) && (_1586_v2);
      _nw245[(new BigNumber(11)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(12)).toNumber()] = (_1586_v2) && (true);
      _nw245[(new BigNumber(13)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(14)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(15)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(16)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(17)).toNumber()] = (_1592_v8).IsDisjointFrom(_dafny.Set.fromElements(_1586_v2));
      _nw245[(new BigNumber(18)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(19)).toNumber()] = ((_1586_v2) ? (!(_1586_v2)) : (true));
      _nw245[(new BigNumber(20)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(21)).toNumber()] = !(((_1586_v2) ? (_1586_v2) : (true)));
      _nw245[(new BigNumber(22)).toNumber()] = false;
      _nw245[(new BigNumber(23)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(24)).toNumber()] = (_1587_v3).isLessThan(_1587_v3);
      _nw245[(new BigNumber(25)).toNumber()] = (_1589_v5).IsSubsetOf(_1589_v5);
      _nw245[(new BigNumber(26)).toNumber()] = _1586_v2;
      _nw245[(new BigNumber(27)).toNumber()] = _1586_v2;
      _1593_v9 = _nw245;
      let _index254 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length));
      (_1593_v9)[_index254] = _1586_v2;
      let _index255 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length));
      (_1593_v9)[_index255] = _1586_v2;
      let _1596_v10;
      _1596_v10 = _dafny.Seq.of(_1587_v3);
      let _1597_v11;
      _1597_v11 = _dafny.Seq.of(_1596_v10, _dafny.Seq.Create(_module.__default.abs(new BigNumber(185)), ((_1598_v10) => function (_1599_i1) {
        return new BigNumber((_1598_v10).length);
      })(_1596_v10)));
      let _1600_v12;
      _1600_v12 = _dafny.Set.fromElements(_dafny.Seq.of(_1587_v3), (_1597_v11)[_module.__default.safeIndex(_1587_v3, new BigNumber((_1597_v11).length))]);
      let _1601_v13;
      _1601_v13 = _module.D15.create_DC37(_1600_v12);
      _1601_v13 = _1601_v13;
      let _1602_v14;
      let _nw246 = new _module.C2();
      _nw246.__ctor();
      _1602_v14 = _nw246;
      let _index256 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length));
      let _rhs173 = _1602_v14;
      let _rhs174 = new BigNumber(689);
      let _rhs175 = _1587_v3;
      let _rhs176 = !((_1593_v9)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length))]) || ((_1593_v9)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length))]);
      let _lhs108 = _1593_v9;
      let _lhs109 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length));
      _1602_v14 = _rhs173;
      _1587_v3 = _rhs174;
      _1587_v3 = _rhs175;
      _lhs108[_lhs109] = _rhs176;
      let _1603_v15;
      _1603_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1587_v3,(_1593_v9)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length))]);
      _1603_v15 = (_1603_v15).update(_1587_v3, (_dafny.Set.fromElements(_1586_v2, _1586_v2, (_1593_v9)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length))])).equals(_1592_v8));
      let _hi10 = _1587_v3;
      for (let _1604_i2 = (_1587_v3).minus(_1587_v3); _1604_i2.isLessThan(_hi10); _1604_i2 = _1604_i2.plus(_dafny.ONE)) {
        _1585_v1 = _1585_v1;
        let _rhs177 = true;
        let _rhs178 = _dafny.Seq.contains(_1596_v10, _module.__default.safeModuloInt(_1587_v3, _1587_v3));
        r2 = _rhs177;
        _1586_v2 = _rhs178;
        let _1605_v16;
        let _nw247 = Array((new BigNumber(3)).toNumber()).fill([]);
        _1605_v16 = _nw247;
        let _1606_v17;
        let _nw248 = Array((new BigNumber(20)).toNumber());
        _nw248[(_dafny.ZERO).toNumber()] = _1585_v1;
        _nw248[(_dafny.ONE).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(2)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(3)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_1585_v1, _module.__default.safeIndex(_1604_i2, new BigNumber((_1585_v1).length)), _1584_v0);
        _nw248[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("oe");
        _nw248[(new BigNumber(6)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(7)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(8)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(9)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(10)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(11)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(12)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(13)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(14)).toNumber()] = _dafny.Seq.UnicodeFromString("diut");
        _nw248[(new BigNumber(15)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(16)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(17)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(18)).toNumber()] = _1585_v1;
        _nw248[(new BigNumber(19)).toNumber()] = _1585_v1;
        _1606_v17 = _nw248;
        let _index257 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_1605_v16).length));
        (_1605_v16)[_index257] = _1606_v17;
        let _1607_v18;
        _1607_v18 = _dafny.Seq.of(_1606_v17);
        let _index258 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_1605_v16).length));
        (_1605_v16)[_index258] = (_1607_v18)[_module.__default.safeIndex(_1587_v3, new BigNumber((_1607_v18).length))];
        let _1608_v19;
        _1608_v19 = _dafny.Set.fromElements(_1593_v9);
        let _index259 = _module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length));
        (_1593_v9)[_index259] = (_1608_v19).IsProperSubsetOf(_1608_v19);
      }
      let _1609_v20;
      _1609_v20 = _dafny.Set.fromElements(_1593_v9, _1593_v9, _1593_v9);
      let _1610_v21;
      _1610_v21 = _1609_v20;
      let _1611_v22;
      _1611_v22 = ((_1610_v21)).Difference(_1609_v20);
      let _source17 = _1610_v21;
      let _1612___mcc_h0 = _source17;
      let _1613_cf36 = _1612___mcc_h0;
      _1587_v3 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeDivisionInt(_1587_v3, _1587_v3), _module.__default.safeDivisionInt(_1587_v3, new BigNumber(317))));
      let _1614_v23;
      let _nw249 = Array((new BigNumber(3)).toNumber()).fill(_dafny.ZERO);
      _1614_v23 = _nw249;
      _1614_v23 = _1614_v23;
      let _1615_v25;
      _1615_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1584_v0,_1587_v3);
      let _1616_v27;
      _1616_v27 = _dafny.Set.fromElements(_1584_v0);
      let _1617_v28;
      _1617_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1584_v0,_dafny.Seq.of((_1593_v9)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length))], (_1593_v9)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length))]));
      let _1618_v30;
      _1618_v30 = _module.D6.create_DC14(_1585_v1);
      let _1619_v31;
      _1619_v31 = _dafny.Map.Empty.slice().updateUnsafe(_1584_v0,_1618_v30);
      let _1620_v32;
      _1620_v32 = _dafny.Map.Empty.slice().updateUnsafe(_1587_v3,_1617_v28);
      let _1621_v33;
      _1621_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1587_v3,_1587_v3);
      let _1622_v35;
      let _nw250 = Array((new BigNumber(23)).toNumber());
      _nw250[(_dafny.ZERO).toNumber()] = function () {
        let _coll59 = new _dafny.Map();
        for (const _compr_59 of (_1615_v25).Keys.Elements) {
          let _1623_v24 = _compr_59;
          if ((_1615_v25).contains(_1623_v24)) {
            _coll59.push([_1623_v24,_1588_v4]);
          }
        }
        return _coll59;
      }();
      _nw250[(_dafny.ONE).toNumber()] = function () {
        let _coll60 = new _dafny.Map();
        for (const _compr_60 of (_1616_v27).Elements) {
          let _1624_v26 = _compr_60;
          if ((_1616_v27).contains(_1624_v26)) {
            _coll60.push([_1624_v26,_1588_v4]);
          }
        }
        return _coll60;
      }();
      _nw250[(new BigNumber(2)).toNumber()] = (((_1593_v9)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length))]) ? (_dafny.Map.Empty.slice().updateUnsafe(_1584_v0,_1588_v4)) : (_1617_v28));
      _nw250[(new BigNumber(3)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(4)).toNumber()] = function () {
        let _coll61 = new _dafny.Map();
        for (const _compr_61 of (_1619_v31).Keys.Elements) {
          let _1625_v29 = _compr_61;
          if ((_1619_v31).contains(_1625_v29)) {
            _coll61.push([_1625_v29,_1588_v4]);
          }
        }
        return _coll61;
      }();
      _nw250[(new BigNumber(5)).toNumber()] = (((_1620_v32).contains((((_1621_v33).contains(_1587_v3)) ? ((_1621_v33).get(_1587_v3)) : (_1587_v3)))) ? ((_1620_v32).get((((_1621_v33).contains(_1587_v3)) ? ((_1621_v33).get(_1587_v3)) : (_1587_v3)))) : (_1617_v28));
      _nw250[(new BigNumber(6)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(7)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(8)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(9)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1584_v0,_1588_v4);
      _nw250[(new BigNumber(10)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(11)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(12)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(13)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(14)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm51(false, _1587_v3, _1587_v3, globalState),_dafny.Seq.of(!((_1593_v9)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length))])));
      _nw250[(new BigNumber(16)).toNumber()] = (_1617_v28).Merge(_1617_v28);
      _nw250[(new BigNumber(17)).toNumber()] = (_1617_v28).Merge(_1617_v28);
      _nw250[(new BigNumber(18)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1584_v0,_1588_v4);
      _nw250[(new BigNumber(19)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(20)).toNumber()] = _1617_v28;
      _nw250[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_1584_v0,_1588_v4);
      _nw250[(new BigNumber(22)).toNumber()] = (function () {
        let _coll62 = new _dafny.Map();
        for (const _compr_62 of (_dafny.Seq.of(_1584_v0)).Elements) {
          let _1626_v34 = _compr_62;
          if (_dafny.Seq.contains(_dafny.Seq.of(_1584_v0), _1626_v34)) {
            _coll62.push([_1626_v34,_1588_v4]);
          }
        }
        return _coll62;
      }()).Merge((_1617_v28).update(_module.__default.fm51((_1593_v9)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length))], _1587_v3, new BigNumber((_dafny.Seq.of(_1587_v3)).length), globalState), _1588_v4));
      _1622_v35 = _nw250;
      let _index260 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_1622_v35).length));
      (_1622_v35)[_index260] = function () {
        let _coll63 = new _dafny.Map();
        for (const _compr_63 of (_1617_v28).Keys.Elements) {
          let _1627_v36 = _compr_63;
          if ((_1617_v28).contains(_1627_v36)) {
            _coll63.push([_1627_v36,_dafny.Seq.of(_1586_v2, (_1593_v9)[_module.__default.safeIndex(new BigNumber(117), new BigNumber((_1593_v9).length))])]);
          }
        }
        return _coll63;
      }();
      let _index261 = _module.__default.safeIndex(new BigNumber(60), new BigNumber((_1622_v35).length));
      (_1622_v35)[_index261] = ((_1617_v28).Merge(function () {
        let _coll64 = new _dafny.Map();
        for (const _compr_64 of (_1585_v1).Elements) {
          let _1628_v37 = _compr_64;
          if (_dafny.Seq.contains(_1585_v1, _1628_v37)) {
            _coll64.push([_1628_v37,_1588_v4]);
          }
        }
        return _coll64;
      }())).Merge(_1617_v28);
      let _1629_v38;
      let _nw251 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
      _1629_v38 = _nw251;
      let _1630_v39;
      let _nw252 = new _module.C1();
      _nw252.__ctor(new BigNumber((_dafny.Seq.Concat(_1585_v1, _1585_v1)).length), _1629_v38, _1596_v10, _1590_v6);
      _1630_v39 = _nw252;
      let _1631_v40;
      let _init35 = ((_1632_v6) => function (_1633_i3) {
        return _1632_v6;
      })(_1590_v6);
      let _nw253 = Array((new BigNumber(3)).toNumber());
      for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw253.length); _i0_35++) {
        _nw253[_i0_35] = _init35(new BigNumber(_i0_35));
      }
      _1631_v40 = _nw253;
      r0 = _1631_v40;
      let _1634_v41;
      let _init36 = ((_1635_v3) => function (_1636_i4) {
        return _module.__default.safeDivisionInt(_1636_i4, (_dafny.ZERO).minus(_1635_v3));
      })(_1587_v3);
      let _nw254 = Array((new BigNumber(3)).toNumber());
      for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw254.length); _i0_36++) {
        _nw254[_i0_36] = _init36(new BigNumber(_i0_36));
      }
      _1634_v41 = _nw254;
      let _1637_v42;
      _1637_v42 = _module.D13.create_DC32(_1585_v1, _1586_v2, new BigNumber((_dafny.MultiSet.fromElements(_1586_v2, false)).cardinality()), _1634_v41);
      let _1638_v43;
      _1638_v43 = _dafny.Seq.of(_1585_v1);
      let _1639_v44;
      _1639_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_1638_v43)[_module.__default.safeIndex(_1587_v3, new BigNumber((_1638_v43).length))]).length),_1634_v41);
      let _1640_v45;
      let _nw255 = Array((new BigNumber(27)).toNumber());
      _nw255[(_dafny.ZERO).toNumber()] = _1634_v41;
      _nw255[(_dafny.ONE).toNumber()] = ((_1586_v2) ? (_1634_v41) : (_1634_v41));
      _nw255[(new BigNumber(2)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(3)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(4)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(5)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(6)).toNumber()] = (_1637_v42).dtor_cf60;
      _nw255[(new BigNumber(7)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(8)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(9)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(10)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(11)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(12)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(13)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(14)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(15)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(16)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(17)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(18)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(19)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(20)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(21)).toNumber()] = (((_1639_v44).contains(_1587_v3)) ? ((_1639_v44).get(_1587_v3)) : (_1634_v41));
      _nw255[(new BigNumber(22)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(23)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(24)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(25)).toNumber()] = _1634_v41;
      _nw255[(new BigNumber(26)).toNumber()] = _1634_v41;
      _1640_v45 = _nw255;
      r1 = _1640_v45;
      r2 = (new BigNumber(396)).isLessThan(new BigNumber((_1585_v1).length));
      return [r0, r1, r2];
    }
    m10(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let _1641_v0;
      _1641_v0 = new _dafny.CodePoint('u'.codePointAt(0));
      let _1642_v1;
      _1642_v1 = _dafny.Map.Empty.slice().updateUnsafe(_1641_v0,p1);
      let _1643_v2;
      _1643_v2 = _1642_v1;
      _1643_v2 = _1642_v1;
      let _hi11 = p0;
      for (let _1644_i0 = new BigNumber(-403); _1644_i0.isLessThan(_hi11); _1644_i0 = _1644_i0.plus(_dafny.ONE)) {
        let _1645_v3;
        _1645_v3 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
        let _1646_v4;
        _1646_v4 = _dafny.MultiSet.fromElements(_1644_i0);
        let _1647_v5;
        _1647_v5 = _module.D0.create_DC1(p1);
        let _1648_v6;
        _1648_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1641_v0,_1647_v5);
        let _1649_v7;
        _1649_v7 = _dafny.Seq.of(!((_this).fm5(_1644_i0, p2, _1648_v6, p2, globalState)));
        let _1650_v8;
        _1650_v8 = _dafny.Seq.of(new BigNumber((_1645_v3).length));
        let _1651_v9;
        let _nw256 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _1651_v9 = _nw256;
        let _1652_v10;
        _1652_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_1651_v9)).length),p1);
        let _1653_v11;
        _1653_v11 = _dafny.Seq.of(_1646_v4);
        let _1654_v12;
        let _nw257 = Array((new BigNumber(23)).toNumber());
        _nw257[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(new BigNumber((((p2) ? (_1645_v3) : (_1645_v3))).length));
        _nw257[(_dafny.ONE).toNumber()] = new BigNumber((_1646_v4).cardinality());
        _nw257[(new BigNumber(2)).toNumber()] = (p0).multipliedBy(new BigNumber((_1649_v7).length));
        _nw257[(new BigNumber(3)).toNumber()] = p0;
        _nw257[(new BigNumber(4)).toNumber()] = new BigNumber((_1650_v8).length);
        _nw257[(new BigNumber(5)).toNumber()] = new BigNumber(205);
        _nw257[(new BigNumber(6)).toNumber()] = _1644_i0;
        _nw257[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw257[(new BigNumber(8)).toNumber()] = p0;
        _nw257[(new BigNumber(9)).toNumber()] = p1;
        _nw257[(new BigNumber(10)).toNumber()] = p1;
        _nw257[(new BigNumber(11)).toNumber()] = p1;
        _nw257[(new BigNumber(12)).toNumber()] = p1;
        _nw257[(new BigNumber(13)).toNumber()] = new BigNumber(-828);
        _nw257[(new BigNumber(14)).toNumber()] = (((_1652_v10).contains((_dafny.ZERO).minus(_1644_i0))) ? ((_1652_v10).get((_dafny.ZERO).minus(_1644_i0))) : (_module.__default.fm20(false, globalState)));
        _nw257[(new BigNumber(15)).toNumber()] = _module.__default.safeDivisionInt(p0, _1644_i0);
        _nw257[(new BigNumber(16)).toNumber()] = new BigNumber(892);
        _nw257[(new BigNumber(17)).toNumber()] = _1644_i0;
        _nw257[(new BigNumber(18)).toNumber()] = (p1).multipliedBy(_1644_i0);
        _nw257[(new BigNumber(19)).toNumber()] = p1;
        _nw257[(new BigNumber(20)).toNumber()] = p1;
        _nw257[(new BigNumber(21)).toNumber()] = _1644_i0;
        _nw257[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(779),new BigNumber((_1653_v11).length))).length));
        _1654_v12 = _nw257;
        let _index262 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_1654_v12).length));
        (_1654_v12)[_index262] = p1;
        let _index263 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_1654_v12).length));
        (_1654_v12)[_index263] = (_dafny.ZERO).minus(p1);
        _1641_v0 = _1641_v0;
        let _index264 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_1654_v12).length));
        (_1654_v12)[_index264] = (_dafny.ZERO).minus(p0);
        let _1655_v13;
        _1655_v13 = _dafny.MultiSet.fromElements(p2);
        r0 = (_1655_v13).IsSubsetOf((_1655_v13).Intersect(_1655_v13));
      }
      let _1656_v14;
      let _init37 = ((_1657_p0) => function (_1658_i1) {
        return _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_1657_p0));
      })(p0);
      let _nw258 = Array((new BigNumber(27)).toNumber());
      for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw258.length); _i0_37++) {
        _nw258[_i0_37] = _init37(new BigNumber(_i0_37));
      }
      _1656_v14 = _nw258;
      let _1659_v15;
      _1659_v15 = _dafny.MultiSet.fromElements(p0);
      let _index265 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1656_v14).length));
      (_1656_v14)[_index265] = _1659_v15;
      let _index266 = _module.__default.safeIndex(new BigNumber(667), new BigNumber((_1656_v14).length));
      (_1656_v14)[_index266] = ((p2) ? (_1659_v15) : ((_1659_v15).Difference(_1659_v15)));
      let _1660_v16;
      _1660_v16 = _dafny.Seq.of(_1641_v0);
      let _1661_v17;
      let _init38 = ((_1662_v1) => function (_1663_i2) {
        return _1662_v1;
      })(_1642_v1);
      let _nw259 = Array((new BigNumber(3)).toNumber());
      for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw259.length); _i0_38++) {
        _nw259[_i0_38] = _init38(new BigNumber(_i0_38));
      }
      _1661_v17 = _nw259;
      let _1664_v18;
      _1664_v18 = _dafny.Set.fromElements(p3, p3, p3);
      let _1665_v19;
      _1665_v19 = _dafny.Seq.of(p0, new BigNumber((_1664_v18).length), p0, p1, p1);
      let _1666_v20;
      let _nw260 = new _module.C1();
      _nw260.__ctor((p1).plus(new BigNumber((_1660_v16).length)), _1661_v17, _dafny.Seq.Concat(_1665_v19, _dafny.Seq.update(_1665_v19, _module.__default.safeIndex(p1, new BigNumber((_1665_v19).length)), new BigNumber(661))), _module.D0.create_DC1(p1));
      _1666_v20 = _nw260;
      let _1667_v21;
      let _nw261 = new _module.C2();
      _nw261.__ctor();
      _1667_v21 = _nw261;
      let _1668_v22;
      _1668_v22 = _module.D0.create_DC1(_1666_v20.f9);
      let _1669_v23;
      _1669_v23 = _dafny.Seq.of(_1668_v22);
      let _1670_v24;
      _1670_v24 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(320),_dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_1669_v23, _module.__default.safeIndex(_1666_v20.f9, new BigNumber((_1669_v23).length)), _1668_v22), _1669_v23));
      _1670_v24 = _1670_v24;
      r0 = p2;
      r1 = _1666_v20.f9;
      return [r0, r1];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f1 = _module.D0.Default();
      this._f0 = _dafny.Seq.of();
      this._f7 = [];
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    set f1(value) {
      let _this = this;
      _this._f1 = value;
    };
    get f0() {
      let _this = this;
      return _this._f0;
    };
    __ctor(f7, f0, f1) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      return;
    }
    fm6(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(487);
    };
    m7(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Set.Empty;
      let r2 = false;
      let _1671_v0;
      _1671_v0 = _dafny.Seq.UnicodeFromString("ibpiei");
      let _1672_v1;
      _1672_v1 = new _dafny.CodePoint('c'.codePointAt(0));
      let _1673_v2;
      let _nw262 = new _module.C0();
      _nw262.__ctor(_dafny.Seq.update(_module.__default.fm16(_1671_v0, p0, globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm16(_1671_v0, p0, globalState)).length)), _1672_v1));
      _1673_v2 = _nw262;
      if (p3) {
        let _1674_v3;
        _1674_v3 = _dafny.Set.fromElements(p3);
        let _index267 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
        ((_this).f7)[_index267] = ((_this).f0)[_module.__default.safeIndex(new BigNumber((_1674_v3).length), new BigNumber(((_this).f0).length))];
        let _1675_v4;
        _1675_v4 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _index268 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
        ((_this).f7)[_index268] = ((((_1675_v4).contains(new BigNumber((_dafny.Seq.UnicodeFromString("wsieh")).length))) ? ((_1675_v4).get(new BigNumber((_dafny.Seq.UnicodeFromString("wsieh")).length))) : (p0))).minus(new BigNumber(333));
        if (_module.__default.fm17(globalState)) {
          r2 = p3;
          let _1676_v5;
          _1676_v5 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))],p3);
          let _1677_v6;
          _1677_v6 = _dafny.Seq.of(p3, (((_1676_v5).contains(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))])) ? ((_1676_v5).get(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))])) : (p3)));
          let _1678_v7;
          let _init39 = ((_1679_p3) => function (_1680_i0) {
            return _1679_p3;
          })(p3);
          let _nw263 = Array((new BigNumber(26)).toNumber());
          for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw263.length); _i0_39++) {
            _nw263[_i0_39] = _init39(new BigNumber(_i0_39));
          }
          _1678_v7 = _nw263;
          let _1681_v8;
          _1681_v8 = _dafny.Map.Empty.slice().updateUnsafe(_1678_v7,_1677_v6);
          _1677_v6 = _dafny.Seq.Concat(_1677_v6, (((_1681_v8).contains(_1678_v7)) ? ((_1681_v8).get(_1678_v7)) : (_1677_v6)));
          r0 = _module.__default.safeDivisionInt((new BigNumber(-93)).plus(new BigNumber(-786)), ((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))]);
          let _index269 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
          let _rhs179 = p0;
          let _rhs180 = (((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))]).minus(p0);
          let _lhs110 = (_this).f7;
          let _lhs111 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
          _lhs110[_lhs111] = _rhs179;
          r0 = _rhs180;
          r0 = ((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))];
        } else {
          let _index270 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
          ((_this).f7)[_index270] = (((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))]).plus(p0);
          let _1682_v9;
          _1682_v9 = _dafny.Set.fromElements(_dafny.Set.fromElements(p3), _1674_v3);
          let _1683_v10;
          let _nw264 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
          _1683_v10 = _nw264;
          let _rhs181 = _module.__default.fm17(globalState);
          let _rhs182 = function () {
            let _coll65 = new _dafny.Set();
            for (const _compr_65 of (_module.__default.fm18(globalState)).Elements) {
              let _1684_v11 = _compr_65;
              if ((_module.__default.fm18(globalState)).contains(_1684_v11)) {
                _coll65.add(_1684_v11);
              }
            }
            return _coll65;
          }();
          let _rhs183 = _1683_v10;
          let _rhs184 = p3;
          let _rhs185 = !(p3);
          r2 = _rhs181;
          _1682_v9 = _rhs182;
          _1683_v10 = _rhs183;
          r2 = _rhs184;
          r2 = _rhs185;
          r2 = p3;
          let _1685_v12;
          _1685_v12 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
          _1685_v12 = (_1685_v12).update((_this).fm6(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))], p0, globalState), p3);
          _1671_v0 = _dafny.Seq.Concat(_1673_v2.f8, _dafny.Seq.Concat(_1671_v0, _1673_v2.f8));
        }
        let _index271 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
        ((_this).f7)[_index271] = p0;
        let _1686_v13;
        _1686_v13 = _dafny.Seq.of((_this).f0);
        r2 = !_dafny.Seq.contains(_1686_v13, (_this).f0);
        if (!(p3)) {
          _1671_v0 = _1673_v2.f8;
          (_this).f1 = _this.f1;
          r0 = _module.__default.safeDivisionInt(new BigNumber(202), p0);
          let _1687_v14;
          _1687_v14 = _dafny.Map.Empty.slice().updateUnsafe(_1672_v1,p0);
          let _1688_v15;
          _1688_v15 = _1687_v14;
          let _1689_v16;
          _1689_v16 = _dafny.Map.Empty.slice().updateUnsafe(false,p3);
          let _1690_v17;
          _1690_v17 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(_module.__default.fm17(globalState), p3, (_1688_v15), p3, globalState),(((_1689_v16).contains(p3)) ? ((_1689_v16).get(p3)) : (!(false))));
          _1690_v17 = ((_1690_v17).Merge(_1690_v17)).Merge(_1690_v17);
          let _1691_v18;
          _1691_v18 = _dafny.MultiSet.fromElements(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))]);
          let _1692_v19;
          _1692_v19 = _dafny.MultiSet.fromElements(p3, p3, p3, !(p3), p3);
          let _1693_v20;
          _1693_v20 = _dafny.Seq.of((((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))]).isLessThanOrEqualTo(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))]), p3, (_1692_v19).IsSubsetOf(_1692_v19));
          let _1694_v21;
          let _nw265 = new _module.C4();
          _nw265.__ctor();
          _1694_v21 = _nw265;
          let _1695_v22;
          _1695_v22 = _dafny.Seq.of(_1693_v20);
          let _1696_v23;
          _1696_v23 = _dafny.Seq.of((_1695_v22)[_module.__default.safeIndex(p0, new BigNumber((_1695_v22).length))]);
          let _index272 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
          let _rhs186 = ((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))];
          let _rhs187 = (_1691_v18).Difference(_1691_v18);
          let _rhs188 = (_1695_v22)[_module.__default.safeIndex(_module.__default.fm32(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))], globalState), new BigNumber((_1695_v22).length))];
          let _rhs189 = _1694_v21;
          let _rhs190 = ((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))];
          let _lhs112 = (_this).f7;
          let _lhs113 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
          r0 = _rhs186;
          _1691_v18 = _rhs187;
          _1693_v20 = _rhs188;
          _1694_v21 = _rhs189;
          _lhs112[_lhs113] = _rhs190;
        } else {
          let _1697_v24;
          let _nw266 = new _module.C2();
          _nw266.__ctor();
          _1697_v24 = _nw266;
          let _1698_v25;
          _1698_v25 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))],p3);
          let _1699_v26;
          _1699_v26 = _dafny.Seq.of(_1698_v25);
          let _1700_v27;
          _1700_v27 = _dafny.Map.Empty.slice().updateUnsafe((_1699_v26)[_module.__default.safeIndex(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))], new BigNumber((_1699_v26).length))],p0);
          let _1701_v28;
          _1701_v28 = _dafny.Map.Empty.slice().updateUnsafe(_1672_v1,_this.f1);
          let _1702_v29;
          _1702_v29 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_1697_v24).fm5(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))], false, _1701_v28, p3, globalState));
          let _1703_v30;
          _1703_v30 = _dafny.Set.fromElements((_this).f0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-806)), function (_1704_i1) {
            return ((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))];
          }));
          let _1705_v31;
          _1705_v31 = _dafny.MultiSet.fromElements(p3, true, p3, p3);
          let _index273 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
          let _index274 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
          let _rhs191 = ((p0).plus(new BigNumber(174))).minus((((_1700_v27).contains(_dafny.Map.Empty.slice().updateUnsafe(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))],p3))) ? ((_1700_v27).get(_dafny.Map.Empty.slice().updateUnsafe(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))],p3))) : (new BigNumber((_1702_v29).length))));
          let _rhs192 = _1673_v2.f8;
          let _rhs193 = !(((p3) ? ((_module.__default.fm58(_module.D16.create_DC42(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))], p0, p0), true, _1703_v30, globalState)).IsSubsetOf(_dafny.MultiSet.fromElements(_1672_v1, _1672_v1))) : (p3)));
          let _rhs194 = (new BigNumber((_1705_v31).cardinality())).multipliedBy(p0);
          let _rhs195 = p3;
          let _lhs114 = (_this).f7;
          let _lhs115 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
          let _lhs116 = (_this).f7;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
          _lhs114[_lhs115] = _rhs191;
          _1671_v0 = _rhs192;
          r2 = _rhs193;
          _lhs116[_lhs117] = _rhs194;
          r2 = _rhs195;
          let _index275 = _module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length));
          ((_this).f7)[_index275] = ((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))];
          r2 = !_dafny.areEqual(_1671_v0, _1673_v2.f8);
          let _1706_v32;
          let _init40 = ((_1707_p3) => function (_1708_i2) {
            return (_1707_p3) || (false);
          })(p3);
          let _nw267 = Array((new BigNumber(15)).toNumber());
          for (let _i0_40 = 0; _i0_40 < new BigNumber(_nw267.length); _i0_40++) {
            _nw267[_i0_40] = _init40(new BigNumber(_i0_40));
          }
          _1706_v32 = _nw267;
          let _1709_v33;
          _1709_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1672_v1,new BigNumber(583));
          let _1710_v34;
          _1710_v34 = _dafny.Set.fromElements(_1709_v33, _1709_v33);
          let _1711_v36;
          _1711_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1672_v1,true);
          let _index276 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_1706_v32).length));
          (_1706_v32)[_index276] = (_1710_v34).equals(_dafny.Set.fromElements(_1709_v33, (function () {
            let _coll66 = new _dafny.Map();
            for (const _compr_66 of (_1711_v36).Keys.Elements) {
              let _1712_v35 = _compr_66;
              if ((_1711_v36).contains(_1712_v35)) {
                _coll66.push([_1712_v35,((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))]]);
              }
            }
            return _coll66;
          }()).update(_1672_v1, (_this).fm6(((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))], ((_this).f7)[_module.__default.safeIndex(new BigNumber(80), new BigNumber(((_this).f7).length))], globalState))));
          let _index277 = _module.__default.safeIndex(new BigNumber(536), new BigNumber((_1706_v32).length));
          (_1706_v32)[_index277] = p3;
        }
      } else {
        r2 = (p0).isLessThanOrEqualTo(p0);
        let _1713_v37;
        _1713_v37 = _dafny.Seq.of(_1673_v2.f8);
        let _1714_v38;
        _1714_v38 = _dafny.Seq.of((_1713_v37)[_module.__default.safeIndex(p0, new BigNumber((_1713_v37).length))]);
        let _1715_v39;
        _1715_v39 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(p0, p0))).cardinality())).minus(new BigNumber((_1713_v37).length)),_dafny.Seq.Concat(_dafny.Seq.of(((_this).f0)[_module.__default.safeIndex(new BigNumber(((_this).f0).length), new BigNumber(((_this).f0).length))], p0, p0, p0, new BigNumber((p1).length)), (_this).f0));
        _1715_v39 = (_1715_v39).update(p0, (_this).f0);
        r2 = p3;
        r2 = !(!(p0).isEqualTo((_dafny.ZERO).minus(p0)));
        let _1716_v40;
        _1716_v40 = _dafny.Set.fromElements(p0);
        _1716_v40 = _1716_v40;
      }
      let _1717_v41;
      _1717_v41 = _dafny.Seq.of(p3);
      let _rhs196 = p3;
      let _rhs197 = ((p3) ? (p3) : (!(_dafny.Seq.IsProperPrefixOf(_1717_v41, _module.__default.fm19(p3, !(false), _dafny.Map.Empty.slice().updateUnsafe(_1672_v1,new BigNumber((_1673_v2.f8).length)), p3, globalState)))));
      let _rhs198 = p0;
      let _rhs199 = (_dafny.ZERO).minus((p0).multipliedBy(p0));
      r2 = _rhs196;
      r2 = _rhs197;
      r0 = _rhs198;
      r0 = _rhs199;
      let _1718_v42;
      _1718_v42 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),p3);
      _1718_v42 = (_1718_v42).Merge(_1718_v42);
      let _1719_v43;
      _1719_v43 = _module.D12.create_DC28(p3, p0);
      _1719_v43 = _1719_v43;
      let _hi12 = p0;
      for (let _1720_i3 = p0; _1720_i3.isLessThan(_hi12); _1720_i3 = _1720_i3.plus(_dafny.ONE)) {
        (_1673_v2).f8 = _1671_v0;
        r0 = p0;
        let _1721_v44;
        _1721_v44 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
        _1721_v44 = (_1721_v44).update((p0).isLessThan(p0), p3);
        let _1722_v45;
        _1722_v45 = _module.D11.create_DC25(_1672_v1, new BigNumber((_1673_v2.f8).length), true);
        let _1723_v46;
        let _nw268 = Array((new BigNumber(21)).toNumber());
        _nw268[(_dafny.ZERO).toNumber()] = p3;
        _nw268[(_dafny.ONE).toNumber()] = p3;
        _nw268[(new BigNumber(2)).toNumber()] = p3;
        _nw268[(new BigNumber(3)).toNumber()] = p3;
        _nw268[(new BigNumber(4)).toNumber()] = p3;
        _nw268[(new BigNumber(5)).toNumber()] = p3;
        _nw268[(new BigNumber(6)).toNumber()] = p3;
        _nw268[(new BigNumber(7)).toNumber()] = p3;
        _nw268[(new BigNumber(8)).toNumber()] = true;
        _nw268[(new BigNumber(9)).toNumber()] = p3;
        _nw268[(new BigNumber(10)).toNumber()] = p3;
        _nw268[(new BigNumber(11)).toNumber()] = (_1722_v45).dtor_cf43;
        _nw268[(new BigNumber(12)).toNumber()] = p3;
        _nw268[(new BigNumber(13)).toNumber()] = p3;
        _nw268[(new BigNumber(14)).toNumber()] = p3;
        _nw268[(new BigNumber(15)).toNumber()] = !(p3);
        _nw268[(new BigNumber(16)).toNumber()] = p3;
        _nw268[(new BigNumber(17)).toNumber()] = p3;
        _nw268[(new BigNumber(18)).toNumber()] = p3;
        _nw268[(new BigNumber(19)).toNumber()] = true;
        _nw268[(new BigNumber(20)).toNumber()] = p3;
        _1723_v46 = _nw268;
        let _1724_v47;
        _1724_v47 = _dafny.Set.fromElements(_1723_v46);
        let _1725_v48;
        _1725_v48 = _dafny.MultiSet.fromElements(new BigNumber((_1724_v47).length), _1720_i3, new BigNumber(239), p0, _1720_i3);
        let _1726_v49;
        _1726_v49 = _dafny.Set.fromElements(p3, (_1725_v48).equals(_1725_v48));
        _1726_v49 = _1726_v49;
      }
      let _1727_v50;
      _1727_v50 = _dafny.MultiSet.fromElements(new BigNumber((_1671_v0).length), (_dafny.ZERO).minus(p0));
      r0 = new BigNumber(((_1727_v50).Union((_1727_v50).Union(_1727_v50))).cardinality());
      r1 = _module.__default.fm59(globalState);
      r2 = true;
      return [r0, r1, r2];
    }
    m8(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Map.Empty;
      let r2 = _dafny.Seq.of();
      let _1728_v0;
      _1728_v0 = _dafny.Seq.UnicodeFromString("fi");
      let _1729_v1;
      let _nw269 = new _module.C0();
      _nw269.__ctor(_1728_v0);
      _1729_v1 = _nw269;
      let _1730_v2;
      _1730_v2 = new _dafny.CodePoint('j'.codePointAt(0));
      let _1731_v3;
      _1731_v3 = new BigNumber(-151);
      let _1732_v4;
      _1732_v4 = _module.D16.create_DC43(_dafny.Seq.UnicodeFromString("ekkqo"), _1730_v2, _1731_v3);
      let _pat_let_tv42 = p1;
      let _pat_let_tv43 = p1;
      let _pat_let_tv44 = p1;
      let _pat_let_tv45 = p3;
      r0 = !(function (_source18) {
        if (_source18.is_DC42) {
          let _1733___mcc_h0 = (_source18).cf76;
          let _1734___mcc_h1 = (_source18).cf77;
          let _1735___mcc_h2 = (_source18).cf78;
          let _1736_cf78 = _1735___mcc_h2;
          let _1737_cf77 = _1734___mcc_h1;
          let _1738_cf76 = _1733___mcc_h0;
          return _pat_let_tv42;
        } else if (_source18.is_DC43) {
          let _1739___mcc_h3 = (_source18).cf79;
          let _1740___mcc_h4 = (_source18).cf80;
          let _1741___mcc_h5 = (_source18).cf81;
          let _1742_cf81 = _1741___mcc_h5;
          let _1743_cf80 = _1740___mcc_h4;
          let _1744_cf79 = _1739___mcc_h3;
          return _pat_let_tv43;
        } else if (_source18.is_DC44) {
          let _1745___mcc_h6 = (_source18).cf82;
          let _1746_cf82 = _1745___mcc_h6;
          return _pat_let_tv44;
        } else {
          let _1747___mcc_h7 = (_source18).cf75;
          let _1748_cf75 = _1747___mcc_h7;
          return _pat_let_tv45;
        }
      }(_1732_v4));
      let _hi13 = _module.__default.fm44(p0, globalState);
      for (let _1749_i0 = (_dafny.ZERO).minus(new BigNumber((_1728_v0).length)); _1749_i0.isLessThan(_hi13); _1749_i0 = _1749_i0.plus(_dafny.ONE)) {
        let _1750_v5;
        let _nw270 = new _module.C0();
        _nw270.__ctor(_dafny.Seq.UnicodeFromString("m"));
        _1750_v5 = _nw270;
        _1731_v3 = new BigNumber((((p1) ? (_1750_v5.f8) : (_dafny.Seq.Concat(_1729_v1.f8, _dafny.Seq.Create(_module.__default.abs(new BigNumber(386)), ((_1751_v2) => function (_1752_i1) {
          return _1751_v2;
        })(_1730_v2)))))).length);
        r0 = false;
        let _1753_v6;
        _1753_v6 = _dafny.Seq.of(true);
        let _1754_v7;
        _1754_v7 = _module.D6.create_DC15(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-836)), ((_1755_v2) => function (_1756_i2) {
  return _1755_v2;
})(_1730_v2))).length), new BigNumber((_1753_v6).length), p3);
        let _1757_v8;
        _1757_v8 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1754_v7);
        let _source19 = _module.D6.create_DC16((((_1757_v8).contains(p3)) ? ((_1757_v8).get(p3)) : (_1754_v7)));
        if (_source19.is_DC15) {
          let _1758___mcc_h8 = (_source19).cf30;
          let _1759___mcc_h9 = (_source19).cf31;
          let _1760___mcc_h10 = (_source19).cf32;
          let _1761_cf32 = _1760___mcc_h10;
          let _1762_cf31 = _1759___mcc_h9;
          let _1763_cf30 = _1758___mcc_h8;
          let _rhs200 = ((_1762_cf31).multipliedBy(new BigNumber(263))).isLessThan((_1763_cf30).multipliedBy(_1763_cf30));
          let _rhs201 = (true) || (p0);
          let _rhs202 = (new BigNumber(191)).minus(_1731_v3);
          let _rhs203 = (_dafny.ZERO).minus(_1763_cf30);
          r0 = _rhs200;
          r0 = _rhs201;
          _1731_v3 = _rhs202;
          _1763_cf30 = _rhs203;
          let _1764_v9;
          let _nw271 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
          _1764_v9 = _nw271;
          let _1765_v13;
          let _init41 = ((_1766_i0, _1767_cf31) => function (_1768_i3) {
            return !(function () {
              let _coll67 = new _dafny.Set();
              for (const _compr_67 of _dafny.IntegerRange(new BigNumber(892), new BigNumber(820))) {
                let _1769_v10 = _compr_67;
                if (((new BigNumber(892)).isLessThanOrEqualTo(_1769_v10)) && ((_1769_v10).isLessThan(new BigNumber(820)))) {
                  _coll67.add((_1769_v10).minus(_1766_i0));
                }
              }
              return _coll67;
            }()).equals(function () {
              let _coll68 = new _dafny.Set();
              for (const _compr_68 of (function () {
                let _coll69 = new _dafny.Map();
                for (const _compr_69 of (_dafny.Set.fromElements(_1766_i0)).Elements) {
                  let _1770_v11 = _compr_69;
                  if ((_dafny.Set.fromElements(_1766_i0)).contains(_1770_v11)) {
                    _coll69.push([_module.__default.safeDivisionInt(_1770_v11, _1767_cf31),new BigNumber(49)]);
                  }
                }
                return _coll69;
              }()).Keys.Elements) {
                let _1771_v12 = _compr_68;
                if ((function () {
                  let _coll70 = new _dafny.Map();
                  for (const _compr_70 of (_dafny.Set.fromElements(_1766_i0)).Elements) {
                    let _1770_v11 = _compr_70;
                    if ((_dafny.Set.fromElements(_1766_i0)).contains(_1770_v11)) {
                      _coll70.push([_module.__default.safeDivisionInt(_1770_v11, _1767_cf31),new BigNumber(49)]);
                    }
                  }
                  return _coll70;
                }()).contains(_1771_v12)) {
                  _coll68.add(_module.__default.safeDivisionInt(_1771_v12, new BigNumber(793)));
                }
              }
              return _coll68;
            }());
          })(_1749_i0, _1762_cf31);
          let _nw272 = Array((new BigNumber(16)).toNumber());
          for (let _i0_41 = 0; _i0_41 < new BigNumber(_nw272.length); _i0_41++) {
            _nw272[_i0_41] = _init41(new BigNumber(_i0_41));
          }
          _1765_v13 = _nw272;
          let _index278 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_1765_v13).length));
          (_1765_v13)[_index278] = _module.__default.fm17(globalState);
          let _1772_v14;
          _1772_v14 = _dafny.Set.fromElements(_1731_v3, new BigNumber(-459));
          let _index279 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_1765_v13).length));
          let _rhs204 = (new BigNumber(459)).isEqualTo(new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_1761_cf32, !(p0), _1761_cf32), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1772_v14).length)), new BigNumber((_dafny.Seq.of(_1761_cf32, !(p0), _1761_cf32)).length)), p3)).length));
          let _rhs205 = _1731_v3;
          let _rhs206 = _1764_v9;
          let _rhs207 = p3;
          let _lhs118 = _1765_v13;
          let _lhs119 = _module.__default.safeIndex(new BigNumber(311), new BigNumber((_1765_v13).length));
          _1761_cf32 = _rhs204;
          _1731_v3 = _rhs205;
          _1764_v9 = _rhs206;
          _lhs118[_lhs119] = _rhs207;
          _1761_cf32 = !_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.of((_this).f0), _module.__default.safeIndex(_module.__default.fm20((_1765_v13)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_1765_v13).length))], globalState), new BigNumber((_dafny.Seq.of((_this).f0)).length)), (_this).f0), _module.__default.safeIndex(_1762_cf31, new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this).f0), _module.__default.safeIndex(_module.__default.fm20((_1765_v13)[_module.__default.safeIndex(new BigNumber(311), new BigNumber((_1765_v13).length))], globalState), new BigNumber((_dafny.Seq.of((_this).f0)).length)), (_this).f0)).length)), _dafny.Seq.of(_1749_i0)), _dafny.Seq.of(_dafny.Seq.of(_1731_v3), _dafny.Seq.update((_this).f0, _module.__default.safeIndex(_1731_v3, new BigNumber(((_this).f0).length)), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(906)), ((_1773_v2) => function (_1774_i4) {
            return _1773_v2;
          })(_1730_v2))).length)))), _dafny.Seq.Concat((_this).f0, (_this).f0));
          r0 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-884)), ((_1775_v2) => function (_1776_i5) {
            return _1775_v2;
          })(_1730_v2)), _1730_v2);
        } else if (_source19.is_DC14) {
          let _1777___mcc_h11 = (_source19).cf29;
          let _1778_cf29 = _1777___mcc_h11;
          r2 = _dafny.Seq.Concat((_this).f0, _dafny.Seq.update(_dafny.Seq.Concat((_this).f0, (_this).f0), _module.__default.safeIndex(((_this).f0)[_module.__default.safeIndex(_module.__default.fm44(_module.__default.fm17(globalState), globalState), new BigNumber(((_this).f0).length))], new BigNumber((_dafny.Seq.Concat((_this).f0, (_this).f0)).length)), _1731_v3));
          let _1779_v15;
          let _nw273 = new _module.C4();
          _nw273.__ctor();
          _1779_v15 = _nw273;
          _1730_v2 = _1730_v2;
          let _1780_v16;
          let _nw274 = new _module.C0();
          _nw274.__ctor(_1728_v0);
          _1780_v16 = _nw274;
        } else {
          let _1781___mcc_h12 = (_source19).cf33;
          let _1782_cf33 = _1781___mcc_h12;
          let _1783_v17;
          let _nw275 = Array((new BigNumber(15)).toNumber()).fill(false);
          _1783_v17 = _nw275;
          let _index280 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_1783_v17).length));
          (_1783_v17)[_index280] = p1;
          let _index281 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_1783_v17).length));
          (_1783_v17)[_index281] = p3;
          _1731_v3 = (new BigNumber(-536)).plus(_1731_v3);
          let _index282 = _module.__default.safeIndex(new BigNumber(971), new BigNumber((_1783_v17).length));
          (_1783_v17)[_index282] = (_dafny.Set.fromElements(new BigNumber(-289))).contains((_dafny.ZERO).minus(_module.__default.fm44(p1, globalState)));
          _1783_v17 = _1783_v17;
        }
      }
      let _1784_i6;
      _1784_i6 = _dafny.ZERO;
      L5: {
        while (false) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1784_i6)) {
              break L5;
            }
            _1784_i6 = (_1784_i6).plus(_dafny.ONE);
            r0 = p0;
            if (!(p0)) {
              let _1785_v18;
              let _nw276 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
              _1785_v18 = _nw276;
              _1785_v18 = (_this).f7;
              let _1786_v19;
              _1786_v19 = _dafny.Set.fromElements(_1731_v3);
              r0 = (_1786_v19).IsProperSubsetOf(_1786_v19);
              _1731_v3 = (_dafny.ZERO).minus((_module.__default.fm44(p1, globalState)).plus(_1731_v3));
              let _index283 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_1785_v18).length));
              (_1785_v18)[_index283] = _1731_v3;
              let _index284 = _module.__default.safeIndex(new BigNumber(12), new BigNumber((_1785_v18).length));
              (_1785_v18)[_index284] = _module.__default.fm20(p1, globalState);
              let _1787_v20;
              let _init42 = ((_1788_v2, _1789_v3) => function (_1790_i8) {
                return _dafny.Map.Empty.slice().updateUnsafe(_1788_v2,_1789_v3);
              })(_1730_v2, _1731_v3);
              let _nw277 = Array((new BigNumber(16)).toNumber());
              for (let _i0_42 = 0; _i0_42 < new BigNumber(_nw277.length); _i0_42++) {
                _nw277[_i0_42] = _init42(new BigNumber(_i0_42));
              }
              _1787_v20 = _nw277;
              let _1791_v21;
              let _nw278 = new _module.C1();
              _nw278.__ctor(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(966)), ((_1792_v2) => function (_1793_i7) {
                return _1792_v2;
              })(_1730_v2)), _dafny.Seq.UnicodeFromString("skvawdoq"))).length), _1787_v20, (_this).f0, _this.f1);
              _1791_v21 = _nw278;
            } else {
              let _1794_v22;
              let _nw279 = Array((new BigNumber(5)).toNumber()).fill([]);
              _1794_v22 = _nw279;
              let _1795_v23;
              let _nw280 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
              _1795_v23 = _nw280;
              let _index285 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_1794_v22).length));
              (_1794_v22)[_index285] = _1795_v23;
              let _index286 = _module.__default.safeIndex(new BigNumber(170), new BigNumber((_1794_v22).length));
              (_1794_v22)[_index286] = _1795_v23;
              let _1796_v24;
              _1796_v24 = _dafny.Set.fromElements(_1731_v3);
              let _1797_v25;
              _1797_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1796_v24,_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(408)), ((_1798_v3) => function (_1799_i9) {
                return _1798_v3;
              })(_1731_v3)), (_this).f0));
              let _1800_v26;
              _1800_v26 = _dafny.Map.Empty.slice().updateUnsafe(_1731_v3,_1796_v24);
              let _1801_v27;
              _1801_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1731_v3,p1);
              _1797_v25 = (_1797_v25).update((((_1800_v26).contains(new BigNumber((_1801_v27).length))) ? ((_1800_v26).get(new BigNumber((_1801_v27).length))) : (function () {
                let _coll71 = new _dafny.Set();
                for (const _compr_71 of _dafny.IntegerRange(new BigNumber(-23), new BigNumber(-358))) {
                  let _1802_v28 = _compr_71;
                  if (((new BigNumber(-23)).isLessThanOrEqualTo(_1802_v28)) && ((_1802_v28).isLessThan(new BigNumber(-358)))) {
                    _coll71.add(_module.__default.safeModuloInt(_1802_v28, _1731_v3));
                  }
                }
                return _coll71;
              }())), (_this).f0);
              let _1803_v29;
              let _nw281 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Set.Empty);
              _1803_v29 = _nw281;
              let _1804_v30;
              _1804_v30 = _dafny.Set.fromElements(p1, p0);
              let _index287 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_1803_v29).length));
              (_1803_v29)[_index287] = _1804_v30;
              let _index288 = _module.__default.safeIndex(new BigNumber(877), new BigNumber((_1803_v29).length));
              (_1803_v29)[_index288] = _1804_v30;
              let _1805_v31;
              _1805_v31 = _dafny.Map.Empty.slice().updateUnsafe(false,_1731_v3);
              let _1806_v32;
              _1806_v32 = _dafny.Set.fromElements(_1805_v31, _1805_v31);
              _1806_v32 = ((p3) ? (_1806_v32) : (_1806_v32));
              let _index289 = _module.__default.safeIndex(new BigNumber(125), new BigNumber(((_this).f7).length));
              ((_this).f7)[_index289] = new BigNumber(136);
              let _1807_v33;
              _1807_v33 = _module.D11.create_DC25(_1730_v2, _1731_v3, !(p1));
              let _1808_v34;
              _1808_v34 = _dafny.Map.Empty.slice().updateUnsafe(_1807_v33,new BigNumber(-795));
              let _1809_v36;
              _1809_v36 = _dafny.Map.Empty.slice().updateUnsafe(_1807_v33,p1);
              let _1810_v37;
              _1810_v37 = _dafny.MultiSet.fromElements((_1808_v34).Merge(_1808_v34), (_1808_v34).Merge(function () {
                let _coll72 = new _dafny.Map();
                for (const _compr_72 of (_1809_v36).Keys.Elements) {
                  let _1811_v35 = _compr_72;
                  if ((_1809_v36).contains(_1811_v35)) {
                    _coll72.push([_1811_v35,_1731_v3]);
                  }
                }
                return _coll72;
              }()));
              let _index290 = _module.__default.safeIndex(new BigNumber(125), new BigNumber(((_this).f7).length));
              ((_this).f7)[_index290] = (((_1810_v37).contains((_module.__default.fm60(_1731_v3, p1, globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1807_v33,_1731_v3)))) ? ((_1810_v37).get((_module.__default.fm60(_1731_v3, p1, globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1807_v33,_1731_v3)))) : (new BigNumber((_1728_v0).length)));
            }
            _1731_v3 = (_1731_v3).minus(_module.__default.safeDivisionInt(_1731_v3, _1731_v3));
            let _1812_v38;
            let _nw282 = Array((new BigNumber(8)).toNumber()).fill(false);
            _1812_v38 = _nw282;
            let _index291 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_1812_v38).length));
            (_1812_v38)[_index291] = p0;
            let _index292 = _module.__default.safeIndex(new BigNumber(153), new BigNumber((_1812_v38).length));
            (_1812_v38)[_index292] = ((true) ? (p1) : (p3));
          }
        }
      }
      let _1813_v39;
      _1813_v39 = _dafny.Seq.of(_module.__default.fm30(_1731_v3, p1, p1, globalState));
      let _1814_v40;
      _1814_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1813_v39).length),_1730_v2);
      _1814_v40 = (_1814_v40).update((_1731_v3).plus((_dafny.ZERO).minus(_1731_v3)), ((true) ? (_1730_v2) : (_1730_v2)));
      let _index293 = _module.__default.safeIndex(new BigNumber(934), new BigNumber(((_this).f7).length));
      ((_this).f7)[_index293] = (_1731_v3).multipliedBy(new BigNumber(446));
      let _index294 = _module.__default.safeIndex(new BigNumber(934), new BigNumber(((_this).f7).length));
      ((_this).f7)[_index294] = _1731_v3;
      let _1815_v41;
      _1815_v41 = _module.D2.create_DC5(new BigNumber(600), (_this).f0, false);
      r0 = (_1815_v41).dtor_cf10;
      r1 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("fhwdeveiq"),p1);
      r2 = (_this).f0;
      return [r0, r1, r2];
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f6 = _dafny.ZERO;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm11(p0, globalState) {
      let _this = this;
      return new BigNumber(261);
    };
    fm12(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(409);
    };
    m6(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _pat_let_tv46 = p1;
      let _source20 = function (_pat_let38_0) {
        return function (_1816_dt__update__tmp_h0) {
          return function (_pat_let39_0) {
            return function (_1817_dt__update_hcf2_h0) {
              return _module.D0.create_DC1(_1817_dt__update_hcf2_h0);
            }(_pat_let39_0);
          }(new BigNumber((_dafny.Set.fromElements((_this).f6, new BigNumber((_pat_let_tv46).length))).length));
        }(_pat_let38_0);
      }(_module.D0.create_DC1(p3));
      if (_source20.is_DC0) {
        let _1818___mcc_h0 = (_source20).cf0;
        let _1819___mcc_h1 = (_source20).cf1;
        let _1820_cf1 = _1819___mcc_h1;
        let _1821_cf0 = _1818___mcc_h0;
        let _1822_v0;
        _1822_v0 = false;
        if (_1822_v0) {
          let _1823_v1;
          _1823_v1 = _dafny.Seq.of(((_1822_v0) ? (_1822_v0) : (_1822_v0)), _1822_v0, _dafny.Seq.IsPrefixOf(p1, p1), _1822_v0);
          _1823_v1 = _1823_v1;
          let _1824_v2;
          _1824_v2 = _dafny.Seq.of(_1820_cf1);
          r3 = (_1824_v2)[_module.__default.safeIndex((_this).fm12(_1822_v0, _1822_v0, globalState), new BigNumber((_1824_v2).length))];
          _1820_cf1 = p3;
          let _1825_v3;
          _1825_v3 = _dafny.MultiSet.fromElements(((_this).f6).multipliedBy(new BigNumber((_1823_v1).length)));
          let _1826_v4;
          _1826_v4 = new _dafny.CodePoint('r'.codePointAt(0));
          let _1827_v5;
          _1827_v5 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f6),_1822_v0);
          let _1828_v6;
          _1828_v6 = _dafny.MultiSet.fromElements(false);
          let _1829_v7;
          _1829_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1828_v6).cardinality()),(_this).f6);
          let _1830_v8;
          _1830_v8 = _dafny.Seq.of(_1829_v7, _1829_v7);
          _1825_v3 = (_module.__default.fm13(_1826_v4, !((((_1827_v5).contains(p3)) ? ((_1827_v5).get(p3)) : (_1822_v0))), globalState)).update(new BigNumber(111), _module.__default.abs((p3).minus(new BigNumber((_1830_v8).length))));
          let _1831_v9;
          let _nw283 = new _module.C2();
          _nw283.__ctor();
          _1831_v9 = _nw283;
        } else {
          let _1832_v10;
          let _nw284 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Map.Empty);
          _1832_v10 = _nw284;
          let _1833_v11;
          _1833_v11 = _dafny.Map.Empty.slice().updateUnsafe(!(_1822_v0),_1822_v0);
          let _index295 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_1832_v10).length));
          (_1832_v10)[_index295] = (_1833_v11).Merge(_1833_v11);
          let _1834_v12;
          _1834_v12 = _dafny.Seq.of(_1822_v0, _1822_v0);
          let _index296 = _module.__default.safeIndex(new BigNumber(55), new BigNumber((_1832_v10).length));
          (_1832_v10)[_index296] = (_1833_v11).Merge(_dafny.Map.Empty.slice().updateUnsafe((_1834_v12)[_module.__default.safeIndex((_this).f6, new BigNumber((_1834_v12).length))],_1822_v0));
          let _1835_v13;
          _1835_v13 = _dafny.Map.Empty.slice().updateUnsafe(_1822_v0,(_this).f6);
          _1835_v13 = (_1835_v13).update((_1834_v12)[_module.__default.safeIndex(new BigNumber(918), new BigNumber((_1834_v12).length))], _1820_cf1);
          let _1836_v14;
          let _nw285 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Map.Empty);
          _1836_v14 = _nw285;
          let _1837_v15;
          _1837_v15 = new _dafny.CodePoint('s'.codePointAt(0));
          let _1838_v16;
          _1838_v16 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('t'.codePointAt(0)), _1837_v15);
          let _1839_v17;
          _1839_v17 = _dafny.Seq.of(p3, (_dafny.ZERO).minus(new BigNumber((_1838_v16).cardinality())), (_this).f6);
          let _1840_v18;
          let _nw286 = new _module.C1();
          _nw286.__ctor(new BigNumber((_1835_v13).length), _1836_v14, _1839_v17, p0);
          _1840_v18 = _nw286;
          let _1841_v19;
          _1841_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1840_v18,_1840_v18.f9);
          _1820_cf1 = ((_this).fm12(_1822_v0, !(_1822_v0), globalState)).plus((((_1841_v19).contains(_1840_v18)) ? ((_1841_v19).get(_1840_v18)) : ((_this).f6)));
          (_1840_v18).f9 = (_1840_v18).fm6(p3, _1820_cf1, globalState);
          let _1842_v20;
          let _init43 = ((_1843_v18, _1844_p1, _1845_v0) => function (_1846_i0) {
            return (new BigNumber((_dafny.MultiSet.fromElements(_1843_v18.f9, (_this).f6)).cardinality())).isLessThan((_module.D4.create_DC11(new BigNumber((_1844_p1).length), false, _1845_v0, _1845_v0, !(_1845_v0))).dtor_cf22);
          })(_1840_v18, p1, _1822_v0);
          let _nw287 = Array((new BigNumber(3)).toNumber());
          for (let _i0_43 = 0; _i0_43 < new BigNumber(_nw287.length); _i0_43++) {
            _nw287[_i0_43] = _init43(new BigNumber(_i0_43));
          }
          _1842_v20 = _nw287;
          let _index297 = _module.__default.safeIndex(new BigNumber(276), new BigNumber((_1842_v20).length));
          (_1842_v20)[_index297] = _1822_v0;
          let _index298 = _module.__default.safeIndex(new BigNumber(276), new BigNumber((_1842_v20).length));
          (_1842_v20)[_index298] = _1822_v0;
        }
        let _1847_v21;
        _1847_v21 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f6);
        _1847_v21 = (_1847_v21).update(p0, new BigNumber(509));
        let _1848_v22;
        _1848_v22 = new _dafny.CodePoint('i'.codePointAt(0));
        let _1849_v23;
        _1849_v23 = _module.D16.create_DC43(p1, new _dafny.CodePoint('u'.codePointAt(0)), new BigNumber(-656));
        let _1850_v24;
        _1850_v24 = _dafny.Seq.of(p1, p1, (_1849_v23).dtor_cf79);
        let _1851_v25;
        _1851_v25 = _dafny.Seq.of((_1850_v24)[_module.__default.safeIndex(p3, new BigNumber((_1850_v24).length))]);
        let _1852_v26;
        _1852_v26 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1822_v0);
        let _1853_v27;
        _1853_v27 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_dafny.Seq.of(p1, p1, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("nydykbd"), _module.__default.safeIndex(new BigNumber(-926), new BigNumber((_dafny.Seq.UnicodeFromString("nydykbd")).length)), _1848_v22), p1, p1), _1851_v25),_1852_v26);
        _1853_v27 = (_1853_v27).update(_1850_v24, _1852_v26);
        let _1854_v28;
        let _nw288 = new _module.C3();
        _nw288.__ctor();
        _1854_v28 = _nw288;
      } else {
        let _1855___mcc_h2 = (_source20).cf2;
        let _1856_cf2 = _1855___mcc_h2;
        let _1857_v29;
        let _nw289 = new _module.C0();
        _nw289.__ctor(_dafny.Seq.UnicodeFromString("jvlu"));
        _1857_v29 = _nw289;
        let _1858_v30;
        _1858_v30 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(461),true);
        let _1859_v31;
        _1859_v31 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,!((((_1858_v30).contains(new BigNumber(207))) ? ((_1858_v30).get(new BigNumber(207))) : (true))));
        let _1860_v32;
        _1860_v32 = true;
        r1 = ((_1858_v30).equals(_1859_v31)) === (_1860_v32);
        r1 = !((_1860_v32) && (_1860_v32)) || (_1860_v32);
        r2 = _1860_v32;
      }
      let _1861_v33;
      _1861_v33 = true;
      let _1862_v34;
      _1862_v34 = _dafny.MultiSet.fromElements(_1861_v33);
      let _1863_v35;
      _1863_v35 = _dafny.Seq.of(_1862_v34, _1862_v34);
      let _1864_v36;
      _1864_v36 = _dafny.Map.Empty.slice().updateUnsafe(p3,_1861_v33);
      let _1865_v37;
      _1865_v37 = _dafny.MultiSet.fromElements((_this).f6);
      let _1866_v38;
      _1866_v38 = _dafny.Set.fromElements(_1861_v33);
      let _1867_v39;
      let _nw290 = Array((new BigNumber(29)).toNumber());
      _nw290[(_dafny.ZERO).toNumber()] = new BigNumber((_1863_v35).length);
      _nw290[(_dafny.ONE).toNumber()] = p3;
      _nw290[(new BigNumber(2)).toNumber()] = (_this).f6;
      _nw290[(new BigNumber(3)).toNumber()] = new BigNumber(-178);
      _nw290[(new BigNumber(4)).toNumber()] = (p3).minus(p3);
      _nw290[(new BigNumber(5)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_1864_v36).length));
      _nw290[(new BigNumber(6)).toNumber()] = ((_dafny.ZERO).minus((_this).f6)).minus((_dafny.ZERO).minus((_this).f6));
      _nw290[(new BigNumber(7)).toNumber()] = _module.__default.safeModuloInt((_this).f6, new BigNumber(762));
      _nw290[(new BigNumber(8)).toNumber()] = (p3).minus((_this).f6);
      _nw290[(new BigNumber(9)).toNumber()] = (_this).f6;
      _nw290[(new BigNumber(10)).toNumber()] = new BigNumber(747);
      _nw290[(new BigNumber(11)).toNumber()] = new BigNumber(-715);
      _nw290[(new BigNumber(12)).toNumber()] = (_this).f6;
      _nw290[(new BigNumber(13)).toNumber()] = ((_this).f6).plus(new BigNumber((((_1862_v34).update(_1861_v33, _module.__default.abs((_this).f6))).update(_1861_v33, _module.__default.abs(new BigNumber((_1865_v37).cardinality())))).cardinality()));
      _nw290[(new BigNumber(14)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f6));
      _nw290[(new BigNumber(15)).toNumber()] = p3;
      _nw290[(new BigNumber(16)).toNumber()] = p3;
      _nw290[(new BigNumber(17)).toNumber()] = (_this).f6;
      _nw290[(new BigNumber(18)).toNumber()] = new BigNumber((_1865_v37).cardinality());
      _nw290[(new BigNumber(19)).toNumber()] = (_this).f6;
      _nw290[(new BigNumber(20)).toNumber()] = (p3).plus((_this).f6);
      _nw290[(new BigNumber(21)).toNumber()] = new BigNumber((_1866_v38).length);
      _nw290[(new BigNumber(22)).toNumber()] = new BigNumber(97);
      _nw290[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus((p3).multipliedBy(p3));
      _nw290[(new BigNumber(24)).toNumber()] = (_this).f6;
      _nw290[(new BigNumber(25)).toNumber()] = (_this).f6;
      _nw290[(new BigNumber(26)).toNumber()] = new BigNumber(296);
      _nw290[(new BigNumber(27)).toNumber()] = (_this).f6;
      _nw290[(new BigNumber(28)).toNumber()] = (_this).f6;
      _1867_v39 = _nw290;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1867_v39).length))) {
        let _1868_i1 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1868_i1)) && ((_1868_i1).isLessThan(new BigNumber((_1867_v39).length))))) {
          (_1867_v39)[(_1868_i1)] = _module.__default.safeDivisionInt(_1868_i1, (_this).f6);
        }
      }
      let _index299 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length));
      (_1867_v39)[_index299] = (_this).f6;
      let _index300 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length));
      (_1867_v39)[_index300] = (_dafny.ZERO).minus(_module.__default.safeDivisionInt((p3).plus(p3), ((_this).f6).plus((_dafny.ZERO).minus((_this).f6))));
      let _1869_v40;
      _1869_v40 = new _dafny.CodePoint('l'.codePointAt(0));
      let _1870_v41;
      _1870_v41 = _module.D15.create_DC38(_1869_v40);
      let _1871_v42;
      _1871_v42 = _dafny.Set.fromElements((_1870_v41).dtor_cf71);
      let _1872_v43;
      _1872_v43 = _dafny.Seq.of((_this).f6, new BigNumber((_1871_v42).length));
      let _1873_i2;
      _1873_i2 = _dafny.ZERO;
      L6: {
        while (_dafny.Seq.IsPrefixOf(_1872_v43, _1872_v43)) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1873_i2)) {
              break L6;
            }
            _1873_i2 = (_1873_i2).plus(_dafny.ONE);
            _1861_v33 = (_1861_v33) && (((_1861_v33) ? (_1861_v33) : (_1861_v33)));
            if (_1861_v33) {
              let _1874_v44;
              _1874_v44 = _module.D4.create_DC10(p3, (_this).f6, p3);
              let _1875_v45;
              _1875_v45 = _module.D15.create_DC39(_dafny.Seq.Concat(_module.__default.fm25(_1874_v44, p1, globalState), p1), _module.__default.fm61(new BigNumber(-243), (_this).f6, globalState));
              let _1876_v46;
              _1876_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1869_v40,(_this).f6);
              let _1877_v47;
              let _nw291 = new _module.C0();
              _nw291.__ctor(_module.__default.fm16(_module.__default.fm46(_1861_v33, p3, _1861_v33, globalState), (_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))], globalState));
              _1877_v47 = _nw291;
              let _1878_v48;
              _1878_v48 = _dafny.Set.fromElements(_1877_v47);
              let _1879_v49;
              _1879_v49 = _module.D4.create_DC9(_1865_v37, _1869_v40, p3, _1878_v48, (_1872_v43)[_module.__default.safeIndex((_this).f6, new BigNumber((_1872_v43).length))]);
              let _1880_v50;
              _1880_v50 = _dafny.Seq.of((_1876_v46).update((_1879_v49).dtor_cf15, p3));
              let _1881_v51;
              _1881_v51 = _module.D7.create_DC19(new BigNumber((_1877_v47.f8).length));
              let _1882_v52;
              _1882_v52 = _dafny.Map.Empty.slice().updateUnsafe(_1881_v51,p3);
              let _1883_v53;
              _1883_v53 = _dafny.Map.Empty.slice().updateUnsafe(_1882_v52,_1861_v33);
              _1875_v45 = _module.__default.fm62(_1880_v50, (((_1883_v53).contains(_1882_v52)) ? ((_1883_v53).get(_1882_v52)) : (_1861_v33)), (_this).f6, (_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))], globalState);
              let _1884_v54;
              _1884_v54 = _module.D13.create_DC32(p1, true, (_this).f6, _1867_v39);
              _1867_v39 = (_1884_v54).dtor_cf60;
              let _1885_v55;
              _1885_v55 = _dafny.Seq.of(true);
              _1861_v33 = (_1885_v55)[_module.__default.safeIndex(new BigNumber((_1885_v55).length), new BigNumber((_1885_v55).length))];
              r2 = _1861_v33;
              r3 = ((new BigNumber((_module.__default.fm50(_1861_v33, _1861_v33, new BigNumber((p1).length), globalState)).length)).multipliedBy(new BigNumber((_dafny.Seq.update(_1877_v47.f8, _module.__default.safeIndex((_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))], new BigNumber((_1877_v47.f8).length)), _1869_v40)).length))).plus((_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))]);
            } else {
              let _1886_v56;
              let _nw292 = Array((new BigNumber(6)).toNumber()).fill([]);
              _1886_v56 = _nw292;
              let _1887_v57;
              _1887_v57 = _dafny.Seq.of(_1861_v33, _1861_v33);
              let _1888_v58;
              let _nw293 = Array((new BigNumber(18)).toNumber());
              _nw293[(_dafny.ZERO).toNumber()] = _1861_v33;
              _nw293[(_dafny.ONE).toNumber()] = _1861_v33;
              _nw293[(new BigNumber(2)).toNumber()] = _1861_v33;
              _nw293[(new BigNumber(3)).toNumber()] = true;
              _nw293[(new BigNumber(4)).toNumber()] = _module.__default.fm17(globalState);
              _nw293[(new BigNumber(5)).toNumber()] = !(!(_1861_v33));
              _nw293[(new BigNumber(6)).toNumber()] = _module.__default.fm17(globalState);
              _nw293[(new BigNumber(7)).toNumber()] = true;
              _nw293[(new BigNumber(8)).toNumber()] = !(false);
              _nw293[(new BigNumber(9)).toNumber()] = true;
              _nw293[(new BigNumber(10)).toNumber()] = _1861_v33;
              _nw293[(new BigNumber(11)).toNumber()] = _1861_v33;
              _nw293[(new BigNumber(12)).toNumber()] = _1861_v33;
              _nw293[(new BigNumber(13)).toNumber()] = false;
              _nw293[(new BigNumber(14)).toNumber()] = !((_1887_v57)[_module.__default.safeIndex(p3, new BigNumber((_1887_v57).length))]);
              _nw293[(new BigNumber(15)).toNumber()] = _1861_v33;
              _nw293[(new BigNumber(16)).toNumber()] = _1861_v33;
              _nw293[(new BigNumber(17)).toNumber()] = !(_1861_v33);
              _1888_v58 = _nw293;
              let _index301 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1886_v56).length));
              (_1886_v56)[_index301] = _1888_v58;
              let _index302 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_1886_v56).length));
              (_1886_v56)[_index302] = _1888_v58;
              _1864_v36 = (_1864_v36).update((_dafny.ZERO).minus(p3), (_1861_v33) && (_module.__default.fm30((_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))], (_1887_v57)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f6), new BigNumber((_1887_v57).length))], !(_1861_v33), globalState)));
              r0 = _1861_v33;
              let _1889_v59;
              _1889_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1869_v40,new BigNumber((_1864_v36).length));
              _1889_v59 = (_1889_v59).update(((_1861_v33) ? (_1869_v40) : (_module.__default.fm29(!(_1861_v33), globalState))), new BigNumber((_dafny.Seq.UnicodeFromString("yyicitvk")).length));
              let _index303 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length));
              (_1867_v39)[_index303] = (_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))];
            }
            let _1890_v60;
            _1890_v60 = _dafny.Seq.of(_1861_v33, _1861_v33, _module.__default.fm17(globalState), _1861_v33, _1861_v33);
            let _1891_v61;
            _1891_v61 = _module.D12.create_DC28((_1890_v60)[_module.__default.safeIndex(p3, new BigNumber((_1890_v60).length))], p3);
            let _1892_v62;
            _1892_v62 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,p3);
            let _index304 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length));
            (_1867_v39)[_index304] = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p3, p3), _dafny.Seq.of(new BigNumber((p1).length), new BigNumber((_1890_v60).length), (_1891_v61).dtor_cf47, new BigNumber((_1892_v62).length)))).length);
            _1864_v36 = (_1864_v36).update(p3, _1861_v33);
          }
        }
      }
      if (((new BigNumber((p1).length)).minus((_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))])).isLessThan((_this).f6)) {
        let _1893_v63;
        _1893_v63 = _module.D12.create_DC29(!(true), (_this).f6, _1861_v33, _module.__default.safeDivisionInt((_this).f6, (_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))]), !(_1861_v33));
        _1893_v63 = _module.D12.create_DC29(_1861_v33, p3, _1861_v33, ((_this).f6).multipliedBy((_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))]), !(_1861_v33));
        let _1894_v64;
        _1894_v64 = _dafny.Seq.of(!(_1866_v38).equals(_1866_v38));
        _1894_v64 = _1894_v64;
        let _1895_v65;
        _1895_v65 = _module.D4.create_DC8(_dafny.Seq.of(p2, p2));
        let _1896_v66;
        _1896_v66 = _dafny.Map.Empty.slice().updateUnsafe(_1861_v33,_1895_v65);
        let _1897_v67;
        _1897_v67 = _dafny.Seq.of(p2, p2);
        _1896_v66 = (_1896_v66).update(!(_1861_v33), _module.D4.create_DC8(_1897_v67));
        let _1898_v68;
        let _nw294 = new _module.C3();
        _nw294.__ctor();
        _1898_v68 = _nw294;
        let _1899_v69;
        _1899_v69 = _dafny.Seq.UnicodeFromString("knsogq");
        _1899_v69 = p1;
      } else {
        let _1900_v70;
        let _nw295 = Array((new BigNumber(12)).toNumber());
        _nw295[(_dafny.ZERO).toNumber()] = _1861_v33;
        _nw295[(_dafny.ONE).toNumber()] = _module.__default.fm17(globalState);
        _nw295[(new BigNumber(2)).toNumber()] = _module.__default.fm30((_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))], true, !(_1861_v33), globalState);
        _nw295[(new BigNumber(3)).toNumber()] = _1861_v33;
        _nw295[(new BigNumber(4)).toNumber()] = true;
        _nw295[(new BigNumber(5)).toNumber()] = false;
        _nw295[(new BigNumber(6)).toNumber()] = ((_dafny.ZERO).minus((((_1865_v37).contains(new BigNumber(272))) ? ((_1865_v37).get(new BigNumber(272))) : (p3)))).isLessThan((_this).f6);
        _nw295[(new BigNumber(7)).toNumber()] = _1861_v33;
        _nw295[(new BigNumber(8)).toNumber()] = _1861_v33;
        _nw295[(new BigNumber(9)).toNumber()] = _1861_v33;
        _nw295[(new BigNumber(10)).toNumber()] = _1861_v33;
        _nw295[(new BigNumber(11)).toNumber()] = _1861_v33;
        _1900_v70 = _nw295;
        let _index305 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_1900_v70).length));
        (_1900_v70)[_index305] = _1861_v33;
        let _index306 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_1900_v70).length));
        (_1900_v70)[_index306] = _1861_v33;
        let _1901_v71;
        _1901_v71 = _dafny.Seq.UnicodeFromString("ydabfju");
        let _1902_v72;
        _1902_v72 = _dafny.MultiSet.fromElements(_dafny.Seq.update(p1, _module.__default.safeIndex(new BigNumber((_1872_v43).length), new BigNumber((p1).length)), _1869_v40), _dafny.Seq.UnicodeFromString("ayhpyhe"));
        let _1903_v73;
        _1903_v73 = _dafny.Set.fromElements((_this).f6, (_this).f6);
        let _index307 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_1900_v70).length));
        let _rhs208 = (((_1861_v33) ? (_1902_v72) : (_1902_v72))).IsSubsetOf(_1902_v72);
        let _rhs209 = (_this).f6;
        let _rhs210 = ((_1861_v33) ? ((_this).f6) : (_module.__default.safeModuloInt((_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))], (_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))])));
        let _rhs211 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("su"), _1901_v71);
        let _rhs212 = (_1903_v73).IsProperSubsetOf(_1903_v73);
        let _lhs120 = _1900_v70;
        let _lhs121 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_1900_v70).length));
        r2 = _rhs208;
        r3 = _rhs209;
        r3 = _rhs210;
        _1901_v71 = _rhs211;
        _lhs120[_lhs121] = _rhs212;
        let _1904_v74;
        _1904_v74 = _module.D12.create_DC28(_1861_v33, p3);
        let _index308 = _module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length));
        (_1867_v39)[_index308] = (_1904_v74).dtor_cf47;
        r3 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f6), (_this).fm12((_1900_v70)[_module.__default.safeIndex(new BigNumber(61), new BigNumber((_1900_v70).length))], (_1900_v70)[_module.__default.safeIndex(new BigNumber(61), new BigNumber((_1900_v70).length))], globalState));
        _1861_v33 = (_1900_v70)[_module.__default.safeIndex(new BigNumber(61), new BigNumber((_1900_v70).length))];
      }
      let _1905_v75;
      _1905_v75 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fm"), p1), _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(p3,_1861_v33)).length)), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("fm"), p1)).length)), _module.__default.fm51(_1861_v33, (_1867_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1867_v39).length))], p3, globalState)), p1, p1, p1, p1);
      _1905_v75 = _1905_v75;
      r0 = _dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-746)), ((_1906_v40) => function (_1907_i3) {
        return _1906_v40;
      })(_1869_v40)), _1869_v40);
      r1 = !(_module.__default.fm17(globalState));
      r2 = _1861_v33;
      let _1908_v76;
      let _init44 = ((_1909_v33, _1910_v36, _1911_v39) => function (_1912_i4) {
        return (((_1910_v36).contains(new BigNumber((_dafny.Seq.of(_1909_v33)).length))) ? ((_1910_v36).get(new BigNumber((_dafny.Seq.of(_1909_v33)).length))) : ((((_1910_v36).contains((_1911_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1911_v39).length))])) ? ((_1910_v36).get((_1911_v39)[_module.__default.safeIndex(new BigNumber(574), new BigNumber((_1911_v39).length))])) : (_1909_v33))));
      })(_1861_v33, _1864_v36, _1867_v39);
      let _nw296 = Array((new BigNumber(24)).toNumber());
      for (let _i0_44 = 0; _i0_44 < new BigNumber(_nw296.length); _i0_44++) {
        _nw296[_i0_44] = _init44(new BigNumber(_i0_44));
      }
      _1908_v76 = _nw296;
      let _1913_v77;
      _1913_v77 = _dafny.Set.fromElements(_1908_v76);
      let _1914_v78;
      _1914_v78 = _1913_v77;
      let _1915_v79;
      _1915_v79 = _dafny.Map.Empty.slice().updateUnsafe(_1861_v33,_1914_v78);
      let _1916_v80;
      _1916_v80 = _1861_v33;
      r3 = _module.__default.safeDivisionInt(new BigNumber(((_1915_v79).Merge((_1915_v79).update((_1916_v80), _1913_v77))).length), new BigNumber(93));
      return [r0, r1, r2, r3];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f1 = _module.D0.Default();
      this._f0 = _dafny.Seq.of();
      this._f4 = _dafny.MultiSet.Empty;
      this._f5 = _module.D0.Default();
    }
    _parentTraits() {
      return [_module.T1];
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    set f1(value) {
      let _this = this;
      _this._f1 = value;
    };
    get f0() {
      let _this = this;
      return _this._f0;
    };
    __ctor(f4, f5, f0, f1) {
      let _this = this;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      return;
    }
    fm6(p0, p1, globalState) {
      let _this = this;
      return (new BigNumber(-76)).minus((new BigNumber((_dafny.Seq.UnicodeFromString("af")).length)).minus(new BigNumber(-575)));
    };
    fm8(p0, globalState) {
      let _this = this;
      if ((!(false)) && (false)) {
        return !(!(new BigNumber((_dafny.Seq.of(false)).length)).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(440)), function (_1917_i0) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })).length)));
      } else {
        return (!(!(!(!(!(!(false))))))) === (false);
      }
    };
    fm9(p0, p1, p2, p3, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(606),new BigNumber(-633))).Merge(function () {
        let _coll73 = new _dafny.Map();
        for (const _compr_73 of _dafny.IntegerRange(new BigNumber(834), new BigNumber(708))) {
          let _1918_v0 = _compr_73;
          if (((new BigNumber(834)).isLessThanOrEqualTo(_1918_v0)) && ((_1918_v0).isLessThan(new BigNumber(708)))) {
            _coll73.push([(_1918_v0).multipliedBy(new BigNumber(597)),new BigNumber(52)]);
          }
        }
        return _coll73;
      }())).Merge((function () {
        let _coll74 = new _dafny.Map();
        for (const _compr_74 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(329),true)).Keys.Elements) {
          let _1919_v1 = _compr_74;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(329),true)).contains(_1919_v1)) {
            _coll74.push([_module.__default.safeModuloInt(_1919_v1, new BigNumber(809)),new BigNumber(743)]);
          }
        }
        return _coll74;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(312)), function (_1920_i0) {
        return new _dafny.CodePoint('j'.codePointAt(0));
      })).length),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("dx")).length)))));
    };
    m5(p0, p1, globalState) {
      let _this = this;
      let _1921_v0;
      _1921_v0 = true;
      let _1922_i0;
      _1922_i0 = _dafny.ZERO;
      L7: {
        while (_1921_v0) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1922_i0)) {
              break L7;
            }
            _1922_i0 = (_1922_i0).plus(_dafny.ONE);
            _1921_v0 = _1921_v0;
            let _1923_v1;
            let _nw297 = Array((new BigNumber(4)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
            _1923_v1 = _nw297;
            let _init45 = function (_1924_i1) {
              return new _dafny.CodePoint('c'.codePointAt(0));
            };
            let _nw298 = Array((new BigNumber(12)).toNumber());
            for (let _i0_45 = 0; _i0_45 < new BigNumber(_nw298.length); _i0_45++) {
              _nw298[_i0_45] = _init45(new BigNumber(_i0_45));
            }
            _1923_v1 = _nw298;
            let _1925_v2;
            let _nw299 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Seq.of());
            _1925_v2 = _nw299;
            let _index309 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1925_v2).length));
            (_1925_v2)[_index309] = (_this).f0;
            let _index310 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_1925_v2).length));
            (_1925_v2)[_index310] = _dafny.Seq.Concat((_this).f0, (_this).f0);
            let _1926_v3;
            let _init46 = function (_1927_i2) {
              return _dafny.Seq.Create(_module.__default.abs(new BigNumber(737)), function (_1928_i3) {
                return new _dafny.CodePoint('u'.codePointAt(0));
              });
            };
            let _nw300 = Array((new BigNumber(25)).toNumber());
            for (let _i0_46 = 0; _i0_46 < new BigNumber(_nw300.length); _i0_46++) {
              _nw300[_i0_46] = _init46(new BigNumber(_i0_46));
            }
            _1926_v3 = _nw300;
            _1926_v3 = _1926_v3;
          }
        }
      }
      let _1929_v4;
      let _nw301 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
      _1929_v4 = _nw301;
      let _1930_v5;
      _1930_v5 = _module.D0.create_DC0(_1929_v4, new BigNumber(600));
      let _1931_v6;
      _1931_v6 = _dafny.MultiSet.fromElements(_1930_v5);
      let _1932_v7;
      _1932_v7 = _dafny.Set.fromElements((_this).fm6(new BigNumber(((_1931_v6).update(_1930_v5, _module.__default.abs(p0))).cardinality()), p0, globalState));
      let _1933_v8;
      _1933_v8 = _dafny.Seq.of(_1932_v7, _1932_v7, _1932_v7, _1932_v7, _1932_v7);
      if (!((_module.__default.fm10(true, globalState)).IsSubsetOf((_1933_v8)[_module.__default.safeIndex(p0, new BigNumber((_1933_v8).length))]))) {
        _1932_v7 = _dafny.Set.fromElements(((_1921_v0) ? (p0) : (p0)), p0);
        let _1934_v9;
        let _nw302 = Array((new BigNumber(7)).toNumber()).fill([]);
        _1934_v9 = _nw302;
        _1934_v9 = _1934_v9;
        let _1935_v10;
        _1935_v10 = new _dafny.CodePoint('c'.codePointAt(0));
        let _1936_v11;
        _1936_v11 = _dafny.Seq.of(_1935_v10, _1935_v10, _1935_v10, _1935_v10);
        let _1937_v12;
        _1937_v12 = _dafny.Set.fromElements(_dafny.Seq.Concat(_1936_v11, _1936_v11));
        let _1938_v13;
        _1938_v13 = _module.D2.create_DC3(_1937_v12);
        let _1939_v14;
        _1939_v14 = _dafny.Seq.of(_1936_v11);
        let _pat_let_tv47 = _1939_v14;
        let _pat_let_tv48 = _1939_v14;
        _1937_v12 = (function (_pat_let40_0) {
          return function (_1940_dt__update__tmp_h0) {
            return function (_pat_let41_0) {
              return function (_1942_dt__update_hcf4_h0) {
                return _module.D2.create_DC3(_1942_dt__update_hcf4_h0);
              }(_pat_let41_0);
            }(function () {
              let _coll75 = new _dafny.Set();
              for (const _compr_75 of (_pat_let_tv47).Elements) {
                let _1941_v15 = _compr_75;
                if (_dafny.Seq.contains(_pat_let_tv48, _1941_v15)) {
                  _coll75.add(_1941_v15);
                }
              }
              return _coll75;
            }());
          }(_pat_let40_0);
        }(_1938_v13)).dtor_cf4;
        let _1943_v16;
        let _nw303 = Array((new BigNumber(3)).toNumber()).fill([]);
        _1943_v16 = _nw303;
        let _1944_v17;
        _1944_v17 = _module.D2.create_DC4(p0, _1921_v0, _1943_v16);
        if ((_1944_v17).dtor_cf6) {
          let _1945_v18;
          _1945_v18 = _dafny.Map.Empty.slice().updateUnsafe(_1935_v10,_1936_v11);
          let _1946_v19;
          _1946_v19 = _dafny.MultiSet.fromElements((((_1945_v18).contains(_1935_v10)) ? ((_1945_v18).get(_1935_v10)) : (_1936_v11)), _1936_v11, _dafny.Seq.Concat(_1936_v11, _1936_v11), _1936_v11);
          _1946_v19 = ((_1921_v0) ? ((_1946_v19).Union(_1946_v19)) : (_1946_v19));
          let _1947_v20;
          _1947_v20 = new BigNumber(221);
          _1947_v20 = _1947_v20;
          let _1948_v21;
          let _nw304 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Map.Empty);
          _1948_v21 = _nw304;
          let _1949_v22;
          let _nw305 = new _module.C1();
          _nw305.__ctor(_module.__default.fm32(_1947_v20, globalState), _1948_v21, (_this).f0, _module.__default.fm49(_1935_v10, globalState));
          _1949_v22 = _nw305;
          let _1950_v23;
          let _nw306 = new _module.C3();
          _nw306.__ctor();
          _1950_v23 = _nw306;
          let _1951_v24;
          _1951_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1947_v20,_1947_v20);
          let _1952_v25;
          _1952_v25 = _dafny.Map.Empty.slice().updateUnsafe((((_1951_v24).contains(_1947_v20)) ? ((_1951_v24).get(_1947_v20)) : ((_dafny.ZERO).minus(_1947_v20))),(_dafny.ZERO).minus(_1947_v20));
          let _1953_v26;
          _1953_v26 = _dafny.Seq.of(_1921_v0, _1921_v0);
          _1952_v25 = (_1952_v25).update(p0, new BigNumber((_1953_v26).length));
        } else {
          let _1954_v27;
          _1954_v27 = _dafny.Map.Empty.slice().updateUnsafe(_1921_v0,p0);
          let _1955_v28;
          _1955_v28 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1954_v27);
          let _1956_v29;
          _1956_v29 = _dafny.MultiSet.fromElements(_1921_v0);
          _1921_v0 = !(_module.__default.fm63((((_1955_v28).contains(p0)) ? ((_1955_v28).get(p0)) : (_dafny.Map.Empty.slice().updateUnsafe(_1921_v0,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_1921_v0),p0)).length)))), p0, globalState)).equals(_1956_v29);
          let _1957_v30;
          _1957_v30 = new BigNumber(838);
          _1957_v30 = _1957_v30;
          _1957_v30 = _1957_v30;
          let _1958_v31;
          _1958_v31 = _module.D12.create_DC28(false, _1957_v30);
          let _1959_v32;
          _1959_v32 = _dafny.Map.Empty.slice().updateUnsafe((_1958_v31).dtor_cf47,(_this).f0);
          let _1960_v33;
          _1960_v33 = _dafny.Map.Empty.slice().updateUnsafe(_1921_v0,false);
          _1936_v11 = _dafny.Seq.update(((_1921_v0) ? (_dafny.Seq.update(_1936_v11, _module.__default.safeIndex(new BigNumber((_1959_v32).length), new BigNumber((_1936_v11).length)), _module.__default.fm51(_1921_v0, new BigNumber((_1960_v33).length), _1957_v30, globalState))) : (_dafny.Seq.Concat(_1936_v11, _dafny.Seq.of(_1935_v10, _1935_v10, _1935_v10)))), _module.__default.safeIndex(_module.__default.safeModuloInt(_1957_v30, p0), new BigNumber((((_1921_v0) ? (_dafny.Seq.update(_1936_v11, _module.__default.safeIndex(new BigNumber((_1959_v32).length), new BigNumber((_1936_v11).length)), _module.__default.fm51(_1921_v0, new BigNumber((_1960_v33).length), _1957_v30, globalState))) : (_dafny.Seq.Concat(_1936_v11, _dafny.Seq.of(_1935_v10, _1935_v10, _1935_v10))))).length)), _1935_v10);
          let _1961_v34;
          _1961_v34 = _dafny.Seq.of(false, _1921_v0, !(p0).isEqualTo(new BigNumber(281)));
          _1961_v34 = _dafny.Seq.Concat(_1961_v34, _1961_v34);
        }
        let _1962_v35;
        let _init47 = ((_1963_v10, _1964_p0) => function (_1965_i4) {
          return _dafny.Map.Empty.slice().updateUnsafe(_1963_v10,_1964_p0);
        })(_1935_v10, p0);
        let _nw307 = Array((_dafny.ONE).toNumber());
        for (let _i0_47 = 0; _i0_47 < new BigNumber(_nw307.length); _i0_47++) {
          _nw307[_i0_47] = _init47(new BigNumber(_i0_47));
        }
        _1962_v35 = _nw307;
        let _1966_v36;
        let _nw308 = new _module.C1();
        _nw308.__ctor(p0, _1962_v35, (_this).f0, (_this).f5);
        _1966_v36 = _nw308;
        let _1967_v37;
        _1967_v37 = _dafny.Set.fromElements(_1966_v36, _1966_v36, _1966_v36);
        let _1968_v38;
        _1968_v38 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),new BigNumber((_1967_v37).length));
        let _1969_v39;
        _1969_v39 = _dafny.Map.Empty.slice().updateUnsafe(_1921_v0,new BigNumber((_1932_v7).length));
        let _1970_v40;
        _1970_v40 = _dafny.Map.Empty.slice().updateUnsafe(_1935_v10,_1966_v36.f9);
        let _1971_v41;
        _1971_v41 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19(_1921_v0, _1921_v0, _1970_v40, false, globalState),p0);
        let _1972_v42;
        _1972_v42 = _dafny.Seq.of(_1968_v38, (_1968_v38).update(new BigNumber((_1969_v39).length), new BigNumber((_1971_v41).length)), _1968_v38, (_1968_v38).Merge(_1968_v38));
        let _1973_v43;
        _1973_v43 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-919)), ((_1974_p0) => function (_1975_i5) {
          return _1974_p0;
        })(p0)),_1972_v42);
        let _1976_v44;
        _1976_v44 = _dafny.MultiSet.fromElements(_1921_v0);
        let _1977_v45;
        _1977_v45 = _dafny.Seq.of((_this).f5, _module.D0.create_DC1(p0), _this.f1, _module.D0.create_DC1(_1966_v36.f9), (_this).f5);
        _1972_v42 = _dafny.Seq.update(_dafny.Seq.update((((p0).isEqualTo(_1966_v36.f9)) ? ((((_1973_v43).contains(_dafny.Seq.update((_this).f0, _module.__default.safeIndex(new BigNumber((_1976_v44).cardinality()), new BigNumber(((_this).f0).length)), _1966_v36.f9))) ? ((_1973_v43).get(_dafny.Seq.update((_this).f0, _module.__default.safeIndex(new BigNumber((_1976_v44).cardinality()), new BigNumber(((_this).f0).length)), _1966_v36.f9))) : (_1972_v42))) : (_dafny.Seq.Concat(_1972_v42, _1972_v42))), _module.__default.safeIndex((_1966_v36).fm6(p0, _1966_v36.f9, globalState), new BigNumber(((((p0).isEqualTo(_1966_v36.f9)) ? ((((_1973_v43).contains(_dafny.Seq.update((_this).f0, _module.__default.safeIndex(new BigNumber((_1976_v44).cardinality()), new BigNumber(((_this).f0).length)), _1966_v36.f9))) ? ((_1973_v43).get(_dafny.Seq.update((_this).f0, _module.__default.safeIndex(new BigNumber((_1976_v44).cardinality()), new BigNumber(((_this).f0).length)), _1966_v36.f9))) : (_1972_v42))) : (_dafny.Seq.Concat(_1972_v42, _1972_v42)))).length)), _1968_v38), _module.__default.safeIndex(new BigNumber(581), new BigNumber((_dafny.Seq.update((((p0).isEqualTo(_1966_v36.f9)) ? ((((_1973_v43).contains(_dafny.Seq.update((_this).f0, _module.__default.safeIndex(new BigNumber((_1976_v44).cardinality()), new BigNumber(((_this).f0).length)), _1966_v36.f9))) ? ((_1973_v43).get(_dafny.Seq.update((_this).f0, _module.__default.safeIndex(new BigNumber((_1976_v44).cardinality()), new BigNumber(((_this).f0).length)), _1966_v36.f9))) : (_1972_v42))) : (_dafny.Seq.Concat(_1972_v42, _1972_v42))), _module.__default.safeIndex((_1966_v36).fm6(p0, _1966_v36.f9, globalState), new BigNumber(((((p0).isEqualTo(_1966_v36.f9)) ? ((((_1973_v43).contains(_dafny.Seq.update((_this).f0, _module.__default.safeIndex(new BigNumber((_1976_v44).cardinality()), new BigNumber(((_this).f0).length)), _1966_v36.f9))) ? ((_1973_v43).get(_dafny.Seq.update((_this).f0, _module.__default.safeIndex(new BigNumber((_1976_v44).cardinality()), new BigNumber(((_this).f0).length)), _1966_v36.f9))) : (_1972_v42))) : (_dafny.Seq.Concat(_1972_v42, _1972_v42)))).length)), _1968_v38)).length)), (_this).fm9(_1977_v45, p0, _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements(true)), _dafny.Seq.of(true), globalState));
      } else {
        let _1978_v46;
        _1978_v46 = new BigNumber(842);
        _1978_v46 = ((_dafny.ZERO).minus(p0)).minus(_1978_v46);
        let _1979_v47;
        _1979_v47 = _module.D4.create_DC10((_1978_v46).plus(p0), _1978_v46, _1978_v46);
        _1979_v47 = _1979_v47;
        let _1980_v48;
        let _nw309 = new _module.C2();
        _nw309.__ctor();
        _1980_v48 = _nw309;
        let _1981_v49;
        _1981_v49 = _module.D16.create_DC44((_dafny.ZERO).minus(p0));
        let _1982_v50;
        let _nw310 = new _module.C2();
        _nw310.__ctor();
        _1982_v50 = _nw310;
        let _rhs213 = _1981_v49;
        let _rhs214 = ((_1921_v0) ? (_1982_v50) : (_1982_v50));
        _1981_v49 = _rhs213;
        _1982_v50 = _rhs214;
        let _1983_v51;
        _1983_v51 = _dafny.Seq.UnicodeFromString("ru");
        let _1984_v52;
        _1984_v52 = _module.D14.create_DC35(_module.__default.safeDivisionInt(p0, _1978_v46), ((true) ? (_1983_v51) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-598)), function (_1985_i6) {
  return new _dafny.CodePoint('w'.codePointAt(0));
}))), ((((_this).f4).contains(_1978_v46)) ? (((_this).f4).get(_1978_v46)) : (_1978_v46)), true, (p0).isLessThanOrEqualTo(p0));
        let _source21 = _1984_v52;
        if (_source21.is_DC35) {
          let _1986___mcc_h0 = (_source21).cf63;
          let _1987___mcc_h1 = (_source21).cf64;
          let _1988___mcc_h2 = (_source21).cf65;
          let _1989___mcc_h3 = (_source21).cf66;
          let _1990___mcc_h4 = (_source21).cf67;
          let _1991_cf67 = _1990___mcc_h4;
          let _1992_cf66 = _1989___mcc_h3;
          let _1993_cf65 = _1988___mcc_h2;
          let _1994_cf64 = _1987___mcc_h1;
          let _1995_cf63 = _1986___mcc_h0;
          _1921_v0 = !(!(_1991_cf67) || (!(_1992_cf66) || (true)));
          let _1996_v53;
          let _nw311 = new _module.C0();
          _nw311.__ctor(_1994_cf64);
          _1996_v53 = _nw311;
          _1995_cf63 = ((_this).f0)[_module.__default.safeIndex(_1995_cf63, new BigNumber(((_this).f0).length))];
          let _1997_v54;
          _1997_v54 = new _dafny.CodePoint('m'.codePointAt(0));
          _1997_v54 = _1997_v54;
        } else if (_source21.is_DC36) {
          let _1998___mcc_h5 = (_source21).cf68;
          let _1999___mcc_h6 = (_source21).cf69;
          let _2000_cf69 = _1999___mcc_h6;
          let _2001_cf68 = _1998___mcc_h5;
          _1921_v0 = _1921_v0;
          let _2002_v55;
          _2002_v55 = _dafny.Set.fromElements(_1921_v0);
          _1921_v0 = (_2002_v55).IsDisjointFrom(_dafny.Set.fromElements(_1921_v0));
          let _2003_v56;
          let _nw312 = new _module.C4();
          _nw312.__ctor();
          _2003_v56 = _nw312;
          let _2004_v57;
          _2004_v57 = _dafny.Seq.of(true, _1921_v0);
          let _2005_v58;
          _2005_v58 = _dafny.Map.Empty.slice().updateUnsafe((((_2004_v57)[_module.__default.safeIndex((((_2001_cf68).contains(p0)) ? ((_2001_cf68).get(p0)) : (p0)), new BigNumber((_2004_v57).length))]) ? (_1932_v7) : (_dafny.Set.fromElements(_1978_v46))),_1921_v0);
          let _2006_v59;
          _2006_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1921_v0,false);
          _2005_v58 = (_2005_v58).update(_dafny.Set.fromElements(new BigNumber(602), _module.__default.fm20(!(_1921_v0), globalState), _1978_v46, new BigNumber((_2006_v59).length), p0), _1921_v0);
        } else {
          let _2007___mcc_h7 = (_source21).cf62;
          let _2008_cf62 = _2007___mcc_h7;
          let _index311 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((p1).length));
          (p1)[_index311] = _1978_v46;
          let _2009_v60;
          _2009_v60 = new _dafny.CodePoint('y'.codePointAt(0));
          let _2010_v61;
          let _init48 = ((_2011_v0) => function (_2012_i7) {
            return _2011_v0;
          })(_1921_v0);
          let _nw313 = Array((new BigNumber(15)).toNumber());
          for (let _i0_48 = 0; _i0_48 < new BigNumber(_nw313.length); _i0_48++) {
            _nw313[_i0_48] = _init48(new BigNumber(_i0_48));
          }
          _2010_v61 = _nw313;
          let _2013_v62;
          _2013_v62 = _module.D7.create_DC18();
          let _2014_v63;
          _2014_v63 = _dafny.Map.Empty.slice().updateUnsafe(true,_1921_v0);
          let _index312 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((p1).length));
          let _rhs215 = !(!(!((_module.__default.fm48(_2013_v62, new BigNumber((_2014_v63).length), p0, p0, globalState)).IsSubsetOf((_this).f4))));
          let _rhs216 = _1978_v46;
          let _rhs217 = _1921_v0;
          let _rhs218 = _2009_v60;
          let _rhs219 = _2010_v61;
          let _lhs122 = p1;
          let _lhs123 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((p1).length));
          _1921_v0 = _rhs215;
          _lhs122[_lhs123] = _rhs216;
          _1921_v0 = _rhs217;
          _2009_v60 = _rhs218;
          _2010_v61 = _rhs219;
          let _2015_v64;
          let _nw314 = new _module.C6();
          _nw314.__ctor(_1978_v46);
          _2015_v64 = _nw314;
          let _index313 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2010_v61).length));
          (_2010_v61)[_index313] = _1921_v0;
          let _2016_v65;
          _2016_v65 = _dafny.MultiSet.fromElements(!(_1921_v0));
          let _2017_v66;
          _2017_v66 = _dafny.Seq.of(_1983_v51);
          let _index314 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2010_v61).length));
          let _index315 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((p1).length));
          let _rhs220 = ((_2016_v65).Union(_2016_v65)).equals(_2016_v65);
          let _rhs221 = (_2015_v64).f6;
          let _rhs222 = _2010_v61;
          let _rhs223 = (_2017_v66)[_module.__default.safeIndex(p0, new BigNumber((_2017_v66).length))];
          let _lhs124 = _2010_v61;
          let _lhs125 = _module.__default.safeIndex(new BigNumber(747), new BigNumber((_2010_v61).length));
          let _lhs126 = p1;
          let _lhs127 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((p1).length));
          _lhs124[_lhs125] = _rhs220;
          _lhs126[_lhs127] = _rhs221;
          _2010_v61 = _rhs222;
          _1983_v51 = _rhs223;
          let _index316 = _module.__default.safeIndex(new BigNumber(885), new BigNumber((p1).length));
          (p1)[_index316] = new BigNumber(-723);
        }
      }
      let _hi14 = p0;
      for (let _2018_i8 = (p0).plus(p0); _2018_i8.isLessThan(_hi14); _2018_i8 = _2018_i8.plus(_dafny.ONE)) {
        let _2019_v67;
        _2019_v67 = _dafny.Map.Empty.slice().updateUnsafe(_2018_i8,_1921_v0);
        let _2020_v68;
        _2020_v68 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2019_v67);
        let _2021_v69;
        _2021_v69 = new BigNumber(611);
        let _rhs224 = _2020_v68;
        let _rhs225 = (p0).minus((_2018_i8).plus((_dafny.ZERO).minus(p0)));
        _2020_v68 = _rhs224;
        _2021_v69 = _rhs225;
        if (_1921_v0) {
          let _2022_v70;
          _2022_v70 = _dafny.Seq.UnicodeFromString("ashcuj");
          let _2023_v71;
          let _nw315 = new _module.C0();
          _nw315.__ctor(_2022_v70);
          _2023_v71 = _nw315;
          _2021_v69 = _2018_i8;
          let _2024_v72;
          let _nw316 = Array((new BigNumber(8)).toNumber()).fill(false);
          _2024_v72 = _nw316;
          let _index317 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2024_v72).length));
          (_2024_v72)[_index317] = _1921_v0;
          let _index318 = _module.__default.safeIndex(new BigNumber(891), new BigNumber((_2024_v72).length));
          (_2024_v72)[_index318] = _1921_v0;
          _2021_v69 = _module.__default.safeModuloInt(p0, _2021_v69);
          let _2025_v73;
          _2025_v73 = _dafny.Map.Empty.slice().updateUnsafe(false,p0);
          let _2026_v74;
          _2026_v74 = _dafny.Seq.of(new BigNumber(((_2025_v73).Merge(_dafny.Map.Empty.slice().updateUnsafe((_2024_v72)[_module.__default.safeIndex(new BigNumber(891), new BigNumber((_2024_v72).length))],_2018_i8))).length), _2021_v69);
          _2026_v74 = _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(11)), ((_2027_i8) => function (_2028_i9) {
            return _2027_i8;
          })(_2018_i8)), _module.__default.safeIndex(_2021_v69, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(11)), ((_2029_i8) => function (_2030_i9) {
            return _2029_i8;
          })(_2018_i8))).length)), _2021_v69);
        } else {
          let _2031_v75;
          _2031_v75 = new _dafny.CodePoint('m'.codePointAt(0));
          let _2032_v76;
          _2032_v76 = _dafny.Seq.UnicodeFromString("unqvelq");
          let _2033_v77;
          let _nw317 = new _module.C0();
          _nw317.__ctor(_2032_v76);
          _2033_v77 = _nw317;
          let _2034_v78;
          _2034_v78 = _dafny.Set.fromElements(_2033_v77, _2033_v77, _2033_v77);
          let _2035_v79;
          let _nw318 = Array((new BigNumber(8)).toNumber()).fill([]);
          _2035_v79 = _nw318;
          let _2036_v80;
          _2036_v80 = _module.D2.create_DC4(new BigNumber((_2033_v77.f8).length), _1921_v0, _2035_v79);
          let _2037_v81;
          _2037_v81 = _module.D2.create_DC6(_2036_v80);
          let _2038_v82;
          _2038_v82 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2032_v76).length),_2037_v81);
          let _2039_v83;
          _2039_v83 = _module.D4.create_DC9((_this).f4, _2031_v75, _2021_v69, _2034_v78, new BigNumber((_2038_v82).length));
          let _2040_v84;
          _2040_v84 = _dafny.Map.Empty.slice().updateUnsafe(_1921_v0,_2033_v77.f8);
          let _2041_v85;
          _2041_v85 = _module.D12.create_DC29(_1921_v0, p0, _1921_v0, new BigNumber(((((_2040_v84).contains(_1921_v0)) ? ((_2040_v84).get(_1921_v0)) : (_2032_v76))).length), _1921_v0);
          let _pat_let_tv49 = _2033_v77;
          let _2042_v86;
          _2042_v86 = _dafny.MultiSet.fromElements(_2041_v85, _2041_v85, function (_pat_let42_0) {
            return function (_2043_dt__update__tmp_h1) {
              return function (_pat_let43_0) {
                return function (_2044_dt__update_hcf49_h0) {
                  return _module.D12.create_DC29((_2043_dt__update__tmp_h1).dtor_cf48, _2044_dt__update_hcf49_h0, (_2043_dt__update__tmp_h1).dtor_cf50, (_2043_dt__update__tmp_h1).dtor_cf51, (_2043_dt__update__tmp_h1).dtor_cf52);
                }(_pat_let43_0);
              }(new BigNumber((_pat_let_tv49.f8).length));
            }(_pat_let42_0);
          }(_2041_v85));
          let _rhs226 = _1921_v0;
          let _rhs227 = p0;
          let _rhs228 = _module.D4.create_DC9((_this).f4, _2031_v75, (((_2042_v86).contains(_2041_v85)) ? ((_2042_v86).get(_2041_v85)) : (p0)), _2034_v78, _module.__default.safeDivisionInt(_2021_v69, (_dafny.ZERO).minus(_2018_i8)));
          let _rhs229 = _1921_v0;
          _1921_v0 = _rhs226;
          _2021_v69 = _rhs227;
          _2039_v83 = _rhs228;
          _1921_v0 = _rhs229;
          _2021_v69 = _2021_v69;
          let _2045_v90;
          _2045_v90 = _module.D17.create_DC45(_dafny.MultiSet.fromElements((_this).f4));
          let _rhs230 = _2032_v76;
          let _rhs231 = _module.__default.safeDivisionInt(new BigNumber(((function () {
            let _coll76 = new _dafny.Map();
            for (const _compr_76 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(-905)), function (_2046_i10) {
              return (_this).f4;
            })).Elements) {
              let _2047_v87 = _compr_76;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-905)), function (_2046_i10) {
                return (_this).f4;
              }), _2047_v87)) {
                _coll76.push([_2047_v87,_1921_v0]);
              }
            }
            return _coll76;
          }()).Merge(function () {
            let _coll77 = new _dafny.Map();
            for (const _compr_77 of (function () {
              let _coll78 = new _dafny.Map();
              for (const _compr_78 of ((_2045_v90).dtor_cf83).Elements) {
                let _2048_v89 = _compr_78;
                if (((_2045_v90).dtor_cf83).contains(_2048_v89)) {
                  _coll78.push([_2048_v89,_dafny.Set.fromElements(_1921_v0, _1921_v0)]);
                }
              }
              return _coll78;
            }()).Keys.Elements) {
              let _2049_v88 = _compr_77;
              if ((function () {
                let _coll79 = new _dafny.Map();
                for (const _compr_79 of ((_2045_v90).dtor_cf83).Elements) {
                  let _2048_v89 = _compr_79;
                  if (((_2045_v90).dtor_cf83).contains(_2048_v89)) {
                    _coll79.push([_2048_v89,_dafny.Set.fromElements(_1921_v0, _1921_v0)]);
                  }
                }
                return _coll79;
              }()).contains(_2049_v88)) {
                _coll77.push([_2049_v88,_1921_v0]);
              }
            }
            return _coll77;
          }())).length), p0);
          let _rhs232 = p0;
          _2032_v76 = _rhs230;
          _2021_v69 = _rhs231;
          _2021_v69 = _rhs232;
          let _2050_v91;
          _2050_v91 = _dafny.Seq.of((((_2019_v67).contains(new BigNumber(((_this).f0).length))) ? ((_2019_v67).get(new BigNumber(((_this).f0).length))) : (_1921_v0)), _1921_v0, _1921_v0);
          _1921_v0 = !(((_1921_v0) ? ((_this).fm6(_2021_v69, new BigNumber((_module.__default.fm24(_1921_v0, globalState)).cardinality()), globalState)) : (new BigNumber((_2050_v91).length)))).isEqualTo(p0);
          let _2051_v92;
          let _init49 = ((_2052_v77, _2053_v76) => function (_2054_i11) {
            return !_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("tvmhkbqyi"), (_2052_v77.f8)[_module.__default.safeIndex(new BigNumber((_2053_v76).length), new BigNumber((_2052_v77.f8).length))]);
          })(_2033_v77, _2032_v76);
          let _nw319 = Array((new BigNumber(5)).toNumber());
          for (let _i0_49 = 0; _i0_49 < new BigNumber(_nw319.length); _i0_49++) {
            _nw319[_i0_49] = _init49(new BigNumber(_i0_49));
          }
          _2051_v92 = _nw319;
          let _index319 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_2051_v92).length));
          (_2051_v92)[_index319] = _1921_v0;
          let _2055_v93;
          let _nw320 = Array((new BigNumber(3)).toNumber()).fill([]);
          _2055_v93 = _nw320;
          let _index320 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_2051_v92).length));
          let _rhs233 = false;
          let _rhs234 = _1921_v0;
          let _rhs235 = _dafny.Seq.UnicodeFromString("yyco");
          let _rhs236 = _module.__default.fm30(_module.__default.safeModuloInt(p0, _2018_i8), _1921_v0, (true) && (_1921_v0), globalState);
          let _rhs237 = _2055_v93;
          let _lhs128 = _2051_v92;
          let _lhs129 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_2051_v92).length));
          let _lhs130 = _2033_v77;
          _lhs128[_lhs129] = _rhs233;
          _1921_v0 = _rhs234;
          _lhs130.f8 = _rhs235;
          _1921_v0 = _rhs236;
          _2055_v93 = _rhs237;
        }
        let _2056_v94;
        _2056_v94 = _module.D4.create_DC8(_dafny.Seq.of(_1930_v5));
        _2056_v94 = _2056_v94;
        _1921_v0 = _1921_v0;
      }
      let _2057_v95;
      _2057_v95 = _dafny.Map.Empty.slice().updateUnsafe(_1921_v0,p0);
      let _2058_v96;
      _2058_v96 = _dafny.Map.Empty.slice().updateUnsafe(_2057_v95,_dafny.Seq.UnicodeFromString("fakxfloyk"));
      let _2059_v97;
      let _nw321 = Array((new BigNumber(3)).toNumber());
      _nw321[(_dafny.ZERO).toNumber()] = _2058_v96;
      _nw321[(_dafny.ONE).toNumber()] = (_2058_v96).Merge(_2058_v96);
      _nw321[(new BigNumber(2)).toNumber()] = (_2058_v96).Merge(_2058_v96);
      _2059_v97 = _nw321;
      let _2060_v98;
      _2060_v98 = _module.D18.create_DC47(_2058_v96);
      let _index321 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_2059_v97).length));
      (_2059_v97)[_index321] = (_2060_v98).dtor_cf85;
      let _index322 = _module.__default.safeIndex(new BigNumber(894), new BigNumber((_2059_v97).length));
      (_2059_v97)[_index322] = _2058_v96;
      let _2061_v99;
      _2061_v99 = _dafny.Map.Empty.slice().updateUnsafe(false,_1921_v0);
      let _2062_v100;
      _2062_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(982),p0);
      let _2063_v101;
      _2063_v101 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_2061_v99).length)), _2062_v100, _2062_v100);
      let _2064_v102;
      _2064_v102 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(816),_1921_v0);
      let _2065_v103;
      _2065_v103 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((((_2064_v102).contains(p0)) ? ((_2064_v102).get(p0)) : (_1921_v0)))).length),_1921_v0);
      let _rhs238 = (_2063_v101).IsProperSubsetOf(_2063_v101);
      let _rhs239 = (_2065_v103).contains((p0).minus((_dafny.ZERO).minus(p0)));
      let _rhs240 = (((_1921_v0) && (_1921_v0)) ? (_1921_v0) : (true));
      _1921_v0 = _rhs238;
      _1921_v0 = _rhs239;
      _1921_v0 = _rhs240;
      let _2066_i12;
      _2066_i12 = _dafny.ZERO;
      L8: {
        while (_1921_v0) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2066_i12)) {
              break L8;
            }
            _2066_i12 = (_2066_i12).plus(_dafny.ONE);
            let _2067_v104;
            _2067_v104 = _dafny.Seq.of(_1921_v0);
            let _2068_v105;
            _2068_v105 = _module.D7.create_DC17(_2067_v104);
            let _2069_v106;
            _2069_v106 = _module.D4.create_DC11(p0, !_dafny.areEqual((_2068_v105).dtor_cf34, _2067_v104), true, !(_1921_v0), _1921_v0);
            _2069_v106 = _2069_v106;
            let _2070_v107;
            _2070_v107 = new BigNumber(399);
            _2070_v107 = _2070_v107;
            let _2071_v108;
            _2071_v108 = _module.D2.create_DC5(_2070_v107, (_this).f0, _1921_v0);
            let _2072_v109;
            let _nw322 = new _module.C5();
            _nw322.__ctor(p1, (_2071_v108).dtor_cf9, _this.f1);
            _2072_v109 = _nw322;
            let _2073_v110;
            _2073_v110 = _module.D12.create_DC29(_1921_v0, _2070_v107, _1921_v0, new BigNumber(331), false);
            let _2074_v111;
            _2074_v111 = _dafny.Seq.of(_2072_v109, _2072_v109, _2072_v109, (((_2073_v110).dtor_cf52) ? (_2072_v109) : (_2072_v109)));
            _2072_v109 = (_2074_v111)[_module.__default.safeIndex(((_this).fm6(_2070_v107, _2070_v107, globalState)).multipliedBy(_2070_v107), new BigNumber((_2074_v111).length))];
            let _2075_v112;
            let _nw323 = new _module.C2();
            _nw323.__ctor();
            _2075_v112 = _nw323;
          }
        }
      }
      return;
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
  };

  $module.C8 = class C8 {
    constructor () {
      this._tname = "_module.C8";
      this._f1 = _module.D0.Default();
      this._f0 = _dafny.Seq.of();
      this._f2 = _dafny.ZERO;
      this._f3 = _module.D0.Default();
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    set f1(value) {
      let _this = this;
      _this._f1 = value;
    };
    get f0() {
      let _this = this;
      return _this._f0;
    };
    __ctor(f2, f3, f0, f1) {
      let _this = this;
      (_this)._f2 = f2;
      (_this)._f3 = f3;
      (_this)._f0 = f0;
      (_this)._f1 = f1;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !((_dafny.MultiSet.FromArray(_dafny.Seq.Create(_module.__default.abs(new BigNumber(191)), function (_2076_i0) {
        return (_this).f2;
      }))).contains(((_this).f0)[_module.__default.safeIndex((_this).f2, new BigNumber(((_this).f0).length))]));
    };
    fm6(p0, p1, globalState) {
      let _this = this;
      if (((_this).f2).isLessThanOrEqualTo(new BigNumber(((_this).f0).length))) {
        return (_this).f2;
      } else {
        return new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,(_this).f2), _dafny.Map.Empty.slice().updateUnsafe(!(false),(_this).f2), _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f2), _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f2))).length);
      }
    };
    fm7(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uyousyj"), _dafny.Seq.UnicodeFromString("b"));
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = false;
      let _2077_v0;
      _2077_v0 = new _dafny.CodePoint('d'.codePointAt(0));
      let _2078_v1;
      _2078_v1 = _dafny.Map.Empty.slice().updateUnsafe(_2077_v0,p0);
      let _2079_v2;
      _2079_v2 = true;
      let _2080_v3;
      _2080_v3 = _dafny.Seq.UnicodeFromString("smiohn");
      let _2081_v5;
      _2081_v5 = _dafny.Map.Empty.slice().updateUnsafe(_2077_v0,_2079_v2);
      let _2082_v6;
      let _nw324 = Array((new BigNumber(25)).toNumber());
      _nw324[(_dafny.ZERO).toNumber()] = (_2078_v1).update(_2077_v0, (_this).f2);
      _nw324[(_dafny.ONE).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(2)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2077_v0,new BigNumber(409));
      _nw324[(new BigNumber(4)).toNumber()] = _module.__default.fm52(_2079_v2, _module.D15.create_DC38(_module.__default.fm51(true, new BigNumber((_2080_v3).length), new BigNumber(535), globalState)), _2079_v2, globalState);
      _nw324[(new BigNumber(5)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(6)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(7)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2077_v0,new BigNumber(237));
      _nw324[(new BigNumber(8)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(9)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(10)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2077_v0,(_this).f2);
      _nw324[(new BigNumber(11)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(12)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),(_this).f2)).update(_2077_v0, (_dafny.ZERO).minus(p0));
      _nw324[(new BigNumber(13)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(14)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(15)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(16)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(17)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2077_v0,p0);
      _nw324[(new BigNumber(18)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(19)).toNumber()] = function () {
        let _coll80 = new _dafny.Map();
        for (const _compr_80 of (_2081_v5).Keys.Elements) {
          let _2083_v4 = _compr_80;
          if ((_2081_v5).contains(_2083_v4)) {
            _coll80.push([_2083_v4,p0]);
          }
        }
        return _coll80;
      }();
      _nw324[(new BigNumber(20)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm29(_2079_v2, globalState),new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(565)), ((_2084_v0) => function (_2085_i0) {
        return _2084_v0;
      })(_2077_v0)), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(565)), ((_2086_v0) => function (_2087_i0) {
        return _2086_v0;
      })(_2077_v0))).length)), _2077_v0)).length));
      _nw324[(new BigNumber(22)).toNumber()] = (_2078_v1).update(_2077_v0, (_this).f2);
      _nw324[(new BigNumber(23)).toNumber()] = _2078_v1;
      _nw324[(new BigNumber(24)).toNumber()] = _2078_v1;
      _2082_v6 = _nw324;
      let _2088_v7;
      let _nw325 = new _module.C1();
      _nw325.__ctor((_this).f2, _2082_v6, ((_2079_v2) ? ((_this).f0) : ((_this).f0)), _this.f1);
      _2088_v7 = _nw325;
      let _2089_v8;
      let _init50 = ((_2090_p0) => function (_2091_i1) {
        return _module.__default.safeDivisionInt(_2091_i1, _2090_p0);
      })(p0);
      let _nw326 = Array((new BigNumber(7)).toNumber());
      for (let _i0_50 = 0; _i0_50 < new BigNumber(_nw326.length); _i0_50++) {
        _nw326[_i0_50] = _init50(new BigNumber(_i0_50));
      }
      _2089_v8 = _nw326;
      let _2092_v9;
      _2092_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2080_v3,_2089_v8);
      let _2093_v10;
      _2093_v10 = _dafny.Map.Empty.slice().updateUnsafe((((_2092_v9).contains(_2080_v3)) ? ((_2092_v9).get(_2080_v3)) : (_2089_v8)),(p0).plus(p0));
      _2093_v10 = (_2093_v10).update(_2089_v8, new BigNumber((_2081_v5).length));
      let _2094_v11;
      _2094_v11 = _dafny.Seq.of(_2079_v2, false);
      let _2095_v12;
      _2095_v12 = _dafny.Seq.of(_2094_v11);
      _2094_v11 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2094_v11, (_2095_v12)[_module.__default.safeIndex(p0, new BigNumber((_2095_v12).length))]), _2094_v11);
      let _2096_v13;
      _2096_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2079_v2);
      let _2097_v14;
      _2097_v14 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2096_v13).length),_2088_v7.f9);
      let _2098_v15;
      _2098_v15 = _dafny.Set.fromElements((_this).f2);
      let _2099_v16;
      _2099_v16 = _dafny.Map.Empty.slice().updateUnsafe(_2079_v2,_dafny.MultiSet.fromElements((_this).f2, new BigNumber((_2097_v14).length), (_dafny.ZERO).minus(new BigNumber((_2098_v15).length))));
      let _2100_v17;
      _2100_v17 = _dafny.MultiSet.fromElements(_module.__default.fm44(_2079_v2, globalState));
      let _hi15 = (_module.__default.fm44(_2079_v2, globalState)).plus(p0);
      for (let _2101_i2 = new BigNumber(((((_2099_v16).contains(_2079_v2)) ? ((_2099_v16).get(_2079_v2)) : (_2100_v17))).cardinality()); _2101_i2.isLessThan(_hi15); _2101_i2 = _2101_i2.plus(_dafny.ONE)) {
        _2079_v2 = false;
        if (_2079_v2) {
          let _2102_v18;
          let _nw327 = Array((new BigNumber(3)).toNumber());
          _nw327[(_dafny.ZERO).toNumber()] = _2080_v3;
          _nw327[(_dafny.ONE).toNumber()] = _2080_v3;
          _nw327[(new BigNumber(2)).toNumber()] = _2080_v3;
          _2102_v18 = _nw327;
          let _index323 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_2102_v18).length));
          (_2102_v18)[_index323] = _2080_v3;
          let _index324 = _module.__default.safeIndex(new BigNumber(389), new BigNumber((_2102_v18).length));
          (_2102_v18)[_index324] = _dafny.Seq.Concat(_2080_v3, _2080_v3);
          let _2103_v19;
          _2103_v19 = _dafny.Map.Empty.slice().updateUnsafe(_2079_v2,!(_2079_v2));
          (_2088_v7).f9 = (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,true), _dafny.Map.Empty.slice().updateUnsafe(_2079_v2,true), ((_dafny.Map.Empty.slice().updateUnsafe(!(_2079_v2),_2079_v2)).update(true, _2079_v2)).Merge(_2103_v19), _2103_v19, _2103_v19)).cardinality()));
          let _2104_v20;
          let _nw328 = Array((new BigNumber(6)).toNumber());
          _nw328[(_dafny.ZERO).toNumber()] = _2096_v13;
          _nw328[(_dafny.ONE).toNumber()] = (_2096_v13).update((_2088_v7).fm28(_2088_v7.f9, _2079_v2, globalState), _2079_v2);
          _nw328[(new BigNumber(2)).toNumber()] = _2096_v13;
          _nw328[(new BigNumber(3)).toNumber()] = _2096_v13;
          _nw328[(new BigNumber(4)).toNumber()] = _2096_v13;
          _nw328[(new BigNumber(5)).toNumber()] = _module.__default.fm50(false, _2079_v2, _2101_i2, globalState);
          _2104_v20 = _nw328;
          let _2105_v21;
          let _nw329 = Array((new BigNumber(8)).toNumber());
          _nw329[(_dafny.ZERO).toNumber()] = _2104_v20;
          _nw329[(_dafny.ONE).toNumber()] = _2104_v20;
          _nw329[(new BigNumber(2)).toNumber()] = _2104_v20;
          _nw329[(new BigNumber(3)).toNumber()] = _2104_v20;
          _nw329[(new BigNumber(4)).toNumber()] = _2104_v20;
          _nw329[(new BigNumber(5)).toNumber()] = _2104_v20;
          _nw329[(new BigNumber(6)).toNumber()] = _2104_v20;
          _nw329[(new BigNumber(7)).toNumber()] = _2104_v20;
          _2105_v21 = _nw329;
          let _index325 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_2105_v21).length));
          (_2105_v21)[_index325] = _2104_v20;
          let _index326 = _module.__default.safeIndex(new BigNumber(775), new BigNumber((_2105_v21).length));
          (_2105_v21)[_index326] = _2104_v20;
          r0 = _2079_v2;
          let _2106_v22;
          _2106_v22 = _dafny.Set.fromElements(_2079_v2);
          let _2107_v23;
          _2107_v23 = _dafny.Map.Empty.slice().updateUnsafe(_2079_v2,new BigNumber((_2106_v22).length));
          r0 = !(_2107_v23).contains(_2079_v2);
        } else {
          let _2108_v24;
          let _nw330 = Array((new BigNumber(24)).toNumber()).fill([]);
          _2108_v24 = _nw330;
          let _2109_v25;
          let _init51 = ((_2110_v2) => function (_2111_i3) {
            return _2110_v2;
          })(_2079_v2);
          let _nw331 = Array((new BigNumber(23)).toNumber());
          for (let _i0_51 = 0; _i0_51 < new BigNumber(_nw331.length); _i0_51++) {
            _nw331[_i0_51] = _init51(new BigNumber(_i0_51));
          }
          _2109_v25 = _nw331;
          let _index327 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_2108_v24).length));
          (_2108_v24)[_index327] = _2109_v25;
          let _2112_v26;
          let _nw332 = Array((new BigNumber(2)).toNumber()).fill(_module.D14.Default());
          _2112_v26 = _nw332;
          let _2113_v27;
          _2113_v27 = _module.D14.create_DC35(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(630)), ((_2114_v7) => function (_2115_i4) {
  return (_dafny.ZERO).minus(_2114_v7.f9);
})(_2088_v7))).length), _2080_v3, _2101_i2, true, _2079_v2);
          let _pat_let_tv50 = _2077_v0;
          let _pat_let_tv51 = _2096_v13;
          let _pat_let_tv52 = _2088_v7;
          let _pat_let_tv53 = _2079_v2;
          let _pat_let_tv54 = _2096_v13;
          let _pat_let_tv55 = _2088_v7;
          let _index328 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_2112_v26).length));
          (_2112_v26)[_index328] = function (_pat_let44_0) {
            return function (_2116_dt__update__tmp_h0) {
              return function (_pat_let45_0) {
                return function (_2119_dt__update_hcf64_h0) {
                  return function (_pat_let46_0) {
                    return function (_2120_dt__update_hcf67_h0) {
                      return _module.D14.create_DC35((_2116_dt__update__tmp_h0).dtor_cf63, _2119_dt__update_hcf64_h0, (_2116_dt__update__tmp_h0).dtor_cf65, (_2116_dt__update__tmp_h0).dtor_cf66, _2120_dt__update_hcf67_h0);
                    }(_pat_let46_0);
                  }((((_pat_let_tv54).contains(_pat_let_tv55.f9)) ? ((_pat_let_tv51).get(_pat_let_tv52.f9)) : (_pat_let_tv53)));
                }(_pat_let45_0);
              }(_dafny.Seq.Create(_module.__default.abs(new BigNumber(892)), ((_2117_v0) => function (_2118_i5) {
                return _2117_v0;
              })(_pat_let_tv50)));
            }(_pat_let44_0);
          }(_2113_v27);
          let _index329 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_2109_v25).length));
          (_2109_v25)[_index329] = _2079_v2;
          let _index330 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_2108_v24).length));
          let _index331 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_2112_v26).length));
          let _index332 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_2109_v25).length));
          let _rhs241 = _2109_v25;
          let _rhs242 = _2113_v27;
          let _rhs243 = _2079_v2;
          let _rhs244 = (_this).f2;
          let _rhs245 = _module.__default.fm17(globalState);
          let _lhs131 = _2108_v24;
          let _lhs132 = _module.__default.safeIndex(new BigNumber(748), new BigNumber((_2108_v24).length));
          let _lhs133 = _2112_v26;
          let _lhs134 = _module.__default.safeIndex(new BigNumber(606), new BigNumber((_2112_v26).length));
          let _lhs135 = _2109_v25;
          let _lhs136 = _module.__default.safeIndex(new BigNumber(581), new BigNumber((_2109_v25).length));
          let _lhs137 = _2088_v7;
          _lhs131[_lhs132] = _rhs241;
          _lhs133[_lhs134] = _rhs242;
          _lhs135[_lhs136] = _rhs243;
          _lhs137.f9 = _rhs244;
          r1 = _rhs245;
          (_2088_v7).f9 = ((_2079_v2) ? (new BigNumber(((_module.__default.fm13(_2077_v0, true, globalState)).Intersect(_2100_v17)).cardinality())) : (_2088_v7.f9));
          let _2121_v28;
          let _2122_v29;
          let _2123_v30;
          let _out49;
          let _out50;
          let _out51;
          let _outcollector16 = (_this).m3(_2101_i2, (_this).f3, _module.__default.fm30(p0, (_2109_v25)[_module.__default.safeIndex(new BigNumber(581), new BigNumber((_2109_v25).length))], (_module.D18.create_DC50((_2109_v25)[_module.__default.safeIndex(new BigNumber(581), new BigNumber((_2109_v25).length))], _2109_v25, _2079_v2, (_2109_v25)[_module.__default.safeIndex(new BigNumber(581), new BigNumber((_2109_v25).length))], _2080_v3)).dtor_cf90, globalState), globalState);
          _out49 = _outcollector16[0];
          _out50 = _outcollector16[1];
          _out51 = _outcollector16[2];
          _2121_v28 = _out49;
          _2122_v29 = _out50;
          _2123_v30 = _out51;
          (_2088_v7).f9 = _module.__default.safeModuloInt((new BigNumber(811)).minus(((_this).f0)[_module.__default.safeIndex(_2101_i2, new BigNumber(((_this).f0).length))]), (((_2100_v17).contains(_2121_v28)) ? ((_2100_v17).get(_2121_v28)) : (_2088_v7.f9)));
          (_2088_v7).f9 = (_this).fm6(_2101_i2, (_dafny.ZERO).minus(((_2079_v2) ? (p0) : (new BigNumber(877)))), globalState);
        }
        let _2124_v31;
        let _init52 = ((_2125_v0) => function (_2126_i6) {
          return _2125_v0;
        })(_2077_v0);
        let _nw333 = Array((new BigNumber(26)).toNumber());
        for (let _i0_52 = 0; _i0_52 < new BigNumber(_nw333.length); _i0_52++) {
          _nw333[_i0_52] = _init52(new BigNumber(_i0_52));
        }
        _2124_v31 = _nw333;
        let _index333 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_2124_v31).length));
        (_2124_v31)[_index333] = _2077_v0;
        let _index334 = _module.__default.safeIndex(new BigNumber(149), new BigNumber((_2124_v31).length));
        (_2124_v31)[_index334] = _2077_v0;
        let _2127_v32;
        _2127_v32 = _dafny.Map.Empty.slice().updateUnsafe(_2079_v2,_2080_v3);
        let _2128_v33;
        _2128_v33 = _dafny.Seq.of(_dafny.Seq.update((((_2127_v32).contains(_2079_v2)) ? ((_2127_v32).get(_2079_v2)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(526)), function (_2129_i7) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }))), _module.__default.safeIndex(_2088_v7.f9, new BigNumber(((((_2127_v32).contains(_2079_v2)) ? ((_2127_v32).get(_2079_v2)) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(526)), function (_2130_i7) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })))).length)), _2077_v0));
        _2128_v33 = _dafny.Seq.of(_dafny.Seq.Concat(_2080_v3, _dafny.Seq.UnicodeFromString("chsubpa")));
      }
      let _2131_v34;
      _2131_v34 = _dafny.Map.Empty.slice().updateUnsafe(_2079_v2,_2077_v0);
      _2131_v34 = (_2131_v34).update(false, _2077_v0);
      let _2132_v36;
      _2132_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2079_v2,p0);
      let _2133_v37;
      _2133_v37 = _module.D18.create_DC47(function () {
  let _coll81 = new _dafny.Map();
  for (const _compr_81 of (_dafny.Seq.of(_2132_v36, _2132_v36)).Elements) {
    let _2134_v35 = _compr_81;
    if (_dafny.Seq.contains(_dafny.Seq.of(_2132_v36, _2132_v36), _2134_v35)) {
      _coll81.push([_2134_v35,_dafny.Seq.UnicodeFromString("w")]);
    }
  }
  return _coll81;
}());
      let _pat_let_tv56 = _2079_v2;
      let _pat_let_tv57 = p0;
      let _pat_let_tv58 = _2098_v15;
      let _pat_let_tv59 = _2098_v15;
      let _rhs246 = function (_source22) {
        if (_source22.is_DC48) {
          return _pat_let_tv56;
        } else if (_source22.is_DC49) {
          let _2135___mcc_h0 = (_source22).cf86;
          let _2136_cf86 = _2135___mcc_h0;
          return ((_this).f2).isLessThan(_pat_let_tv57);
        } else if (_source22.is_DC50) {
          let _2137___mcc_h1 = (_source22).cf87;
          let _2138___mcc_h2 = (_source22).cf88;
          let _2139___mcc_h3 = (_source22).cf89;
          let _2140___mcc_h4 = (_source22).cf90;
          let _2141___mcc_h5 = (_source22).cf91;
          let _2142_cf91 = _2141___mcc_h5;
          let _2143_cf90 = _2140___mcc_h4;
          let _2144_cf89 = _2139___mcc_h3;
          let _2145_cf88 = _2138___mcc_h2;
          let _2146_cf87 = _2137___mcc_h1;
          return _2143_cf90;
        } else {
          let _2147___mcc_h6 = (_source22).cf85;
          let _2148_cf85 = _2147___mcc_h6;
          return (_pat_let_tv58).IsProperSubsetOf(_pat_let_tv59);
        }
      }(_2133_v37);
      let _rhs247 = !(((_this).f2).isLessThan(p0));
      r1 = _rhs246;
      r1 = _rhs247;
      r0 = ((_2079_v2) ? (false) : (_2079_v2));
      r1 = _module.__default.fm17(globalState);
      return [r0, r1];
    }
    m2(p0, globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let _2149_v0;
      _2149_v0 = true;
      _2149_v0 = _2149_v0;
      let _2150_v1;
      let _nw334 = Array((new BigNumber(20)).toNumber()).fill(false);
      _2150_v1 = _nw334;
      let _index335 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_2150_v1).length));
      (_2150_v1)[_index335] = _2149_v0;
      let _2151_v2;
      _2151_v2 = _2149_v0;
      let _index336 = _module.__default.safeIndex(new BigNumber(807), new BigNumber((_2150_v1).length));
      (_2150_v1)[_index336] = _2151_v2;
      if ((p0).isLessThanOrEqualTo(p0)) {
        let _2152_v3;
        _2152_v3 = new BigNumber(14);
        let _2153_v4;
        let _nw335 = Array((new BigNumber(10)).toNumber()).fill(false);
        _2153_v4 = _nw335;
        let _2154_v5;
        _2154_v5 = _dafny.MultiSet.fromElements(_2153_v4);
        let _index337 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_2153_v4).length));
        (_2153_v4)[_index337] = !(!(_2149_v0));
        let _index338 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_2153_v4).length));
        let _rhs248 = (_this).f2;
        let _rhs249 = _2154_v5;
        let _rhs250 = !((_2152_v3).isEqualTo(_2152_v3));
        let _rhs251 = !(!(_2149_v0));
        let _lhs138 = _2153_v4;
        let _lhs139 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_2153_v4).length));
        _2152_v3 = _rhs248;
        _2154_v5 = _rhs249;
        _lhs138[_lhs139] = _rhs250;
        _2149_v0 = _rhs251;
        let _2155_v6;
        _2155_v6 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2153_v4)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_2153_v4).length))]);
        let _2156_v9;
        _2156_v9 = _dafny.Map.Empty.slice().updateUnsafe(_2149_v0,(_2153_v4)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_2153_v4).length))]);
        let _2157_v10;
        _2157_v10 = _dafny.Set.fromElements(_2155_v6, _2155_v6, _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2156_v9).length),_2149_v0));
        let _index339 = _module.__default.safeIndex(new BigNumber(23), new BigNumber((_2153_v4).length));
        (_2153_v4)[_index339] = (_dafny.Set.fromElements(_2155_v6, function () {
          let _coll82 = new _dafny.Map();
          for (const _compr_82 of _dafny.IntegerRange(new BigNumber(567), new BigNumber(-187))) {
            let _2158_v7 = _compr_82;
            if (((new BigNumber(567)).isLessThanOrEqualTo(_2158_v7)) && ((_2158_v7).isLessThan(new BigNumber(-187)))) {
              _coll82.push([(_2158_v7).minus((_this).f2),_2149_v0]);
            }
          }
          return _coll82;
        }(), function () {
          let _coll83 = new _dafny.Map();
          for (const _compr_83 of _dafny.IntegerRange(new BigNumber(-911), new BigNumber(126))) {
            let _2159_v8 = _compr_83;
            if (((new BigNumber(-911)).isLessThanOrEqualTo(_2159_v8)) && ((_2159_v8).isLessThan(new BigNumber(126)))) {
              _coll83.push([(_2159_v8).minus((_this).f2),!((_2153_v4)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_2153_v4).length))])]);
            }
          }
          return _coll83;
        }())).IsDisjointFrom(_2157_v10);
        let _2160_v11;
        _2160_v11 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(((_2154_v5).Union(_2154_v5)).cardinality())),p0);
        _2160_v11 = (_2160_v11).update((new BigNumber(-793)).plus(new BigNumber((_dafny.Set.fromElements((_this).f2)).length)), (_this).f2);
        let _2161_v12;
        _2161_v12 = _dafny.Seq.UnicodeFromString("i");
        let _2162_v13;
        _2162_v13 = _dafny.Map.Empty.slice().updateUnsafe((_2152_v3).isLessThanOrEqualTo(_2152_v3),_2161_v12);
        let _2163_v14;
        _2163_v14 = new _dafny.CodePoint('u'.codePointAt(0));
        _2162_v13 = (_2162_v13).update(_2149_v0, _dafny.Seq.Concat(_dafny.Seq.update(_2161_v12, _module.__default.safeIndex((_this).f2, new BigNumber((_2161_v12).length)), _2163_v14), _2161_v12));
        let _2164_v15;
        _2164_v15 = _dafny.Seq.of(_2149_v0);
        let _2165_v16;
        _2165_v16 = _dafny.MultiSet.fromElements(_2152_v3);
        let _2166_v17;
        _2166_v17 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-438)), ((_2167_v14) => function (_2168_i0) {
          return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(492)), ((_2169_v14) => function (_2170_i1) {
            return _2169_v14;
          })(_2167_v14))).length);
        })(_2163_v14))).length), _2152_v3, (_this).f2);
        let _2171_v18;
        _2171_v18 = _dafny.Set.fromElements(new BigNumber((_2166_v17).length), (_this).f2);
        let _2172_v19;
        let _nw336 = Array((new BigNumber(27)).toNumber());
        _nw336[(_dafny.ZERO).toNumber()] = p0;
        _nw336[(_dafny.ONE).toNumber()] = _2152_v3;
        _nw336[(new BigNumber(2)).toNumber()] = p0;
        _nw336[(new BigNumber(3)).toNumber()] = (_this).f2;
        _nw336[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.of(_2152_v3, p0)).length);
        _nw336[(new BigNumber(5)).toNumber()] = new BigNumber((_2164_v15).length);
        _nw336[(new BigNumber(6)).toNumber()] = _2152_v3;
        _nw336[(new BigNumber(7)).toNumber()] = new BigNumber(292);
        _nw336[(new BigNumber(8)).toNumber()] = (_this).f2;
        _nw336[(new BigNumber(9)).toNumber()] = (_dafny.ZERO).minus(p0);
        _nw336[(new BigNumber(10)).toNumber()] = p0;
        _nw336[(new BigNumber(11)).toNumber()] = (_this).f2;
        _nw336[(new BigNumber(12)).toNumber()] = p0;
        _nw336[(new BigNumber(13)).toNumber()] = p0;
        _nw336[(new BigNumber(14)).toNumber()] = _2152_v3;
        _nw336[(new BigNumber(15)).toNumber()] = (_this).f2;
        _nw336[(new BigNumber(16)).toNumber()] = (_this).f2;
        _nw336[(new BigNumber(17)).toNumber()] = p0;
        _nw336[(new BigNumber(18)).toNumber()] = new BigNumber((_2161_v12).length);
        _nw336[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_2164_v15)).cardinality());
        _nw336[(new BigNumber(20)).toNumber()] = _2152_v3;
        _nw336[(new BigNumber(21)).toNumber()] = _2152_v3;
        _nw336[(new BigNumber(22)).toNumber()] = new BigNumber((_2160_v11).length);
        _nw336[(new BigNumber(23)).toNumber()] = new BigNumber((_2165_v16).cardinality());
        _nw336[(new BigNumber(24)).toNumber()] = new BigNumber((_2166_v17).length);
        _nw336[(new BigNumber(25)).toNumber()] = _2152_v3;
        _nw336[(new BigNumber(26)).toNumber()] = _2152_v3;
        _2172_v19 = _nw336;
        let _pat_let_tv60 = _2152_v3;
        let _2173_v20;
        let _nw337 = new _module.C5();
        _nw337.__ctor(_2172_v19, (_module.D2.create_DC5(p0, (_this).f0, (_2153_v4)[_module.__default.safeIndex(new BigNumber(23), new BigNumber((_2153_v4).length))])).dtor_cf9, function (_pat_let47_0) {
          return function (_2174_dt__update__tmp_h0) {
            return function (_pat_let48_0) {
              return function (_2175_dt__update_hcf2_h0) {
                return _module.D0.create_DC1(_2175_dt__update_hcf2_h0);
              }(_pat_let48_0);
            }((_dafny.ZERO).minus((_dafny.ZERO).minus(_pat_let_tv60)));
          }(_pat_let47_0);
        }(_this.f1));
        _2173_v20 = _nw337;
        let _2176_v21;
        _2176_v21 = _dafny.MultiSet.fromElements(_2173_v20, _2173_v20);
        _2176_v21 = _2176_v21;
      } else {
        let _2177_v22;
        _2177_v22 = new BigNumber(150);
        _2177_v22 = ((!(true)) ? (_2177_v22) : ((_this).f2));
        let _2178_v23;
        _2178_v23 = _dafny.Seq.UnicodeFromString("uefbomho");
        let _2179_v24;
        _2179_v24 = _module.D6.create_DC14(_2178_v23);
        let _source23 = _2179_v24;
        if (_source23.is_DC15) {
          let _2180___mcc_h0 = (_source23).cf30;
          let _2181___mcc_h1 = (_source23).cf31;
          let _2182___mcc_h2 = (_source23).cf32;
          let _2183_cf32 = _2182___mcc_h2;
          let _2184_cf31 = _2181___mcc_h1;
          let _2185_cf30 = _2180___mcc_h0;
          let _2186_v25;
          let _init53 = ((_2187_cf32) => function (_2188_i2) {
            return _dafny.Set.fromElements(!(_2187_cf32), _2187_cf32);
          })(_2183_cf32);
          let _nw338 = Array((new BigNumber(20)).toNumber());
          for (let _i0_53 = 0; _i0_53 < new BigNumber(_nw338.length); _i0_53++) {
            _nw338[_i0_53] = _init53(new BigNumber(_i0_53));
          }
          _2186_v25 = _nw338;
          let _2189_v26;
          _2189_v26 = _dafny.Map.Empty.slice().updateUnsafe(((false) ? (_2177_v22) : (_2177_v22)),_2186_v25);
          _2189_v26 = (_2189_v26).update(_2177_v22, _2186_v25);
          _2183_cf32 = true;
          let _2190_v27;
          let _nw339 = Array((new BigNumber(6)).toNumber()).fill(false);
          _2190_v27 = _nw339;
          let _index340 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_2190_v27).length));
          (_2190_v27)[_index340] = _2149_v0;
          let _index341 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_2190_v27).length));
          (_2190_v27)[_index341] = _2183_cf32;
          let _2191_v28;
          _2191_v28 = _dafny.Seq.of(_2183_cf32, _2149_v0, _2183_cf32, _2183_cf32);
          let _index342 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_2190_v27).length));
          let _index343 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_2190_v27).length));
          let _rhs252 = (((p0).isLessThanOrEqualTo(((_this).f0)[_module.__default.safeIndex(new BigNumber((_2178_v23).length), new BigNumber(((_this).f0).length))])) ? (_2149_v0) : ((_2191_v28)[_module.__default.safeIndex((_this).f2, new BigNumber((_2191_v28).length))]));
          let _rhs253 = _2149_v0;
          let _rhs254 = _2183_cf32;
          let _rhs255 = true;
          let _rhs256 = _2149_v0;
          let _lhs140 = _2190_v27;
          let _lhs141 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_2190_v27).length));
          let _lhs142 = _2190_v27;
          let _lhs143 = _module.__default.safeIndex(new BigNumber(9), new BigNumber((_2190_v27).length));
          _lhs140[_lhs141] = _rhs252;
          _lhs142[_lhs143] = _rhs253;
          _2149_v0 = _rhs254;
          _2149_v0 = _rhs255;
          _2149_v0 = _rhs256;
          let _2192_v29;
          _2192_v29 = _dafny.MultiSet.fromElements(_2149_v0);
          (_this).m4(_2149_v0, new BigNumber(105), _2192_v29, p0, globalState);
        } else if (_source23.is_DC14) {
          let _2193___mcc_h3 = (_source23).cf29;
          let _2194_cf29 = _2193___mcc_h3;
          let _2195_v30;
          let _nw340 = Array((new BigNumber(19)).toNumber()).fill(false);
          _2195_v30 = _nw340;
          let _2196_v31;
          _2196_v31 = _module.D18.create_DC50(!(_2149_v0), _2195_v30, false, false, _2194_cf29);
          let _index344 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length));
          (_2195_v30)[_index344] = (_2196_v31).dtor_cf87;
          let _index345 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length));
          (_2195_v30)[_index345] = _2149_v0;
          let _2197_v32;
          let _nw341 = new _module.C3();
          _nw341.__ctor();
          _2197_v32 = _nw341;
          let _2198_v33;
          _2198_v33 = new _dafny.CodePoint('m'.codePointAt(0));
          let _2199_v34;
          _2199_v34 = _dafny.MultiSet.fromElements(_2198_v33, _2198_v33);
          let _2200_v35;
          _2200_v35 = _dafny.Seq.of(_2199_v34);
          let _index346 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length));
          let _index347 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length));
          let _rhs257 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_2200_v35, _module.__default.safeIndex(new BigNumber(287), new BigNumber((_2200_v35).length)), _2199_v34), _dafny.Seq.Create(_module.__default.abs(new BigNumber(127)), ((_2201_v34) => function (_2202_i3) {
            return _2201_v34;
          })(_2199_v34)));
          let _rhs258 = (_2195_v30)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length))];
          let _rhs259 = !(((new BigNumber((_2194_cf29).length)).isLessThan((_this).f2)) && (!((_2195_v30)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length))])));
          let _rhs260 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).fm6((_this).f2, (_this).f2, globalState)), new BigNumber(833))).minus((_this).f2)));
          let _lhs144 = _2195_v30;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length));
          let _lhs146 = _2195_v30;
          let _lhs147 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length));
          _lhs144[_lhs145] = _rhs257;
          _lhs146[_lhs147] = _rhs258;
          _2149_v0 = _rhs259;
          _2177_v22 = _rhs260;
          let _2203_v36;
          _2203_v36 = _dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0)));
          let _2204_v37;
          _2204_v37 = _dafny.Set.fromElements(_2177_v22, (_this).f2, new BigNumber(25));
          let _2205_v38;
          _2205_v38 = _dafny.Set.fromElements(_2149_v0);
          let _2206_v39;
          _2206_v39 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_2177_v22),(_dafny.ZERO).minus((_this).f2));
          let _2207_v40;
          _2207_v40 = _dafny.Map.Empty.slice().updateUnsafe(_2198_v33,_2198_v33);
          let _2208_v41;
          _2208_v41 = _dafny.Seq.of(_2207_v40, _2207_v40);
          let _2209_v42;
          _2209_v42 = _dafny.MultiSet.fromElements(_2149_v0);
          let _2210_v43;
          _2210_v43 = _dafny.Seq.of(false);
          let _2211_v44;
          let _nw342 = Array((new BigNumber(25)).toNumber());
          _nw342[(_dafny.ZERO).toNumber()] = new BigNumber(994);
          _nw342[(_dafny.ONE).toNumber()] = _2177_v22;
          _nw342[(new BigNumber(2)).toNumber()] = (_this).f2;
          _nw342[(new BigNumber(3)).toNumber()] = _2177_v22;
          _nw342[(new BigNumber(4)).toNumber()] = new BigNumber(314);
          _nw342[(new BigNumber(5)).toNumber()] = new BigNumber((_2205_v38).length);
          _nw342[(new BigNumber(6)).toNumber()] = (_this).f2;
          _nw342[(new BigNumber(7)).toNumber()] = _2177_v22;
          _nw342[(new BigNumber(8)).toNumber()] = (((_2206_v39).contains(p0)) ? ((_2206_v39).get(p0)) : (_module.__default.fm44(_2149_v0, globalState)));
          _nw342[(new BigNumber(9)).toNumber()] = p0;
          _nw342[(new BigNumber(10)).toNumber()] = new BigNumber(((_this).f0).length);
          _nw342[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(((_this).f0).length));
          _nw342[(new BigNumber(12)).toNumber()] = (_this).f2;
          _nw342[(new BigNumber(13)).toNumber()] = new BigNumber(((_2208_v41)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_2209_v42).cardinality())), new BigNumber((_2208_v41).length))]).length);
          _nw342[(new BigNumber(14)).toNumber()] = new BigNumber((_2194_cf29).length);
          _nw342[(new BigNumber(15)).toNumber()] = _2177_v22;
          _nw342[(new BigNumber(16)).toNumber()] = (_this).f2;
          _nw342[(new BigNumber(17)).toNumber()] = new BigNumber(((_this).f0).length);
          _nw342[(new BigNumber(18)).toNumber()] = p0;
          _nw342[(new BigNumber(19)).toNumber()] = _2177_v22;
          _nw342[(new BigNumber(20)).toNumber()] = new BigNumber(266);
          _nw342[(new BigNumber(21)).toNumber()] = _2177_v22;
          _nw342[(new BigNumber(22)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_2210_v43)).cardinality());
          _nw342[(new BigNumber(23)).toNumber()] = (_this).f2;
          _nw342[(new BigNumber(24)).toNumber()] = p0;
          _2211_v44 = _nw342;
          let _2212_v45;
          _2212_v45 = _dafny.Seq.of(_2211_v44, _2211_v44, _2211_v44);
          let _2213_v46;
          _2213_v46 = _dafny.Map.Empty.slice().updateUnsafe((_2195_v30)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length))],p0);
          let _2214_v47;
          _2214_v47 = _dafny.Map.Empty.slice().updateUnsafe((_2212_v45)[_module.__default.safeIndex(new BigNumber((_2213_v46).length), new BigNumber((_2212_v45).length))],(_this).f2);
          let _2215_v48;
          _2215_v48 = _dafny.MultiSet.fromElements(new BigNumber((_module.__default.fm25(_module.__default.fm64(globalState), _dafny.Seq.update(_module.__default.fm16(_2178_v23, new BigNumber((_dafny.MultiSet.fromElements((_this).f2, new BigNumber((_2210_v43).length), _2177_v22, (_this).f2, new BigNumber((_2204_v37).length))).cardinality()), globalState), _module.__default.safeIndex((_this).f2, new BigNumber((_module.__default.fm16(_2178_v23, new BigNumber((_dafny.MultiSet.fromElements((_this).f2, new BigNumber((_2210_v43).length), _2177_v22, (_this).f2, new BigNumber((_2204_v37).length))).cardinality()), globalState)).length)), _2198_v33), globalState)).length), _2177_v22);
          let _2216_v49;
          _2216_v49 = _dafny.Map.Empty.slice().updateUnsafe(_2198_v33,p0);
          let _2217_v50;
          _2217_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2195_v30)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length))]);
          let _index348 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length));
          let _rhs261 = _dafny.Set.fromElements(_2198_v33, _2198_v33, _2198_v33, _2198_v33, _2198_v33);
          let _rhs262 = _2149_v0;
          let _rhs263 = (_dafny.Set.fromElements(new BigNumber((_2215_v48).cardinality()), new BigNumber(12), new BigNumber(535), new BigNumber((_2216_v49).length))).Difference((function () {
            let _coll84 = new _dafny.Set();
            for (const _compr_84 of (_2217_v50).Keys.Elements) {
              let _2218_v51 = _compr_84;
              if ((_2217_v50).contains(_2218_v51)) {
                _coll84.add(_module.__default.safeDivisionInt(_2218_v51, new BigNumber((_dafny.Seq.of(new BigNumber(((_module.D2.create_DC5(new BigNumber((_dafny.Seq.of(true, false, true, true)).length), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(671),new BigNumber(-9))).length)), true)).dtor_cf9).length))).length)));
              }
            }
            return _coll84;
          }()).Intersect(_2204_v37));
          let _rhs264 = (_2214_v47).Merge(_2214_v47);
          let _rhs265 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, (((_2195_v30)[_module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length))]) ? (_2177_v22) : ((_this).f2))));
          let _lhs148 = _2195_v30;
          let _lhs149 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_2195_v30).length));
          _2203_v36 = _rhs261;
          _lhs148[_lhs149] = _rhs262;
          _2204_v37 = _rhs263;
          _2214_v47 = _rhs264;
          _2177_v22 = _rhs265;
        } else {
          let _2219___mcc_h4 = (_source23).cf33;
          let _2220_cf33 = _2219___mcc_h4;
          let _2221_v52;
          let _nw343 = new _module.C4();
          _nw343.__ctor();
          _2221_v52 = _nw343;
          _2149_v0 = _2149_v0;
          let _2222_v53;
          let _nw344 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Seq.of());
          _2222_v53 = _nw344;
          let _2223_v54;
          _2223_v54 = _dafny.Seq.of(true);
          let _index349 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_2222_v53).length));
          (_2222_v53)[_index349] = _2223_v54;
          let _2224_v55;
          _2224_v55 = new _dafny.CodePoint('g'.codePointAt(0));
          let _2225_v56;
          _2225_v56 = _dafny.Map.Empty.slice().updateUnsafe(_2224_v55,_2223_v54);
          let _2226_v57;
          _2226_v57 = _module.D15.create_DC38(_2224_v55);
          let _index350 = _module.__default.safeIndex(new BigNumber(956), new BigNumber((_2222_v53).length));
          (_2222_v53)[_index350] = (((_2225_v56).contains((_2226_v57).dtor_cf71)) ? ((_2225_v56).get((_2226_v57).dtor_cf71)) : (_dafny.Seq.update(_dafny.Seq.Concat(_2223_v54, _2223_v54), _module.__default.safeIndex(_2177_v22, new BigNumber((_dafny.Seq.Concat(_2223_v54, _2223_v54)).length)), _2149_v0)));
          let _2227_v58;
          let _nw345 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
          _2227_v58 = _nw345;
          let _index351 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_2227_v58).length));
          (_2227_v58)[_index351] = (_this).f2;
          let _2228_v59;
          _2228_v59 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('i'.codePointAt(0)));
          let _index352 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_2227_v58).length));
          (_2227_v58)[_index352] = new BigNumber((_2228_v59).cardinality());
          let _2229_v60;
          _2229_v60 = _dafny.Seq.of(_2227_v58, _2227_v58, _2227_v58, _2227_v58);
          let _2230_v61;
          _2230_v61 = _dafny.MultiSet.fromElements(false);
          let _index353 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_2227_v58).length));
          let _index354 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_2227_v58).length));
          let _rhs266 = new BigNumber((_2229_v60).length);
          let _rhs267 = new BigNumber((_2230_v61).cardinality());
          let _lhs150 = _2227_v58;
          let _lhs151 = _module.__default.safeIndex(new BigNumber(442), new BigNumber((_2227_v58).length));
          let _lhs152 = _2227_v58;
          let _lhs153 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_2227_v58).length));
          _lhs150[_lhs151] = _rhs266;
          _lhs152[_lhs153] = _rhs267;
        }
        let _pat_let_tv61 = _2177_v22;
        let _2231_v62;
        let _2232_v63;
        let _2233_v64;
        let _out52;
        let _out53;
        let _out54;
        let _outcollector17 = (_this).m3(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2178_v23).length),_2149_v0)).length), function (_pat_let49_0) {
          return function (_2234_dt__update__tmp_h1) {
            return function (_pat_let50_0) {
              return function (_2235_dt__update_hcf1_h0) {
                return _module.D0.create_DC0((_2234_dt__update__tmp_h1).dtor_cf0, _2235_dt__update_hcf1_h0);
              }(_pat_let50_0);
            }(_pat_let_tv61);
          }(_pat_let49_0);
        }((_this).f3), true, globalState);
        _out52 = _outcollector17[0];
        _out53 = _outcollector17[1];
        _out54 = _outcollector17[2];
        _2231_v62 = _out52;
        _2232_v63 = _out53;
        _2233_v64 = _out54;
        let _2236_v65;
        _2236_v65 = _module.D7.create_DC19((_this).f2);
        let _source24 = _2236_v65;
        if (_source24.is_DC18) {
          let _2237_v66;
          _2237_v66 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('s'.codePointAt(0)),_this.f1);
          _2149_v0 = (_module.__default.fm44(!((_this).fm5((_this).f2, !(_2149_v0), _2237_v66, _2149_v0, globalState)), globalState)).isEqualTo(new BigNumber(-859));
          let _2238_v67;
          _2238_v67 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_2231_v62,_2149_v0));
          let _2239_v68;
          _2239_v68 = _dafny.Seq.of(new BigNumber(((_2238_v67)[_module.__default.safeIndex(_2231_v62, new BigNumber((_2238_v67).length))]).length), new BigNumber(-681));
          let _2240_v69;
          _2240_v69 = _dafny.Seq.of(_2149_v0);
          let _2241_v70;
          let _nw346 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
          _2241_v70 = _nw346;
          let _2242_v71;
          _2242_v71 = _module.D13.create_DC32(_2178_v23, (_2240_v69)[_module.__default.safeIndex(new BigNumber(10), new BigNumber((_2240_v69).length))], _2177_v22, _2241_v70);
          let _rhs268 = (_2242_v71).dtor_cf58;
          let _rhs269 = _dafny.Seq.Concat(_dafny.Seq.Concat(_2239_v68, _2239_v68), _2239_v68);
          _2149_v0 = _rhs268;
          _2239_v68 = _rhs269;
          let _2243_v72;
          let _nw347 = Array((new BigNumber(13)).toNumber()).fill(false);
          _2243_v72 = _nw347;
          _2243_v72 = _2243_v72;
          _2149_v0 = _2149_v0;
        } else if (_source24.is_DC19) {
          let _2244___mcc_h5 = (_source24).cf35;
          let _2245_cf35 = _2244___mcc_h5;
          _2149_v0 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_2178_v23, _2178_v23), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), function (_2246_i4) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }), _2178_v23));
          _2233_v64 = _module.__default.fm44(true, globalState);
          _2149_v0 = _2149_v0;
          _2233_v64 = (_this).f2;
        } else {
          let _2247___mcc_h6 = (_source24).cf34;
          let _2248_cf34 = _2247___mcc_h6;
          let _2249_v73;
          _2249_v73 = _dafny.MultiSet.fromElements(_2231_v62, _2177_v22, (_this).fm6(_module.__default.fm32(new BigNumber(990), globalState), p0, globalState), p0, _2233_v64);
          _2149_v0 = (_2249_v73).IsSubsetOf((_2249_v73).Difference(_dafny.MultiSet.fromElements((_this).f2)));
          let _2250_v74;
          let _nw348 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
          _2250_v74 = _nw348;
          let _index355 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_2250_v74).length));
          (_2250_v74)[_index355] = new BigNumber(341);
          let _index356 = _module.__default.safeIndex(new BigNumber(230), new BigNumber((_2250_v74).length));
          (_2250_v74)[_index356] = ((_dafny.ZERO).minus(new BigNumber(((_this).f0).length))).minus(new BigNumber(-205));
          let _2251_v75;
          _2251_v75 = new _dafny.CodePoint('x'.codePointAt(0));
          let _2252_v76;
          _2252_v76 = _dafny.MultiSet.fromElements(_2251_v75, _2251_v75);
          let _2253_v77;
          let _2254_v78;
          let _2255_v79;
          let _out55;
          let _out56;
          let _out57;
          let _outcollector18 = (_this).m3(new BigNumber(147), (_this).f3, (_2252_v76).IsSubsetOf(_2252_v76), globalState);
          _out55 = _outcollector18[0];
          _out56 = _outcollector18[1];
          _out57 = _outcollector18[2];
          _2253_v77 = _out55;
          _2254_v78 = _out56;
          _2255_v79 = _out57;
          let _2256_v80;
          _2256_v80 = _dafny.Map.Empty.slice().updateUnsafe(_2149_v0,(_2248_cf34)[_module.__default.safeIndex(new BigNumber((_2178_v23).length), new BigNumber((_2248_cf34).length))]);
          _2149_v0 = (((_2256_v80).contains(_2149_v0)) ? ((_2256_v80).get(_2149_v0)) : (_2149_v0));
        }
        let _2257_v81;
        _2257_v81 = _dafny.Seq.of(_2149_v0, _2149_v0, _module.__default.fm17(globalState));
        let _2258_v82;
        _2258_v82 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm7(_2233_v64, globalState),_2149_v0);
        let _2259_v83;
        _2259_v83 = _dafny.Map.Empty.slice().updateUnsafe(_2177_v22,_2149_v0);
        let _rhs270 = _dafny.Seq.of((((_2258_v82).contains(_2178_v23)) ? ((_2258_v82).get(_2178_v23)) : ((((_2259_v83).contains(_2233_v64)) ? ((_2259_v83).get(_2233_v64)) : (_2149_v0)))));
        let _rhs271 = _2233_v64;
        _2257_v81 = _rhs270;
        _2231_v62 = _rhs271;
      }
      let _2260_v84;
      let _init54 = ((_2261_v0) => function (_2262_i5) {
        return _2261_v0;
      })(_2149_v0);
      let _nw349 = Array((new BigNumber(5)).toNumber());
      for (let _i0_54 = 0; _i0_54 < new BigNumber(_nw349.length); _i0_54++) {
        _nw349[_i0_54] = _init54(new BigNumber(_i0_54));
      }
      _2260_v84 = _nw349;
      _2260_v84 = ((_2149_v0) ? (_2260_v84) : (_2260_v84));
      let _2263_v85;
      _2263_v85 = new BigNumber(844);
      _2263_v85 = (_dafny.ZERO).minus(p0);
      let _2264_v86;
      _2264_v86 = _dafny.Seq.of(_2149_v0, !(_2149_v0) || (_2149_v0));
      let _2265_i6;
      _2265_i6 = _dafny.ZERO;
      L9: {
        while ((_2264_v86)[_module.__default.safeIndex(((_2149_v0) ? (new BigNumber((_2264_v86).length)) : (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(73)), function (_2293_i7) {
          return new _dafny.CodePoint('q'.codePointAt(0));
        })).length))), new BigNumber((_2264_v86).length))]) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2265_i6)) {
              break L9;
            }
            _2265_i6 = (_2265_i6).plus(_dafny.ONE);
            _2263_v85 = new BigNumber(724);
            if (!((_module.D6.create_DC15(p0, _module.__default.fm20(_2149_v0, globalState), !(_2149_v0))).dtor_cf32)) {
              let _2266_v87;
              _2266_v87 = _dafny.Seq.of(((_2149_v0) ? (_2260_v84) : (_2260_v84)), _2260_v84, _2260_v84);
              _2266_v87 = _2266_v87;
              let _2267_v88;
              let _init55 = ((_2268_v86) => function (_2269_i8) {
                return _2268_v86;
              })(_2264_v86);
              let _nw350 = Array((new BigNumber(18)).toNumber());
              for (let _i0_55 = 0; _i0_55 < new BigNumber(_nw350.length); _i0_55++) {
                _nw350[_i0_55] = _init55(new BigNumber(_i0_55));
              }
              _2267_v88 = _nw350;
              let _index357 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_2267_v88).length));
              (_2267_v88)[_index357] = _2264_v86;
              let _2270_v89;
              _2270_v89 = _module.D16.create_DC44((_this).f2);
              let _2271_v90;
              _2271_v90 = _dafny.Map.Empty.slice().updateUnsafe((_2270_v89).dtor_cf82,_2264_v86);
              let _2272_v91;
              _2272_v91 = _dafny.MultiSet.fromElements(_2264_v86);
              let _index358 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_2267_v88).length));
              (_2267_v88)[_index358] = (((_2271_v90).contains(new BigNumber(((_2272_v91).update(_dafny.Seq.of(_2149_v0, _2149_v0, false, !(_2149_v0), _module.__default.fm17(globalState)), _module.__default.abs(_2263_v85))).cardinality()))) ? ((_2271_v90).get(new BigNumber(((_2272_v91).update(_dafny.Seq.of(_2149_v0, _2149_v0, false, !(_2149_v0), _module.__default.fm17(globalState)), _module.__default.abs(_2263_v85))).cardinality()))) : (_2264_v86));
              _2263_v85 = p0;
              let _2273_v92;
              let _init56 = ((_2274_p0) => function (_2275_i9) {
                return _module.__default.safeDivisionInt(_2275_i9, _2274_p0);
              })(p0);
              let _nw351 = Array((new BigNumber(2)).toNumber());
              for (let _i0_56 = 0; _i0_56 < new BigNumber(_nw351.length); _i0_56++) {
                _nw351[_i0_56] = _init56(new BigNumber(_i0_56));
              }
              _2273_v92 = _nw351;
              let _2276_v93;
              let _nw352 = new _module.C5();
              _nw352.__ctor(_2273_v92, (_this).f0, _this.f1);
              _2276_v93 = _nw352;
              let _2277_v94;
              _2277_v94 = _dafny.Seq.of(_2276_v93, _2276_v93);
              let _2278_v95;
              let _nw353 = Array((new BigNumber(13)).toNumber());
              _nw353[(_dafny.ZERO).toNumber()] = _2276_v93;
              _nw353[(_dafny.ONE).toNumber()] = ((_2149_v0) ? (_2276_v93) : (_2276_v93));
              _nw353[(new BigNumber(2)).toNumber()] = _2276_v93;
              _nw353[(new BigNumber(3)).toNumber()] = _2276_v93;
              _nw353[(new BigNumber(4)).toNumber()] = _2276_v93;
              _nw353[(new BigNumber(5)).toNumber()] = _2276_v93;
              _nw353[(new BigNumber(6)).toNumber()] = (_2277_v94)[_module.__default.safeIndex(_2263_v85, new BigNumber((_2277_v94).length))];
              _nw353[(new BigNumber(7)).toNumber()] = _2276_v93;
              _nw353[(new BigNumber(8)).toNumber()] = _2276_v93;
              _nw353[(new BigNumber(9)).toNumber()] = _2276_v93;
              _nw353[(new BigNumber(10)).toNumber()] = _2276_v93;
              _nw353[(new BigNumber(11)).toNumber()] = _2276_v93;
              _nw353[(new BigNumber(12)).toNumber()] = _2276_v93;
              _2278_v95 = _nw353;
              let _index359 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2278_v95).length));
              (_2278_v95)[_index359] = _2276_v93;
              let _index360 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_2278_v95).length));
              (_2278_v95)[_index360] = _2276_v93;
              let _2279_v96;
              _2279_v96 = new _dafny.CodePoint('r'.codePointAt(0));
              let _2280_v97;
              _2280_v97 = _dafny.Map.Empty.slice().updateUnsafe(_2279_v96,_2263_v85);
              let _2281_v98;
              _2281_v98 = _dafny.Seq.of(_2264_v86);
              let _2282_v99;
              _2282_v99 = _dafny.Set.fromElements(_module.__default.fm19(_2149_v0, _2149_v0, _2280_v97, !(_2149_v0), globalState), (_2281_v98)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f2), new BigNumber((_2281_v98).length))], _module.__default.fm19(_2149_v0, _2149_v0, _2280_v97, _2149_v0, globalState));
              _2263_v85 = new BigNumber((_2282_v99).length);
            } else {
              _2263_v85 = p0;
              let _2283_v100;
              let _nw354 = Array((new BigNumber(17)).toNumber()).fill(_dafny.ZERO);
              _2283_v100 = _nw354;
              let _2284_v101;
              let _nw355 = new _module.C5();
              _nw355.__ctor(_2283_v100, _dafny.Seq.Concat((_this).f0, (_this).f0), _module.D0.create_DC1(_2263_v85));
              _2284_v101 = _nw355;
              let _2285_v102;
              _2285_v102 = _dafny.Seq.UnicodeFromString("v");
              _2285_v102 = _dafny.Seq.Concat(_module.__default.fm16(_2285_v102, (_dafny.ZERO).minus(new BigNumber((_2285_v102).length)), globalState), _2285_v102);
              let _2286_v103;
              _2286_v103 = _dafny.Map.Empty.slice().updateUnsafe(_2149_v0,_2149_v0);
              let _2287_v104;
              _2287_v104 = new _dafny.CodePoint('y'.codePointAt(0));
              let _2288_v105;
              _2288_v105 = _dafny.Map.Empty.slice().updateUnsafe((_2284_v101).f7,_dafny.Map.Empty.slice().updateUnsafe(_2149_v0,_module.__default.fm65(new BigNumber((_2286_v103).length), _2287_v104, (_this).f2, globalState)));
              let _2289_v106;
              _2289_v106 = _dafny.Map.Empty.slice().updateUnsafe(_2149_v0,p0);
              let _2290_v107;
              _2290_v107 = _dafny.Map.Empty.slice().updateUnsafe(true,_2289_v106);
              _2288_v105 = (_2288_v105).update((_2284_v101).f7, _2290_v107);
              let _2291_v108;
              _2291_v108 = _dafny.Set.fromElements((_this).f2);
              let _2292_v109;
              _2292_v109 = _dafny.Map.Empty.slice().updateUnsafe(_2260_v84,_2263_v85);
              let _rhs272 = (_2291_v108).Union(_2291_v108);
              let _rhs273 = _module.__default.fm17(globalState);
              let _rhs274 = new BigNumber(363);
              let _rhs275 = (_2292_v109).Merge(_2292_v109);
              let _rhs276 = (_dafny.ZERO).minus((_dafny.ZERO).minus(((_this).f2).minus((_2284_v101).fm6((_this).f2, (_dafny.ZERO).minus(new BigNumber(-447)), globalState))));
              _2291_v108 = _rhs272;
              _2149_v0 = _rhs273;
              _2263_v85 = _rhs274;
              _2292_v109 = _rhs275;
              _2263_v85 = _rhs276;
            }
            _2263_v85 = p0;
            _2149_v0 = !(!(_2149_v0));
          }
        }
      }
      let _2294_v110;
      let _nw356 = Array((new BigNumber(7)).toNumber());
      _2294_v110 = _nw356;
      let _2295_v111;
      _2295_v111 = _dafny.Map.Empty.slice().updateUnsafe(_2149_v0,_dafny.MultiSet.fromElements((_this).f2));
      let _2296_v112;
      _2296_v112 = _dafny.Map.Empty.slice().updateUnsafe(_2294_v110,_2295_v111);
      let _2297_v113;
      _2297_v113 = _dafny.MultiSet.fromElements(_2263_v85);
      r0 = ((((_2296_v112).contains(_2294_v110)) ? ((_2296_v112).get(_2294_v110)) : (_dafny.Map.Empty.slice().updateUnsafe(_2149_v0,_2297_v113)))).Merge(_2295_v111);
      return r0;
    }
    m3(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.Map.Empty;
      let r2 = _dafny.ZERO;
      let _2298_i0;
      _2298_i0 = _dafny.ZERO;
      L10: {
        while (!(p2)) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2298_i0)) {
              break L10;
            }
            _2298_i0 = (_2298_i0).plus(_dafny.ONE);
            let _2299_v0;
            _2299_v0 = _dafny.MultiSet.fromElements(p2);
            let _2300_v1;
            _2300_v1 = _dafny.MultiSet.fromElements((_this).f2, new BigNumber(949), (_this).f2, new BigNumber(303), (_this).f2);
            let _2301_v2;
            _2301_v2 = _dafny.Map.Empty.slice().updateUnsafe(!(p2),(_this).f2);
            let _2302_v3;
            _2302_v3 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_dafny.ZERO).minus(new BigNumber((_2301_v2).length)));
            let _2303_v4;
            _2303_v4 = new _dafny.CodePoint('m'.codePointAt(0));
            let _2304_v5;
            _2304_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-939),p0);
            let _2305_v6;
            _2305_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(981),p2);
            let _2306_v7;
            _2306_v7 = _dafny.Map.Empty.slice().updateUnsafe(p2,new _dafny.CodePoint('m'.codePointAt(0)));
            let _2307_v8;
            _2307_v8 = _dafny.Set.fromElements(_2306_v7);
            let _2308_v9;
            let _nw357 = Array((new BigNumber(2)).toNumber());
            _nw357[(_dafny.ZERO).toNumber()] = false;
            _nw357[(_dafny.ONE).toNumber()] = (((_2305_v6).contains(new BigNumber(532))) ? ((_2305_v6).get(new BigNumber(532))) : (p2));
            _2308_v9 = _nw357;
            let _2309_v10;
            _2309_v10 = _dafny.Map.Empty.slice().updateUnsafe(_2308_v9,p0);
            let _2310_v11;
            let _nw358 = Array((new BigNumber(22)).toNumber());
            _nw358[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(810)), function (_2311_i1) {
              return new _dafny.CodePoint('p'.codePointAt(0));
            })).length);
            _nw358[(_dafny.ONE).toNumber()] = new BigNumber((_2299_v0).cardinality());
            _nw358[(new BigNumber(2)).toNumber()] = (_this).f2;
            _nw358[(new BigNumber(3)).toNumber()] = ((_dafny.ZERO).minus((_this).f2)).multipliedBy(new BigNumber((_2300_v1).cardinality()));
            _nw358[(new BigNumber(4)).toNumber()] = ((p2) ? (new BigNumber((_module.__default.fm31(new BigNumber(-301), p2, p0, globalState)).length)) : (p0));
            _nw358[(new BigNumber(5)).toNumber()] = p0;
            _nw358[(new BigNumber(6)).toNumber()] = new BigNumber((_2301_v2).length);
            _nw358[(new BigNumber(7)).toNumber()] = ((_this).f2).plus((_this).f2);
            _nw358[(new BigNumber(8)).toNumber()] = new BigNumber(23);
            _nw358[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.of(p2, (_this).fm5((_this).f2, p2, _dafny.Map.Empty.slice().updateUnsafe(_2303_v4,_this.f1), p2, globalState))).length);
            _nw358[(new BigNumber(10)).toNumber()] = (_this).f2;
            _nw358[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f2, (_dafny.ZERO).minus((((_2304_v5).contains((_dafny.ZERO).minus((_this).f2))) ? ((_2304_v5).get((_dafny.ZERO).minus((_this).f2))) : ((_this).f2))), (_this).f2, new BigNumber((_2305_v6).length))).length));
            _nw358[(new BigNumber(12)).toNumber()] = (_this).f2;
            _nw358[(new BigNumber(13)).toNumber()] = (p0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2307_v8).length),(_this).f2)).length));
            _nw358[(new BigNumber(14)).toNumber()] = new BigNumber(502);
            _nw358[(new BigNumber(15)).toNumber()] = (_this).fm6((_dafny.ZERO).minus((_this).f2), (_this).f2, globalState);
            _nw358[(new BigNumber(16)).toNumber()] = (_this).f2;
            _nw358[(new BigNumber(17)).toNumber()] = (new BigNumber((_2309_v10).length)).minus(_module.__default.fm44(p2, globalState));
            _nw358[(new BigNumber(18)).toNumber()] = p0;
            _nw358[(new BigNumber(19)).toNumber()] = (((_2301_v2).contains(!(p2))) ? ((_2301_v2).get(!(p2))) : (new BigNumber(800)));
            _nw358[(new BigNumber(20)).toNumber()] = new BigNumber(4);
            _nw358[(new BigNumber(21)).toNumber()] = p0;
            _2310_v11 = _nw358;
            let _index361 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_2310_v11).length));
            (_2310_v11)[_index361] = p0;
            let _index362 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_2310_v11).length));
            (_2310_v11)[_index362] = (p0).plus(p0);
            let _2312_v12;
            _2312_v12 = _dafny.Seq.of(p2, p2);
            let _2313_v13;
            _2313_v13 = _dafny.Map.Empty.slice().updateUnsafe((_module.D7.create_DC17(_2312_v12)).dtor_cf34,p2);
            r0 = ((_this).f0)[_module.__default.safeIndex(new BigNumber((_2313_v13).length), new BigNumber(((_this).f0).length))];
            let _2314_v14;
            _2314_v14 = true;
            let _2315_v15;
            _2315_v15 = _dafny.Seq.UnicodeFromString("hivvd");
            _2314_v14 = (_this).fm5(new BigNumber((_dafny.Seq.Concat(_2315_v15, _2315_v15)).length), ((true) ? (_2314_v14) : (p2)), _dafny.Map.Empty.slice().updateUnsafe(_2303_v4,_this.f1), ((_2310_v11)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_2310_v11).length))]).isLessThanOrEqualTo(new BigNumber((_2315_v15).length)), globalState);
            let _2316_v16;
            let _nw359 = new _module.C6();
            _nw359.__ctor(_module.__default.safeModuloInt(new BigNumber(((_this).f0).length), p0));
            _2316_v16 = _nw359;
          }
        }
      }
      let _2317_i2;
      _2317_i2 = _dafny.ZERO;
      L11: {
        while ((new BigNumber(642)).isLessThanOrEqualTo(_module.__default.safeModuloInt((_this).f2, p0))) {
          C11: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2317_i2)) {
              break L11;
            }
            _2317_i2 = (_2317_i2).plus(_dafny.ONE);
            if (p2) {
              let _2318_v17;
              let _nw360 = Array((new BigNumber(23)).toNumber()).fill([]);
              _2318_v17 = _nw360;
              _2318_v17 = _2318_v17;
              let _2319_v18;
              _2319_v18 = _dafny.Seq.UnicodeFromString("xkbca");
              let _2320_v19;
              _2320_v19 = new _dafny.CodePoint('h'.codePointAt(0));
              let _2321_v20;
              _2321_v20 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Concat(_2319_v18, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("snuobpymy"), _module.__default.safeIndex((_this).f2, new BigNumber((_dafny.Seq.UnicodeFromString("snuobpymy")).length)), _2320_v19)),(_this).f2);
              _2321_v20 = (_2321_v20).update(_2319_v18, _module.__default.safeDivisionInt(new BigNumber(459), (_dafny.ZERO).minus(p0)));
              let _2322_v21;
              _2322_v21 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,(_this).f2);
              let _2323_v22;
              _2323_v22 = _module.D13.create_DC31(true, p0, (_this).f2);
              let _2324_v23;
              _2324_v23 = _dafny.Map.Empty.slice().updateUnsafe(false,true);
              let _2325_v24;
              let _nw361 = Array((new BigNumber(15)).toNumber());
              _nw361[(_dafny.ZERO).toNumber()] = (_this).f2;
              _nw361[(_dafny.ONE).toNumber()] = new BigNumber((((p2) ? (_dafny.Seq.update(_2319_v18, _module.__default.safeIndex(p0, new BigNumber((_2319_v18).length)), _2320_v19)) : (_dafny.Seq.UnicodeFromString("b")))).length);
              _nw361[(new BigNumber(2)).toNumber()] = new BigNumber((_2319_v18).length);
              _nw361[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Set.fromElements(p0, p0)).length), p0);
              _nw361[(new BigNumber(4)).toNumber()] = _module.__default.safeModuloInt((_this).f2, (_module.__default.fm66(new BigNumber((_2322_v21).length), p2, globalState)).dtor_cf78);
              _nw361[(new BigNumber(5)).toNumber()] = ((p2) ? (p0) : (p0));
              _nw361[(new BigNumber(6)).toNumber()] = (_2323_v22).dtor_cf55;
              _nw361[(new BigNumber(7)).toNumber()] = (_this).f2;
              _nw361[(new BigNumber(8)).toNumber()] = p0;
              _nw361[(new BigNumber(9)).toNumber()] = new BigNumber((_2324_v23).length);
              _nw361[(new BigNumber(10)).toNumber()] = p0;
              _nw361[(new BigNumber(11)).toNumber()] = p0;
              _nw361[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("gk")).length);
              _nw361[(new BigNumber(13)).toNumber()] = ((_this).f2).minus(_module.__default.fm32(p0, globalState));
              _nw361[(new BigNumber(14)).toNumber()] = p0;
              _2325_v24 = _nw361;
              let _2326_v25;
              _2326_v25 = _dafny.MultiSet.fromElements((_this).f2);
              let _index363 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_2325_v24).length));
              (_2325_v24)[_index363] = new BigNumber((_2326_v25).cardinality());
              let _2327_v26;
              let _init57 = ((_2328_p2) => function (_2329_i3) {
                return _2328_p2;
              })(p2);
              let _nw362 = Array((new BigNumber(25)).toNumber());
              for (let _i0_57 = 0; _i0_57 < new BigNumber(_nw362.length); _i0_57++) {
                _nw362[_i0_57] = _init57(new BigNumber(_i0_57));
              }
              _2327_v26 = _nw362;
              let _2330_v27;
              _2330_v27 = _dafny.Set.fromElements(((p2) ? (_2327_v26) : (_2327_v26)), _2327_v26, _2327_v26);
              let _2331_v28;
              _2331_v28 = false;
              let _index364 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_2325_v24).length));
              let _rhs277 = _module.__default.safeModuloInt((_this).f2, p0);
              let _rhs278 = p0;
              let _rhs279 = _2330_v27;
              let _rhs280 = _2331_v28;
              let _lhs154 = _2325_v24;
              let _lhs155 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_2325_v24).length));
              r2 = _rhs277;
              _lhs154[_lhs155] = _rhs278;
              _2330_v27 = _rhs279;
              _2331_v28 = _rhs280;
              let _index365 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2327_v26).length));
              (_2327_v26)[_index365] = !(_2331_v28) || (p2);
              let _index366 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2327_v26).length));
              let _index367 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_2325_v24).length));
              let _rhs281 = !(_2331_v28);
              let _rhs282 = _2331_v28;
              let _rhs283 = (_dafny.ZERO).minus((_2325_v24)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_2325_v24).length))]);
              let _lhs156 = _2327_v26;
              let _lhs157 = _module.__default.safeIndex(new BigNumber(228), new BigNumber((_2327_v26).length));
              let _lhs158 = _2325_v24;
              let _lhs159 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_2325_v24).length));
              _2331_v28 = _rhs281;
              _lhs156[_lhs157] = _rhs282;
              _lhs158[_lhs159] = _rhs283;
              let _index368 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_2325_v24).length));
              (_2325_v24)[_index368] = ((_2325_v24)[_module.__default.safeIndex(new BigNumber(102), new BigNumber((_2325_v24).length))]).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(141)), ((_2332_v19) => function (_2333_i4) {
                return _2332_v19;
              })(_2320_v19))).length));
            } else {
              let _2334_v29;
              _2334_v29 = _dafny.MultiSet.fromElements(p2, p2, p2, p2);
              let _2335_v30;
              _2335_v30 = _dafny.MultiSet.fromElements(p0);
              let _2336_v31;
              _2336_v31 = _dafny.Set.fromElements((_this).f2, (((_2334_v29).contains(p2)) ? ((_2334_v29).get(p2)) : ((((_2335_v30).contains((_this).f2)) ? ((_2335_v30).get((_this).f2)) : ((_this).f2)))));
              let _2337_v32;
              _2337_v32 = _module.D9.create_DC22(p0);
              let _2338_v33;
              _2338_v33 = _dafny.Map.Empty.slice().updateUnsafe(_2336_v31,_2337_v32);
              _2338_v33 = _2338_v33;
              let _2339_v34;
              let _nw363 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
              _2339_v34 = _nw363;
              let _index369 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_2339_v34).length));
              (_2339_v34)[_index369] = p0;
              let _index370 = _module.__default.safeIndex(new BigNumber(39), new BigNumber((_2339_v34).length));
              (_2339_v34)[_index370] = (_dafny.ZERO).minus((p0).minus((_this).f2));
              let _2340_v35;
              _2340_v35 = _dafny.Seq.UnicodeFromString("ovve");
              _2340_v35 = _2340_v35;
              let _2341_v36;
              _2341_v36 = false;
              _2341_v36 = _2341_v36;
              _2341_v36 = !(_module.__default.safeDivisionInt(p0, (_2339_v34)[_module.__default.safeIndex(new BigNumber(39), new BigNumber((_2339_v34).length))])).isEqualTo(new BigNumber((((_2341_v36) ? ((_this).f0) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(388)), function (_2342_i5) {
                return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-225),new BigNumber(((_this).f0).length))).length);
              })))).length));
            }
            let _2343_v37;
            _2343_v37 = new _dafny.CodePoint('l'.codePointAt(0));
            let _2344_v38;
            _2344_v38 = _module.D11.create_DC25(_2343_v37, p0, !(p0).isEqualTo(new BigNumber(800)));
            let _2345_v39;
            _2345_v39 = _dafny.Set.fromElements((_this).f2);
            let _2346_v40;
            _2346_v40 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,p2);
            let _2347_v41;
            _2347_v41 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2345_v39).length),(((_2346_v40).contains(p0)) ? ((_2346_v40).get(p0)) : (p2)));
            let _2348_v42;
            _2348_v42 = _module.D16.create_DC44(p0);
            let _2349_v43;
            _2349_v43 = _dafny.Set.fromElements(new BigNumber(((_2346_v40).Merge(_dafny.Map.Empty.slice().updateUnsafe(p0,!(p2)))).length), ((_this).f2).minus((_this).f2), (_2348_v42).dtor_cf82, p0);
            let _2350_v44;
            _2350_v44 = _dafny.Seq.UnicodeFromString("ntyajnp");
            let _2351_v45;
            _2351_v45 = true;
            let _2352_v46;
            _2352_v46 = _dafny.Map.Empty.slice().updateUnsafe(!(_2351_v45),_2351_v45);
            let _2353_v47;
            _2353_v47 = _dafny.Seq.of(p2, p2, p2, p2, p2);
            let _2354_v48;
            _2354_v48 = _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(p2,p2), _dafny.Map.Empty.slice().updateUnsafe(p2,_2351_v45), _dafny.Map.Empty.slice().updateUnsafe((_2353_v47)[_module.__default.safeIndex((_this).f2, new BigNumber((_2353_v47).length))],p2));
            let _rhs284 = _module.__default.fm67((_dafny.Set.fromElements(_2352_v46)).Intersect(_2354_v48), globalState);
            let _rhs285 = (_dafny.Set.fromElements(p0)).Union(_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements((_this).f2)).cardinality())));
            let _rhs286 = (_this).f2;
            let _rhs287 = _2350_v44;
            let _rhs288 = p2;
            _2344_v38 = _rhs284;
            _2345_v39 = _rhs285;
            r0 = _rhs286;
            _2350_v44 = _rhs287;
            _2351_v45 = _rhs288;
            let _2355_v49;
            _2355_v49 = _dafny.Set.fromElements(_2350_v44, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-745)), ((_2356_v37) => function (_2357_i6) {
              return _2356_v37;
            })(_2343_v37)));
            let _2358_v50;
            _2358_v50 = _dafny.Map.Empty.slice().updateUnsafe(_2353_v47,(_2355_v49).Difference(_2355_v49));
            _2358_v50 = (_2358_v50).update(_2353_v47, _dafny.Set.fromElements(_2350_v44));
            let _2359_v51;
            let _nw364 = new _module.C2();
            _nw364.__ctor();
            _2359_v51 = _nw364;
          }
        }
      }
      let _2360_v52;
      _2360_v52 = _module.D16.create_DC44((_this).f2);
      let _pat_let_tv62 = p2;
      let _pat_let_tv63 = p2;
      let _pat_let_tv64 = p2;
      let _pat_let_tv65 = p2;
      let _pat_let_tv66 = p2;
      let _pat_let_tv67 = p2;
      let _pat_let_tv68 = p2;
      let _pat_let_tv69 = p2;
      let _pat_let_tv70 = p2;
      r2 = function (_source25) {
        if (_source25.is_DC42) {
          let _2361___mcc_h0 = (_source25).cf76;
          let _2362___mcc_h1 = (_source25).cf77;
          let _2363___mcc_h2 = (_source25).cf78;
          let _2364_cf78 = _2363___mcc_h2;
          let _2365_cf77 = _2362___mcc_h1;
          let _2366_cf76 = _2361___mcc_h0;
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_pat_let_tv62)).length));
        } else if (_source25.is_DC43) {
          let _2367___mcc_h3 = (_source25).cf79;
          let _2368___mcc_h4 = (_source25).cf80;
          let _2369___mcc_h5 = (_source25).cf81;
          let _2370_cf81 = _2369___mcc_h5;
          let _2371_cf80 = _2368___mcc_h4;
          let _2372_cf79 = _2367___mcc_h3;
          return new BigNumber((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,_pat_let_tv63), _dafny.Map.Empty.slice().updateUnsafe(!(_pat_let_tv64),_pat_let_tv65), _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv66,_pat_let_tv67), _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv68,true), _dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv69,_pat_let_tv70))).length);
        } else if (_source25.is_DC44) {
          let _2373___mcc_h6 = (_source25).cf82;
          let _2374_cf82 = _2373___mcc_h6;
          return _module.__default.safeDivisionInt((_this).f2, (_this).f2);
        } else {
          let _2375___mcc_h7 = (_source25).cf75;
          let _2376_cf75 = _2375___mcc_h7;
          return (_this).f2;
        }
      }(_2360_v52);
      let _2377_v53;
      let _nw365 = Array((new BigNumber(7)).toNumber()).fill(false);
      _2377_v53 = _nw365;
      let _index371 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length));
      (_2377_v53)[_index371] = p2;
      let _2378_v54;
      _2378_v54 = true;
      let _index372 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length));
      let _rhs289 = _2378_v54;
      let _rhs290 = !(p2);
      let _lhs160 = _2377_v53;
      let _lhs161 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length));
      _lhs160[_lhs161] = _rhs289;
      _2378_v54 = _rhs290;
      let _hi16 = (((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))]) ? ((_dafny.ZERO).minus(new BigNumber(-773))) : (p0));
      for (let _2379_i7 = (new BigNumber(945)).minus(p0); _2379_i7.isLessThan(_hi16); _2379_i7 = _2379_i7.plus(_dafny.ONE)) {
        _2378_v54 = ((p2) ? (((_this).f2).isLessThanOrEqualTo(p0)) : ((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))]));
        let _2380_v55;
        _2380_v55 = _module.D12.create_DC29(true, p0, p2, _2379_i7, (_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))]);
        let _index373 = _module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length));
        (_2377_v53)[_index373] = ((((p2) ? ((_2380_v55).dtor_cf50) : ((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))]))) ? (p2) : ((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))]));
        let _2381_v56;
        let _nw366 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2381_v56 = _nw366;
        _2381_v56 = _2381_v56;
        let _2382_v57;
        let _nw367 = new _module.C2();
        _nw367.__ctor();
        _2382_v57 = _nw367;
      }
      let _2383_v58;
      _2383_v58 = _module.D12.create_DC28(true, (_this).f2);
      let _2384_v59;
      _2384_v59 = _dafny.Map.Empty.slice().updateUnsafe(!((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))]),(_2383_v58).dtor_cf46);
      let _2385_v61;
      _2385_v61 = _dafny.Seq.of(_2378_v54);
      let _2386_v62;
      _2386_v62 = _dafny.Map.Empty.slice().updateUnsafe(_2385_v61,p2);
      let _2387_v63;
      _2387_v63 = _dafny.Map.Empty.slice().updateUnsafe((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))],function () {
        let _coll85 = new _dafny.Map();
        for (const _compr_85 of (_2386_v62).Keys.Elements) {
          let _2388_v60 = _compr_85;
          if ((_2386_v62).contains(_2388_v60)) {
            _coll85.push([_2388_v60,(_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))]]);
          }
        }
        return _coll85;
      }());
      let _2389_v64;
      _2389_v64 = _dafny.Set.fromElements(_2378_v54);
      let _2390_v65;
      _2390_v65 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _2391_v66;
      let _init58 = ((_2392_p0) => function (_2393_i8) {
        return _module.__default.safeDivisionInt(_2393_i8, _2392_p0);
      })(p0);
      let _nw368 = Array((new BigNumber(18)).toNumber());
      for (let _i0_58 = 0; _i0_58 < new BigNumber(_nw368.length); _i0_58++) {
        _nw368[_i0_58] = _init58(new BigNumber(_i0_58));
      }
      _2391_v66 = _nw368;
      let _2394_v67;
      _2394_v67 = new _dafny.CodePoint('i'.codePointAt(0));
      let _2395_v68;
      _2395_v68 = _dafny.Map.Empty.slice().updateUnsafe(_2391_v66,_module.__default.fm65((_dafny.ZERO).minus((_this).f2), _2394_v67, p0, globalState));
      let _2396_v69;
      _2396_v69 = _dafny.Seq.UnicodeFromString("haw");
      let _2397_v70;
      let _nw369 = new _module.C0();
      _nw369.__ctor(_2396_v69);
      _2397_v70 = _nw369;
      let _2398_v71;
      _2398_v71 = _dafny.Seq.of(_2397_v70);
      let _2399_v72;
      let _nw370 = Array((new BigNumber(27)).toNumber());
      _nw370[(_dafny.ZERO).toNumber()] = ((_this).f2).minus(p0);
      _nw370[(_dafny.ONE).toNumber()] = (new BigNumber(941)).multipliedBy((_this).f2);
      _nw370[(new BigNumber(2)).toNumber()] = p0;
      _nw370[(new BigNumber(3)).toNumber()] = (_this).f2;
      _nw370[(new BigNumber(4)).toNumber()] = new BigNumber(((_this).f0).length);
      _nw370[(new BigNumber(5)).toNumber()] = ((_this).f2).plus((_this).f2);
      _nw370[(new BigNumber(6)).toNumber()] = p0;
      _nw370[(new BigNumber(7)).toNumber()] = (_this).f2;
      _nw370[(new BigNumber(8)).toNumber()] = (new BigNumber((_dafny.Seq.of((((_2384_v59).contains((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))])) ? ((_2384_v59).get((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))])) : (p2)))).length)).plus(new BigNumber(323));
      _nw370[(new BigNumber(9)).toNumber()] = (_this).f2;
      _nw370[(new BigNumber(10)).toNumber()] = (new BigNumber(((((_2387_v63).contains(!(p2))) ? ((_2387_v63).get(!(p2))) : (_dafny.Map.Empty.slice().updateUnsafe(_2385_v61,_2378_v54)))).length)).plus((_this).f2);
      _nw370[(new BigNumber(11)).toNumber()] = new BigNumber(((_2389_v64).Difference(_2389_v64)).length);
      _nw370[(new BigNumber(12)).toNumber()] = (_this).f2;
      _nw370[(new BigNumber(13)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_2390_v65).length), new BigNumber(((((_2395_v68).contains(_2391_v66)) ? ((_2395_v68).get(_2391_v66)) : (_dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f2)))).length));
      _nw370[(new BigNumber(14)).toNumber()] = _module.__default.safeDivisionInt((_this).f2, new BigNumber(355));
      _nw370[(new BigNumber(15)).toNumber()] = (p0).multipliedBy(new BigNumber((_2396_v69).length));
      _nw370[(new BigNumber(16)).toNumber()] = (p0).minus(new BigNumber((_dafny.Set.fromElements((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))], !((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))]), p2, p2)).length));
      _nw370[(new BigNumber(17)).toNumber()] = (_dafny.ZERO).minus((_this).f2);
      _nw370[(new BigNumber(18)).toNumber()] = p0;
      _nw370[(new BigNumber(19)).toNumber()] = (((_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))]) ? (new BigNumber((_2398_v71).length)) : ((_this).f2));
      _nw370[(new BigNumber(20)).toNumber()] = (p0).multipliedBy(_module.__default.fm32(new BigNumber(502), globalState));
      _nw370[(new BigNumber(21)).toNumber()] = (_dafny.ZERO).minus((new BigNumber((_2390_v65).length)).plus(p0));
      _nw370[(new BigNumber(22)).toNumber()] = p0;
      _nw370[(new BigNumber(23)).toNumber()] = (_this).f2;
      _nw370[(new BigNumber(24)).toNumber()] = new BigNumber(((_this).f0).length);
      _nw370[(new BigNumber(25)).toNumber()] = (new BigNumber((_dafny.Seq.of((_this).f2)).length)).minus(new BigNumber(-469));
      _nw370[(new BigNumber(26)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(p0)).length), (_this).f2);
      _2399_v72 = _nw370;
      let _2400_v73;
      _2400_v73 = _module.D12.create_DC29(_2378_v54, new BigNumber((_2385_v61).length), (_2377_v53)[_module.__default.safeIndex(new BigNumber(674), new BigNumber((_2377_v53).length))], new BigNumber(196), p2);
      let _index374 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_2399_v72).length));
      (_2399_v72)[_index374] = (_2400_v73).dtor_cf49;
      let _index375 = _module.__default.safeIndex(new BigNumber(644), new BigNumber((_2399_v72).length));
      (_2399_v72)[_index375] = (((_dafny.ZERO).minus(p0)).plus(new BigNumber(447))).plus(new BigNumber(383));
      let _2401_v74;
      _2401_v74 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_dafny.ZERO).minus((_2399_v72)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_2399_v72).length))]), (_this).f2),_2377_v53);
      r0 = new BigNumber((_2401_v74).length);
      let _2402_v75;
      _2402_v75 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f0)[_module.__default.safeIndex((_2399_v72)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_2399_v72).length))], new BigNumber(((_this).f0).length))],_2394_v67);
      r1 = (_2402_v75).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2396_v69).length),_2394_v67));
      r2 = _module.__default.safeModuloInt((_2399_v72)[_module.__default.safeIndex(new BigNumber(644), new BigNumber((_2399_v72).length))], p0);
      return [r0, r1, r2];
    }
    m4(p0, p1, p2, p3, globalState) {
      let _this = this;
      let _2403_v0;
      let _init59 = function (_2404_i0) {
        return false;
      };
      let _nw371 = Array((new BigNumber(23)).toNumber());
      for (let _i0_59 = 0; _i0_59 < new BigNumber(_nw371.length); _i0_59++) {
        _nw371[_i0_59] = _init59(new BigNumber(_i0_59));
      }
      _2403_v0 = _nw371;
      let _2405_v1;
      let _nw372 = Array((new BigNumber(7)).toNumber());
      _nw372[(_dafny.ZERO).toNumber()] = _2403_v0;
      _nw372[(_dafny.ONE).toNumber()] = _2403_v0;
      _nw372[(new BigNumber(2)).toNumber()] = _2403_v0;
      _nw372[(new BigNumber(3)).toNumber()] = _2403_v0;
      _nw372[(new BigNumber(4)).toNumber()] = _2403_v0;
      _nw372[(new BigNumber(5)).toNumber()] = _2403_v0;
      _nw372[(new BigNumber(6)).toNumber()] = _2403_v0;
      _2405_v1 = _nw372;
      let _2406_v2;
      _2406_v2 = new _dafny.CodePoint('a'.codePointAt(0));
      let _2407_v3;
      _2407_v3 = _dafny.Map.Empty.slice().updateUnsafe(_2406_v2,_this.f1);
      let _2408_v4;
      _2408_v4 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm5(p3, !(p0), _2407_v3, p0, globalState),_2405_v1);
      _2405_v1 = (((_2408_v4).contains((new BigNumber(((_this).f0).length)).isEqualTo(p3))) ? ((_2408_v4).get((new BigNumber(((_this).f0).length)).isEqualTo(p3))) : (_2405_v1));
      let _2409_v5;
      _2409_v5 = _dafny.Seq.of(!(p0), p0, p0);
      let _2410_v6;
      _2410_v6 = _dafny.Seq.of(_2409_v5);
      if (!_dafny.Seq.contains(_dafny.Seq.Concat(_dafny.Seq.of(_2409_v5), _2410_v6), _dafny.Seq.of(p0))) {
        let _2411_v7;
        _2411_v7 = _dafny.Set.fromElements(_2403_v0);
        let _2412_v8;
        _2412_v8 = _2411_v7;
        _2412_v8 = _2412_v8;
        let _2413_v9;
        _2413_v9 = false;
        _2413_v9 = _2413_v9;
        if (p0) {
          let _2414_v10;
          _2414_v10 = new BigNumber(249);
          _2414_v10 = p1;
          let _2415_v11;
          _2415_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(868),!(p0));
          let _2416_v12;
          _2416_v12 = _dafny.Map.Empty.slice().updateUnsafe(_2403_v0,!(_2413_v9));
          let _2417_v13;
          let _nw373 = Array((new BigNumber(26)).toNumber());
          _nw373[(_dafny.ZERO).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_2403_v0,(((_2415_v11).contains(new BigNumber(601))) ? ((_2415_v11).get(new BigNumber(601))) : (_2413_v9)));
          _nw373[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2403_v0,p0)).Merge(_2416_v12);
          _nw373[(new BigNumber(2)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(3)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2403_v0,p0)).Merge(_2416_v12);
          _nw373[(new BigNumber(4)).toNumber()] = ((p0) ? (_2416_v12) : (_dafny.Map.Empty.slice().updateUnsafe(_2403_v0,p0)));
          _nw373[(new BigNumber(5)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(6)).toNumber()] = (_2416_v12).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2403_v0,_2413_v9));
          _nw373[(new BigNumber(7)).toNumber()] = (_2416_v12).Merge(_2416_v12);
          _nw373[(new BigNumber(8)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2403_v0,p0)).Merge(_2416_v12);
          _nw373[(new BigNumber(9)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(10)).toNumber()] = (_2416_v12).Merge(_2416_v12);
          _nw373[(new BigNumber(11)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(12)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(13)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(14)).toNumber()] = (_2416_v12).Merge(_2416_v12);
          _nw373[(new BigNumber(15)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(16)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(17)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(18)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(19)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(20)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(21)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_2403_v0,p0)).Merge(_2416_v12);
          _nw373[(new BigNumber(22)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(23)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(24)).toNumber()] = _2416_v12;
          _nw373[(new BigNumber(25)).toNumber()] = _2416_v12;
          _2417_v13 = _nw373;
          let _index376 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2417_v13).length));
          (_2417_v13)[_index376] = _2416_v12;
          let _2418_v14;
          _2418_v14 = _dafny.Map.Empty.slice().updateUnsafe(p1,(_dafny.ZERO).minus((_this).f2));
          let _2419_v15;
          let _init60 = ((_2420_v10) => function (_2421_i1) {
            return (_2421_i1).multipliedBy(_2420_v10);
          })(_2414_v10);
          let _nw374 = Array((new BigNumber(28)).toNumber());
          for (let _i0_60 = 0; _i0_60 < new BigNumber(_nw374.length); _i0_60++) {
            _nw374[_i0_60] = _init60(new BigNumber(_i0_60));
          }
          _2419_v15 = _nw374;
          let _2422_v16;
          let _nw375 = new _module.C5();
          _nw375.__ctor(_2419_v15, _dafny.Seq.of((_this).f2, p1), _this.f1);
          _2422_v16 = _nw375;
          let _2423_v17;
          let _nw376 = Array((new BigNumber(21)).toNumber());
          _nw376[(_dafny.ZERO).toNumber()] = _2422_v16;
          _nw376[(_dafny.ONE).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(2)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(3)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(4)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(5)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(6)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(7)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(8)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(9)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(10)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(11)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(12)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(13)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(14)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(15)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(16)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(17)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(18)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(19)).toNumber()] = _2422_v16;
          _nw376[(new BigNumber(20)).toNumber()] = _2422_v16;
          _2423_v17 = _nw376;
          let _2424_v18;
          _2424_v18 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2423_v17);
          let _2425_v19;
          _2425_v19 = _dafny.Seq.of((((_2424_v18).contains(p1)) ? ((_2424_v18).get(p1)) : (_2423_v17)));
          let _2426_v20;
          _2426_v20 = _module.D12.create_DC29(_2413_v9, p3, !(p0), new BigNumber((_2425_v19).length), _module.__default.fm30(_2414_v10, false, p0, globalState));
          let _index377 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2417_v13).length));
          let _rhs291 = (_module.__default.safeModuloInt(new BigNumber((_2418_v14).length), new BigNumber(832))).minus((_2426_v20).dtor_cf49);
          let _rhs292 = p0;
          let _rhs293 = _2406_v2;
          let _rhs294 = p0;
          let _rhs295 = ((_2416_v12).Merge(_2416_v12)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_2403_v0,p0));
          let _lhs162 = _2417_v13;
          let _lhs163 = _module.__default.safeIndex(new BigNumber(351), new BigNumber((_2417_v13).length));
          _2414_v10 = _rhs291;
          _2413_v9 = _rhs292;
          _2406_v2 = _rhs293;
          _2413_v9 = _rhs294;
          _lhs162[_lhs163] = _rhs295;
          let _2427_v21;
          _2427_v21 = _dafny.Seq.of(((_this).f2).minus(p3), p1, p1);
          _2427_v21 = _dafny.Seq.Concat((_this).f0, _2427_v21);
          _2413_v9 = _2413_v9;
          _2414_v10 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_2427_v21, _2427_v21), (_this).f0)).length);
        } else {
          let _2428_v22;
          _2428_v22 = _dafny.Seq.UnicodeFromString("tqhfxuca");
          let _2429_v23;
          let _nw377 = Array((new BigNumber(4)).toNumber());
          _nw377[(_dafny.ZERO).toNumber()] = p1;
          _nw377[(_dafny.ONE).toNumber()] = (_this).f2;
          _nw377[(new BigNumber(2)).toNumber()] = (_this).f2;
          _nw377[(new BigNumber(3)).toNumber()] = new BigNumber((_2428_v22).length);
          _2429_v23 = _nw377;
          let _2430_v24;
          let _nw378 = new _module.C5();
          _nw378.__ctor(_2429_v23, _dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), ((_2431_p3) => function (_2432_i2) {
            return _2431_p3;
          })(p3)), _this.f1);
          _2430_v24 = _nw378;
          let _2433_v25;
          _2433_v25 = _2430_v24;
          let _2434_v26;
          _2434_v26 = _dafny.Set.fromElements((_2433_v25), _2430_v24);
          let _2435_v27;
          _2435_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2409_v5)[_module.__default.safeIndex((_this).f2, new BigNumber((_2409_v5).length))]);
          let _2436_v28;
          _2436_v28 = _dafny.Map.Empty.slice().updateUnsafe((_2434_v26).Union(_2434_v26),new BigNumber(((_2435_v27).Merge(_2435_v27)).length));
          _2436_v28 = (_2436_v28).update(_2434_v26, (_this).f2);
          let _rhs296 = ((_2413_v9) ? (p0) : (_2413_v9));
          let _rhs297 = p0;
          _2413_v9 = _rhs296;
          _2413_v9 = _rhs297;
          let _2437_v29;
          _2437_v29 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(((_this).f0)[_module.__default.safeIndex(p3, new BigNumber(((_this).f0).length))]),_dafny.Seq.update(_2428_v22, _module.__default.safeIndex((_this).f2, new BigNumber((_2428_v22).length)), _2406_v2));
          let _2438_v30;
          _2438_v30 = _module.D14.create_DC34(_2437_v29);
          let _2439_v32;
          _2439_v32 = _dafny.Map.Empty.slice().updateUnsafe((_this).f2,p0);
          let _2440_v33;
          _2440_v33 = _dafny.Map.Empty.slice().updateUnsafe(p1,_2406_v2);
          let _2441_v35;
          _2441_v35 = _dafny.Set.fromElements(p3, p1);
          let _pat_let_tv71 = p1;
          let _pat_let_tv72 = _2428_v22;
          let _pat_let_tv73 = _2437_v29;
          let _pat_let_tv74 = _2437_v29;
          let _2442_v36;
          let _nw379 = Array((new BigNumber(23)).toNumber());
          _nw379[(_dafny.ZERO).toNumber()] = _2438_v30;
          _nw379[(_dafny.ONE).toNumber()] = _module.D14.create_DC34(function () {
  let _coll86 = new _dafny.Map();
  for (const _compr_86 of (_2439_v32).Keys.Elements) {
    let _2443_v31 = _compr_86;
    if ((_2439_v32).contains(_2443_v31)) {
      _coll86.push([(_2443_v31).multipliedBy(p3),_2428_v22]);
    }
  }
  return _coll86;
}());
          _nw379[(new BigNumber(2)).toNumber()] = function (_pat_let51_0) {
            return function (_2444_dt__update__tmp_h0) {
              return function (_pat_let52_0) {
                return function (_2445_dt__update_hcf62_h0) {
                  return _module.D14.create_DC34(_2445_dt__update_hcf62_h0);
                }(_pat_let52_0);
              }(_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv71,_pat_let_tv72));
            }(_pat_let51_0);
          }(_module.D14.create_DC34(_2437_v29));
          _nw379[(new BigNumber(3)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(4)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(5)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(6)).toNumber()] = _module.__default.fm68(p0, _2428_v22, globalState);
          _nw379[(new BigNumber(7)).toNumber()] = _module.__default.fm68(!(p0), _2428_v22, globalState);
          _nw379[(new BigNumber(8)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(9)).toNumber()] = _module.D14.create_DC34(_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_2440_v33).length)),_2428_v22));
          _nw379[(new BigNumber(10)).toNumber()] = _module.D14.create_DC34(function () {
  let _coll87 = new _dafny.Map();
  for (const _compr_87 of (_2441_v35).Elements) {
    let _2446_v34 = _compr_87;
    if ((_2441_v35).contains(_2446_v34)) {
      _coll87.push([_module.__default.safeModuloInt(_2446_v34, p3),_dafny.Seq.Create(_module.__default.abs(new BigNumber(-583)), ((_2447_v2) => function (_2448_i3) {
        return _2447_v2;
      })(_2406_v2))]);
    }
  }
  return _coll87;
}());
          _nw379[(new BigNumber(11)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(12)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(13)).toNumber()] = _module.__default.fm68(p0, _2428_v22, globalState);
          _nw379[(new BigNumber(14)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(15)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(16)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(17)).toNumber()] = function (_pat_let53_0) {
            return function (_2449_dt__update__tmp_h1) {
              return function (_pat_let54_0) {
                return function (_2450_dt__update_hcf62_h1) {
                  return _module.D14.create_DC34(_2450_dt__update_hcf62_h1);
                }(_pat_let54_0);
              }(_pat_let_tv73);
            }(_pat_let53_0);
          }(_2438_v30);
          _nw379[(new BigNumber(18)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(19)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(20)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(21)).toNumber()] = _2438_v30;
          _nw379[(new BigNumber(22)).toNumber()] = function (_pat_let55_0) {
            return function (_2451_dt__update__tmp_h2) {
              return function (_pat_let56_0) {
                return function (_2452_dt__update_hcf62_h2) {
                  return _module.D14.create_DC34(_2452_dt__update_hcf62_h2);
                }(_pat_let56_0);
              }(_pat_let_tv74);
            }(_pat_let55_0);
          }(_2438_v30);
          _2442_v36 = _nw379;
          let _index378 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_2442_v36).length));
          (_2442_v36)[_index378] = _2438_v30;
          let _index379 = _module.__default.safeIndex(new BigNumber(751), new BigNumber((_2442_v36).length));
          (_2442_v36)[_index379] = _2438_v30;
          _2439_v32 = (_module.__default.fm21(globalState)).update(p1, _2413_v9);
          _2413_v9 = p0;
        }
        let _2453_v37;
        let _nw380 = new _module.C3();
        _nw380.__ctor();
        _2453_v37 = _nw380;
        let _2454_v38;
        let _nw381 = Array((new BigNumber(26)).toNumber()).fill(_module.D2.Default());
        _2454_v38 = _nw381;
        let _2455_v39;
        _2455_v39 = _dafny.Seq.of(_2406_v2);
        let _2456_v40;
        _2456_v40 = _dafny.MultiSet.fromElements(_2455_v39);
        let _2457_v42;
        _2457_v42 = _module.D2.create_DC3(function () {
  let _coll88 = new _dafny.Set();
  for (const _compr_88 of (_2456_v40).Elements) {
    let _2458_v41 = _compr_88;
    if ((_2456_v40).contains(_2458_v41)) {
      _coll88.add(_2458_v41);
    }
  }
  return _coll88;
}());
        let _pat_let_tv75 = _2455_v39;
        let _index380 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_2454_v38).length));
        (_2454_v38)[_index380] = function (_pat_let57_0) {
          return function (_2459_dt__update__tmp_h3) {
            return function (_pat_let58_0) {
              return function (_2460_dt__update_hcf4_h0) {
                return _module.D2.create_DC3(_2460_dt__update_hcf4_h0);
              }(_pat_let58_0);
            }(_dafny.Set.fromElements(_pat_let_tv75));
          }(_pat_let57_0);
        }(_2457_v42);
        let _index381 = _module.__default.safeIndex(new BigNumber(580), new BigNumber((_2454_v38).length));
        (_2454_v38)[_index381] = _2457_v42;
      } else {
        let _index382 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length));
        (_2403_v0)[_index382] = p0;
        let _2461_v43;
        _2461_v43 = new BigNumber(950);
        let _2462_v44;
        _2462_v44 = _dafny.Seq.UnicodeFromString("wvwvfkb");
        let _2463_v45;
        let _nw382 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _2463_v45 = _nw382;
        let _index383 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length));
        (_2463_v45)[_index383] = _module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f0,(_this).f2)).length), new BigNumber(232));
        let _2464_v46;
        _2464_v46 = _dafny.Set.fromElements(false, p0);
        let _2465_v47;
        _2465_v47 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2464_v46).length),_2463_v45);
        let _index384 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length));
        let _index385 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length));
        let _rhs298 = p0;
        let _rhs299 = _module.__default.fm20(!(p0), globalState);
        let _rhs300 = _dafny.Seq.update(_2462_v44, _module.__default.safeIndex(p3, new BigNumber((_2462_v44).length)), ((p0) ? (new _dafny.CodePoint('u'.codePointAt(0))) : (new _dafny.CodePoint('q'.codePointAt(0)))));
        let _rhs301 = (((_2465_v47).contains((_this).f2)) ? (_module.__default.safeDivisionInt(_module.__default.fm44(p0, globalState), _2461_v43)) : (p3));
        let _lhs164 = _2403_v0;
        let _lhs165 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length));
        let _lhs166 = _2463_v45;
        let _lhs167 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length));
        _lhs164[_lhs165] = _rhs298;
        _2461_v43 = _rhs299;
        _2462_v44 = _rhs300;
        _lhs166[_lhs167] = _rhs301;
        if ((_2403_v0)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length))]) {
          let _index386 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length));
          (_2463_v45)[_index386] = (_this).f2;
          let _index387 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length));
          (_2403_v0)[_index387] = !((_2403_v0)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length))]);
          let _index388 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length));
          (_2463_v45)[_index388] = (_dafny.ZERO).minus(((_this).f0)[_module.__default.safeIndex((_this).f2, new BigNumber(((_this).f0).length))]);
          let _2466_v48;
          let _nw383 = new _module.C3();
          _nw383.__ctor();
          _2466_v48 = _nw383;
          _2466_v48 = _2466_v48;
          _2462_v44 = _dafny.Seq.Concat(_2462_v44, _dafny.Seq.Concat(_2462_v44, _dafny.Seq.UnicodeFromString("oo")));
        } else {
          let _index389 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length));
          (_2403_v0)[_index389] = ((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_2462_v44).length), _2461_v43))).isEqualTo(((true) ? ((_this).f2) : (_module.__default.fm20(_module.__default.fm17(globalState), globalState))));
          let _2467_v49;
          _2467_v49 = _dafny.Map.Empty.slice().updateUnsafe(_2461_v43,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-869)), function (_2468_i4) {
            return new BigNumber(538);
          }));
          let _index390 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length));
          (_2463_v45)[_index390] = (((_2403_v0)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length))]) ? ((_this).f2) : (_module.__default.safeDivisionInt(new BigNumber(((((_2467_v49).contains(p1)) ? ((_2467_v49).get(p1)) : ((_this).f0))).length), new BigNumber((_dafny.Set.fromElements((_this).f2)).length))));
          let _2469_v50;
          _2469_v50 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_2409_v5).length),new BigNumber(930));
          let _2470_v51;
          _2470_v51 = _dafny.MultiSet.fromElements(new BigNumber((_2469_v50).length));
          let _2471_v52;
          _2471_v52 = _dafny.MultiSet.fromElements((_2463_v45)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length))], (((_2469_v50).contains(new BigNumber(-461))) ? ((_2469_v50).get(new BigNumber(-461))) : ((_this).f2)), p3, (_this).f2, (new BigNumber((_2470_v51).cardinality())).plus(p3));
          let _index391 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length));
          (_2463_v45)[_index391] = (((_2470_v51).contains((new BigNumber(936)).multipliedBy(p3))) ? ((_2470_v51).get((new BigNumber(936)).multipliedBy(p3))) : (_module.__default.fm44(p0, globalState)));
          let _index392 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length));
          (_2463_v45)[_index392] = ((_this).f0)[_module.__default.safeIndex(new BigNumber(404), new BigNumber(((_this).f0).length))];
          let _2472_v53;
          let _nw384 = new _module.C4();
          _nw384.__ctor();
          _2472_v53 = _nw384;
          let _2473_v54;
          _2473_v54 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2472_v53);
          _2473_v54 = (_2473_v54).update((_2403_v0)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length))], _2472_v53);
        }
        let _2474_v55;
        _2474_v55 = _dafny.Map.Empty.slice().updateUnsafe(p3,_2406_v2);
        _2406_v2 = (((_2474_v55).contains(p1)) ? ((_2474_v55).get(p1)) : (_2406_v2));
        let _2475_v56;
        _2475_v56 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_2463_v45)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length))]);
        let _2476_v57;
        _2476_v57 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(303),_2475_v56);
        if (!(new BigNumber(((_2476_v57).update(p3, _2475_v56)).length)).isEqualTo(p3)) {
          let _2477_v58;
          let _nw385 = new _module.C2();
          _nw385.__ctor();
          _2477_v58 = _nw385;
          let _2478_v59;
          _2478_v59 = _dafny.Set.fromElements(_2477_v58);
          let _2479_v60;
          let _nw386 = Array((new BigNumber(15)).toNumber()).fill([]);
          _2479_v60 = _nw386;
          let _2480_v61;
          _2480_v61 = _dafny.Map.Empty.slice().updateUnsafe((_2478_v59).Union(_2478_v59),_2479_v60);
          _2480_v61 = (_2480_v61).update(_dafny.Set.fromElements(_2477_v58), _2479_v60);
          _2461_v43 = _2461_v43;
          let _2481_v62;
          _2481_v62 = _dafny.Seq.of((_this).f2);
          _2481_v62 = (_module.D2.create_DC5(_2461_v43, _2481_v62, false)).dtor_cf9;
          let _2482_v63;
          let _init61 = ((_2483_v2) => function (_2484_i5) {
            return _2483_v2;
          })(_2406_v2);
          let _nw387 = Array((new BigNumber(5)).toNumber());
          for (let _i0_61 = 0; _i0_61 < new BigNumber(_nw387.length); _i0_61++) {
            _nw387[_i0_61] = _init61(new BigNumber(_i0_61));
          }
          _2482_v63 = _nw387;
          let _2485_v64;
          _2485_v64 = _dafny.Seq.of(_2482_v63);
          _2485_v64 = _2485_v64;
        } else {
          let _index393 = _module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length));
          (_2463_v45)[_index393] = ((p0) ? (_2461_v43) : ((_2463_v45)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length))]));
          _2462_v44 = _2462_v44;
          let _2486_v65;
          _2486_v65 = _dafny.Seq.of((_2463_v45)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length))]);
          _2486_v65 = _2486_v65;
          _2406_v2 = (_2462_v44)[_module.__default.safeIndex(_module.__default.safeModuloInt(p3, new BigNumber((_dafny.MultiSet.fromElements(_2462_v44)).cardinality())), new BigNumber((_2462_v44).length))];
          let _2487_v66;
          let _nw388 = Array((new BigNumber(25)).toNumber()).fill([]);
          _2487_v66 = _nw388;
          let _2488_v67;
          let _nw389 = Array((new BigNumber(29)).toNumber()).fill(_dafny.Map.Empty);
          _2488_v67 = _nw389;
          let _index394 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_2487_v66).length));
          (_2487_v66)[_index394] = (((_2409_v5)[_module.__default.safeIndex(p3, new BigNumber((_2409_v5).length))]) ? (_2488_v67) : (_2488_v67));
          let _2489_v68;
          _2489_v68 = _dafny.Set.fromElements(new BigNumber((_2462_v44).length));
          let _2490_v69;
          _2490_v69 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of((_2463_v45)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length))])).length),(_2403_v0)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length))]);
          let _2491_v70;
          _2491_v70 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_2462_v44).length), (_2463_v45)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length))], (_2463_v45)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length))], (_2463_v45)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length))], p1),_2490_v69);
          let _index395 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_2487_v66).length));
          let _index396 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length));
          let _rhs302 = _2488_v67;
          let _rhs303 = (_this).f2;
          let _rhs304 = !((_2491_v70).Merge(_2491_v70)).contains(_2489_v68);
          let _lhs168 = _2487_v66;
          let _lhs169 = _module.__default.safeIndex(new BigNumber(29), new BigNumber((_2487_v66).length));
          let _lhs170 = _2403_v0;
          let _lhs171 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length));
          _lhs168[_lhs169] = _rhs302;
          _2461_v43 = _rhs303;
          _lhs170[_lhs171] = _rhs304;
        }
        let _2492_v71;
        _2492_v71 = _module.D13.create_DC32(_2462_v44, p0, _2461_v43, _2463_v45);
        let _2493_v72;
        _2493_v72 = _dafny.MultiSet.fromElements(_2462_v44, (_2492_v71).dtor_cf57);
        let _index397 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_2403_v0).length));
        (_2403_v0)[_index397] = (_this).fm5(p3, p0, _2407_v3, ((_2463_v45)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length))]).isLessThan((((_2493_v72).contains(_dafny.Seq.UnicodeFromString("cu"))) ? ((_2493_v72).get(_dafny.Seq.UnicodeFromString("cu"))) : ((_2463_v45)[_module.__default.safeIndex(new BigNumber(315), new BigNumber((_2463_v45).length))]))), globalState);
      }
      let _2494_v73;
      let _nw390 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _2494_v73 = _nw390;
      let _index398 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_2494_v73).length));
      (_2494_v73)[_index398] = p3;
      let _index399 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_2494_v73).length));
      (_2494_v73)[_index399] = (p1).multipliedBy((_this).f2);
      let _index400 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_2403_v0).length));
      (_2403_v0)[_index400] = p0;
      let _2495_v74;
      _2495_v74 = _dafny.Set.fromElements((_this).f0, _dafny.Seq.of(p3));
      let _2496_v75;
      _2496_v75 = _module.D15.create_DC37(_2495_v74);
      let _2497_v76;
      _2497_v76 = _dafny.Map.Empty.slice().updateUnsafe((_2494_v73)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_2494_v73).length))],_2496_v75);
      let _2498_v77;
      let _nw391 = Array((new BigNumber(12)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _2498_v77 = _nw391;
      let _2499_v78;
      _2499_v78 = _module.D20.create_DC52(_2498_v77);
      let _2500_v79;
      let _nw392 = Array((new BigNumber(22)).toNumber());
      _nw392[(_dafny.ZERO).toNumber()] = _2498_v77;
      _nw392[(_dafny.ONE).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(2)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(3)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(4)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(5)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(6)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(7)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(8)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(9)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(10)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(11)).toNumber()] = (_2499_v78).dtor_cf93;
      _nw392[(new BigNumber(12)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(13)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(14)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(15)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(16)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(17)).toNumber()] = (_2499_v78).dtor_cf93;
      _nw392[(new BigNumber(18)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(19)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(20)).toNumber()] = _2498_v77;
      _nw392[(new BigNumber(21)).toNumber()] = _2498_v77;
      _2500_v79 = _nw392;
      let _index401 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_2500_v79).length));
      (_2500_v79)[_index401] = _2498_v77;
      let _2501_v81;
      _2501_v81 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements((_this).f2, (_2494_v73)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_2494_v73).length))]),_module.__default.fm69(false, globalState));
      let _2502_v82;
      _2502_v82 = _dafny.Set.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(p3)));
      let _2503_v83;
      _2503_v83 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f2);
      let _index402 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_2403_v0).length));
      let _index403 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_2500_v79).length));
      let _index404 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_2494_v73).length));
      let _rhs305 = p0;
      let _rhs306 = (function () {
        let _coll89 = new _dafny.Map();
        for (const _compr_89 of _dafny.IntegerRange(new BigNumber(311), new BigNumber(702))) {
          let _2504_v80 = _compr_89;
          if (((new BigNumber(311)).isLessThanOrEqualTo(_2504_v80)) && ((_2504_v80).isLessThan(new BigNumber(702)))) {
            _coll89.push([(_2504_v80).minus(new BigNumber(-843)),_2496_v75]);
          }
        }
        return _coll89;
      }()).Merge((((_2501_v81).contains(_2502_v82)) ? ((_2501_v81).get(_2502_v82)) : (_2497_v76)));
      let _rhs307 = _2498_v77;
      let _rhs308 = ((_this).f0)[_module.__default.safeIndex((((_2503_v83).contains(false)) ? ((_2503_v83).get(false)) : ((_this).f2)), new BigNumber(((_this).f0).length))];
      let _lhs172 = _2403_v0;
      let _lhs173 = _module.__default.safeIndex(new BigNumber(846), new BigNumber((_2403_v0).length));
      let _lhs174 = _2500_v79;
      let _lhs175 = _module.__default.safeIndex(new BigNumber(419), new BigNumber((_2500_v79).length));
      let _lhs176 = _2494_v73;
      let _lhs177 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_2494_v73).length));
      _lhs172[_lhs173] = _rhs305;
      _2497_v76 = _rhs306;
      _lhs174[_lhs175] = _rhs307;
      _lhs176[_lhs177] = _rhs308;
      _2406_v2 = new _dafny.CodePoint('x'.codePointAt(0));
      let _2505_v84;
      let _nw393 = new _module.C4();
      _nw393.__ctor();
      _2505_v84 = _nw393;
      return;
    }
    get f2() {
      let _this = this;
      return _this._f2;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
  };

  $module.C9 = class C9 {
    constructor () {
      this._tname = "_module.C9";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm0(p0, p1, globalState) {
      let _this = this;
      return new BigNumber(63);
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.of();
      let _2506_v0;
      _2506_v0 = false;
      let _2507_v1;
      _2507_v1 = _dafny.MultiSet.fromElements(_2506_v0);
      r0 = (_this).fm0(_2506_v0, new BigNumber((_2507_v1).cardinality()), globalState);
      let _2508_v2;
      _2508_v2 = new _dafny.CodePoint('q'.codePointAt(0));
      let _2509_v3;
      _2509_v3 = _dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), _2508_v2);
      let _2510_v4;
      _2510_v4 = _dafny.Seq.of(_2509_v3, _2509_v3, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-249)), function (_2511_i0) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(398)), ((_2512_v2) => function (_2513_i1) {
        return _2512_v2;
      })(_2508_v2)), _2509_v3), _dafny.Seq.of(_2508_v2));
      _2510_v4 = _2510_v4;
      let _2514_v5;
      _2514_v5 = _dafny.MultiSet.fromElements(_2508_v2);
      let _2515_v6;
      let _nw394 = Array((_dafny.ONE).toNumber());
      _nw394[(_dafny.ZERO).toNumber()] = _2514_v5;
      _2515_v6 = _nw394;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_2515_v6).length))) {
        let _2516_i2 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_2516_i2)) && ((_2516_i2).isLessThan(new BigNumber((_2515_v6).length))))) {
          (_2515_v6)[(_2516_i2)] = _2514_v5;
        }
      }
      let _2517_v7;
      _2517_v7 = _dafny.Set.fromElements(_2506_v0);
      let _2518_i3;
      _2518_i3 = _dafny.ZERO;
      L12: {
        while ((_2517_v7).IsProperSubsetOf(_2517_v7)) {
          C12: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_2518_i3)) {
              break L12;
            }
            _2518_i3 = (_2518_i3).plus(_dafny.ONE);
            let _2519_v8;
            _2519_v8 = _dafny.Map.Empty.slice().updateUnsafe((_2517_v7).IsDisjointFrom(_2517_v7),p0);
            let _2520_v9;
            _2520_v9 = _dafny.Seq.of(true);
            let _2521_v10;
            let _nw395 = Array((new BigNumber(13)).toNumber());
            _nw395[(_dafny.ZERO).toNumber()] = _2506_v0;
            _nw395[(_dafny.ONE).toNumber()] = _2506_v0;
            _nw395[(new BigNumber(2)).toNumber()] = false;
            _nw395[(new BigNumber(3)).toNumber()] = _2506_v0;
            _nw395[(new BigNumber(4)).toNumber()] = _2506_v0;
            _nw395[(new BigNumber(5)).toNumber()] = !(!(_2506_v0));
            _nw395[(new BigNumber(6)).toNumber()] = _2506_v0;
            _nw395[(new BigNumber(7)).toNumber()] = _2506_v0;
            _nw395[(new BigNumber(8)).toNumber()] = _2506_v0;
            _nw395[(new BigNumber(9)).toNumber()] = _2506_v0;
            _nw395[(new BigNumber(10)).toNumber()] = (_2520_v9)[_module.__default.safeIndex(new BigNumber((_2509_v3).length), new BigNumber((_2520_v9).length))];
            _nw395[(new BigNumber(11)).toNumber()] = _2506_v0;
            _nw395[(new BigNumber(12)).toNumber()] = _2506_v0;
            _2521_v10 = _nw395;
            let _2522_v11;
            _2522_v11 = _dafny.Set.fromElements(_2521_v10);
            _2519_v8 = (_2519_v8).update(_2506_v0, (_dafny.ZERO).minus(new BigNumber(((_2522_v11).Difference(_2522_v11)).length)));
            r1 = (_dafny.ZERO).minus(p0);
            let _2523_v12;
            _2523_v12 = _dafny.Seq.of((p0).plus(p0));
            _2523_v12 = _dafny.Seq.update(_dafny.Seq.of(p0, new BigNumber((_2509_v3).length), p0), _module.__default.safeIndex((new BigNumber((_module.__default.fm1(p0, _2506_v0, p0, _2506_v0, globalState)).length)).minus((_this).fm0(false, p0, globalState)), new BigNumber((_dafny.Seq.of(p0, new BigNumber((_2509_v3).length), p0)).length)), p0);
            _2506_v0 = _2506_v0;
          }
        }
      }
      let _2524_v13;
      _2524_v13 = _dafny.Map.Empty.slice().updateUnsafe(p0,_2506_v0);
      _2524_v13 = _2524_v13;
      let _2525_v14;
      let _init62 = ((_2526_v7) => function (_2527_i4) {
        return _2526_v7;
      })(_2517_v7);
      let _nw396 = Array((new BigNumber(3)).toNumber());
      for (let _i0_62 = 0; _i0_62 < new BigNumber(_nw396.length); _i0_62++) {
        _nw396[_i0_62] = _init62(new BigNumber(_i0_62));
      }
      _2525_v14 = _nw396;
      let _2528_v15;
      _2528_v15 = _module.D0.create_DC0(((_2506_v0) ? (_2525_v14) : (_2525_v14)), p0);
      let _source26 = _2528_v15;
      if (_source26.is_DC0) {
        let _2529___mcc_h0 = (_source26).cf0;
        let _2530___mcc_h1 = (_source26).cf1;
        let _2531_cf1 = _2530___mcc_h1;
        let _2532_cf0 = _2529___mcc_h0;
        r0 = (_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("moityagb")).length), p0));
        let _2533_v16;
        _2533_v16 = _dafny.Seq.of(_2506_v0, _2506_v0, true, _2506_v0);
        let _2534_v17;
        let _nw397 = Array((new BigNumber(9)).toNumber());
        _nw397[(_dafny.ZERO).toNumber()] = !(_2506_v0) || (!(!(_2506_v0)));
        _nw397[(_dafny.ONE).toNumber()] = (_2533_v16)[_module.__default.safeIndex((_this).fm0(_2506_v0, p0, globalState), new BigNumber((_2533_v16).length))];
        _nw397[(new BigNumber(2)).toNumber()] = (_2533_v16)[_module.__default.safeIndex(_2531_cf1, new BigNumber((_2533_v16).length))];
        _nw397[(new BigNumber(3)).toNumber()] = !_dafny.areEqual(_2533_v16, _2533_v16);
        _nw397[(new BigNumber(4)).toNumber()] = ((_2506_v0) ? (_2506_v0) : (_2506_v0));
        _nw397[(new BigNumber(5)).toNumber()] = ((_dafny.ZERO).minus(_2531_cf1)).isLessThanOrEqualTo(p0);
        _nw397[(new BigNumber(6)).toNumber()] = ((_2506_v0) ? (!(_2506_v0)) : (_2506_v0));
        _nw397[(new BigNumber(7)).toNumber()] = _2506_v0;
        _nw397[(new BigNumber(8)).toNumber()] = !(_2506_v0);
        _2534_v17 = _nw397;
        let _index405 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2534_v17).length));
        (_2534_v17)[_index405] = _2506_v0;
        let _2535_v18;
        _2535_v18 = _dafny.Map.Empty.slice().updateUnsafe((_2506_v0) && (_2506_v0),_2506_v0);
        let _index406 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2534_v17).length));
        (_2534_v17)[_index406] = !((((_2535_v18).contains(!(_2506_v0))) ? ((_2535_v18).get(!(_2506_v0))) : (_2506_v0)));
        let _2536_v19;
        _2536_v19 = _dafny.Seq.of(_2517_v7, _dafny.Set.fromElements(_2506_v0), _2517_v7, _2517_v7, _dafny.Set.fromElements((_2534_v17)[_module.__default.safeIndex(new BigNumber(505), new BigNumber((_2534_v17).length))]));
        let _2537_v20;
        _2537_v20 = _dafny.Set.fromElements(p0, _2531_cf1, new BigNumber((_2536_v19).length));
        let _index407 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2534_v17).length));
        let _rhs309 = (new BigNumber(568)).plus(p0);
        let _rhs310 = _2509_v3;
        let _rhs311 = !(_2506_v0);
        let _rhs312 = (_2537_v20).IsProperSubsetOf(_2537_v20);
        let _lhs178 = _2534_v17;
        let _lhs179 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2534_v17).length));
        r0 = _rhs309;
        _2509_v3 = _rhs310;
        _lhs178[_lhs179] = _rhs311;
        _2506_v0 = _rhs312;
        let _2538_v21;
        let _nw398 = Array((new BigNumber(10)).toNumber()).fill([]);
        _2538_v21 = _nw398;
        let _2539_v22;
        let _nw399 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _2539_v22 = _nw399;
        let _index408 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_2538_v21).length));
        (_2538_v21)[_index408] = _2539_v22;
        let _index409 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_2538_v21).length));
        let _index410 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2534_v17).length));
        let _rhs313 = _2539_v22;
        let _rhs314 = (_2533_v16)[_module.__default.safeIndex(_2531_cf1, new BigNumber((_2533_v16).length))];
        let _rhs315 = _dafny.Seq.Concat(_2510_v4, _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(363)), ((_2540_v2) => function (_2541_i5) {
          return _2540_v2;
        })(_2508_v2)), _2509_v3));
        let _lhs180 = _2538_v21;
        let _lhs181 = _module.__default.safeIndex(new BigNumber(401), new BigNumber((_2538_v21).length));
        let _lhs182 = _2534_v17;
        let _lhs183 = _module.__default.safeIndex(new BigNumber(505), new BigNumber((_2534_v17).length));
        _lhs180[_lhs181] = _rhs313;
        _lhs182[_lhs183] = _rhs314;
        _2510_v4 = _rhs315;
      } else {
        let _2542___mcc_h2 = (_source26).cf2;
        let _2543_cf2 = _2542___mcc_h2;
        if (!(!(((_2506_v0) ? (_module.__default.fm2((_2506_v0), _2543_cf2, globalState)) : (_2506_v0))))) {
          _2543_cf2 = ((_dafny.ZERO).minus(_2543_cf2)).plus(p0);
          _2506_v0 = _2506_v0;
          let _2544_v23;
          _2544_v23 = _dafny.Seq.of(_2506_v0);
          _2506_v0 = _module.__default.fm2((_2544_v23)[_module.__default.safeIndex(p0, new BigNumber((_2544_v23).length))], p0, globalState);
          _2506_v0 = _module.__default.fm2(_2506_v0, _module.__default.safeModuloInt(p0, new BigNumber(829)), globalState);
          let _2545_v24;
          _2545_v24 = _dafny.Map.Empty.slice().updateUnsafe(_2544_v23,_2508_v2);
          _2545_v24 = _module.__default.fm3(_2508_v2, (_2543_cf2).multipliedBy((_dafny.ZERO).minus(_2543_cf2)), _2506_v0, _2506_v0, globalState);
        } else {
          let _2546_v25;
          let _nw400 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
          _2546_v25 = _nw400;
          let _2547_v26;
          _2547_v26 = _dafny.Map.Empty.slice().updateUnsafe(_2506_v0,_2543_cf2);
          let _index411 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_2546_v25).length));
          (_2546_v25)[_index411] = (((_2547_v26).contains(true)) ? ((_2547_v26).get(true)) : (new BigNumber((_2509_v3).length)));
          let _index412 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_2546_v25).length));
          (_2546_v25)[_index412] = (_this).fm0((_2506_v0) && (_2506_v0), (new BigNumber((_module.__default.fm4(_2509_v3, _2508_v2, _2506_v0, globalState)).length)).minus(p0), globalState);
          _2509_v3 = _dafny.Seq.Concat(_2509_v3, _dafny.Seq.UnicodeFromString("rbslwhc"));
          let _2548_v27;
          _2548_v27 = _dafny.Map.Empty.slice().updateUnsafe(_2506_v0,_2508_v2);
          let _2549_v28;
          _2549_v28 = false;
          _2548_v27 = (_2548_v27).update(_2549_v28, new _dafny.CodePoint('o'.codePointAt(0)));
          let _2550_v29;
          let _nw401 = Array((new BigNumber(5)).toNumber());
          _nw401[(_dafny.ZERO).toNumber()] = !(_2506_v0) || (_2506_v0);
          _nw401[(_dafny.ONE).toNumber()] = true;
          _nw401[(new BigNumber(2)).toNumber()] = _2506_v0;
          _nw401[(new BigNumber(3)).toNumber()] = _2506_v0;
          _nw401[(new BigNumber(4)).toNumber()] = true;
          _2550_v29 = _nw401;
          let _index413 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_2550_v29).length));
          (_2550_v29)[_index413] = false;
          let _2551_v30;
          _2551_v30 = _module.D0.create_DC1((_dafny.ZERO).minus((_2546_v25)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_2546_v25).length))]));
          let _2552_v31;
          _2552_v31 = _dafny.Map.Empty.slice().updateUnsafe(_2551_v30,_2517_v7);
          let _index414 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2550_v29).length));
          (_2550_v29)[_index414] = _2506_v0;
          let _2553_v32;
          _2553_v32 = _dafny.Map.Empty.slice().updateUnsafe(false,_2506_v0);
          let _index415 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_2550_v29).length));
          let _index416 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2550_v29).length));
          let _rhs316 = _module.__default.safeModuloInt((_2546_v25)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_2546_v25).length))], p0);
          let _rhs317 = !(_2506_v0);
          let _rhs318 = _module.__default.safeModuloInt(_2543_cf2, (_2546_v25)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_2546_v25).length))]);
          let _rhs319 = (((_2553_v32).contains(_2506_v0)) ? (_2552_v31) : ((_2552_v31).update(_2551_v30, _2517_v7)));
          let _rhs320 = true;
          let _lhs184 = _2550_v29;
          let _lhs185 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_2550_v29).length));
          let _lhs186 = _2550_v29;
          let _lhs187 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_2550_v29).length));
          r0 = _rhs316;
          _lhs184[_lhs185] = _rhs317;
          r0 = _rhs318;
          _2552_v31 = _rhs319;
          _lhs186[_lhs187] = _rhs320;
          let _index417 = _module.__default.safeIndex(new BigNumber(550), new BigNumber((_2550_v29).length));
          (_2550_v29)[_index417] = ((_2546_v25)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_2546_v25).length))]).isLessThanOrEqualTo(new BigNumber(470));
        }
        let _2554_v33;
        let _nw402 = new _module.C3();
        _nw402.__ctor();
        _2554_v33 = _nw402;
        _2543_cf2 = (p0).multipliedBy(p0);
        _2506_v0 = !((_2506_v0) === (!((_dafny.Set.fromElements(true)).IsSubsetOf(_2517_v7))));
      }
      r0 = p0;
      let _2555_v34;
      let _nw403 = new _module.C3();
      _nw403.__ctor();
      _2555_v34 = _nw403;
      let _2556_v35;
      _2556_v35 = _dafny.Seq.of(_2506_v0, _2506_v0, _2506_v0, _2506_v0);
      let _2557_v36;
      _2557_v36 = _dafny.Map.Empty.slice().updateUnsafe(_2555_v34,_2556_v35);
      r1 = new BigNumber(((((_2557_v36).contains(_2555_v34)) ? ((_2557_v36).get(_2555_v34)) : (_2556_v35))).length);
      r2 = _dafny.Seq.Concat(_2556_v35, _2556_v35);
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
