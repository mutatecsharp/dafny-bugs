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
    static fm0(globalState) {
      return new BigNumber(-214);
    };
    static fm1(p0, p1, globalState) {
      return _dafny.Seq.of(!(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(420), new BigNumber((_dafny.Seq.of(false, (_module.D0.create_DC2(_dafny.MultiSet.fromElements(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('h'.codePointAt(0))), !(!(true)))).dtor_cf6)).length), new BigNumber(630), new BigNumber(-203), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false, true)).cardinality())))).cardinality()))).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-421))));
    };
    static fm2(p0, globalState) {
      return ((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(-551), new BigNumber((_dafny.Seq.UnicodeFromString("apvmft")).length))))).isLessThanOrEqualTo(new BigNumber(-770));
    };
    static fm4(p0, p1, p2, p3, globalState) {
      if (((true) ? (false) : (false))) {
        return _dafny.Seq.of(new BigNumber(811));
      } else {
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(552)), function (_0_i0) {
          return (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(927),(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true, false)).cardinality())))).length));
        }), _dafny.Seq.of(new BigNumber((_dafny.Seq.of(new BigNumber(-559), new BigNumber(-695), new BigNumber(903))).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(false, false)).length))));
      }
    };
    static fm6(p0, p1, globalState) {
      let _source0 = _module.D10.create_DC29(false, new _dafny.CodePoint('f'.codePointAt(0)));
      if (_source0.is_DC29) {
        let _1___mcc_h0 = (_source0).cf42;
        let _2___mcc_h1 = (_source0).cf43;
        let _3_cf43 = _2___mcc_h1;
        let _4_cf42 = _1___mcc_h0;
        return _dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-730)), _dafny.Seq.of(new BigNumber(14), new BigNumber(18))));
      } else if (_source0.is_DC30) {
        let _5___mcc_h2 = (_source0).cf44;
        let _6___mcc_h3 = (_source0).cf45;
        let _7___mcc_h4 = (_source0).cf46;
        let _8___mcc_h5 = (_source0).cf47;
        let _9___mcc_h6 = (_source0).cf48;
        let _10_cf48 = _9___mcc_h6;
        let _11_cf47 = _8___mcc_h5;
        let _12_cf46 = _7___mcc_h4;
        let _13_cf45 = _6___mcc_h3;
        let _14_cf44 = _5___mcc_h2;
        if (!(_12_cf46)) {
          return _dafny.MultiSet.fromElements(_10_cf48, new BigNumber(212), _11_cf47);
        } else {
          return _dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_12_cf46, _12_cf46)).length), new BigNumber(848)));
        }
      } else {
        let _15___mcc_h7 = (_source0).cf41;
        let _16_cf41 = _15___mcc_h7;
        return _dafny.MultiSet.fromElements(new BigNumber(723));
      }
    };
    static fm7(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(false)).Difference(_dafny.Set.fromElements(!(false)))).Intersect((_dafny.Set.fromElements(true, true, true, true, false)).Union(_dafny.Set.fromElements(true)));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      return _dafny.Seq.Create(_module.__default.abs(new BigNumber(468)), function (_17_i0) {
        return ((true) ? (new _dafny.CodePoint('k'.codePointAt(0))) : (new _dafny.CodePoint('a'.codePointAt(0))));
      });
    };
    static fm9(p0, p1, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(true),true));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      let _source1 = _module.D13.create_DC38(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
      if (_source1.is_DC39) {
        let _18___mcc_h0 = (_source1).cf69;
        let _19_cf69 = _18___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(-396)), _dafny.Seq.of(new BigNumber(367), new BigNumber(184), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(838),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-492)), function (_20_i0) {
          return new BigNumber(-647);
        })).length)))).length)));
      } else if (_source1.is_DC40) {
        let _21___mcc_h1 = (_source1).cf70;
        let _22___mcc_h2 = (_source1).cf71;
        let _23_cf71 = _22___mcc_h2;
        let _24_cf70 = _21___mcc_h1;
        return _dafny.Seq.of(_23_cf71);
      } else if (_source1.is_DC38) {
        let _25___mcc_h3 = (_source1).cf68;
        let _26_cf68 = _25___mcc_h3;
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(557)), function (_27_i1) {
          return new BigNumber((_dafny.Set.fromElements(true)).length);
        });
      } else {
        let _28___mcc_h4 = (_source1).cf72;
        let _29_cf72 = _28___mcc_h4;
        return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-849)), function (_30_i2) {
          return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),new BigNumber(603))).length),new BigNumber(497))).length);
        }), _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('f'.codePointAt(0)),false), _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('u'.codePointAt(0)),true))).length)), new BigNumber(212)));
      }
    };
    static fm11(p0, globalState) {
      if ((new BigNumber(331)).isEqualTo(new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ao")).length)))).cardinality()))) {
        return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(!(false)),!(false))).length),(_dafny.ZERO).minus(new BigNumber(-592))));
      } else {
        return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(104),new BigNumber((_dafny.Seq.of(new BigNumber(-845), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),false)).length), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("rdjmrl")).length))).cardinality()))).length))).length),new BigNumber(371));
      }
    };
    static fm14(p0, p1, globalState) {
      return ((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).Intersect(_dafny.MultiSet.fromElements(true, !(true)))).Difference(_dafny.MultiSet.fromElements(false));
    };
    static fm15(p0, p1, p2, globalState) {
      return (((!(false)) ? (_module.D16.create_DC46(_dafny.Seq.UnicodeFromString("wrs"), _dafny.Set.fromElements(false), false)) : (_module.D16.create_DC46(_dafny.Seq.UnicodeFromString("kbcxn"), _dafny.Set.fromElements(false), true)))).dtor_cf79;
    };
    static fm16(p0, p1, globalState) {
      return (((false) ? (_dafny.MultiSet.fromElements(_dafny.Seq.of(false, false, !(!(false))), _dafny.Seq.of(true))) : (_dafny.MultiSet.fromElements(_dafny.Seq.of(true), _dafny.Seq.of(true))))).Difference((_dafny.MultiSet.fromElements((_module.D6.create_DC17(_dafny.Seq.of(true, true, false, !(false)))).dtor_cf23, _dafny.Seq.of(false, false), _dafny.Seq.of(false, true))).Intersect(_dafny.MultiSet.fromElements(_dafny.Seq.of(false, false, false, false))));
    };
    static fm17(p0, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('r'.codePointAt(0)),true);
    };
    static fm18(p0, p1, p2, globalState) {
      return (_dafny.Set.fromElements((_module.D8.create_DC23(true, !(false), false, new _dafny.CodePoint('x'.codePointAt(0)))).dtor_cf35)).Difference(_dafny.Set.fromElements(new _dafny.CodePoint('s'.codePointAt(0))));
    };
    static fm19(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber(513), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality())), new BigNumber(420), new BigNumber(698)), _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(true),new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Seq.UnicodeFromString("qa")).length)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber(-314)))).length)))).length))).length))), _dafny.Seq.of(new BigNumber(48), new BigNumber(460))));
    };
    static fm20(p0, globalState) {
      return _dafny.Seq.UnicodeFromString("lt");
    };
    static fm21(p0, globalState) {
      if (true) {
        return _dafny.Set.fromElements(_module.D0.create_DC3(), _module.D0.create_DC3(), _module.D0.create_DC3());
      } else {
        return _dafny.Set.fromElements(_module.D0.create_DC3());
      }
    };
    static fm22(p0, p1, globalState) {
      return ((_dafny.Set.fromElements(true, true)).Difference(_dafny.Set.fromElements(!(true)))).Union((_dafny.Set.fromElements(!(true))).Difference(_dafny.Set.fromElements(true, false, true)));
    };
    static fm23(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(210),new BigNumber(454))).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(417)), function (_31_i0) {
        return _module.D8.create_DC23(!(false), true, false, new _dafny.CodePoint('i'.codePointAt(0)));
      })).length),new BigNumber(794)))).Merge((function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(134), new BigNumber(361))) {
          let _32_v0 = _compr_0;
          if (((new BigNumber(134)).isLessThanOrEqualTo(_32_v0)) && ((_32_v0).isLessThan(new BigNumber(361)))) {
            _coll0.push([(_32_v0).minus(new BigNumber((_dafny.Seq.of(false)).length)),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("vu")).length))]);
          }
        }
        return _coll0;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(666),new BigNumber(-331))));
    };
    static fm24(globalState) {
      return new _dafny.CodePoint('h'.codePointAt(0));
    };
    static fm25(p0, globalState) {
      return _module.D0.create_DC0((function () {
  let _coll1 = new _dafny.Set();
  for (const _compr_1 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(249)), function (_33_i0) {
    return _dafny.Seq.of(new BigNumber(428));
  })).Elements) {
    let _34_v0 = _compr_1;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(249)), function (_33_i0) {
      return _dafny.Seq.of(new BigNumber(428));
    }), _34_v0)) {
      _coll1.add(_34_v0);
    }
  }
  return _coll1;
}()).IsDisjointFrom(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-168)), function (_35_i1) {
  return new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),false)).length);
})).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false, (_module.D16.create_DC46(_dafny.Seq.Create(_module.__default.abs(new BigNumber(210)), function (_36_i2) {
  return new _dafny.CodePoint('v'.codePointAt(0));
}), _dafny.Set.fromElements(false), false)).dtor_cf81)).length))))));
    };
    static fm26(globalState) {
      return _module.D8.create_DC24(((true) ? (_module.D8.create_DC23(false, !(false), !(true), new _dafny.CodePoint('f'.codePointAt(0)))) : (_module.D8.create_DC23(true, true, true, new _dafny.CodePoint('t'.codePointAt(0))))));
    };
    static fm27(globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(186),false)).Merge((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()),false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("ogslplvme")).length),true)));
    };
    static fm28(p0, p1, p2, globalState) {
      if ((new BigNumber((function () {
        let _coll2 = new _dafny.Map();
        for (const _compr_2 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false, false, false),true)).Keys.Elements) {
          let _37_v0 = _compr_2;
          if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(false, false, false),true)).contains(_37_v0)) {
            _coll2.push([_37_v0,new BigNumber((_dafny.Seq.of(true, true, false)).length)]);
          }
        }
        return _coll2;
      }()).length)).isLessThan((_module.D11.create_DC33(false, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,false)).length), new BigNumber(-513))).dtor_cf54)) {
        return (function () {
          let _coll3 = new _dafny.Map();
          for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, true)).length))),new _dafny.CodePoint('m'.codePointAt(0)))).Keys.Elements) {
            let _38_v1 = _compr_3;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, true)).length))),new _dafny.CodePoint('m'.codePointAt(0)))).contains(_38_v1)) {
              _coll3.push([_38_v1,_dafny.Seq.UnicodeFromString("mkjgpgh")]);
            }
          }
          return _coll3;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber(-199))),_dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), function (_39_i0) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        })));
      } else {
        return (function () {
          let _coll4 = new _dafny.Map();
          for (const _compr_4 of (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(-778), new BigNumber(723)))).Elements) {
            let _40_v2 = _compr_4;
            if ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(-778), new BigNumber(723)))).contains(_40_v2)) {
              _coll4.push([_40_v2,_dafny.Seq.UnicodeFromString("ivuqtjs")]);
            }
          }
          return _coll4;
        }()).Merge(function () {
          let _coll5 = new _dafny.Map();
          for (const _compr_5 of (_dafny.Seq.of(_dafny.Seq.of(new BigNumber(196)))).Elements) {
            let _41_v3 = _compr_5;
            if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(196))), _41_v3)) {
              _coll5.push([_41_v3,_dafny.Seq.UnicodeFromString("qpbh")]);
            }
          }
          return _coll5;
        }());
      }
    };
    static fm29(globalState) {
      if ((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('g'.codePointAt(0)))).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('i'.codePointAt(0))))) {
        return _module.D6.create_DC18(true, new BigNumber((_dafny.Seq.UnicodeFromString("a")).length));
      } else {
        return _module.D6.create_DC18(!(true), new BigNumber(760));
      }
    };
    static fm30(p0, p1, p2, globalState) {
      return _module.D12.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(false,false), !(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements((_module.D12.create_DC36(false, false, new BigNumber((_dafny.Seq.of(new BigNumber(486))).length))).dtor_cf61, !(true))).cardinality()))).length),false)).length),new BigNumber(15))).length)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D0.create_DC3(),new BigNumber((_dafny.Set.fromElements(true, !(false))).length))).length)), new BigNumber(713), new BigNumber(217));
    };
    static fm31(globalState) {
      let _source2 = _module.D12.create_DC36(false, false, (_module.D12.create_DC37(_dafny.Map.Empty.slice().updateUnsafe(!(false),true), true, new BigNumber(527), new BigNumber(89))).dtor_cf67);
      if (_source2.is_DC35) {
        let _42___mcc_h0 = (_source2).cf56;
        let _43___mcc_h1 = (_source2).cf57;
        let _44___mcc_h2 = (_source2).cf58;
        let _45___mcc_h3 = (_source2).cf59;
        let _46___mcc_h4 = (_source2).cf60;
        let _47_cf60 = _46___mcc_h4;
        let _48_cf59 = _45___mcc_h3;
        let _49_cf58 = _44___mcc_h2;
        let _50_cf57 = _43___mcc_h1;
        let _51_cf56 = _42___mcc_h0;
        return _dafny.Set.fromElements(new BigNumber(401));
      } else if (_source2.is_DC36) {
        let _52___mcc_h5 = (_source2).cf61;
        let _53___mcc_h6 = (_source2).cf62;
        let _54___mcc_h7 = (_source2).cf63;
        let _55_cf63 = _54___mcc_h7;
        let _56_cf62 = _53___mcc_h6;
        let _57_cf61 = _52___mcc_h5;
        return _dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(_56_cf62, _57_cf61)).length));
      } else if (_source2.is_DC37) {
        let _58___mcc_h8 = (_source2).cf64;
        let _59___mcc_h9 = (_source2).cf65;
        let _60___mcc_h10 = (_source2).cf66;
        let _61___mcc_h11 = (_source2).cf67;
        let _62_cf67 = _61___mcc_h11;
        let _63_cf66 = _60___mcc_h10;
        let _64_cf65 = _59___mcc_h9;
        let _65_cf64 = _58___mcc_h8;
        return _dafny.Set.fromElements(_62_cf67);
      } else {
        let _66___mcc_h12 = (_source2).cf55;
        let _67_cf55 = _66___mcc_h12;
        return (_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, (_module.D16.create_DC46(_dafny.Seq.UnicodeFromString("yoa"), _dafny.Set.fromElements(false, true), false)).dtor_cf81)).length)), new BigNumber(725))).Union((_module.D4.create_DC11(function () {
  let _coll6 = new _dafny.Set();
  for (const _compr_6 of _dafny.IntegerRange(new BigNumber(784), new BigNumber(-950))) {
    let _68_v0 = _compr_6;
    if (((new BigNumber(784)).isLessThanOrEqualTo(_68_v0)) && ((_68_v0).isLessThan(new BigNumber(-950)))) {
      _coll6.add((_68_v0).minus(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('o'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)), new _dafny.CodePoint('d'.codePointAt(0)), new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('t'.codePointAt(0)))).cardinality())));
    }
  }
  return _coll6;
}())).dtor_cf18);
      }
    };
    static Main(__noArgsParameter) {
      let _69_v0;
      let _init0 = function (_70_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      };
      let _nw0 = Array((new BigNumber(26)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw0.length); _i0_0++) {
        _nw0[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _69_v0 = _nw0;
      let _71_globalState;
      let _nw1 = new _module.GlobalState();
      _nw1.__ctor(_69_v0, new BigNumber(-562), false);
      _71_globalState = _nw1;
      let _72_v1;
      _72_v1 = false;
      if (_72_v1) {
        let _73_v2;
        let _init1 = ((_74_v1) => function (_75_i1) {
          return _74_v1;
        })(_72_v1);
        let _nw2 = Array((new BigNumber(13)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw2.length); _i0_1++) {
          _nw2[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _73_v2 = _nw2;
        let _76_v3;
        _76_v3 = new BigNumber(-419);
        let _77_v4;
        _77_v4 = _dafny.Map.Empty.slice().updateUnsafe(_76_v3,_73_v2);
        _73_v2 = (((_77_v4).contains(_76_v3)) ? ((_77_v4).get(_76_v3)) : (_73_v2));
        _76_v3 = new BigNumber(587);
        (_71_globalState).f2 = _72_v1;
        let _78_v5;
        _78_v5 = new _dafny.CodePoint('v'.codePointAt(0));
        _78_v5 = _78_v5;
        let _79_v6;
        _79_v6 = _dafny.MultiSet.fromElements(_78_v5, _78_v5);
        let _80_v7;
        _80_v7 = _dafny.Seq.of(new BigNumber(43), new BigNumber(206), ((_72_v1) ? (_76_v3) : (new BigNumber(862))), new BigNumber(((_79_v6).update(_78_v5, _module.__default.abs(_76_v3))).cardinality()), _76_v3);
        let _81_v8;
        _81_v8 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(349),_72_v1);
        let _82_v9;
        _82_v9 = _module.D0.create_DC2(_79_v6, _72_v1);
        let _83_v10;
        _83_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_81_v8).length),(_82_v9).dtor_cf6);
        let _84_v11;
        _84_v11 = _dafny.Map.Empty.slice().updateUnsafe(_72_v1,_81_v8);
        _80_v7 = _dafny.Seq.Concat(_dafny.Seq.Concat(_80_v7, _80_v7), _dafny.Seq.of(_76_v3, new BigNumber((_84_v11).length), _76_v3, new BigNumber(140), new BigNumber(-53)));
      } else {
        let _85_v12;
        let _nw3 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
        _85_v12 = _nw3;
        let _86_v13;
        _86_v13 = new BigNumber(744);
        let _index0 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_85_v12).length));
        (_85_v12)[_index0] = _86_v13;
        let _index1 = _module.__default.safeIndex(new BigNumber(164), new BigNumber((_85_v12).length));
        (_85_v12)[_index1] = _module.__default.fm0(_71_globalState);
        let _87_v14;
        _87_v14 = _dafny.Map.Empty.slice().updateUnsafe(_86_v13,_module.__default.fm0(_71_globalState));
        let _88_v15;
        _88_v15 = _dafny.Seq.UnicodeFromString("t");
        let _89_v16;
        _89_v16 = _dafny.Map.Empty.slice().updateUnsafe(_86_v13,_88_v15);
        let _90_v17;
        _90_v17 = _dafny.Seq.of(_86_v13);
        let _91_v18;
        _91_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_85_v12)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_85_v12).length))], (_85_v12)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_85_v12).length))]),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_89_v16).length),(_90_v17)[_module.__default.safeIndex((_85_v12)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_85_v12).length))], new BigNumber((_90_v17).length))]));
        _87_v14 = (((_91_v18).contains(_86_v13)) ? ((_91_v18).get(_86_v13)) : ((_87_v14).update(new BigNumber(-505), _86_v13)));
        let _92_v19;
        _92_v19 = _dafny.Seq.of(_85_v12, _85_v12, _85_v12);
        let _93_v20;
        _93_v20 = _dafny.Seq.of(!(_72_v1));
        let _94_v21;
        let _nw4 = Array((new BigNumber(16)).toNumber());
        _nw4[(_dafny.ZERO).toNumber()] = _85_v12;
        _nw4[(_dafny.ONE).toNumber()] = _85_v12;
        _nw4[(new BigNumber(2)).toNumber()] = ((_72_v1) ? (_85_v12) : ((_92_v19)[_module.__default.safeIndex(_86_v13, new BigNumber((_92_v19).length))]));
        _nw4[(new BigNumber(3)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(4)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(5)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(6)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(7)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(8)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(9)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(10)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(11)).toNumber()] = (_92_v19)[_module.__default.safeIndex(new BigNumber((_93_v20).length), new BigNumber((_92_v19).length))];
        _nw4[(new BigNumber(12)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(13)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(14)).toNumber()] = _85_v12;
        _nw4[(new BigNumber(15)).toNumber()] = _85_v12;
        _94_v21 = _nw4;
        let _index2 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_94_v21).length));
        (_94_v21)[_index2] = _85_v12;
        let _95_v22;
        _95_v22 = _dafny.Map.Empty.slice().updateUnsafe((_85_v12)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_85_v12).length))],false);
        let _96_v23;
        _96_v23 = new _dafny.CodePoint('m'.codePointAt(0));
        let _97_v24;
        _97_v24 = _dafny.Seq.of(_dafny.Seq.update(_88_v15, _module.__default.safeIndex(new BigNumber((_95_v22).length), new BigNumber((_88_v15).length)), _96_v23));
        let _98_v25;
        _98_v25 = _dafny.Set.fromElements(_module.__default.fm0(_71_globalState), (_85_v12)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_85_v12).length))], _86_v13);
        let _99_v26;
        _99_v26 = _dafny.MultiSet.fromElements(_72_v1);
        let _100_v27;
        _100_v27 = _dafny.Map.Empty.slice().updateUnsafe((_99_v26).IsDisjointFrom(_dafny.MultiSet.FromArray(_module.__default.fm1(_96_v23, _72_v1, _71_globalState))),!(_module.__default.fm2(_86_v13, _71_globalState)));
        let _101_v28;
        let _init2 = ((_102_v1) => function (_103_i2) {
          return _102_v1;
        })(_72_v1);
        let _nw5 = Array((new BigNumber(21)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw5.length); _i0_2++) {
          _nw5[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _101_v28 = _nw5;
        let _index3 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_94_v21).length));
        let _rhs0 = (_98_v25).contains((_dafny.ZERO).minus((_85_v12)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_85_v12).length))]));
        let _rhs1 = _85_v12;
        let _rhs2 = _dafny.Seq.Concat(_97_v24, _dafny.Seq.Concat(_97_v24, _97_v24));
        let _rhs3 = (((_100_v27).contains(!(!(_module.__default.fm0(_71_globalState)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_99_v26).cardinality()),_101_v28)).length))))) ? ((_100_v27).get(!(!(_module.__default.fm0(_71_globalState)).isEqualTo(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_99_v26).cardinality()),_101_v28)).length))))) : (_72_v1));
        let _rhs4 = _88_v15;
        let _lhs0 = _94_v21;
        let _lhs1 = _module.__default.safeIndex(new BigNumber(223), new BigNumber((_94_v21).length));
        _72_v1 = _rhs0;
        _lhs0[_lhs1] = _rhs1;
        _97_v24 = _rhs2;
        _72_v1 = _rhs3;
        _88_v15 = _rhs4;
        let _104_v29;
        let _nw6 = new _module.C1();
        _nw6.__ctor(_86_v13, _module.__default.safeDivisionInt((_85_v12)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_85_v12).length))], _86_v13), _module.__default.fm2(new BigNumber(132), _71_globalState));
        _104_v29 = _nw6;
        let _105_v30;
        let _out0;
        _out0 = (_104_v29).m1((_85_v12)[_module.__default.safeIndex(new BigNumber(164), new BigNumber((_85_v12).length))], _71_globalState);
        _105_v30 = _out0;
      }
      let _106_v31;
      let _nw7 = Array((new BigNumber(11)).toNumber());
      _nw7[(_dafny.ZERO).toNumber()] = _72_v1;
      _nw7[(_dafny.ONE).toNumber()] = _72_v1;
      _nw7[(new BigNumber(2)).toNumber()] = _72_v1;
      _nw7[(new BigNumber(3)).toNumber()] = _72_v1;
      _nw7[(new BigNumber(4)).toNumber()] = _72_v1;
      _nw7[(new BigNumber(5)).toNumber()] = _72_v1;
      _nw7[(new BigNumber(6)).toNumber()] = _72_v1;
      _nw7[(new BigNumber(7)).toNumber()] = _72_v1;
      _nw7[(new BigNumber(8)).toNumber()] = !(_72_v1);
      _nw7[(new BigNumber(9)).toNumber()] = _72_v1;
      _nw7[(new BigNumber(10)).toNumber()] = _72_v1;
      _106_v31 = _nw7;
      let _107_v32;
      _107_v32 = new BigNumber(714);
      let _108_v33;
      _108_v33 = _dafny.Set.fromElements(_72_v1);
      let _109_v34;
      let _nw8 = new _module.C5();
      _nw8.__ctor(_106_v31, ((_dafny.ZERO).minus(_107_v32)).minus(new BigNumber((_108_v33).length)), (_107_v32).minus(_module.__default.fm0(_71_globalState)), (_108_v33).IsProperSubsetOf(_108_v33));
      _109_v34 = _nw8;
      let _110_v35;
      _110_v35 = new _dafny.CodePoint('s'.codePointAt(0));
      let _111_v36;
      _111_v36 = _dafny.Seq.UnicodeFromString("ndf");
      let _112_v37;
      let _nw9 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
      _112_v37 = _nw9;
      let _113_v38;
      let _nw10 = new _module.C3();
      _nw10.__ctor(_110_v35, _111_v36, _107_v32, _72_v1, (_module.D10.create_DC30(new BigNumber((_module.__default.fm1(_110_v35, _72_v1, _71_globalState)).length), _112_v37, _72_v1, _107_v32, _107_v32)).dtor_cf44);
      _113_v38 = _nw10;
      let _114_v39;
      _114_v39 = _dafny.Map.Empty.slice().updateUnsafe(_113_v38,_107_v32);
      let _115_v40;
      _115_v40 = _dafny.Seq.of(_72_v1, _72_v1, _72_v1, _72_v1);
      let _116_v41;
      _116_v41 = _dafny.Seq.of((_114_v39).update(_113_v38, new BigNumber((_dafny.MultiSet.FromArray(_115_v40)).cardinality())));
      let _hi0 = ((_72_v1) ? (_107_v32) : (_107_v32));
      for (let _117_i3 = new BigNumber((_116_v41).length); _117_i3.isLessThan(_hi0); _117_i3 = _117_i3.plus(_dafny.ONE)) {
        let _118_v42;
        let _nw11 = Array((new BigNumber(27)).toNumber()).fill(_dafny.Map.Empty);
        _118_v42 = _nw11;
        let _119_v43;
        let _nw12 = Array((new BigNumber(25)).toNumber());
        _119_v43 = _nw12;
        let _120_v44;
        _120_v44 = _dafny.Map.Empty.slice().updateUnsafe(_107_v32,_119_v43);
        let _index4 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_118_v42).length));
        (_118_v42)[_index4] = (_120_v44).Merge(_120_v44);
        let _121_v45;
        let _nw13 = new _module.C4();
        _nw13.__ctor(new BigNumber(472), new BigNumber((_module.__default.fm31(_71_globalState)).length), _72_v1);
        _121_v45 = _nw13;
        let _122_v46;
        _122_v46 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(_117_i3, new BigNumber((_dafny.Seq.of(_121_v45, _121_v45)).length)),_120_v44);
        let _index5 = _module.__default.safeIndex(new BigNumber(952), new BigNumber((_118_v42).length));
        (_118_v42)[_index5] = (((_122_v46).contains((_dafny.ZERO).minus((_107_v32).plus(_117_i3)))) ? ((_122_v46).get((_dafny.ZERO).minus((_107_v32).plus(_117_i3)))) : (_120_v44));
        let _123_v47;
        _123_v47 = _dafny.Map.Empty.slice().updateUnsafe(_111_v36,_113_v38.f13);
        let _124_v48;
        let _nw14 = Array((new BigNumber(29)).toNumber());
        _nw14[(_dafny.ZERO).toNumber()] = _dafny.Seq.UnicodeFromString("wiusmh");
        _nw14[(_dafny.ONE).toNumber()] = _module.__default.fm20((_121_v45).f3, _71_globalState);
        _nw14[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(584)), ((_125_v38) => function (_126_i4) {
          return (_125_v38).f12;
        })(_113_v38)), _dafny.Seq.UnicodeFromString("tr"));
        _nw14[(new BigNumber(3)).toNumber()] = _113_v38.f13;
        _nw14[(new BigNumber(4)).toNumber()] = _dafny.Seq.UnicodeFromString("bybpuk");
        _nw14[(new BigNumber(5)).toNumber()] = _dafny.Seq.UnicodeFromString("qkhmicrt");
        _nw14[(new BigNumber(6)).toNumber()] = _111_v36;
        _nw14[(new BigNumber(7)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-349)), ((_127_v35) => function (_128_i5) {
          return _127_v35;
        })(_110_v35));
        _nw14[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_111_v36, _dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(293)), ((_129_v38) => function (_130_i6) {
          return (_129_v38).f12;
        })(_113_v38)), _module.__default.safeIndex(new BigNumber(53), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(293)), ((_131_v38) => function (_132_i6) {
          return (_131_v38).f12;
        })(_113_v38))).length)), (_module.D8.create_DC23(_121_v45.f4, _72_v1, true, (_113_v38).f12)).dtor_cf35));
        _nw14[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_module.__default.fm20(new BigNumber(-245), _71_globalState), _111_v36);
        _nw14[(new BigNumber(10)).toNumber()] = _113_v38.f13;
        _nw14[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_113_v38.f13, _111_v36);
        _nw14[(new BigNumber(12)).toNumber()] = _dafny.Seq.UnicodeFromString("crcifcx");
        _nw14[(new BigNumber(13)).toNumber()] = _113_v38.f13;
        _nw14[(new BigNumber(14)).toNumber()] = _113_v38.f13;
        _nw14[(new BigNumber(15)).toNumber()] = _113_v38.f13;
        _nw14[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_111_v36, _113_v38.f13);
        _nw14[(new BigNumber(17)).toNumber()] = _module.__default.fm20((_121_v45).f7, _71_globalState);
        _nw14[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_113_v38.f13, _111_v36);
        _nw14[(new BigNumber(19)).toNumber()] = _dafny.Seq.UnicodeFromString("jluu");
        _nw14[(new BigNumber(20)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wagffgkjj"), _113_v38.f13), _module.__default.safeIndex(_107_v32, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("wagffgkjj"), _113_v38.f13)).length)), (_113_v38).f12);
        _nw14[(new BigNumber(21)).toNumber()] = _111_v36;
        _nw14[(new BigNumber(22)).toNumber()] = _dafny.Seq.UnicodeFromString("agrk");
        _nw14[(new BigNumber(23)).toNumber()] = ((_72_v1) ? (_113_v38.f13) : ((((_123_v47).contains(_dafny.Seq.UnicodeFromString("x"))) ? ((_123_v47).get(_dafny.Seq.UnicodeFromString("x"))) : (_dafny.Seq.UnicodeFromString("k")))));
        _nw14[(new BigNumber(24)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-793)), ((_133_v35) => function (_134_i7) {
          return _133_v35;
        })(_110_v35));
        _nw14[(new BigNumber(25)).toNumber()] = _113_v38.f13;
        _nw14[(new BigNumber(26)).toNumber()] = _113_v38.f13;
        _nw14[(new BigNumber(27)).toNumber()] = _dafny.Seq.Concat(_111_v36, _113_v38.f13);
        _nw14[(new BigNumber(28)).toNumber()] = ((_121_v45.f4) ? (_dafny.Seq.Create(_module.__default.abs(new BigNumber(854)), ((_135_v35) => function (_136_i8) {
          return _135_v35;
        })(_110_v35))) : (_111_v36));
        _124_v48 = _nw14;
        _124_v48 = _124_v48;
        if (!(_121_v45.f4) || (!(true) || (_72_v1))) {
          let _137_v49;
          _137_v49 = _dafny.Map.Empty.slice().updateUnsafe(((_121_v45).f3).plus((_121_v45).f3),_110_v35);
          _137_v49 = (_137_v49).update((_113_v38).fm12(!(_121_v45.f4), _121_v45.f4, _71_globalState), _110_v35);
          (_113_v38).m5(_71_globalState);
          (_71_globalState).f1 = _117_i3;
          let _138_v50;
          let _nw15 = Array((new BigNumber(16)).toNumber());
          _138_v50 = _nw15;
          let _index6 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_138_v50).length));
          (_138_v50)[_index6] = _121_v45;
          let _index7 = _module.__default.safeIndex(new BigNumber(202), new BigNumber((_138_v50).length));
          let _nw16 = new _module.C5();
          _nw16.__ctor((_109_v34).f9, (_121_v45).f3, new BigNumber(-406), _121_v45.f4);
          (_138_v50)[_index7] = _nw16;
          let _139_v51;
          let _nw17 = new _module.C7();
          _nw17.__ctor(_121_v45.f4, (_121_v45).f7, (_121_v45).f7, (_108_v33).IsProperSubsetOf(_108_v33));
          _139_v51 = _nw17;
        } else {
          (_121_v45).f4 = _121_v45.f4;
          let _140_v52;
          let _nw18 = new _module.C7();
          _nw18.__ctor(_121_v45.f4, (_121_v45).f7, new BigNumber((_111_v36).length), _72_v1);
          _140_v52 = _nw18;
          let _141_v53;
          _141_v53 = _dafny.Seq.of(_140_v52, _140_v52, _140_v52);
          let _index8 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_112_v37).length));
          (_112_v37)[_index8] = new BigNumber((_141_v53).length);
          let _index9 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_112_v37).length));
          (_112_v37)[_index9] = (_140_v52).f6;
          let _142_v54;
          _142_v54 = _dafny.Set.fromElements(_109_v34, _109_v34);
          _142_v54 = _142_v54;
          let _143_v55;
          let _nw19 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _143_v55 = _nw19;
          let _144_v56;
          let _nw20 = new _module.C2();
          _nw20.__ctor((_121_v45).f7, _143_v55, _107_v32, _72_v1);
          _144_v56 = _nw20;
          let _145_v57;
          _145_v57 = _dafny.Set.fromElements(_144_v56);
          let _146_v58;
          _146_v58 = _dafny.Map.Empty.slice().updateUnsafe(_121_v45.f4,new BigNumber((_145_v57).length));
          let _147_v59;
          _147_v59 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_146_v58).Merge(_146_v58)).length),_121_v45.f4);
          let _rhs5 = !((((_147_v59).contains((_dafny.ZERO).minus(_117_i3))) ? ((_147_v59).get((_dafny.ZERO).minus(_117_i3))) : (_121_v45.f4)));
          let _rhs6 = _121_v45.f4;
          let _rhs7 = (_121_v45).f7;
          let _lhs2 = _140_v52;
          let _lhs3 = _71_globalState;
          _72_v1 = _rhs5;
          _lhs2.f5 = _rhs6;
          _lhs3.f1 = _rhs7;
          _143_v55 = _143_v55;
        }
        let _148_v60;
        _148_v60 = _dafny.Map.Empty.slice().updateUnsafe(_121_v45.f4,_111_v36);
        (_121_v45).f4 = (((_113_v38).fm12(_module.__default.fm2(new BigNumber(((((_148_v60).contains((_113_v38).fm13(_71_globalState))) ? ((_148_v60).get((_113_v38).fm13(_71_globalState))) : (_dafny.Seq.UnicodeFromString("epeew")))).length), _71_globalState), _72_v1, _71_globalState)).minus((_121_v45).f7)).isLessThanOrEqualTo((_121_v45).f3);
      }
      let _hi1 = _107_v32;
      for (let _149_i9 = _107_v32; _149_i9.isLessThan(_hi1); _149_i9 = _149_i9.plus(_dafny.ONE)) {
        let _150_v61;
        _150_v61 = _module.D8.create_DC23(_72_v1, !(_72_v1), false, new _dafny.CodePoint('e'.codePointAt(0)));
        let _151_v62;
        _151_v62 = _module.D8.create_DC24(_150_v61);
        let _152_v63;
        _152_v63 = _module.D8.create_DC24(_150_v61);
        let _153_v64;
        _153_v64 = _module.D13.create_DC40(_152_v63, (_dafny.ZERO).minus(_107_v32));
        let _154_v65;
        _154_v65 = _module.D12.create_DC36(_72_v1, !(_72_v1), _module.__default.fm0(_71_globalState));
        let _155_v66;
        _155_v66 = _module.D12.create_DC36(_72_v1, _72_v1, (_154_v65).dtor_cf63);
        let _rhs8 = ((_154_v65).dtor_cf63).isLessThanOrEqualTo(_107_v32);
        let _rhs9 = _72_v1;
        let _rhs10 = _153_v64;
        let _lhs4 = _71_globalState;
        _72_v1 = _rhs8;
        _lhs4.f2 = _rhs9;
        _153_v64 = _rhs10;
        let _156_v67;
        _156_v67 = _dafny.MultiSet.fromElements((_113_v38).f12);
        let _157_v68;
        let _nw21 = new _module.C4();
        _nw21.__ctor(_107_v32, new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_72_v1))).cardinality()), (_module.D0.create_DC2(_156_v67, !(_72_v1))).dtor_cf6);
        _157_v68 = _nw21;
        let _rhs11 = _113_v38;
        let _rhs12 = _157_v68;
        _113_v38 = _rhs11;
        _157_v68 = _rhs12;
        _107_v32 = _107_v32;
        let _nw22 = Array((new BigNumber(10)).toNumber()).fill(false);
        _106_v31 = _nw22;
      }
      let _158_v69;
      let _nw23 = new _module.C1();
      _nw23.__ctor(new BigNumber((_111_v36).length), new BigNumber(840), (_115_v40)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("p")).length), new BigNumber((_115_v40).length))]);
      _158_v69 = _nw23;
      let _159_v70;
      _159_v70 = _dafny.Seq.of(_158_v69);
      let _160_v71;
      _160_v71 = _dafny.Set.fromElements(_158_v69, _158_v69, _158_v69, (_159_v70)[_module.__default.safeIndex(_107_v32, new BigNumber((_159_v70).length))]);
      _160_v71 = (_160_v71).Intersect(_160_v71);
      let _hi2 = (new BigNumber(273)).minus(_107_v32);
      for (let _161_i10 = _107_v32; _161_i10.isLessThan(_hi2); _161_i10 = _161_i10.plus(_dafny.ONE)) {
        let _162_v72;
        _162_v72 = _dafny.MultiSet.fromElements(_72_v1, _72_v1);
        let _163_v73;
        _163_v73 = _dafny.MultiSet.fromElements(_161_i10);
        (_113_v38).f13 = _module.__default.fm15((_72_v1) || (_72_v1), _module.__default.safeModuloInt((((_162_v72).contains(!(_72_v1))) ? ((_162_v72).get(!(_72_v1))) : (new BigNumber(-565))), _161_i10), (((_163_v73).contains(_161_i10)) ? ((_163_v73).get(_161_i10)) : (_107_v32)), _71_globalState);
        let _index10 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_106_v31).length));
        (_106_v31)[_index10] = _72_v1;
        let _index11 = _module.__default.safeIndex(new BigNumber(840), new BigNumber((_106_v31).length));
        (_106_v31)[_index11] = (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_72_v1,_161_i10)).length)).isLessThan(new BigNumber(-570));
        let _164_v74;
        _164_v74 = _dafny.Set.fromElements(_113_v38, _113_v38, _113_v38, _113_v38, _113_v38);
        _164_v74 = _164_v74;
        let _165_v75;
        let _init3 = ((_166_v31) => function (_167_i11) {
          return !((_166_v31)[_module.__default.safeIndex(new BigNumber(840), new BigNumber((_166_v31).length))]) || ((_166_v31)[_module.__default.safeIndex(new BigNumber(840), new BigNumber((_166_v31).length))]);
        })(_106_v31);
        let _nw24 = Array((_dafny.ONE).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw24.length); _i0_3++) {
          _nw24[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _165_v75 = _nw24;
        let _index12 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_69_v0).length));
        (_69_v0)[_index12] = (_113_v38).f12;
        let _index13 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_69_v0).length));
        let _rhs13 = _165_v75;
        let _rhs14 = _module.__default.safeModuloInt((_107_v32).minus(_107_v32), _107_v32);
        let _rhs15 = _161_i10;
        let _rhs16 = (_113_v38).f12;
        let _rhs17 = _72_v1;
        let _lhs5 = _71_globalState;
        let _lhs6 = _71_globalState;
        let _lhs7 = _69_v0;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_69_v0).length));
        _165_v75 = _rhs13;
        _lhs5.f1 = _rhs14;
        _lhs6.f1 = _rhs15;
        _lhs7[_lhs8] = _rhs16;
        _72_v1 = _rhs17;
      }
      _72_v1 = _72_v1;
      let _168_v76;
      _168_v76 = _dafny.Set.fromElements(_107_v32);
      let _169_i12;
      _169_i12 = _dafny.ZERO;
      L0: {
        while (!(!(_107_v32).isEqualTo(_module.__default.safeModuloInt(new BigNumber((_168_v76).length), (_dafny.ZERO).minus(_107_v32))))) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_169_i12)) {
              break L0;
            }
            _169_i12 = (_169_i12).plus(_dafny.ONE);
            _107_v32 = _107_v32;
            let _index14 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_112_v37).length));
            (_112_v37)[_index14] = _module.__default.safeDivisionInt(_107_v32, _module.__default.fm0(_71_globalState));
            let _170_v77;
            let _nw25 = Array((new BigNumber(17)).toNumber()).fill([]);
            _170_v77 = _nw25;
            let _171_v78;
            let _nw26 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
            _171_v78 = _nw26;
            let _172_v79;
            _172_v79 = _dafny.Set.fromElements((_113_v38).f12, new _dafny.CodePoint('k'.codePointAt(0)));
            let _173_v80;
            let _nw27 = new _module.C5();
            _nw27.__ctor((_109_v34).f9, new BigNumber((_113_v38.f13).length), new BigNumber((_172_v79).length), _72_v1);
            _173_v80 = _nw27;
            let _174_v81;
            _174_v81 = _dafny.Set.fromElements(_173_v80);
            let _175_v82;
            _175_v82 = _dafny.Map.Empty.slice().updateUnsafe(_72_v1,_174_v81);
            let _index15 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_171_v78).length));
            (_171_v78)[_index15] = _175_v82;
            let _index16 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_112_v37).length));
            let _index17 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_171_v78).length));
            let _rhs18 = new BigNumber((_168_v76).length);
            let _rhs19 = _170_v77;
            let _rhs20 = (((_175_v82).update(_173_v80.f4, _174_v81)).Merge(_175_v82)).update(_72_v1, _174_v81);
            let _lhs9 = _112_v37;
            let _lhs10 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_112_v37).length));
            let _lhs11 = _171_v78;
            let _lhs12 = _module.__default.safeIndex(new BigNumber(388), new BigNumber((_171_v78).length));
            _lhs9[_lhs10] = _rhs18;
            _170_v77 = _rhs19;
            _lhs11[_lhs12] = _rhs20;
            if (_173_v80.f4) {
              let _176_v83;
              let _init4 = ((_177_v80) => function (_178_i13) {
                return (_178_i13).multipliedBy((_177_v80).f7);
              })(_173_v80);
              let _nw28 = Array((new BigNumber(9)).toNumber());
              for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw28.length); _i0_4++) {
                _nw28[_i0_4] = _init4(new BigNumber(_i0_4));
              }
              _176_v83 = _nw28;
              let _179_v84;
              let _nw29 = Array((new BigNumber(29)).toNumber());
              _179_v84 = _nw29;
              let _180_v85;
              _180_v85 = _dafny.Map.Empty.slice().updateUnsafe(_179_v84,_113_v38);
              let _rhs21 = _176_v83;
              let _rhs22 = (((_180_v85).contains(_179_v84)) ? ((_180_v85).get(_179_v84)) : (_113_v38));
              _176_v83 = _rhs21;
              _113_v38 = _rhs22;
              let _181_v86;
              let _nw30 = new _module.C2();
              _nw30.__ctor(_107_v32, _176_v83, (_173_v80).f7, _173_v80.f4);
              _181_v86 = _nw30;
              let _182_v87;
              _182_v87 = _dafny.Map.Empty.slice().updateUnsafe(_173_v80.f4,_173_v80.f4);
              let _183_v88;
              _183_v88 = _module.D12.create_DC37(_182_v87, _72_v1, (_112_v37)[_module.__default.safeIndex(new BigNumber(699), new BigNumber((_112_v37).length))], (_173_v80).f7);
              let _184_v89;
              _184_v89 = _dafny.Seq.of((_173_v80).f7, _module.__default.safeModuloInt((_173_v80).f7, (_173_v80).f7), (_183_v88).dtor_cf66);
              let _185_v90;
              _185_v90 = _dafny.Map.Empty.slice().updateUnsafe((_181_v86).f14,_dafny.Seq.of(_173_v80.f4, _173_v80.f4));
              let _186_v91;
              _186_v91 = _dafny.MultiSet.fromElements(_113_v38.f13, _111_v36);
              let _187_v92;
              _187_v92 = _dafny.MultiSet.fromElements((_173_v80).f3);
              let _index18 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_112_v37).length));
              let _rhs23 = _181_v86;
              let _rhs24 = _dafny.Seq.of(((_173_v80).f7).isEqualTo((_dafny.ZERO).minus(new BigNumber(((((_185_v90).contains((_173_v80).f7)) ? ((_185_v90).get((_173_v80).f7)) : (_115_v40))).length))), !(new BigNumber((_186_v91).cardinality())).isEqualTo((_173_v80).f3), _173_v80.f4);
              let _rhs25 = _module.__default.fm0(_71_globalState);
              let _rhs26 = _184_v89;
              let _rhs27 = _module.__default.fm4(_173_v80.f4, _72_v1, _module.__default.fm2((_dafny.ZERO).minus((_173_v80).f3), _71_globalState), (_107_v32).isLessThan(new BigNumber((_187_v92).cardinality())), _71_globalState);
              let _lhs13 = _112_v37;
              let _lhs14 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_112_v37).length));
              _181_v86 = _rhs23;
              _115_v40 = _rhs24;
              _lhs13[_lhs14] = _rhs25;
              _184_v89 = _rhs26;
              _184_v89 = _rhs27;
              let _188_v93;
              _188_v93 = _dafny.Map.Empty.slice().updateUnsafe((_173_v80).f7,_173_v80.f4);
              (_71_globalState).f2 = (((_188_v93).contains((_181_v86).f14)) ? ((_188_v93).get((_181_v86).f14)) : (_173_v80.f4));
              (_173_v80).f4 = _module.__default.fm2((_dafny.ZERO).minus(_module.__default.safeDivisionInt((_173_v80).f7, (((_187_v92).contains((_173_v80).f7)) ? ((_187_v92).get((_173_v80).f7)) : ((_173_v80).f3)))), _71_globalState);
              let _189_v94;
              let _nw31 = new _module.C7();
              _nw31.__ctor(!(false), _module.__default.fm0(_71_globalState), _107_v32, _173_v80.f4);
              _189_v94 = _nw31;
              let _190_v95;
              let _nw32 = Array((new BigNumber(5)).toNumber());
              _nw32[(_dafny.ZERO).toNumber()] = _189_v94;
              _nw32[(_dafny.ONE).toNumber()] = _189_v94;
              _nw32[(new BigNumber(2)).toNumber()] = _189_v94;
              _nw32[(new BigNumber(3)).toNumber()] = _189_v94;
              _nw32[(new BigNumber(4)).toNumber()] = _189_v94;
              _190_v95 = _nw32;
              let _index19 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_190_v95).length));
              (_190_v95)[_index19] = _189_v94;
              let _index20 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_112_v37).length));
              let _index21 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_190_v95).length));
              let _rhs28 = (_181_v86).f14;
              let _rhs29 = _module.__default.fm0(_71_globalState);
              let _rhs30 = (_173_v80).f7;
              let _rhs31 = (_181_v86).f15;
              let _rhs32 = _189_v94;
              let _lhs15 = _71_globalState;
              let _lhs16 = _112_v37;
              let _lhs17 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_112_v37).length));
              let _lhs18 = _71_globalState;
              let _lhs19 = _190_v95;
              let _lhs20 = _module.__default.safeIndex(new BigNumber(150), new BigNumber((_190_v95).length));
              _lhs15.f1 = _rhs28;
              _lhs16[_lhs17] = _rhs29;
              _lhs18.f1 = _rhs30;
              _176_v83 = _rhs31;
              _lhs19[_lhs20] = _rhs32;
            } else {
              let _191_v96;
              _191_v96 = _dafny.Map.Empty.slice().updateUnsafe((_173_v80).f7,_69_v0);
              let _192_v97;
              _192_v97 = _dafny.MultiSet.fromElements((((_191_v96).contains((_173_v80).f7)) ? ((_191_v96).get((_173_v80).f7)) : (_69_v0)));
              _192_v97 = ((_192_v97).Difference(_dafny.MultiSet.fromElements(_69_v0))).Union((_192_v97).Difference(_192_v97));
              let _193_v98;
              _193_v98 = _dafny.Map.Empty.slice().updateUnsafe(_109_v34,_111_v36);
              let _194_v99;
              _194_v99 = _dafny.MultiSet.fromElements(true);
              let _195_v100;
              _195_v100 = _dafny.Map.Empty.slice().updateUnsafe(((_193_v98).update(_109_v34, _113_v38.f13)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_109_v34,_113_v38.f13)),(new BigNumber((_108_v33).length)).multipliedBy(new BigNumber((_194_v99).cardinality())));
              _195_v100 = (_195_v100).update(_193_v98, (_173_v80).f3);
              let _index22 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_112_v37).length));
              (_112_v37)[_index22] = (_113_v38).fm12(_module.__default.fm2(_107_v32, _71_globalState), _dafny.areEqual(_dafny.Seq.UnicodeFromString("jd"), _111_v36), _71_globalState);
              (_173_v80).f4 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.update(_dafny.Seq.Concat(_111_v36, _113_v38.f13), _module.__default.safeIndex((_173_v80).f3, new BigNumber((_dafny.Seq.Concat(_111_v36, _113_v38.f13)).length)), (_113_v38).f12), _111_v36);
              let _196_v101;
              _196_v101 = _module.D12.create_DC35((_109_v34).f9, true, _dafny.Seq.of(_72_v1), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-514)), ((_197_v38) => function (_198_i14) {
  return (_197_v38).f12;
})(_113_v38)), (_173_v80).f7);
              let _pat_let_tv0 = _113_v38;
              let _pat_let_tv1 = _173_v80;
              let _pat_let_tv2 = _115_v40;
              let _pat_let_tv3 = _106_v31;
              _111_v36 = (function (_pat_let0_0) {
                return function (_199_dt__update__tmp_h0) {
                  return function (_pat_let1_0) {
                    return function (_200_dt__update_hcf59_h0) {
                      return function (_pat_let2_0) {
                        return function (_201_dt__update_hcf60_h0) {
                          return function (_pat_let3_0) {
                            return function (_202_dt__update_hcf58_h0) {
                              return function (_pat_let4_0) {
                                return function (_203_dt__update_hcf56_h0) {
                                  return _module.D12.create_DC35(_203_dt__update_hcf56_h0, (_199_dt__update__tmp_h0).dtor_cf57, _202_dt__update_hcf58_h0, _200_dt__update_hcf59_h0, _201_dt__update_hcf60_h0);
                                }(_pat_let4_0);
                              }(_pat_let_tv3);
                            }(_pat_let3_0);
                          }(_pat_let_tv2);
                        }(_pat_let2_0);
                      }((_pat_let_tv1).f7);
                    }(_pat_let1_0);
                  }(_pat_let_tv0.f13);
                }(_pat_let0_0);
              }(_196_v101)).dtor_cf59;
            }
            (_113_v38).m5(_71_globalState);
          }
        }
      }
      let _204_v102;
      _204_v102 = _dafny.Map.Empty.slice().updateUnsafe(_72_v1,_110_v35);
      let _205_v103;
      _205_v103 = _dafny.Seq.of(_107_v32);
      let _206_i15;
      _206_i15 = _dafny.ZERO;
      L1: {
        while (!(_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-792)), ((_217_v32) => function (_218_i16) {
          return _217_v32;
        })(_107_v32)), _205_v103)) || (_dafny.areEqual(_dafny.Seq.update(_113_v38.f13, _module.__default.safeIndex(_107_v32, new BigNumber((_113_v38.f13).length)), (((_204_v102).contains(_72_v1)) ? ((_204_v102).get(_72_v1)) : ((_113_v38).f12))), _111_v36))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_206_i15)) {
              break L1;
            }
            _206_i15 = (_206_i15).plus(_dafny.ONE);
            let _207_v104;
            _207_v104 = _dafny.MultiSet.fromElements(_72_v1);
            (_71_globalState).f1 = ((new BigNumber((_113_v38.f13).length)).minus((_dafny.ZERO).minus(_107_v32))).minus(new BigNumber((_207_v104).cardinality()));
            let _208_v105;
            _208_v105 = _dafny.Map.Empty.slice().updateUnsafe((_72_v1) && (_72_v1),_72_v1);
            _208_v105 = (_208_v105).update((_107_v32).isLessThanOrEqualTo(_107_v32), _72_v1);
            let _209_v106;
            let _nw33 = new _module.C4();
            _nw33.__ctor(new BigNumber((_207_v104).cardinality()), (_dafny.ZERO).minus(new BigNumber((_205_v103).length)), _module.__default.fm2(_107_v32, _71_globalState));
            _209_v106 = _nw33;
            let _210_v107;
            _210_v107 = _dafny.Map.Empty.slice().updateUnsafe(_72_v1,_209_v106);
            let _211_v108;
            _211_v108 = _dafny.Seq.of(_209_v106, _209_v106);
            let _212_v109;
            let _nw34 = Array((new BigNumber(14)).toNumber());
            _nw34[(_dafny.ZERO).toNumber()] = _209_v106;
            _nw34[(_dafny.ONE).toNumber()] = _209_v106;
            _nw34[(new BigNumber(2)).toNumber()] = _209_v106;
            _nw34[(new BigNumber(3)).toNumber()] = _209_v106;
            _nw34[(new BigNumber(4)).toNumber()] = _209_v106;
            _nw34[(new BigNumber(5)).toNumber()] = _209_v106;
            _nw34[(new BigNumber(6)).toNumber()] = _209_v106;
            _nw34[(new BigNumber(7)).toNumber()] = (((_210_v107).contains(true)) ? ((_210_v107).get(true)) : (_209_v106));
            _nw34[(new BigNumber(8)).toNumber()] = _209_v106;
            _nw34[(new BigNumber(9)).toNumber()] = _209_v106;
            _nw34[(new BigNumber(10)).toNumber()] = _209_v106;
            _nw34[(new BigNumber(11)).toNumber()] = _209_v106;
            _nw34[(new BigNumber(12)).toNumber()] = (_211_v108)[_module.__default.safeIndex(_107_v32, new BigNumber((_211_v108).length))];
            _nw34[(new BigNumber(13)).toNumber()] = _209_v106;
            _212_v109 = _nw34;
            let _index23 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_212_v109).length));
            (_212_v109)[_index23] = _209_v106;
            let _index24 = _module.__default.safeIndex(new BigNumber(199), new BigNumber((_212_v109).length));
            (_212_v109)[_index24] = _209_v106;
            let _213_v110;
            let _214_v111;
            let _215_v112;
            let _216_v113;
            let _out1;
            let _out2;
            let _out3;
            let _out4;
            let _outcollector0 = (_109_v34).m2((_72_v1) === (true), _112_v37, (_208_v105).Merge(_208_v105), _72_v1, _71_globalState);
            _out1 = _outcollector0[0];
            _out2 = _outcollector0[1];
            _out3 = _outcollector0[2];
            _out4 = _outcollector0[3];
            _213_v110 = _out1;
            _214_v111 = _out2;
            _215_v112 = _out3;
            _216_v113 = _out4;
          }
        }
      }
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_106_v31).length))) {
        let _219_i17 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_219_i17)) && ((_219_i17).isLessThan(new BigNumber((_106_v31).length))))) {
          (_106_v31)[(_219_i17)] = !(!(!(_72_v1)));
        }
      }
      let _220_v114;
      let _nw35 = Array((new BigNumber(4)).toNumber());
      _nw35[(_dafny.ZERO).toNumber()] = _113_v38;
      _nw35[(_dafny.ONE).toNumber()] = ((_72_v1) ? (_113_v38) : (_113_v38));
      _nw35[(new BigNumber(2)).toNumber()] = _113_v38;
      _nw35[(new BigNumber(3)).toNumber()] = _113_v38;
      _220_v114 = _nw35;
      let _index25 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_220_v114).length));
      (_220_v114)[_index25] = _113_v38;
      let _index26 = _module.__default.safeIndex(new BigNumber(998), new BigNumber((_220_v114).length));
      (_220_v114)[_index26] = _113_v38;
      let _221_v115;
      let _nw36 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Map.Empty);
      _221_v115 = _nw36;
      let _222_v116;
      let _nw37 = new _module.C3();
      _nw37.__ctor((_113_v38).f12, _111_v36, new BigNumber((_113_v38.f13).length), true, _107_v32);
      _222_v116 = _nw37;
      let _223_v117;
      let _nw38 = new _module.C2();
      _nw38.__ctor((_222_v116).f3, _112_v37, (_222_v116).f3, _72_v1);
      _223_v117 = _nw38;
      let _224_v118;
      _224_v118 = _dafny.Map.Empty.slice().updateUnsafe(_222_v116,_223_v117);
      let _index27 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_221_v115).length));
      (_221_v115)[_index27] = (_224_v118).Merge(_224_v118);
      let _index28 = _module.__default.safeIndex(new BigNumber(91), new BigNumber((_221_v115).length));
      (_221_v115)[_index28] = _224_v118;
      let _225_v119;
      _225_v119 = _dafny.Seq.of(_112_v37, _112_v37);
      let _226_v120;
      _226_v120 = _dafny.Seq.of(_225_v119, (_module.D15.create_DC43(_225_v119)).dtor_cf74);
      let _index29 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length));
      (_106_v31)[_index29] = !_dafny.Seq.contains((_226_v120)[_module.__default.safeIndex((_223_v117).f14, new BigNumber((_226_v120).length))], (_223_v117).f15);
      let _index30 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length));
      (_106_v31)[_index30] = _72_v1;
      let _227_i18;
      _227_i18 = _dafny.ZERO;
      L2: {
        while (_72_v1) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_227_i18)) {
              break L2;
            }
            _227_i18 = (_227_i18).plus(_dafny.ONE);
            let _228_v121;
            _228_v121 = _dafny.Map.Empty.slice().updateUnsafe(_222_v116.f4,_222_v116.f4);
            _228_v121 = (_228_v121).update((_106_v31)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length))], _222_v116.f4);
            let _229_v122;
            _229_v122 = _dafny.MultiSet.fromElements(_107_v32);
            (_71_globalState).f2 = _module.__default.fm2(new BigNumber(((_229_v122).Intersect(_229_v122)).cardinality()), _71_globalState);
            let _230_v123;
            let _nw39 = new _module.C0();
            _nw39.__ctor(_222_v116.f4, (_109_v34).f9);
            _230_v123 = _nw39;
            (_222_v116).f4 = !(_222_v116.f4);
          }
        }
      }
      let _index31 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_112_v37).length));
      (_112_v37)[_index31] = (_222_v116).f3;
      let _index32 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_112_v37).length));
      (_112_v37)[_index32] = _107_v32;
      let _231_v124;
      _231_v124 = _dafny.Map.Empty.slice().updateUnsafe(_72_v1,false);
      let _232_v125;
      _232_v125 = _module.D13.create_DC38(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(144)), ((_233_v116, _234_v31) => function (_235_i19) {
  return _dafny.Map.Empty.slice().updateUnsafe(_233_v116.f4,(_234_v31)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_234_v31).length))]);
})(_222_v116, _106_v31)), _module.__default.safeIndex((_223_v117).f14, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(144)), ((_236_v116, _237_v31) => function (_238_i19) {
  return _dafny.Map.Empty.slice().updateUnsafe(_236_v116.f4,(_237_v31)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_237_v31).length))]);
})(_222_v116, _106_v31))).length)), _231_v124));
      let _source3 = _232_v125;
      if (_source3.is_DC39) {
        let _239___mcc_h0 = (_source3).cf69;
        let _240_cf69 = _239___mcc_h0;
        (_71_globalState).f2 = _222_v116.f4;
        _111_v36 = _dafny.Seq.Concat(_module.__default.fm15((_106_v31)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length))], (_223_v117).f14, (_222_v116).f3, _71_globalState), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-238)), ((_241_v38) => function (_242_i20) {
          return (_241_v38).f12;
        })(_113_v38)));
        let _243_v126;
        _243_v126 = _dafny.MultiSet.fromElements(_222_v116.f4);
        let _244_v127;
        _244_v127 = _dafny.Map.Empty.slice().updateUnsafe((_223_v117).f14,new BigNumber((_dafny.Seq.UnicodeFromString("sqq")).length));
        (_71_globalState).f1 = new BigNumber((((((_113_v38).fm13(_71_globalState)) ? (_243_v126) : (_243_v126))).update((_module.__default.fm2(_107_v32, _71_globalState)) && (_222_v116.f4), _module.__default.abs((((_244_v127).contains(new BigNumber((_dafny.Set.fromElements(_109_v34)).length))) ? ((_244_v127).get(new BigNumber((_dafny.Set.fromElements(_109_v34)).length))) : (_module.__default.fm0(_71_globalState)))))).cardinality());
        let _245_v128;
        let _nw40 = new _module.C3();
        _nw40.__ctor((_113_v38).f12, _113_v38.f13, new BigNumber((_168_v76).length), _222_v116.f4, (_223_v117).f14);
        _245_v128 = _nw40;
        let _246_v129;
        _246_v129 = _dafny.MultiSet.fromElements(_245_v128);
        let _247_v130;
        _247_v130 = _module.D15.create_DC44(_245_v128, _110_v35, _107_v32);
        let _248_v131;
        _248_v131 = _dafny.Seq.of(_245_v128);
        _246_v129 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of((_247_v130).dtor_cf75), _248_v131))).Difference(_dafny.MultiSet.FromArray(_248_v131));
      } else if (_source3.is_DC40) {
        let _249___mcc_h1 = (_source3).cf70;
        let _250___mcc_h2 = (_source3).cf71;
        let _251_cf71 = _250___mcc_h2;
        let _252_cf70 = _249___mcc_h1;
        let _253_v132;
        _253_v132 = _dafny.Seq.of(_115_v40, _dafny.Seq.of(_222_v116.f4, _72_v1, (_106_v31)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length))], _222_v116.f4));
        let _254_v133;
        let _nw41 = new _module.C6();
        _nw41.__ctor((_223_v117).f14, _251_cf71, _222_v116.f4, new BigNumber(((_253_v132)[_module.__default.safeIndex(_107_v32, new BigNumber((_253_v132).length))]).length));
        _254_v133 = _nw41;
        let _255_v134;
        _255_v134 = _dafny.Seq.of(_254_v133, _254_v133, _254_v133);
        let _256_v135;
        _256_v135 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.MultiSet.FromArray(_205_v103)).cardinality()), (_223_v117).f14, new BigNumber(844));
        let _257_v136;
        _257_v136 = _module.D3.create_DC10(new BigNumber((_256_v135).cardinality()), (_109_v34).f9, new BigNumber((_dafny.MultiSet.fromElements((_dafny.ZERO).minus((((_256_v135).contains((_223_v117).f14)) ? ((_256_v135).get((_223_v117).f14)) : ((_112_v37)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_112_v37).length))]))), (_223_v117).f14, (_222_v116).f3)).cardinality()));
        let _258_v137;
        let _nw42 = Array((new BigNumber(24)).toNumber());
        _nw42[(_dafny.ZERO).toNumber()] = _254_v133;
        _nw42[(_dafny.ONE).toNumber()] = _254_v133;
        _nw42[(new BigNumber(2)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(3)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(4)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(5)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(6)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(7)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(8)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(9)).toNumber()] = ((_72_v1) ? (_254_v133) : (_254_v133));
        _nw42[(new BigNumber(10)).toNumber()] = (_255_v134)[_module.__default.safeIndex((_257_v136).dtor_cf15, new BigNumber((_255_v134).length))];
        _nw42[(new BigNumber(11)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(12)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(13)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(14)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(15)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(16)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(17)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(18)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(19)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(20)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(21)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(22)).toNumber()] = _254_v133;
        _nw42[(new BigNumber(23)).toNumber()] = _254_v133;
        _258_v137 = _nw42;
        let _rhs33 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("qxab"), _113_v38.f13)).length);
        let _rhs34 = _258_v137;
        _107_v32 = _rhs33;
        _258_v137 = _rhs34;
        let _259_v138;
        _259_v138 = _dafny.Map.Empty.slice().updateUnsafe(false,(_222_v116).f3);
        let _260_v139;
        _260_v139 = _module.D12.create_DC37(_231_v124, _72_v1, new BigNumber((_111_v36).length), new BigNumber((_259_v138).length));
        let _261_v140;
        _261_v140 = _module.D1.create_DC6(_222_v116, _dafny.Seq.UnicodeFromString("pvomejghm"), (_106_v31)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length))], _222_v116.f4);
        let _source4 = ((_222_v116.f4) ? (_260_v139) : (_module.D12.create_DC37(_231_v124, (_261_v140).dtor_cf12, (_222_v116).f3, (_223_v117).f14)));
        if (_source4.is_DC35) {
          let _262___mcc_h5 = (_source4).cf56;
          let _263___mcc_h6 = (_source4).cf57;
          let _264___mcc_h7 = (_source4).cf58;
          let _265___mcc_h8 = (_source4).cf59;
          let _266___mcc_h9 = (_source4).cf60;
          let _267_cf60 = _266___mcc_h9;
          let _268_cf59 = _265___mcc_h8;
          let _269_cf58 = _264___mcc_h7;
          let _270_cf57 = _263___mcc_h6;
          let _271_cf56 = _262___mcc_h5;
          let _272_v141;
          let _nw43 = new _module.C4();
          _nw43.__ctor(new BigNumber((_108_v33).length), (_254_v133).f8, _222_v116.f4);
          _272_v141 = _nw43;
          let _273_v142;
          _273_v142 = _dafny.Map.Empty.slice().updateUnsafe(_107_v32,_272_v141);
          let _274_v143;
          _274_v143 = _dafny.Map.Empty.slice().updateUnsafe(true,((_222_v116.f4) ? ((((_273_v142).contains(_251_cf71)) ? ((_273_v142).get(_251_cf71)) : (_272_v141))) : (_272_v141)));
          _274_v143 = _274_v143;
          let _275_v144;
          let _init5 = ((_276_v40) => function (_277_i21) {
            return _276_v40;
          })(_115_v40);
          let _nw44 = Array((new BigNumber(12)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw44.length); _i0_5++) {
            _nw44[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _275_v144 = _nw44;
          let _278_v145;
          _278_v145 = _dafny.Map.Empty.slice().updateUnsafe(_112_v37,(_254_v133).f8);
          (_223_v117).m6(_275_v144, _251_cf71, (new BigNumber((_278_v145).length)).plus((_223_v117).f14), _71_globalState);
          _231_v124 = (_231_v124).update(!(!(true)), _222_v116.f4);
          let _279_v146;
          let _nw45 = new _module.C7();
          _nw45.__ctor(_270_cf57, _module.__default.fm0(_71_globalState), (new BigNumber((_168_v76).length)).multipliedBy((_223_v117).f14), true);
          _279_v146 = _nw45;
        } else if (_source4.is_DC36) {
          let _280___mcc_h10 = (_source4).cf61;
          let _281___mcc_h11 = (_source4).cf62;
          let _282___mcc_h12 = (_source4).cf63;
          let _283_cf63 = _282___mcc_h12;
          let _284_cf62 = _281___mcc_h11;
          let _285_cf61 = _280___mcc_h10;
          let _286_v147;
          let _nw46 = Array((new BigNumber(2)).toNumber()).fill([]);
          _286_v147 = _nw46;
          let _287_v148;
          let _nw47 = new _module.C0();
          _nw47.__ctor(false, (_109_v34).f9);
          _287_v148 = _nw47;
          let _288_v149;
          let _nw48 = Array((new BigNumber(3)).toNumber());
          _nw48[(_dafny.ZERO).toNumber()] = _287_v148;
          _nw48[(_dafny.ONE).toNumber()] = _287_v148;
          _nw48[(new BigNumber(2)).toNumber()] = _287_v148;
          _288_v149 = _nw48;
          let _index33 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_286_v147).length));
          (_286_v147)[_index33] = _288_v149;
          let _289_v150;
          _289_v150 = _module.D12.create_DC35((_109_v34).f9, _72_v1, _115_v40, _113_v38.f13, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-843)), ((_290_v38) => function (_291_i22) {
  return (_290_v38).f12;
})(_113_v38))).length));
          let _292_v151;
          _292_v151 = _module.D4.create_DC11(_dafny.Set.fromElements(new BigNumber(737), (_289_v150).dtor_cf60));
          let _293_v152;
          _293_v152 = _dafny.Map.Empty.slice().updateUnsafe(_292_v151,_256_v135);
          let _index34 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_286_v147).length));
          let _rhs35 = (((_293_v152).contains(_292_v151)) ? ((_293_v152).get(_292_v151)) : ((_256_v135).Difference(_256_v135)));
          let _rhs36 = (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.FromArray(_205_v103)).cardinality()));
          let _rhs37 = _113_v38.f13;
          let _rhs38 = _115_v40;
          let _rhs39 = ((_222_v116.f4) ? (_288_v149) : (_288_v149));
          let _lhs21 = _71_globalState;
          let _lhs22 = _113_v38;
          let _lhs23 = _286_v147;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(524), new BigNumber((_286_v147).length));
          _256_v135 = _rhs35;
          _lhs21.f1 = _rhs36;
          _lhs22.f13 = _rhs37;
          _115_v40 = _rhs38;
          _lhs23[_lhs24] = _rhs39;
          let _294_v153;
          _294_v153 = _dafny.Set.fromElements(_112_v37, (_223_v117).f15, (_223_v117).f15, (_223_v117).f15);
          _294_v153 = ((_294_v153).Intersect(_dafny.Set.fromElements((_223_v117).f15))).Difference((_294_v153).Intersect(_294_v153));
          let _295_v155;
          _295_v155 = _dafny.MultiSet.fromElements(_110_v35, (_113_v38).f12, (_113_v38).f12, (_113_v38).f12, (_113_v38).f12);
          let _296_v156;
          _296_v156 = _dafny.Seq.of(function () {
            let _coll7 = new _dafny.Map();
            for (const _compr_7 of (_295_v155).Elements) {
              let _297_v154 = _compr_7;
              if ((_295_v155).contains(_297_v154)) {
                _coll7.push([_297_v154,_287_v148.f10]);
              }
            }
            return _coll7;
          }());
          let _298_v157;
          _298_v157 = _dafny.MultiSet.fromElements((_296_v156)[_module.__default.safeIndex((_222_v116).f3, new BigNumber((_296_v156).length))]);
          let _299_v158;
          _299_v158 = _module.D10.create_DC29(!(_72_v1), (_113_v38).f12);
          let _300_v159;
          _300_v159 = _dafny.Map.Empty.slice().updateUnsafe((_299_v158).dtor_cf43,_222_v116.f4);
          _251_cf71 = (_dafny.ZERO).minus((((_298_v157).contains(_300_v159)) ? ((_298_v157).get(_300_v159)) : (_107_v32)));
          (_287_v148).f10 = !(((_dafny.ZERO).minus(_module.__default.fm0(_71_globalState))).plus(new BigNumber(690))).isEqualTo((_222_v116).f3);
        } else if (_source4.is_DC37) {
          let _301___mcc_h13 = (_source4).cf64;
          let _302___mcc_h14 = (_source4).cf65;
          let _303___mcc_h15 = (_source4).cf66;
          let _304___mcc_h16 = (_source4).cf67;
          let _305_cf67 = _304___mcc_h16;
          let _306_cf66 = _303___mcc_h15;
          let _307_cf65 = _302___mcc_h14;
          let _308_cf64 = _301___mcc_h13;
          _106_v31 = _106_v31;
          (_71_globalState).f2 = _module.__default.fm2(_107_v32, _71_globalState);
          (_222_v116).f4 = _222_v116.f4;
          let _309_v160;
          _309_v160 = _dafny.Map.Empty.slice().updateUnsafe(_72_v1,_113_v38.f13);
          let _310_v161;
          _310_v161 = _module.D16.create_DC45(_309_v160);
          let _311_v162;
          let _nw49 = Array((new BigNumber(4)).toNumber());
          _nw49[(_dafny.ZERO).toNumber()] = _309_v160;
          _nw49[(_dafny.ONE).toNumber()] = (_309_v160).Merge(_309_v160);
          _nw49[(new BigNumber(2)).toNumber()] = (_309_v160).Merge((_dafny.Map.Empty.slice().updateUnsafe(_307_cf65,_111_v36)).update(false, _111_v36));
          _nw49[(new BigNumber(3)).toNumber()] = (_310_v161).dtor_cf78;
          _311_v162 = _nw49;
          let _index35 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_311_v162).length));
          (_311_v162)[_index35] = ((_module.D16.create_DC45(_dafny.Map.Empty.slice().updateUnsafe(!(_72_v1),_111_v36))).dtor_cf78).Merge(_309_v160);
          let _index36 = _module.__default.safeIndex(new BigNumber(349), new BigNumber((_311_v162).length));
          (_311_v162)[_index36] = _309_v160;
        } else {
          let _312___mcc_h17 = (_source4).cf55;
          let _313_cf55 = _312___mcc_h17;
          let _314_v163;
          let _nw50 = Array((new BigNumber(26)).toNumber());
          _314_v163 = _nw50;
          let _315_v164;
          _315_v164 = _dafny.Set.fromElements(_314_v163, _314_v163, _314_v163);
          _315_v164 = (_315_v164).Union(_dafny.Set.fromElements(_314_v163, _314_v163));
          (_71_globalState).f1 = new BigNumber((_113_v38.f13).length);
          let _316_v165;
          let _317_v166;
          let _318_v167;
          let _319_v168;
          let _out5;
          let _out6;
          let _out7;
          let _out8;
          let _outcollector1 = (_109_v34).m2((_106_v31)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length))], (_223_v117).f15, (_231_v124).Merge(_231_v124), _72_v1, _71_globalState);
          _out5 = _outcollector1[0];
          _out6 = _outcollector1[1];
          _out7 = _outcollector1[2];
          _out8 = _outcollector1[3];
          _316_v165 = _out5;
          _317_v166 = _out6;
          _318_v167 = _out7;
          _319_v168 = _out8;
          let _320_v169;
          let _out9;
          _out9 = (_109_v34).m1(new BigNumber(424), _71_globalState);
          _320_v169 = _out9;
        }
        let _321_v170;
        let _nw51 = new _module.C5();
        _nw51.__ctor((_109_v34).f9, (_223_v117).f14, (_223_v117).f14, _72_v1);
        _321_v170 = _nw51;
        if (_222_v116.f4) {
          let _322_v171;
          let _nw52 = Array((_dafny.ONE).toNumber()).fill([]);
          _322_v171 = _nw52;
          let _323_v172;
          let _nw53 = Array((new BigNumber(4)).toNumber()).fill(_module.D8.Default());
          _323_v172 = _nw53;
          let _index37 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_322_v171).length));
          (_322_v171)[_index37] = ((false) ? (_323_v172) : (_323_v172));
          let _index38 = _module.__default.safeIndex(new BigNumber(771), new BigNumber((_322_v171).length));
          (_322_v171)[_index38] = _323_v172;
          let _index39 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_112_v37).length));
          (_112_v37)[_index39] = new BigNumber(185);
          let _324_v173;
          let _nw54 = new _module.C0();
          _nw54.__ctor(_222_v116.f4, (_109_v34).f9);
          _324_v173 = _nw54;
          _324_v173 = _324_v173;
          let _325_v174;
          let _nw55 = new _module.C5();
          _nw55.__ctor((_109_v34).f9, _107_v32, _107_v32, true);
          _325_v174 = _nw55;
          let _326_v175;
          _326_v175 = _dafny.Seq.of(_325_v174, _325_v174);
          let _327_v176;
          _327_v176 = _dafny.Map.Empty.slice().updateUnsafe(_326_v175,(_254_v133).f8);
          _327_v176 = _327_v176;
          let _index40 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length));
          let _index41 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length));
          let _index42 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length));
          let _rhs40 = _325_v174.f4;
          let _rhs41 = (false) || ((_324_v173).fm5((_254_v133).f8, (_325_v174).f3, _110_v35, _72_v1, _71_globalState));
          let _rhs42 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_205_v103, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-229)), ((_328_v38) => function (_329_i23) {
            return new BigNumber((_328_v38.f13).length);
          })(_113_v38))), _205_v103)).length);
          let _rhs43 = _222_v116.f4;
          let _lhs25 = _106_v31;
          let _lhs26 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length));
          let _lhs27 = _106_v31;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length));
          let _lhs29 = _71_globalState;
          let _lhs30 = _106_v31;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length));
          _lhs25[_lhs26] = _rhs40;
          _lhs27[_lhs28] = _rhs41;
          _lhs29.f1 = _rhs42;
          _lhs30[_lhs31] = _rhs43;
        } else {
          _72_v1 = _222_v116.f4;
          let _330_v177;
          let _nw56 = new _module.C6();
          _nw56.__ctor((_dafny.ZERO).minus(new BigNumber((_111_v36).length)), (_223_v117).f14, _222_v116.f4, _251_cf71);
          _330_v177 = _nw56;
          let _331_v178;
          let _init6 = ((_332_v177) => function (_333_i24) {
            return _dafny.Seq.of((_332_v177).f8);
          })(_330_v177);
          let _nw57 = Array((new BigNumber(6)).toNumber());
          for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw57.length); _i0_6++) {
            _nw57[_i0_6] = _init6(new BigNumber(_i0_6));
          }
          _331_v178 = _nw57;
          let _nw58 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
          _331_v178 = _nw58;
          let _index43 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_112_v37).length));
          (_112_v37)[_index43] = (_222_v116).f3;
          let _334_v179;
          _334_v179 = _dafny.Map.Empty.slice().updateUnsafe((_109_v34).f9,_dafny.Set.fromElements(_106_v31));
          let _335_v180;
          _335_v180 = _dafny.Set.fromElements((_109_v34).f9, (_109_v34).f9);
          _334_v179 = (_334_v179).update((_109_v34).f9, _335_v180);
        }
      } else if (_source3.is_DC38) {
        let _336___mcc_h3 = (_source3).cf68;
        let _337_cf68 = _336___mcc_h3;
        let _index44 = _module.__default.safeIndex(new BigNumber(911), new BigNumber((_112_v37).length));
        (_112_v37)[_index44] = (_222_v116).f3;
        let _338_v181;
        _338_v181 = _dafny.Map.Empty.slice().updateUnsafe(_223_v117,new BigNumber(-124));
        _338_v181 = (_338_v181).update(_223_v117, (_112_v37)[_module.__default.safeIndex(new BigNumber(911), new BigNumber((_112_v37).length))]);
        let _339_v182;
        let _nw59 = new _module.C7();
        _nw59.__ctor(_222_v116.f4, _107_v32, (_223_v117).f14, _72_v1);
        _339_v182 = _nw59;
        let _340_v183;
        _340_v183 = _dafny.Map.Empty.slice().updateUnsafe(_339_v182,_115_v40);
        let _341_v184;
        _341_v184 = _dafny.MultiSet.fromElements(_dafny.Seq.of(_72_v1, _72_v1), (((_340_v183).contains(_339_v182)) ? ((_340_v183).get(_339_v182)) : (_dafny.Seq.update(_115_v40, _module.__default.safeIndex(new BigNumber((_111_v36).length), new BigNumber((_115_v40).length)), false))), _115_v40);
        _341_v184 = ((_341_v184).update(_115_v40, _module.__default.abs(new BigNumber((_168_v76).length)))).Union((_module.D17.create_DC48(_dafny.MultiSet.fromElements(_dafny.Seq.of(_222_v116.f4)))).dtor_cf83);
        let _index45 = _module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length));
        (_106_v31)[_index45] = !((_106_v31)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length))]) || (_339_v182.f5);
      } else {
        let _342___mcc_h4 = (_source3).cf72;
        let _343_cf72 = _342___mcc_h4;
        _112_v37 = (_223_v117).f15;
        (_71_globalState).f1 = (_222_v116).f3;
        (_71_globalState).f2 = _72_v1;
        let _344_v185;
        _344_v185 = _dafny.Set.fromElements(_69_v0);
        let _rhs44 = !((_344_v185).IsProperSubsetOf((((_106_v31)[_module.__default.safeIndex(new BigNumber(209), new BigNumber((_106_v31).length))]) ? (_344_v185) : (_344_v185))));
        let _rhs45 = _72_v1;
        let _rhs46 = _113_v38.f13;
        let _lhs32 = _71_globalState;
        let _lhs33 = _113_v38;
        _lhs32.f2 = _rhs44;
        _72_v1 = _rhs45;
        _lhs33.f13 = _rhs46;
      }
      process.stdout.write(_dafny.toString((_69_v0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_69_v0)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_71_globalState).f0)[new BigNumber(25)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_71_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_71_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_72_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_106_v31)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_107_v32));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_108_v33).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_109_v34).f9)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_110_v35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_111_v36).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_112_v37)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_113_v38).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write((_113_v38.f13).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_114_v39).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_115_v40, _dafny.Seq.of(false, false, false, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_116_v41).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_159_v70).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_160_v71).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_168_v76).equals(_dafny.Set.fromElements(new BigNumber(714)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_169_i12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_204_v102).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new _dafny.CodePoint('s'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_205_v103, _dafny.Seq.of(new BigNumber(714)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_206_i15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[_dafny.ZERO]).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_220_v114)[_dafny.ZERO].f13).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[_dafny.ZERO]).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v114)[_dafny.ZERO].f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[_dafny.ZERO]).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[_dafny.ONE]).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_220_v114)[_dafny.ONE].f13).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[_dafny.ONE]).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v114)[_dafny.ONE].f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[_dafny.ONE]).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[new BigNumber(2)]).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_220_v114)[new BigNumber(2)].f13).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[new BigNumber(2)]).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v114)[new BigNumber(2)].f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[new BigNumber(2)]).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[new BigNumber(3)]).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_220_v114)[new BigNumber(3)].f13).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[new BigNumber(3)]).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v114)[new BigNumber(3)].f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_220_v114)[new BigNumber(3)]).f7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_221_v115)[new BigNumber(3)]).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v116).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_222_v116.f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v117).f14));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_223_v117).f15)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_224_v118).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_225_v119).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_226_v120).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_227_i18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_231_v124).equals(_dafny.Map.Empty.slice().updateUnsafe(false,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_232_v125).dtor_cf68, _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(false,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false), _dafny.Map.Empty.slice().updateUnsafe(true,false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1, cf2, cf3, cf4) {
      let $dt = new D0(0);
      $dt.cf1 = cf1;
      $dt.cf2 = cf2;
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC2(cf5, cf6) {
      let $dt = new D0(1);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      return $dt;
    }
    static create_DC3() {
      let $dt = new D0(2);
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(3);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC4(cf7) {
      let $dt = new D0(4);
      $dt.cf7 = cf7;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC3() { return this.$tag === 2; }
    get is_DC0() { return this.$tag === 3; }
    get is_DC4() { return this.$tag === 4; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf7() { return this.cf7; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC3";
      } else if (this.$tag === 3) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 4) {
        return "D0.DC4" + "(" + _dafny.toString(this.cf7) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf1 === other.cf1 && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5) && this.cf6 === other.cf6;
      } else if (this.$tag === 2) {
        return other.$tag === 2;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf0 === other.cf0;
      } else if (this.$tag === 4) {
        return other.$tag === 4 && _dafny.areEqual(this.cf7, other.cf7);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(false, false, new _dafny.CodePoint('D'.codePointAt(0)), []);
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
    static create_DC6(cf9, cf10, cf11, cf12) {
      let $dt = new D1(0);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      $dt.cf12 = cf12;
      return $dt;
    }
    static create_DC7() {
      let $dt = new D1(1);
      return $dt;
    }
    static create_DC5(cf8) {
      let $dt = new D1(2);
      $dt.cf8 = cf8;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC5() { return this.$tag === 2; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf12() { return this.cf12; }
    get dtor_cf8() { return this.cf8; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC6" + "(" + _dafny.toString(this.cf9) + ", " + this.cf10.toVerbatimString(true) + ", " + _dafny.toString(this.cf11) + ", " + _dafny.toString(this.cf12) + ")";
      } else if (this.$tag === 1) {
        return "D1.DC7";
      } else if (this.$tag === 2) {
        return "D1.DC5" + "(" + _dafny.toString(this.cf8) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf9 === other.cf9 && _dafny.areEqual(this.cf10, other.cf10) && this.cf11 === other.cf11 && this.cf12 === other.cf12;
      } else if (this.$tag === 1) {
        return other.$tag === 1;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf8, other.cf8);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC6(null, _dafny.Seq.UnicodeFromString(""), false, false);
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
    static create_DC8(cf13) {
      let $dt = new D2(0);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC8" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf13 === other.cf13;
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
          return D2.Default();
        }
      };
    }
  }

  $module.D3 = class D3 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC10(cf15, cf16, cf17) {
      let $dt = new D3(0);
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      $dt.cf17 = cf17;
      return $dt;
    }
    static create_DC9(cf14) {
      let $dt = new D3(1);
      $dt.cf14 = cf14;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf14() { return this.cf14; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC10" + "(" + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ", " + _dafny.toString(this.cf17) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf14) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf15, other.cf15) && this.cf16 === other.cf16 && _dafny.areEqual(this.cf17, other.cf17);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf14 === other.cf14;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC10(_dafny.ZERO, [], _dafny.ZERO);
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
    static create_DC12() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC11(cf18) {
      let $dt = new D4(1);
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC13(cf19) {
      let $dt = new D4(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get is_DC13() { return this.$tag === 2; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC12";
      } else if (this.$tag === 1) {
        return "D4.DC11" + "(" + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 2) {
        return "D4.DC13" + "(" + _dafny.toString(this.cf19) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC12();
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
    static create_DC15(cf21) {
      let $dt = new D5(0);
      $dt.cf21 = cf21;
      return $dt;
    }
    static create_DC14(cf20) {
      let $dt = new D5(1);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC16(cf22) {
      let $dt = new D5(2);
      $dt.cf22 = cf22;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get is_DC14() { return this.$tag === 1; }
    get is_DC16() { return this.$tag === 2; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf22() { return this.cf22; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC15" + "(" + _dafny.toString(this.cf21) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC16" + "(" + _dafny.toString(this.cf22) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf21, other.cf21);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf22, other.cf22);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC15(_module.D0.Default());
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
    static create_DC18(cf24, cf25) {
      let $dt = new D6(0);
      $dt.cf24 = cf24;
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC19(cf26) {
      let $dt = new D6(1);
      $dt.cf26 = cf26;
      return $dt;
    }
    static create_DC17(cf23) {
      let $dt = new D6(2);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf26() { return this.cf26; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC18" + "(" + _dafny.toString(this.cf24) + ", " + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC19" + "(" + _dafny.toString(this.cf26) + ")";
      } else if (this.$tag === 2) {
        return "D6.DC17" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf24 === other.cf24 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf26, other.cf26);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf23, other.cf23);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC18(false, _dafny.ZERO);
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
    static create_DC21(cf28, cf29, cf30) {
      let $dt = new D7(0);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      return $dt;
    }
    static create_DC20(cf27) {
      let $dt = new D7(1);
      $dt.cf27 = cf27;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf27() { return this.cf27; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC21" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC20" + "(" + _dafny.toString(this.cf27) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf28 === other.cf28 && this.cf29 === other.cf29 && this.cf30 === other.cf30;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf27, other.cf27);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC21([], false, false);
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
    static create_DC23(cf32, cf33, cf34, cf35) {
      let $dt = new D8(0);
      $dt.cf32 = cf32;
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC22(cf31) {
      let $dt = new D8(1);
      $dt.cf31 = cf31;
      return $dt;
    }
    static create_DC24(cf36) {
      let $dt = new D8(2);
      $dt.cf36 = cf36;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC22() { return this.$tag === 1; }
    get is_DC24() { return this.$tag === 2; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf36() { return this.cf36; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC23" + "(" + _dafny.toString(this.cf32) + ", " + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC22" + "(" + _dafny.toString(this.cf31) + ")";
      } else if (this.$tag === 2) {
        return "D8.DC24" + "(" + _dafny.toString(this.cf36) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf32 === other.cf32 && this.cf33 === other.cf33 && this.cf34 === other.cf34 && _dafny.areEqual(this.cf35, other.cf35);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf31, other.cf31);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf36, other.cf36);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC23(false, false, false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC26(cf38, cf39) {
      let $dt = new D9(0);
      $dt.cf38 = cf38;
      $dt.cf39 = cf39;
      return $dt;
    }
    static create_DC25(cf37) {
      let $dt = new D9(1);
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC27(cf40) {
      let $dt = new D9(2);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC26() { return this.$tag === 0; }
    get is_DC25() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf39() { return this.cf39; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC26" + "(" + _dafny.toString(this.cf38) + ", " + _dafny.toString(this.cf39) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC25" + "(" + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D9.DC27" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf38, other.cf38) && this.cf39 === other.cf39;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf40, other.cf40);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC26(_dafny.ZERO, false);
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
    static create_DC29(cf42, cf43) {
      let $dt = new D10(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      return $dt;
    }
    static create_DC30(cf44, cf45, cf46, cf47, cf48) {
      let $dt = new D10(1);
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC28(cf41) {
      let $dt = new D10(2);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC30() { return this.$tag === 1; }
    get is_DC28() { return this.$tag === 2; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC29" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC30" + "(" + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + _dafny.toString(this.cf48) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC28" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf42 === other.cf42 && _dafny.areEqual(this.cf43, other.cf43);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf44, other.cf44) && this.cf45 === other.cf45 && this.cf46 === other.cf46 && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC29(false, new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC32(cf50, cf51) {
      let $dt = new D11(0);
      $dt.cf50 = cf50;
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC33(cf52, cf53, cf54) {
      let $dt = new D11(1);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      $dt.cf54 = cf54;
      return $dt;
    }
    static create_DC31(cf49) {
      let $dt = new D11(2);
      $dt.cf49 = cf49;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf54() { return this.cf54; }
    get dtor_cf49() { return this.cf49; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC32" + "(" + _dafny.toString(this.cf50) + ", " + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC33" + "(" + _dafny.toString(this.cf52) + ", " + _dafny.toString(this.cf53) + ", " + _dafny.toString(this.cf54) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC31" + "(" + _dafny.toString(this.cf49) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf50 === other.cf50 && _dafny.areEqual(this.cf51, other.cf51);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf52 === other.cf52 && _dafny.areEqual(this.cf53, other.cf53) && _dafny.areEqual(this.cf54, other.cf54);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf49, other.cf49);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC32(false, _dafny.ZERO);
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
    static create_DC35(cf56, cf57, cf58, cf59, cf60) {
      let $dt = new D12(0);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC36(cf61, cf62, cf63) {
      let $dt = new D12(1);
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      return $dt;
    }
    static create_DC37(cf64, cf65, cf66, cf67) {
      let $dt = new D12(2);
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      $dt.cf66 = cf66;
      $dt.cf67 = cf67;
      return $dt;
    }
    static create_DC34(cf55) {
      let $dt = new D12(3);
      $dt.cf55 = cf55;
      return $dt;
    }
    get is_DC35() { return this.$tag === 0; }
    get is_DC36() { return this.$tag === 1; }
    get is_DC37() { return this.$tag === 2; }
    get is_DC34() { return this.$tag === 3; }
    get dtor_cf56() { return this.cf56; }
    get dtor_cf57() { return this.cf57; }
    get dtor_cf58() { return this.cf58; }
    get dtor_cf59() { return this.cf59; }
    get dtor_cf60() { return this.cf60; }
    get dtor_cf61() { return this.cf61; }
    get dtor_cf62() { return this.cf62; }
    get dtor_cf63() { return this.cf63; }
    get dtor_cf64() { return this.cf64; }
    get dtor_cf65() { return this.cf65; }
    get dtor_cf66() { return this.cf66; }
    get dtor_cf67() { return this.cf67; }
    get dtor_cf55() { return this.cf55; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC35" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + this.cf59.toVerbatimString(true) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC36" + "(" + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC37" + "(" + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ", " + _dafny.toString(this.cf66) + ", " + _dafny.toString(this.cf67) + ")";
      } else if (this.$tag === 3) {
        return "D12.DC34" + "(" + _dafny.toString(this.cf55) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf56 === other.cf56 && this.cf57 === other.cf57 && _dafny.areEqual(this.cf58, other.cf58) && _dafny.areEqual(this.cf59, other.cf59) && _dafny.areEqual(this.cf60, other.cf60);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf61 === other.cf61 && this.cf62 === other.cf62 && _dafny.areEqual(this.cf63, other.cf63);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf64, other.cf64) && this.cf65 === other.cf65 && _dafny.areEqual(this.cf66, other.cf66) && _dafny.areEqual(this.cf67, other.cf67);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && this.cf55 === other.cf55;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC35([], false, _dafny.Seq.of(), _dafny.Seq.UnicodeFromString(""), _dafny.ZERO);
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
    static create_DC39(cf69) {
      let $dt = new D13(0);
      $dt.cf69 = cf69;
      return $dt;
    }
    static create_DC40(cf70, cf71) {
      let $dt = new D13(1);
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC38(cf68) {
      let $dt = new D13(2);
      $dt.cf68 = cf68;
      return $dt;
    }
    static create_DC41(cf72) {
      let $dt = new D13(3);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC38() { return this.$tag === 2; }
    get is_DC41() { return this.$tag === 3; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC39" + "(" + _dafny.toString(this.cf69) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC40" + "(" + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC38" + "(" + _dafny.toString(this.cf68) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC41" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf69, other.cf69);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf70, other.cf70) && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf68, other.cf68);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf72, other.cf72);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC39(_dafny.Map.Empty);
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
    static create_DC42(cf73) {
      let $dt = new D14(0);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC42() { return this.$tag === 0; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC42" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.of();
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
    static create_DC44(cf75, cf76, cf77) {
      let $dt = new D15(0);
      $dt.cf75 = cf75;
      $dt.cf76 = cf76;
      $dt.cf77 = cf77;
      return $dt;
    }
    static create_DC43(cf74) {
      let $dt = new D15(1);
      $dt.cf74 = cf74;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf74() { return this.cf74; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC44" + "(" + _dafny.toString(this.cf75) + ", " + _dafny.toString(this.cf76) + ", " + _dafny.toString(this.cf77) + ")";
      } else if (this.$tag === 1) {
        return "D15.DC43" + "(" + _dafny.toString(this.cf74) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf75 === other.cf75 && _dafny.areEqual(this.cf76, other.cf76) && _dafny.areEqual(this.cf77, other.cf77);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf74, other.cf74);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D15.create_DC44(null, new _dafny.CodePoint('D'.codePointAt(0)), _dafny.ZERO);
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
    static create_DC46(cf79, cf80, cf81) {
      let $dt = new D16(0);
      $dt.cf79 = cf79;
      $dt.cf80 = cf80;
      $dt.cf81 = cf81;
      return $dt;
    }
    static create_DC45(cf78) {
      let $dt = new D16(1);
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC47(cf82) {
      let $dt = new D16(2);
      $dt.cf82 = cf82;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get is_DC47() { return this.$tag === 2; }
    get dtor_cf79() { return this.cf79; }
    get dtor_cf80() { return this.cf80; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf82() { return this.cf82; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC46" + "(" + this.cf79.toVerbatimString(true) + ", " + _dafny.toString(this.cf80) + ", " + _dafny.toString(this.cf81) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC45" + "(" + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC47" + "(" + _dafny.toString(this.cf82) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf79, other.cf79) && _dafny.areEqual(this.cf80, other.cf80) && this.cf81 === other.cf81;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf78, other.cf78);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf82, other.cf82);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC46(_dafny.Seq.UnicodeFromString(""), _dafny.Set.Empty, false);
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
    static create_DC49(cf84, cf85, cf86) {
      let $dt = new D17(0);
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      $dt.cf86 = cf86;
      return $dt;
    }
    static create_DC50(cf87, cf88, cf89, cf90, cf91) {
      let $dt = new D17(1);
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      return $dt;
    }
    static create_DC48(cf83) {
      let $dt = new D17(2);
      $dt.cf83 = cf83;
      return $dt;
    }
    get is_DC49() { return this.$tag === 0; }
    get is_DC50() { return this.$tag === 1; }
    get is_DC48() { return this.$tag === 2; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf86() { return this.cf86; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf83() { return this.cf83; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC49" + "(" + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ", " + _dafny.toString(this.cf86) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC50" + "(" + _dafny.toString(this.cf87) + ", " + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC48" + "(" + _dafny.toString(this.cf83) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf84, other.cf84) && _dafny.areEqual(this.cf85, other.cf85) && this.cf86 === other.cf86;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf87 === other.cf87 && this.cf88 === other.cf88 && this.cf89 === other.cf89 && _dafny.areEqual(this.cf90, other.cf90) && this.cf91 === other.cf91;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf83, other.cf83);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC49(_dafny.ZERO, _dafny.ZERO, null);
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
    static create_DC51(cf92) {
      let $dt = new D18(0);
      $dt.cf92 = cf92;
      return $dt;
    }
    get is_DC51() { return this.$tag === 0; }
    get dtor_cf92() { return this.cf92; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC51" + "(" + _dafny.toString(this.cf92) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf92, other.cf92);
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
          return D18.Default();
        }
      };
    }
  }

  $module.D19 = class D19 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC52(cf93) {
      let $dt = new D19(0);
      $dt.cf93 = cf93;
      return $dt;
    }
    get is_DC52() { return this.$tag === 0; }
    get dtor_cf93() { return this.cf93; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC52" + "(" + _dafny.toString(this.cf93) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf93, other.cf93);
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
          return D19.Default();
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
      this.f1 = _dafny.ZERO;
      this.f2 = false;
      this._f0 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2) {
      let _this = this;
      (_this)._f0 = f0;
      (_this).f1 = f1;
      (_this).f2 = f2;
      return;
    }
    get f0() {
      let _this = this;
      return _this._f0;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this.f10 = false;
      this._f11 = [];
    }
    _parentTraits() {
      return [];
    }
    __ctor(f10, f11) {
      let _this = this;
      (_this).f10 = f10;
      (_this)._f11 = f11;
      return;
    }
    fm5(p0, p1, p2, p3, globalState) {
      let _this = this;
      return !(new BigNumber((_dafny.Seq.UnicodeFromString("cdnmimuc")).length)).isEqualTo(new BigNumber((_dafny.Set.fromElements(!(_this.f10))).length));
    };
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f4 = false;
      this._f7 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f7, f3, f4) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      if (!(_this.f4)) {
        let _345_v0;
        let _nw60 = Array((new BigNumber(19)).toNumber()).fill(false);
        _345_v0 = _nw60;
        let _346_v1;
        let _nw61 = new _module.C0();
        _nw61.__ctor(!(_module.__default.fm2((_this).f7, globalState)), _345_v0);
        _346_v1 = _nw61;
        let _347_v2;
        let _nw62 = new _module.C0();
        _nw62.__ctor(_346_v1.f10, (_346_v1).f11);
        _347_v2 = _nw62;
        let _index46 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_345_v0).length));
        (_345_v0)[_index46] = ((_this.f4) ? (_347_v2.f10) : (_this.f4));
        let _348_v3;
        _348_v3 = _dafny.Seq.of(_346_v1.f10, _346_v1.f10);
        let _349_v4;
        _349_v4 = _dafny.Seq.of(new BigNumber((_module.__default.fm19(p0, globalState)).length));
        let _350_v5;
        _350_v5 = _dafny.Seq.UnicodeFromString("hxpnn");
        let _351_v6;
        _351_v6 = _module.D3.create_DC10(new BigNumber((_348_v3).length), (_346_v1).f11, (_349_v4)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_349_v4, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_350_v5).length)), new BigNumber((_349_v4).length)), p0)).length), new BigNumber((_349_v4).length))]);
        let _index47 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_345_v0).length));
        let _rhs47 = (_351_v6).dtor_cf17;
        let _rhs48 = !(_this.f4);
        let _lhs34 = globalState;
        let _lhs35 = _345_v0;
        let _lhs36 = _module.__default.safeIndex(new BigNumber(845), new BigNumber((_345_v0).length));
        _lhs34.f1 = _rhs47;
        _lhs35[_lhs36] = _rhs48;
        let _352_v7;
        _352_v7 = new _dafny.CodePoint('j'.codePointAt(0));
        _352_v7 = new _dafny.CodePoint('c'.codePointAt(0));
        (globalState).f1 = (_this).f7;
      } else {
        let _353_v8;
        _353_v8 = _dafny.Seq.UnicodeFromString("xpxqoyf");
        let _354_v9;
        let _init7 = function (_355_i0) {
          return _this.f4;
        };
        let _nw63 = Array((new BigNumber(20)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw63.length); _i0_7++) {
          _nw63[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _354_v9 = _nw63;
        let _356_v10;
        _356_v10 = _module.D3.create_DC10(new BigNumber((_353_v8).length), _354_v9, (_this).f7);
        _353_v8 = _module.__default.fm20(((_356_v10).dtor_cf15).multipliedBy((_this).f3), globalState);
        (globalState).f1 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_this).f3), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,_this.f4)).length));
        if (_dafny.areEqual(_dafny.Seq.UnicodeFromString("ukt"), _module.__default.fm20((_dafny.ZERO).minus(new BigNumber((_353_v8).length)), globalState))) {
          (globalState).f1 = (_this).f3;
          let _357_v11;
          _357_v11 = _dafny.MultiSet.fromElements(_this.f4);
          _357_v11 = _357_v11;
          (globalState).f2 = ((_this.f4) ? (((_dafny.ZERO).minus((_this).f7)).isLessThan(p0)) : (false));
          (globalState).f1 = _module.__default.safeModuloInt(new BigNumber(791), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f3,_this.f4)).length));
          _356_v10 = _356_v10;
        } else {
          (globalState).f2 = _module.__default.fm2(_module.__default.safeModuloInt(p0, (_this).f3), globalState);
          let _358_v12;
          _358_v12 = _module.D0.create_DC0(_this.f4);
          let _359_v13;
          _359_v13 = _dafny.MultiSet.fromElements(_this.f4, _this.f4, _this.f4, (_358_v12).dtor_cf0, _this.f4);
          let _360_v14;
          _360_v14 = _dafny.Map.Empty.slice().updateUnsafe(_353_v8,_359_v13);
          _360_v14 = (_360_v14).update(_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-476)), function (_361_i1) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          }), _module.__default.safeIndex((_this).f7, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-476)), function (_362_i1) {
            return new _dafny.CodePoint('y'.codePointAt(0));
          })).length)), new _dafny.CodePoint('q'.codePointAt(0))), _359_v13);
          (globalState).f1 = (_dafny.ZERO).minus((_this).f7);
          let _363_v15;
          let _nw64 = Array((new BigNumber(11)).toNumber());
          _nw64[(_dafny.ZERO).toNumber()] = p0;
          _nw64[(_dafny.ONE).toNumber()] = (_this).f3;
          _nw64[(new BigNumber(2)).toNumber()] = (_this).f7;
          _nw64[(new BigNumber(3)).toNumber()] = _module.__default.fm0(globalState);
          _nw64[(new BigNumber(4)).toNumber()] = (_this).f3;
          _nw64[(new BigNumber(5)).toNumber()] = new BigNumber((_353_v8).length);
          _nw64[(new BigNumber(6)).toNumber()] = p0;
          _nw64[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_this).f7);
          _nw64[(new BigNumber(8)).toNumber()] = (_this).f7;
          _nw64[(new BigNumber(9)).toNumber()] = p0;
          _nw64[(new BigNumber(10)).toNumber()] = new BigNumber(-233);
          _363_v15 = _nw64;
          let _index48 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_363_v15).length));
          (_363_v15)[_index48] = ((_dafny.ZERO).minus((_this).f7)).multipliedBy(p0);
          let _index49 = _module.__default.safeIndex(new BigNumber(868), new BigNumber((_363_v15).length));
          (_363_v15)[_index49] = _module.__default.fm0(globalState);
          (globalState).f2 = true;
        }
        let _364_v16;
        _364_v16 = _dafny.Set.fromElements(p0);
        r0 = new BigNumber((_364_v16).length);
        _354_v9 = _354_v9;
      }
      let _hi3 = _module.__default.safeModuloInt(new BigNumber(338), (_this).f3);
      for (let _365_i2 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(40)), function (_378_i3) {
        return (_this).f3;
      })).length); _365_i2.isLessThan(_hi3); _365_i2 = _365_i2.plus(_dafny.ONE)) {
        let _366_v17;
        let _init8 = function (_367_i4) {
          return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("kjyvte"), _dafny.Seq.UnicodeFromString("ydnlt"));
        };
        let _nw65 = Array((new BigNumber(20)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw65.length); _i0_8++) {
          _nw65[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _366_v17 = _nw65;
        let _368_v18;
        _368_v18 = _dafny.Seq.UnicodeFromString("wjj");
        let _index50 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_366_v17).length));
        (_366_v17)[_index50] = _368_v18;
        let _index51 = _module.__default.safeIndex(new BigNumber(329), new BigNumber((_366_v17).length));
        (_366_v17)[_index51] = _dafny.Seq.UnicodeFromString("hyb");
        let _369_v19;
        let _init9 = function (_370_i5) {
          return (_370_i5).multipliedBy(new BigNumber((_dafny.Seq.of(_this.f4)).length));
        };
        let _nw66 = Array((new BigNumber(20)).toNumber());
        for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw66.length); _i0_9++) {
          _nw66[_i0_9] = _init9(new BigNumber(_i0_9));
        }
        _369_v19 = _nw66;
        let _index52 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_369_v19).length));
        (_369_v19)[_index52] = _365_i2;
        let _index53 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_369_v19).length));
        (_369_v19)[_index53] = (_this).f3;
        r0 = _module.__default.fm0(globalState);
        let _371_v20;
        let _nw67 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _371_v20 = _nw67;
        let _372_v21;
        _372_v21 = new _dafny.CodePoint('u'.codePointAt(0));
        let _index54 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_371_v20).length));
        (_371_v20)[_index54] = _372_v21;
        let _373_v22;
        _373_v22 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_this.f4);
        let _374_v23;
        _374_v23 = _dafny.Seq.of(_373_v22);
        let _375_v24;
        _375_v24 = _dafny.Seq.of(_module.__default.fm21(_374_v23, globalState));
        let _376_v25;
        _376_v25 = _dafny.Seq.of(_this.f4);
        let _377_v26;
        _377_v26 = _dafny.Seq.of(new BigNumber((_376_v25).length));
        let _index55 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_371_v20).length));
        let _index56 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_369_v19).length));
        let _rhs49 = _this.f4;
        let _rhs50 = _372_v21;
        let _rhs51 = new BigNumber(((_375_v24)[_module.__default.safeIndex(_module.__default.safeDivisionInt((_this).f7, new BigNumber((_377_v26).length)), new BigNumber((_375_v24).length))]).length);
        let _lhs37 = globalState;
        let _lhs38 = _371_v20;
        let _lhs39 = _module.__default.safeIndex(new BigNumber(584), new BigNumber((_371_v20).length));
        let _lhs40 = _369_v19;
        let _lhs41 = _module.__default.safeIndex(new BigNumber(812), new BigNumber((_369_v19).length));
        _lhs37.f2 = _rhs49;
        _lhs38[_lhs39] = _rhs50;
        _lhs40[_lhs41] = _rhs51;
      }
      let _379_v27;
      _379_v27 = _dafny.Seq.UnicodeFromString("s");
      _379_v27 = _dafny.Seq.Concat(_379_v27, _dafny.Seq.Concat(_379_v27, _dafny.Seq.UnicodeFromString("asujtxoi")));
      let _380_v28;
      _380_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_379_v27).length),new BigNumber(235));
      let _381_v29;
      _381_v29 = _dafny.Seq.of(_380_v28, _dafny.Map.Empty.slice().updateUnsafe((_this).f3,new BigNumber((_379_v27).length)), _380_v28);
      let _hi4 = ((_this).f3).plus(new BigNumber((_381_v29).length));
      for (let _382_i6 = _module.__default.safeDivisionInt((_this).f7, new BigNumber(357)); _382_i6.isLessThan(_hi4); _382_i6 = _382_i6.plus(_dafny.ONE)) {
        let _383_v30;
        _383_v30 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_dafny.ZERO).minus(p0)));
        (_this).f4 = (((_383_v30).IsSubsetOf(_383_v30)) ? (_this.f4) : (((_this).f7).isLessThan(_module.__default.fm0(globalState))));
        if (_module.__default.fm2((_this).f3, globalState)) {
          let _384_v31;
          let _init10 = function (_385_i7) {
            return _this.f4;
          };
          let _nw68 = Array((new BigNumber(27)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw68.length); _i0_10++) {
            _nw68[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _384_v31 = _nw68;
          let _386_v32;
          let _nw69 = Array((new BigNumber(11)).toNumber());
          _nw69[(_dafny.ZERO).toNumber()] = _384_v31;
          _nw69[(_dafny.ONE).toNumber()] = _384_v31;
          _nw69[(new BigNumber(2)).toNumber()] = _384_v31;
          _nw69[(new BigNumber(3)).toNumber()] = _384_v31;
          _nw69[(new BigNumber(4)).toNumber()] = _384_v31;
          _nw69[(new BigNumber(5)).toNumber()] = _384_v31;
          _nw69[(new BigNumber(6)).toNumber()] = _384_v31;
          _nw69[(new BigNumber(7)).toNumber()] = ((_this.f4) ? (_384_v31) : (_384_v31));
          _nw69[(new BigNumber(8)).toNumber()] = _384_v31;
          _nw69[(new BigNumber(9)).toNumber()] = _384_v31;
          _nw69[(new BigNumber(10)).toNumber()] = _384_v31;
          _386_v32 = _nw69;
          let _387_v33;
          let _nw70 = Array((new BigNumber(22)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _387_v33 = _nw70;
          let _388_v34;
          _388_v34 = new _dafny.CodePoint('r'.codePointAt(0));
          let _index57 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_387_v33).length));
          (_387_v33)[_index57] = _388_v34;
          let _index58 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_387_v33).length));
          let _rhs52 = _386_v32;
          let _rhs53 = _module.__default.safeModuloInt(_module.__default.fm0(globalState), _382_i6);
          let _rhs54 = _this.f4;
          let _rhs55 = !(_this.f4);
          let _rhs56 = _388_v34;
          let _lhs42 = globalState;
          let _lhs43 = _this;
          let _lhs44 = _this;
          let _lhs45 = _387_v33;
          let _lhs46 = _module.__default.safeIndex(new BigNumber(463), new BigNumber((_387_v33).length));
          _386_v32 = _rhs52;
          _lhs42.f1 = _rhs53;
          _lhs43.f4 = _rhs54;
          _lhs44.f4 = _rhs55;
          _lhs45[_lhs46] = _rhs56;
          let _389_v35;
          _389_v35 = _dafny.Seq.of(_379_v27);
          let _390_v36;
          _390_v36 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_389_v35);
          _390_v36 = (_390_v36).update((_dafny.ZERO).minus(p0), _dafny.Seq.Concat(_389_v35, _389_v35));
          let _391_v37;
          let _init11 = ((_392_i6) => function (_393_i8) {
            return (_393_i8).minus(_392_i6);
          })(_382_i6);
          let _nw71 = Array((new BigNumber(26)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw71.length); _i0_11++) {
            _nw71[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _391_v37 = _nw71;
          _391_v37 = _391_v37;
          (globalState).f1 = (_this).f3;
          let _394_v38;
          _394_v38 = _dafny.Seq.of(_this.f4, _this.f4, _this.f4);
          let _395_v39;
          _395_v39 = _dafny.MultiSet.fromElements((_387_v33)[_module.__default.safeIndex(new BigNumber(463), new BigNumber((_387_v33).length))]);
          let _396_v40;
          _396_v40 = _module.D0.create_DC2(_395_v39, _this.f4);
          let _397_v41;
          _397_v41 = _module.D0.create_DC2((_396_v40).dtor_cf5, _this.f4);
          let _398_v42;
          let _nw72 = new _module.C0();
          _nw72.__ctor(!((_396_v40).dtor_cf6), _384_v31);
          _398_v42 = _nw72;
          let _399_v43;
          _399_v43 = _module.D6.create_DC17(_394_v38);
          let _rhs57 = (_this).f3;
          let _rhs58 = _dafny.Seq.Concat((_399_v43).dtor_cf23, _394_v38);
          let _rhs59 = _this.f4;
          let _rhs60 = p0;
          let _rhs61 = _398_v42;
          let _lhs47 = globalState;
          let _lhs48 = _this;
          let _lhs49 = globalState;
          _lhs47.f1 = _rhs57;
          _394_v38 = _rhs58;
          _lhs48.f4 = _rhs59;
          _lhs49.f1 = _rhs60;
          _398_v42 = _rhs61;
        } else {
          let _400_v44;
          let _nw73 = Array((new BigNumber(14)).toNumber()).fill([]);
          _400_v44 = _nw73;
          _400_v44 = _400_v44;
          (_this).f4 = _dafny.Seq.IsPrefixOf(_379_v27, _379_v27);
          let _401_v45;
          _401_v45 = new _dafny.CodePoint('o'.codePointAt(0));
          _401_v45 = _401_v45;
          (globalState).f2 = _module.__default.fm2(new BigNumber(-802), globalState);
          let _402_v46;
          let _nw74 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _402_v46 = _nw74;
          let _index59 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_402_v46).length));
          (_402_v46)[_index59] = _dafny.Seq.Concat(_379_v27, _dafny.Seq.Create(_module.__default.abs(new BigNumber(940)), ((_403_v45) => function (_404_i9) {
            return _403_v45;
          })(_401_v45)));
          let _405_v47;
          _405_v47 = _module.D1.create_DC5(_382_i6);
          let _406_v48;
          _406_v48 = _dafny.Map.Empty.slice().updateUnsafe(_405_v47,new BigNumber(220));
          let _407_v49;
          _407_v49 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_this.f4);
          let _408_v50;
          _408_v50 = _dafny.Map.Empty.slice().updateUnsafe(_407_v49,_382_i6);
          let _index60 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_402_v46).length));
          let _rhs62 = (((_406_v48).contains(_405_v47)) ? ((_406_v48).get(_405_v47)) : (new BigNumber((_408_v50).length)));
          let _rhs63 = _module.__default.fm20((_this).f3, globalState);
          let _lhs50 = globalState;
          let _lhs51 = _402_v46;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_402_v46).length));
          _lhs50.f1 = _rhs62;
          _lhs51[_lhs52] = _rhs63;
        }
        (globalState).f2 = _this.f4;
        (globalState).f2 = true;
      }
      let _409_i10;
      _409_i10 = _dafny.ZERO;
      L3: {
        while (!(_this.f4) || (_this.f4)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_409_i10)) {
              break L3;
            }
            _409_i10 = (_409_i10).plus(_dafny.ONE);
            let _410_v51;
            let _nw75 = Array((new BigNumber(29)).toNumber());
            _nw75[(_dafny.ZERO).toNumber()] = _this.f4;
            _nw75[(_dafny.ONE).toNumber()] = _this.f4;
            _nw75[(new BigNumber(2)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(3)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(4)).toNumber()] = true;
            _nw75[(new BigNumber(5)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(6)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(7)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(8)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(9)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(10)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(11)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(12)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(13)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(14)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(15)).toNumber()] = true;
            _nw75[(new BigNumber(16)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(17)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(18)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(19)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(20)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(21)).toNumber()] = true;
            _nw75[(new BigNumber(22)).toNumber()] = _module.__default.fm2((_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_379_v27).length))), globalState);
            _nw75[(new BigNumber(23)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(24)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(25)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(26)).toNumber()] = _module.__default.fm2(p0, globalState);
            _nw75[(new BigNumber(27)).toNumber()] = _this.f4;
            _nw75[(new BigNumber(28)).toNumber()] = true;
            _410_v51 = _nw75;
            let _411_v52;
            let _nw76 = new _module.C0();
            _nw76.__ctor(_this.f4, _410_v51);
            _411_v52 = _nw76;
            let _412_v53;
            _412_v53 = _dafny.Set.fromElements(p0, (_this).f7, (_this).f3);
            let _413_v54;
            _413_v54 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f7).isLessThan(p0),_379_v27);
            let _414_v55;
            _414_v55 = new _dafny.CodePoint('f'.codePointAt(0));
            let _rhs64 = (((_413_v54).contains(((_module.__default.fm2((_this).f7, globalState)) ? ((_411_v52).fm5(p0, (_this).f7, _414_v55, _411_v52.f10, globalState)) : (_this.f4)))) ? ((_413_v54).get(((_module.__default.fm2((_this).f7, globalState)) ? ((_411_v52).fm5(p0, (_this).f7, _414_v55, _411_v52.f10, globalState)) : (_this.f4)))) : (_379_v27));
            let _rhs65 = _412_v53;
            let _rhs66 = ((_381_v29)[_module.__default.safeIndex(p0, new BigNumber((_381_v29).length))]).Merge(_380_v28);
            let _rhs67 = ((_this.f4) ? (_410_v51) : (_410_v51));
            _379_v27 = _rhs64;
            _412_v53 = _rhs65;
            _380_v28 = _rhs66;
            _410_v51 = _rhs67;
            let _415_v56;
            _415_v56 = _dafny.Seq.of(_411_v52.f10, _this.f4);
            let _416_v57;
            let _nw77 = Array((new BigNumber(4)).toNumber());
            _nw77[(_dafny.ZERO).toNumber()] = _415_v56;
            _nw77[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_415_v56, _dafny.Seq.of(_411_v52.f10));
            _nw77[(new BigNumber(2)).toNumber()] = _415_v56;
            _nw77[(new BigNumber(3)).toNumber()] = _415_v56;
            _416_v57 = _nw77;
            let _index61 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_416_v57).length));
            (_416_v57)[_index61] = _415_v56;
            let _417_v58;
            let _nw78 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
            _417_v58 = _nw78;
            let _418_v59;
            _418_v59 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("vfxj"));
            let _index62 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_417_v58).length));
            (_417_v58)[_index62] = ((_this).f3).multipliedBy(new BigNumber((_418_v59).length));
            let _419_v60;
            _419_v60 = _dafny.Seq.of(_412_v53, _412_v53);
            let _420_v61;
            _420_v61 = _dafny.Map.Empty.slice().updateUnsafe(_410_v51,_414_v55);
            let _index63 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_416_v57).length));
            let _index64 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_417_v58).length));
            let _rhs68 = _415_v56;
            let _rhs69 = new BigNumber((_dafny.Seq.Concat(_419_v60, _419_v60)).length);
            let _rhs70 = new BigNumber((_420_v61).length);
            let _rhs71 = (_dafny.Set.fromElements(p0, p0)).Intersect(_412_v53);
            let _lhs53 = _416_v57;
            let _lhs54 = _module.__default.safeIndex(new BigNumber(598), new BigNumber((_416_v57).length));
            let _lhs55 = globalState;
            let _lhs56 = _417_v58;
            let _lhs57 = _module.__default.safeIndex(new BigNumber(910), new BigNumber((_417_v58).length));
            _lhs53[_lhs54] = _rhs68;
            _lhs55.f1 = _rhs69;
            _lhs56[_lhs57] = _rhs70;
            _412_v53 = _rhs71;
            (globalState).f1 = (_this).f3;
          }
        }
      }
      let _421_v62;
      _421_v62 = _dafny.Set.fromElements(!(_this.f4));
      let _422_i11;
      _422_i11 = _dafny.ZERO;
      L4: {
        while ((_421_v62).IsProperSubsetOf((_module.__default.fm22(_this.f4, _this.f4, globalState)).Intersect(_dafny.Set.fromElements(_this.f4)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_422_i11)) {
              break L4;
            }
            _422_i11 = (_422_i11).plus(_dafny.ONE);
            let _423_v63;
            let _nw79 = Array((new BigNumber(6)).toNumber()).fill(_dafny.ZERO);
            _423_v63 = _nw79;
            let _424_v64;
            let _nw80 = Array((new BigNumber(11)).toNumber());
            _nw80[(_dafny.ZERO).toNumber()] = _423_v63;
            _nw80[(_dafny.ONE).toNumber()] = _423_v63;
            _nw80[(new BigNumber(2)).toNumber()] = _423_v63;
            _nw80[(new BigNumber(3)).toNumber()] = _423_v63;
            _nw80[(new BigNumber(4)).toNumber()] = _423_v63;
            _nw80[(new BigNumber(5)).toNumber()] = _423_v63;
            _nw80[(new BigNumber(6)).toNumber()] = _423_v63;
            _nw80[(new BigNumber(7)).toNumber()] = _423_v63;
            _nw80[(new BigNumber(8)).toNumber()] = _423_v63;
            _nw80[(new BigNumber(9)).toNumber()] = ((_this.f4) ? (_423_v63) : (_423_v63));
            _nw80[(new BigNumber(10)).toNumber()] = _423_v63;
            _424_v64 = _nw80;
            let _index65 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_424_v64).length));
            (_424_v64)[_index65] = _423_v63;
            let _425_v65;
            let _init12 = ((_426_v27) => function (_427_i12) {
              return _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_426_v27);
            })(_379_v27);
            let _nw81 = Array((_dafny.ONE).toNumber());
            for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw81.length); _i0_12++) {
              _nw81[_i0_12] = _init12(new BigNumber(_i0_12));
            }
            _425_v65 = _nw81;
            let _428_v66;
            _428_v66 = _module.D7.create_DC20(_dafny.Map.Empty.slice().updateUnsafe((_this).f3,_dafny.Seq.UnicodeFromString("tfjfsyru")));
            let _index66 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_425_v65).length));
            (_425_v65)[_index66] = (_428_v66).dtor_cf27;
            let _429_v67;
            _429_v67 = new _dafny.CodePoint('y'.codePointAt(0));
            let _430_v68;
            _430_v68 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p0),_379_v27);
            let _index67 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_424_v64).length));
            let _index68 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_425_v65).length));
            let _rhs72 = _423_v63;
            let _rhs73 = _430_v68;
            let _rhs74 = (_this).f3;
            let _rhs75 = _429_v67;
            let _lhs58 = _424_v64;
            let _lhs59 = _module.__default.safeIndex(new BigNumber(441), new BigNumber((_424_v64).length));
            let _lhs60 = _425_v65;
            let _lhs61 = _module.__default.safeIndex(new BigNumber(196), new BigNumber((_425_v65).length));
            let _lhs62 = globalState;
            _lhs58[_lhs59] = _rhs72;
            _lhs60[_lhs61] = _rhs73;
            _lhs62.f1 = _rhs74;
            _429_v67 = _rhs75;
            r0 = _module.__default.fm0(globalState);
            let _431_v69;
            let _init13 = function (_432_i13) {
              return _this.f4;
            };
            let _nw82 = Array((new BigNumber(16)).toNumber());
            for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw82.length); _i0_13++) {
              _nw82[_i0_13] = _init13(new BigNumber(_i0_13));
            }
            _431_v69 = _nw82;
            let _433_v70;
            _433_v70 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_this.f4);
            let _index69 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
            (_431_v69)[_index69] = (((_433_v70).contains(_this.f4)) ? ((_433_v70).get(_this.f4)) : (_this.f4));
            let _index70 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
            (_431_v69)[_index70] = !_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-39)), ((_434_v67) => function (_435_i14) {
              return _434_v67;
            })(_429_v67)), _429_v67);
            let _436_v71;
            _436_v71 = _module.D0.create_DC3();
            let _source5 = _436_v71;
            if (_source5.is_DC1) {
              let _437___mcc_h0 = (_source5).cf1;
              let _438___mcc_h1 = (_source5).cf2;
              let _439___mcc_h2 = (_source5).cf3;
              let _440___mcc_h3 = (_source5).cf4;
              let _441_cf4 = _440___mcc_h3;
              let _442_cf3 = _439___mcc_h2;
              let _443_cf2 = _438___mcc_h1;
              let _444_cf1 = _437___mcc_h0;
              let _index71 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_423_v63).length));
              (_423_v63)[_index71] = new BigNumber((_dafny.Seq.UnicodeFromString("cnff")).length);
              let _index72 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_423_v63).length));
              let _index73 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
              let _rhs76 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(909)), function (_445_i15) {
                return new _dafny.CodePoint('t'.codePointAt(0));
              }), _dafny.Seq.update(_dafny.Seq.UnicodeFromString("hmi"), _module.__default.safeIndex((_this).f7, new BigNumber((_dafny.Seq.UnicodeFromString("hmi")).length)), _429_v67))).length);
              let _rhs77 = (_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))];
              let _rhs78 = (_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))];
              let _lhs63 = _423_v63;
              let _lhs64 = _module.__default.safeIndex(new BigNumber(193), new BigNumber((_423_v63).length));
              let _lhs65 = _431_v69;
              let _lhs66 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
              let _lhs67 = globalState;
              _lhs63[_lhs64] = _rhs76;
              _lhs65[_lhs66] = _rhs77;
              _lhs67.f2 = _rhs78;
              let _446_v72;
              _446_v72 = _dafny.Seq.of(((_dafny.ZERO).minus((((_380_v28).contains(p0)) ? ((_380_v28).get(p0)) : (p0)))).isLessThanOrEqualTo((_423_v63)[_module.__default.safeIndex(new BigNumber(193), new BigNumber((_423_v63).length))]), !(_443_cf2));
              let _index74 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
              (_431_v69)[_index74] = (_446_v72)[_module.__default.safeIndex((_dafny.ZERO).minus((p0).minus((_this).f3)), new BigNumber((_446_v72).length))];
              let _447_v73;
              _447_v73 = _module.D6.create_DC17(_module.__default.fm1(_429_v67, _443_cf2, globalState));
              let _448_v74;
              _448_v74 = _module.D5.create_DC14(_dafny.Map.Empty.slice().updateUnsafe(_443_cf2,_dafny.MultiSet.FromArray((_447_v73).dtor_cf23)));
              _448_v74 = _448_v74;
              let _index75 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
              (_431_v69)[_index75] = false;
            } else if (_source5.is_DC2) {
              let _449___mcc_h4 = (_source5).cf5;
              let _450___mcc_h5 = (_source5).cf6;
              let _451_cf6 = _450___mcc_h5;
              let _452_cf5 = _449___mcc_h4;
              let _index76 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
              (_431_v69)[_index76] = !((new BigNumber((_379_v27).length)).isLessThan((_this).f3));
              (globalState).f2 = !((_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))]) || (_451_cf6);
              _379_v27 = _379_v27;
              let _453_v75;
              _453_v75 = _dafny.Seq.of((((_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))]) ? (_this.f4) : ((_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))])));
              let _454_v76;
              _454_v76 = _dafny.Seq.of(_453_v75, _453_v75, _dafny.Seq.of(_this.f4), _453_v75);
              let _455_v77;
              _455_v77 = _dafny.MultiSet.fromElements((_this).f3, new BigNumber((_453_v75).length));
              _453_v75 = (_454_v76)[_module.__default.safeIndex((_dafny.ZERO).minus(_module.__default.safeModuloInt((((_455_v77).contains((_this).f7)) ? ((_455_v77).get((_this).f7)) : (p0)), (_this).f3)), new BigNumber((_454_v76).length))];
            } else if (_source5.is_DC3) {
              let _456_v78;
              let _nw83 = Array((new BigNumber(27)).toNumber()).fill([]);
              _456_v78 = _nw83;
              let _457_v79;
              let _nw84 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
              _457_v79 = _nw84;
              let _index77 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_456_v78).length));
              (_456_v78)[_index77] = _457_v79;
              let _index78 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_456_v78).length));
              let _index79 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
              let _rhs79 = _457_v79;
              let _rhs80 = _429_v67;
              let _rhs81 = ((_421_v62).IsDisjointFrom(_421_v62)) && (_this.f4);
              let _lhs68 = _456_v78;
              let _lhs69 = _module.__default.safeIndex(new BigNumber(158), new BigNumber((_456_v78).length));
              let _lhs70 = _431_v69;
              let _lhs71 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
              _lhs68[_lhs69] = _rhs79;
              _429_v67 = _rhs80;
              _lhs70[_lhs71] = _rhs81;
              let _458_v80;
              _458_v80 = _dafny.Map.Empty.slice().updateUnsafe(_379_v27,_431_v69);
              _458_v80 = (_458_v80).update(_dafny.Seq.Concat(_379_v27, _379_v27), _431_v69);
              let _459_v81;
              _459_v81 = _dafny.Seq.of(new BigNumber(560));
              let _460_v82;
              _460_v82 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_459_v81).length),!((_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))]) || ((_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))]));
              _460_v82 = (_460_v82).update((_this).f3, (_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))]);
              let _461_v83;
              _461_v83 = _dafny.MultiSet.fromElements((_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))], !(p0).isEqualTo(p0), (_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))]);
              _461_v83 = _461_v83;
            } else if (_source5.is_DC0) {
              let _462___mcc_h6 = (_source5).cf0;
              let _463_cf0 = _462___mcc_h6;
              _379_v27 = _dafny.Seq.Concat(_379_v27, _dafny.Seq.UnicodeFromString("uclqne"));
              let _464_v84;
              _464_v84 = _431_v69;
              let _465_v85;
              let _nw85 = new _module.C0();
              _nw85.__ctor(_module.__default.fm2(p0, globalState), (_464_v84));
              _465_v85 = _nw85;
              let _466_v86;
              _466_v86 = _dafny.Map.Empty.slice().updateUnsafe(_465_v85,(_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))]);
              _466_v86 = (_466_v86).update(_465_v85, false);
              let _467_v87;
              _467_v87 = _dafny.MultiSet.fromElements(_429_v67);
              let _468_v88;
              _468_v88 = _module.D0.create_DC2(_467_v87, (_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))]);
              (_465_v85).f10 = (((_463_cf0) ? (_468_v88) : (_468_v88))).dtor_cf6;
              let _469_v89;
              _469_v89 = _dafny.MultiSet.fromElements(false, _module.__default.fm2((_this).f7, globalState), _465_v85.f10);
              let _470_v90;
              _470_v90 = _dafny.Seq.of(_469_v89, _469_v89, (_469_v89).Union(_469_v89), ((_469_v89).update(_this.f4, _module.__default.abs((_dafny.ZERO).minus((_this).f7)))).Difference(_469_v89), (_dafny.MultiSet.fromElements(_465_v85.f10)).Difference(_469_v89));
              let _471_v91;
              _471_v91 = _dafny.MultiSet.fromElements(p0);
              _470_v90 = _dafny.Seq.update(_470_v90, _module.__default.safeIndex((((_471_v91).contains(p0)) ? ((_471_v91).get(p0)) : ((_this).f3)), new BigNumber((_470_v90).length)), _469_v89);
            } else {
              let _472___mcc_h7 = (_source5).cf7;
              let _473_cf7 = _472___mcc_h7;
              let _474_v92;
              _474_v92 = _dafny.Seq.of(!(false));
              (globalState).f2 = (_this.f4) || (!_dafny.Seq.contains(_474_v92, (_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))]));
              let _475_v93;
              _475_v93 = _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements((_this).f7, (_this).f7)).length));
              let _index80 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
              (_431_v69)[_index80] = (new BigNumber(563)).isLessThan((_475_v93)[_module.__default.safeIndex((_this).f7, new BigNumber((_475_v93).length))]);
              (globalState).f2 = (_431_v69)[_module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length))];
              let _index81 = _module.__default.safeIndex(new BigNumber(813), new BigNumber((_431_v69).length));
              (_431_v69)[_index81] = _this.f4;
            }
          }
        }
      }
      r0 = (_this).f7;
      return r0;
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f4 = false;
      this._f3 = _dafny.ZERO;
      this._f14 = _dafny.ZERO;
      this._f15 = [];
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f14, f15, f3, f4) {
      let _this = this;
      (_this)._f14 = f14;
      (_this)._f15 = f15;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    m6(p0, p1, p2, globalState) {
      let _this = this;
      if (_this.f4) {
        (_this).f4 = _this.f4;
        (globalState).f2 = (p1).isLessThanOrEqualTo(p2);
        if (true) {
          let _476_v0;
          let _init14 = function (_477_i0) {
            return (_module.D0.create_DC0(false)).dtor_cf0;
          };
          let _nw86 = Array((new BigNumber(27)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw86.length); _i0_14++) {
            _nw86[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _476_v0 = _nw86;
          let _478_v1;
          let _nw87 = new _module.C0();
          _nw87.__ctor(_this.f4, _476_v0);
          _478_v1 = _nw87;
          let _479_v2;
          let _nw88 = new _module.C0();
          _nw88.__ctor(((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ajxdmcbun")).length))).isEqualTo((_this).f3), _476_v0);
          _479_v2 = _nw88;
          let _index82 = _module.__default.safeIndex(new BigNumber(460), new BigNumber(((_this).f15).length));
          ((_this).f15)[_index82] = p1;
          let _index83 = _module.__default.safeIndex(new BigNumber(460), new BigNumber(((_this).f15).length));
          ((_this).f15)[_index83] = (_this).f3;
          let _480_v3;
          _480_v3 = _dafny.Seq.UnicodeFromString("di");
          let _481_v4;
          _481_v4 = _dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.Concat(_480_v3, _480_v3));
          _481_v4 = (_481_v4).update(_this.f4, _dafny.Seq.UnicodeFromString("oeaya"));
          let _482_v5;
          _482_v5 = _dafny.Seq.of(_this.f4, _478_v1.f10, _479_v2.f10);
          let _483_v6;
          _483_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.FromArray(_482_v5)).cardinality()),(p2).plus((_this).f14));
          let _484_v7;
          _484_v7 = _dafny.MultiSet.fromElements((_this).f14);
          let _index84 = _module.__default.safeIndex(new BigNumber(460), new BigNumber(((_this).f15).length));
          let _rhs82 = (((_483_v6).contains(((_478_v1.f10) ? (new BigNumber(435)) : ((_dafny.ZERO).minus((_this).f14))))) ? ((_483_v6).get(((_478_v1.f10) ? (new BigNumber(435)) : ((_dafny.ZERO).minus((_this).f14))))) : (p1));
          let _rhs83 = _module.__default.safeDivisionInt((_this).f14, ((_dafny.ZERO).minus(p2)).multipliedBy(p2));
          let _rhs84 = ((_484_v7).Intersect(_484_v7)).IsProperSubsetOf(_484_v7);
          let _lhs72 = (_this).f15;
          let _lhs73 = _module.__default.safeIndex(new BigNumber(460), new BigNumber(((_this).f15).length));
          let _lhs74 = globalState;
          let _lhs75 = _479_v2;
          _lhs72[_lhs73] = _rhs82;
          _lhs74.f1 = _rhs83;
          _lhs75.f10 = _rhs84;
        } else {
          let _485_v8;
          _485_v8 = _dafny.Seq.of(new BigNumber(-723), (_this).f14, (_this).f3, p2, (_this).f3);
          (globalState).f1 = (_485_v8)[_module.__default.safeIndex(p2, new BigNumber((_485_v8).length))];
          (globalState).f1 = p1;
          let _486_v9;
          let _nw89 = Array((new BigNumber(2)).toNumber()).fill(false);
          _486_v9 = _nw89;
          let _index85 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_486_v9).length));
          (_486_v9)[_index85] = _this.f4;
          let _index86 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_486_v9).length));
          (_486_v9)[_index86] = _this.f4;
          let _487_v10;
          _487_v10 = _dafny.Map.Empty.slice().updateUnsafe((_486_v9)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_486_v9).length))],(_this).f3);
          (globalState).f1 = new BigNumber(((((_486_v9)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_486_v9).length))]) ? (_487_v10) : (_487_v10))).length);
          let _488_v11;
          let _nw90 = new _module.C0();
          _nw90.__ctor((p1).isLessThan((_this).f3), _486_v9);
          _488_v11 = _nw90;
        }
        let _489_v12;
        let _nw91 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
        _489_v12 = _nw91;
        let _490_v13;
        let _init15 = function (_491_i1) {
          return _this.f4;
        };
        let _nw92 = Array((new BigNumber(15)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw92.length); _i0_15++) {
          _nw92[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _490_v13 = _nw92;
        let _492_v14;
        _492_v14 = _dafny.MultiSet.fromElements(_490_v13, _490_v13);
        let _index87 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_489_v12).length));
        (_489_v12)[_index87] = _492_v14;
        let _index88 = _module.__default.safeIndex(new BigNumber(404), new BigNumber((_489_v12).length));
        (_489_v12)[_index88] = _dafny.MultiSet.fromElements(_490_v13);
        (globalState).f2 = false;
      } else {
        let _493_v15;
        let _nw93 = Array((new BigNumber(11)).toNumber()).fill(_dafny.MultiSet.Empty);
        _493_v15 = _nw93;
        let _494_v16;
        _494_v16 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f14));
        let _index89 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_493_v15).length));
        (_493_v15)[_index89] = _494_v16;
        let _495_v17;
        _495_v17 = new _dafny.CodePoint('f'.codePointAt(0));
        let _496_v18;
        _496_v18 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_495_v17);
        let _index90 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_493_v15).length));
        (_493_v15)[_index90] = (_494_v16).Union(_dafny.MultiSet.fromElements(new BigNumber((_496_v18).length)));
        let _rhs85 = (_this).f14;
        let _rhs86 = _this.f4;
        let _lhs76 = globalState;
        let _lhs77 = globalState;
        _lhs76.f1 = _rhs85;
        _lhs77.f2 = _rhs86;
        let _497_v19;
        let _init16 = function (_498_i2) {
          return _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_dafny.MultiSet.fromElements(_this.f4, _this.f4, _this.f4));
        };
        let _nw94 = Array((new BigNumber(4)).toNumber());
        for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw94.length); _i0_16++) {
          _nw94[_i0_16] = _init16(new BigNumber(_i0_16));
        }
        _497_v19 = _nw94;
        let _499_v20;
        _499_v20 = _dafny.MultiSet.fromElements(_this.f4, false);
        let _500_v21;
        _500_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_this).f14, globalState),_499_v20);
        let _index91 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_497_v19).length));
        (_497_v19)[_index91] = _500_v21;
        let _501_v22;
        _501_v22 = _module.D5.create_DC14(_500_v21);
        let _index92 = _module.__default.safeIndex(new BigNumber(931), new BigNumber((_497_v19).length));
        (_497_v19)[_index92] = (((_this.f4) ? (_501_v22) : (_501_v22))).dtor_cf20;
        (globalState).f2 = !(_this.f4);
        let _502_v23;
        _502_v23 = _dafny.Seq.UnicodeFromString("arajbrv");
        if (_dafny.Seq.IsProperPrefixOf(_502_v23, _502_v23)) {
          let _503_v24;
          let _init17 = function (_504_i3) {
            return !(_this.f4);
          };
          let _nw95 = Array((new BigNumber(5)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw95.length); _i0_17++) {
            _nw95[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _503_v24 = _nw95;
          let _505_v25;
          let _nw96 = Array((new BigNumber(25)).toNumber());
          _nw96[(_dafny.ZERO).toNumber()] = ((_this.f4) ? (_503_v24) : (_503_v24));
          _nw96[(_dafny.ONE).toNumber()] = _503_v24;
          _nw96[(new BigNumber(2)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(3)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(4)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(5)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(6)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(7)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(8)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(9)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(10)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(11)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(12)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(13)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(14)).toNumber()] = ((!(_this.f4)) ? (_503_v24) : (_503_v24));
          _nw96[(new BigNumber(15)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(16)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(17)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(18)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(19)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(20)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(21)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(22)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(23)).toNumber()] = _503_v24;
          _nw96[(new BigNumber(24)).toNumber()] = _503_v24;
          _505_v25 = _nw96;
          let _index93 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_505_v25).length));
          (_505_v25)[_index93] = _503_v24;
          let _index94 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_505_v25).length));
          let _rhs87 = _503_v24;
          let _rhs88 = _module.__default.fm2(((_this).f14).plus(p1), globalState);
          let _lhs78 = _505_v25;
          let _lhs79 = _module.__default.safeIndex(new BigNumber(325), new BigNumber((_505_v25).length));
          let _lhs80 = globalState;
          _lhs78[_lhs79] = _rhs87;
          _lhs80.f2 = _rhs88;
          (globalState).f1 = (p2).multipliedBy(p2);
          let _index95 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_493_v15).length));
          let _rhs89 = (_493_v15)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_493_v15).length))];
          let _rhs90 = _this.f4;
          let _lhs81 = _493_v15;
          let _lhs82 = _module.__default.safeIndex(new BigNumber(319), new BigNumber((_493_v15).length));
          let _lhs83 = globalState;
          _lhs81[_lhs82] = _rhs89;
          _lhs83.f2 = _rhs90;
          _502_v23 = _502_v23;
          let _506_v26;
          _506_v26 = _dafny.Set.fromElements(((_493_v15)[_module.__default.safeIndex(new BigNumber(319), new BigNumber((_493_v15).length))]).IsProperSubsetOf(_494_v16), _this.f4, _this.f4, _this.f4);
          let _507_v27;
          _507_v27 = _dafny.Seq.of(_506_v26);
          let _508_v28;
          _508_v28 = _dafny.Seq.of(p2, p2);
          _506_v26 = (_507_v27)[_module.__default.safeIndex(new BigNumber((_508_v28).length), new BigNumber((_507_v27).length))];
        } else {
          let _509_v29;
          _509_v29 = _dafny.Set.fromElements((_this).f15, (_this).f15, (_this).f15, (_this).f15, (_this).f15);
          let _510_v30;
          _510_v30 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_509_v29);
          _510_v30 = (_510_v30).update(false, _509_v29);
          let _511_v31;
          _511_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(214),(p1).isEqualTo(p2));
          _511_v31 = (_511_v31).update(p1, _this.f4);
          (globalState).f1 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_this).f14, (_this).f14), (_this).f14);
          let _512_v32;
          let _init18 = ((_513_v21) => function (_514_i4) {
            return _module.D5.create_DC16(_module.D5.create_DC16(_module.D5.create_DC16(_module.D5.create_DC16(_module.D5.create_DC14((_513_v21).update(_this.f4, _dafny.MultiSet.fromElements(_this.f4, _this.f4)))))));
          })(_500_v21);
          let _nw97 = Array((new BigNumber(4)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw97.length); _i0_18++) {
            _nw97[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _512_v32 = _nw97;
          let _index96 = _module.__default.safeIndex(new BigNumber(114), new BigNumber(((_this).f15).length));
          ((_this).f15)[_index96] = new BigNumber((_499_v20).cardinality());
          let _index97 = _module.__default.safeIndex(new BigNumber(114), new BigNumber(((_this).f15).length));
          let _rhs91 = _512_v32;
          let _rhs92 = !(_this.f4) || (_this.f4);
          let _rhs93 = _this.f4;
          let _rhs94 = _module.__default.safeModuloInt((_this).f14, p1);
          let _rhs95 = ((_dafny.ZERO).minus((_this).f14)).plus(((_this).f14).minus(p1));
          let _lhs84 = globalState;
          let _lhs85 = globalState;
          let _lhs86 = globalState;
          let _lhs87 = (_this).f15;
          let _lhs88 = _module.__default.safeIndex(new BigNumber(114), new BigNumber(((_this).f15).length));
          _512_v32 = _rhs91;
          _lhs84.f2 = _rhs92;
          _lhs85.f2 = _rhs93;
          _lhs86.f1 = _rhs94;
          _lhs87[_lhs88] = _rhs95;
          let _515_v33;
          let _init19 = function (_516_i5) {
            return (_this.f4) === (_this.f4);
          };
          let _nw98 = Array((new BigNumber(15)).toNumber());
          for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw98.length); _i0_19++) {
            _nw98[_i0_19] = _init19(new BigNumber(_i0_19));
          }
          _515_v33 = _nw98;
          let _nw99 = Array((new BigNumber(19)).toNumber()).fill(false);
          _515_v33 = _nw99;
        }
      }
      let _517_i6;
      _517_i6 = _dafny.ZERO;
      L5: {
        while (_module.__default.fm2((_dafny.ZERO).minus((_this).f14), globalState)) {
          C5: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_517_i6)) {
              break L5;
            }
            _517_i6 = (_517_i6).plus(_dafny.ONE);
            let _518_v34;
            _518_v34 = new _dafny.CodePoint('x'.codePointAt(0));
            let _519_v35;
            _519_v35 = _dafny.MultiSet.fromElements(_518_v34, _518_v34);
            let _520_v36;
            _520_v36 = _module.D0.create_DC2(_519_v35, _this.f4);
            let _521_v37;
            _521_v37 = _module.D0.create_DC4(_520_v36);
            let _rhs96 = p2;
            let _rhs97 = (_module.__default.fm0(globalState)).isEqualTo(new BigNumber(776));
            let _rhs98 = _this.f4;
            let _rhs99 = ((_this.f4) ? (_521_v37) : (_521_v37));
            let _lhs89 = globalState;
            let _lhs90 = globalState;
            let _lhs91 = globalState;
            _lhs89.f1 = _rhs96;
            _lhs90.f2 = _rhs97;
            _lhs91.f2 = _rhs98;
            _521_v37 = _rhs99;
            let _522_v38;
            let _nw100 = Array((new BigNumber(8)).toNumber()).fill([]);
            _522_v38 = _nw100;
            let _523_v39;
            let _nw101 = Array((new BigNumber(18)).toNumber());
            _nw101[(_dafny.ZERO).toNumber()] = _this.f4;
            _nw101[(_dafny.ONE).toNumber()] = _this.f4;
            _nw101[(new BigNumber(2)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(3)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(4)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(5)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(6)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(7)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(8)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(9)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(10)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(11)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(12)).toNumber()] = !(_this.f4);
            _nw101[(new BigNumber(13)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(14)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(15)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(16)).toNumber()] = _this.f4;
            _nw101[(new BigNumber(17)).toNumber()] = _module.__default.fm2(p1, globalState);
            _523_v39 = _nw101;
            let _index98 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_522_v38).length));
            (_522_v38)[_index98] = _523_v39;
            let _524_v40;
            _524_v40 = _dafny.Seq.of(((_this.f4) ? (_this.f4) : (_this.f4)), false, _this.f4, _this.f4);
            let _index99 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_523_v39).length));
            (_523_v39)[_index99] = _this.f4;
            let _index100 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_522_v38).length));
            let _index101 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_523_v39).length));
            let _rhs100 = _523_v39;
            let _rhs101 = _524_v40;
            let _rhs102 = (p1).isEqualTo((_this).f3);
            let _lhs92 = _522_v38;
            let _lhs93 = _module.__default.safeIndex(new BigNumber(909), new BigNumber((_522_v38).length));
            let _lhs94 = _523_v39;
            let _lhs95 = _module.__default.safeIndex(new BigNumber(221), new BigNumber((_523_v39).length));
            _lhs92[_lhs93] = _rhs100;
            _524_v40 = _rhs101;
            _lhs94[_lhs95] = _rhs102;
            let _525_v41;
            _525_v41 = _dafny.Map.Empty.slice().updateUnsafe((_523_v39)[_module.__default.safeIndex(new BigNumber(221), new BigNumber((_523_v39).length))],(_this).f3);
            let _526_v42;
            _526_v42 = _dafny.MultiSet.fromElements(new BigNumber(850));
            let _527_v43;
            _527_v43 = _dafny.MultiSet.fromElements(false, (_523_v39)[_module.__default.safeIndex(new BigNumber(221), new BigNumber((_523_v39).length))]);
            let _528_v44;
            _528_v44 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_519_v35).cardinality()),_this.f4);
            let _529_v45;
            _529_v45 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,new BigNumber((_528_v44).length));
            let _530_v46;
            let _nw102 = Array((new BigNumber(25)).toNumber());
            _nw102[(_dafny.ZERO).toNumber()] = _module.__default.safeModuloInt(p1, (_this).f3);
            _nw102[(_dafny.ONE).toNumber()] = (new BigNumber((_525_v41).length)).plus(new BigNumber(688));
            _nw102[(new BigNumber(2)).toNumber()] = p1;
            _nw102[(new BigNumber(3)).toNumber()] = (p1).multipliedBy((_this).f3);
            _nw102[(new BigNumber(4)).toNumber()] = ((_this).f14).multipliedBy((_this).f14);
            _nw102[(new BigNumber(5)).toNumber()] = _module.__default.fm0(globalState);
            _nw102[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt((_this).f3, (_this).f3);
            _nw102[(new BigNumber(7)).toNumber()] = (p2).minus(new BigNumber(-388));
            _nw102[(new BigNumber(8)).toNumber()] = (_dafny.ZERO).minus((_this).f14);
            _nw102[(new BigNumber(9)).toNumber()] = ((_module.D1.create_DC5(_module.__default.fm0(globalState))).dtor_cf8).minus(p1);
            _nw102[(new BigNumber(10)).toNumber()] = (((_526_v42).contains(p2)) ? ((_526_v42).get(p2)) : ((_this).f14));
            _nw102[(new BigNumber(11)).toNumber()] = (_this).f14;
            _nw102[(new BigNumber(12)).toNumber()] = (_this).f14;
            _nw102[(new BigNumber(13)).toNumber()] = p2;
            _nw102[(new BigNumber(14)).toNumber()] = p2;
            _nw102[(new BigNumber(15)).toNumber()] = p2;
            _nw102[(new BigNumber(16)).toNumber()] = (_this).f3;
            _nw102[(new BigNumber(17)).toNumber()] = p2;
            _nw102[(new BigNumber(18)).toNumber()] = p1;
            _nw102[(new BigNumber(19)).toNumber()] = p1;
            _nw102[(new BigNumber(20)).toNumber()] = (_this).f3;
            _nw102[(new BigNumber(21)).toNumber()] = p1;
            _nw102[(new BigNumber(22)).toNumber()] = (((_527_v43).contains((_523_v39)[_module.__default.safeIndex(new BigNumber(221), new BigNumber((_523_v39).length))])) ? ((_527_v43).get((_523_v39)[_module.__default.safeIndex(new BigNumber(221), new BigNumber((_523_v39).length))])) : (new BigNumber((_529_v45).length)));
            _nw102[(new BigNumber(23)).toNumber()] = (_this).f3;
            _nw102[(new BigNumber(24)).toNumber()] = new BigNumber(-440);
            _530_v46 = _nw102;
            let _531_v47;
            let _nw103 = Array((new BigNumber(19)).toNumber());
            _531_v47 = _nw103;
            let _532_v48;
            _532_v48 = _module.D8.create_DC22(_528_v44);
            let _533_v49;
            let _nw104 = new _module.C1();
            _nw104.__ctor((_this).f3, new BigNumber(((_532_v48).dtor_cf31).length), false);
            _533_v49 = _nw104;
            let _index102 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_531_v47).length));
            (_531_v47)[_index102] = _533_v49;
            let _534_v50;
            _534_v50 = _dafny.Map.Empty.slice().updateUnsafe((_529_v45).update(p2, (_this).f3),_530_v46);
            let _index103 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_531_v47).length));
            let _rhs103 = (_523_v39)[_module.__default.safeIndex(new BigNumber(221), new BigNumber((_523_v39).length))];
            let _rhs104 = p1;
            let _rhs105 = (((_534_v50).contains((_module.__default.fm23(globalState)).update((_533_v49).f7, (_dafny.ZERO).minus((_533_v49).f7)))) ? ((_534_v50).get((_module.__default.fm23(globalState)).update((_533_v49).f7, (_dafny.ZERO).minus((_533_v49).f7)))) : (_530_v46));
            let _rhs106 = _533_v49;
            let _rhs107 = (_533_v49).f3;
            let _lhs96 = _this;
            let _lhs97 = globalState;
            let _lhs98 = _531_v47;
            let _lhs99 = _module.__default.safeIndex(new BigNumber(327), new BigNumber((_531_v47).length));
            let _lhs100 = globalState;
            _lhs96.f4 = _rhs103;
            _lhs97.f1 = _rhs104;
            _530_v46 = _rhs105;
            _lhs98[_lhs99] = _rhs106;
            _lhs100.f1 = _rhs107;
            let _535_v51;
            let _nw105 = Array((new BigNumber(17)).toNumber());
            _nw105[(_dafny.ZERO).toNumber()] = _518_v34;
            _nw105[(_dafny.ONE).toNumber()] = _518_v34;
            _nw105[(new BigNumber(2)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(3)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(4)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(5)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(6)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(7)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(8)).toNumber()] = ((_this.f4) ? (new _dafny.CodePoint('k'.codePointAt(0))) : (_518_v34));
            _nw105[(new BigNumber(9)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(10)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(11)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(12)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(13)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(14)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(15)).toNumber()] = _518_v34;
            _nw105[(new BigNumber(16)).toNumber()] = _518_v34;
            _535_v51 = _nw105;
            let _index104 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_535_v51).length));
            (_535_v51)[_index104] = _518_v34;
            let _536_v52;
            _536_v52 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(p1, p1),new _dafny.CodePoint('n'.codePointAt(0)));
            let _index105 = _module.__default.safeIndex(new BigNumber(317), new BigNumber((_535_v51).length));
            (_535_v51)[_index105] = (((_536_v52).contains(new BigNumber(693))) ? ((_536_v52).get(new BigNumber(693))) : (new _dafny.CodePoint('g'.codePointAt(0))));
          }
        }
      }
      let _537_v53;
      let _init20 = function (_538_i8) {
        return false;
      };
      let _nw106 = Array((new BigNumber(3)).toNumber());
      for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw106.length); _i0_20++) {
        _nw106[_i0_20] = _init20(new BigNumber(_i0_20));
      }
      _537_v53 = _nw106;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_537_v53).length))) {
        let _539_i7 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_539_i7)) && ((_539_i7).isLessThan(new BigNumber((_537_v53).length))))) {
          (_537_v53)[(_539_i7)] = _this.f4;
        }
      }
      let _540_v54;
      _540_v54 = _dafny.Seq.of(_this.f4, _this.f4);
      let _541_v55;
      _541_v55 = _module.D7.create_DC21(_537_v53, _dafny.Seq.IsProperPrefixOf(_540_v54, _dafny.Seq.of(_this.f4, _this.f4, _this.f4, _this.f4)), !(_this.f4));
      _541_v55 = _541_v55;
      let _542_v56;
      _542_v56 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,(p1).plus(p2));
      let _543_v57;
      _543_v57 = new _dafny.CodePoint('v'.codePointAt(0));
      let _544_v58;
      _544_v58 = _dafny.MultiSet.fromElements(new _dafny.CodePoint('w'.codePointAt(0)), _543_v57, new _dafny.CodePoint('r'.codePointAt(0)));
      _542_v56 = (_542_v56).update((_module.D0.create_DC2(_544_v58, _this.f4)).dtor_cf6, _module.__default.safeDivisionInt((_this).f3, (_this).f3));
      if (!((_this.f4) || (true)) || (_this.f4)) {
        let _545_v59;
        _545_v59 = _module.D8.create_DC23(_this.f4, _this.f4, _this.f4, _543_v57);
        let _546_v60;
        _546_v60 = _dafny.Seq.UnicodeFromString("dakcdomv");
        let _pat_let_tv4 = globalState;
        let _rhs108 = function (_pat_let5_0) {
          return function (_547_dt__update__tmp_h0) {
            return function (_pat_let6_0) {
              return function (_548_dt__update_hcf35_h0) {
                return function (_pat_let7_0) {
                  return function (_549_dt__update_hcf32_h0) {
                    return _module.D8.create_DC23(_549_dt__update_hcf32_h0, (_547_dt__update__tmp_h0).dtor_cf33, (_547_dt__update__tmp_h0).dtor_cf34, _548_dt__update_hcf35_h0);
                  }(_pat_let7_0);
                }(_this.f4);
              }(_pat_let6_0);
            }(_module.__default.fm24(_pat_let_tv4));
          }(_pat_let5_0);
        }(((true) ? (_545_v59) : (_545_v59)));
        let _rhs109 = _this.f4;
        let _rhs110 = _dafny.Seq.contains(_546_v60, _543_v57);
        let _lhs101 = globalState;
        let _lhs102 = globalState;
        _545_v59 = _rhs108;
        _lhs101.f2 = _rhs109;
        _lhs102.f2 = _rhs110;
        _543_v57 = _543_v57;
        let _550_v61;
        _550_v61 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,(_this).f14);
        let _index106 = _module.__default.safeIndex(new BigNumber(271), new BigNumber(((_this).f15).length));
        ((_this).f15)[_index106] = (((_550_v61).contains((_this).f3)) ? ((_550_v61).get((_this).f3)) : ((_dafny.ZERO).minus((_this).f3)));
        let _551_v62;
        _551_v62 = _dafny.Set.fromElements(_this.f4);
        let _index107 = _module.__default.safeIndex(new BigNumber(271), new BigNumber(((_this).f15).length));
        ((_this).f15)[_index107] = (p1).plus(new BigNumber(((_551_v62).Union(_551_v62)).length));
        _540_v54 = _540_v54;
        let _552_v63;
        _552_v63 = _dafny.Seq.of(p1, p2);
        (globalState).f1 = new BigNumber((_dafny.Seq.Concat(_552_v63, _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(222)), ((_553_p2) => function (_554_i9) {
          return _553_p2;
        })(p2)), _552_v63), _module.__default.safeIndex(((_this).f15)[_module.__default.safeIndex(new BigNumber(271), new BigNumber(((_this).f15).length))], new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(222)), ((_555_p2) => function (_556_i9) {
          return _555_p2;
        })(p2)), _552_v63)).length)), p1))).length);
      } else {
        let _557_v64;
        _557_v64 = _dafny.Map.Empty.slice().updateUnsafe(_540_v54,(_this).f3);
        (globalState).f1 = (((_557_v64).contains(_dafny.Seq.update(_dafny.Seq.Concat(_540_v54, _540_v54), _module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_dafny.Seq.Concat(_540_v54, _540_v54)).length)), true))) ? ((_557_v64).get(_dafny.Seq.update(_dafny.Seq.Concat(_540_v54, _540_v54), _module.__default.safeIndex(_module.__default.fm0(globalState), new BigNumber((_dafny.Seq.Concat(_540_v54, _540_v54)).length)), true))) : (p1));
        let _558_v65;
        _558_v65 = _module.D8.create_DC23(_this.f4, _this.f4, _this.f4, _543_v57);
        let _559_v66;
        _559_v66 = _module.D8.create_DC24(_558_v65);
        let _560_v67;
        _560_v67 = _module.D8.create_DC24(_559_v66);
        let _561_v68;
        _561_v68 = _module.D8.create_DC24(_560_v67);
        let _562_v69;
        _562_v69 = _dafny.Map.Empty.slice().updateUnsafe(true,_561_v68);
        _562_v69 = ((_562_v69).Merge(_562_v69)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_this.f4,_561_v68));
        let _563_v70;
        let _nw107 = new _module.C1();
        _nw107.__ctor((_this).f14, p2, _this.f4);
        _563_v70 = _nw107;
        let _564_v71;
        _564_v71 = _dafny.Seq.UnicodeFromString("axxyv");
        _564_v71 = _564_v71;
        let _565_v72;
        let _init21 = ((_566_v58) => function (_567_i10) {
          return _dafny.Seq.of(new BigNumber((_566_v58).cardinality()));
        })(_544_v58);
        let _nw108 = Array((new BigNumber(2)).toNumber());
        for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw108.length); _i0_21++) {
          _nw108[_i0_21] = _init21(new BigNumber(_i0_21));
        }
        _565_v72 = _nw108;
        let _568_v73;
        _568_v73 = _dafny.Set.fromElements((_this).f3);
        let _569_v74;
        _569_v74 = _dafny.Map.Empty.slice().updateUnsafe(true,true);
        let _rhs111 = new BigNumber(-860);
        let _rhs112 = !((_this.f4) && ((_568_v73).equals(_568_v73)));
        let _rhs113 = _module.__default.fm2(new BigNumber(((_569_v74).Merge((_569_v74).update(_this.f4, _this.f4))).length), globalState);
        let _rhs114 = _module.__default.fm2(p2, globalState);
        let _rhs115 = _565_v72;
        let _lhs103 = globalState;
        let _lhs104 = globalState;
        let _lhs105 = globalState;
        let _lhs106 = globalState;
        _lhs103.f1 = _rhs111;
        _lhs104.f2 = _rhs112;
        _lhs105.f2 = _rhs113;
        _lhs106.f2 = _rhs114;
        _565_v72 = _rhs115;
      }
      return;
    }
    m7(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      r0 = (_this).f14;
      let _570_v0;
      _570_v0 = _dafny.Seq.of(new BigNumber(-638));
      let _571_v1;
      _571_v1 = _dafny.Seq.of(new BigNumber((_570_v0).length), (p0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_this).f3)).length)), (new BigNumber(-673)).plus((_this).f14));
      _571_v1 = _570_v0;
      let _hi5 = _module.__default.fm0(globalState);
      for (let _572_i0 = (_this).f14; _572_i0.isLessThan(_hi5); _572_i0 = _572_i0.plus(_dafny.ONE)) {
        let _573_v2;
        _573_v2 = new _dafny.CodePoint('i'.codePointAt(0));
        let _574_v3;
        _574_v3 = _dafny.Set.fromElements(_573_v2);
        let _575_v4;
        let _nw109 = Array((new BigNumber(6)).toNumber());
        _nw109[(_dafny.ZERO).toNumber()] = _574_v3;
        _nw109[(_dafny.ONE).toNumber()] = _574_v3;
        _nw109[(new BigNumber(2)).toNumber()] = (_574_v3).Intersect(_574_v3);
        _nw109[(new BigNumber(3)).toNumber()] = _574_v3;
        _nw109[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements(_module.__default.fm24(globalState), _573_v2, _573_v2);
        _nw109[(new BigNumber(5)).toNumber()] = ((_this.f4) ? (_574_v3) : (_574_v3));
        _575_v4 = _nw109;
        _575_v4 = _575_v4;
        let _source6 = _module.__default.fm25(((_this.f4) ? (_module.__default.fm2(new BigNumber(-606), globalState)) : (_this.f4)), globalState);
        if (_source6.is_DC1) {
          let _576___mcc_h0 = (_source6).cf1;
          let _577___mcc_h1 = (_source6).cf2;
          let _578___mcc_h2 = (_source6).cf3;
          let _579___mcc_h3 = (_source6).cf4;
          let _580_cf4 = _579___mcc_h3;
          let _581_cf3 = _578___mcc_h2;
          let _582_cf2 = _577___mcc_h1;
          let _583_cf1 = _576___mcc_h0;
          let _584_v5;
          _584_v5 = _module.D0.create_DC1(false, _583_cf1, _581_cf3, _580_cf4);
          (globalState).f1 = new BigNumber((_dafny.Seq.of((_584_v5).dtor_cf3)).length);
          let _585_v6;
          _585_v6 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f4);
          let _586_v7;
          _586_v7 = _dafny.Set.fromElements(_583_cf1);
          let _587_v8;
          _587_v8 = _dafny.Map.Empty.slice().updateUnsafe(_583_cf1,_586_v7);
          let _588_v9;
          _588_v9 = _dafny.Seq.of(_586_v7, _dafny.Set.fromElements(false));
          let _589_v10;
          _589_v10 = _dafny.Seq.UnicodeFromString("ltfp");
          let _590_v11;
          _590_v11 = _dafny.MultiSet.fromElements(_this.f4, true);
          let _591_v12;
          _591_v12 = _dafny.Map.Empty.slice().updateUnsafe(_582_cf2,_590_v11);
          let _592_v13;
          let _nw110 = Array((new BigNumber(14)).toNumber());
          _nw110[(_dafny.ZERO).toNumber()] = _this.f4;
          _nw110[(_dafny.ONE).toNumber()] = _this.f4;
          _nw110[(new BigNumber(2)).toNumber()] = (((_585_v6).contains(_module.__default.fm2((_this).f3, globalState))) ? ((_585_v6).get(_module.__default.fm2((_this).f3, globalState))) : (_this.f4));
          _nw110[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_module.__default.fm20(new BigNumber(((((_587_v8).contains(_module.__default.fm2(p0, globalState))) ? ((_587_v8).get(_module.__default.fm2(p0, globalState))) : ((_588_v9)[_module.__default.safeIndex(_572_i0, new BigNumber((_588_v9).length))]))).length), globalState), _589_v10);
          _nw110[(new BigNumber(4)).toNumber()] = (_this.f4) || (_module.__default.fm2(new BigNumber((_591_v12).length), globalState));
          _nw110[(new BigNumber(5)).toNumber()] = false;
          _nw110[(new BigNumber(6)).toNumber()] = _582_cf2;
          _nw110[(new BigNumber(7)).toNumber()] = !(_this.f4);
          _nw110[(new BigNumber(8)).toNumber()] = ((_this).f14).isLessThan(_572_i0);
          _nw110[(new BigNumber(9)).toNumber()] = ((_582_cf2) ? (_583_cf1) : (_582_cf2));
          _nw110[(new BigNumber(10)).toNumber()] = (((_585_v6).contains(_582_cf2)) ? ((_585_v6).get(_582_cf2)) : (_582_cf2));
          _nw110[(new BigNumber(11)).toNumber()] = !(_this.f4);
          _nw110[(new BigNumber(12)).toNumber()] = _this.f4;
          _nw110[(new BigNumber(13)).toNumber()] = !(_module.__default.fm0(globalState)).isEqualTo(p0);
          _592_v13 = _nw110;
          let _593_v14;
          _593_v14 = _dafny.MultiSet.fromElements(_590_v11, _590_v11);
          let _index108 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_592_v13).length));
          (_592_v13)[_index108] = (_593_v14).equals(_dafny.MultiSet.FromArray(_dafny.Seq.of(_590_v11)));
          let _index109 = _module.__default.safeIndex(new BigNumber(435), new BigNumber((_592_v13).length));
          (_592_v13)[_index109] = (_572_i0).isLessThan(_572_i0);
          let _594_v15;
          let _nw111 = new _module.C0();
          _nw111.__ctor((_592_v13)[_module.__default.safeIndex(new BigNumber(435), new BigNumber((_592_v13).length))], _592_v13);
          _594_v15 = _nw111;
          let _index110 = _module.__default.safeIndex(new BigNumber(232), new BigNumber(((_this).f15).length));
          ((_this).f15)[_index110] = (_this).f3;
          let _index111 = _module.__default.safeIndex(new BigNumber(232), new BigNumber(((_this).f15).length));
          let _rhs116 = (_this).f15;
          let _rhs117 = ((_dafny.ZERO).minus((_this).f14)).minus((_dafny.ZERO).minus((_this).f14));
          let _lhs107 = (_this).f15;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(232), new BigNumber(((_this).f15).length));
          _580_cf4 = _rhs116;
          _lhs107[_lhs108] = _rhs117;
        } else if (_source6.is_DC2) {
          let _595___mcc_h4 = (_source6).cf5;
          let _596___mcc_h5 = (_source6).cf6;
          let _597_cf6 = _596___mcc_h5;
          let _598_cf5 = _595___mcc_h4;
          let _599_v16;
          let _nw112 = Array((new BigNumber(18)).toNumber()).fill([]);
          _599_v16 = _nw112;
          let _index112 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_599_v16).length));
          (_599_v16)[_index112] = ((_597_cf6) ? ((_this).f15) : ((_this).f15));
          let _600_v17;
          let _nw113 = Array((new BigNumber(26)).toNumber()).fill(_module.D8.Default());
          _600_v17 = _nw113;
          let _601_v18;
          _601_v18 = _module.D8.create_DC23(false, _597_cf6, _this.f4, _573_v2);
          let _602_v19;
          _602_v19 = _module.D8.create_DC24(_601_v18);
          let _index113 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_600_v17).length));
          (_600_v17)[_index113] = _602_v19;
          let _603_v20;
          _603_v20 = _dafny.MultiSet.fromElements(_572_i0, (_this).f14);
          let _604_v21;
          _604_v21 = _dafny.Seq.of(_597_cf6, _597_cf6);
          let _index114 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_599_v16).length));
          let _index115 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_600_v17).length));
          let _rhs118 = (_this).f15;
          let _rhs119 = _module.__default.fm26(globalState);
          let _rhs120 = (_603_v20).update(new BigNumber((_604_v21).length), _module.__default.abs((_this).f3));
          let _lhs109 = _599_v16;
          let _lhs110 = _module.__default.safeIndex(new BigNumber(570), new BigNumber((_599_v16).length));
          let _lhs111 = _600_v17;
          let _lhs112 = _module.__default.safeIndex(new BigNumber(626), new BigNumber((_600_v17).length));
          _lhs109[_lhs110] = _rhs118;
          _lhs111[_lhs112] = _rhs119;
          _603_v20 = _rhs120;
          let _605_v22;
          _605_v22 = _dafny.Map.Empty.slice().updateUnsafe(_572_i0,_this.f4);
          _605_v22 = (_605_v22).update((_dafny.ZERO).minus((_this).f3), _this.f4);
          r0 = _572_i0;
          r0 = (_this).f3;
        } else if (_source6.is_DC3) {
          let _606_v23;
          _606_v23 = _dafny.Seq.UnicodeFromString("af");
          let _607_v24;
          _607_v24 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(-271)).isLessThanOrEqualTo(_572_i0),_dafny.Seq.of(_606_v23));
          let _608_v25;
          _608_v25 = _dafny.Seq.of(_this.f4);
          let _609_v26;
          _609_v26 = _dafny.Seq.of(_606_v23);
          _607_v24 = (_607_v24).update(!((_this.f4) === ((_608_v25)[_module.__default.safeIndex(_572_i0, new BigNumber((_608_v25).length))])), _609_v26);
          let _610_v27;
          let _nw114 = Array((new BigNumber(17)).toNumber());
          _nw114[(_dafny.ZERO).toNumber()] = _this.f4;
          _nw114[(_dafny.ONE).toNumber()] = true;
          _nw114[(new BigNumber(2)).toNumber()] = _this.f4;
          _nw114[(new BigNumber(3)).toNumber()] = _this.f4;
          _nw114[(new BigNumber(4)).toNumber()] = !(_this.f4);
          _nw114[(new BigNumber(5)).toNumber()] = _this.f4;
          _nw114[(new BigNumber(6)).toNumber()] = _module.__default.fm2(_572_i0, globalState);
          _nw114[(new BigNumber(7)).toNumber()] = _this.f4;
          _nw114[(new BigNumber(8)).toNumber()] = true;
          _nw114[(new BigNumber(9)).toNumber()] = _this.f4;
          _nw114[(new BigNumber(10)).toNumber()] = _this.f4;
          _nw114[(new BigNumber(11)).toNumber()] = false;
          _nw114[(new BigNumber(12)).toNumber()] = true;
          _nw114[(new BigNumber(13)).toNumber()] = false;
          _nw114[(new BigNumber(14)).toNumber()] = _this.f4;
          _nw114[(new BigNumber(15)).toNumber()] = _this.f4;
          _nw114[(new BigNumber(16)).toNumber()] = _this.f4;
          _610_v27 = _nw114;
          let _611_v28;
          _611_v28 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_610_v27);
          _611_v28 = (_611_v28).update(true, _610_v27);
          let _612_v29;
          let _nw115 = new _module.C0();
          _nw115.__ctor(_this.f4, _610_v27);
          _612_v29 = _nw115;
          (_this).f4 = _612_v29.f10;
        } else if (_source6.is_DC0) {
          let _613___mcc_h6 = (_source6).cf0;
          let _614_cf0 = _613___mcc_h6;
          let _615_v30;
          _615_v30 = _dafny.Seq.UnicodeFromString("gjdgtgnld");
          let _616_v31;
          _616_v31 = _dafny.Map.Empty.slice().updateUnsafe(_615_v30,_572_i0);
          let _617_v32;
          _617_v32 = _dafny.Map.Empty.slice().updateUnsafe(_614_cf0,false);
          _616_v31 = (_616_v31).update(_615_v30, new BigNumber((_dafny.MultiSet.fromElements(_614_cf0, (((_617_v32).contains(_this.f4)) ? ((_617_v32).get(_this.f4)) : (true)))).cardinality()));
          let _index116 = _module.__default.safeIndex(new BigNumber(586), new BigNumber(((_this).f15).length));
          ((_this).f15)[_index116] = (_this).f14;
          let _index117 = _module.__default.safeIndex(new BigNumber(586), new BigNumber(((_this).f15).length));
          ((_this).f15)[_index117] = (_this).f14;
          let _618_v33;
          let _nw116 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _618_v33 = _nw116;
          let _index118 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_618_v33).length));
          (_618_v33)[_index118] = _615_v30;
          let _619_v34;
          _619_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,((_this).f15)[_module.__default.safeIndex(new BigNumber(586), new BigNumber(((_this).f15).length))]);
          let _620_v36;
          _620_v36 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,new BigNumber((function () {
            let _coll8 = new _dafny.Set();
            for (const _compr_8 of (_619_v34).Keys.Elements) {
              let _621_v35 = _compr_8;
              if ((_619_v34).contains(_621_v35)) {
                _coll8.add((_621_v35).multipliedBy(new BigNumber(-756)));
              }
            }
            return _coll8;
          }()).length));
          let _622_v37;
          _622_v37 = _dafny.MultiSet.fromElements(_620_v36, _620_v36);
          let _623_v38;
          let _nw117 = Array((new BigNumber(20)).toNumber()).fill(false);
          _623_v38 = _nw117;
          let _624_v39;
          let _nw118 = new _module.C0();
          _nw118.__ctor(_614_cf0, _623_v38);
          _624_v39 = _nw118;
          let _625_v40;
          _625_v40 = _dafny.Map.Empty.slice().updateUnsafe(_624_v39,_615_v30);
          let _626_v41;
          _626_v41 = _dafny.Map.Empty.slice().updateUnsafe((_622_v37).update(_620_v36, _module.__default.abs((_this).f3)),(((_625_v40).contains(_624_v39)) ? ((_625_v40).get(_624_v39)) : (_dafny.Seq.UnicodeFromString("tpoiv"))));
          let _index119 = _module.__default.safeIndex(new BigNumber(836), new BigNumber((_618_v33).length));
          (_618_v33)[_index119] = (((_626_v41).contains(_622_v37)) ? ((_626_v41).get(_622_v37)) : (_615_v30));
          _570_v0 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(4)), function (_627_i1) {
            return (_this).f3;
          });
        } else {
          let _628___mcc_h7 = (_source6).cf7;
          let _629_cf7 = _628___mcc_h7;
          r1 = _module.__default.fm2((_dafny.ZERO).minus((_this).f3), globalState);
          let _630_v42;
          _630_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f14,true);
          _630_v42 = (_630_v42).update((_this).f3, !(((_this).f14).isLessThan(p0)));
          let _631_v43;
          _631_v43 = _dafny.Seq.UnicodeFromString("viwcggtyo");
          let _632_v44;
          _632_v44 = _dafny.Set.fromElements(p0, new BigNumber((_631_v43).length));
          let _633_v45;
          _633_v45 = _dafny.MultiSet.fromElements((_this).f14);
          r0 = (new BigNumber((_632_v44).length)).plus(((_this.f4) ? ((_this).f3) : (new BigNumber((_633_v45).cardinality()))));
          let _634_v46;
          let _init22 = function (_635_i2) {
            return _this.f4;
          };
          let _nw119 = Array((new BigNumber(29)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw119.length); _i0_22++) {
            _nw119[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _634_v46 = _nw119;
          let _636_v47;
          _636_v47 = _module.D7.create_DC21(_634_v46, _this.f4, _this.f4);
          let _pat_let_tv5 = _634_v46;
          let _637_v48;
          let _nw120 = Array((new BigNumber(15)).toNumber());
          _nw120[(_dafny.ZERO).toNumber()] = _module.D7.create_DC21(_634_v46, _this.f4, _this.f4);
          _nw120[(_dafny.ONE).toNumber()] = _636_v47;
          _nw120[(new BigNumber(2)).toNumber()] = _636_v47;
          _nw120[(new BigNumber(3)).toNumber()] = _636_v47;
          _nw120[(new BigNumber(4)).toNumber()] = _636_v47;
          _nw120[(new BigNumber(5)).toNumber()] = _636_v47;
          _nw120[(new BigNumber(6)).toNumber()] = _636_v47;
          _nw120[(new BigNumber(7)).toNumber()] = _636_v47;
          _nw120[(new BigNumber(8)).toNumber()] = _636_v47;
          _nw120[(new BigNumber(9)).toNumber()] = _636_v47;
          _nw120[(new BigNumber(10)).toNumber()] = function (_pat_let8_0) {
            return function (_638_dt__update__tmp_h0) {
              return function (_pat_let9_0) {
                return function (_639_dt__update_hcf29_h0) {
                  return _module.D7.create_DC21((_638_dt__update__tmp_h0).dtor_cf28, _639_dt__update_hcf29_h0, (_638_dt__update__tmp_h0).dtor_cf30);
                }(_pat_let9_0);
              }(_this.f4);
            }(_pat_let8_0);
          }(_636_v47);
          _nw120[(new BigNumber(11)).toNumber()] = _module.D7.create_DC21(_634_v46, _this.f4, _this.f4);
          _nw120[(new BigNumber(12)).toNumber()] = function (_pat_let10_0) {
            return function (_640_dt__update__tmp_h1) {
              return function (_pat_let11_0) {
                return function (_641_dt__update_hcf28_h0) {
                  return _module.D7.create_DC21(_641_dt__update_hcf28_h0, (_640_dt__update__tmp_h1).dtor_cf29, (_640_dt__update__tmp_h1).dtor_cf30);
                }(_pat_let11_0);
              }(_pat_let_tv5);
            }(_pat_let10_0);
          }(_module.D7.create_DC21(_634_v46, false, _this.f4));
          _nw120[(new BigNumber(13)).toNumber()] = _636_v47;
          _nw120[(new BigNumber(14)).toNumber()] = _module.D7.create_DC21(_634_v46, _this.f4, _this.f4);
          _637_v48 = _nw120;
          let _642_v49;
          _642_v49 = _dafny.Set.fromElements(_637_v48, _637_v48);
          let _643_v50;
          _643_v50 = _dafny.Seq.of(_dafny.Set.fromElements(_637_v48, _637_v48));
          _642_v49 = ((_643_v50)[_module.__default.safeIndex(p0, new BigNumber((_643_v50).length))]).Union(_642_v49);
        }
        let _644_v52;
        _644_v52 = _module.D7.create_DC20(function () {
  let _coll9 = new _dafny.Map();
  for (const _compr_9 of _dafny.IntegerRange(new BigNumber(353), new BigNumber(82))) {
    let _645_v51 = _compr_9;
    if (((new BigNumber(353)).isLessThanOrEqualTo(_645_v51)) && ((_645_v51).isLessThan(new BigNumber(82)))) {
      _coll9.push([_module.__default.safeModuloInt(_645_v51, p0),_dafny.Seq.UnicodeFromString("xoulfc")]);
    }
  }
  return _coll9;
}());
        let _source7 = _644_v52;
        if (_source7.is_DC21) {
          let _646___mcc_h8 = (_source7).cf28;
          let _647___mcc_h9 = (_source7).cf29;
          let _648___mcc_h10 = (_source7).cf30;
          let _649_cf30 = _648___mcc_h10;
          let _650_cf29 = _647___mcc_h9;
          let _651_cf28 = _646___mcc_h8;
          (_this).f4 = _650_cf29;
          let _652_v53;
          _652_v53 = _dafny.Map.Empty.slice().updateUnsafe(p0,true);
          let _index120 = _module.__default.safeIndex(new BigNumber(354), new BigNumber(((_this).f15).length));
          ((_this).f15)[_index120] = (p0).plus(new BigNumber((_652_v53).length));
          let _653_v54;
          _653_v54 = _module.D7.create_DC21(_651_cf28, _649_cf30, !(_650_cf29));
          let _654_v55;
          _654_v55 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_572_i0);
          let _655_v56;
          _655_v56 = _module.D3.create_DC10((_572_i0).plus((_dafny.ZERO).minus(_572_i0)), (_653_v54).dtor_cf28, (((_654_v55).contains(_649_cf30)) ? ((_654_v55).get(_649_cf30)) : ((_this).f14)));
          let _656_v57;
          _656_v57 = _dafny.Seq.of(_651_cf28);
          let _index121 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((p1).length));
          (p1)[_index121] = (_656_v57)[_module.__default.safeIndex((_this).f3, new BigNumber((_656_v57).length))];
          let _index122 = _module.__default.safeIndex(new BigNumber(354), new BigNumber(((_this).f15).length));
          let _index123 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((p1).length));
          let _rhs121 = new BigNumber(-163);
          let _rhs122 = ((_module.__default.fm2((_this).f3, globalState)) ? (_module.__default.fm2(_572_i0, globalState)) : (true));
          let _rhs123 = p0;
          let _rhs124 = _655_v56;
          let _rhs125 = _651_cf28;
          let _lhs113 = globalState;
          let _lhs114 = globalState;
          let _lhs115 = (_this).f15;
          let _lhs116 = _module.__default.safeIndex(new BigNumber(354), new BigNumber(((_this).f15).length));
          let _lhs117 = p1;
          let _lhs118 = _module.__default.safeIndex(new BigNumber(686), new BigNumber((p1).length));
          _lhs113.f1 = _rhs121;
          _lhs114.f2 = _rhs122;
          _lhs115[_lhs116] = _rhs123;
          _655_v56 = _rhs124;
          _lhs117[_lhs118] = _rhs125;
          let _index124 = _module.__default.safeIndex(new BigNumber(354), new BigNumber(((_this).f15).length));
          ((_this).f15)[_index124] = _module.__default.safeDivisionInt(_module.__default.fm0(globalState), ((_650_cf29) ? (((_this).f15)[_module.__default.safeIndex(new BigNumber(354), new BigNumber(((_this).f15).length))]) : ((_this).f14)));
          let _657_v58;
          let _nw121 = new _module.C0();
          _nw121.__ctor(_649_cf30, (p1)[_module.__default.safeIndex(new BigNumber(686), new BigNumber((p1).length))]);
          _657_v58 = _nw121;
        } else {
          let _658___mcc_h11 = (_source7).cf27;
          let _659_cf27 = _658___mcc_h11;
          let _660_v59;
          _660_v59 = _dafny.Seq.of(_this.f4, false);
          r0 = new BigNumber((_dafny.Seq.Concat(_660_v59, _dafny.Seq.Concat(_660_v59, _660_v59))).length);
          let _661_v60;
          _661_v60 = _dafny.Seq.UnicodeFromString("yqnnbe");
          r0 = new BigNumber((_661_v60).length);
          let _662_v61;
          _662_v61 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_this.f4);
          let _663_v62;
          _663_v62 = _module.D6.create_DC18(_this.f4, p0);
          let _664_v63;
          _664_v63 = _dafny.Set.fromElements(_572_i0, _572_i0, new BigNumber((_662_v61).length), (_663_v62).dtor_cf25, _572_i0);
          let _665_v64;
          let _nw122 = Array((new BigNumber(15)).toNumber());
          _nw122[(_dafny.ZERO).toNumber()] = _this.f4;
          _nw122[(_dafny.ONE).toNumber()] = (new BigNumber((_664_v63).length)).isLessThan(new BigNumber(656));
          _nw122[(new BigNumber(2)).toNumber()] = _this.f4;
          _nw122[(new BigNumber(3)).toNumber()] = (_662_v61).contains(_module.__default.fm2((_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.fm0(globalState))), globalState));
          _nw122[(new BigNumber(4)).toNumber()] = ((_this.f4) ? (_this.f4) : (!(true)));
          _nw122[(new BigNumber(5)).toNumber()] = ((_this).f3).isLessThanOrEqualTo(p0);
          _nw122[(new BigNumber(6)).toNumber()] = _this.f4;
          _nw122[(new BigNumber(7)).toNumber()] = _this.f4;
          _nw122[(new BigNumber(8)).toNumber()] = !(false);
          _nw122[(new BigNumber(9)).toNumber()] = (_662_v61).contains(!(_this.f4));
          _nw122[(new BigNumber(10)).toNumber()] = _this.f4;
          _nw122[(new BigNumber(11)).toNumber()] = _this.f4;
          _nw122[(new BigNumber(12)).toNumber()] = _this.f4;
          _nw122[(new BigNumber(13)).toNumber()] = _this.f4;
          _nw122[(new BigNumber(14)).toNumber()] = _this.f4;
          _665_v64 = _nw122;
          let _index125 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_665_v64).length));
          (_665_v64)[_index125] = _this.f4;
          let _index126 = _module.__default.safeIndex(new BigNumber(330), new BigNumber((_665_v64).length));
          (_665_v64)[_index126] = _this.f4;
          _665_v64 = _665_v64;
        }
        let _666_v65;
        _666_v65 = _dafny.Seq.UnicodeFromString("png");
        let _667_v66;
        _667_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_666_v65).length),_this.f4);
        let _668_v67;
        let _nw123 = new _module.C1();
        _nw123.__ctor((_dafny.ZERO).minus(_572_i0), new BigNumber((_667_v66).length), _this.f4);
        _668_v67 = _nw123;
      }
      let _669_v68;
      _669_v68 = new _dafny.CodePoint('l'.codePointAt(0));
      let _pat_let_tv6 = _669_v68;
      let _pat_let_tv7 = _669_v68;
      let _pat_let_tv8 = _669_v68;
      _669_v68 = function (_source8) {
        if (_source8.is_DC6) {
          let _670___mcc_h12 = (_source8).cf9;
          let _671___mcc_h13 = (_source8).cf10;
          let _672___mcc_h14 = (_source8).cf11;
          let _673___mcc_h15 = (_source8).cf12;
          let _674_cf12 = _673___mcc_h15;
          let _675_cf11 = _672___mcc_h14;
          let _676_cf10 = _671___mcc_h13;
          let _677_cf9 = _670___mcc_h12;
          if (_674_cf12) {
            return _pat_let_tv6;
          } else {
            return new _dafny.CodePoint('c'.codePointAt(0));
          }
        } else if (_source8.is_DC7) {
          return _pat_let_tv7;
        } else {
          let _678___mcc_h16 = (_source8).cf8;
          let _679_cf8 = _678___mcc_h16;
          return _pat_let_tv8;
        }
      }(_module.D1.create_DC5((_this).f14));
      r0 = (p0).multipliedBy((_this).f14);
      (globalState).f2 = _this.f4;
      r0 = (((new BigNumber(906)).isLessThan(p0)) ? ((_this).f14) : (_module.__default.safeDivisionInt((_this).f14, (_dafny.ZERO).minus((_this).f3))));
      r1 = _module.__default.fm2(_module.__default.safeModuloInt((_this).f14, p0), globalState);
      return [r0, r1];
    }
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f15() {
      let _this = this;
      return _this._f15;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f4 = false;
      this._f7 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this.f13 = _dafny.Seq.UnicodeFromString("");
      this._f12 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f12, f13, f3, f4, f7) {
      let _this = this;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f7 = f7;
      return;
    }
    fm12(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((_this).f7, new BigNumber(14));
    };
    fm13(globalState) {
      let _this = this;
      return _this.f4;
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _680_v0;
      _680_v0 = _dafny.Set.fromElements(((_this).f7).minus((_this).f7), (_this).f7, new BigNumber((_this.f13).length));
      _680_v0 = _dafny.Set.fromElements(_module.__default.safeDivisionInt((_this).f3, (_this).f7));
      let _681_v1;
      _681_v1 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.__default.fm14(!(_this.f4), (_this).fm13(globalState), globalState)),_this.f4);
      let _682_v2;
      _682_v2 = _dafny.MultiSet.fromElements(true);
      (globalState).f2 = (((_681_v1).contains(_dafny.Seq.of(_dafny.MultiSet.fromElements(_this.f4), _682_v2))) ? ((_681_v1).get(_dafny.Seq.of(_dafny.MultiSet.fromElements(_this.f4), _682_v2))) : (((_this).f3).isLessThan(p0)));
      let _683_i0;
      _683_i0 = _dafny.ZERO;
      L6: {
        while (_this.f4) {
          C6: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_683_i0)) {
              break L6;
            }
            _683_i0 = (_683_i0).plus(_dafny.ONE);
            (_this).m5(globalState);
            (globalState).f2 = !((_this).fm13(globalState));
            if (_this.f4) {
              let _684_v3;
              _684_v3 = _module.D0.create_DC0(_this.f4);
              let _685_v4;
              let _nw124 = Array((new BigNumber(23)).toNumber());
              _nw124[(_dafny.ZERO).toNumber()] = !(false);
              _nw124[(_dafny.ONE).toNumber()] = _this.f4;
              _nw124[(new BigNumber(2)).toNumber()] = (_684_v3).dtor_cf0;
              _nw124[(new BigNumber(3)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(4)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(5)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(6)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(7)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(8)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(9)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(10)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(11)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(12)).toNumber()] = false;
              _nw124[(new BigNumber(13)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(14)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(15)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(16)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(17)).toNumber()] = false;
              _nw124[(new BigNumber(18)).toNumber()] = false;
              _nw124[(new BigNumber(19)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(20)).toNumber()] = _this.f4;
              _nw124[(new BigNumber(21)).toNumber()] = false;
              _nw124[(new BigNumber(22)).toNumber()] = _this.f4;
              _685_v4 = _nw124;
              let _686_v5;
              let _nw125 = new _module.C0();
              _nw125.__ctor(((!(_this.f4)) ? (_this.f4) : (_this.f4)), _685_v4);
              _686_v5 = _nw125;
              let _687_v6;
              _687_v6 = _dafny.Seq.of((_this).f3);
              let _688_v7;
              let _nw126 = new _module.C0();
              _nw126.__ctor((new BigNumber((_dafny.Seq.update(_module.__default.fm15(!(_686_v5.f10), (_this).f3, new BigNumber((_687_v6).length), globalState), _module.__default.safeIndex(p0, new BigNumber((_module.__default.fm15(!(_686_v5.f10), (_this).f3, new BigNumber((_687_v6).length), globalState)).length)), (_this).f12)).length)).isEqualTo(p0), (_686_v5).f11);
              _688_v7 = _nw126;
              let _689_v8;
              _689_v8 = _module.D4.create_DC11(_680_v0);
              let _690_v9;
              _690_v9 = _module.D4.create_DC11((_689_v8).dtor_cf18);
              _680_v0 = (_689_v8).dtor_cf18;
              let _691_v10;
              let _init23 = ((_692_p0) => function (_693_i1) {
                return (_693_i1).multipliedBy(_692_p0);
              })(p0);
              let _nw127 = Array((new BigNumber(5)).toNumber());
              for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw127.length); _i0_23++) {
                _nw127[_i0_23] = _init23(new BigNumber(_i0_23));
              }
              _691_v10 = _nw127;
              _691_v10 = _691_v10;
              let _694_v11;
              _694_v11 = _dafny.MultiSet.fromElements((_688_v7).f11, (_688_v7).f11, (_688_v7).f11, (_686_v5).f11);
              let _695_v12;
              _695_v12 = _dafny.Map.Empty.slice().updateUnsafe(_688_v7.f10,(_688_v7).fm5(new BigNumber(292), p0, (_this).f12, _688_v7.f10, globalState));
              let _696_v13;
              _696_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_695_v12).length),_686_v5.f10);
              let _rhs126 = _module.__default.safeDivisionInt((_this).f3, (new BigNumber((_696_v13).length)).minus(p0));
              let _rhs127 = (_694_v11).Difference(_694_v11);
              let _rhs128 = (_this).f7;
              let _rhs129 = _691_v10;
              let _rhs130 = _687_v6;
              let _lhs119 = globalState;
              r0 = _rhs126;
              _694_v11 = _rhs127;
              _lhs119.f1 = _rhs128;
              _691_v10 = _rhs129;
              _687_v6 = _rhs130;
            } else {
              let _697_v14;
              let _nw128 = Array((new BigNumber(7)).toNumber());
              _nw128[(_dafny.ZERO).toNumber()] = (_this).f12;
              _nw128[(_dafny.ONE).toNumber()] = (_this).f12;
              _nw128[(new BigNumber(2)).toNumber()] = (_this).f12;
              _nw128[(new BigNumber(3)).toNumber()] = (_this).f12;
              _nw128[(new BigNumber(4)).toNumber()] = (_this).f12;
              _nw128[(new BigNumber(5)).toNumber()] = (_this).f12;
              _nw128[(new BigNumber(6)).toNumber()] = (_this).f12;
              _697_v14 = _nw128;
              let _698_v15;
              _698_v15 = _dafny.Seq.of(_697_v14);
              let _699_v16;
              _699_v16 = _dafny.MultiSet.fromElements((_this).f3);
              let _700_v17;
              let _nw129 = Array((new BigNumber(13)).toNumber());
              _nw129[(_dafny.ZERO).toNumber()] = _697_v14;
              _nw129[(_dafny.ONE).toNumber()] = _697_v14;
              _nw129[(new BigNumber(2)).toNumber()] = _697_v14;
              _nw129[(new BigNumber(3)).toNumber()] = (_698_v15)[_module.__default.safeIndex(new BigNumber((_699_v16).cardinality()), new BigNumber((_698_v15).length))];
              _nw129[(new BigNumber(4)).toNumber()] = _697_v14;
              _nw129[(new BigNumber(5)).toNumber()] = (_698_v15)[_module.__default.safeIndex((_this).f7, new BigNumber((_698_v15).length))];
              _nw129[(new BigNumber(6)).toNumber()] = _697_v14;
              _nw129[(new BigNumber(7)).toNumber()] = _697_v14;
              _nw129[(new BigNumber(8)).toNumber()] = _697_v14;
              _nw129[(new BigNumber(9)).toNumber()] = _697_v14;
              _nw129[(new BigNumber(10)).toNumber()] = _697_v14;
              _nw129[(new BigNumber(11)).toNumber()] = _697_v14;
              _nw129[(new BigNumber(12)).toNumber()] = _697_v14;
              _700_v17 = _nw129;
              let _index127 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_700_v17).length));
              (_700_v17)[_index127] = _697_v14;
              let _701_v18;
              _701_v18 = _module.D3.create_DC9(_697_v14);
              let _index128 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_700_v17).length));
              let _rhs131 = (_701_v18).dtor_cf14;
              let _rhs132 = (_this.f4) || (_this.f4);
              let _lhs120 = _700_v17;
              let _lhs121 = _module.__default.safeIndex(new BigNumber(448), new BigNumber((_700_v17).length));
              let _lhs122 = globalState;
              _lhs120[_lhs121] = _rhs131;
              _lhs122.f2 = _rhs132;
              let _702_v19;
              let _nw130 = Array((new BigNumber(11)).toNumber()).fill(false);
              _702_v19 = _nw130;
              let _index129 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_702_v19).length));
              (_702_v19)[_index129] = _this.f4;
              let _703_v20;
              _703_v20 = _dafny.Seq.of(p0);
              let _704_v21;
              _704_v21 = _module.D1.create_DC5((_this).f7);
              let _705_v22;
              _705_v22 = _dafny.Map.Empty.slice().updateUnsafe(_704_v21,(_this).f3);
              let _706_v23;
              _706_v23 = _module.D1.create_DC5((((_705_v22).contains(_704_v21)) ? ((_705_v22).get(_704_v21)) : ((_this).f7)));
              let _707_v24;
              _707_v24 = _dafny.Seq.of(_this.f4, _this.f4, _this.f4);
              let _708_v25;
              _708_v25 = _dafny.Map.Empty.slice().updateUnsafe((_707_v24)[_module.__default.safeIndex(p0, new BigNumber((_707_v24).length))],p0);
              let _709_v26;
              _709_v26 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,_708_v25);
              let _710_v27;
              let _nw131 = Array((new BigNumber(21)).toNumber());
              _nw131[(_dafny.ZERO).toNumber()] = (_this).f7;
              _nw131[(_dafny.ONE).toNumber()] = (_this).f3;
              _nw131[(new BigNumber(2)).toNumber()] = (_this).f7;
              _nw131[(new BigNumber(3)).toNumber()] = (_this).f3;
              _nw131[(new BigNumber(4)).toNumber()] = p0;
              _nw131[(new BigNumber(5)).toNumber()] = ((_this).f7).minus(new BigNumber((_703_v20).length));
              _nw131[(new BigNumber(6)).toNumber()] = p0;
              _nw131[(new BigNumber(7)).toNumber()] = ((_this).f7).minus(new BigNumber(166));
              _nw131[(new BigNumber(8)).toNumber()] = (_this).f7;
              _nw131[(new BigNumber(9)).toNumber()] = (_704_v21).dtor_cf8;
              _nw131[(new BigNumber(10)).toNumber()] = _module.__default.fm0(globalState);
              _nw131[(new BigNumber(11)).toNumber()] = new BigNumber(-367);
              _nw131[(new BigNumber(12)).toNumber()] = (new BigNumber(407)).multipliedBy((_this).f3);
              _nw131[(new BigNumber(13)).toNumber()] = new BigNumber((_709_v26).length);
              _nw131[(new BigNumber(14)).toNumber()] = _module.__default.safeModuloInt((_module.D1.create_DC5((_this).f7)).dtor_cf8, p0);
              _nw131[(new BigNumber(15)).toNumber()] = new BigNumber(-272);
              _nw131[(new BigNumber(16)).toNumber()] = p0;
              _nw131[(new BigNumber(17)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_707_v24).length), (_this).f3);
              _nw131[(new BigNumber(18)).toNumber()] = (_this).f7;
              _nw131[(new BigNumber(19)).toNumber()] = _module.__default.safeModuloInt((_this).f7, (_this).f7);
              _nw131[(new BigNumber(20)).toNumber()] = (_this).f3;
              _710_v27 = _nw131;
              let _index130 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_702_v19).length));
              let _rhs133 = _this.f4;
              let _rhs134 = _710_v27;
              let _lhs123 = _702_v19;
              let _lhs124 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_702_v19).length));
              _lhs123[_lhs124] = _rhs133;
              _710_v27 = _rhs134;
              let _711_v28;
              let _nw132 = new _module.C0();
              _nw132.__ctor((_702_v19)[_module.__default.safeIndex(new BigNumber(379), new BigNumber((_702_v19).length))], _702_v19);
              _711_v28 = _nw132;
              (globalState).f1 = p0;
              let _712_v29;
              let _init24 = ((_713_v0) => function (_714_i2) {
                return _713_v0;
              })(_680_v0);
              let _nw133 = Array((new BigNumber(2)).toNumber());
              for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw133.length); _i0_24++) {
                _nw133[_i0_24] = _init24(new BigNumber(_i0_24));
              }
              _712_v29 = _nw133;
              let _715_v30;
              let _nw134 = Array((new BigNumber(3)).toNumber()).fill(_module.D3.Default());
              _715_v30 = _nw134;
              let _716_v31;
              _716_v31 = _module.D3.create_DC10((_dafny.ZERO).minus(p0), (_711_v28).f11, (_this).f7);
              let _index131 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_715_v30).length));
              (_715_v30)[_index131] = _716_v31;
              let _717_v32;
              _717_v32 = _dafny.Seq.of(_this.f13);
              let _718_v33;
              _718_v33 = _dafny.MultiSet.fromElements(_this.f13, _this.f13);
              let _pat_let_tv9 = _702_v19;
              let _index132 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_702_v19).length));
              let _index133 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_715_v30).length));
              let _rhs135 = _712_v29;
              let _rhs136 = (new BigNumber((((true) ? (_dafny.MultiSet.FromArray(_717_v32)) : (_718_v33))).cardinality())).isLessThanOrEqualTo((_this).f7);
              let _rhs137 = (p0).isEqualTo(_module.__default.safeModuloInt(p0, (_this).f7));
              let _rhs138 = function (_pat_let12_0) {
                return function (_719_dt__update__tmp_h0) {
                  return function (_pat_let13_0) {
                    return function (_720_dt__update_hcf15_h0) {
                      return function (_pat_let14_0) {
                        return function (_721_dt__update_hcf16_h0) {
                          return _module.D3.create_DC10(_720_dt__update_hcf15_h0, _721_dt__update_hcf16_h0, (_719_dt__update__tmp_h0).dtor_cf17);
                        }(_pat_let14_0);
                      }(_pat_let_tv9);
                    }(_pat_let13_0);
                  }((new BigNumber(997)).plus((_this).f3));
                }(_pat_let12_0);
              }(_716_v31);
              let _lhs125 = _702_v19;
              let _lhs126 = _module.__default.safeIndex(new BigNumber(379), new BigNumber((_702_v19).length));
              let _lhs127 = globalState;
              let _lhs128 = _715_v30;
              let _lhs129 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_715_v30).length));
              _712_v29 = _rhs135;
              _lhs125[_lhs126] = _rhs136;
              _lhs127.f2 = _rhs137;
              _lhs128[_lhs129] = _rhs138;
            }
            let _722_v34;
            let _nw135 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
            _722_v34 = _nw135;
            let _index134 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_722_v34).length));
            (_722_v34)[_index134] = p0;
            let _index135 = _module.__default.safeIndex(new BigNumber(57), new BigNumber((_722_v34).length));
            (_722_v34)[_index135] = (_this).f7;
          }
        }
      }
      let _723_i3;
      _723_i3 = _dafny.ZERO;
      L7: {
        while (_this.f4) {
          C7: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_723_i3)) {
              break L7;
            }
            _723_i3 = (_723_i3).plus(_dafny.ONE);
            let _724_v35;
            _724_v35 = _dafny.Seq.of(p0);
            let _725_v36;
            _725_v36 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f4),_724_v35);
            _725_v36 = (_725_v36).update(!(_this.f4), _724_v35);
            _680_v0 = _680_v0;
            let _726_v37;
            _726_v37 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,new BigNumber(637));
            let _727_v38;
            let _nw136 = Array((new BigNumber(5)).toNumber());
            _nw136[(_dafny.ZERO).toNumber()] = (_726_v37).contains(!(_this.f4));
            _nw136[(_dafny.ONE).toNumber()] = true;
            _nw136[(new BigNumber(2)).toNumber()] = _this.f4;
            _nw136[(new BigNumber(3)).toNumber()] = true;
            _nw136[(new BigNumber(4)).toNumber()] = _this.f4;
            _727_v38 = _nw136;
            let _index136 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_727_v38).length));
            (_727_v38)[_index136] = !(_this.f4);
            let _728_v39;
            _728_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(546),(_this).f3);
            let _index137 = _module.__default.safeIndex(new BigNumber(436), new BigNumber((_727_v38).length));
            (_727_v38)[_index137] = ((_this).f7).isEqualTo(new BigNumber((_728_v39).length));
            let _729_v40;
            let _nw137 = new _module.C0();
            _nw137.__ctor(((_this).f3).isLessThanOrEqualTo((_this).f7), _727_v38);
            _729_v40 = _nw137;
          }
        }
      }
      (globalState).f2 = false;
      let _730_v41;
      _730_v41 = _module.D1.create_DC7();
      let _rhs139 = (_this).fm12(_module.__default.fm2(new BigNumber((_680_v0).length), globalState), _this.f4, globalState);
      let _rhs140 = function (_source9) {
        if (_source9.is_DC6) {
          let _731___mcc_h0 = (_source9).cf9;
          let _732___mcc_h1 = (_source9).cf10;
          let _733___mcc_h2 = (_source9).cf11;
          let _734___mcc_h3 = (_source9).cf12;
          let _735_cf12 = _734___mcc_h3;
          let _736_cf11 = _733___mcc_h2;
          let _737_cf10 = _732___mcc_h1;
          let _738_cf9 = _731___mcc_h0;
          if (!(false)) {
            return false;
          } else {
            return true;
          }
        } else if (_source9.is_DC7) {
          return _this.f4;
        } else {
          let _739___mcc_h4 = (_source9).cf8;
          let _740_cf8 = _739___mcc_h4;
          return (_dafny.MultiSet.fromElements((_this).f3)).contains(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(300)), ((_741_cf8) => function (_742_i4) {
            return _741_cf8;
          })(_740_cf8))).length));
        }
      }(_730_v41);
      let _rhs141 = ((p0).minus(new BigNumber((_this.f13).length))).multipliedBy((_this).f3);
      let _lhs130 = globalState;
      let _lhs131 = globalState;
      let _lhs132 = globalState;
      _lhs130.f1 = _rhs139;
      _lhs131.f2 = _rhs140;
      _lhs132.f1 = _rhs141;
      let _743_v42;
      _743_v42 = _dafny.Map.Empty.slice().updateUnsafe(_682_v2,(_this).f7);
      r0 = _module.__default.safeModuloInt(_module.__default.safeDivisionInt((((_743_v42).contains(_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f4)))) ? ((_743_v42).get(_dafny.MultiSet.FromArray(_dafny.Seq.of(_this.f4)))) : ((_this).f3)), new BigNumber((_680_v0).length)), (_this).f3);
      return r0;
    }
    m5(globalState) {
      let _this = this;
      let _hi6 = _module.__default.safeDivisionInt((_this).f7, (_this).f3);
      for (let _744_i0 = (_this).f7; _744_i0.isLessThan(_hi6); _744_i0 = _744_i0.plus(_dafny.ONE)) {
        let _745_v0;
        let _nw138 = Array((new BigNumber(25)).toNumber()).fill(false);
        _745_v0 = _nw138;
        let _746_v1;
        let _nw139 = new _module.C0();
        _nw139.__ctor(!(((_dafny.ZERO).minus((_this).f7)).isLessThan(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(591)), function (_747_i1) {
          return (_this).f12;
        })).length))), _745_v0);
        _746_v1 = _nw139;
        (globalState).f1 = (_this).f3;
        let _748_v2;
        _748_v2 = _dafny.Seq.of(_746_v1.f10);
        let _749_v3;
        _749_v3 = _dafny.MultiSet.fromElements(_748_v2, _748_v2);
        (globalState).f1 = new BigNumber((((_749_v3).Intersect(_749_v3)).Difference((_module.__default.fm16(_this.f4, _this.f4, globalState)).Difference(_module.__default.fm16(_746_v1.f10, true, globalState)))).cardinality());
        let _750_v4;
        _750_v4 = _dafny.Map.Empty.slice().updateUnsafe(_746_v1.f10,_this.f4);
        let _751_v5;
        let _nw140 = new _module.C0();
        _nw140.__ctor((((_750_v4).contains(_this.f4)) ? ((_750_v4).get(_this.f4)) : (_746_v1.f10)), (_746_v1).f11);
        _751_v5 = _nw140;
      }
      let _752_v6;
      let _nw141 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
      _752_v6 = _nw141;
      let _753_v7;
      _753_v7 = _module.D0.create_DC0(_dafny.areEqual(_dafny.Seq.of(_752_v6, _752_v6, _752_v6), _dafny.Seq.of(_752_v6)));
      let _source10 = _753_v7;
      if (_source10.is_DC1) {
        let _754___mcc_h0 = (_source10).cf1;
        let _755___mcc_h1 = (_source10).cf2;
        let _756___mcc_h2 = (_source10).cf3;
        let _757___mcc_h3 = (_source10).cf4;
        let _758_cf4 = _757___mcc_h3;
        let _759_cf3 = _756___mcc_h2;
        let _760_cf2 = _755___mcc_h1;
        let _761_cf1 = _754___mcc_h0;
        (globalState).f1 = new BigNumber(-77);
        let _762_v8;
        let _init25 = ((_763_cf1) => function (_764_i2) {
          return _763_cf1;
        })(_761_cf1);
        let _nw142 = Array((new BigNumber(11)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw142.length); _i0_25++) {
          _nw142[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _762_v8 = _nw142;
        let _index138 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_762_v8).length));
        (_762_v8)[_index138] = _760_cf2;
        let _765_v9;
        _765_v9 = _dafny.Seq.of(_752_v6);
        let _766_v10;
        _766_v10 = _dafny.MultiSet.fromElements((_765_v9)[_module.__default.safeIndex((_this).f3, new BigNumber((_765_v9).length))]);
        let _767_v11;
        _767_v11 = _dafny.Map.Empty.slice().updateUnsafe(!(_760_cf2),new BigNumber((_dafny.Seq.update(_this.f13, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_766_v10).cardinality())), new BigNumber((_this.f13).length)), (_this).f12)).length));
        let _index139 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_762_v8).length));
        let _rhs142 = _761_cf1;
        let _rhs143 = _767_v11;
        let _lhs133 = _762_v8;
        let _lhs134 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_762_v8).length));
        _lhs133[_lhs134] = _rhs142;
        _767_v11 = _rhs143;
        (globalState).f2 = ((_this).f7).isLessThanOrEqualTo(((_this).f7).minus((_this).f7));
        let _index140 = _module.__default.safeIndex(new BigNumber(967), new BigNumber((_762_v8).length));
        (_762_v8)[_index140] = _760_cf2;
      } else if (_source10.is_DC2) {
        let _768___mcc_h4 = (_source10).cf5;
        let _769___mcc_h5 = (_source10).cf6;
        let _770_cf6 = _769___mcc_h5;
        let _771_cf5 = _768___mcc_h4;
        let _772_v12;
        _772_v12 = _dafny.Seq.of(_770_cf6, _770_cf6);
        let _773_v13;
        _773_v13 = _dafny.Seq.of(_dafny.Seq.of(_this.f4), _dafny.Seq.update(_772_v12, _module.__default.safeIndex((_this).f7, new BigNumber((_772_v12).length)), _this.f4), _772_v12, _772_v12);
        let _774_v14;
        _774_v14 = _module.D0.create_DC1(false, _this.f4, (_this).f12, _752_v6);
        let _775_v15;
        _775_v15 = _dafny.Map.Empty.slice().updateUnsafe((_773_v13)[_module.__default.safeIndex(new BigNumber((_module.__default.fm15(_770_cf6, (_dafny.ZERO).minus((_this).f3), new BigNumber((_this.f13).length), globalState)).length), new BigNumber((_773_v13).length))],_dafny.Seq.Concat(_this.f13, _dafny.Seq.update(_this.f13, _module.__default.safeIndex((_dafny.ZERO).minus((_this).f3), new BigNumber((_this.f13).length)), (_774_v14).dtor_cf3)));
        _775_v15 = _775_v15;
        (_this).f4 = _770_cf6;
        let _776_v16;
        let _init26 = ((_777_cf6) => function (_778_i3) {
          return _777_cf6;
        })(_770_cf6);
        let _nw143 = Array((new BigNumber(27)).toNumber());
        for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw143.length); _i0_26++) {
          _nw143[_i0_26] = _init26(new BigNumber(_i0_26));
        }
        _776_v16 = _nw143;
        let _779_v17;
        let _nw144 = new _module.C0();
        _nw144.__ctor(_this.f4, _776_v16);
        _779_v17 = _nw144;
        let _780_v18;
        _780_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_770_cf6);
        let _781_v20;
        _781_v20 = _dafny.Seq.of(_780_v18, (function () {
          let _coll10 = new _dafny.Map();
          for (const _compr_10 of (_this.f13).Elements) {
            let _782_v19 = _compr_10;
            if (_dafny.Seq.contains(_this.f13, _782_v19)) {
              _coll10.push([_782_v19,_779_v17.f10]);
            }
          }
          return _coll10;
        }()).update((_this).f12, _this.f4), _module.__default.fm17(_dafny.Map.Empty.slice().updateUnsafe(_779_v17.f10,_779_v17.f10), globalState), _780_v18, _780_v18);
        (globalState).f1 = (new BigNumber(((_781_v20)[_module.__default.safeIndex((_this).f7, new BigNumber((_781_v20).length))]).length)).plus((_this).f3);
      } else if (_source10.is_DC3) {
        let _783_v21;
        let _nw145 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _783_v21 = _nw145;
        let _784_v22;
        _784_v22 = _module.D3.create_DC9(_783_v21);
        let _785_v23;
        _785_v23 = _dafny.Seq.of(_784_v22, _784_v22, _784_v22, _784_v22);
        let _786_v24;
        let _nw146 = Array((new BigNumber(29)).toNumber());
        _nw146[(_dafny.ZERO).toNumber()] = _785_v23;
        _nw146[(_dafny.ONE).toNumber()] = _dafny.Seq.of(_784_v22);
        _nw146[(new BigNumber(2)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_785_v23, _785_v23);
        _nw146[(new BigNumber(4)).toNumber()] = _dafny.Seq.of(_784_v22, _784_v22, _784_v22);
        _nw146[(new BigNumber(5)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(6)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(7)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(8)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(9)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_785_v23, _785_v23), _module.__default.safeIndex((_this).f7, new BigNumber((_dafny.Seq.Concat(_785_v23, _785_v23)).length)), _784_v22);
        _nw146[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_784_v22);
        _nw146[(new BigNumber(12)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(13)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(14)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(15)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(_784_v22, _784_v22);
        _nw146[(new BigNumber(17)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(18)).toNumber()] = _dafny.Seq.of(_784_v22);
        _nw146[(new BigNumber(19)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(20)).toNumber()] = _dafny.Seq.Concat(_785_v23, _785_v23);
        _nw146[(new BigNumber(21)).toNumber()] = _dafny.Seq.of(_784_v22);
        _nw146[(new BigNumber(22)).toNumber()] = _dafny.Seq.of(_784_v22, _784_v22);
        _nw146[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_785_v23, _785_v23);
        _nw146[(new BigNumber(24)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(25)).toNumber()] = _dafny.Seq.of(_module.D3.create_DC9(_783_v21));
        _nw146[(new BigNumber(26)).toNumber()] = _dafny.Seq.of(_784_v22, _module.D3.create_DC9(_783_v21), _784_v22);
        _nw146[(new BigNumber(27)).toNumber()] = _785_v23;
        _nw146[(new BigNumber(28)).toNumber()] = _785_v23;
        _786_v24 = _nw146;
        let _index141 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_786_v24).length));
        (_786_v24)[_index141] = _785_v23;
        let _index142 = _module.__default.safeIndex(new BigNumber(83), new BigNumber((_786_v24).length));
        (_786_v24)[_index142] = ((_this.f4) ? (_785_v23) : (_dafny.Seq.Concat(_785_v23, _785_v23)));
        let _787_v25;
        _787_v25 = _dafny.Map.Empty.slice().updateUnsafe(true,(_this).f7);
        let _788_v26;
        _788_v26 = _dafny.Seq.of((_787_v25).Merge(_787_v25), _787_v25);
        let _789_v27;
        _789_v27 = _dafny.Seq.of(false);
        let _790_v28;
        _790_v28 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(globalState),_this.f4);
        let _791_v30;
        _791_v30 = _dafny.MultiSet.fromElements((_this).f12, (_this).f12);
        let _792_v31;
        _792_v31 = _module.D0.create_DC2(_791_v30, (_module.D0.create_DC2(_791_v30, _this.f4)).dtor_cf6);
        let _793_v32;
        let _nw147 = new _module.C2();
        _nw147.__ctor(new BigNumber(704), _752_v6, (_this).f7, _this.f4);
        _793_v32 = _nw147;
        let _794_v33;
        _794_v33 = _dafny.Set.fromElements(false);
        let _795_v34;
        _795_v34 = _module.D8.create_DC23(_this.f4, _793_v32.f4, _793_v32.f4, (_this).f12);
        let _796_v35;
        _796_v35 = _dafny.Map.Empty.slice().updateUnsafe(_793_v32.f4,_this.f4);
        let _797_v36;
        let _nw148 = Array((new BigNumber(22)).toNumber());
        _nw148[(_dafny.ZERO).toNumber()] = _this.f4;
        _nw148[(_dafny.ONE).toNumber()] = _this.f4;
        _nw148[(new BigNumber(2)).toNumber()] = _this.f4;
        _nw148[(new BigNumber(3)).toNumber()] = _this.f4;
        _nw148[(new BigNumber(4)).toNumber()] = (_this.f4) && (true);
        _nw148[(new BigNumber(5)).toNumber()] = _this.f4;
        _nw148[(new BigNumber(6)).toNumber()] = (_789_v27)[_module.__default.safeIndex((_this).f7, new BigNumber((_789_v27).length))];
        _nw148[(new BigNumber(7)).toNumber()] = !(_790_v28).contains((_this).f7);
        _nw148[(new BigNumber(8)).toNumber()] = _this.f4;
        _nw148[(new BigNumber(9)).toNumber()] = _this.f4;
        _nw148[(new BigNumber(10)).toNumber()] = _this.f4;
        _nw148[(new BigNumber(11)).toNumber()] = (new BigNumber((_module.__default.fm14(_this.f4, _this.f4, globalState)).cardinality())).isEqualTo(new BigNumber(-524));
        _nw148[(new BigNumber(12)).toNumber()] = _this.f4;
        _nw148[(new BigNumber(13)).toNumber()] = (new BigNumber(-434)).isLessThanOrEqualTo(new BigNumber(-900));
        _nw148[(new BigNumber(14)).toNumber()] = (_module.__default.fm18(_792_v31, (_module.D1.create_DC6(_793_v32, _module.__default.fm20(new BigNumber((_794_v33).length), globalState), _this.f4, _793_v32.f4)).dtor_cf11, _793_v32.f4, globalState)).IsSubsetOf(function () {
          let _coll11 = new _dafny.Set();
          for (const _compr_11 of (_this.f13).Elements) {
            let _798_v29 = _compr_11;
            if (_dafny.Seq.contains(_this.f13, _798_v29)) {
              _coll11.add(_798_v29);
            }
          }
          return _coll11;
        }());
        _nw148[(new BigNumber(15)).toNumber()] = _793_v32.f4;
        _nw148[(new BigNumber(16)).toNumber()] = _793_v32.f4;
        _nw148[(new BigNumber(17)).toNumber()] = (_795_v34).dtor_cf32;
        _nw148[(new BigNumber(18)).toNumber()] = _793_v32.f4;
        _nw148[(new BigNumber(19)).toNumber()] = _this.f4;
        _nw148[(new BigNumber(20)).toNumber()] = (((_796_v35).contains(_this.f4)) ? ((_796_v35).get(_this.f4)) : (_this.f4));
        _nw148[(new BigNumber(21)).toNumber()] = _this.f4;
        _797_v36 = _nw148;
        let _799_v37;
        _799_v37 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_this.f13);
        let _index143 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_797_v36).length));
        (_797_v36)[_index143] = (new BigNumber((_794_v33).length)).isLessThanOrEqualTo(new BigNumber((_799_v37).length));
        let _800_v38;
        _800_v38 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("kmrehf")).length));
        let _index144 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_797_v36).length));
        let _rhs144 = ((((_795_v34).dtor_cf32) ? ((_this).f7) : (new BigNumber(900)))).multipliedBy(new BigNumber((_module.__default.fm27(globalState)).length));
        let _rhs145 = _dafny.Seq.update(_788_v26, _module.__default.safeIndex((_793_v32).f3, new BigNumber((_788_v26).length)), _787_v25);
        let _rhs146 = _dafny.Seq.contains(_module.__default.fm15(_this.f4, new BigNumber(318), (_dafny.ZERO).minus(new BigNumber((_800_v38).length)), globalState), (_this.f13)[_module.__default.safeIndex((_793_v32).f3, new BigNumber((_this.f13).length))]);
        let _lhs135 = globalState;
        let _lhs136 = _797_v36;
        let _lhs137 = _module.__default.safeIndex(new BigNumber(521), new BigNumber((_797_v36).length));
        _lhs135.f1 = _rhs144;
        _788_v26 = _rhs145;
        _lhs136[_lhs137] = _rhs146;
        (_this).f4 = (_this).fm13(globalState);
        let _801_v39;
        _801_v39 = new _dafny.CodePoint('l'.codePointAt(0));
        _801_v39 = (_this).f12;
      } else if (_source10.is_DC0) {
        let _802___mcc_h6 = (_source10).cf0;
        let _803_cf0 = _802___mcc_h6;
        let _804_v40;
        let _nw149 = Array((new BigNumber(6)).toNumber()).fill(false);
        _804_v40 = _nw149;
        let _805_v41;
        _805_v41 = _dafny.Map.Empty.slice().updateUnsafe(_804_v40,new BigNumber((_dafny.Seq.of((_this).f7)).length));
        _805_v41 = (_805_v41).update(_804_v40, (_this).f3);
        if ((_this).fm13(globalState)) {
          let _806_v42;
          _806_v42 = _module.D3.create_DC10((_this).f3, _804_v40, ((_this).f7).multipliedBy((_this).f3));
          _806_v42 = _806_v42;
          let _807_v43;
          _807_v43 = _dafny.Seq.of((_module.D0.create_DC0(_803_cf0)).dtor_cf0);
          let _808_v44;
          _808_v44 = _dafny.Map.Empty.slice().updateUnsafe((_this).f12,_module.__default.fm2(new BigNumber((_807_v43).length), globalState));
          let _809_v45;
          let _nw150 = new _module.C2();
          _nw150.__ctor((_dafny.ZERO).minus((_this).fm12(_803_cf0, true, globalState)), _752_v6, new BigNumber(((_808_v44).Merge(_808_v44)).length), _this.f4);
          _809_v45 = _nw150;
          (globalState).f2 = (!(_803_cf0)) || (_803_cf0);
          let _810_v46;
          _810_v46 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_803_cf0);
          let _811_v47;
          _811_v47 = _module.D6.create_DC18((((_810_v46).contains(_this.f4)) ? ((_810_v46).get(_this.f4)) : (_803_cf0)), (_this).f7);
          let _812_v48;
          _812_v48 = _dafny.Map.Empty.slice().updateUnsafe((_809_v45).f14,_module.__default.fm2((_811_v47).dtor_cf25, globalState));
          let _813_v49;
          _813_v49 = _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length));
          let _rhs147 = ((!_dafny.areEqual(_813_v49, _813_v49)) ? (_812_v48) : (_812_v48));
          let _rhs148 = (_809_v45).f14;
          let _rhs149 = _module.__default.safeDivisionInt((_809_v45).f14, (_this).f7);
          let _rhs150 = new BigNumber(128);
          let _lhs138 = globalState;
          let _lhs139 = globalState;
          let _lhs140 = globalState;
          _812_v48 = _rhs147;
          _lhs138.f1 = _rhs148;
          _lhs139.f1 = _rhs149;
          _lhs140.f1 = _rhs150;
          (_this).f13 = _dafny.Seq.Concat(_this.f13, _dafny.Seq.Create(_module.__default.abs(new BigNumber(659)), function (_814_i4) {
            return (_this).f12;
          }));
        } else {
          let _815_v50;
          _815_v50 = _dafny.Seq.of(!(_803_cf0), _803_cf0);
          (globalState).f2 = (_815_v50)[_module.__default.safeIndex(new BigNumber((_this.f13).length), new BigNumber((_815_v50).length))];
          let _816_v51;
          _816_v51 = _dafny.Set.fromElements(true);
          let _817_v52;
          _817_v52 = _module.D9.create_DC25(_816_v51);
          let _rhs151 = ((_817_v52).dtor_cf37).Difference(_816_v51);
          let _rhs152 = ((_this).fm13(globalState)) || (_this.f4);
          let _lhs141 = globalState;
          _816_v51 = _rhs151;
          _lhs141.f2 = _rhs152;
          let _index145 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_804_v40).length));
          (_804_v40)[_index145] = !(_803_cf0);
          let _818_v53;
          _818_v53 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("gfobxgt")).length));
          let _819_v54;
          _819_v54 = _dafny.MultiSet.fromElements(_module.__default.fm15(false, (_dafny.ZERO).minus((_this).f3), new BigNumber((_dafny.Seq.UnicodeFromString("vbacq")).length), globalState), _this.f13);
          let _820_v55;
          _820_v55 = _dafny.Set.fromElements((_this).f12, (_this).f12, (_this).f12);
          let _index146 = _module.__default.safeIndex(new BigNumber(555), new BigNumber((_804_v40).length));
          (_804_v40)[_index146] = !(!((_818_v53).IsDisjointFrom(_dafny.MultiSet.FromArray(_dafny.Seq.of(new BigNumber((_819_v54).cardinality()), new BigNumber((_820_v55).length)))))) || (_this.f4);
          (globalState).f1 = (_this).f3;
          (globalState).f2 = _this.f4;
        }
        let _index147 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_804_v40).length));
        (_804_v40)[_index147] = (new BigNumber((_dafny.Seq.UnicodeFromString("mibpnhp")).length)).isLessThan(_module.__default.fm0(globalState));
        let _821_v56;
        _821_v56 = _dafny.MultiSet.fromElements((_this).f3, (_this).f3, (_this).f7, (_this).f7);
        let _index148 = _module.__default.safeIndex(new BigNumber(283), new BigNumber((_804_v40).length));
        (_804_v40)[_index148] = (_821_v56).IsDisjointFrom(_821_v56);
        let _index149 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_752_v6).length));
        (_752_v6)[_index149] = _module.__default.safeModuloInt((_this).f3, (_this).f7);
        let _822_v57;
        _822_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_821_v56);
        let _index150 = _module.__default.safeIndex(new BigNumber(216), new BigNumber((_752_v6).length));
        (_752_v6)[_index150] = new BigNumber(((_822_v57).Merge(_822_v57)).length);
      } else {
        let _823___mcc_h7 = (_source10).cf7;
        let _824_cf7 = _823___mcc_h7;
        let _825_v58;
        _825_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,!(_dafny.Seq.contains(_this.f13, (_this).f12)));
        let _826_v59;
        _826_v59 = _module.D9.create_DC26((_this).f3, true);
        _825_v58 = (_825_v58).update((_this).f7, (_826_v59).dtor_cf39);
        (globalState).f2 = _this.f4;
        let _827_v60;
        _827_v60 = _module.D8.create_DC23(_this.f4, false, true, (_this).f12);
        let _828_v61;
        let _nw151 = new _module.C2();
        _nw151.__ctor((_this).f3, _752_v6, ((_this).f3).minus((_this).f7), (_827_v60).dtor_cf33);
        _828_v61 = _nw151;
        let _829_v62;
        _829_v62 = _dafny.Set.fromElements(false, _this.f4, _this.f4);
        let _830_v63;
        let _nw152 = Array((new BigNumber(22)).toNumber());
        _nw152[(_dafny.ZERO).toNumber()] = _this.f4;
        _nw152[(_dafny.ONE).toNumber()] = _this.f4;
        _nw152[(new BigNumber(2)).toNumber()] = !_dafny.Seq.contains(_dafny.Seq.of((_828_v61).f14), (_this).f3);
        _nw152[(new BigNumber(3)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(4)).toNumber()] = !(_this.f4);
        _nw152[(new BigNumber(5)).toNumber()] = (_this.f4) && (_this.f4);
        _nw152[(new BigNumber(6)).toNumber()] = (new BigNumber(759)).isLessThan((_828_v61).f14);
        _nw152[(new BigNumber(7)).toNumber()] = ((false) ? (_this.f4) : (_this.f4));
        _nw152[(new BigNumber(8)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(9)).toNumber()] = !((_this).f7).isEqualTo((_this).f3);
        _nw152[(new BigNumber(10)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(11)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(12)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(13)).toNumber()] = !(false);
        _nw152[(new BigNumber(14)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(15)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(16)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(17)).toNumber()] = (_dafny.Set.fromElements(_this.f4)).IsDisjointFrom(_829_v62);
        _nw152[(new BigNumber(18)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(19)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(20)).toNumber()] = _this.f4;
        _nw152[(new BigNumber(21)).toNumber()] = (_this).fm13(globalState);
        _830_v63 = _nw152;
        let _index151 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_830_v63).length));
        (_830_v63)[_index151] = _this.f4;
        let _index152 = _module.__default.safeIndex(new BigNumber(513), new BigNumber((_830_v63).length));
        (_830_v63)[_index152] = ((((_this.f4) ? (false) : (!(_this.f4)))) ? (!((!(true)) && (_this.f4))) : (_this.f4));
      }
      let _831_v64;
      _831_v64 = new _dafny.CodePoint('g'.codePointAt(0));
      let _832_v65;
      _832_v65 = _module.D8.create_DC23(_this.f4, _this.f4, _this.f4, _831_v64);
      _831_v64 = (_832_v65).dtor_cf35;
      let _hi7 = (_dafny.ZERO).minus((_this).f3);
      for (let _833_i5 = (_this).f7; _833_i5.isLessThan(_hi7); _833_i5 = _833_i5.plus(_dafny.ONE)) {
        let _834_v66;
        let _nw153 = Array((new BigNumber(21)).toNumber()).fill(_module.D6.Default());
        _834_v66 = _nw153;
        let _835_v67;
        _835_v67 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_this.f4);
        let _836_v68;
        _836_v68 = _module.D6.create_DC19(new BigNumber((_835_v67).length));
        let _index153 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_834_v66).length));
        (_834_v66)[_index153] = _836_v68;
        let _index154 = _module.__default.safeIndex(new BigNumber(13), new BigNumber((_834_v66).length));
        (_834_v66)[_index154] = function (_pat_let15_0) {
          return function (_837_dt__update__tmp_h0) {
            return function (_pat_let16_0) {
              return function (_838_dt__update_hcf26_h0) {
                return _module.D6.create_DC19(_838_dt__update_hcf26_h0);
              }(_pat_let16_0);
            }(_833_i5);
          }(_pat_let15_0);
        }(_836_v68);
        let _839_v69;
        _839_v69 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm28(!(_this.f4), (_this).f3, _this.f4, globalState),(_this).f3);
        let _840_v70;
        _840_v70 = _dafny.Seq.of((_dafny.ZERO).minus((_this).f7));
        let _841_v71;
        _841_v71 = _dafny.Map.Empty.slice().updateUnsafe(_840_v70,_dafny.Seq.Create(_module.__default.abs(new BigNumber(-752)), ((_842_v64) => function (_843_i6) {
          return _842_v64;
        })(_831_v64)));
        _839_v69 = (_839_v69).update((_module.D10.create_DC28(_841_v71)).dtor_cf41, _833_i5);
        (globalState).f2 = !(((_this).f7).isLessThanOrEqualTo(_833_i5)) || (_this.f4);
        let _844_v72;
        let _init27 = function (_845_i7) {
          return _this.f4;
        };
        let _nw154 = Array((new BigNumber(3)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw154.length); _i0_27++) {
          _nw154[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _844_v72 = _nw154;
        let _index155 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_844_v72).length));
        (_844_v72)[_index155] = _this.f4;
        let _index156 = _module.__default.safeIndex(new BigNumber(363), new BigNumber((_844_v72).length));
        (_844_v72)[_index156] = _this.f4;
      }
      let _846_v73;
      _846_v73 = _module.D0.create_DC3();
      let _source11 = _846_v73;
      if (_source11.is_DC1) {
        let _847___mcc_h8 = (_source11).cf1;
        let _848___mcc_h9 = (_source11).cf2;
        let _849___mcc_h10 = (_source11).cf3;
        let _850___mcc_h11 = (_source11).cf4;
        let _851_cf4 = _850___mcc_h11;
        let _852_cf3 = _849___mcc_h10;
        let _853_cf2 = _848___mcc_h9;
        let _854_cf1 = _847___mcc_h8;
        (globalState).f1 = ((false) ? ((_this).f7) : (new BigNumber(338)));
        (_this).f4 = _854_cf1;
        let _855_v74;
        _855_v74 = _dafny.Map.Empty.slice().updateUnsafe(_852_cf3,_module.D9.create_DC26((_this).f7, false));
        let _856_v75;
        _856_v75 = _dafny.Set.fromElements((_this).f7, new BigNumber((_855_v74).length));
        (_this).f13 = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ymicw"), _this.f13), _module.__default.safeIndex(new BigNumber((_856_v75).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ymicw"), _this.f13)).length)), _831_v64);
        let _index157 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_752_v6).length));
        (_752_v6)[_index157] = ((_853_cf2) ? ((_this).f3) : (new BigNumber(476)));
        let _857_v76;
        let _init28 = ((_858_cf1) => function (_859_i8) {
          return _858_cf1;
        })(_854_cf1);
        let _nw155 = Array((_dafny.ONE).toNumber());
        for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw155.length); _i0_28++) {
          _nw155[_i0_28] = _init28(new BigNumber(_i0_28));
        }
        _857_v76 = _nw155;
        let _index158 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_857_v76).length));
        (_857_v76)[_index158] = _this.f4;
        let _860_v77;
        _860_v77 = _dafny.Seq.of(!(true));
        let _index159 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_752_v6).length));
        let _index160 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_857_v76).length));
        let _rhs153 = !((_860_v77)[_module.__default.safeIndex((_dafny.ZERO).minus((_this).f3), new BigNumber((_860_v77).length))]) || (_853_cf2);
        let _rhs154 = new BigNumber(52);
        let _rhs155 = _dafny.areEqual(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("dcb"), _this.f13), _this.f13);
        let _rhs156 = !((_860_v77)[_module.__default.safeIndex(_module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.UnicodeFromString("athho")).length), (_this).f3), new BigNumber((_860_v77).length))]);
        let _rhs157 = _853_cf2;
        let _lhs142 = _this;
        let _lhs143 = _752_v6;
        let _lhs144 = _module.__default.safeIndex(new BigNumber(829), new BigNumber((_752_v6).length));
        let _lhs145 = _this;
        let _lhs146 = _857_v76;
        let _lhs147 = _module.__default.safeIndex(new BigNumber(613), new BigNumber((_857_v76).length));
        let _lhs148 = globalState;
        _lhs142.f4 = _rhs153;
        _lhs143[_lhs144] = _rhs154;
        _lhs145.f4 = _rhs155;
        _lhs146[_lhs147] = _rhs156;
        _lhs148.f2 = _rhs157;
      } else if (_source11.is_DC2) {
        let _861___mcc_h12 = (_source11).cf5;
        let _862___mcc_h13 = (_source11).cf6;
        let _863_cf6 = _862___mcc_h13;
        let _864_cf5 = _861___mcc_h12;
        (globalState).f1 = _module.__default.safeDivisionInt((_this).f3, (_this).f3);
        (_this).f4 = !(new BigNumber((_dafny.Seq.Concat(_this.f13, _this.f13)).length)).isEqualTo((_this).f3);
        let _865_v78;
        let _nw156 = Array((new BigNumber(23)).toNumber()).fill(false);
        _865_v78 = _nw156;
        let _866_v79;
        let _nw157 = new _module.C0();
        _nw157.__ctor(!(_863_cf6), _865_v78);
        _866_v79 = _nw157;
        let _index161 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_865_v78).length));
        (_865_v78)[_index161] = _863_cf6;
        let _index162 = _module.__default.safeIndex(new BigNumber(624), new BigNumber((_865_v78).length));
        (_865_v78)[_index162] = _dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(728)), ((_867_v64) => function (_868_i9) {
          return _867_v64;
        })(_831_v64)), _this.f13);
      } else if (_source11.is_DC3) {
        (globalState).f1 = (_dafny.ZERO).minus((_this).f3);
        let _869_v80;
        _869_v80 = _dafny.MultiSet.fromElements((_this).f7, (_this).f7, (_this).f3, new BigNumber(551));
        let _870_v81;
        _870_v81 = _dafny.MultiSet.fromElements(new BigNumber((_869_v80).cardinality()));
        (globalState).f2 = (_870_v81).IsSubsetOf(_870_v81);
        let _871_v82;
        _871_v82 = _dafny.Seq.of(_this.f4, false, _this.f4, _this.f4);
        let _872_v83;
        _872_v83 = _dafny.MultiSet.fromElements((_871_v82)[_module.__default.safeIndex((_this).f3, new BigNumber((_871_v82).length))], _this.f4);
        (globalState).f1 = _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber(((_872_v83).Difference(_872_v83)).cardinality())), (_this).f3);
        (globalState).f1 = (((new BigNumber(561)).isEqualTo((_this).f7)) ? (new BigNumber(279)) : ((_this).f7));
      } else if (_source11.is_DC0) {
        let _873___mcc_h14 = (_source11).cf0;
        let _874_cf0 = _873___mcc_h14;
        let _rhs158 = _874_cf0;
        let _rhs159 = _module.__default.fm2((_this).f3, globalState);
        let _lhs149 = globalState;
        let _lhs150 = _this;
        _lhs149.f2 = _rhs158;
        _lhs150.f4 = _rhs159;
        let _pat_let_tv10 = _831_v64;
        let _pat_let_tv11 = _874_cf0;
        let _875_v84;
        let _nw158 = Array((new BigNumber(13)).toNumber());
        _nw158[(_dafny.ZERO).toNumber()] = _832_v65;
        _nw158[(_dafny.ONE).toNumber()] = _module.D8.create_DC23(_this.f4, _this.f4, true, (_this).f12);
        _nw158[(new BigNumber(2)).toNumber()] = _832_v65;
        _nw158[(new BigNumber(3)).toNumber()] = _module.D8.create_DC23(_this.f4, true, _this.f4, _831_v64);
        _nw158[(new BigNumber(4)).toNumber()] = _832_v65;
        _nw158[(new BigNumber(5)).toNumber()] = _832_v65;
        _nw158[(new BigNumber(6)).toNumber()] = _module.D8.create_DC23(_this.f4, _874_cf0, _this.f4, _831_v64);
        _nw158[(new BigNumber(7)).toNumber()] = _832_v65;
        _nw158[(new BigNumber(8)).toNumber()] = _module.D8.create_DC23(_874_cf0, _874_cf0, _this.f4, (_this).f12);
        _nw158[(new BigNumber(9)).toNumber()] = _832_v65;
        _nw158[(new BigNumber(10)).toNumber()] = _832_v65;
        _nw158[(new BigNumber(11)).toNumber()] = _832_v65;
        _nw158[(new BigNumber(12)).toNumber()] = function (_pat_let17_0) {
          return function (_876_dt__update__tmp_h1) {
            return function (_pat_let18_0) {
              return function (_877_dt__update_hcf35_h0) {
                return function (_pat_let19_0) {
                  return function (_878_dt__update_hcf32_h0) {
                    return _module.D8.create_DC23(_878_dt__update_hcf32_h0, (_876_dt__update__tmp_h1).dtor_cf33, (_876_dt__update__tmp_h1).dtor_cf34, _877_dt__update_hcf35_h0);
                  }(_pat_let19_0);
                }(_pat_let_tv11);
              }(_pat_let18_0);
            }(_pat_let_tv10);
          }(_pat_let17_0);
        }(_832_v65);
        _875_v84 = _nw158;
        let _879_v85;
        _879_v85 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_875_v84);
        _879_v85 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(638),_875_v84);
        (_this).f4 = true;
        let _880_v86;
        _880_v86 = _dafny.MultiSet.fromElements(_this.f4, _this.f4, _874_cf0, false, _this.f4);
        let _881_v87;
        _881_v87 = _dafny.MultiSet.fromElements(new BigNumber((_this.f13).length), (_this).f7, _module.__default.safeModuloInt((_this).f3, new BigNumber((_880_v86).cardinality())), (_dafny.ZERO).minus((_this).f7));
        let _882_v88;
        _882_v88 = _dafny.Seq.of((_this).f3, (_this).f3, (_this).f3);
        let _883_v89;
        _883_v89 = _module.D11.create_DC31(_882_v88);
        _881_v87 = _dafny.MultiSet.FromArray((_883_v89).dtor_cf49);
      } else {
        let _884___mcc_h15 = (_source11).cf7;
        let _885_cf7 = _884___mcc_h15;
        let _886_v91;
        _886_v91 = _dafny.Set.fromElements((_this).f3);
        (globalState).f2 = (_886_v91).IsSubsetOf(function () {
          let _coll12 = new _dafny.Set();
          for (const _compr_12 of _dafny.IntegerRange(new BigNumber(819), new BigNumber(503))) {
            let _887_v90 = _compr_12;
            if (((new BigNumber(819)).isLessThanOrEqualTo(_887_v90)) && ((_887_v90).isLessThan(new BigNumber(503)))) {
              _coll12.add(_module.__default.safeModuloInt(_887_v90, (_this).f3));
            }
          }
          return _coll12;
        }());
        let _888_v92;
        _888_v92 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("upfnlrddl")).length),_this.f4);
        let _889_v94;
        _889_v94 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),new BigNumber((_886_v91).length));
        let _890_v97;
        _890_v97 = _dafny.Seq.of(new BigNumber(44), (_this).f3);
        let _891_v98;
        _891_v98 = _dafny.Seq.of(_this.f4);
        let _892_v100;
        let _nw159 = Array((new BigNumber(28)).toNumber());
        _nw159[(_dafny.ZERO).toNumber()] = (_888_v92).Merge(_888_v92);
        _nw159[(_dafny.ONE).toNumber()] = (_888_v92).update(_module.__default.fm0(globalState), _this.f4);
        _nw159[(new BigNumber(2)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(3)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(4)).toNumber()] = ((_888_v92).update((_this).f3, _this.f4)).Merge(_888_v92);
        _nw159[(new BigNumber(5)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_this.f4);
        _nw159[(new BigNumber(6)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(7)).toNumber()] = (_888_v92).update((_this).f7, _this.f4);
        _nw159[(new BigNumber(8)).toNumber()] = (_888_v92).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).f3,_this.f4));
        _nw159[(new BigNumber(9)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(10)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(11)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(12)).toNumber()] = (_888_v92).Merge(function () {
          let _coll13 = new _dafny.Map();
          for (const _compr_13 of _dafny.IntegerRange(new BigNumber(-833), new BigNumber(-986))) {
            let _893_v93 = _compr_13;
            if (((new BigNumber(-833)).isLessThanOrEqualTo(_893_v93)) && ((_893_v93).isLessThan(new BigNumber(-986)))) {
              _coll13.push([(_893_v93).plus((_this).f7),_this.f4]);
            }
          }
          return _coll13;
        }());
        _nw159[(new BigNumber(13)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((((_889_v94).contains(_831_v64)) ? ((_889_v94).get(_831_v64)) : ((_this).f7)),_this.f4);
        _nw159[(new BigNumber(15)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(16)).toNumber()] = (function () {
          let _coll14 = new _dafny.Map();
          for (const _compr_14 of _dafny.IntegerRange(new BigNumber(626), new BigNumber(259))) {
            let _894_v95 = _compr_14;
            if (((new BigNumber(626)).isLessThanOrEqualTo(_894_v95)) && ((_894_v95).isLessThan(new BigNumber(259)))) {
              _coll14.push([(_894_v95).minus((_this).f7),_this.f4]);
            }
          }
          return _coll14;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(607),_this.f4));
        _nw159[(new BigNumber(17)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(18)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(19)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(20)).toNumber()] = _888_v92;
        _nw159[(new BigNumber(21)).toNumber()] = function () {
          let _coll15 = new _dafny.Map();
          for (const _compr_15 of (_890_v97).Elements) {
            let _895_v96 = _compr_15;
            if (_dafny.Seq.contains(_890_v97, _895_v96)) {
              _coll15.push([_module.__default.safeModuloInt(_895_v96, (_this).f3),_this.f4]);
            }
          }
          return _coll15;
        }();
        _nw159[(new BigNumber(22)).toNumber()] = (_888_v92).Merge(_888_v92);
        _nw159[(new BigNumber(23)).toNumber()] = (_888_v92).Merge(_888_v92);
        _nw159[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,!(_this.f4));
        _nw159[(new BigNumber(25)).toNumber()] = ((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,_this.f4)).update((_this).f3, false)).Merge(_888_v92);
        _nw159[(new BigNumber(26)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(211),(_891_v98)[_module.__default.safeIndex((_this).f7, new BigNumber((_891_v98).length))]);
        _nw159[(new BigNumber(27)).toNumber()] = (_888_v92).Merge(function () {
          let _coll16 = new _dafny.Map();
          for (const _compr_16 of _dafny.IntegerRange(new BigNumber(937), new BigNumber(258))) {
            let _896_v99 = _compr_16;
            if (((new BigNumber(937)).isLessThanOrEqualTo(_896_v99)) && ((_896_v99).isLessThan(new BigNumber(258)))) {
              _coll16.push([(_896_v99).minus((_this).f7),_this.f4]);
            }
          }
          return _coll16;
        }());
        _892_v100 = _nw159;
        let _index163 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_892_v100).length));
        (_892_v100)[_index163] = _888_v92;
        let _897_v101;
        _897_v101 = _dafny.MultiSet.fromElements(_this.f4, _this.f4);
        let _index164 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_892_v100).length));
        let _rhs160 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,((_this).f7).isLessThan(new BigNumber((_890_v97).length)));
        let _rhs161 = (_897_v101).IsProperSubsetOf(_897_v101);
        let _lhs151 = _892_v100;
        let _lhs152 = _module.__default.safeIndex(new BigNumber(496), new BigNumber((_892_v100).length));
        let _lhs153 = globalState;
        _lhs151[_lhs152] = _rhs160;
        _lhs153.f2 = _rhs161;
        (globalState).f2 = _this.f4;
        let _898_v102;
        let _nw160 = new _module.C1();
        _nw160.__ctor((_this).f7, new BigNumber(885), _module.__default.fm2((_this).f3, globalState));
        _898_v102 = _nw160;
      }
      (_this).f4 = true;
      return;
    }
    get f12() {
      let _this = this;
      return _this._f12;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f4 = false;
      this._f7 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f7, f3, f4) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      (globalState).f2 = _this.f4;
      let _899_v0;
      let _init29 = function (_900_i0) {
        return _dafny.areEqual(_dafny.Seq.of(!(_this.f4)), _dafny.Seq.of(!(_this.f4), _this.f4));
      };
      let _nw161 = Array((new BigNumber(7)).toNumber());
      for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw161.length); _i0_29++) {
        _nw161[_i0_29] = _init29(new BigNumber(_i0_29));
      }
      _899_v0 = _nw161;
      let _index165 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_899_v0).length));
      (_899_v0)[_index165] = _this.f4;
      let _index166 = _module.__default.safeIndex(new BigNumber(80), new BigNumber((_899_v0).length));
      (_899_v0)[_index166] = true;
      let _hi8 = new BigNumber(972);
      for (let _901_i1 = (_this).f7; _901_i1.isLessThan(_hi8); _901_i1 = _901_i1.plus(_dafny.ONE)) {
        let _902_v1;
        let _nw162 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
        _902_v1 = _nw162;
        let _903_v2;
        _903_v2 = _dafny.Set.fromElements(_902_v1);
        let _904_v3;
        let _nw163 = Array((_dafny.ONE).toNumber()).fill(_dafny.Seq.of());
        _904_v3 = _nw163;
        let _905_v4;
        _905_v4 = _dafny.Seq.of((_this).f3, p0, (_this).f3, _901_i1);
        let _index167 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_904_v3).length));
        (_904_v3)[_index167] = _dafny.Seq.Concat(_905_v4, _905_v4);
        let _906_v5;
        _906_v5 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_module.__default.fm0(globalState)),_903_v2);
        let _index168 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_904_v3).length));
        let _rhs162 = _this.f4;
        let _rhs163 = (((_906_v5).contains(p0)) ? ((_906_v5).get(p0)) : (_903_v2));
        let _rhs164 = _905_v4;
        let _lhs154 = globalState;
        let _lhs155 = _904_v3;
        let _lhs156 = _module.__default.safeIndex(new BigNumber(549), new BigNumber((_904_v3).length));
        _lhs154.f2 = _rhs162;
        _903_v2 = _rhs163;
        _lhs155[_lhs156] = _rhs164;
        let _907_v6;
        _907_v6 = _dafny.Seq.UnicodeFromString("g");
        (globalState).f1 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_907_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-923)), function (_908_i2) {
          return new _dafny.CodePoint('k'.codePointAt(0));
        }))).length));
        let _909_v7;
        let _nw164 = new _module.C0();
        _nw164.__ctor(_dafny.areEqual(_dafny.Seq.UnicodeFromString("dmaqyb"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(486)), function (_910_i3) {
          return new _dafny.CodePoint('d'.codePointAt(0));
        })), _899_v0);
        _909_v7 = _nw164;
        let _911_v8;
        _911_v8 = _dafny.MultiSet.fromElements(_901_i1);
        let _912_v9;
        _912_v9 = new _dafny.CodePoint('a'.codePointAt(0));
        let _913_v10;
        _913_v10 = _dafny.MultiSet.fromElements(_912_v9);
        let _914_v11;
        _914_v11 = _module.D0.create_DC2(_913_v10, _909_v7.f10);
        let _915_v12;
        _915_v12 = _dafny.Map.Empty.slice().updateUnsafe(_914_v11,(_899_v0)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_899_v0).length))]);
        let _pat_let_tv12 = _913_v10;
        let _pat_let_tv13 = _913_v10;
        if ((_911_v8).IsDisjointFrom(_module.__default.fm6((((_915_v12).contains(function (_pat_let22_0) {
          return function (_918_dt__update__tmp_h0) {
            return function (_pat_let23_0) {
              return function (_919_dt__update_hcf5_h0) {
                return _module.D0.create_DC2(_919_dt__update_hcf5_h0, (_918_dt__update__tmp_h0).dtor_cf6);
              }(_pat_let23_0);
            }(_pat_let_tv13);
          }(_pat_let22_0);
        }(_module.D0.create_DC2(_913_v10, true)))) ? ((_915_v12).get(function (_pat_let20_0) {
          return function (_916_dt__update__tmp_h1) {
            return function (_pat_let21_0) {
              return function (_917_dt__update_hcf5_h1) {
                return _module.D0.create_DC2(_917_dt__update_hcf5_h1, (_916_dt__update__tmp_h1).dtor_cf6);
              }(_pat_let21_0);
            }(_pat_let_tv12);
          }(_pat_let20_0);
        }(_module.D0.create_DC2(_913_v10, true)))) : (_this.f4)), (_899_v0)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_899_v0).length))], globalState))) {
          let _920_v13;
          _920_v13 = _module.D0.create_DC1(!(_this.f4), _909_v7.f10, _912_v9, _902_v1);
          let _921_v14;
          _921_v14 = _dafny.Seq.of((_920_v13).dtor_cf1);
          let _922_v15;
          _922_v15 = _dafny.Map.Empty.slice().updateUnsafe(_921_v14,_907_v6);
          let _923_v16;
          _923_v16 = _dafny.Set.fromElements(_901_i1);
          let _924_v18;
          _924_v18 = _dafny.Map.Empty.slice().updateUnsafe(_907_v6,(_this).f7);
          let _rhs165 = ((!(_this.f4) || (_909_v7.f10)) ? (_912_v9) : (_912_v9));
          let _rhs166 = (_922_v15).Merge((_922_v15).Merge(_922_v15));
          let _rhs167 = _923_v16;
          let _rhs168 = new BigNumber((_module.__default.fm7(function () {
            let _coll17 = new _dafny.Map();
            for (const _compr_17 of _dafny.IntegerRange(new BigNumber(792), new BigNumber(-411))) {
              let _925_v17 = _compr_17;
              if (((new BigNumber(792)).isLessThanOrEqualTo(_925_v17)) && ((_925_v17).isLessThan(new BigNumber(-411)))) {
                _coll17.push([(_925_v17).multipliedBy((((_911_v8).contains((_this).f3)) ? ((_911_v8).get((_this).f3)) : ((_this).f7))),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_901_i1,_909_v7.f10)).length)]);
              }
            }
            return _coll17;
          }(), (_this).f3, _924_v18, globalState)).length);
          let _lhs157 = globalState;
          _912_v9 = _rhs165;
          _922_v15 = _rhs166;
          _923_v16 = _rhs167;
          _lhs157.f1 = _rhs168;
          (globalState).f1 = _module.__default.safeModuloInt((_this).f7, p0);
          let _926_v19;
          _926_v19 = _dafny.Map.Empty.slice().updateUnsafe(_909_v7.f10,(_this).f7);
          let _index169 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_902_v1).length));
          (_902_v1)[_index169] = new BigNumber((_926_v19).length);
          let _927_v20;
          _927_v20 = _module.D1.create_DC7();
          let _928_v21;
          _928_v21 = _dafny.Map.Empty.slice().updateUnsafe(_module.D1.create_DC7(),_this.f4);
          let _929_v22;
          _929_v22 = _dafny.MultiSet.fromElements(_this.f4, _909_v7.f10, _909_v7.f10);
          let _index170 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_902_v1).length));
          let _rhs169 = !(_909_v7.f10);
          let _rhs170 = (_this).f7;
          let _rhs171 = (_928_v21).contains(_927_v20);
          let _rhs172 = ((new BigNumber((_907_v6).length)).multipliedBy(new BigNumber((_929_v22).cardinality()))).minus((_this).f7);
          let _lhs158 = globalState;
          let _lhs159 = globalState;
          let _lhs160 = globalState;
          let _lhs161 = _902_v1;
          let _lhs162 = _module.__default.safeIndex(new BigNumber(696), new BigNumber((_902_v1).length));
          _lhs158.f2 = _rhs169;
          _lhs159.f1 = _rhs170;
          _lhs160.f2 = _rhs171;
          _lhs161[_lhs162] = _rhs172;
          let _930_v23;
          let _nw165 = new _module.C0();
          _nw165.__ctor((_this.f4) || (_this.f4), (_909_v7).f11);
          _930_v23 = _nw165;
          let _931_v24;
          _931_v24 = _module.D0.create_DC3();
          _931_v24 = _931_v24;
        } else {
          (globalState).f2 = (_899_v0)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_899_v0).length))];
          (globalState).f1 = (_this).f7;
          let _932_v25;
          _932_v25 = _dafny.Map.Empty.slice().updateUnsafe(_909_v7.f10,!(_this.f4));
          let _index171 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_902_v1).length));
          (_902_v1)[_index171] = (_901_i1).multipliedBy((_this).f7);
          let _index172 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_902_v1).length));
          let _rhs173 = _909_v7.f10;
          let _rhs174 = _932_v25;
          let _rhs175 = new BigNumber((function () {
            let _coll18 = new _dafny.Set();
            for (const _compr_18 of _dafny.IntegerRange(new BigNumber(149), new BigNumber(861))) {
              let _933_v26 = _compr_18;
              if (((new BigNumber(149)).isLessThanOrEqualTo(_933_v26)) && ((_933_v26).isLessThan(new BigNumber(861)))) {
                _coll18.add(_module.__default.safeModuloInt(_933_v26, _901_i1));
              }
            }
            return _coll18;
          }()).length);
          let _rhs176 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(p0, (_this).f3)).multipliedBy(_module.__default.fm0(globalState)));
          let _lhs163 = _this;
          let _lhs164 = globalState;
          let _lhs165 = _902_v1;
          let _lhs166 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_902_v1).length));
          _lhs163.f4 = _rhs173;
          _932_v25 = _rhs174;
          _lhs164.f1 = _rhs175;
          _lhs165[_lhs166] = _rhs176;
          let _934_v27;
          _934_v27 = _module.D0.create_DC0(_909_v7.f10);
          _907_v6 = _module.__default.fm8(false, _module.D1.create_DC7(), _934_v27, p0, globalState);
          (globalState).f2 = !(true);
        }
      }
      let _935_v28;
      _935_v28 = _dafny.Seq.of(new BigNumber(231), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements((_899_v0)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_899_v0).length))], _this.f4)).length)));
      r0 = (_935_v28)[_module.__default.safeIndex((p0).minus((_this).f3), new BigNumber((_935_v28).length))];
      let _936_v29;
      _936_v29 = _899_v0;
      let _937_v30;
      _937_v30 = _dafny.Seq.of((_936_v29));
      let _938_v31;
      let _nw166 = Array((new BigNumber(19)).toNumber());
      _nw166[(_dafny.ZERO).toNumber()] = _899_v0;
      _nw166[(_dafny.ONE).toNumber()] = (((_899_v0)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_899_v0).length))]) ? (_899_v0) : (_899_v0));
      _nw166[(new BigNumber(2)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(3)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(4)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(5)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(6)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(7)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(8)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(9)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(10)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(11)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(12)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(13)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(14)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(15)).toNumber()] = _899_v0;
      _nw166[(new BigNumber(16)).toNumber()] = (_937_v30)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_937_v30).length))];
      _nw166[(new BigNumber(17)).toNumber()] = (((_899_v0)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_899_v0).length))]) ? (_899_v0) : (_899_v0));
      _nw166[(new BigNumber(18)).toNumber()] = _899_v0;
      _938_v31 = _nw166;
      let _index173 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_938_v31).length));
      (_938_v31)[_index173] = _899_v0;
      let _index174 = _module.__default.safeIndex(new BigNumber(203), new BigNumber((_938_v31).length));
      (_938_v31)[_index174] = _899_v0;
      let _939_v32;
      let _nw167 = new _module.C0();
      _nw167.__ctor((_899_v0)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_899_v0).length))], (_937_v30)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_937_v30).length))]);
      _939_v32 = _nw167;
      let _940_v33;
      _940_v33 = _module.D0.create_DC0(_939_v32.f10);
      let _pat_let_tv14 = p0;
      let _pat_let_tv15 = _899_v0;
      let _pat_let_tv16 = _899_v0;
      r0 = (_dafny.ZERO).minus(function (_source12) {
        if (_source12.is_DC1) {
          let _941___mcc_h0 = (_source12).cf1;
          let _942___mcc_h1 = (_source12).cf2;
          let _943___mcc_h2 = (_source12).cf3;
          let _944___mcc_h3 = (_source12).cf4;
          let _945_cf4 = _944___mcc_h3;
          let _946_cf3 = _943___mcc_h2;
          let _947_cf2 = _942___mcc_h1;
          let _948_cf1 = _941___mcc_h0;
          return (_pat_let_tv14).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_pat_let_tv16)[_module.__default.safeIndex(new BigNumber(80), new BigNumber((_pat_let_tv15).length))],true)).length));
        } else if (_source12.is_DC2) {
          let _949___mcc_h4 = (_source12).cf5;
          let _950___mcc_h5 = (_source12).cf6;
          let _951_cf6 = _950___mcc_h5;
          let _952_cf5 = _949___mcc_h4;
          return (_this).f7;
        } else if (_source12.is_DC3) {
          return new BigNumber(480);
        } else if (_source12.is_DC0) {
          let _953___mcc_h6 = (_source12).cf0;
          let _954_cf0 = _953___mcc_h6;
          return (_this).f7;
        } else {
          let _955___mcc_h7 = (_source12).cf7;
          let _956_cf7 = _955___mcc_h7;
          return (_this).f7;
        }
      }(_940_v33));
      return r0;
    }
    m4(p0, globalState) {
      let _this = this;
      let _957_i0;
      _957_i0 = _dafny.ZERO;
      L8: {
        while (_this.f4) {
          C8: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_957_i0)) {
              break L8;
            }
            _957_i0 = (_957_i0).plus(_dafny.ONE);
            let _958_v0;
            _958_v0 = new _dafny.CodePoint('w'.codePointAt(0));
            let _959_v1;
            _959_v1 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f4),_958_v0);
            let _960_v2;
            let _nw168 = Array((new BigNumber(6)).toNumber());
            _nw168[(_dafny.ZERO).toNumber()] = _958_v0;
            _nw168[(_dafny.ONE).toNumber()] = _958_v0;
            _nw168[(new BigNumber(2)).toNumber()] = _958_v0;
            _nw168[(new BigNumber(3)).toNumber()] = (((_959_v1).contains(_module.__default.fm2(new BigNumber(866), globalState))) ? ((_959_v1).get(_module.__default.fm2(new BigNumber(866), globalState))) : (_958_v0));
            _nw168[(new BigNumber(4)).toNumber()] = _958_v0;
            _nw168[(new BigNumber(5)).toNumber()] = _958_v0;
            _960_v2 = _nw168;
            let _961_v3;
            _961_v3 = _dafny.Seq.of(_960_v2, _960_v2);
            let _962_v4;
            _962_v4 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,(_this).f7);
            let _963_v5;
            _963_v5 = _dafny.MultiSet.fromElements(_module.__default.fm0(globalState), (_this).f3, (_this).f3, (_this).f3);
            let _964_v6;
            _964_v6 = _module.D3.create_DC9(_960_v2);
            let _965_v7;
            _965_v7 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.UnicodeFromString("hynbt")).length),_960_v2);
            let _966_v8;
            let _nw169 = Array((new BigNumber(22)).toNumber());
            _nw169[(_dafny.ZERO).toNumber()] = _960_v2;
            _nw169[(_dafny.ONE).toNumber()] = _960_v2;
            _nw169[(new BigNumber(2)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(3)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(4)).toNumber()] = ((_this.f4) ? (_960_v2) : (_960_v2));
            _nw169[(new BigNumber(5)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(6)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(7)).toNumber()] = (_961_v3)[_module.__default.safeIndex(new BigNumber((_module.__default.fm9((p0)[_module.__default.safeIndex(new BigNumber((_962_v4).length), new BigNumber((p0).length))], _this.f4, globalState)).length), new BigNumber((_961_v3).length))];
            _nw169[(new BigNumber(8)).toNumber()] = (_961_v3)[_module.__default.safeIndex(new BigNumber((_963_v5).cardinality()), new BigNumber((_961_v3).length))];
            _nw169[(new BigNumber(9)).toNumber()] = (_964_v6).dtor_cf14;
            _nw169[(new BigNumber(10)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(11)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(12)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(13)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(14)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(15)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(16)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(17)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(18)).toNumber()] = ((_this.f4) ? ((((_965_v7).contains((_this).f3)) ? ((_965_v7).get((_this).f3)) : (_960_v2))) : (_960_v2));
            _nw169[(new BigNumber(19)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(20)).toNumber()] = _960_v2;
            _nw169[(new BigNumber(21)).toNumber()] = _960_v2;
            _966_v8 = _nw169;
            _966_v8 = _966_v8;
            (globalState).f1 = ((_this).f7).minus(new BigNumber((_dafny.Seq.update(_dafny.Seq.of((_this).f3), _module.__default.safeIndex(new BigNumber((_module.__default.fm10(_this.f4, _this.f4, (_this).f3, _dafny.Map.Empty.slice().updateUnsafe(_this.f4,p0), globalState)).length), new BigNumber((_dafny.Seq.of((_this).f3)).length)), new BigNumber(730))).length));
            (_this).f4 = !(_this.f4);
            if (!(true)) {
              (globalState).f2 = false;
              (globalState).f1 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber(-964), (_this).f3), (_this).f3));
              (globalState).f2 = _this.f4;
              (_this).f4 = false;
              (globalState).f2 = true;
            } else {
              let _967_v9;
              _967_v9 = _dafny.Seq.UnicodeFromString("evx");
              _967_v9 = p0;
              let _968_v10;
              let _nw170 = Array((new BigNumber(26)).toNumber()).fill(false);
              _968_v10 = _nw170;
              let _index175 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_968_v10).length));
              (_968_v10)[_index175] = _this.f4;
              let _index176 = _module.__default.safeIndex(new BigNumber(52), new BigNumber((_968_v10).length));
              (_968_v10)[_index176] = _this.f4;
              let _969_v11;
              _969_v11 = _dafny.Set.fromElements((_this).f7);
              let _970_v13;
              _970_v13 = _dafny.Seq.of(new BigNumber((function () {
                let _coll19 = new _dafny.Map();
                for (const _compr_19 of _dafny.IntegerRange(new BigNumber(309), new BigNumber(185))) {
                  let _971_v12 = _compr_19;
                  if (((new BigNumber(309)).isLessThanOrEqualTo(_971_v12)) && ((_971_v12).isLessThan(new BigNumber(185)))) {
                    _coll19.push([(_971_v12).plus(new BigNumber(514)),true]);
                  }
                }
                return _coll19;
              }()).length), (_this).f3);
              let _972_v14;
              let _nw171 = Array((new BigNumber(16)).toNumber());
              _nw171[(_dafny.ZERO).toNumber()] = new BigNumber((_969_v11).length);
              _nw171[(_dafny.ONE).toNumber()] = new BigNumber(613);
              _nw171[(new BigNumber(2)).toNumber()] = (_this).f3;
              _nw171[(new BigNumber(3)).toNumber()] = (_this).f3;
              _nw171[(new BigNumber(4)).toNumber()] = _module.__default.fm0(globalState);
              _nw171[(new BigNumber(5)).toNumber()] = (_this).f3;
              _nw171[(new BigNumber(6)).toNumber()] = (_this).f3;
              _nw171[(new BigNumber(7)).toNumber()] = new BigNumber(373);
              _nw171[(new BigNumber(8)).toNumber()] = new BigNumber((_970_v13).length);
              _nw171[(new BigNumber(9)).toNumber()] = (_this).f7;
              _nw171[(new BigNumber(10)).toNumber()] = new BigNumber((_969_v11).length);
              _nw171[(new BigNumber(11)).toNumber()] = new BigNumber(473);
              _nw171[(new BigNumber(12)).toNumber()] = (_this).f7;
              _nw171[(new BigNumber(13)).toNumber()] = (_this).f7;
              _nw171[(new BigNumber(14)).toNumber()] = (_this).f3;
              _nw171[(new BigNumber(15)).toNumber()] = (_this).f7;
              _972_v14 = _nw171;
              let _973_v15;
              let _nw172 = Array((new BigNumber(18)).toNumber());
              _nw172[(_dafny.ZERO).toNumber()] = _972_v14;
              _nw172[(_dafny.ONE).toNumber()] = _972_v14;
              _nw172[(new BigNumber(2)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(3)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(4)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(5)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(6)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(7)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(8)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(9)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(10)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(11)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(12)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(13)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(14)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(15)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(16)).toNumber()] = _972_v14;
              _nw172[(new BigNumber(17)).toNumber()] = _972_v14;
              _973_v15 = _nw172;
              let _974_v16;
              _974_v16 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,_973_v15);
              _974_v16 = (_974_v16).update(new BigNumber((_dafny.Seq.UnicodeFromString("h")).length), _973_v15);
              let _975_v17;
              _975_v17 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_967_v9);
              (globalState).f1 = ((!_dafny.areEqual(_970_v13, _dafny.Seq.of(new BigNumber(((((_975_v17).contains(_this.f4)) ? ((_975_v17).get(_this.f4)) : (_dafny.Seq.UnicodeFromString("eu")))).length)))) ? ((_dafny.ZERO).minus((_this).f3)) : (_module.__default.safeDivisionInt(new BigNumber(122), (_dafny.ZERO).minus((_this).f3))));
              let _976_v18;
              let _nw173 = new _module.C0();
              _nw173.__ctor(_this.f4, _968_v10);
              _976_v18 = _nw173;
            }
          }
        }
      }
      let _977_v19;
      _977_v19 = new _dafny.CodePoint('k'.codePointAt(0));
      _977_v19 = _977_v19;
      let _978_v20;
      _978_v20 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f3).multipliedBy((_this).f3),_this.f4);
      let _979_v21;
      _979_v21 = _dafny.Set.fromElements(_this.f4, _this.f4);
      let _980_v22;
      _980_v22 = _dafny.MultiSet.fromElements((_this).f3, (_dafny.ZERO).minus(new BigNumber((_979_v21).length)), (_this).f3, (_this).f7, (_dafny.ZERO).minus((_this).f3));
      if ((((_978_v20).contains(new BigNumber((_980_v22).cardinality()))) ? ((_978_v20).get(new BigNumber((_980_v22).cardinality()))) : (_this.f4))) {
        let _981_v23;
        let _nw174 = Array((new BigNumber(23)).toNumber()).fill(false);
        _981_v23 = _nw174;
        let _982_v24;
        let _nw175 = new _module.C0();
        _nw175.__ctor(_this.f4, _981_v23);
        _982_v24 = _nw175;
        let _983_v25;
        _983_v25 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,(_this).f7);
        _983_v25 = (_983_v25).update((_dafny.ZERO).minus(_module.__default.fm0(globalState)), new BigNumber(((_979_v21).Union(_979_v21)).length));
        let _984_v26;
        _984_v26 = _dafny.Map.Empty.slice().updateUnsafe(_982_v24.f10,p0);
        (globalState).f1 = (new BigNumber((_module.__default.fm10(_982_v24.f10, _this.f4, (_this).f7, _984_v26, globalState)).length)).minus((_this).f3);
        (_this).f4 = true;
        (globalState).f1 = (_this).f7;
      } else {
        let _985_v27;
        _985_v27 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_this.f4);
        _985_v27 = (_985_v27).update(_this.f4, _this.f4);
        let _986_v28;
        let _nw176 = Array((new BigNumber(2)).toNumber());
        _nw176[(_dafny.ZERO).toNumber()] = (_this).f3;
        _nw176[(_dafny.ONE).toNumber()] = (_this).f7;
        _986_v28 = _nw176;
        let _index177 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_986_v28).length));
        (_986_v28)[_index177] = _module.__default.fm0(globalState);
        let _987_v29;
        _987_v29 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_this).f7);
        let _988_v30;
        _988_v30 = _dafny.Seq.of(_987_v29);
        let _989_v31;
        _989_v31 = _dafny.Map.Empty.slice().updateUnsafe(p0,_this.f4);
        let _990_v32;
        let _nw177 = Array((new BigNumber(10)).toNumber());
        _nw177[(_dafny.ZERO).toNumber()] = _this.f4;
        _nw177[(_dafny.ONE).toNumber()] = true;
        _nw177[(new BigNumber(2)).toNumber()] = _this.f4;
        _nw177[(new BigNumber(3)).toNumber()] = _this.f4;
        _nw177[(new BigNumber(4)).toNumber()] = true;
        _nw177[(new BigNumber(5)).toNumber()] = _this.f4;
        _nw177[(new BigNumber(6)).toNumber()] = _this.f4;
        _nw177[(new BigNumber(7)).toNumber()] = (((_989_v31).contains(p0)) ? ((_989_v31).get(p0)) : (_this.f4));
        _nw177[(new BigNumber(8)).toNumber()] = _this.f4;
        _nw177[(new BigNumber(9)).toNumber()] = _this.f4;
        _990_v32 = _nw177;
        let _991_v33;
        let _nw178 = new _module.C0();
        _nw178.__ctor(_this.f4, _990_v32);
        _991_v33 = _nw178;
        let _992_v34;
        _992_v34 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,_991_v33);
        let _993_v35;
        _993_v35 = _dafny.Seq.of(_992_v34, _992_v34);
        let _index178 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_986_v28).length));
        (_986_v28)[_index178] = new BigNumber((_dafny.Seq.update(((!((_988_v30)[_module.__default.safeIndex((_this).f3, new BigNumber((_988_v30).length))]).equals((_module.__default.fm11((_this).f7, globalState)).update((_this).f3, (_this).f7))) ? (_993_v35) : (_993_v35)), _module.__default.safeIndex((_dafny.ZERO).minus(((_991_v33.f10) ? ((_this).f3) : ((_this).f7))), new BigNumber((((!((_988_v30)[_module.__default.safeIndex((_this).f3, new BigNumber((_988_v30).length))]).equals((_module.__default.fm11((_this).f7, globalState)).update((_this).f3, (_this).f7))) ? (_993_v35) : (_993_v35))).length)), _992_v34)).length);
        let _994_v36;
        _994_v36 = _module.D11.create_DC32(_991_v33.f10, (_this).f3);
        let _995_v37;
        let _nw179 = new _module.C3();
        _nw179.__ctor(_977_v19, p0, (_994_v36).dtor_cf51, _this.f4, (_this).f7);
        _995_v37 = _nw179;
        let _996_v38;
        let _nw180 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
        _996_v38 = _nw180;
        let _997_v39;
        _997_v39 = _dafny.Seq.of((_this).f7, (_995_v37).f7);
        let _index179 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_996_v38).length));
        (_996_v38)[_index179] = _dafny.Seq.Concat(_997_v39, _997_v39);
        let _998_v40;
        _998_v40 = _module.D1.create_DC7();
        let _index180 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_986_v28).length));
        let _index181 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_996_v38).length));
        let _rhs177 = ((_this).f7).minus((_this).f7);
        let _rhs178 = _995_v37;
        let _rhs179 = _997_v39;
        let _rhs180 = _998_v40;
        let _lhs167 = _986_v28;
        let _lhs168 = _module.__default.safeIndex(new BigNumber(657), new BigNumber((_986_v28).length));
        let _lhs169 = _996_v38;
        let _lhs170 = _module.__default.safeIndex(new BigNumber(856), new BigNumber((_996_v38).length));
        _lhs167[_lhs168] = _rhs177;
        _995_v37 = _rhs178;
        _lhs169[_lhs170] = _rhs179;
        _998_v40 = _rhs180;
        let _999_v41;
        _999_v41 = _module.D7.create_DC20(_dafny.Map.Empty.slice().updateUnsafe((_986_v28)[_module.__default.safeIndex(new BigNumber(657), new BigNumber((_986_v28).length))],p0));
        let _1000_v42;
        _1000_v42 = _dafny.Set.fromElements((_995_v37).f3);
        let _1001_v43;
        _1001_v43 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1000_v42).length),p0);
        let _pat_let_tv17 = _1001_v43;
        _999_v41 = function (_pat_let24_0) {
          return function (_1002_dt__update__tmp_h0) {
            return function (_pat_let25_0) {
              return function (_1003_dt__update_hcf27_h0) {
                return _module.D7.create_DC20(_1003_dt__update_hcf27_h0);
              }(_pat_let25_0);
            }(_pat_let_tv17);
          }(_pat_let24_0);
        }(_999_v41);
        let _1004_v44;
        _1004_v44 = _dafny.Seq.of(p0, p0, p0);
        if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(p0, p0, p0, _dafny.Seq.UnicodeFromString("gjq")), _1004_v44)) {
          let _1005_v45;
          _1005_v45 = _dafny.MultiSet.fromElements(_986_v28);
          let _1006_v46;
          _1006_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1005_v45,(_986_v28)[_module.__default.safeIndex(new BigNumber(657), new BigNumber((_986_v28).length))]);
          _1006_v46 = (_1006_v46).update((_1005_v45).Union(_1005_v45), (_995_v37).f3);
          (_991_v33).f10 = _991_v33.f10;
          (_991_v33).f10 = !(_991_v33.f10);
          let _1007_v47;
          _1007_v47 = _dafny.MultiSet.fromElements(_997_v39);
          (globalState).f1 = (_dafny.ZERO).minus(new BigNumber((_1007_v47).cardinality()));
          (globalState).f1 = (_995_v37).f3;
        } else {
          (_991_v33).f10 = _module.__default.fm2((_995_v37).f7, globalState);
          let _1008_v48;
          _1008_v48 = _dafny.Seq.UnicodeFromString("gxc");
          let _rhs181 = !(_991_v33.f10);
          let _rhs182 = _dafny.Seq.Concat(_dafny.Seq.Concat(p0, _1008_v48), _dafny.Seq.Concat(_1008_v48, _1008_v48));
          let _lhs171 = _995_v37;
          _lhs171.f4 = _rhs181;
          _1008_v48 = _rhs182;
          let _1009_v49;
          _1009_v49 = _dafny.MultiSet.fromElements(_985_v27, (_dafny.Map.Empty.slice().updateUnsafe(!(_this.f4),true)).update(_995_v37.f4, _995_v37.f4), _985_v27);
          let _1010_v50;
          _1010_v50 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f3, (_986_v28)[_module.__default.safeIndex(new BigNumber(657), new BigNumber((_986_v28).length))]),_1009_v49);
          _1010_v50 = (_1010_v50).update((_dafny.ZERO).minus((_995_v37).f3), _1009_v49);
          let _index182 = _module.__default.safeIndex(new BigNumber(986), new BigNumber(((_991_v33).f11).length));
          ((_991_v33).f11)[_index182] = _this.f4;
          let _index183 = _module.__default.safeIndex(new BigNumber(986), new BigNumber(((_991_v33).f11).length));
          let _rhs183 = _995_v37.f4;
          let _rhs184 = ((_995_v37).f3).isLessThanOrEqualTo((_995_v37).f3);
          let _lhs172 = globalState;
          let _lhs173 = (_991_v33).f11;
          let _lhs174 = _module.__default.safeIndex(new BigNumber(986), new BigNumber(((_991_v33).f11).length));
          _lhs172.f2 = _rhs183;
          _lhs173[_lhs174] = _rhs184;
          let _1011_v51;
          _1011_v51 = _dafny.Seq.of((false) && (_this.f4));
          _1011_v51 = _1011_v51;
        }
      }
      (_this).f4 = _this.f4;
      (globalState).f1 = (_this).f7;
      let _1012_v52;
      let _nw181 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.of());
      _1012_v52 = _nw181;
      _1012_v52 = _1012_v52;
      return;
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f4 = false;
      this._f7 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this._f9 = [];
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f9, f7, f3, f4) {
      let _this = this;
      (_this)._f9 = f9;
      (_this)._f7 = f7;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let _1013_v0;
      let _nw182 = Array((new BigNumber(9)).toNumber());
      _1013_v0 = _nw182;
      let _1014_v1;
      let _nw183 = Array((new BigNumber(6)).toNumber()).fill([]);
      _1014_v1 = _nw183;
      let _1015_v2;
      _1015_v2 = _dafny.Seq.of(_1013_v0, _1013_v0);
      let _rhs185 = (_1015_v2)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(p0, (_this).f3, (_this).f3)).cardinality()), new BigNumber((_1015_v2).length))];
      let _rhs186 = _1014_v1;
      let _rhs187 = (((_this.f4) ? (p0) : ((_this).f7))).multipliedBy(p0);
      let _lhs175 = globalState;
      _1013_v0 = _rhs185;
      _1014_v1 = _rhs186;
      _lhs175.f1 = _rhs187;
      let _hi9 = ((_this).f3).plus(new BigNumber((_dafny.Seq.UnicodeFromString("ptajbtlxg")).length));
      for (let _1016_i0 = (_this).f3; _1016_i0.isLessThan(_hi9); _1016_i0 = _1016_i0.plus(_dafny.ONE)) {
        let _1017_v3;
        let _nw184 = Array((new BigNumber(8)).toNumber()).fill(false);
        _1017_v3 = _nw184;
        _1017_v3 = _1017_v3;
        let _1018_v4;
        _1018_v4 = _dafny.Seq.of(p0, _1016_i0, (_this).f3);
        let _1019_v5;
        _1019_v5 = _dafny.Set.fromElements(_1018_v4, _1018_v4);
        let _1020_v6;
        _1020_v6 = _dafny.Map.Empty.slice().updateUnsafe(_1016_i0,_dafny.Set.fromElements(_this.f4, !(false)));
        let _1021_v7;
        _1021_v7 = _dafny.Seq.UnicodeFromString("efq");
        let _1022_v8;
        let _nw185 = Array((new BigNumber(17)).toNumber());
        _nw185[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(606)), ((_1023_i0) => function (_1024_i1) {
          return (_dafny.ZERO).minus(_1023_i0);
        })(_1016_i0))).length);
        _nw185[(_dafny.ONE).toNumber()] = p0;
        _nw185[(new BigNumber(2)).toNumber()] = p0;
        _nw185[(new BigNumber(3)).toNumber()] = new BigNumber((_1019_v5).length);
        _nw185[(new BigNumber(4)).toNumber()] = new BigNumber(-369);
        _nw185[(new BigNumber(5)).toNumber()] = (_this).f3;
        _nw185[(new BigNumber(6)).toNumber()] = (_this).f3;
        _nw185[(new BigNumber(7)).toNumber()] = p0;
        _nw185[(new BigNumber(8)).toNumber()] = p0;
        _nw185[(new BigNumber(9)).toNumber()] = p0;
        _nw185[(new BigNumber(10)).toNumber()] = new BigNumber((_1020_v6).length);
        _nw185[(new BigNumber(11)).toNumber()] = new BigNumber(687);
        _nw185[(new BigNumber(12)).toNumber()] = (_this).f7;
        _nw185[(new BigNumber(13)).toNumber()] = _1016_i0;
        _nw185[(new BigNumber(14)).toNumber()] = (_this).f3;
        _nw185[(new BigNumber(15)).toNumber()] = (_this).f3;
        _nw185[(new BigNumber(16)).toNumber()] = new BigNumber((_1021_v7).length);
        _1022_v8 = _nw185;
        let _1025_v9;
        let _nw186 = new _module.C2();
        _nw186.__ctor((_1016_i0).plus(_1016_i0), _1022_v8, (_this).f7, _this.f4);
        _1025_v9 = _nw186;
        let _1026_v10;
        _1026_v10 = _dafny.Seq.of(_this.f4);
        let _1027_v11;
        _1027_v11 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1026_v10).length),_1021_v7);
        let _1028_v12;
        _1028_v12 = _dafny.Seq.of(_this.f4, _this.f4, _this.f4, !(_1027_v11).equals(_1027_v11));
        let _1029_v13;
        _1029_v13 = _dafny.Seq.of(_1028_v12);
        _1028_v12 = (_1029_v13)[_module.__default.safeIndex((_this).f7, new BigNumber((_1029_v13).length))];
        let _1030_v14;
        let _init30 = function (_1031_i2) {
          return _module.D6.create_DC18(false, (_this).f7);
        };
        let _nw187 = Array((new BigNumber(3)).toNumber());
        for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw187.length); _i0_30++) {
          _nw187[_i0_30] = _init30(new BigNumber(_i0_30));
        }
        _1030_v14 = _nw187;
        let _index184 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1030_v14).length));
        (_1030_v14)[_index184] = _module.__default.fm29(globalState);
        let _1032_v15;
        _1032_v15 = _module.D6.create_DC18(!(!(((_this.f4) ? (_this.f4) : (_this.f4)))), ((_1025_v9).f14).minus(p0));
        let _index185 = _module.__default.safeIndex(new BigNumber(98), new BigNumber((_1030_v14).length));
        (_1030_v14)[_index185] = _1032_v15;
      }
      let _1033_v16;
      _1033_v16 = new _dafny.CodePoint('s'.codePointAt(0));
      let _1034_v17;
      _1034_v17 = _dafny.MultiSet.fromElements(_1033_v16);
      let _1035_v18;
      _1035_v18 = _dafny.Seq.of((_this).f7);
      let _1036_v19;
      _1036_v19 = _dafny.Map.Empty.slice().updateUnsafe(_1035_v18,_1033_v16);
      let _1037_v20;
      _1037_v20 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1036_v19).length),_this.f4);
      let _1038_v21;
      _1038_v21 = _dafny.MultiSet.fromElements((((_1034_v17).contains(_1033_v16)) ? ((_1034_v17).get(_1033_v16)) : ((_this).f7)), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f4,(((_1037_v20).contains(new BigNumber(-496))) ? ((_1037_v20).get(new BigNumber(-496))) : (_this.f4)))).length));
      let _1039_i3;
      _1039_i3 = _dafny.ZERO;
      L9: {
        while ((_1038_v21).IsDisjointFrom(_1038_v21)) {
          C9: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1039_i3)) {
              break L9;
            }
            _1039_i3 = (_1039_i3).plus(_dafny.ONE);
            let _1040_v22;
            _1040_v22 = _dafny.Set.fromElements(_this.f4, _this.f4);
            let _1041_v23;
            _1041_v23 = _module.D6.create_DC19(new BigNumber((_1040_v22).length));
            let _pat_let_tv18 = p0;
            let _1042_v24;
            _1042_v24 = _dafny.Set.fromElements(_1041_v23, _1041_v23, _1041_v23, function (_pat_let26_0) {
              return function (_1043_dt__update__tmp_h0) {
                return function (_pat_let27_0) {
                  return function (_1044_dt__update_hcf26_h0) {
                    return _module.D6.create_DC19(_1044_dt__update_hcf26_h0);
                  }(_pat_let27_0);
                }((_dafny.ZERO).minus(_pat_let_tv18));
              }(_pat_let26_0);
            }(_1041_v23), _module.D6.create_DC19((_this).f3));
            _1042_v24 = _1042_v24;
            let _1045_v25;
            _1045_v25 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,((_dafny.ZERO).minus(p0)).plus(new BigNumber(829)));
            _1045_v25 = (_1045_v25).update(((_dafny.ZERO).minus((_this).f7)).isLessThanOrEqualTo(p0), (_this).f3);
            r0 = p0;
            let _1046_v26;
            _1046_v26 = _dafny.Seq.UnicodeFromString("rkni");
            let _1047_v27;
            _1047_v27 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1046_v26);
            _1047_v27 = (_1047_v27).update((_this).f3, _dafny.Seq.Concat((((_1047_v27).contains(p0)) ? ((_1047_v27).get(p0)) : (_1046_v26)), _1046_v26));
          }
        }
      }
      (_this).f4 = _this.f4;
      let _1048_v28;
      let _init31 = function (_1049_i4) {
        return _module.__default.safeModuloInt(_1049_i4, new BigNumber(453));
      };
      let _nw188 = Array((new BigNumber(9)).toNumber());
      for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw188.length); _i0_31++) {
        _nw188[_i0_31] = _init31(new BigNumber(_i0_31));
      }
      _1048_v28 = _nw188;
      let _1050_v29;
      _1050_v29 = _dafny.Set.fromElements((_this).f3);
      let _1051_v30;
      _1051_v30 = _dafny.Seq.UnicodeFromString("hnjykaw");
      let _nw189 = Array((new BigNumber(8)).toNumber());
      _nw189[(_dafny.ZERO).toNumber()] = p0;
      _nw189[(_dafny.ONE).toNumber()] = p0;
      _nw189[(new BigNumber(2)).toNumber()] = p0;
      _nw189[(new BigNumber(3)).toNumber()] = (_this).f7;
      _nw189[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_1050_v29).length), new BigNumber((_1051_v30).length));
      _nw189[(new BigNumber(5)).toNumber()] = ((_this).f3).multipliedBy((_this).f3);
      _nw189[(new BigNumber(6)).toNumber()] = new BigNumber(327);
      _nw189[(new BigNumber(7)).toNumber()] = (_this).f7;
      _1048_v28 = _nw189;
      let _rhs188 = (_this).f7;
      let _rhs189 = _this.f4;
      let _rhs190 = _1051_v30;
      let _lhs176 = globalState;
      let _lhs177 = globalState;
      _lhs176.f1 = _rhs188;
      _lhs177.f2 = _rhs189;
      _1051_v30 = _rhs190;
      r0 = new BigNumber(195);
      return r0;
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.MultiSet.Empty;
      let r1 = new _dafny.CodePoint('D'.codePointAt(0));
      let r2 = _dafny.Map.Empty;
      let r3 = false;
      let _1052_v0;
      _1052_v0 = _dafny.Seq.of(_this.f4);
      let _1053_v1;
      _1053_v1 = _module.D1.create_DC7();
      let _1054_v2;
      _1054_v2 = _module.D0.create_DC0(_this.f4);
      (globalState).f1 = (((_this).f7).minus(new BigNumber((_1052_v0).length))).multipliedBy(new BigNumber((_module.__default.fm8(_this.f4, _1053_v1, _1054_v2, (_this).f3, globalState)).length));
      let _1055_v3;
      let _init32 = ((_1056_p0) => function (_1057_i1) {
        return _1056_p0;
      })(p0);
      let _nw190 = Array((new BigNumber(9)).toNumber());
      for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw190.length); _i0_32++) {
        _nw190[_i0_32] = _init32(new BigNumber(_i0_32));
      }
      _1055_v3 = _nw190;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1055_v3).length))) {
        let _1058_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1058_i0)) && ((_1058_i0).isLessThan(new BigNumber((_1055_v3).length))))) {
          (_1055_v3)[(_1058_i0)] = (_1052_v0)[_module.__default.safeIndex(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,p3)).length), new BigNumber((_1052_v0).length))];
        }
      }
      let _1059_v4;
      _1059_v4 = _dafny.Seq.UnicodeFromString("lgks");
      let _1060_v5;
      _1060_v5 = new _dafny.CodePoint('f'.codePointAt(0));
      let _1061_v6;
      let _nw191 = new _module.C3();
      _nw191.__ctor(_1060_v5, _1059_v4, (_this).f7, _this.f4, (_this).f3);
      _1061_v6 = _nw191;
      let _1062_v7;
      _1062_v7 = _module.D1.create_DC6(_1061_v6, _1059_v4, _1061_v6.f4, p3);
      let _rhs191 = _1059_v4;
      let _rhs192 = _dafny.Seq.Concat((_1062_v7).dtor_cf10, _dafny.Seq.update(_dafny.Seq.UnicodeFromString("uuopatwnv"), _module.__default.safeIndex(new BigNumber(840), new BigNumber((_dafny.Seq.UnicodeFromString("uuopatwnv")).length)), _1060_v5));
      _1059_v4 = _rhs191;
      _1059_v4 = _rhs192;
      let _rhs193 = (_1061_v6).f3;
      let _rhs194 = _1052_v0;
      let _rhs195 = (_1061_v6).f3;
      let _lhs178 = globalState;
      let _lhs179 = globalState;
      _lhs178.f1 = _rhs193;
      _1052_v0 = _rhs194;
      _lhs179.f1 = _rhs195;
      let _1063_i2;
      _1063_i2 = _dafny.ZERO;
      L10: {
        while (_this.f4) {
          C10: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_1063_i2)) {
              break L10;
            }
            _1063_i2 = (_1063_i2).plus(_dafny.ONE);
            let _1064_v8;
            _1064_v8 = _dafny.Map.Empty.slice().updateUnsafe((_1061_v6).f3,_1061_v6.f4);
            let _1065_v9;
            _1065_v9 = _module.D8.create_DC22(_1064_v8);
            let _1066_v10;
            let _nw192 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
            _1066_v10 = _nw192;
            let _1067_v11;
            _1067_v11 = _dafny.Seq.of(_1066_v10);
            let _1068_v12;
            _1068_v12 = _module.D12.create_DC34(_1066_v10);
            let _1069_v13;
            let _nw193 = Array((new BigNumber(24)).toNumber());
            _nw193[(_dafny.ZERO).toNumber()] = _1066_v10;
            _nw193[(_dafny.ONE).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(2)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(3)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(4)).toNumber()] = (_1067_v11)[_module.__default.safeIndex((_1061_v6).f3, new BigNumber((_1067_v11).length))];
            _nw193[(new BigNumber(5)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(6)).toNumber()] = ((_1061_v6.f4) ? (_1066_v10) : (_1066_v10));
            _nw193[(new BigNumber(7)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(8)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(9)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(10)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(11)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(12)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(13)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(14)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(15)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(16)).toNumber()] = (_module.D12.create_DC34(_1066_v10)).dtor_cf55;
            _nw193[(new BigNumber(17)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(18)).toNumber()] = (_1068_v12).dtor_cf55;
            _nw193[(new BigNumber(19)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(20)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(21)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(22)).toNumber()] = _1066_v10;
            _nw193[(new BigNumber(23)).toNumber()] = _1066_v10;
            _1069_v13 = _nw193;
            let _index186 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_1069_v13).length));
            (_1069_v13)[_index186] = _1066_v10;
            let _index187 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1055_v3).length));
            (_1055_v3)[_index187] = p3;
            let _1070_v14;
            _1070_v14 = _dafny.MultiSet.fromElements((_this).f3, (_dafny.ZERO).minus((_this).f7));
            let _pat_let_tv19 = _1064_v8;
            let _pat_let_tv20 = _1061_v6;
            let _index188 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_1069_v13).length));
            let _index189 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1055_v3).length));
            let _rhs196 = function (_pat_let28_0) {
              return function (_1071_dt__update__tmp_h0) {
                return function (_pat_let29_0) {
                  return function (_1072_dt__update_hcf31_h0) {
                    return _module.D8.create_DC22(_1072_dt__update_hcf31_h0);
                  }(_pat_let29_0);
                }((_pat_let_tv19).update((_pat_let_tv20).f3, _this.f4));
              }(_pat_let28_0);
            }(_1065_v9);
            let _rhs197 = _1066_v10;
            let _rhs198 = !((new BigNumber((_1070_v14).cardinality())).isEqualTo(((_1061_v6.f4) ? ((_this).f7) : (_module.__default.fm0(globalState)))));
            let _rhs199 = ((p3) ? (p3) : (p3));
            let _lhs180 = _1069_v13;
            let _lhs181 = _module.__default.safeIndex(new BigNumber(69), new BigNumber((_1069_v13).length));
            let _lhs182 = _1055_v3;
            let _lhs183 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1055_v3).length));
            let _lhs184 = _1061_v6;
            _1065_v9 = _rhs196;
            _lhs180[_lhs181] = _rhs197;
            _lhs182[_lhs183] = _rhs198;
            _lhs184.f4 = _rhs199;
            _1059_v4 = _dafny.Seq.UnicodeFromString("ncpxl");
            let _1073_v15;
            _1073_v15 = _dafny.Seq.of(p2);
            let _1074_v16;
            _1074_v16 = _module.D13.create_DC38(_1073_v15);
            let _index190 = _module.__default.safeIndex(new BigNumber(745), new BigNumber((_1055_v3).length));
            (_1055_v3)[_index190] = _dafny.Seq.IsPrefixOf((_1074_v16).dtor_cf68, _dafny.Seq.of(p2));
            let _1075_v17;
            let _1076_v18;
            let _1077_v19;
            let _1078_v20;
            let _out10;
            let _out11;
            let _out12;
            let _out13;
            let _outcollector2 = (_this).m3((_module.D13.create_DC39(_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm2((_1061_v6).f3, globalState),(_this).f7))).dtor_cf69, new BigNumber((_dafny.Seq.Concat(_1059_v4, _1059_v4)).length), globalState);
            _out10 = _outcollector2[0];
            _out11 = _outcollector2[1];
            _out12 = _outcollector2[2];
            _out13 = _outcollector2[3];
            _1075_v17 = _out10;
            _1076_v18 = _out11;
            _1077_v19 = _out12;
            _1078_v20 = _out13;
          }
        }
      }
      let _1079_v21;
      _1079_v21 = _dafny.Seq.of(new BigNumber(534));
      let _1080_v22;
      _1080_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1079_v21,_1059_v4);
      let _1081_v23;
      _1081_v23 = _module.D10.create_DC28(_1080_v22);
      let _pat_let_tv21 = _1059_v4;
      let _pat_let_tv22 = _1060_v5;
      let _pat_let_tv23 = _1080_v22;
      (globalState).f1 = new BigNumber((function (_source13) {
        if (_source13.is_DC29) {
          let _1082___mcc_h0 = (_source13).cf42;
          let _1083___mcc_h1 = (_source13).cf43;
          let _1084_cf43 = _1083___mcc_h1;
          let _1085_cf42 = _1082___mcc_h0;
          return _pat_let_tv21;
        } else if (_source13.is_DC30) {
          let _1086___mcc_h2 = (_source13).cf44;
          let _1087___mcc_h3 = (_source13).cf45;
          let _1088___mcc_h4 = (_source13).cf46;
          let _1089___mcc_h5 = (_source13).cf47;
          let _1090___mcc_h6 = (_source13).cf48;
          let _1091_cf48 = _1090___mcc_h6;
          let _1092_cf47 = _1089___mcc_h5;
          let _1093_cf46 = _1088___mcc_h4;
          let _1094_cf45 = _1087___mcc_h3;
          let _1095_cf44 = _1086___mcc_h2;
          return _dafny.Seq.UnicodeFromString("whfgm");
        } else {
          let _1096___mcc_h7 = (_source13).cf41;
          let _1097_cf41 = _1096___mcc_h7;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(894)), ((_1098_v5) => function (_1099_i3) {
            return _1098_v5;
          })(_pat_let_tv22));
        }
      }(function (_pat_let30_0) {
        return function (_1100_dt__update__tmp_h1) {
          return function (_pat_let31_0) {
            return function (_1101_dt__update_hcf41_h0) {
              return _module.D10.create_DC28(_1101_dt__update_hcf41_h0);
            }(_pat_let31_0);
          }(_pat_let_tv23);
        }(_pat_let30_0);
      }(_1081_v23))).length);
      let _1102_v24;
      _1102_v24 = _module.D12.create_DC36(p0, _1061_v6.f4, (_1061_v6).f3);
      let _1103_v25;
      _1103_v25 = _dafny.MultiSet.fromElements((((_1102_v24).dtor_cf61) ? (_this.f4) : (p3)));
      r0 = _1103_v25;
      let _1104_v26;
      _1104_v26 = _dafny.MultiSet.fromElements(new BigNumber((_1103_v25).cardinality()));
      r1 = ((!(_1104_v26).equals(_1104_v26)) ? (_1060_v5) : (new _dafny.CodePoint('o'.codePointAt(0))));
      r2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_1060_v5);
      r3 = _this.f4;
      return [r0, r1, r2, r3];
    }
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let r3 = _dafny.ZERO;
      let _hi10 = _module.__default.fm0(globalState);
      for (let _1105_i0 = (_dafny.ZERO).minus((_this).f3); _1105_i0.isLessThan(_hi10); _1105_i0 = _1105_i0.plus(_dafny.ONE)) {
        let _1106_v0;
        let _nw194 = Array((new BigNumber(4)).toNumber()).fill([]);
        _1106_v0 = _nw194;
        let _rhs200 = _1106_v0;
        let _rhs201 = ((_this).f3).plus(((_module.__default.fm30(_this.f4, (_this).f7, (_this).f3, globalState)).dtor_cf66).minus((_this).f3));
        _1106_v0 = _rhs200;
        r1 = _rhs201;
        let _1107_v1;
        _1107_v1 = _dafny.Seq.of(_this.f4);
        let _1108_v2;
        _1108_v2 = _dafny.Seq.UnicodeFromString("epm");
        let _1109_v3;
        _1109_v3 = _dafny.Seq.of(true, (_1107_v1)[_module.__default.safeIndex(new BigNumber((_1108_v2).length), new BigNumber((_1107_v1).length))]);
        _1107_v1 = _1107_v1;
        let _1110_v4;
        _1110_v4 = _dafny.Set.fromElements(_this.f4, _this.f4, _this.f4, _this.f4, _this.f4);
        let _1111_v5;
        _1111_v5 = _dafny.Seq.of((new BigNumber((_dafny.Seq.of((_this).f9)).length)).minus((_this).f3));
        let _rhs202 = ((_1110_v4).Difference(_1110_v4)).Difference(_1110_v4);
        let _rhs203 = (_dafny.ZERO).minus((_dafny.ZERO).minus((_1105_i0).multipliedBy(_1105_i0)));
        let _rhs204 = (_this.f4) || (true);
        let _rhs205 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1111_v5, _1111_v5), _1111_v5);
        let _lhs185 = globalState;
        let _lhs186 = globalState;
        _1110_v4 = _rhs202;
        _lhs185.f1 = _rhs203;
        _lhs186.f2 = _rhs204;
        _1111_v5 = _rhs205;
        let _1112_v6;
        let _init33 = ((_1113_v2) => function (_1114_i1) {
          return (_1114_i1).plus(new BigNumber((_1113_v2).length));
        })(_1108_v2);
        let _nw195 = Array((new BigNumber(19)).toNumber());
        for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw195.length); _i0_33++) {
          _nw195[_i0_33] = _init33(new BigNumber(_i0_33));
        }
        _1112_v6 = _nw195;
        let _index191 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1112_v6).length));
        (_1112_v6)[_index191] = ((_this).f3).multipliedBy((_this).f3);
        let _index192 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1112_v6).length));
        let _rhs206 = _module.__default.fm0(globalState);
        let _rhs207 = (_1110_v4).IsSubsetOf(_dafny.Set.fromElements(_this.f4));
        let _rhs208 = (_dafny.ZERO).minus((_1105_i0).minus(p1));
        let _lhs187 = globalState;
        let _lhs188 = _1112_v6;
        let _lhs189 = _module.__default.safeIndex(new BigNumber(468), new BigNumber((_1112_v6).length));
        _lhs187.f1 = _rhs206;
        r0 = _rhs207;
        _lhs188[_lhs189] = _rhs208;
      }
      let _hi11 = (_this).f3;
      for (let _1115_i2 = p1; _1115_i2.isLessThan(_hi11); _1115_i2 = _1115_i2.plus(_dafny.ONE)) {
        let _index193 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f9).length));
        ((_this).f9)[_index193] = (new BigNumber((_dafny.Seq.UnicodeFromString("hpexysybx")).length)).isLessThanOrEqualTo((_this).f7);
        let _1116_v7;
        _1116_v7 = _dafny.MultiSet.fromElements(_this.f4, _this.f4, _this.f4);
        let _1117_v8;
        _1117_v8 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt((_this).f7, (_this).f7),_module.__default.fm2(_1115_i2, globalState));
        let _1118_v9;
        _1118_v9 = _dafny.Seq.UnicodeFromString("he");
        let _index194 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f9).length));
        let _rhs209 = (_1116_v7).IsSubsetOf(_module.__default.fm14(_this.f4, false, globalState));
        let _rhs210 = (_this).f7;
        let _rhs211 = (((_1117_v8).contains((_this).f7)) ? ((_1117_v8).get((_this).f7)) : (!(_this.f4)));
        let _rhs212 = ((_this).f7).multipliedBy(new BigNumber((_1118_v9).length));
        let _rhs213 = (((_this.f4) ? ((_this).f7) : (p1))).minus((_dafny.ZERO).minus((_this).f3));
        let _lhs190 = (_this).f9;
        let _lhs191 = _module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f9).length));
        let _lhs192 = globalState;
        let _lhs193 = globalState;
        _lhs190[_lhs191] = _rhs209;
        r1 = _rhs210;
        _lhs192.f2 = _rhs211;
        _lhs193.f1 = _rhs212;
        r3 = _rhs213;
        let _1119_v10;
        let _nw196 = Array((new BigNumber(2)).toNumber()).fill([]);
        _1119_v10 = _nw196;
        let _1120_v11;
        _1120_v11 = _module.D6.create_DC18(((_this).f9)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f9).length))], new BigNumber((_1118_v9).length));
        let _1121_v12;
        let _nw197 = Array((new BigNumber(5)).toNumber());
        _nw197[(_dafny.ZERO).toNumber()] = ((_this).f9)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f9).length))];
        _nw197[(_dafny.ONE).toNumber()] = _this.f4;
        _nw197[(new BigNumber(2)).toNumber()] = _this.f4;
        _nw197[(new BigNumber(3)).toNumber()] = (_1120_v11).dtor_cf24;
        _nw197[(new BigNumber(4)).toNumber()] = !(_module.__default.fm2(p1, globalState));
        _1121_v12 = _nw197;
        let _index195 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_1119_v10).length));
        (_1119_v10)[_index195] = _1121_v12;
        let _1122_v13;
        _1122_v13 = _dafny.Seq.of(true, ((_this).f9)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f9).length))], ((_this).f9)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f9).length))], !(((_this).f9)[_module.__default.safeIndex(new BigNumber(104), new BigNumber(((_this).f9).length))]));
        let _1123_v14;
        _1123_v14 = _module.D12.create_DC35(_1121_v12, _this.f4, _1122_v13, _dafny.Seq.UnicodeFromString("tgykad"), p1);
        let _index196 = _module.__default.safeIndex(new BigNumber(923), new BigNumber((_1119_v10).length));
        (_1119_v10)[_index196] = (_1123_v14).dtor_cf56;
        r1 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt((_this).f7, (_this).f3), (_dafny.ZERO).minus(p1));
        let _1124_v15;
        _1124_v15 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,false);
        let _1125_v16;
        let _nw198 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Map.Empty);
        _1125_v16 = _nw198;
        let _1126_v17;
        _1126_v17 = _module.D12.create_DC34(_1125_v16);
        let _1127_v18;
        _1127_v18 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1124_v15).length),_1126_v17);
        let _1128_v19;
        _1128_v19 = _dafny.MultiSet.fromElements(_1127_v18);
        _1128_v19 = ((false) ? (((true) ? (_1128_v19) : (_1128_v19))) : (_1128_v19));
      }
      let _1129_v20;
      _1129_v20 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_dafny.Seq.UnicodeFromString("q"));
      let _1130_v21;
      _1130_v21 = _module.D7.create_DC20(_1129_v20);
      let _1131_v22;
      _1131_v22 = _dafny.MultiSet.fromElements(_1130_v21, _1130_v21);
      if (((_1131_v22).Union(_1131_v22)).contains(((_this.f4) ? (_module.D7.create_DC20(_1129_v20)) : (_1130_v21)))) {
        let _1132_v23;
        _1132_v23 = new _dafny.CodePoint('v'.codePointAt(0));
        let _1133_v24;
        _1133_v24 = _dafny.Set.fromElements(_1132_v23);
        let _1134_v25;
        _1134_v25 = _dafny.Map.Empty.slice().updateUnsafe(_1133_v24,_1132_v23);
        let _1135_v26;
        let _nw199 = new _module.C1();
        _nw199.__ctor(new BigNumber((_1134_v25).length), (_this).f3, _this.f4);
        _1135_v26 = _nw199;
        let _1136_v27;
        _1136_v27 = _dafny.Set.fromElements(_1135_v26, _1135_v26, _1135_v26, _1135_v26);
        _1136_v27 = (_dafny.Set.fromElements(_1135_v26, _1135_v26, _1135_v26)).Union(_dafny.Set.fromElements(_1135_v26));
        let _1137_v28;
        let _nw200 = new _module.C4();
        _nw200.__ctor((_this).f3, (_this).f3, ((_this).f7).isLessThanOrEqualTo(p1));
        _1137_v28 = _nw200;
        r1 = (_this).f7;
        let _1138_v29;
        let _nw201 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
        _1138_v29 = _nw201;
        let _1139_v30;
        _1139_v30 = _dafny.Seq.of(_this.f4);
        let _1140_v31;
        _1140_v31 = _dafny.Seq.UnicodeFromString("s");
        let _1141_v32;
        _1141_v32 = _module.D12.create_DC35((_this).f9, true, _1139_v30, _1140_v31, (_this).f3);
        let _1142_v33;
        _1142_v33 = _dafny.Seq.of((_1141_v32).dtor_cf57);
        let _index197 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1138_v29).length));
        (_1138_v29)[_index197] = _1139_v30;
        let _index198 = _module.__default.safeIndex(new BigNumber(652), new BigNumber((_1138_v29).length));
        (_1138_v29)[_index198] = _1139_v30;
        let _1143_v34;
        let _nw202 = Array((new BigNumber(2)).toNumber()).fill(false);
        _1143_v34 = _nw202;
        _1143_v34 = (_this).f9;
      } else {
        let _1144_v35;
        _1144_v35 = _dafny.Set.fromElements(_this.f4);
        let _1145_v36;
        _1145_v36 = _dafny.Seq.UnicodeFromString("tjqkmu");
        let _rhs214 = ((_1144_v35).Intersect(_dafny.Set.fromElements(false, _this.f4))).Union(_1144_v35);
        let _rhs215 = (((_this.f4) ? (p1) : ((_this).f7))).isLessThanOrEqualTo(new BigNumber((_dafny.Seq.Concat(_1145_v36, _dafny.Seq.UnicodeFromString("spyelrt"))).length));
        let _lhs194 = _this;
        _1144_v35 = _rhs214;
        _lhs194.f4 = _rhs215;
        let _1146_v37;
        _1146_v37 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_this.f4);
        _1146_v37 = (_1146_v37).update(_this.f4, ((_this).f3).isEqualTo(p1));
        let _1147_v38;
        _1147_v38 = _dafny.MultiSet.fromElements((_this).f7, (new BigNumber(362)).plus((_this).f3), (_this).f7);
        _1147_v38 = ((_1147_v38).update((_dafny.ZERO).minus((_this).f3), _module.__default.abs(p1))).update((((p0).contains(_this.f4)) ? ((p0).get(_this.f4)) : (p1)), _module.__default.abs((_dafny.ZERO).minus((p1).minus(p1))));
        let _1148_v39;
        _1148_v39 = new _dafny.CodePoint('o'.codePointAt(0));
        let _1149_v40;
        _1149_v40 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_1145_v36, _module.__default.safeIndex(p1, new BigNumber((_1145_v36).length)), _1148_v39),(_this).f7);
        let _1150_v41;
        _1150_v41 = _module.D12.create_DC37(_1146_v37, _this.f4, (((_1149_v40).contains(_dafny.Seq.UnicodeFromString("cmrcf"))) ? ((_1149_v40).get(_dafny.Seq.UnicodeFromString("cmrcf"))) : ((_this).f3)), (_this).f3);
        r3 = (_1150_v41).dtor_cf67;
        let _1151_v42;
        _1151_v42 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f3).plus(new BigNumber(-416)),(_this).f7);
        _1151_v42 = (_1151_v42).update((_this).f3, (_this).f3);
      }
      let _1152_v43;
      _1152_v43 = _dafny.Seq.of(_this.f4);
      (globalState).f1 = (((_1152_v43)[_module.__default.safeIndex((_this).f7, new BigNumber((_1152_v43).length))]) ? (_module.__default.safeModuloInt(p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-567)), function (_1153_i3) {
        return new _dafny.CodePoint('s'.codePointAt(0));
      })).length))) : (new BigNumber(762)));
      let _1154_v44;
      let _nw203 = new _module.C4();
      _nw203.__ctor((_dafny.ZERO).minus(((_this).f7).minus((_this).f3)), (_dafny.ZERO).minus(_module.__default.fm0(globalState)), _module.__default.fm2(new BigNumber((_1152_v43).length), globalState));
      _1154_v44 = _nw203;
      let _1155_v45;
      let _nw204 = Array((new BigNumber(13)).toNumber()).fill(_dafny.ZERO);
      _1155_v45 = _nw204;
      _1155_v45 = _1155_v45;
      r0 = _this.f4;
      let _1156_v46;
      _1156_v46 = _dafny.Seq.UnicodeFromString("gtmwd");
      r1 = new BigNumber((_1156_v46).length);
      r2 = _module.__default.fm15(_this.f4, _module.__default.safeModuloInt(p1, (_this).f3), new BigNumber(639), globalState);
      r3 = (_dafny.ZERO).minus((_this).f3);
      return [r0, r1, r2, r3];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f4 = false;
      this._f7 = _dafny.ZERO;
      this._f3 = _dafny.ZERO;
      this._f8 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f8, f3, f4, f7) {
      let _this = this;
      (_this)._f8 = f8;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this)._f7 = f7;
      return;
    }
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      if (_this.f4) {
        let _1157_v0;
        _1157_v0 = _dafny.Seq.of(_this.f4);
        let _1158_v1;
        _1158_v1 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,new BigNumber((_1157_v0).length));
        _1158_v1 = (_1158_v1).update((_this.f4) && (_this.f4), ((_this.f4) ? ((_this).f3) : ((_this).f3)));
        let _1159_v2;
        let _nw205 = new _module.C4();
        _nw205.__ctor(_module.__default.safeModuloInt(p0, new BigNumber((_dafny.Seq.UnicodeFromString("qvoflbfav")).length)), ((_this).f3).multipliedBy((_this).f3), _this.f4);
        _1159_v2 = _nw205;
        let _1160_v3;
        let _nw206 = Array((new BigNumber(21)).toNumber()).fill(false);
        _1160_v3 = _nw206;
        let _1161_v4;
        let _nw207 = Array((new BigNumber(4)).toNumber());
        _nw207[(_dafny.ZERO).toNumber()] = _1160_v3;
        _nw207[(_dafny.ONE).toNumber()] = _1160_v3;
        _nw207[(new BigNumber(2)).toNumber()] = _1160_v3;
        _nw207[(new BigNumber(3)).toNumber()] = _1160_v3;
        _1161_v4 = _nw207;
        _1161_v4 = _1161_v4;
        (globalState).f1 = ((_this).f7).plus((_dafny.ZERO).minus((_dafny.ZERO).minus((_this).f8)));
        let _1162_v5;
        let _nw208 = new _module.C1();
        _nw208.__ctor(_module.__default.fm0(globalState), (_this).f8, _this.f4);
        _1162_v5 = _nw208;
      } else {
        let _1163_v6;
        _1163_v6 = _dafny.Seq.UnicodeFromString("kbtb");
        let _1164_v7;
        _1164_v7 = _dafny.MultiSet.fromElements(_dafny.Seq.Concat(_1163_v6, _1163_v6), _1163_v6);
        let _1165_v8;
        _1165_v8 = _dafny.Seq.of(_1163_v6, _1163_v6);
        _1164_v7 = (((_this.f4) ? (_1164_v7) : ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("oeofpknod"), _1163_v6, _dafny.Seq.UnicodeFromString("ou"), _1163_v6)).update(_1163_v6, _module.__default.abs(p0))))).Difference((_1164_v7).Intersect((_dafny.MultiSet.FromArray(_1165_v8)).update(_1163_v6, _module.__default.abs(p0))));
        let _1166_v9;
        let _nw209 = Array((new BigNumber(13)).toNumber()).fill(false);
        _1166_v9 = _nw209;
        let _index199 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1166_v9).length));
        (_1166_v9)[_index199] = _this.f4;
        let _index200 = _module.__default.safeIndex(new BigNumber(608), new BigNumber((_1166_v9).length));
        (_1166_v9)[_index200] = _this.f4;
        let _1167_v10;
        _1167_v10 = new _dafny.CodePoint('o'.codePointAt(0));
        _1167_v10 = _1167_v10;
        let _1168_v11;
        let _nw210 = Array((new BigNumber(2)).toNumber()).fill(_dafny.Map.Empty);
        _1168_v11 = _nw210;
        let _index201 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1168_v11).length));
        (_1168_v11)[_index201] = _dafny.Map.Empty.slice().updateUnsafe(p0,(_1166_v9)[_module.__default.safeIndex(new BigNumber(608), new BigNumber((_1166_v9).length))]);
        let _1169_v12;
        _1169_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,_module.__default.fm2(p0, globalState));
        let _index202 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1168_v11).length));
        (_1168_v11)[_index202] = (_1169_v12).Merge(_1169_v12);
        let _rhs216 = (((_1166_v9)[_module.__default.safeIndex(new BigNumber(608), new BigNumber((_1166_v9).length))]) ? ((_this).f3) : ((_this).f7));
        let _rhs217 = _dafny.Seq.update(_dafny.Seq.Concat(_1163_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), ((_1170_v10) => function (_1171_i0) {
          return _1170_v10;
        })(_1167_v10))), _module.__default.safeIndex(_module.__default.safeModuloInt((_this).f8, p0), new BigNumber((_dafny.Seq.Concat(_1163_v6, _dafny.Seq.Create(_module.__default.abs(new BigNumber(588)), ((_1172_v10) => function (_1173_i0) {
          return _1172_v10;
        })(_1167_v10)))).length)), _1167_v10);
        let _rhs218 = (_1166_v9)[_module.__default.safeIndex(new BigNumber(608), new BigNumber((_1166_v9).length))];
        let _rhs219 = _dafny.Seq.Concat(_module.__default.fm20((_this).f8, globalState), _1163_v6);
        let _rhs220 = true;
        let _lhs195 = globalState;
        let _lhs196 = globalState;
        let _lhs197 = globalState;
        _lhs195.f1 = _rhs216;
        _1163_v6 = _rhs217;
        _lhs196.f2 = _rhs218;
        _1163_v6 = _rhs219;
        _lhs197.f2 = _rhs220;
      }
      let _1174_v13;
      let _init34 = function (_1175_i1) {
        return _this.f4;
      };
      let _nw211 = Array((new BigNumber(4)).toNumber());
      for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw211.length); _i0_34++) {
        _nw211[_i0_34] = _init34(new BigNumber(_i0_34));
      }
      _1174_v13 = _nw211;
      let _1176_v14;
      let _nw212 = new _module.C5();
      _nw212.__ctor(_1174_v13, ((!(_this.f4)) ? ((_this).f8) : ((_this).f8)), (_this).f7, !(_this.f4));
      _1176_v14 = _nw212;
      let _1177_v15;
      _1177_v15 = _dafny.Seq.UnicodeFromString("euk");
      let _1178_v16;
      _1178_v16 = _dafny.Seq.of(_1177_v15, _1177_v15);
      if (!(!_dafny.Seq.contains(_dafny.Seq.Concat(_1178_v16, _1178_v16), _1177_v15))) {
        let _1179_v17;
        _1179_v17 = _dafny.Map.Empty.slice().updateUnsafe(false,_this.f4);
        let _1180_v18;
        _1180_v18 = _dafny.Set.fromElements(new BigNumber((_1179_v17).length), (_this).f8, new BigNumber((_dafny.Set.fromElements((_this).f7, (_this).f7)).length));
        let _1181_v19;
        _1181_v19 = _module.D4.create_DC11(_1180_v18);
        let _source14 = _1181_v19;
        if (_source14.is_DC12) {
          let _1182_v20;
          let _init35 = function (_1183_i2) {
            return _module.__default.safeModuloInt(_1183_i2, (_this).f3);
          };
          let _nw213 = Array((new BigNumber(15)).toNumber());
          for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw213.length); _i0_35++) {
            _nw213[_i0_35] = _init35(new BigNumber(_i0_35));
          }
          _1182_v20 = _nw213;
          let _index203 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_1182_v20).length));
          (_1182_v20)[_index203] = (_this).f7;
          let _index204 = _module.__default.safeIndex(new BigNumber(61), new BigNumber((_1182_v20).length));
          (_1182_v20)[_index204] = _module.__default.fm0(globalState);
          let _1184_v21;
          let _nw214 = new _module.C2();
          _nw214.__ctor(_module.__default.safeModuloInt((_this).f3, new BigNumber(-240)), _1182_v20, (_this).f3, _this.f4);
          _1184_v21 = _nw214;
          (globalState).f2 = _this.f4;
          (globalState).f1 = new BigNumber(-741);
        } else if (_source14.is_DC11) {
          let _1185___mcc_h0 = (_source14).cf18;
          let _1186_cf18 = _1185___mcc_h0;
          (globalState).f1 = (new BigNumber((_dafny.Seq.UnicodeFromString("jacm")).length)).minus(_module.__default.safeModuloInt((_this).f8, (_this).f7));
          _1177_v15 = _dafny.Seq.UnicodeFromString("cmeqw");
          let _1187_v22;
          _1187_v22 = _dafny.Set.fromElements(_this.f4);
          let _1188_v23;
          _1188_v23 = _dafny.MultiSet.fromElements((_this).f3, p0, new BigNumber((_1187_v22).length));
          let _1189_v24;
          _1189_v24 = _dafny.Map.Empty.slice().updateUnsafe((_this).f3,_this.f4);
          let _1190_v25;
          _1190_v25 = _dafny.Seq.of(_this.f4, _this.f4, (((_1189_v24).contains(p0)) ? ((_1189_v24).get(p0)) : (_this.f4)));
          let _1191_v26;
          _1191_v26 = _module.D3.create_DC10(new BigNumber((_1177_v15).length), (_1176_v14).f9, (((_1188_v23).contains((_dafny.ZERO).minus((_this).f8))) ? ((_1188_v23).get((_dafny.ZERO).minus((_this).f8))) : (new BigNumber((_1190_v25).length))));
          let _1192_v28;
          let _nw215 = new _module.C5();
          _nw215.__ctor(((_this.f4) ? ((_1191_v26).dtor_cf16) : ((_1176_v14).f9)), new BigNumber((function () {
            let _coll20 = new _dafny.Set();
            for (const _compr_20 of (_1189_v24).Keys.Elements) {
              let _1193_v27 = _compr_20;
              if ((_1189_v24).contains(_1193_v27)) {
                _coll20.add((_1193_v27).plus(new BigNumber(-533)));
              }
            }
            return _coll20;
          }()).length), new BigNumber((_dafny.MultiSet.FromArray(_1190_v25)).cardinality()), _module.__default.fm2((_this).f8, globalState));
          _1192_v28 = _nw215;
          let _1194_v29;
          _1194_v29 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_1174_v13);
          let _1195_v30;
          let _nw216 = new _module.C0();
          _nw216.__ctor(_this.f4, (((_1194_v29).contains(_this.f4)) ? ((_1194_v29).get(_this.f4)) : ((_1176_v14).f9)));
          _1195_v30 = _nw216;
        } else {
          let _1196___mcc_h1 = (_source14).cf19;
          let _1197_cf19 = _1196___mcc_h1;
          let _1198_v31;
          let _init36 = function (_1199_i3) {
            return (_1199_i3).multipliedBy(new BigNumber(862));
          };
          let _nw217 = Array((new BigNumber(28)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw217.length); _i0_36++) {
            _nw217[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1198_v31 = _nw217;
          let _index205 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1198_v31).length));
          (_1198_v31)[_index205] = p0;
          let _1200_v32;
          _1200_v32 = new _dafny.CodePoint('j'.codePointAt(0));
          let _index206 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1198_v31).length));
          let _rhs221 = _module.__default.safeModuloInt(((_dafny.ZERO).minus((_this).f8)).plus((_this).f7), (_dafny.ZERO).minus(new BigNumber((_module.__default.fm9(_1200_v32, _this.f4, globalState)).length)));
          let _rhs222 = !(_this.f4);
          let _lhs198 = _1198_v31;
          let _lhs199 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1198_v31).length));
          let _lhs200 = globalState;
          _lhs198[_lhs199] = _rhs221;
          _lhs200.f2 = _rhs222;
          let _index207 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1198_v31).length));
          (_1198_v31)[_index207] = ((_this.f4) ? (p0) : (new BigNumber(819)));
          let _1201_v33;
          _1201_v33 = _dafny.Seq.of(((_this).f3).minus(p0), (_this).f3, (_this).f8);
          let _1202_v34;
          _1202_v34 = _dafny.Seq.of(_this.f4);
          let _1203_v35;
          _1203_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_1200_v32);
          let _1204_v36;
          _1204_v36 = _dafny.Set.fromElements(_1200_v32, (((_1203_v35).contains((_1198_v31)[_module.__default.safeIndex(new BigNumber(5), new BigNumber((_1198_v31).length))])) ? ((_1203_v35).get((_1198_v31)[_module.__default.safeIndex(new BigNumber(5), new BigNumber((_1198_v31).length))])) : (new _dafny.CodePoint('k'.codePointAt(0)))), _1200_v32);
          let _1205_v37;
          _1205_v37 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1204_v36).length),_this.f4);
          let _1206_v38;
          let _nw218 = new _module.C1();
          _nw218.__ctor((_this).f3, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(542)), function (_1207_i4) {
            return (_this).f8;
          })).length), _this.f4);
          _1206_v38 = _nw218;
          let _1208_v39;
          _1208_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(233),_1206_v38);
          let _1209_v40;
          _1209_v40 = _dafny.MultiSet.fromElements(((_this.f4) ? (_1206_v38) : (_1206_v38)), _1206_v38, _1206_v38, (((_1208_v39).contains((_this).f8)) ? ((_1208_v39).get((_this).f8)) : (_1206_v38)), _1206_v38);
          let _index208 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1198_v31).length));
          let _rhs223 = _dafny.Seq.update(_1201_v33, _module.__default.safeIndex(new BigNumber((((_this.f4) ? (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1202_v34).length),_this.f4)) : (_1205_v37))).length), new BigNumber((_1201_v33).length)), p0);
          let _rhs224 = (((_1209_v40).contains(((_this.f4) ? (_1206_v38) : (_1206_v38)))) ? ((_1209_v40).get(((_this.f4) ? (_1206_v38) : (_1206_v38)))) : (_module.__default.safeModuloInt(p0, (_this).f8)));
          let _rhs225 = (_1202_v34)[_module.__default.safeIndex(((_this).f3).plus(_module.__default.fm0(globalState)), new BigNumber((_1202_v34).length))];
          let _rhs226 = _this.f4;
          let _rhs227 = (_1198_v31)[_module.__default.safeIndex(new BigNumber(5), new BigNumber((_1198_v31).length))];
          let _lhs201 = _1198_v31;
          let _lhs202 = _module.__default.safeIndex(new BigNumber(5), new BigNumber((_1198_v31).length));
          let _lhs203 = globalState;
          let _lhs204 = _this;
          let _lhs205 = globalState;
          _1201_v33 = _rhs223;
          _lhs201[_lhs202] = _rhs224;
          _lhs203.f2 = _rhs225;
          _lhs204.f4 = _rhs226;
          _lhs205.f1 = _rhs227;
          _1198_v31 = _1198_v31;
        }
        let _1210_v41;
        _1210_v41 = _dafny.Map.Empty.slice().updateUnsafe((_this.f4) || (!(_this.f4)),(_this).f3);
        let _1211_v42;
        let _nw219 = new _module.C4();
        _nw219.__ctor(new BigNumber(-621), (_this).f7, _this.f4);
        _1211_v42 = _nw219;
        let _1212_v43;
        _1212_v43 = _dafny.Seq.of(_1211_v42, _1211_v42);
        let _1213_v44;
        _1213_v44 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f4),_1212_v43);
        let _1214_v45;
        _1214_v45 = (((_1213_v44).contains(true)) ? ((_1213_v44).get(true)) : (_1212_v43));
        _1210_v41 = (_1210_v41).update(_this.f4, new BigNumber((_dafny.Seq.Concat((_1214_v45), _1212_v43)).length));
        let _1215_v46;
        _1215_v46 = new _dafny.CodePoint('f'.codePointAt(0));
        let _1216_v47;
        _1216_v47 = _dafny.Seq.of(_this.f4);
        let _1217_v48;
        let _nw220 = new _module.C3();
        _nw220.__ctor(_1215_v46, _dafny.Seq.UnicodeFromString("iqrkkklva"), (_this).f8, _this.f4, new BigNumber((_1216_v47).length));
        _1217_v48 = _nw220;
        let _1218_v49;
        _1218_v49 = _dafny.Map.Empty.slice().updateUnsafe(_1217_v48,_this.f4);
        let _1219_v50;
        _1219_v50 = _dafny.Map.Empty.slice().updateUnsafe(_this.f4,_1218_v49);
        _1218_v49 = ((_1218_v49).Merge(_1218_v49)).Merge(((((_1219_v50).contains(_this.f4)) ? ((_1219_v50).get(_this.f4)) : ((_dafny.Map.Empty.slice().updateUnsafe(_1217_v48,_this.f4)).update(_1217_v48, _this.f4)))).Merge(_1218_v49));
        let _1220_v51;
        let _init37 = function (_1221_i5) {
          return (_1221_i5).minus((_this).f8);
        };
        let _nw221 = Array((new BigNumber(15)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw221.length); _i0_37++) {
          _nw221[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1220_v51 = _nw221;
        let _1222_v52;
        _1222_v52 = _dafny.MultiSet.fromElements(_1220_v51, _1220_v51);
        let _index209 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_1220_v51).length));
        (_1220_v51)[_index209] = new BigNumber((_1222_v52).cardinality());
        let _index210 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_1220_v51).length));
        (_1220_v51)[_index210] = _module.__default.safeDivisionInt(new BigNumber(46), (_this).f7);
        let _1223_v53;
        _1223_v53 = _dafny.Set.fromElements(_this.f4);
        let _1224_v54;
        _1224_v54 = _dafny.Map.Empty.slice().updateUnsafe((_this).f8,_1223_v53);
        let _index211 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_1220_v51).length));
        (_1220_v51)[_index211] = new BigNumber(((((((_1224_v54).contains(p0)) ? ((_1224_v54).get(p0)) : (_1223_v53))).Difference(_1223_v53)).Difference(_1223_v53)).length);
      } else {
        let _1225_v55;
        _1225_v55 = _dafny.Seq.of((_1176_v14).f9, (_1176_v14).f9, (_1176_v14).f9);
        _1225_v55 = _dafny.Seq.Concat(_dafny.Seq.update(_1225_v55, _module.__default.safeIndex((_this).f7, new BigNumber((_1225_v55).length)), (_1176_v14).f9), _dafny.Seq.Concat(_dafny.Seq.update(_1225_v55, _module.__default.safeIndex(p0, new BigNumber((_1225_v55).length)), _1174_v13), _1225_v55));
        let _1226_v56;
        _1226_v56 = _dafny.Map.Empty.slice().updateUnsafe(!(_this.f4),_this.f4);
        _1226_v56 = _1226_v56;
        let _1227_v57;
        _1227_v57 = _dafny.Seq.of(_this.f4, _this.f4);
        let _1228_v58;
        _1228_v58 = _dafny.Set.fromElements(_dafny.Seq.Concat(_dafny.Seq.update(_1227_v57, _module.__default.safeIndex((_this).f8, new BigNumber((_1227_v57).length)), _this.f4), _1227_v57), _1227_v57, _1227_v57, _1227_v57);
        let _1229_v59;
        _1229_v59 = _dafny.Seq.of(_dafny.Set.fromElements(_1227_v57), (_1228_v58).Union(_dafny.Set.fromElements(_1227_v57)));
        _1228_v58 = (_1229_v59)[_module.__default.safeIndex(p0, new BigNumber((_1229_v59).length))];
        let _1230_v60;
        _1230_v60 = _dafny.MultiSet.fromElements(_this.f4);
        let _1231_v61;
        _1231_v61 = _dafny.Set.fromElements(new BigNumber((_1230_v60).cardinality()), new BigNumber(818));
        _1177_v15 = _module.__default.fm15((new BigNumber((_1231_v61).length)).isLessThan((_this).f8), (_this).f8, (_this).f8, globalState);
        r0 = (_dafny.ZERO).minus((_dafny.ZERO).minus((((_this.f4) || (_this.f4)) ? ((_this).f7) : ((_dafny.ZERO).minus(new BigNumber(-893))))));
      }
      (globalState).f2 = _this.f4;
      let _1232_v62;
      _1232_v62 = _module.D4.create_DC12();
      let _1233_v63;
      _1233_v63 = _dafny.Set.fromElements(_this.f4);
      let _1234_v64;
      _1234_v64 = _dafny.Seq.of(_this.f4);
      let _1235_v65;
      _1235_v65 = _dafny.Seq.of(p0);
      let _1236_v66;
      _1236_v66 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1235_v65).length),(_this).f8);
      let _1237_v67;
      _1237_v67 = _module.D3.create_DC10((_this).f7, (_1176_v14).f9, (_dafny.ZERO).minus((_this).f3));
      let _1238_v68;
      let _nw222 = Array((new BigNumber(24)).toNumber());
      _nw222[(_dafny.ZERO).toNumber()] = (_this).f3;
      _nw222[(_dafny.ONE).toNumber()] = p0;
      _nw222[(new BigNumber(2)).toNumber()] = (_this).f8;
      _nw222[(new BigNumber(3)).toNumber()] = new BigNumber(688);
      _nw222[(new BigNumber(4)).toNumber()] = new BigNumber((_1177_v15).length);
      _nw222[(new BigNumber(5)).toNumber()] = new BigNumber((_1233_v63).length);
      _nw222[(new BigNumber(6)).toNumber()] = p0;
      _nw222[(new BigNumber(7)).toNumber()] = (_this).f8;
      _nw222[(new BigNumber(8)).toNumber()] = (_this).f7;
      _nw222[(new BigNumber(9)).toNumber()] = (_this).f8;
      _nw222[(new BigNumber(10)).toNumber()] = new BigNumber((_1234_v64).length);
      _nw222[(new BigNumber(11)).toNumber()] = _module.__default.fm0(globalState);
      _nw222[(new BigNumber(12)).toNumber()] = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(831)), function (_1239_i6) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length);
      _nw222[(new BigNumber(13)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_1235_v65)).cardinality());
      _nw222[(new BigNumber(14)).toNumber()] = new BigNumber(((_1236_v66).update(p0, p0)).length);
      _nw222[(new BigNumber(15)).toNumber()] = p0;
      _nw222[(new BigNumber(16)).toNumber()] = (_1237_v67).dtor_cf17;
      _nw222[(new BigNumber(17)).toNumber()] = new BigNumber((_module.__default.fm31(globalState)).length);
      _nw222[(new BigNumber(18)).toNumber()] = (_this).f3;
      _nw222[(new BigNumber(19)).toNumber()] = p0;
      _nw222[(new BigNumber(20)).toNumber()] = (_this).f8;
      _nw222[(new BigNumber(21)).toNumber()] = p0;
      _nw222[(new BigNumber(22)).toNumber()] = (_this).f8;
      _nw222[(new BigNumber(23)).toNumber()] = (_this).f3;
      _1238_v68 = _nw222;
      let _1240_v69;
      _1240_v69 = _dafny.Seq.of(_1238_v68);
      let _1241_v70;
      let _nw223 = new _module.C2();
      _nw223.__ctor(p0, _1238_v68, (_this).f7, _this.f4);
      _1241_v70 = _nw223;
      let _1242_v71;
      _1242_v71 = _dafny.Seq.of(_1241_v70, _1241_v70, _1241_v70);
      let _1243_v72;
      _1243_v72 = _dafny.Seq.of(_1240_v69);
      let _1244_v73;
      _1244_v73 = new _dafny.CodePoint('b'.codePointAt(0));
      let _1245_v74;
      _1245_v74 = _module.D8.create_DC23(_this.f4, _this.f4, _this.f4, _1244_v73);
      let _1246_v75;
      _1246_v75 = _module.D0.create_DC1(_this.f4, _this.f4, _1244_v73, (_1241_v70).f15);
      let _pat_let_tv24 = _1235_v65;
      let _pat_let_tv25 = _1235_v65;
      let _pat_let_tv26 = _1235_v65;
      let _pat_let_tv27 = _1241_v70;
      let _pat_let_tv28 = _1177_v15;
      let _pat_let_tv29 = _1241_v70;
      let _pat_let_tv30 = _1244_v73;
      let _pat_let_tv31 = _1244_v73;
      let _rhs228 = ((_this.f4) ? (_1232_v62) : (_module.D4.create_DC12()));
      let _rhs229 = (_1243_v72)[_module.__default.safeIndex((_1241_v70).f14, new BigNumber((_1243_v72).length))];
      let _rhs230 = function (_source15) {
        if (_source15.is_DC23) {
          let _1247___mcc_h2 = (_source15).cf32;
          let _1248___mcc_h3 = (_source15).cf33;
          let _1249___mcc_h4 = (_source15).cf34;
          let _1250___mcc_h5 = (_source15).cf35;
          let _1251_cf35 = _1250___mcc_h5;
          let _1252_cf34 = _1249___mcc_h4;
          let _1253_cf33 = _1248___mcc_h3;
          let _1254_cf32 = _1247___mcc_h2;
          return _dafny.Seq.update(_pat_let_tv24, _module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements((_module.D0.create_DC0(_1252_cf34)).dtor_cf0)).cardinality()), new BigNumber((_pat_let_tv25).length)), (_this).f8);
        } else if (_source15.is_DC22) {
          let _1255___mcc_h6 = (_source15).cf31;
          let _1256_cf31 = _1255___mcc_h6;
          return _dafny.Seq.Concat(_pat_let_tv26, _dafny.Seq.of((_dafny.ZERO).minus((_pat_let_tv27).f14), new BigNumber((_dafny.MultiSet.fromElements(_this.f4)).cardinality())));
        } else {
          let _1257___mcc_h7 = (_source15).cf36;
          let _1258_cf36 = _1257___mcc_h7;
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(724)), ((_1259_v15, _1260_v70, _1261_v73) => function (_1262_i7) {
            return new BigNumber((_dafny.Seq.update(_1259_v15, _module.__default.safeIndex((_1260_v70).f14, new BigNumber((_1259_v15).length)), _1261_v73)).length);
          })(_pat_let_tv28, _pat_let_tv29, _pat_let_tv30));
        }
      }(_1245_v74);
      let _rhs231 = (function (_pat_let32_0) {
        return function (_1263_dt__update__tmp_h0) {
          return function (_pat_let33_0) {
            return function (_1264_dt__update_hcf3_h0) {
              return _module.D0.create_DC1((_1263_dt__update__tmp_h0).dtor_cf1, (_1263_dt__update__tmp_h0).dtor_cf2, _1264_dt__update_hcf3_h0, (_1263_dt__update__tmp_h0).dtor_cf4);
            }(_pat_let33_0);
          }(_pat_let_tv31);
        }(_pat_let32_0);
      }(_1246_v75)).dtor_cf2;
      let _rhs232 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1242_v71, _1242_v71), _1242_v71);
      let _lhs206 = globalState;
      _1232_v62 = _rhs228;
      _1240_v69 = _rhs229;
      _1235_v65 = _rhs230;
      _lhs206.f2 = _rhs231;
      _1242_v71 = _rhs232;
      let _1265_v76;
      _1265_v76 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_1234_v64).length));
      let _1266_v77;
      _1266_v77 = _dafny.MultiSet.fromElements(new BigNumber((_1235_v65).length), (((_1265_v76).contains(_this.f4)) ? ((_1265_v76).get(_this.f4)) : ((_this).f3)), (_this).f7, (_this).f8);
      let _1267_v78;
      let _nw224 = new _module.C3();
      _nw224.__ctor(_1244_v73, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-483)), ((_1268_v73) => function (_1269_i8) {
        return _1268_v73;
      })(_1244_v73)), new BigNumber((_1177_v15).length), _this.f4, new BigNumber((((_1266_v77).update(p0, _module.__default.abs(p0))).update((_this).f8, _module.__default.abs(new BigNumber(-529)))).cardinality()));
      _1267_v78 = _nw224;
      r0 = (_1241_v70).f14;
      return r0;
    }
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C7 = class C7 {
    constructor () {
      this._tname = "_module.C7";
      this._f4 = false;
      this._f3 = _dafny.ZERO;
      this.f5 = false;
      this._f6 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f4() {
      let _this = this;
      return _this._f4;
    };
    set f4(value) {
      let _this = this;
      _this._f4 = value;
    };
    get f3() {
      let _this = this;
      return _this._f3;
    };
    __ctor(f5, f6, f3, f4) {
      let _this = this;
      (_this).f5 = f5;
      (_this)._f6 = f6;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      return new _dafny.CodePoint('f'.codePointAt(0));
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.Seq.of();
      let r1 = _dafny.ZERO;
      (_this).f5 = _this.f5;
      let _1270_v0;
      _1270_v0 = _module.D1.create_DC5((_this).f3);
      let _hi12 = ((_this).f3).plus((_1270_v0).dtor_cf8);
      for (let _1271_i0 = (_this).f6; _1271_i0.isLessThan(_hi12); _1271_i0 = _1271_i0.plus(_dafny.ONE)) {
        let _1272_v1;
        _1272_v1 = _dafny.Seq.of(_this.f4);
        let _1273_v2;
        _1273_v2 = _dafny.Seq.of(!((_1272_v1)[_module.__default.safeIndex(new BigNumber(242), new BigNumber((_1272_v1).length))]), _this.f5, _this.f4, _this.f4, _this.f4);
        let _1274_v3;
        let _init38 = function (_1275_i1) {
          return _this.f4;
        };
        let _nw225 = Array((new BigNumber(4)).toNumber());
        for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw225.length); _i0_38++) {
          _nw225[_i0_38] = _init38(new BigNumber(_i0_38));
        }
        _1274_v3 = _nw225;
        let _1276_v4;
        _1276_v4 = _dafny.Map.Empty.slice().updateUnsafe(_1274_v3,_this.f4);
        let _1277_v5;
        _1277_v5 = _dafny.Seq.of((_this).f6);
        let _1278_v6;
        _1278_v6 = _module.D0.create_DC0(_this.f4);
        let _1279_v7;
        _1279_v7 = _dafny.Seq.UnicodeFromString("wyblf");
        let _1280_v8;
        _1280_v8 = _dafny.MultiSet.fromElements(_this.f5);
        let _1281_v9;
        let _nw226 = Array((new BigNumber(24)).toNumber());
        _nw226[(_dafny.ZERO).toNumber()] = _this.f5;
        _nw226[(_dafny.ONE).toNumber()] = (_1272_v1)[_module.__default.safeIndex((_this).f6, new BigNumber((_1272_v1).length))];
        _nw226[(new BigNumber(2)).toNumber()] = _this.f4;
        _nw226[(new BigNumber(3)).toNumber()] = _this.f5;
        _nw226[(new BigNumber(4)).toNumber()] = _this.f4;
        _nw226[(new BigNumber(5)).toNumber()] = _dafny.Seq.IsProperPrefixOf(_module.__default.fm4(_this.f5, (((_1276_v4).contains(_1274_v3)) ? ((_1276_v4).get(_1274_v3)) : (_this.f5)), _this.f5, _this.f4, globalState), _1277_v5);
        _nw226[(new BigNumber(6)).toNumber()] = _this.f5;
        _nw226[(new BigNumber(7)).toNumber()] = !((_this).f3).isEqualTo(_1271_i0);
        _nw226[(new BigNumber(8)).toNumber()] = _this.f4;
        _nw226[(new BigNumber(9)).toNumber()] = (_1278_v6).dtor_cf0;
        _nw226[(new BigNumber(10)).toNumber()] = _this.f4;
        _nw226[(new BigNumber(11)).toNumber()] = !(_this.f4);
        _nw226[(new BigNumber(12)).toNumber()] = _this.f5;
        _nw226[(new BigNumber(13)).toNumber()] = _this.f5;
        _nw226[(new BigNumber(14)).toNumber()] = _this.f4;
        _nw226[(new BigNumber(15)).toNumber()] = ((_this).f6).isLessThan(new BigNumber((_1279_v7).length));
        _nw226[(new BigNumber(16)).toNumber()] = _this.f5;
        _nw226[(new BigNumber(17)).toNumber()] = _this.f5;
        _nw226[(new BigNumber(18)).toNumber()] = (_dafny.MultiSet.fromElements(_this.f5, _this.f4)).IsDisjointFrom(_1280_v8);
        _nw226[(new BigNumber(19)).toNumber()] = _this.f4;
        _nw226[(new BigNumber(20)).toNumber()] = !(_this.f5) || (_this.f5);
        _nw226[(new BigNumber(21)).toNumber()] = _this.f4;
        _nw226[(new BigNumber(22)).toNumber()] = _this.f4;
        _nw226[(new BigNumber(23)).toNumber()] = !(!(!(_this.f4)));
        _1281_v9 = _nw226;
        let _1282_v10;
        _1282_v10 = _dafny.Seq.of(_1281_v9, _1274_v3, _1281_v9);
        _1281_v9 = (_1282_v10)[_module.__default.safeIndex((_this).f3, new BigNumber((_1282_v10).length))];
        let _1283_v11;
        _1283_v11 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt((_dafny.ZERO).minus(_1271_i0), (_this).f6),_dafny.Seq.UnicodeFromString("u"));
        _1283_v11 = (_1283_v11).update((_this).f3, _1279_v7);
        (globalState).f1 = (_this).f6;
        let _1284_v12;
        let _nw227 = new _module.C6();
        _nw227.__ctor(_module.__default.safeDivisionInt((_this).f3, (_dafny.ZERO).minus((_this).f3)), (_this).f3, true, (_this).f6);
        _1284_v12 = _nw227;
      }
      let _1285_v13;
      let _init39 = function (_1286_i3) {
        return (_1286_i3).plus((_this).f3);
      };
      let _nw228 = Array((new BigNumber(27)).toNumber());
      for (let _i0_39 = 0; _i0_39 < new BigNumber(_nw228.length); _i0_39++) {
        _nw228[_i0_39] = _init39(new BigNumber(_i0_39));
      }
      _1285_v13 = _nw228;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1285_v13).length))) {
        let _1287_i2 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1287_i2)) && ((_1287_i2).isLessThan(new BigNumber((_1285_v13).length))))) {
          (_1285_v13)[(_1287_i2)] = _module.__default.safeModuloInt(_1287_i2, new BigNumber(947));
        }
      }
      let _1288_v14;
      _1288_v14 = _dafny.Seq.of(_this.f4);
      let _index212 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_1285_v13).length));
      (_1285_v13)[_index212] = new BigNumber((_1288_v14).length);
      let _index213 = _module.__default.safeIndex(new BigNumber(527), new BigNumber((_1285_v13).length));
      (_1285_v13)[_index213] = _module.__default.fm0(globalState);
      (_this).f4 = _this.f4;
      (_this).f4 = (_this.f4) && (_this.f4);
      r0 = _1288_v14;
      r1 = (_this).f3;
      return [r0, r1];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
