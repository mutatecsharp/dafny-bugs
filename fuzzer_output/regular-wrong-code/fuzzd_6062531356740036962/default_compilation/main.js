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
    static fm0(p0, p1, p2, p3, globalState) {
      return (!(true)) || ((!(true)) || (false));
    };
    static fm5(p0, p1, globalState) {
      return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber(299), new BigNumber(553)), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(162)), function (_0_i0) {
        return new _dafny.CodePoint('r'.codePointAt(0));
      })).length)), function () {
        let _coll0 = new _dafny.Set();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(955), new BigNumber(500))) {
          let _1_v0 = _compr_0;
          if (((new BigNumber(955)).isLessThanOrEqualTo(_1_v0)) && ((_1_v0).isLessThan(new BigNumber(500)))) {
            _coll0.add((_1_v0).plus(new BigNumber((_dafny.Set.fromElements(true, !(true))).length)));
          }
        }
        return _coll0;
      }(), function () {
        let _coll1 = new _dafny.Set();
        for (const _compr_1 of (_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length))).Elements) {
          let _2_v1 = _compr_1;
          if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,true)).length)), _2_v1)) {
            _coll1.add((_2_v1).minus(new BigNumber(601)));
          }
        }
        return _coll1;
      }())).length), new BigNumber(528)), _dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(429)), function (_3_i1) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      })).length), new BigNumber((_dafny.Seq.UnicodeFromString("smswnlmsn")).length)))).length);
    };
    static fm7(globalState) {
      if (false) {
        return _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(true));
      } else {
        return _dafny.Seq.of(true);
      }
    };
    static fm8(p0, p1, p2, globalState) {
      return ((_dafny.MultiSet.fromElements(!(true))).Intersect(_dafny.MultiSet.fromElements(false))).Intersect((_dafny.MultiSet.fromElements(false)).Union(_dafny.MultiSet.fromElements(!(true), false)));
    };
    static fm9(p0, p1, globalState) {
      return ((function () {
        let _coll2 = new _dafny.Set();
        for (const _compr_2 of (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(713),true))).Elements) {
          let _4_v0 = _compr_2;
          if ((_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(713),true))).contains(_4_v0)) {
            _coll2.add(_4_v0);
          }
        }
        return _coll2;
      }()).Intersect(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()),false)))).Difference(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(413),true), _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("k")).length)),false), function () {
        let _coll3 = new _dafny.Map();
        for (const _compr_3 of _dafny.IntegerRange(new BigNumber(-871), new BigNumber(-831))) {
          let _5_v1 = _compr_3;
          if (((new BigNumber(-871)).isLessThanOrEqualTo(_5_v1)) && ((_5_v1).isLessThan(new BigNumber(-831)))) {
            _coll3.push([(_5_v1).plus(new BigNumber(-737)),false]);
          }
        }
        return _coll3;
      }()));
    };
    static fm10(p0, p1, p2, p3, globalState) {
      return new _dafny.CodePoint('n'.codePointAt(0));
    };
    static fm11(p0, p1, globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(255),new BigNumber((((!(!(true))) ? (_dafny.Set.fromElements(_module.D12.create_DC26(_dafny.Set.fromElements(new BigNumber(387))), _module.D12.create_DC26(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(569),true)).length))), _module.D12.create_DC26(_dafny.Set.fromElements(new BigNumber(427), new BigNumber(687), new BigNumber(659), new BigNumber(249))), _module.D12.create_DC26(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(true))).length),false)).length))))) : (_dafny.Set.fromElements(_module.D12.create_DC26(function () {
  let _coll4 = new _dafny.Set();
  for (const _compr_4 of _dafny.IntegerRange(new BigNumber(71), new BigNumber(108))) {
    let _6_v0 = _compr_4;
    if (((new BigNumber(71)).isLessThanOrEqualTo(_6_v0)) && ((_6_v0).isLessThan(new BigNumber(108)))) {
      _coll4.add((_6_v0).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(!(true)))).length)));
    }
  }
  return _coll4;
}()), _module.D12.create_DC26(_dafny.Set.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("cao")).length), new BigNumber(353))), _module.D12.create_DC26(function () {
  let _coll5 = new _dafny.Set();
  for (const _compr_5 of _dafny.IntegerRange(new BigNumber(388), new BigNumber(551))) {
    let _7_v1 = _compr_5;
    if (((new BigNumber(388)).isLessThanOrEqualTo(_7_v1)) && ((_7_v1).isLessThan(new BigNumber(551)))) {
      _coll5.add(_module.__default.safeDivisionInt(_7_v1, new BigNumber((function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0)))).Elements) {
          let _8_v2 = _compr_6;
          if (_dafny.Seq.contains(_dafny.Seq.of(new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('e'.codePointAt(0))), _8_v2)) {
            _coll6.push([_8_v2,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('h'.codePointAt(0)))).length)]);
          }
        }
        return _coll6;
      }()).length)));
    }
  }
  return _coll5;
}()), _module.D12.create_DC26(_dafny.Set.fromElements(new BigNumber(129))))))).length));
    };
    static fm12(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(true)).Intersect(_dafny.Set.fromElements(true))).Union((_dafny.Set.fromElements(false, true, !(true))).Difference(_dafny.Set.fromElements(!(!(true)))));
    };
    static fm13(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,_module.D3.create_DC6(new BigNumber((_dafny.MultiSet.fromElements(!(true))).cardinality()), new BigNumber(-906), false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D3.create_DC6(new BigNumber(799), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-306)), function (_9_i0) {
  return new BigNumber((_dafny.Seq.UnicodeFromString("wmk")).length);
})).length), true)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D3.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("qi")).length), new BigNumber((function () {
  let _coll7 = new _dafny.Map();
  for (const _compr_7 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("tmyf"),!(false))).Keys.Elements) {
    let _10_v0 = _compr_7;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("tmyf"),!(false))).contains(_10_v0)) {
      _coll7.push([_10_v0,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("ecjrmps")).length))]);
    }
  }
  return _coll7;
}()).length), true)));
    };
    static fm14(p0, globalState) {
      let _source0 = _module.D6.create_DC12(true, new BigNumber(932));
      if (_source0.is_DC12) {
        let _11___mcc_h0 = (_source0).cf17;
        let _12___mcc_h1 = (_source0).cf18;
        let _13_cf18 = _12___mcc_h1;
        let _14_cf17 = _11___mcc_h0;
        return _module.D10.create_DC18(_dafny.Set.fromElements(_14_cf17));
      } else {
        let _15___mcc_h2 = (_source0).cf16;
        let _16_cf16 = _15___mcc_h2;
        return _module.D10.create_DC18(_dafny.Set.fromElements(!(true)));
      }
    };
    static fm15(p0, p1, p2, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(10)), function (_17_i0) {
        return new _dafny.CodePoint('e'.codePointAt(0));
      }), _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-385)), function (_18_i1) {
        return new _dafny.CodePoint('n'.codePointAt(0));
      }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-51)), function (_19_i2) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })));
    };
    static fm16(p0, globalState) {
      return _module.D10.create_DC20((_dafny.Set.fromElements(new BigNumber(266))).IsDisjointFrom(_dafny.Set.fromElements(new BigNumber(246))), (new BigNumber(-107)).plus(new BigNumber(449)));
    };
    static fm17(p0, p1, p2, globalState) {
      let _source1 = ((true) ? (_module.D2.create_DC2(false)) : (_module.D2.create_DC2(true)));
      if (_source1.is_DC3) {
        let _20___mcc_h0 = (_source1).cf3;
        let _21___mcc_h1 = (_source1).cf4;
        let _22_cf4 = _21___mcc_h1;
        let _23_cf3 = _20___mcc_h0;
        return _23_cf3;
      } else {
        let _24___mcc_h2 = (_source1).cf2;
        let _25_cf2 = _24___mcc_h2;
        return (_dafny.ZERO).minus((_module.D11.create_DC24(_25_cf2, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_25_cf2,new BigNumber(205))).length))).dtor_cf37);
      }
    };
    static fm19(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,!(!(true)))).Merge(((true) ? (_dafny.Map.Empty.slice().updateUnsafe(true,true)) : (_dafny.Map.Empty.slice().updateUnsafe((_module.D11.create_DC23(new BigNumber(355), !(true))).dtor_cf35,!(false)))));
    };
    static fm20(p0, p1, globalState) {
      let _source2 = ((!(true)) ? (_module.D11.create_DC24(true, new BigNumber(961))) : (_module.D11.create_DC24(true, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new _dafny.CodePoint('a'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0)), new _dafny.CodePoint('a'.codePointAt(0)))).cardinality()),new BigNumber(634))).length))));
      if (_source2.is_DC23) {
        let _26___mcc_h0 = (_source2).cf34;
        let _27___mcc_h1 = (_source2).cf35;
        let _28_cf35 = _27___mcc_h1;
        let _29_cf34 = _26___mcc_h0;
        return _module.D10.create_DC19(new _dafny.CodePoint('r'.codePointAt(0)));
      } else if (_source2.is_DC24) {
        let _30___mcc_h2 = (_source2).cf36;
        let _31___mcc_h3 = (_source2).cf37;
        let _32_cf37 = _31___mcc_h3;
        let _33_cf36 = _30___mcc_h2;
        return _module.D10.create_DC19(new _dafny.CodePoint('x'.codePointAt(0)));
      } else if (_source2.is_DC22) {
        let _34___mcc_h4 = (_source2).cf33;
        let _35_cf33 = _34___mcc_h4;
        return _module.D10.create_DC19(new _dafny.CodePoint('k'.codePointAt(0)));
      } else {
        let _36___mcc_h5 = (_source2).cf38;
        let _37_cf38 = _36___mcc_h5;
        return _module.D10.create_DC19(new _dafny.CodePoint('x'.codePointAt(0)));
      }
    };
    static fm21(p0, p1, globalState) {
      let _source3 = ((true) ? (_module.D9.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(true,_module.D3.create_DC6(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(46),new BigNumber(327))).length), new BigNumber(723), !(false))))) : (_module.D9.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D3.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("nmqjwa")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(210)), function (_38_i0) {
  return new _dafny.CodePoint('w'.codePointAt(0));
})).length), false)))));
      if (_source3.is_DC17) {
        let _39___mcc_h0 = (_source3).cf25;
        let _40_cf25 = _39___mcc_h0;
        return _dafny.Seq.Concat(_dafny.Seq.of(_module.D7.create_DC13(function () {
  let _coll8 = new _dafny.Map();
  for (const _compr_8 of (_dafny.Seq.of(_40_cf25)).Elements) {
    let _41_v0 = _compr_8;
    if (_dafny.Seq.contains(_dafny.Seq.of(_40_cf25), _41_v0)) {
      _coll8.push([(_41_v0).minus(_40_cf25),false]);
    }
  }
  return _coll8;
}())), _dafny.Seq.of(_module.D7.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(_40_cf25,!(false)))));
      } else {
        let _42___mcc_h1 = (_source3).cf24;
        let _43_cf24 = _42___mcc_h1;
        return _dafny.Seq.of(_module.D7.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-943)), function (_44_i1) {
  return new _dafny.CodePoint('i'.codePointAt(0));
})).length),false)), _module.D7.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(853),true)));
      }
    };
    static fm22(p0, p1, p2, globalState) {
      return _module.D11.create_DC23((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(true, !(false), true),new BigNumber(132))).length)), !(new BigNumber((_dafny.Set.fromElements(new BigNumber(249))).length)).isEqualTo(new BigNumber((_dafny.Seq.UnicodeFromString("b")).length)));
    };
    static fm23(globalState) {
      return _module.D3.create_DC6((new BigNumber(8)).plus(new BigNumber((_dafny.Seq.of(true)).length)), new BigNumber((((false) ? (_dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)))) : (_dafny.Seq.of(new _dafny.CodePoint('l'.codePointAt(0)))))).length), false);
    };
    static fm24(p0, p1, p2, globalState) {
      if (!(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(290), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(977)), function (_45_i0) {
        return new _dafny.CodePoint('x'.codePointAt(0));
      })).length), new BigNumber(-498), new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(752))).cardinality()))).cardinality()))).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(618)))) {
        return _module.D9.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D3.create_DC6(new BigNumber((_dafny.MultiSet.fromElements(false)).cardinality()), new BigNumber(630), true)));
      } else {
        return _module.D9.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(false,_module.D3.create_DC6(new BigNumber((_dafny.Seq.UnicodeFromString("cnempwcsy")).length), new BigNumber((_dafny.Set.fromElements(true)).length), false)));
      }
    };
    static fm25(p0, p1, p2, globalState) {
      return ((function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of (function () {
          let _coll10 = new _dafny.Set();
          for (const _compr_10 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(633))).cardinality()))).Keys.Elements) {
            let _46_v1 = _compr_10;
            if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(633))).cardinality()))).contains(_46_v1)) {
              _coll10.add(_46_v1);
            }
          }
          return _coll10;
        }()).Elements) {
          let _47_v0 = _compr_9;
          if ((function () {
            let _coll11 = new _dafny.Set();
            for (const _compr_11 of (_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(633))).cardinality()))).Keys.Elements) {
              let _48_v1 = _compr_11;
              if ((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('i'.codePointAt(0)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(633))).cardinality()))).contains(_48_v1)) {
                _coll11.add(_48_v1);
              }
            }
            return _coll11;
          }()).contains(_47_v0)) {
            _coll9.push([_47_v0,false]);
          }
        }
        return _coll9;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('h'.codePointAt(0)),true))).Merge((_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('p'.codePointAt(0)),false)));
    };
    static fm26(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(639),false)).length)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(680)), function (_49_i0) {
        return new BigNumber((_dafny.Set.fromElements(!(true))).length);
      }));
    };
    static fm27(p0, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements(false)).length)))).length)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false, false, false, false),new BigNumber(-332)));
    };
    static fm28(p0, p1, p2, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(_module.D9.create_DC17(new BigNumber((_dafny.Set.fromElements(false, true)).length)),false)).Merge((_dafny.Map.Empty.slice().updateUnsafe(_module.D9.create_DC17(new BigNumber(288)),!(true))).Merge(function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Set.fromElements(_module.D9.create_DC17(new BigNumber(859)), _module.D9.create_DC17(new BigNumber(413)), _module.D9.create_DC17(new BigNumber((function () {
  let _coll13 = new _dafny.Set();
  for (const _compr_13 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(152)), function (_50_i0) {
    return _dafny.Seq.UnicodeFromString("ujowcoxtg");
  })).Elements) {
    let _51_v1 = _compr_13;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(152)), function (_50_i0) {
      return _dafny.Seq.UnicodeFromString("ujowcoxtg");
    }), _51_v1)) {
      _coll13.add(_51_v1);
    }
  }
  return _coll13;
}()).length)))).Elements) {
          let _52_v0 = _compr_12;
          if ((_dafny.Set.fromElements(_module.D9.create_DC17(new BigNumber(859)), _module.D9.create_DC17(new BigNumber(413)), _module.D9.create_DC17(new BigNumber((function () {
  let _coll14 = new _dafny.Set();
  for (const _compr_14 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(152)), function (_50_i0) {
    return _dafny.Seq.UnicodeFromString("ujowcoxtg");
  })).Elements) {
    let _53_v1 = _compr_14;
    if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(152)), function (_50_i0) {
      return _dafny.Seq.UnicodeFromString("ujowcoxtg");
    }), _53_v1)) {
      _coll14.add(_53_v1);
    }
  }
  return _coll14;
}()).length)))).contains(_52_v0)) {
            _coll12.push([_52_v0,!(true)]);
          }
        }
        return _coll12;
      }()));
    };
    static fm29(p0, globalState) {
      if (false) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(857)), function (_54_i0) {
          return new _dafny.CodePoint('a'.codePointAt(0));
        });
      } else {
        return _dafny.Seq.UnicodeFromString("qt");
      }
    };
    static fm30(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll15 = new _dafny.Map();
        for (const _compr_15 of ((function () {
          let _coll16 = new _dafny.Set();
          for (const _compr_16 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), function (_55_i0) {
            return _dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.UnicodeFromString("ro"));
          })).Elements) {
            let _56_v1 = _compr_16;
            if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), function (_55_i0) {
              return _dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.UnicodeFromString("ro"));
            }), _56_v1)) {
              _coll16.add(_56_v1);
            }
          }
          return _coll16;
        }()).Intersect(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("qduxisjho"))))).Elements) {
          let _57_v0 = _compr_15;
          if (((function () {
            let _coll17 = new _dafny.Set();
            for (const _compr_17 of (_dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), function (_55_i0) {
              return _dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.UnicodeFromString("ro"));
            })).Elements) {
              let _58_v1 = _compr_17;
              if (_dafny.Seq.contains(_dafny.Seq.Create(_module.__default.abs(new BigNumber(590)), function (_55_i0) {
                return _dafny.Map.Empty.slice().updateUnsafe(!(true),_dafny.Seq.UnicodeFromString("ro"));
              }), _58_v1)) {
                _coll17.add(_58_v1);
              }
            }
            return _coll17;
          }()).Intersect(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.UnicodeFromString("qduxisjho"))))).contains(_57_v0)) {
            _coll15.push([_57_v0,_dafny.areEqual(_dafny.Seq.UnicodeFromString("fyriywjbp"), _dafny.Seq.UnicodeFromString("mniamr"))]);
          }
        }
        return _coll15;
      }();
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return (_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(-432))).Merge((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(152))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,(_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(true, !(false))).cardinality())))));
    };
    static fm32(p0, p1, globalState) {
      return function () {
        let _coll18 = new _dafny.Map();
        for (const _compr_18 of _dafny.IntegerRange(new BigNumber(672), new BigNumber(958))) {
          let _59_v0 = _compr_18;
          if (((new BigNumber(672)).isLessThanOrEqualTo(_59_v0)) && ((_59_v0).isLessThan(new BigNumber(958)))) {
            _coll18.push([_module.__default.safeModuloInt(_59_v0, new BigNumber((_dafny.Seq.of(false)).length)),_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_dafny.Seq.of(!(true)), _dafny.Seq.of(true, false)))]);
          }
        }
        return _coll18;
      }();
    };
    static fm33(globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("gvtucu"),new BigNumber((function () {
        let _coll19 = new _dafny.Map();
        for (const _compr_19 of _dafny.IntegerRange(new BigNumber(370), new BigNumber(-582))) {
          let _60_v0 = _compr_19;
          if (((new BigNumber(370)).isLessThanOrEqualTo(_60_v0)) && ((_60_v0).isLessThan(new BigNumber(-582)))) {
            _coll19.push([_module.__default.safeDivisionInt(_60_v0, new BigNumber(-876)),false]);
          }
        }
        return _coll19;
      }()).length))).Merge((_module.D16.create_DC33(function () {
  let _coll20 = new _dafny.Map();
  for (const _compr_20 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("xtjcm"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(607)), function (_61_i0) {
    return new _dafny.CodePoint('g'.codePointAt(0));
  }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(662)), function (_62_i1) {
    return new _dafny.CodePoint('f'.codePointAt(0));
  }))).Elements) {
    let _63_v1 = _compr_20;
    if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("xtjcm"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(607)), function (_61_i0) {
      return new _dafny.CodePoint('g'.codePointAt(0));
    }), _dafny.Seq.Create(_module.__default.abs(new BigNumber(662)), function (_62_i1) {
      return new _dafny.CodePoint('f'.codePointAt(0));
    }))).contains(_63_v1)) {
      _coll20.push([_63_v1,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(656)), function (_64_i2) {
        return new BigNumber(473);
      })).length),true)).length)]);
    }
  }
  return _coll20;
}())).dtor_cf49)).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("o"),new BigNumber(612)));
    };
    static fm34(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(-661)), _dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("qhld")).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Set.fromElements(false, true)).length))))).Intersect(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber((function () {
        let _coll21 = new _dafny.Set();
        for (const _compr_21 of _dafny.IntegerRange(new BigNumber(918), new BigNumber(553))) {
          let _65_v0 = _compr_21;
          if (((new BigNumber(918)).isLessThanOrEqualTo(_65_v0)) && ((_65_v0).isLessThan(new BigNumber(553)))) {
            _coll21.add((_65_v0).plus(new BigNumber(-644)));
          }
        }
        return _coll21;
      }()).length))))).Union((function () {
        let _coll22 = new _dafny.Set();
        for (const _compr_22 of (function () {
          let _coll23 = new _dafny.Set();
          for (const _compr_23 of (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(601)))).Elements) {
            let _66_v1 = _compr_23;
            if ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(601)))).contains(_66_v1)) {
              _coll23.add(_66_v1);
            }
          }
          return _coll23;
        }()).Elements) {
          let _67_v2 = _compr_22;
          if ((function () {
            let _coll24 = new _dafny.Set();
            for (const _compr_24 of (_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(601)))).Elements) {
              let _68_v1 = _compr_24;
              if ((_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(601)))).contains(_68_v1)) {
                _coll24.add(_68_v1);
              }
            }
            return _coll24;
          }()).contains(_67_v2)) {
            _coll22.add(_67_v2);
          }
        }
        return _coll22;
      }()).Difference(_dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(894)), _dafny.Seq.of(new BigNumber(((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false)).length)))).length)), _dafny.Seq.of(new BigNumber(-476), new BigNumber(198)), _dafny.Seq.of(new BigNumber(439)))));
    };
    static m6(p0, p1, p2, p3, globalState) {
      let r0 = new _dafny.CodePoint('D'.codePointAt(0));
      let _hi0 = (_dafny.ZERO).minus(p0);
      for (let _69_i0 = p0; _69_i0.isLessThan(_hi0); _69_i0 = _69_i0.plus(_dafny.ONE)) {
        if (p1) {
          let _70_v0;
          let _nw0 = Array((new BigNumber(16)).toNumber()).fill(false);
          _70_v0 = _nw0;
          let _index0 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_70_v0).length));
          (_70_v0)[_index0] = p2;
          let _index1 = _module.__default.safeIndex(new BigNumber(394), new BigNumber((_70_v0).length));
          (_70_v0)[_index1] = p3;
          let _71_v1;
          _71_v1 = new _dafny.CodePoint('f'.codePointAt(0));
          let _72_v2;
          let _nw1 = new _module.C3();
          _nw1.__ctor(_module.__default.fm0(_69_i0, (_70_v0)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_70_v0).length))], p3, p0, globalState), _71_v1, p0);
          _72_v2 = _nw1;
          _72_v2 = _72_v2;
          let _73_v3;
          _73_v3 = _dafny.MultiSet.fromElements(_module.__default.fm0(p0, (_70_v0)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_70_v0).length))], _72_v2.f31, new BigNumber(-620), globalState), p3, _module.__default.fm0(p0, (_70_v0)[_module.__default.safeIndex(new BigNumber(394), new BigNumber((_70_v0).length))], _72_v2.f31, _69_i0, globalState), p1);
          let _74_v4;
          _74_v4 = _dafny.Map.Empty.slice().updateUnsafe(_73_v3,_69_i0);
          let _75_v5;
          _75_v5 = _dafny.Seq.UnicodeFromString("vah");
          let _76_v6;
          let _nw2 = new _module.C4();
          _nw2.__ctor(_74_v4, _dafny.Seq.IsProperPrefixOf(_75_v5, _75_v5), _71_v1, p0);
          _76_v6 = _nw2;
          let _77_v7;
          let _nw3 = new _module.C3();
          _nw3.__ctor(p2, _71_v1, p0);
          _77_v7 = _nw3;
          let _78_v8;
          _78_v8 = _dafny.Seq.of(_77_v7, _77_v7);
          let _79_v9;
          let _nw4 = new _module.C5();
          _nw4.__ctor((_76_v6).f30, _71_v1, (_69_i0).multipliedBy(new BigNumber((_78_v8).length)));
          _79_v9 = _nw4;
          _79_v9 = _79_v9;
        } else {
          r0 = _module.__default.fm10((_69_i0).isLessThanOrEqualTo(_69_i0), p2, p1, p3, globalState);
          let _80_v10;
          _80_v10 = _dafny.Seq.UnicodeFromString("rthyuxu");
          let _81_v11;
          _81_v11 = _dafny.MultiSet.fromElements(_module.__default.fm5(new BigNumber((_80_v10).length), _69_i0, globalState));
          let _82_v12;
          let _nw5 = Array((new BigNumber(8)).toNumber());
          _nw5[(_dafny.ZERO).toNumber()] = p2;
          _nw5[(_dafny.ONE).toNumber()] = ((p2) ? (p3) : (p3));
          _nw5[(new BigNumber(2)).toNumber()] = (_81_v11).IsSubsetOf(_81_v11);
          _nw5[(new BigNumber(3)).toNumber()] = p2;
          _nw5[(new BigNumber(4)).toNumber()] = true;
          _nw5[(new BigNumber(5)).toNumber()] = p1;
          _nw5[(new BigNumber(6)).toNumber()] = !(p0).isEqualTo((_dafny.ZERO).minus((_dafny.ZERO).minus(_69_i0)));
          _nw5[(new BigNumber(7)).toNumber()] = (p3) || (p3);
          _82_v12 = _nw5;
          let _index2 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_82_v12).length));
          (_82_v12)[_index2] = p2;
          let _index3 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_82_v12).length));
          (_82_v12)[_index3] = p3;
          let _index4 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_82_v12).length));
          (_82_v12)[_index4] = p2;
          let _83_v13;
          _83_v13 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-334),p1);
          _83_v13 = (_83_v13).update(_69_i0, (_82_v12)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_82_v12).length))]);
          let _84_v14;
          let _init0 = ((_85_v12, _86_p1, _87_p2) => function (_88_i1) {
            return _dafny.Seq.of(!((_85_v12)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_85_v12).length))]), _86_p1, _87_p2, !(_86_p1));
          })(_82_v12, p1, p2);
          let _nw6 = Array((_dafny.ONE).toNumber());
          for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw6.length); _i0_0++) {
            _nw6[_i0_0] = _init0(new BigNumber(_i0_0));
          }
          _84_v14 = _nw6;
          let _89_v15;
          _89_v15 = _dafny.Seq.of(p3);
          let _index5 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_84_v14).length));
          (_84_v14)[_index5] = _89_v15;
          let _90_v16;
          _90_v16 = _dafny.Set.fromElements(_83_v13);
          let _91_v17;
          _91_v17 = _module.D3.create_DC4(_90_v16);
          let _92_v18;
          let _nw7 = new _module.C0();
          _nw7.__ctor();
          _92_v18 = _nw7;
          let _93_v19;
          _93_v19 = new _dafny.CodePoint('w'.codePointAt(0));
          let _94_v21;
          _94_v21 = _dafny.Set.fromElements(_93_v19);
          let _95_v22;
          _95_v22 = _module.D10.create_DC21(new BigNumber(787), (_69_i0).isEqualTo(new BigNumber(296)), !(function () {
  let _coll25 = new _dafny.Set();
  for (const _compr_25 of (_dafny.Map.Empty.slice().updateUnsafe(_93_v19,new BigNumber(736))).Keys.Elements) {
    let _96_v20 = _compr_25;
    if ((_dafny.Map.Empty.slice().updateUnsafe(_93_v19,new BigNumber(736))).contains(_96_v20)) {
      _coll25.add(_96_v20);
    }
  }
  return _coll25;
}()).equals(_94_v21));
          let _index6 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_84_v14).length));
          let _index7 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_82_v12).length));
          let _rhs0 = _89_v15;
          let _rhs1 = ((p0).isLessThan(_69_i0)) || ((new BigNumber((_80_v10).length)).isLessThanOrEqualTo(p0));
          let _rhs2 = _91_v17;
          let _rhs3 = _92_v18;
          let _rhs4 = _95_v22;
          let _lhs0 = _84_v14;
          let _lhs1 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_84_v14).length));
          let _lhs2 = _82_v12;
          let _lhs3 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_82_v12).length));
          _lhs0[_lhs1] = _rhs0;
          _lhs2[_lhs3] = _rhs1;
          _91_v17 = _rhs2;
          _92_v18 = _rhs3;
          _95_v22 = _rhs4;
        }
        let _97_v23;
        _97_v23 = _dafny.Seq.of(!(p1), p3);
        let _98_v24;
        let _nw8 = Array((new BigNumber(17)).toNumber());
        _nw8[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_97_v23, _module.__default.safeIndex(p0, new BigNumber((_97_v23).length)), p1), _97_v23);
        _nw8[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_97_v23, _97_v23);
        _nw8[(new BigNumber(2)).toNumber()] = _97_v23;
        _nw8[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_97_v23, _97_v23);
        _nw8[(new BigNumber(4)).toNumber()] = ((p2) ? (_97_v23) : (_97_v23));
        _nw8[(new BigNumber(5)).toNumber()] = _97_v23;
        _nw8[(new BigNumber(6)).toNumber()] = _97_v23;
        _nw8[(new BigNumber(7)).toNumber()] = _97_v23;
        _nw8[(new BigNumber(8)).toNumber()] = _97_v23;
        _nw8[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_97_v23, _module.__default.safeIndex(p0, new BigNumber((_97_v23).length)), !(p1)), _97_v23);
        _nw8[(new BigNumber(10)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.of(p3), _97_v23), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(p3), _97_v23)).length)), p1);
        _nw8[(new BigNumber(11)).toNumber()] = _dafny.Seq.update(_97_v23, _module.__default.safeIndex(_69_i0, new BigNumber((_97_v23).length)), p1);
        _nw8[(new BigNumber(12)).toNumber()] = _97_v23;
        _nw8[(new BigNumber(13)).toNumber()] = _97_v23;
        _nw8[(new BigNumber(14)).toNumber()] = _97_v23;
        _nw8[(new BigNumber(15)).toNumber()] = _97_v23;
        _nw8[(new BigNumber(16)).toNumber()] = _97_v23;
        _98_v24 = _nw8;
        let _index8 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_98_v24).length));
        (_98_v24)[_index8] = _97_v23;
        let _index9 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_98_v24).length));
        (_98_v24)[_index9] = _module.__default.fm7(globalState);
        let _99_v25;
        _99_v25 = _module.D11.create_DC23(p0, p1);
        let _100_v26;
        _100_v26 = _module.D11.create_DC25(_99_v25);
        let _101_v27;
        _101_v27 = _dafny.Map.Empty.slice().updateUnsafe(_100_v26,p1);
        let _rhs5 = (_dafny.ZERO).minus(_module.__default.fm17((_101_v27).equals(_101_v27), !(p1) || (p1), p3, globalState));
        let _rhs6 = p3;
        let _lhs4 = globalState;
        let _lhs5 = globalState;
        _lhs4.f2 = _rhs5;
        _lhs5.f18 = _rhs6;
        let _102_v28;
        _102_v28 = _dafny.Seq.UnicodeFromString("rkvodkl");
        if (!_dafny.areEqual(_102_v28, _102_v28)) {
          let _103_v29;
          let _nw9 = new _module.C2();
          _nw9.__ctor((_69_i0).plus(_69_i0), !(p2) || (p2));
          _103_v29 = _nw9;
          (globalState).f18 = true;
          let _104_v31;
          _104_v31 = new _dafny.CodePoint('l'.codePointAt(0));
          let _105_v32;
          let _nw10 = new _module.C1();
          _nw10.__ctor(_104_v31, p0);
          _105_v32 = _nw10;
          let _106_v33;
          _106_v33 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(330),_105_v32);
          let _107_v34;
          _107_v34 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_103_v29).f28);
          let _108_v35;
          let _nw11 = new _module.C6();
          _nw11.__ctor(((function () {
            let _coll26 = new _dafny.Map();
            for (const _compr_26 of _dafny.IntegerRange(new BigNumber(606), new BigNumber(984))) {
              let _109_v30 = _compr_26;
              if (((new BigNumber(606)).isLessThanOrEqualTo(_109_v30)) && ((_109_v30).isLessThan(new BigNumber(984)))) {
                _coll26.push([(_109_v30).multipliedBy((_103_v29).f27),!(false)]);
              }
            }
            return _coll26;
          }()).update(p0, (_97_v23)[_module.__default.safeIndex(new BigNumber((_106_v33).length), new BigNumber((_97_v23).length))])).Merge(_107_v34));
          _108_v35 = _nw11;
          (globalState).f3 = (_dafny.Seq.UnicodeFromString("rohb"));
          (globalState).f16 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt((_103_v29).f27, _69_i0)).minus(((p3) ? (p0) : ((_dafny.ZERO).minus(p0)))));
        } else {
          let _110_v36;
          _110_v36 = _dafny.MultiSet.fromElements(p3);
          let _111_v37;
          _111_v37 = _dafny.Map.Empty.slice().updateUnsafe(p0,_110_v36);
          (globalState).f18 = ((((_111_v37).contains(new BigNumber(-178))) ? ((_111_v37).get(new BigNumber(-178))) : (_dafny.MultiSet.fromElements(false)))).IsDisjointFrom(_110_v36);
          let _112_v38;
          _112_v38 = _dafny.Seq.of(_dafny.Seq.Concat((_98_v24)[_module.__default.safeIndex(new BigNumber(814), new BigNumber((_98_v24).length))], _97_v23), _dafny.Seq.of(!(p2), p1, p3, p3, p2), _dafny.Seq.of(_module.__default.fm0(new BigNumber((_102_v28).length), p1, p3, _69_i0, globalState)), _97_v23, _97_v23);
          _112_v38 = _112_v38;
          let _113_v39;
          _113_v39 = _dafny.Set.fromElements(_102_v28);
          let _114_v40;
          _114_v40 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_113_v39).length),_69_i0);
          let _115_v41;
          _115_v41 = _dafny.Seq.of(_114_v40);
          (globalState).f18 = _module.__default.fm0(_69_i0, true, p1, (p0).multipliedBy(new BigNumber(((_115_v41)[_module.__default.safeIndex(p0, new BigNumber((_115_v41).length))]).length)), globalState);
          (globalState).f10 = (p0).plus(_69_i0);
          (globalState).f0 = (_69_i0).minus((_dafny.ZERO).minus(p0));
        }
      }
      let _116_v42;
      _116_v42 = _dafny.Seq.of(p1, p3);
      let _117_v43;
      let _nw12 = new _module.C5();
      _nw12.__ctor(p3, _module.__default.fm10(true, (_116_v42)[_module.__default.safeIndex(p0, new BigNumber((_116_v42).length))], p1, false, globalState), _module.__default.safeDivisionInt(p0, new BigNumber(12)));
      _117_v43 = _nw12;
      if (p2) {
        (globalState).f10 = (_dafny.ZERO).minus(p0);
        let _118_v44;
        _118_v44 = new _dafny.CodePoint('r'.codePointAt(0));
        let _119_v45;
        let _nw13 = new _module.C3();
        _nw13.__ctor(true, _118_v44, (p0).multipliedBy(new BigNumber((_dafny.Seq.UnicodeFromString("qicyr")).length)));
        _119_v45 = _nw13;
        let _120_v46;
        _120_v46 = _dafny.MultiSet.fromElements(p1);
        let _121_v47;
        _121_v47 = _dafny.Map.Empty.slice().updateUnsafe(_120_v46,p0);
        let _122_v48;
        let _nw14 = new _module.C4();
        _nw14.__ctor((_121_v47).Merge(_121_v47), (_117_v43).f26, _118_v44, p0);
        _122_v48 = _nw14;
        _122_v48 = _122_v48;
        let _123_v49;
        _123_v49 = _dafny.Set.fromElements((_117_v43).f26);
        let _124_v50;
        _124_v50 = _dafny.Seq.of(new BigNumber(754), _122_v48.f25);
        let _125_v51;
        let _nw15 = new _module.C4();
        _nw15.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_120_v46,_122_v48.f25), (p0).isLessThan(new BigNumber((_123_v49).length)), _module.__default.fm10(!(p3), !(p1), (_119_v45).fm18(globalState), (_117_v43).f26, globalState), new BigNumber((_124_v50).length));
        _125_v51 = _nw15;
        let _126_v52;
        let _init1 = ((_127_v51) => function (_128_i2) {
          return (_127_v51).f30;
        })(_125_v51);
        let _nw16 = Array((new BigNumber(3)).toNumber());
        for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw16.length); _i0_1++) {
          _nw16[_i0_1] = _init1(new BigNumber(_i0_1));
        }
        _126_v52 = _nw16;
        let _index10 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_126_v52).length));
        (_126_v52)[_index10] = (_module.__default.fm17((_117_v43).f26, p3, _119_v45.f31, globalState)).isLessThanOrEqualTo(_module.__default.fm17(p1, p3, !((_117_v43).f26), globalState));
        let _129_v53;
        let _init2 = function (_130_i3) {
          return (_130_i3).plus(new BigNumber(928));
        };
        let _nw17 = Array((new BigNumber(19)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw17.length); _i0_2++) {
          _nw17[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _129_v53 = _nw17;
        let _index11 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_129_v53).length));
        (_129_v53)[_index11] = p0;
        let _131_v54;
        _131_v54 = _dafny.Seq.of(_dafny.Seq.update(_124_v50, _module.__default.safeIndex(p0, new BigNumber((_124_v50).length)), new BigNumber(86)), _124_v50, _dafny.Seq.of(p0), _124_v50, _124_v50);
        let _index12 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_126_v52).length));
        let _index13 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_129_v53).length));
        let _rhs7 = _122_v48.f25;
        let _rhs8 = _125_v51;
        let _rhs9 = (_dafny.Set.fromElements(p1)).contains(p1);
        let _rhs10 = new BigNumber(-14);
        let _rhs11 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_122_v48.f25, new BigNumber((_dafny.Seq.Concat(_131_v54, _131_v54)).length)));
        let _lhs6 = globalState;
        let _lhs7 = _126_v52;
        let _lhs8 = _module.__default.safeIndex(new BigNumber(36), new BigNumber((_126_v52).length));
        let _lhs9 = globalState;
        let _lhs10 = _129_v53;
        let _lhs11 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_129_v53).length));
        _lhs6.f10 = _rhs7;
        _125_v51 = _rhs8;
        _lhs7[_lhs8] = _rhs9;
        _lhs9.f11 = _rhs10;
        _lhs10[_lhs11] = _rhs11;
        let _132_v55;
        _132_v55 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_129_v53)[_module.__default.safeIndex(new BigNumber(568), new BigNumber((_129_v53).length))]);
        let _133_v56;
        let _nw18 = Array((new BigNumber(2)).toNumber());
        _nw18[(_dafny.ZERO).toNumber()] = (((_125_v51).f30) ? (_132_v55) : (_132_v55));
        _nw18[(_dafny.ONE).toNumber()] = _132_v55;
        _133_v56 = _nw18;
        let _134_v57;
        _134_v57 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_116_v42).length)),!((_117_v43).f26));
        let _index14 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_133_v56).length));
        (_133_v56)[_index14] = _module.__default.fm31(new BigNumber((_134_v57).length), (_117_v43).f26, _module.__default.fm0(new BigNumber(967), p3, p3, new BigNumber((_116_v42).length), globalState), new BigNumber(262), globalState);
        let _index15 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_133_v56).length));
        (_133_v56)[_index15] = _132_v55;
      } else {
        let _135_v58;
        let _init3 = ((_136_p0) => function (_137_i4) {
          return _module.__default.safeDivisionInt(_137_i4, _136_p0);
        })(p0);
        let _nw19 = Array((new BigNumber(5)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw19.length); _i0_3++) {
          _nw19[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _135_v58 = _nw19;
        let _138_v59;
        _138_v59 = _dafny.Seq.UnicodeFromString("qdj");
        let _139_v60;
        _139_v60 = _dafny.Seq.of(_138_v59);
        let _140_v61;
        _140_v61 = new _dafny.CodePoint('q'.codePointAt(0));
        let _141_v62;
        _141_v62 = _dafny.Map.Empty.slice().updateUnsafe(p0,_140_v61);
        let _142_v63;
        let _nw20 = Array((new BigNumber(13)).toNumber());
        _nw20[(_dafny.ZERO).toNumber()] = p2;
        _nw20[(_dafny.ONE).toNumber()] = true;
        _nw20[(new BigNumber(2)).toNumber()] = p2;
        _nw20[(new BigNumber(3)).toNumber()] = _dafny.Seq.IsPrefixOf(_139_v60, _139_v60);
        _nw20[(new BigNumber(4)).toNumber()] = p3;
        _nw20[(new BigNumber(5)).toNumber()] = true;
        _nw20[(new BigNumber(6)).toNumber()] = p2;
        _nw20[(new BigNumber(7)).toNumber()] = (_117_v43).f26;
        _nw20[(new BigNumber(8)).toNumber()] = !_dafny.Seq.contains(_138_v59, _140_v61);
        _nw20[(new BigNumber(9)).toNumber()] = p3;
        _nw20[(new BigNumber(10)).toNumber()] = !_dafny.Seq.contains(_116_v42, p2);
        _nw20[(new BigNumber(11)).toNumber()] = !(_141_v62).equals(_141_v62);
        _nw20[(new BigNumber(12)).toNumber()] = (p0).isLessThanOrEqualTo(new BigNumber(-372));
        _142_v63 = _nw20;
        let _rhs12 = !((_117_v43).f26) || (p3);
        let _rhs13 = _135_v58;
        let _rhs14 = new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(34)), function (_143_i5) {
          return new _dafny.CodePoint('n'.codePointAt(0));
        })).length);
        let _rhs15 = (p0).plus(p0);
        let _rhs16 = _142_v63;
        let _lhs12 = globalState;
        let _lhs13 = globalState;
        let _lhs14 = globalState;
        _lhs12.f18 = _rhs12;
        _135_v58 = _rhs13;
        _lhs13.f16 = _rhs14;
        _lhs14.f11 = _rhs15;
        _142_v63 = _rhs16;
        let _144_v64;
        _144_v64 = _module.D13.create_DC29((_117_v43).f26, (_117_v43).f26, _142_v63);
        let _145_v65;
        _145_v65 = _dafny.Map.Empty.slice().updateUnsafe(_module.D13.create_DC29(p2, _module.__default.fm0(p0, p3, p2, p0, globalState), _142_v63),_135_v58);
        let _rhs17 = p1;
        let _rhs18 = (_145_v65).contains(_144_v64);
        let _rhs19 = p0;
        let _rhs20 = ((!(p1)) ? (p0) : (_module.__default.fm5(p0, p0, globalState)));
        let _lhs15 = globalState;
        let _lhs16 = globalState;
        let _lhs17 = globalState;
        let _lhs18 = globalState;
        _lhs15.f18 = _rhs17;
        _lhs16.f18 = _rhs18;
        _lhs17.f16 = _rhs19;
        _lhs18.f22 = _rhs20;
        let _146_v66;
        let _nw21 = new _module.C2();
        _nw21.__ctor(p0, (_117_v43).f26);
        _146_v66 = _nw21;
        let _147_v67;
        _147_v67 = _dafny.Set.fromElements(_146_v66, _146_v66, _146_v66, _146_v66, _146_v66);
        let _148_v68;
        _148_v68 = _dafny.Seq.of(_147_v67, _147_v67, _147_v67, _147_v67, _147_v67);
        _148_v68 = _dafny.Seq.Concat(_148_v68, _148_v68);
        let _149_v69;
        _149_v69 = _dafny.Map.Empty.slice().updateUnsafe((_146_v66).f28,(_117_v43).f26);
        _149_v69 = (_149_v69).update(p1, p1);
        let _150_v70;
        let _nw22 = new _module.C3();
        _nw22.__ctor(p1, _140_v61, p0);
        _150_v70 = _nw22;
        let _151_v71;
        _151_v71 = _dafny.Seq.of(_150_v70, _150_v70, _150_v70, _150_v70);
        let _152_v72;
        _152_v72 = _dafny.Map.Empty.slice().updateUnsafe((_117_v43).f26,_151_v71);
        let _153_v73;
        _153_v73 = _dafny.Seq.of(_151_v71, _151_v71, _151_v71, _151_v71);
        let _154_v74;
        _154_v74 = _dafny.Map.Empty.slice().updateUnsafe((_146_v66).f27,_151_v71);
        let _155_v75;
        let _nw23 = new _module.C1();
        _nw23.__ctor(new _dafny.CodePoint('g'.codePointAt(0)), p0);
        _155_v75 = _nw23;
        let _156_v76;
        _156_v76 = _dafny.Map.Empty.slice().updateUnsafe(_155_v75,_151_v71);
        let _157_v77;
        let _nw24 = Array((new BigNumber(21)).toNumber());
        _nw24[(_dafny.ZERO).toNumber()] = ((p2) ? (_151_v71) : (_151_v71));
        _nw24[(_dafny.ONE).toNumber()] = _151_v71;
        _nw24[(new BigNumber(2)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(3)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(4)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(_150_v70);
        _nw24[(new BigNumber(6)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(7)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_151_v71, _dafny.Seq.update(_151_v71, _module.__default.safeIndex((_146_v66).f27, new BigNumber((_151_v71).length)), _150_v70));
        _nw24[(new BigNumber(9)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(10)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(11)).toNumber()] = _dafny.Seq.of(_150_v70, _150_v70);
        _nw24[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat((((_152_v72).contains(!(p2))) ? ((_152_v72).get(!(p2))) : (_151_v71)), _dafny.Seq.update((_153_v73)[_module.__default.safeIndex(p0, new BigNumber((_153_v73).length))], _module.__default.safeIndex(p0, new BigNumber(((_153_v73)[_module.__default.safeIndex(p0, new BigNumber((_153_v73).length))]).length)), _150_v70));
        _nw24[(new BigNumber(13)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(14)).toNumber()] = (((_154_v74).contains((_146_v66).f27)) ? ((_154_v74).get((_146_v66).f27)) : (_151_v71));
        _nw24[(new BigNumber(15)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(16)).toNumber()] = _dafny.Seq.of(_150_v70);
        _nw24[(new BigNumber(17)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(18)).toNumber()] = _151_v71;
        _nw24[(new BigNumber(19)).toNumber()] = (((_156_v76).contains(_155_v75)) ? ((_156_v76).get(_155_v75)) : (_151_v71));
        _nw24[(new BigNumber(20)).toNumber()] = _151_v71;
        _157_v77 = _nw24;
        _157_v77 = _157_v77;
      }
      let _158_v78;
      let _init4 = ((_159_v43) => function (_160_i7) {
        return (_159_v43).f26;
      })(_117_v43);
      let _nw25 = Array((new BigNumber(22)).toNumber());
      for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw25.length); _i0_4++) {
        _nw25[_i0_4] = _init4(new BigNumber(_i0_4));
      }
      _158_v78 = _nw25;
      let _source4 = _module.D13.create_DC29(p2, !_dafny.areEqual(_dafny.Seq.Create(_module.__default.abs(new BigNumber(995)), ((_161_p0) => function (_162_i6) {
  return _161_p0;
})(p0)), _dafny.Seq.of(new BigNumber(7), p0)), _158_v78);
      if (_source4.is_DC29) {
        let _163___mcc_h0 = (_source4).cf43;
        let _164___mcc_h1 = (_source4).cf44;
        let _165___mcc_h2 = (_source4).cf45;
        let _166_cf45 = _165___mcc_h2;
        let _167_cf44 = _164___mcc_h1;
        let _168_cf43 = _163___mcc_h0;
        let _169_v79;
        _169_v79 = _dafny.Map.Empty.slice().updateUnsafe(true,p0);
        let _170_v80;
        _170_v80 = _dafny.Map.Empty.slice().updateUnsafe(p0,_169_v79);
        let _171_v81;
        _171_v81 = _dafny.Seq.UnicodeFromString("mw");
        (globalState).f16 = new BigNumber(((_169_v79).Merge((((_170_v80).contains(new BigNumber((_171_v81).length))) ? ((_170_v80).get(new BigNumber((_171_v81).length))) : (_169_v79)))).length);
        let _172_v82;
        let _nw26 = new _module.C0();
        _nw26.__ctor();
        _172_v82 = _nw26;
        let _173_v83;
        let _nw27 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Set.Empty);
        _173_v83 = _nw27;
        let _174_v84;
        _174_v84 = _module.D10.create_DC20(_167_cf44, _module.__default.fm5(p0, p0, globalState));
        let _175_v85;
        _175_v85 = _dafny.Set.fromElements(_174_v84);
        let _index16 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_173_v83).length));
        (_173_v83)[_index16] = _175_v85;
        let _176_v87;
        _176_v87 = _dafny.Seq.of(_174_v84, _174_v84);
        let _177_v88;
        _177_v88 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((function () {
          let _coll27 = new _dafny.Set();
          for (const _compr_27 of _dafny.IntegerRange(new BigNumber(961), new BigNumber(914))) {
            let _178_v86 = _compr_27;
            if (((new BigNumber(961)).isLessThanOrEqualTo(_178_v86)) && ((_178_v86).isLessThan(new BigNumber(914)))) {
              _coll27.add((_178_v86).plus(new BigNumber((_171_v81).length)));
            }
          }
          return _coll27;
        }()).length)),_176_v87);
        let _179_v89;
        _179_v89 = _dafny.Seq.of(p0);
        let _index17 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_173_v83).length));
        (_173_v83)[_index17] = (_175_v85).Union(function () {
          let _coll28 = new _dafny.Set();
          for (const _compr_28 of ((((_177_v88).contains(new BigNumber((_179_v89).length))) ? ((_177_v88).get(new BigNumber((_179_v89).length))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), ((_180_cf44, _181_v43, _182_p2, _183_p0) => function (_184_i8) {
            return _module.D10.create_DC20((_module.D6.create_DC12(_180_cf44, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_181_v43).f26,_182_p2)).length))).dtor_cf17, _183_p0);
          })(_167_cf44, _117_v43, p2, p0))))).Elements) {
            let _185_v90 = _compr_28;
            if (_dafny.Seq.contains((((_177_v88).contains(new BigNumber((_179_v89).length))) ? ((_177_v88).get(new BigNumber((_179_v89).length))) : (_dafny.Seq.Create(_module.__default.abs(new BigNumber(782)), ((_186_cf44, _187_v43, _188_p2, _189_p0) => function (_184_i8) {
              return _module.D10.create_DC20((_module.D6.create_DC12(_186_cf44, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_187_v43).f26,_188_p2)).length))).dtor_cf17, _189_p0);
            })(_167_cf44, _117_v43, p2, p0)))), _185_v90)) {
              _coll28.add(_185_v90);
            }
          }
          return _coll28;
        }());
        let _190_v91;
        let _nw28 = Array((new BigNumber(28)).toNumber()).fill(_dafny.ZERO);
        _190_v91 = _nw28;
        let _191_v92;
        _191_v92 = _dafny.Map.Empty.slice().updateUnsafe(p0,_117_v43);
        let _192_v93;
        _192_v93 = _module.D2.create_DC3(new BigNumber((_171_v81).length), new BigNumber((_191_v92).length));
        let _193_v94;
        _193_v94 = _module.D7.create_DC14(p2, p0, (_192_v93).dtor_cf3);
        let _index18 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_190_v91).length));
        (_190_v91)[_index18] = _module.__default.safeDivisionInt((_193_v94).dtor_cf22, new BigNumber((_module.__default.fm26(globalState)).length));
        let _194_v95;
        let _nw29 = new _module.C3();
        _nw29.__ctor(_167_cf44, (_171_v81)[_module.__default.safeIndex(p0, new BigNumber((_171_v81).length))], new BigNumber(547));
        _194_v95 = _nw29;
        let _195_v96;
        _195_v96 = _dafny.Map.Empty.slice().updateUnsafe(_194_v95,_171_v81);
        let _index19 = _module.__default.safeIndex(new BigNumber(765), new BigNumber((_190_v91).length));
        (_190_v91)[_index19] = _module.__default.fm17((_195_v96).contains(_194_v95), p2, (_117_v43).f26, globalState);
      } else if (_source4.is_DC28) {
        let _196___mcc_h3 = (_source4).cf42;
        let _197_cf42 = _196___mcc_h3;
        (globalState).f10 = (p0).multipliedBy((new BigNumber(667)).minus(new BigNumber(450)));
        let _198_v97;
        _198_v97 = _dafny.MultiSet.fromElements(p2);
        let _199_v98;
        _199_v98 = _dafny.Map.Empty.slice().updateUnsafe(_198_v97,p0);
        let _200_v99;
        _200_v99 = new _dafny.CodePoint('v'.codePointAt(0));
        let _201_v100;
        _201_v100 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
        let _202_v101;
        let _nw30 = new _module.C4();
        _nw30.__ctor(_199_v98, p3, _200_v99, p0);
        _202_v101 = _nw30;
        let _203_v102;
        _203_v102 = _dafny.Map.Empty.slice().updateUnsafe(_202_v101,p0);
        let _204_v103;
        _204_v103 = _dafny.MultiSet.fromElements(new BigNumber((_203_v102).length), _202_v101.f25);
        let _205_v104;
        _205_v104 = _dafny.Seq.of(p0, new BigNumber((_201_v100).length), new BigNumber((_204_v103).cardinality()));
        let _206_v105;
        let _nw31 = new _module.C4();
        _nw31.__ctor(_199_v98, p3, _200_v99, new BigNumber((_205_v104).length));
        _206_v105 = _nw31;
        let _207_v106;
        _207_v106 = _dafny.MultiSet.fromElements(((p2) ? (_206_v105) : (_206_v105)), _206_v105);
        let _rhs21 = (_207_v106).Difference(_207_v106);
        let _rhs22 = !((p0).plus((_dafny.ZERO).minus(_206_v105.f25))).isEqualTo(p0);
        let _lhs19 = globalState;
        _207_v106 = _rhs21;
        _lhs19.f18 = _rhs22;
        if (!((!(p3)) && (((p3) ? ((_117_v43).f26) : (p2))))) {
          let _208_v107;
          let _nw32 = new _module.C3();
          _nw32.__ctor((_117_v43).f26, (_206_v105).f24, _206_v105.f25);
          _208_v107 = _nw32;
          let _209_v108;
          _209_v108 = _208_v107;
          let _210_v109;
          _210_v109 = _dafny.Map.Empty.slice().updateUnsafe((_209_v108),_dafny.Seq.of(new BigNumber(563)));
          let _211_v110;
          _211_v110 = _dafny.MultiSet.fromElements(_205_v104, _205_v104, _dafny.Seq.of(_206_v105.f25));
          let _212_v111;
          let _nw33 = Array((new BigNumber(20)).toNumber());
          _nw33[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_205_v104, _module.__default.safeIndex(new BigNumber(-929), new BigNumber((_205_v104).length)), (_205_v104)[_module.__default.safeIndex(_module.__default.fm17((_117_v43).f26, p1, false, globalState), new BigNumber((_205_v104).length))]);
          _nw33[(_dafny.ONE).toNumber()] = _205_v104;
          _nw33[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(817)), ((_213_v101) => function (_214_i9) {
            return _213_v101.f25;
          })(_202_v101));
          _nw33[(new BigNumber(3)).toNumber()] = (((_210_v109).contains(_208_v107)) ? ((_210_v109).get(_208_v107)) : (_205_v104));
          _nw33[(new BigNumber(4)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(5)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(6)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(7)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(8)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(9)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-330)), ((_215_v97) => function (_216_i10) {
            return new BigNumber((_215_v97).cardinality());
          })(_198_v97));
          _nw33[(new BigNumber(10)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(11)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(12)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(13)).toNumber()] = _dafny.Seq.of(_206_v105.f25, new BigNumber((_116_v42).length), _202_v101.f25);
          _nw33[(new BigNumber(14)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(15)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(16)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(-866)), ((_217_v104) => function (_218_i11) {
            return (_217_v104)[_module.__default.safeIndex(new BigNumber(237), new BigNumber((_217_v104).length))];
          })(_205_v104));
          _nw33[(new BigNumber(17)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(18)).toNumber()] = _205_v104;
          _nw33[(new BigNumber(19)).toNumber()] = _dafny.Seq.update(_205_v104, _module.__default.safeIndex(_206_v105.f25, new BigNumber((_205_v104).length)), new BigNumber((_211_v110).cardinality()));
          _212_v111 = _nw33;
          let _219_v112;
          _219_v112 = _dafny.Map.Empty.slice().updateUnsafe((_117_v43).f26,_212_v111);
          _219_v112 = (_219_v112).update((_117_v43).f26, _212_v111);
          let _index20 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_212_v111).length));
          (_212_v111)[_index20] = _205_v104;
          let _index21 = _module.__default.safeIndex(new BigNumber(434), new BigNumber((_212_v111).length));
          (_212_v111)[_index21] = _dafny.Seq.Concat(_dafny.Seq.of(_206_v105.f25, _202_v101.f25), _dafny.Seq.Concat(_205_v104, _205_v104));
          let _220_v113;
          _220_v113 = _dafny.Map.Empty.slice().updateUnsafe(p3,(_116_v42)[_module.__default.safeIndex(p0, new BigNumber((_116_v42).length))]);
          let _221_v114;
          _221_v114 = _module.D11.create_DC24((_117_v43).f26, new BigNumber((_module.__default.fm34(p2, new BigNumber((_220_v113).length), (_117_v43).f26, globalState)).length));
          let _222_v115;
          let _nw34 = new _module.C1();
          _nw34.__ctor((_202_v101).f24, (_202_v101.f25).minus((_221_v114).dtor_cf37));
          _222_v115 = _nw34;
          let _223_v116;
          _223_v116 = _module.D3.create_DC5(_116_v42, _202_v101.f25, _202_v101.f25);
          let _224_v117;
          _224_v117 = _module.D10.create_DC20((_117_v43).f26, (_223_v116).dtor_cf7);
          (globalState).f18 = (_224_v117).dtor_cf28;
          _200_v99 = new _dafny.CodePoint('a'.codePointAt(0));
        } else {
          let _225_v118;
          _225_v118 = _dafny.Map.Empty.slice().updateUnsafe((_117_v43).f26,p0);
          (globalState).f11 = (((_225_v118).contains(p3)) ? ((_225_v118).get(p3)) : (_202_v101.f25));
          let _226_v119;
          _226_v119 = _module.D3.create_DC5(_116_v42, _206_v105.f25, _206_v105.f25);
          _226_v119 = _226_v119;
          let _index22 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_158_v78).length));
          (_158_v78)[_index22] = (_198_v97).IsDisjointFrom(_198_v97);
          let _index23 = _module.__default.safeIndex(new BigNumber(116), new BigNumber((_158_v78).length));
          (_158_v78)[_index23] = (_117_v43).f26;
          let _227_v120;
          let _nw35 = Array((new BigNumber(19)).toNumber()).fill(_dafny.ZERO);
          _227_v120 = _nw35;
          let _index24 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_227_v120).length));
          (_227_v120)[_index24] = _202_v101.f25;
          let _228_v121;
          _228_v121 = _module.D10.create_DC21(_202_v101.f25, !((_158_v78)[_module.__default.safeIndex(new BigNumber(116), new BigNumber((_158_v78).length))]), (_117_v43).f26);
          let _229_v122;
          _229_v122 = _module.D11.create_DC25(_module.D11.create_DC24(p3, _module.__default.fm17((_module.D11.create_DC24(p3, _206_v105.f25)).dtor_cf36, (_228_v121).dtor_cf32, p2, globalState)));
          let _230_v123;
          _230_v123 = _dafny.MultiSet.fromElements(_229_v122, _229_v122, _229_v122);
          let _index25 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_227_v120).length));
          let _rhs23 = new BigNumber(795);
          let _rhs24 = (((_225_v118).contains((_117_v43).f26)) ? ((_225_v118).get((_117_v43).f26)) : (new BigNumber((_230_v123).cardinality())));
          let _rhs25 = _module.__default.safeModuloInt((new BigNumber(-545)).plus(_206_v105.f25), _202_v101.f25);
          let _lhs20 = _227_v120;
          let _lhs21 = _module.__default.safeIndex(new BigNumber(232), new BigNumber((_227_v120).length));
          let _lhs22 = globalState;
          let _lhs23 = globalState;
          _lhs20[_lhs21] = _rhs23;
          _lhs22.f22 = _rhs24;
          _lhs23.f10 = _rhs25;
          let _231_v124;
          let _init5 = ((_232_v101, _233_v120) => function (_234_i12) {
            return (_dafny.MultiSet.FromArray(_dafny.Seq.of(_dafny.Set.fromElements((_233_v120)[_module.__default.safeIndex(new BigNumber(232), new BigNumber((_233_v120).length))], _232_v101.f25), _dafny.Set.fromElements(_232_v101.f25)))).IsSubsetOf(_dafny.MultiSet.fromElements(_dafny.Set.fromElements(_232_v101.f25)));
          })(_202_v101, _227_v120);
          let _nw36 = Array((new BigNumber(29)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw36.length); _i0_5++) {
            _nw36[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _231_v124 = _nw36;
          _231_v124 = ((p1) ? (_231_v124) : (_231_v124));
        }
        let _235_v125;
        let _init6 = ((_236_v104) => function (_237_i13) {
          return _236_v104;
        })(_205_v104);
        let _nw37 = Array((new BigNumber(23)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw37.length); _i0_6++) {
          _nw37[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _235_v125 = _nw37;
        let _index26 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_235_v125).length));
        (_235_v125)[_index26] = _205_v104;
        let _238_v126;
        _238_v126 = _module.D10.create_DC19((_206_v105).f24);
        let _239_v127;
        _239_v127 = _dafny.Map.Empty.slice().updateUnsafe(_158_v78,(_117_v43).f26);
        let _index27 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_235_v125).length));
        let _rhs26 = (_206_v105.f25).multipliedBy(_206_v105.f25);
        let _rhs27 = (_dafny.ZERO).minus(p0);
        let _rhs28 = _dafny.Seq.of((new BigNumber(703)).plus(new BigNumber(-504)), new BigNumber((_dafny.Seq.UnicodeFromString("pvsknr")).length));
        let _rhs29 = (new BigNumber((_239_v127).length)).plus(_206_v105.f25);
        let _rhs30 = _238_v126;
        let _lhs24 = globalState;
        let _lhs25 = _202_v101;
        let _lhs26 = _235_v125;
        let _lhs27 = _module.__default.safeIndex(new BigNumber(828), new BigNumber((_235_v125).length));
        let _lhs28 = globalState;
        _lhs24.f0 = _rhs26;
        _lhs25.f25 = _rhs27;
        _lhs26[_lhs27] = _rhs28;
        _lhs28.f2 = _rhs29;
        _238_v126 = _rhs30;
      } else {
        let _240___mcc_h4 = (_source4).cf46;
        let _241_cf46 = _240___mcc_h4;
        let _242_v128;
        _242_v128 = _dafny.Set.fromElements(p0);
        let _243_v129;
        _243_v129 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_242_v128).length));
        let _244_v130;
        _244_v130 = _module.D11.create_DC24(p2, new BigNumber((_116_v42).length));
        let _245_v131;
        _245_v131 = _dafny.Map.Empty.slice().updateUnsafe((((_243_v129).contains((_244_v130).dtor_cf36)) ? ((_243_v129).get((_244_v130).dtor_cf36)) : (p0)),p0);
        (globalState).f10 = (((_245_v131).contains(p0)) ? ((_245_v131).get(p0)) : (p0));
        let _index28 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_158_v78).length));
        (_158_v78)[_index28] = p2;
        let _index29 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_158_v78).length));
        let _rhs31 = (_117_v43).f26;
        let _rhs32 = (_dafny.ZERO).minus((_dafny.ZERO).minus(p0));
        let _lhs29 = _158_v78;
        let _lhs30 = _module.__default.safeIndex(new BigNumber(236), new BigNumber((_158_v78).length));
        let _lhs31 = globalState;
        _lhs29[_lhs30] = _rhs31;
        _lhs31.f0 = _rhs32;
        let _246_v132;
        _246_v132 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
        let _247_v133;
        _247_v133 = _dafny.Set.fromElements(_246_v132);
        let _248_v134;
        _248_v134 = _module.D3.create_DC4(_247_v133);
        let _249_v135;
        _249_v135 = _116_v42;
        let _250_v136;
        _250_v136 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber(((_249_v135)).length)),_module.D3.create_DC4(_247_v133));
        let _251_v137;
        let _init7 = function (_252_i14) {
          return new _dafny.CodePoint('t'.codePointAt(0));
        };
        let _nw38 = Array((new BigNumber(16)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw38.length); _i0_7++) {
          _nw38[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _251_v137 = _nw38;
        let _253_v138;
        _253_v138 = _dafny.Seq.of(_251_v137, _251_v137);
        let _254_v139;
        let _255_v140;
        let _256_v141;
        let _out0;
        let _out1;
        let _out2;
        let _outcollector0 = (_117_v43).m1((_dafny.Map.Empty.slice().updateUnsafe(p0,_248_v134)).Merge(_250_v136), ((_116_v42)[_module.__default.safeIndex(p0, new BigNumber((_116_v42).length))]) === (p3), _253_v138, globalState);
        _out0 = _outcollector0[0];
        _out1 = _outcollector0[1];
        _out2 = _outcollector0[2];
        _254_v139 = _out0;
        _255_v140 = _out1;
        _256_v141 = _out2;
        let _257_v142;
        let _nw39 = new _module.C0();
        _nw39.__ctor();
        _257_v142 = _nw39;
      }
      let _258_v143;
      _258_v143 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
      let _259_v144;
      _259_v144 = _dafny.Map.Empty.slice().updateUnsafe(_258_v143,_158_v78);
      _259_v144 = (_259_v144).update(function () {
        let _coll29 = new _dafny.Map();
        for (const _compr_29 of _dafny.IntegerRange(new BigNumber(691), new BigNumber(699))) {
          let _260_v145 = _compr_29;
          if (((new BigNumber(691)).isLessThanOrEqualTo(_260_v145)) && ((_260_v145).isLessThan(new BigNumber(699)))) {
            _coll29.push([(_260_v145).multipliedBy(new BigNumber(919)),p3]);
          }
        }
        return _coll29;
      }(), _158_v78);
      let _261_v146;
      let _init8 = ((_262_p0) => function (_263_i15) {
        return _module.__default.safeModuloInt(_263_i15, _262_p0);
      })(p0);
      let _nw40 = Array((new BigNumber(10)).toNumber());
      for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw40.length); _i0_8++) {
        _nw40[_i0_8] = _init8(new BigNumber(_i0_8));
      }
      _261_v146 = _nw40;
      let _index30 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_261_v146).length));
      (_261_v146)[_index30] = p0;
      let _index31 = _module.__default.safeIndex(new BigNumber(237), new BigNumber((_261_v146).length));
      (_261_v146)[_index31] = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.update(_dafny.Seq.Concat(_116_v42, _116_v42), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_116_v42, _116_v42)).length)), p3))).cardinality());
      let _264_v147;
      _264_v147 = _module.D10.create_DC21(new BigNumber(124), p2, p1);
      r0 = function (_source5) {
        if (_source5.is_DC19) {
          let _265___mcc_h5 = (_source5).cf27;
          let _266_cf27 = _265___mcc_h5;
          return _266_cf27;
        } else if (_source5.is_DC20) {
          let _267___mcc_h6 = (_source5).cf28;
          let _268___mcc_h7 = (_source5).cf29;
          let _269_cf29 = _268___mcc_h7;
          let _270_cf28 = _267___mcc_h6;
          return new _dafny.CodePoint('l'.codePointAt(0));
        } else if (_source5.is_DC21) {
          let _271___mcc_h8 = (_source5).cf30;
          let _272___mcc_h9 = (_source5).cf31;
          let _273___mcc_h10 = (_source5).cf32;
          let _274_cf32 = _273___mcc_h10;
          let _275_cf31 = _272___mcc_h9;
          let _276_cf30 = _271___mcc_h8;
          return new _dafny.CodePoint('n'.codePointAt(0));
        } else {
          let _277___mcc_h11 = (_source5).cf26;
          let _278_cf26 = _277___mcc_h11;
          return new _dafny.CodePoint('o'.codePointAt(0));
        }
      }(_264_v147);
      return r0;
    }
    static Main(__noArgsParameter) {
      let _279_v0;
      _279_v0 = _dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)));
      let _280_v1;
      _280_v1 = true;
      let _281_v2;
      _281_v2 = new BigNumber(742);
      let _282_v3;
      _282_v3 = _dafny.Map.Empty.slice().updateUnsafe(_280_v1,_281_v2);
      let _283_v4;
      let _nw41 = Array((new BigNumber(27)).toNumber()).fill(false);
      _283_v4 = _nw41;
      let _284_v5;
      let _nw42 = Array((new BigNumber(25)).toNumber()).fill(_dafny.ZERO);
      _284_v5 = _nw42;
      let _285_v6;
      _285_v6 = _dafny.Map.Empty.slice().updateUnsafe(_283_v4,_284_v5);
      let _286_v7;
      _286_v7 = _dafny.MultiSet.fromElements(_280_v1);
      let _287_v8;
      _287_v8 = _dafny.Seq.of(new BigNumber((_286_v7).cardinality()));
      let _288_v9;
      _288_v9 = _dafny.Set.fromElements(_287_v8);
      let _289_globalState;
      let _nw43 = new _module.GlobalState();
      _nw43.__ctor(new BigNumber(-928), new BigNumber(141), new BigNumber(124), _dafny.Seq.Concat(_279_v0, _279_v0), new BigNumber(227), false, new _dafny.CodePoint('d'.codePointAt(0)), _282_v3, _285_v6, new BigNumber(-50), new BigNumber(-492), new BigNumber(-627), false, (_282_v3).Merge(_282_v3), _288_v9, new BigNumber(968), new BigNumber(343), false, false, _287_v8, new BigNumber(425), true, new BigNumber(-118));
      _289_globalState = _nw43;
      let _290_v10;
      _290_v10 = _dafny.MultiSet.fromElements(_281_v2, _281_v2);
      let _291_v11;
      _291_v11 = _dafny.Seq.of(_280_v1, _module.__default.fm0(new BigNumber((_290_v10).cardinality()), _280_v1, _280_v1, new BigNumber((_282_v3).length), _289_globalState));
      let _292_v12;
      _292_v12 = _291_v11;
      let _293_v13;
      _293_v13 = _dafny.Seq.of(_284_v5, _284_v5, _284_v5, _284_v5);
      let _294_v14;
      _294_v14 = _279_v0;
      let _295_v15;
      _295_v15 = _dafny.Seq.of((_294_v14), _279_v0, _279_v0, _dafny.Seq.UnicodeFromString("fln"), _279_v0);
      let _296_v16;
      _296_v16 = new _dafny.CodePoint('j'.codePointAt(0));
      let _rhs33 = _279_v0;
      let _rhs34 = _281_v2;
      let _rhs35 = _281_v2;
      let _rhs36 = (_dafny.ZERO).minus((_281_v2).minus(new BigNumber((_dafny.Seq.Concat((_292_v12), _291_v11)).length)));
      let _rhs37 = (_293_v13)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update((_295_v15)[_module.__default.safeIndex(_281_v2, new BigNumber((_295_v15).length))], _module.__default.safeIndex(new BigNumber(((_295_v15)[_module.__default.safeIndex(_281_v2, new BigNumber((_295_v15).length))]).length), new BigNumber(((_295_v15)[_module.__default.safeIndex(_281_v2, new BigNumber((_295_v15).length))]).length)), _296_v16)).length), new BigNumber((_293_v13).length))];
      let _lhs32 = _289_globalState;
      let _lhs33 = _289_globalState;
      let _lhs34 = _289_globalState;
      _279_v0 = _rhs33;
      _lhs32.f11 = _rhs34;
      _lhs33.f10 = _rhs35;
      _lhs34.f16 = _rhs36;
      _284_v5 = _rhs37;
      (_289_globalState).f2 = (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_287_v8, _287_v8), _287_v8)).length));
      let _297_i0;
      _297_i0 = _dafny.ZERO;
      L0: {
        while (true) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_297_i0)) {
              break L0;
            }
            _297_i0 = (_297_i0).plus(_dafny.ONE);
            let _298_v17;
            let _nw44 = new _module.C4();
            _nw44.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_286_v7,new BigNumber((_291_v11).length)), !(_280_v1), _296_v16, _281_v2);
            _298_v17 = _nw44;
            _296_v16 = ((!(!(_280_v1)) || (_280_v1)) ? (_296_v16) : (_296_v16));
            let _source6 = _294_v14;
            let _299___mcc_h0 = _source6;
            let _300_cf1 = _299___mcc_h0;
            let _index32 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_283_v4).length));
            (_283_v4)[_index32] = false;
            let _301_v18;
            _301_v18 = _dafny.Map.Empty.slice().updateUnsafe(_281_v2,(((_282_v3).contains((_291_v11)[_module.__default.safeIndex(_281_v2, new BigNumber((_291_v11).length))])) ? ((_282_v3).get((_291_v11)[_module.__default.safeIndex(_281_v2, new BigNumber((_291_v11).length))])) : (_281_v2)));
            let _302_v19;
            _302_v19 = _module.D3.create_DC6((((_301_v18).contains(_281_v2)) ? ((_301_v18).get(_281_v2)) : (_281_v2)), _281_v2, (_298_v17).f30);
            let _index33 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_283_v4).length));
            (_283_v4)[_index33] = (_302_v19).dtor_cf11;
            (_289_globalState).f0 = (_287_v8)[_module.__default.safeIndex(_281_v2, new BigNumber((_287_v8).length))];
            (_289_globalState).f15 = _281_v2;
            let _index34 = _module.__default.safeIndex(new BigNumber(947), new BigNumber((_283_v4).length));
            (_283_v4)[_index34] = false;
            let _303_v20;
            _303_v20 = _module.D5.create_DC10(false, true);
            let _304_v21;
            _304_v21 = _dafny.Map.Empty.slice().updateUnsafe(!(!((_298_v17).f30)),_303_v20);
            (_298_v17).m5((_304_v21).Merge(_304_v21), true, _281_v2, _289_globalState);
          }
        }
      }
      let _305_v22;
      _305_v22 = _dafny.Set.fromElements(_281_v2, new BigNumber((_291_v11).length));
      let _306_v23;
      _306_v23 = _dafny.Map.Empty.slice().updateUnsafe((_287_v8)[_module.__default.safeIndex(_281_v2, new BigNumber((_287_v8).length))],_module.__default.fm0(_281_v2, _280_v1, false, (((_282_v3).contains(_280_v1)) ? ((_282_v3).get(_280_v1)) : (_281_v2)), _289_globalState));
      (_289_globalState).f2 = ((false) ? (new BigNumber((_305_v22).length)) : (new BigNumber((_306_v23).length)));
      let _307_v24;
      let _out3;
      _out3 = _module.__default.m6(_281_v2, _280_v1, _280_v1, _280_v1, _289_globalState);
      _307_v24 = _out3;
      let _308_v25;
      _308_v25 = _module.D11.create_DC23(_281_v2, _280_v1);
      let _309_v26;
      _309_v26 = _module.D11.create_DC25(_308_v25);
      let _310_v27;
      let _nw45 = Array((new BigNumber(4)).toNumber());
      _nw45[(_dafny.ZERO).toNumber()] = _309_v26;
      _nw45[(_dafny.ONE).toNumber()] = _309_v26;
      _nw45[(new BigNumber(2)).toNumber()] = _309_v26;
      _nw45[(new BigNumber(3)).toNumber()] = _module.D11.create_DC25(_308_v25);
      _310_v27 = _nw45;
      let _311_v28;
      _311_v28 = _dafny.Set.fromElements(_310_v27, _310_v27);
      let _312_v29;
      _312_v29 = _dafny.Map.Empty.slice().updateUnsafe(_280_v1,_311_v28);
      _312_v29 = (_312_v29).update(_280_v1, _311_v28);
      if (_280_v1) {
        let _313_v30;
        let _nw46 = new _module.C3();
        _nw46.__ctor(true, _296_v16, (_287_v8)[_module.__default.safeIndex(_281_v2, new BigNumber((_287_v8).length))]);
        _313_v30 = _nw46;
        let _314_v31;
        _314_v31 = _dafny.MultiSet.fromElements(_313_v30);
        let _315_v32;
        let _out4;
        _out4 = _module.__default.m6(_281_v2, (_281_v2).isEqualTo(_module.__default.fm5(_281_v2, new BigNumber((_314_v31).cardinality()), _289_globalState)), _dafny.areEqual(_279_v0, _279_v0), (_281_v2).isLessThan(_281_v2), _289_globalState);
        _315_v32 = _out4;
        if (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xvkixbphw"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(618)), ((_316_v24) => function (_317_i1) {
          return _316_v24;
        })(_307_v24))), _dafny.Seq.Concat(_279_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(912)), ((_318_v16) => function (_319_i2) {
          return _318_v16;
        })(_296_v16))))) {
          _280_v1 = _dafny.Seq.contains(_dafny.Seq.Concat(_287_v8, _287_v8), (new BigNumber((_dafny.Set.fromElements(_313_v30)).length)).multipliedBy(_313_v30.f25));
          (_289_globalState).f18 = _280_v1;
          let _320_v33;
          let _out5;
          _out5 = _module.__default.m6((_313_v30.f25).multipliedBy(_281_v2), _280_v1, (_286_v7).IsDisjointFrom(_dafny.MultiSet.fromElements(_280_v1, _280_v1)), (_313_v30.f25).isLessThanOrEqualTo(_281_v2), _289_globalState);
          _320_v33 = _out5;
          let _321_v34;
          _321_v34 = _dafny.Map.Empty.slice().updateUnsafe(_286_v7,_313_v30.f25);
          let _322_v36;
          _322_v36 = _dafny.Set.fromElements(_module.__default.fm8(_313_v30.f25, true, _dafny.Seq.of(_280_v1), _289_globalState));
          let _323_v37;
          let _nw47 = new _module.C4();
          _nw47.__ctor((_321_v34).Merge(function () {
            let _coll30 = new _dafny.Map();
            for (const _compr_30 of (_322_v36).Elements) {
              let _324_v35 = _compr_30;
              if ((_322_v36).contains(_324_v35)) {
                _coll30.push([_324_v35,_313_v30.f25]);
              }
            }
            return _coll30;
          }()), ((true) ? (_280_v1) : (_280_v1)), new _dafny.CodePoint('b'.codePointAt(0)), _313_v30.f25);
          _323_v37 = _nw47;
          (_289_globalState).f3 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(941)), ((_325_v30) => function (_326_i3) {
            return (_325_v30).f24;
          })(_313_v30));
        } else {
          let _327_v38;
          _327_v38 = _dafny.Map.Empty.slice().updateUnsafe((_290_v10).IsDisjointFrom(_290_v10),_283_v4);
          let _328_v39;
          _328_v39 = _dafny.Map.Empty.slice().updateUnsafe((_313_v30).f24,_280_v1);
          let _329_v40;
          _329_v40 = _module.D13.create_DC29(_280_v1, _280_v1, _283_v4);
          _327_v38 = ((_dafny.Map.Empty.slice().updateUnsafe(_280_v1,_283_v4)).update((((_328_v39).contains((_313_v30).f24)) ? ((_328_v39).get((_313_v30).f24)) : (_280_v1)), (_329_v40).dtor_cf45)).Merge(_327_v38);
          (_289_globalState).f18 = _280_v1;
          let _330_v41;
          let _nw48 = new _module.C6();
          _nw48.__ctor(_306_v23);
          _330_v41 = _nw48;
          let _331_v42;
          _331_v42 = _dafny.Map.Empty.slice().updateUnsafe(_280_v1,new _dafny.CodePoint('u'.codePointAt(0)));
          _331_v42 = (_331_v42).update(_280_v1, (_313_v30).f24);
          let _332_v43;
          let _nw49 = Array((new BigNumber(3)).toNumber());
          _332_v43 = _nw49;
          _332_v43 = ((false) ? (_332_v43) : (_332_v43));
        }
        (_289_globalState).f18 = (_dafny.Map.Empty.slice().updateUnsafe(_296_v16,_280_v1)).contains(_307_v24);
        (_289_globalState).f16 = _281_v2;
        if ((((_280_v1) ? (_280_v1) : (_280_v1))) === (false)) {
          let _333_v44;
          _333_v44 = _dafny.Set.fromElements(!(_313_v30.f25).isEqualTo(_281_v2));
          let _334_v45;
          let _nw50 = Array((new BigNumber(13)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _334_v45 = _nw50;
          let _index35 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_334_v45).length));
          (_334_v45)[_index35] = _315_v32;
          let _index36 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_334_v45).length));
          let _rhs38 = (((_291_v11)[_module.__default.safeIndex(_281_v2, new BigNumber((_291_v11).length))]) ? ((_281_v2).plus(_281_v2)) : (_313_v30.f25));
          let _rhs39 = ((_dafny.Set.fromElements(_280_v1)).Difference(_333_v44)).Difference(_333_v44);
          let _rhs40 = _280_v1;
          let _rhs41 = new _dafny.CodePoint('y'.codePointAt(0));
          let _rhs42 = ((_280_v1) ? (_284_v5) : (_284_v5));
          let _lhs35 = _289_globalState;
          let _lhs36 = _289_globalState;
          let _lhs37 = _334_v45;
          let _lhs38 = _module.__default.safeIndex(new BigNumber(843), new BigNumber((_334_v45).length));
          _lhs35.f16 = _rhs38;
          _333_v44 = _rhs39;
          _lhs36.f18 = _rhs40;
          _lhs37[_lhs38] = _rhs41;
          _284_v5 = _rhs42;
          let _index37 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_284_v5).length));
          (_284_v5)[_index37] = new BigNumber((_dafny.Seq.Concat(_287_v8, _287_v8)).length);
          let _index38 = _module.__default.safeIndex(new BigNumber(553), new BigNumber((_284_v5).length));
          (_284_v5)[_index38] = _313_v30.f25;
          let _335_v46;
          _335_v46 = _dafny.MultiSet.fromElements(_279_v0);
          (_289_globalState).f16 = (_dafny.ZERO).minus(new BigNumber((_335_v46).cardinality()));
          _306_v23 = (_306_v23).update(_313_v30.f25, _280_v1);
          (_289_globalState).f18 = _280_v1;
        } else {
          let _336_v47;
          _336_v47 = _dafny.Map.Empty.slice().updateUnsafe(_286_v7,_313_v30.f25);
          let _337_v48;
          _337_v48 = _module.D3.create_DC6(_281_v2, _313_v30.f25, _280_v1);
          let _338_v49;
          _338_v49 = _module.D5.create_DC10(_280_v1, (_337_v48).dtor_cf11);
          let _339_v50;
          let _nw51 = new _module.C4();
          _nw51.__ctor((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of((((_306_v23).contains(_281_v2)) ? ((_306_v23).get(_281_v2)) : (_280_v1)), _280_v1)),_281_v2)).Merge(_336_v47), _280_v1, _module.__default.fm10(_280_v1, _280_v1, !((_338_v49).dtor_cf15), _280_v1, _289_globalState), _281_v2);
          _339_v50 = _nw51;
          let _340_v51;
          _340_v51 = _dafny.Set.fromElements(true);
          let _341_v52;
          _341_v52 = _module.D6.create_DC11(_339_v50);
          _339_v50 = ((_module.__default.fm0(new BigNumber((_340_v51).length), false, _280_v1, _281_v2, _289_globalState)) ? (_339_v50) : ((_341_v52).dtor_cf16));
          (_289_globalState).f18 = _280_v1;
          let _index39 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_284_v5).length));
          (_284_v5)[_index39] = (_281_v2).multipliedBy(_339_v50.f25);
          let _index40 = _module.__default.safeIndex(new BigNumber(375), new BigNumber((_284_v5).length));
          (_284_v5)[_index40] = (_dafny.ZERO).minus(new BigNumber((_291_v11).length));
          (_289_globalState).f18 = _280_v1;
          let _342_v53;
          let _nw52 = new _module.C0();
          _nw52.__ctor();
          _342_v53 = _nw52;
        }
      } else {
        let _index41 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_284_v5).length));
        (_284_v5)[_index41] = _281_v2;
        let _index42 = _module.__default.safeIndex(new BigNumber(814), new BigNumber((_284_v5).length));
        (_284_v5)[_index42] = _module.__default.safeModuloInt(_module.__default.safeDivisionInt(new BigNumber(161), new BigNumber((_dafny.Seq.of(_281_v2)).length)), _281_v2);
        (_289_globalState).f11 = new BigNumber((_279_v0).length);
        let _343_v54;
        let _nw53 = new _module.C2();
        _nw53.__ctor(new BigNumber((_module.__default.fm7(_289_globalState)).length), _280_v1);
        _343_v54 = _nw53;
        let _344_v55;
        _344_v55 = _dafny.Seq.of(((_280_v1) ? (_343_v54) : (_343_v54)));
        _344_v55 = _dafny.Seq.Concat(_344_v55, _dafny.Seq.Concat(_344_v55, _344_v55));
        _291_v11 = _291_v11;
        _343_v54 = _343_v54;
      }
      if (_280_v1) {
        let _345_v56;
        let _out6;
        _out6 = _module.__default.m6(new BigNumber(138), !((_281_v2).isLessThan(new BigNumber(-166))), (_280_v1) === (_280_v1), _280_v1, _289_globalState);
        _345_v56 = _out6;
        let _346_v57;
        let _nw54 = new _module.C0();
        _nw54.__ctor();
        _346_v57 = _nw54;
        let _347_v58;
        _347_v58 = _dafny.Map.Empty.slice().updateUnsafe(_346_v57,_dafny.Seq.UnicodeFromString("kmijxxnsm"));
        _347_v58 = (_347_v58).update(_346_v57, _dafny.Seq.UnicodeFromString("ogexx"));
        let _348_v59;
        let _nw55 = new _module.C5();
        _nw55.__ctor(_280_v1, _307_v24, _281_v2);
        _348_v59 = _nw55;
        _348_v59 = _348_v59;
        _284_v5 = _284_v5;
        let _349_v60;
        _349_v60 = _dafny.Set.fromElements(_284_v5);
        let _350_v61;
        _350_v61 = _module.D13.create_DC28(_349_v60);
        _350_v61 = _350_v61;
      } else {
        let _351_v62;
        _351_v62 = _dafny.Map.Empty.slice().updateUnsafe(_281_v2,_286_v7);
        let _352_v63;
        _352_v63 = _dafny.Map.Empty.slice().updateUnsafe(_351_v62,new BigNumber((_279_v0).length));
        _352_v63 = (_352_v63).update(_module.__default.fm32(_280_v1, _280_v1, _289_globalState), (_dafny.ZERO).minus((((_286_v7).contains(_280_v1)) ? ((_286_v7).get(_280_v1)) : (_281_v2))));
        _281_v2 = (_dafny.ZERO).minus(_281_v2);
        (_289_globalState).f19 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(2)), ((_353_v2) => function (_354_i4) {
          return (_353_v2).multipliedBy(_353_v2);
        })(_281_v2));
        _280_v1 = !(_280_v1);
        _280_v1 = (new BigNumber((_module.__default.fm15(new BigNumber((_282_v3).length), _281_v2, _281_v2, _289_globalState)).length)).isLessThan(_281_v2);
      }
      let _355_v64;
      _355_v64 = _module.D3.create_DC6(_281_v2, (_dafny.ZERO).minus(_281_v2), _280_v1);
      let _356_v65;
      _356_v65 = _dafny.Map.Empty.slice().updateUnsafe(_281_v2,_281_v2);
      let _357_v66;
      _357_v66 = _dafny.Map.Empty.slice().updateUnsafe(_281_v2,new BigNumber((_356_v65).length));
      let _358_v67;
      let _out7;
      _out7 = _module.__default.m6(_281_v2, _module.__default.fm0(_281_v2, _module.__default.fm0((_355_v64).dtor_cf10, _280_v1, _280_v1, _281_v2, _289_globalState), _280_v1, new BigNumber((_279_v0).length), _289_globalState), _280_v1, ((((_357_v66).contains(_module.__default.fm5(_281_v2, _module.__default.fm17(_280_v1, _280_v1, _280_v1, _289_globalState), _289_globalState))) ? ((_357_v66).get(_module.__default.fm5(_281_v2, _module.__default.fm17(_280_v1, _280_v1, _280_v1, _289_globalState), _289_globalState))) : (new BigNumber((_279_v0).length)))).isLessThanOrEqualTo(_281_v2), _289_globalState);
      _358_v67 = _out7;
      let _359_i5;
      _359_i5 = _dafny.ZERO;
      L1: {
        while (_280_v1) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_359_i5)) {
              break L1;
            }
            _359_i5 = (_359_i5).plus(_dafny.ONE);
            if ((false) || (_280_v1)) {
              (_289_globalState).f18 = _280_v1;
              let _360_v68;
              _360_v68 = _dafny.Map.Empty.slice().updateUnsafe(_286_v7,_281_v2);
              let _361_v69;
              let _nw56 = new _module.C4();
              _nw56.__ctor(_360_v68, !((_280_v1) === (_280_v1)), new _dafny.CodePoint('f'.codePointAt(0)), (_dafny.ZERO).minus(_module.__default.safeModuloInt(_281_v2, _281_v2)));
              _361_v69 = _nw56;
              let _362_v70;
              _362_v70 = _module.D12.create_DC27(_281_v2, _280_v1);
              (_289_globalState).f15 = (_281_v2).minus((_362_v70).dtor_cf40);
              let _363_v71;
              _363_v71 = _module.D5.create_DC10((_361_v69).f30, _280_v1);
              let _pat_let_tv0 = _280_v1;
              let _364_v72;
              _364_v72 = _dafny.Map.Empty.slice().updateUnsafe((_361_v69).f30,function (_pat_let0_0) {
                return function (_365_dt__update__tmp_h0) {
                  return function (_pat_let1_0) {
                    return function (_366_dt__update_hcf14_h0) {
                      return _module.D5.create_DC10(_366_dt__update_hcf14_h0, (_365_dt__update__tmp_h0).dtor_cf15);
                    }(_pat_let1_0);
                  }(_pat_let_tv0);
                }(_pat_let0_0);
              }(_363_v71));
              (_361_v69).m5(_364_v72, _280_v1, _module.__default.safeModuloInt(_281_v2, _281_v2), _289_globalState);
              _293_v13 = _293_v13;
            } else {
              let _367_v73;
              let _nw57 = new _module.C1();
              _nw57.__ctor(_296_v16, (_281_v2).minus(new BigNumber((_279_v0).length)));
              _367_v73 = _nw57;
              (_289_globalState).f10 = new BigNumber((_dafny.Seq.Concat(_279_v0, _dafny.Seq.UnicodeFromString("untmsi"))).length);
              let _368_v74;
              _368_v74 = _dafny.Map.Empty.slice().updateUnsafe(_307_v24,_280_v1);
              let _369_v75;
              let _nw58 = Array((new BigNumber(25)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
              _369_v75 = _nw58;
              let _370_v76;
              _370_v76 = _dafny.Map.Empty.slice().updateUnsafe(((_280_v1) ? (new BigNumber((_368_v74).length)) : (_281_v2)),_369_v75);
              _370_v76 = (_370_v76).update((_367_v73).fm6(_289_globalState), _369_v75);
              let _371_v77;
              let _nw59 = new _module.C3();
              _nw59.__ctor(true, _307_v24, new BigNumber((_287_v8).length));
              _371_v77 = _nw59;
              let _372_v78;
              _372_v78 = _dafny.Map.Empty.slice().updateUnsafe(_371_v77,_280_v1);
              _280_v1 = ((_280_v1) ? ((_372_v78).contains(_371_v77)) : (((((_306_v23).contains(_281_v2)) ? ((_306_v23).get(_281_v2)) : (_280_v1))) === (_module.__default.fm0(new BigNumber(763), true, _280_v1, _371_v77.f25, _289_globalState))));
              let _373_v79;
              _373_v79 = _dafny.Map.Empty.slice().updateUnsafe(_286_v7,_371_v77.f25);
              let _374_v81;
              _374_v81 = _dafny.Set.fromElements(_dafny.MultiSet.fromElements(!(_280_v1), _280_v1));
              let _375_v82;
              let _nw60 = new _module.C4();
              _nw60.__ctor((_373_v79).Merge(function () {
                let _coll31 = new _dafny.Map();
                for (const _compr_31 of (_374_v81).Elements) {
                  let _376_v80 = _compr_31;
                  if ((_374_v81).contains(_376_v80)) {
                    _coll31.push([_376_v80,_371_v77.f25]);
                  }
                }
                return _coll31;
              }()), _280_v1, _307_v24, _281_v2);
              _375_v82 = _nw60;
            }
            let _377_v83;
            _377_v83 = _dafny.Set.fromElements(_284_v5, _284_v5, _284_v5);
            let _378_v84;
            _378_v84 = _dafny.Map.Empty.slice().updateUnsafe(((_291_v11)[_module.__default.safeIndex(new BigNumber((_279_v0).length), new BigNumber((_291_v11).length))]) || (_280_v1),_377_v83);
            _378_v84 = (_378_v84).update(_280_v1, ((_280_v1) ? (_377_v83) : (_377_v83)));
            let _index43 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_283_v4).length));
            (_283_v4)[_index43] = _280_v1;
            let _379_v85;
            let _init9 = ((_380_v0) => function (_381_i6) {
              return _380_v0;
            })(_279_v0);
            let _nw61 = Array((new BigNumber(6)).toNumber());
            for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw61.length); _i0_9++) {
              _nw61[_i0_9] = _init9(new BigNumber(_i0_9));
            }
            _379_v85 = _nw61;
            let _index44 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_379_v85).length));
            (_379_v85)[_index44] = _279_v0;
            let _382_v86;
            _382_v86 = _dafny.MultiSet.fromElements(_358_v67);
            let _index45 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_283_v4).length));
            let _index46 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_379_v85).length));
            let _rhs43 = _module.__default.fm5(_281_v2, new BigNumber(((_382_v86).Difference(_dafny.MultiSet.fromElements(new _dafny.CodePoint('f'.codePointAt(0))))).cardinality()), _289_globalState);
            let _rhs44 = _280_v1;
            let _rhs45 = _dafny.Seq.Concat(_279_v0, _dafny.Seq.UnicodeFromString("wdfjarp"));
            let _lhs39 = _289_globalState;
            let _lhs40 = _283_v4;
            let _lhs41 = _module.__default.safeIndex(new BigNumber(816), new BigNumber((_283_v4).length));
            let _lhs42 = _379_v85;
            let _lhs43 = _module.__default.safeIndex(new BigNumber(206), new BigNumber((_379_v85).length));
            _lhs39.f2 = _rhs43;
            _lhs40[_lhs41] = _rhs44;
            _lhs42[_lhs43] = _rhs45;
            let _index47 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_284_v5).length));
            (_284_v5)[_index47] = (_287_v8)[_module.__default.safeIndex(_281_v2, new BigNumber((_287_v8).length))];
            let _index48 = _module.__default.safeIndex(new BigNumber(950), new BigNumber((_284_v5).length));
            (_284_v5)[_index48] = (((_290_v10).contains(_281_v2)) ? ((_290_v10).get(_281_v2)) : (_281_v2));
          }
        }
      }
      _281_v2 = _281_v2;
      (_289_globalState).f15 = ((_280_v1) ? (_281_v2) : (_281_v2));
      let _383_v87;
      _383_v87 = _dafny.Map.Empty.slice().updateUnsafe(false,_284_v5);
      let _384_v88;
      let _nw62 = Array((new BigNumber(11)).toNumber());
      _nw62[(_dafny.ZERO).toNumber()] = _284_v5;
      _nw62[(_dafny.ONE).toNumber()] = _284_v5;
      _nw62[(new BigNumber(2)).toNumber()] = _284_v5;
      _nw62[(new BigNumber(3)).toNumber()] = _284_v5;
      _nw62[(new BigNumber(4)).toNumber()] = (((_383_v87).contains(_280_v1)) ? ((_383_v87).get(_280_v1)) : (_284_v5));
      _nw62[(new BigNumber(5)).toNumber()] = _284_v5;
      _nw62[(new BigNumber(6)).toNumber()] = _284_v5;
      _nw62[(new BigNumber(7)).toNumber()] = _284_v5;
      _nw62[(new BigNumber(8)).toNumber()] = _284_v5;
      _nw62[(new BigNumber(9)).toNumber()] = _284_v5;
      _nw62[(new BigNumber(10)).toNumber()] = _284_v5;
      _384_v88 = _nw62;
      _384_v88 = _384_v88;
      let _385_v89;
      let _init10 = function (_386_i8) {
        return _dafny.Seq.Create(_module.__default.abs(new BigNumber(903)), function (_387_i9) {
          return new _dafny.CodePoint('y'.codePointAt(0));
        });
      };
      let _nw63 = Array((new BigNumber(14)).toNumber());
      for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw63.length); _i0_10++) {
        _nw63[_i0_10] = _init10(new BigNumber(_i0_10));
      }
      _385_v89 = _nw63;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_385_v89).length))) {
        let _388_i7 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_388_i7)) && ((_388_i7).isLessThan(new BigNumber((_385_v89).length))))) {
          (_385_v89)[(_388_i7)] = _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sn"), _279_v0), _dafny.Seq.UnicodeFromString("infvwkvh"));
        }
      }
      let _389_v90;
      _389_v90 = _dafny.Map.Empty.slice().updateUnsafe((_291_v11)[_module.__default.safeIndex(_module.__default.fm17(!(_280_v1), _280_v1, _280_v1, _289_globalState), new BigNumber((_291_v11).length))],_280_v1);
      if (!((((_389_v90).contains(_280_v1)) ? ((_389_v90).get(_280_v1)) : (((_dafny.ZERO).minus(_281_v2)).isLessThan(_281_v2))))) {
        let _390_v91;
        let _nw64 = new _module.C6();
        _nw64.__ctor(((_280_v1) ? ((_306_v23).update(_281_v2, false)) : (_306_v23)));
        _390_v91 = _nw64;
        let _391_v92;
        let _nw65 = new _module.C3();
        _nw65.__ctor(_280_v1, _358_v67, new BigNumber(224));
        _391_v92 = _nw65;
        let _392_v93;
        let _nw66 = new _module.C1();
        _nw66.__ctor(_358_v67, (_287_v8)[_module.__default.safeIndex(_281_v2, new BigNumber((_287_v8).length))]);
        _392_v93 = _nw66;
        _392_v93 = ((!(_391_v92.f31)) ? (_392_v93) : (_392_v93));
        let _index49 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_284_v5).length));
        (_284_v5)[_index49] = _281_v2;
        let _index50 = _module.__default.safeIndex(new BigNumber(919), new BigNumber((_284_v5).length));
        (_284_v5)[_index50] = (((_356_v65).contains(_281_v2)) ? ((_356_v65).get(_281_v2)) : (_module.__default.safeDivisionInt(_281_v2, new BigNumber((_291_v11).length))));
        (_289_globalState).f15 = (((_284_v5)[_module.__default.safeIndex(new BigNumber(919), new BigNumber((_284_v5).length))]).minus(_281_v2)).minus(_281_v2);
      } else {
        let _393_v94;
        _393_v94 = _module.D7.create_DC13(_306_v23);
        let _394_v95;
        let _nw67 = new _module.C1();
        _nw67.__ctor(_307_v24, (new BigNumber(((_393_v94).dtor_cf19).length)).minus(_281_v2));
        _394_v95 = _nw67;
        let _index51 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_284_v5).length));
        (_284_v5)[_index51] = _281_v2;
        let _index52 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_385_v89).length));
        (_385_v89)[_index52] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("yj"), _279_v0);
        let _395_v96;
        _395_v96 = _module.D10.create_DC21((((_282_v3).contains(_280_v1)) ? ((_282_v3).get(_280_v1)) : (_281_v2)), _280_v1, !(false));
        let _396_v97;
        _396_v97 = _dafny.MultiSet.fromElements(_395_v96, _395_v96, _395_v96, _395_v96, _module.D10.create_DC21(_281_v2, !(_module.__default.fm0(new BigNumber((_dafny.MultiSet.FromArray(_291_v11)).cardinality()), _280_v1, _280_v1, new BigNumber((_279_v0).length), _289_globalState)), _280_v1));
        let _pat_let_tv1 = _280_v1;
        let _index53 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_284_v5).length));
        let _index54 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_385_v89).length));
        let _rhs46 = _281_v2;
        let _rhs47 = _279_v0;
        let _rhs48 = (_305_v22).Union(_305_v22);
        let _rhs49 = function (_pat_let2_0) {
          return function (_397_dt__update__tmp_h1) {
            return function (_pat_let3_0) {
              return function (_398_dt__update_hcf32_h0) {
                return _module.D10.create_DC21((_397_dt__update__tmp_h1).dtor_cf30, (_397_dt__update__tmp_h1).dtor_cf31, _398_dt__update_hcf32_h0);
              }(_pat_let3_0);
            }(_pat_let_tv1);
          }(_pat_let2_0);
        }(_395_v96);
        let _rhs50 = (((_280_v1) ? (_396_v97) : (_396_v97))).Intersect((_dafny.MultiSet.fromElements(_395_v96)).Union(_396_v97));
        let _lhs44 = _284_v5;
        let _lhs45 = _module.__default.safeIndex(new BigNumber(706), new BigNumber((_284_v5).length));
        let _lhs46 = _385_v89;
        let _lhs47 = _module.__default.safeIndex(new BigNumber(617), new BigNumber((_385_v89).length));
        _lhs44[_lhs45] = _rhs46;
        _lhs46[_lhs47] = _rhs47;
        _305_v22 = _rhs48;
        _395_v96 = _rhs49;
        _396_v97 = _rhs50;
        let _399_v98;
        _399_v98 = _dafny.Map.Empty.slice().updateUnsafe(_286_v7,new BigNumber(683));
        let _400_v99;
        let _nw68 = new _module.C4();
        _nw68.__ctor(_399_v98, _280_v1, _296_v16, new BigNumber(350));
        _400_v99 = _nw68;
        let _401_v100;
        _401_v100 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(482),_400_v99);
        _401_v100 = (_401_v100).update(((_280_v1) ? (new BigNumber((_282_v3).length)) : (_281_v2)), _400_v99);
        let _402_v101;
        _402_v101 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update(_dafny.Seq.UnicodeFromString("b"), _module.__default.safeIndex(_281_v2, new BigNumber((_dafny.Seq.UnicodeFromString("b")).length)), _296_v16),_281_v2);
        let _403_v102;
        _403_v102 = _dafny.Map.Empty.slice().updateUnsafe((_394_v95).fm6(_289_globalState),_402_v101);
        _403_v102 = (_403_v102).update(_400_v99.f25, _module.__default.fm33(_289_globalState));
        _279_v0 = (_385_v89)[_module.__default.safeIndex(new BigNumber(617), new BigNumber((_385_v89).length))];
      }
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_284_v5).length))) {
        let _404_i10 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_404_i10)) && ((_404_i10).isLessThan(new BigNumber((_284_v5).length))))) {
          (_284_v5)[(_404_i10)] = _module.__default.safeModuloInt(_404_i10, _281_v2);
        }
      }
      process.stdout.write(_dafny.toString(_dafny.areEqual(_279_v0, _dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_280_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_281_v2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_282_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(742)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_283_v4)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_283_v4)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_284_v5)[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_285_v6).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_286_v7).equals(_dafny.MultiSet.fromElements(true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_287_v8, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_288_v9).equals(_dafny.Set.fromElements(_dafny.Seq.of(_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_289_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_globalState).f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_289_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_289_globalState.f3, _dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)), new _dafny.CodePoint('w'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_globalState).f4));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_globalState).f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_globalState).f6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_289_globalState).f7).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(742)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber(((_289_globalState).f8).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_globalState).f9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_289_globalState.f10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_289_globalState.f11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_globalState).f12));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_globalState.f13).equals(_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(742)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_289_globalState).f14).equals(_dafny.Set.fromElements(_dafny.Seq.of(_dafny.ONE)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_289_globalState.f15));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_289_globalState.f16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_globalState).f17));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_289_globalState.f18));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_289_globalState.f19, _dafny.Seq.of(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_globalState).f20));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_289_globalState).f21));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_289_globalState.f22));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_290_v10).equals(_dafny.MultiSet.fromElements(new BigNumber(742), new BigNumber(742)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_291_v11, _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_292_v12), _dafny.Seq.of(true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_293_v13).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_294_v14), _dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual(_295_v15, _dafny.Seq.of(_dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0))), _dafny.Seq.UnicodeFromString("fln"), _dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0)))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_296_v16));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_297_i0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_305_v22).equals(_dafny.Set.fromElements(new BigNumber(742), new BigNumber(2)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v23).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_307_v24));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_308_v25).dtor_cf34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_308_v25).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_309_v26).dtor_cf38).dtor_cf34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_309_v26).dtor_cf38).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_310_v27)[_dafny.ZERO]).dtor_cf38).dtor_cf34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_310_v27)[_dafny.ZERO]).dtor_cf38).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_310_v27)[_dafny.ONE]).dtor_cf38).dtor_cf34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_310_v27)[_dafny.ONE]).dtor_cf38).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_310_v27)[new BigNumber(2)]).dtor_cf38).dtor_cf34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_310_v27)[new BigNumber(2)]).dtor_cf38).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_310_v27)[new BigNumber(3)]).dtor_cf38).dtor_cf34));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((((_310_v27)[new BigNumber(3)]).dtor_cf38).dtor_cf35));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_311_v28).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_312_v29).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_355_v64).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_355_v64).dtor_cf10));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_355_v64).dtor_cf11));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_356_v65).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(742),new BigNumber(742)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_357_v66).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(742),_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_358_v67));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_359_i5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_383_v87).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ZERO])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[_dafny.ONE])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(2)])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(3)])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(4)])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(5)])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(6)])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(7)])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(8)])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(9)])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(12)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(13)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(14)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(15)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(16)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(17)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(18)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(19)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(20)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(21)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(22)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(23)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_384_v88)[new BigNumber(10)])[new BigNumber(24)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[_dafny.ZERO]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_dafny.areEqual((_385_v89)[_dafny.ONE], _dafny.Seq.of(new _dafny.CodePoint('w'.codePointAt(0))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(2)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(3)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(4)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(5)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(6)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(7)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(8)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(9)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(10)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(11)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(12)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_385_v89)[new BigNumber(13)]).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_389_v90).equals(_dafny.Map.Empty.slice().updateUnsafe(true,true))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      return;
    }
  };

  $module.D0 = class D0 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC0(cf0) {
      let $dt = new D0(0);
      $dt.cf0 = cf0;
      return $dt;
    }
    get is_DC0() { return this.$tag === 0; }
    get dtor_cf0() { return this.cf0; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf0, other.cf0);
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
          return D0.Default();
        }
      };
    }
  }

  $module.D1 = class D1 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC1(cf1) {
      let $dt = new D1(0);
      $dt.cf1 = cf1;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get dtor_cf1() { return this.cf1; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC1" + "(" + this.cf1.toVerbatimString(true) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _dafny.Seq.UnicodeFromString("");
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
    static create_DC3(cf3, cf4) {
      let $dt = new D2(0);
      $dt.cf3 = cf3;
      $dt.cf4 = cf4;
      return $dt;
    }
    static create_DC2(cf2) {
      let $dt = new D2(1);
      $dt.cf2 = cf2;
      return $dt;
    }
    get is_DC3() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf2() { return this.cf2; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC3" + "(" + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D2.DC2" + "(" + _dafny.toString(this.cf2) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf3, other.cf3) && _dafny.areEqual(this.cf4, other.cf4);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf2 === other.cf2;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D2.create_DC3(_dafny.ZERO, _dafny.ZERO);
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
    static create_DC5(cf6, cf7, cf8) {
      let $dt = new D3(0);
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      return $dt;
    }
    static create_DC6(cf9, cf10, cf11) {
      let $dt = new D3(1);
      $dt.cf9 = cf9;
      $dt.cf10 = cf10;
      $dt.cf11 = cf11;
      return $dt;
    }
    static create_DC4(cf5) {
      let $dt = new D3(2);
      $dt.cf5 = cf5;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC6() { return this.$tag === 1; }
    get is_DC4() { return this.$tag === 2; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf10() { return this.cf10; }
    get dtor_cf11() { return this.cf11; }
    get dtor_cf5() { return this.cf5; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC5" + "(" + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC6" + "(" + _dafny.toString(this.cf9) + ", " + _dafny.toString(this.cf10) + ", " + _dafny.toString(this.cf11) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC4" + "(" + _dafny.toString(this.cf5) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf9, other.cf9) && _dafny.areEqual(this.cf10, other.cf10) && this.cf11 === other.cf11;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf5, other.cf5);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC5(_dafny.Seq.of(), _dafny.ZERO, _dafny.ZERO);
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
    static create_DC8() {
      let $dt = new D4(0);
      return $dt;
    }
    static create_DC7(cf12) {
      let $dt = new D4(1);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC8";
      } else if (this.$tag === 1) {
        return "D4.DC7" + "(" + _dafny.toString(this.cf12) + ")";
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
        return other.$tag === 1 && this.cf12 === other.cf12;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D4.create_DC8();
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
    static create_DC10(cf14, cf15) {
      let $dt = new D5(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      return $dt;
    }
    static create_DC9(cf13) {
      let $dt = new D5(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get is_DC9() { return this.$tag === 1; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf13() { return this.cf13; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC10" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC9" + "(" + _dafny.toString(this.cf13) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf14 === other.cf14 && this.cf15 === other.cf15;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf13 === other.cf13;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC10(false, false);
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
    static create_DC12(cf17, cf18) {
      let $dt = new D6(0);
      $dt.cf17 = cf17;
      $dt.cf18 = cf18;
      return $dt;
    }
    static create_DC11(cf16) {
      let $dt = new D6(1);
      $dt.cf16 = cf16;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC11() { return this.$tag === 1; }
    get dtor_cf17() { return this.cf17; }
    get dtor_cf18() { return this.cf18; }
    get dtor_cf16() { return this.cf16; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC12" + "(" + _dafny.toString(this.cf17) + ", " + _dafny.toString(this.cf18) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC11" + "(" + _dafny.toString(this.cf16) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf17 === other.cf17 && _dafny.areEqual(this.cf18, other.cf18);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf16 === other.cf16;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC12(false, _dafny.ZERO);
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
    static create_DC14(cf20, cf21, cf22) {
      let $dt = new D7(0);
      $dt.cf20 = cf20;
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      return $dt;
    }
    static create_DC13(cf19) {
      let $dt = new D7(1);
      $dt.cf19 = cf19;
      return $dt;
    }
    get is_DC14() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf19() { return this.cf19; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC14" + "(" + _dafny.toString(this.cf20) + ", " + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC13" + "(" + _dafny.toString(this.cf19) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf20 === other.cf20 && _dafny.areEqual(this.cf21, other.cf21) && _dafny.areEqual(this.cf22, other.cf22);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf19, other.cf19);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC14(false, _dafny.ZERO, _dafny.ZERO);
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
    static create_DC15(cf23) {
      let $dt = new D8(0);
      $dt.cf23 = cf23;
      return $dt;
    }
    get is_DC15() { return this.$tag === 0; }
    get dtor_cf23() { return this.cf23; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC15" + "(" + _dafny.toString(this.cf23) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf23, other.cf23);
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
          return D8.Default();
        }
      };
    }
  }

  $module.D9 = class D9 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC17(cf25) {
      let $dt = new D9(0);
      $dt.cf25 = cf25;
      return $dt;
    }
    static create_DC16(cf24) {
      let $dt = new D9(1);
      $dt.cf24 = cf24;
      return $dt;
    }
    get is_DC17() { return this.$tag === 0; }
    get is_DC16() { return this.$tag === 1; }
    get dtor_cf25() { return this.cf25; }
    get dtor_cf24() { return this.cf24; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC17" + "(" + _dafny.toString(this.cf25) + ")";
      } else if (this.$tag === 1) {
        return "D9.DC16" + "(" + _dafny.toString(this.cf24) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf25, other.cf25);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf24, other.cf24);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D9.create_DC17(_dafny.ZERO);
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
    static create_DC19(cf27) {
      let $dt = new D10(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC20(cf28, cf29) {
      let $dt = new D10(1);
      $dt.cf28 = cf28;
      $dt.cf29 = cf29;
      return $dt;
    }
    static create_DC21(cf30, cf31, cf32) {
      let $dt = new D10(2);
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC18(cf26) {
      let $dt = new D10(3);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC19() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get is_DC21() { return this.$tag === 2; }
    get is_DC18() { return this.$tag === 3; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf28() { return this.cf28; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC19" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D10.DC20" + "(" + _dafny.toString(this.cf28) + ", " + _dafny.toString(this.cf29) + ")";
      } else if (this.$tag === 2) {
        return "D10.DC21" + "(" + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 3) {
        return "D10.DC18" + "(" + _dafny.toString(this.cf26) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf27, other.cf27);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf28 === other.cf28 && _dafny.areEqual(this.cf29, other.cf29);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf30, other.cf30) && this.cf31 === other.cf31 && this.cf32 === other.cf32;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D10.create_DC19(new _dafny.CodePoint('D'.codePointAt(0)));
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
    static create_DC23(cf34, cf35) {
      let $dt = new D11(0);
      $dt.cf34 = cf34;
      $dt.cf35 = cf35;
      return $dt;
    }
    static create_DC24(cf36, cf37) {
      let $dt = new D11(1);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      return $dt;
    }
    static create_DC22(cf33) {
      let $dt = new D11(2);
      $dt.cf33 = cf33;
      return $dt;
    }
    static create_DC25(cf38) {
      let $dt = new D11(3);
      $dt.cf38 = cf38;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get is_DC24() { return this.$tag === 1; }
    get is_DC22() { return this.$tag === 2; }
    get is_DC25() { return this.$tag === 3; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf35() { return this.cf35; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf38() { return this.cf38; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC23" + "(" + _dafny.toString(this.cf34) + ", " + _dafny.toString(this.cf35) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC22" + "(" + _dafny.toString(this.cf33) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf38) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf34, other.cf34) && this.cf35 === other.cf35;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf36 === other.cf36 && _dafny.areEqual(this.cf37, other.cf37);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && this.cf33 === other.cf33;
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf38, other.cf38);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC23(_dafny.ZERO, false);
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
    static create_DC27(cf40, cf41) {
      let $dt = new D12(0);
      $dt.cf40 = cf40;
      $dt.cf41 = cf41;
      return $dt;
    }
    static create_DC26(cf39) {
      let $dt = new D12(1);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC27() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get dtor_cf40() { return this.cf40; }
    get dtor_cf41() { return this.cf41; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC27" + "(" + _dafny.toString(this.cf40) + ", " + _dafny.toString(this.cf41) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC26" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf40, other.cf40) && this.cf41 === other.cf41;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf39, other.cf39);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC27(_dafny.ZERO, false);
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
    static create_DC29(cf43, cf44, cf45) {
      let $dt = new D13(0);
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      $dt.cf45 = cf45;
      return $dt;
    }
    static create_DC28(cf42) {
      let $dt = new D13(1);
      $dt.cf42 = cf42;
      return $dt;
    }
    static create_DC30(cf46) {
      let $dt = new D13(2);
      $dt.cf46 = cf46;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf46() { return this.cf46; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC29" + "(" + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ", " + _dafny.toString(this.cf45) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC28" + "(" + _dafny.toString(this.cf42) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC30" + "(" + _dafny.toString(this.cf46) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf43 === other.cf43 && this.cf44 === other.cf44 && this.cf45 === other.cf45;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf42, other.cf42);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf46, other.cf46);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC29(false, false, []);
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
    static create_DC31(cf47) {
      let $dt = new D14(0);
      $dt.cf47 = cf47;
      return $dt;
    }
    get is_DC31() { return this.$tag === 0; }
    get dtor_cf47() { return this.cf47; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC31" + "(" + _dafny.toString(this.cf47) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf47, other.cf47);
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
    static create_DC32(cf48) {
      let $dt = new D15(0);
      $dt.cf48 = cf48;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get dtor_cf48() { return this.cf48; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC32" + "(" + _dafny.toString(this.cf48) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf48 === other.cf48;
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
          return D15.Default();
        }
      };
    }
  }

  $module.D16 = class D16 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC34(cf50) {
      let $dt = new D16(0);
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC33(cf49) {
      let $dt = new D16(1);
      $dt.cf49 = cf49;
      return $dt;
    }
    static create_DC35(cf51) {
      let $dt = new D16(2);
      $dt.cf51 = cf51;
      return $dt;
    }
    get is_DC34() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC35() { return this.$tag === 2; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf51() { return this.cf51; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC34" + "(" + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC33" + "(" + _dafny.toString(this.cf49) + ")";
      } else if (this.$tag === 2) {
        return "D16.DC35" + "(" + _dafny.toString(this.cf51) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf50 === other.cf50;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf49, other.cf49);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf51, other.cf51);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC34(false);
    }
    static Rtd() {
      return class {
        static get Default() {
          return D16.Default();
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
      this.f0 = _dafny.ZERO;
      this.f2 = _dafny.ZERO;
      this.f3 = _dafny.Seq.UnicodeFromString("");
      this.f10 = _dafny.ZERO;
      this.f11 = _dafny.ZERO;
      this.f13 = _dafny.Map.Empty;
      this.f15 = _dafny.ZERO;
      this.f16 = _dafny.ZERO;
      this.f18 = false;
      this.f19 = _dafny.Seq.of();
      this.f22 = _dafny.ZERO;
      this._f1 = _dafny.ZERO;
      this._f4 = _dafny.ZERO;
      this._f5 = false;
      this._f6 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f7 = _dafny.Map.Empty;
      this._f8 = _dafny.Map.Empty;
      this._f9 = _dafny.ZERO;
      this._f12 = false;
      this._f14 = _dafny.Set.Empty;
      this._f17 = false;
      this._f20 = _dafny.ZERO;
      this._f21 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22) {
      let _this = this;
      (_this).f0 = f0;
      (_this)._f1 = f1;
      (_this).f2 = f2;
      (_this).f3 = f3;
      (_this)._f4 = f4;
      (_this)._f5 = f5;
      (_this)._f6 = f6;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      (_this)._f9 = f9;
      (_this).f10 = f10;
      (_this).f11 = f11;
      (_this)._f12 = f12;
      (_this).f13 = f13;
      (_this)._f14 = f14;
      (_this).f15 = f15;
      (_this).f16 = f16;
      (_this)._f17 = f17;
      (_this).f18 = f18;
      (_this).f19 = f19;
      (_this)._f20 = f20;
      (_this)._f21 = f21;
      (_this).f22 = f22;
      return;
    }
    get f1() {
      let _this = this;
      return _this._f1;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
    get f5() {
      let _this = this;
      return _this._f5;
    };
    get f6() {
      let _this = this;
      return _this._f6;
    };
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f12() {
      let _this = this;
      return _this._f12;
    };
    get f14() {
      let _this = this;
      return _this._f14;
    };
    get f17() {
      let _this = this;
      return _this._f17;
    };
    get f20() {
      let _this = this;
      return _this._f20;
    };
    get f21() {
      let _this = this;
      return _this._f21;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
    }
    _parentTraits() {
      return [];
    }
    __ctor() {
      let _this = this;
      return;
    }
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
      this._f25 = _dafny.ZERO;
      this._f24 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    set f25(value) {
      let _this = this;
      _this._f25 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    __ctor(f24, f25) {
      let _this = this;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ywk"), _dafny.Seq.UnicodeFromString("nyrhpcet"));
    };
    fm6(globalState) {
      let _this = this;
      return _module.__default.safeModuloInt(_module.__default.safeModuloInt(_this.f25, _this.f25), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(false, true),new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-385)), function (_405_i0) {
        return _this.f25;
      })).length))).length))).length));
    };
    m4(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = _dafny.ZERO;
      let _406_v0;
      _406_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
      let _407_v1;
      let _init11 = ((_408_p0) => function (_409_i0) {
        return _408_p0;
      })(p0);
      let _nw69 = Array((new BigNumber(9)).toNumber());
      for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw69.length); _i0_11++) {
        _nw69[_i0_11] = _init11(new BigNumber(_i0_11));
      }
      _407_v1 = _nw69;
      let _410_v2;
      _410_v2 = _dafny.Set.fromElements(_407_v1);
      _406_v0 = (_406_v0).update((_410_v2).IsSubsetOf(_410_v2), p0);
      (_this).f25 = p1;
      let _hi1 = (_this).fm6(globalState);
      for (let _411_i1 = p1; _411_i1.isLessThan(_hi1); _411_i1 = _411_i1.plus(_dafny.ONE)) {
        let _412_v3;
        let _nw70 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
        _412_v3 = _nw70;
        _412_v3 = _412_v3;
        let _413_v4;
        _413_v4 = _module.D5.create_DC10(p0, !(p0));
        let _source7 = _413_v4;
        if (_source7.is_DC10) {
          let _414___mcc_h0 = (_source7).cf14;
          let _415___mcc_h1 = (_source7).cf15;
          let _416_cf15 = _415___mcc_h1;
          let _417_cf14 = _414___mcc_h0;
          let _418_v5;
          let _nw71 = Array((new BigNumber(29)).toNumber()).fill(_dafny.ZERO);
          _418_v5 = _nw71;
          let _index55 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_418_v5).length));
          (_418_v5)[_index55] = new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(_416_cf15))).cardinality());
          let _index56 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_407_v1).length));
          (_407_v1)[_index56] = _417_cf14;
          let _index57 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_418_v5).length));
          (_418_v5)[_index57] = _411_i1;
          let _419_v6;
          _419_v6 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,p0);
          let _420_v7;
          _420_v7 = _dafny.Seq.UnicodeFromString("tpjon");
          let _421_v8;
          _421_v8 = _dafny.Seq.of(_416_cf15, p0);
          let _422_v10;
          _422_v10 = _dafny.Seq.of(_this.f25);
          let _423_v11;
          _423_v11 = _dafny.Set.fromElements(p0);
          let _424_v12;
          _424_v12 = _module.D7.create_DC13(_419_v6);
          let _425_v14;
          let _nw72 = Array((new BigNumber(28)).toNumber());
          _nw72[(_dafny.ZERO).toNumber()] = _419_v6;
          _nw72[(_dafny.ONE).toNumber()] = _419_v6;
          _nw72[(new BigNumber(2)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(3)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_420_v7).length),(_421_v8)[_module.__default.safeIndex(p1, new BigNumber((_421_v8).length))]);
          _nw72[(new BigNumber(4)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(5)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(6)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(7)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(8)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(9)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(10)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(11)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(12)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(13)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_420_v7).length),_417_cf14);
          _nw72[(new BigNumber(14)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(15)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(16)).toNumber()] = (_419_v6).update((_this).fm6(globalState), _416_cf15);
          _nw72[(new BigNumber(17)).toNumber()] = function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of (_422_v10).Elements) {
              let _426_v9 = _compr_32;
              if (_dafny.Seq.contains(_422_v10, _426_v9)) {
                _coll32.push([_module.__default.safeModuloInt(_426_v9, new BigNumber(560)),_417_cf14]);
              }
            }
            return _coll32;
          }();
          _nw72[(new BigNumber(18)).toNumber()] = (_419_v6).update(new BigNumber((_423_v11).length), _416_cf15);
          _nw72[(new BigNumber(19)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(20)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(21)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(22)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(23)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(24)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_417_cf14);
          _nw72[(new BigNumber(25)).toNumber()] = (_424_v12).dtor_cf19;
          _nw72[(new BigNumber(26)).toNumber()] = _419_v6;
          _nw72[(new BigNumber(27)).toNumber()] = function () {
            let _coll33 = new _dafny.Map();
            for (const _compr_33 of (_422_v10).Elements) {
              let _427_v13 = _compr_33;
              if (_dafny.Seq.contains(_422_v10, _427_v13)) {
                _coll33.push([(_427_v13).minus(_411_i1),_417_cf14]);
              }
            }
            return _coll33;
          }();
          _425_v14 = _nw72;
          let _428_v15;
          _428_v15 = _module.D4.create_DC7(_425_v14);
          let _429_v16;
          _429_v16 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Seq.of(_411_i1));
          let _index58 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_418_v5).length));
          let _index59 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_407_v1).length));
          let _index60 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_418_v5).length));
          let _rhs51 = (_411_i1).minus((_411_i1).minus(_411_i1));
          let _rhs52 = new BigNumber((_dafny.Seq.UnicodeFromString("duj")).length);
          let _rhs53 = _417_cf14;
          let _rhs54 = (_dafny.ZERO).minus((_411_i1).multipliedBy(_module.__default.safeModuloInt(new BigNumber((_429_v16).length), _this.f25)));
          let _rhs55 = _428_v15;
          let _lhs48 = _418_v5;
          let _lhs49 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_418_v5).length));
          let _lhs50 = globalState;
          let _lhs51 = _407_v1;
          let _lhs52 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_407_v1).length));
          let _lhs53 = _418_v5;
          let _lhs54 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_418_v5).length));
          _lhs48[_lhs49] = _rhs51;
          _lhs50.f22 = _rhs52;
          _lhs51[_lhs52] = _rhs53;
          _lhs53[_lhs54] = _rhs54;
          _428_v15 = _rhs55;
          let _430_v17;
          let _nw73 = Array((new BigNumber(19)).toNumber()).fill([]);
          _430_v17 = _nw73;
          _430_v17 = _430_v17;
          let _index61 = _module.__default.safeIndex(new BigNumber(929), new BigNumber((_418_v5).length));
          (_418_v5)[_index61] = (_module.__default.safeDivisionInt((_418_v5)[_module.__default.safeIndex(new BigNumber(929), new BigNumber((_418_v5).length))], p1)).multipliedBy(_411_i1);
          _418_v5 = ((((_416_cf15) ? ((_407_v1)[_module.__default.safeIndex(new BigNumber(844), new BigNumber((_407_v1).length))]) : (true))) ? (_418_v5) : (_418_v5));
        } else {
          let _431___mcc_h2 = (_source7).cf13;
          let _432_cf13 = _431___mcc_h2;
          let _index62 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_432_cf13).length));
          (_432_cf13)[_index62] = _411_i1;
          let _index63 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_432_cf13).length));
          (_432_cf13)[_index63] = (_dafny.ZERO).minus((_411_i1).minus(_411_i1));
          (globalState).f2 = (_432_cf13)[_module.__default.safeIndex(new BigNumber(487), new BigNumber((_432_cf13).length))];
          let _433_v18;
          _433_v18 = _dafny.Map.Empty.slice().updateUnsafe(p0,_432_cf13);
          let _434_v19;
          _434_v19 = _dafny.MultiSet.fromElements(p0);
          let _435_v20;
          _435_v20 = _dafny.Map.Empty.slice().updateUnsafe(_411_i1,true);
          let _index64 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_432_cf13).length));
          let _rhs56 = (((_434_v19).contains(!(true) || (_module.__default.fm0(_411_i1, p0, p0, _this.f25, globalState)))) ? ((_434_v19).get(!(true) || (_module.__default.fm0(_411_i1, p0, p0, _this.f25, globalState)))) : (new BigNumber((_435_v20).length)));
          let _rhs57 = (_433_v18).Merge(_433_v18);
          let _rhs58 = _411_i1;
          let _rhs59 = p0;
          let _rhs60 = _432_cf13;
          let _lhs55 = globalState;
          let _lhs56 = _432_cf13;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(487), new BigNumber((_432_cf13).length));
          let _lhs58 = globalState;
          _lhs55.f15 = _rhs56;
          _433_v18 = _rhs57;
          _lhs56[_lhs57] = _rhs58;
          _lhs58.f18 = _rhs59;
          _432_cf13 = _rhs60;
          let _436_v21;
          let _nw74 = new _module.C0();
          _nw74.__ctor();
          _436_v21 = _nw74;
        }
        let _437_v22;
        let _nw75 = Array((new BigNumber(21)).toNumber());
        _nw75[(_dafny.ZERO).toNumber()] = ((p0) ? (_407_v1) : (_407_v1));
        _nw75[(_dafny.ONE).toNumber()] = _407_v1;
        _nw75[(new BigNumber(2)).toNumber()] = ((p0) ? (_407_v1) : (_407_v1));
        _nw75[(new BigNumber(3)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(4)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(5)).toNumber()] = ((p0) ? (_407_v1) : (_407_v1));
        _nw75[(new BigNumber(6)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(7)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(8)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(9)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(10)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(11)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(12)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(13)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(14)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(15)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(16)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(17)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(18)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(19)).toNumber()] = _407_v1;
        _nw75[(new BigNumber(20)).toNumber()] = _407_v1;
        _437_v22 = _nw75;
        let _438_v23;
        _438_v23 = _dafny.Seq.of(_module.D5.create_DC10(p0, false));
        let _439_v24;
        _439_v24 = _dafny.Seq.of(!(p0));
        let _440_v25;
        _440_v25 = _dafny.Seq.UnicodeFromString("nnhp");
        let _441_v26;
        _441_v26 = _dafny.Seq.of(new BigNumber((_440_v25).length));
        let _442_v27;
        let _nw76 = Array((new BigNumber(14)).toNumber());
        _nw76[(_dafny.ZERO).toNumber()] = (_411_i1).minus(_this.f25);
        _nw76[(_dafny.ONE).toNumber()] = _411_i1;
        _nw76[(new BigNumber(2)).toNumber()] = _this.f25;
        _nw76[(new BigNumber(3)).toNumber()] = ((p0) ? (new BigNumber((_438_v23).length)) : (_this.f25));
        _nw76[(new BigNumber(4)).toNumber()] = _411_i1;
        _nw76[(new BigNumber(5)).toNumber()] = _this.f25;
        _nw76[(new BigNumber(6)).toNumber()] = new BigNumber((_439_v24).length);
        _nw76[(new BigNumber(7)).toNumber()] = _411_i1;
        _nw76[(new BigNumber(8)).toNumber()] = p1;
        _nw76[(new BigNumber(9)).toNumber()] = _411_i1;
        _nw76[(new BigNumber(10)).toNumber()] = (_411_i1).minus(_411_i1);
        _nw76[(new BigNumber(11)).toNumber()] = p1;
        _nw76[(new BigNumber(12)).toNumber()] = _this.f25;
        _nw76[(new BigNumber(13)).toNumber()] = (_441_v26)[_module.__default.safeIndex(p1, new BigNumber((_441_v26).length))];
        _442_v27 = _nw76;
        let _index65 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_442_v27).length));
        (_442_v27)[_index65] = _this.f25;
        let _443_v28;
        _443_v28 = _dafny.Set.fromElements(p0, false, p0, p0);
        let _index66 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_442_v27).length));
        let _rhs61 = ((_443_v28).Intersect(_443_v28)).IsProperSubsetOf((_dafny.Set.fromElements(p0)).Difference(_443_v28));
        let _rhs62 = _437_v22;
        let _rhs63 = ((_this.f25).minus((_dafny.ZERO).minus(p1))).multipliedBy(new BigNumber(-958));
        let _rhs64 = _413_v4;
        let _lhs59 = globalState;
        let _lhs60 = _442_v27;
        let _lhs61 = _module.__default.safeIndex(new BigNumber(949), new BigNumber((_442_v27).length));
        _lhs59.f18 = _rhs61;
        _437_v22 = _rhs62;
        _lhs60[_lhs61] = _rhs63;
        _413_v4 = _rhs64;
        let _444_v29;
        _444_v29 = _module.D6.create_DC12(p0, _this.f25);
        (globalState).f18 = (_444_v29).dtor_cf17;
      }
      let _445_v30;
      _445_v30 = _dafny.Set.fromElements(!(p0));
      let _446_v31;
      _446_v31 = _dafny.Seq.of(new BigNumber((_445_v30).length), p1, p1);
      r3 = ((p0) ? (new BigNumber(-490)) : ((_446_v31)[_module.__default.safeIndex((_dafny.ZERO).minus(p1), new BigNumber((_446_v31).length))]));
      (globalState).f18 = p0;
      let _447_v32;
      _447_v32 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(295),_this.f25);
      if (((((_447_v32).contains(p1)) ? ((_447_v32).get(p1)) : (_this.f25))).isLessThanOrEqualTo(_this.f25)) {
        if (!(((p0) ? (p0) : (p0)))) {
          let _index67 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_407_v1).length));
          (_407_v1)[_index67] = ((p0) ? (p0) : (p0));
          let _448_v33;
          _448_v33 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,p0);
          let _index68 = _module.__default.safeIndex(new BigNumber(397), new BigNumber((_407_v1).length));
          (_407_v1)[_index68] = (((p0) ? ((_448_v33).update(_this.f25, false)) : (_448_v33))).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(176),true));
          let _449_v34;
          let _init12 = ((_450_v1) => function (_451_i2) {
            return (_450_v1)[_module.__default.safeIndex(new BigNumber(397), new BigNumber((_450_v1).length))];
          })(_407_v1);
          let _nw77 = Array((new BigNumber(13)).toNumber());
          for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw77.length); _i0_12++) {
            _nw77[_i0_12] = _init12(new BigNumber(_i0_12));
          }
          _449_v34 = _nw77;
          _449_v34 = _449_v34;
          r0 = new BigNumber(-777);
          let _452_v35;
          _452_v35 = _dafny.Seq.of(!((_407_v1)[_module.__default.safeIndex(new BigNumber(397), new BigNumber((_407_v1).length))]));
          _452_v35 = _module.__default.fm7(globalState);
          (globalState).f18 = _module.__default.fm0((p1).multipliedBy(_this.f25), (_407_v1)[_module.__default.safeIndex(new BigNumber(397), new BigNumber((_407_v1).length))], p0, (_this).fm6(globalState), globalState);
        } else {
          let _453_v36;
          _453_v36 = _dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(309)), function (_454_i3) {
            return (_this).f24;
          })).length));
          let _455_v37;
          _455_v37 = _module.D7.create_DC14(p0, new BigNumber(882), _this.f25);
          (globalState).f18 = ((new BigNumber((_453_v36).cardinality())).minus((_455_v37).dtor_cf21)).isLessThan(new BigNumber((_445_v30).length));
          (globalState).f18 = p0;
          let _456_v38;
          _456_v38 = _dafny.Seq.of(p0);
          let _457_v39;
          _457_v39 = _dafny.Seq.UnicodeFromString("ybuioxvd");
          let _458_v40;
          _458_v40 = _dafny.MultiSet.fromElements(p0, p0, p0, p0);
          let _459_v41;
          let _nw78 = Array((new BigNumber(12)).toNumber());
          _nw78[(_dafny.ZERO).toNumber()] = (new BigNumber((_dafny.Seq.of(_this.f25)).length)).plus(_this.f25);
          _nw78[(_dafny.ONE).toNumber()] = (_this.f25).minus((_dafny.ZERO).minus(p1));
          _nw78[(new BigNumber(2)).toNumber()] = (_this).fm6(globalState);
          _nw78[(new BigNumber(3)).toNumber()] = p1;
          _nw78[(new BigNumber(4)).toNumber()] = _this.f25;
          _nw78[(new BigNumber(5)).toNumber()] = new BigNumber(((_module.__default.fm8(new BigNumber((_457_v39).length), p0, _456_v38, globalState)).Difference(_458_v40)).cardinality());
          _nw78[(new BigNumber(6)).toNumber()] = p1;
          _nw78[(new BigNumber(7)).toNumber()] = (_this.f25).plus(p1);
          _nw78[(new BigNumber(8)).toNumber()] = (new BigNumber(687)).multipliedBy(_this.f25);
          _nw78[(new BigNumber(9)).toNumber()] = (_this).fm6(globalState);
          _nw78[(new BigNumber(10)).toNumber()] = p1;
          _nw78[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(p1);
          _459_v41 = _nw78;
          let _index69 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_459_v41).length));
          (_459_v41)[_index69] = (_this).fm6(globalState);
          let _460_v42;
          _460_v42 = _module.D6.create_DC12(p0, p1);
          let _461_v43;
          _461_v43 = _module.D3.create_DC4(_module.__default.fm9(true, _dafny.Seq.update(_456_v38, _module.__default.safeIndex(p1, new BigNumber((_456_v38).length)), (_460_v42).dtor_cf17), globalState));
          let _index70 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_459_v41).length));
          let _rhs65 = _dafny.Seq.update(_dafny.Seq.of(p0), _module.__default.safeIndex(_this.f25, new BigNumber((_dafny.Seq.of(p0)).length)), _module.__default.fm0(p1, false, p0, _this.f25, globalState));
          let _rhs66 = (_this).fm3((_this).f24, globalState);
          let _rhs67 = (_this).fm6(globalState);
          let _rhs68 = _461_v43;
          let _lhs62 = globalState;
          let _lhs63 = _459_v41;
          let _lhs64 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_459_v41).length));
          _456_v38 = _rhs65;
          _lhs62.f3 = _rhs66;
          _lhs63[_lhs64] = _rhs67;
          _461_v43 = _rhs68;
          let _462_v44;
          _462_v44 = _module.D5.create_DC9(_459_v41);
          let _pat_let_tv2 = _459_v41;
          let _463_v45;
          let _nw79 = Array((new BigNumber(16)).toNumber());
          _nw79[(_dafny.ZERO).toNumber()] = _module.D5.create_DC9(_459_v41);
          _nw79[(_dafny.ONE).toNumber()] = _462_v44;
          _nw79[(new BigNumber(2)).toNumber()] = _462_v44;
          _nw79[(new BigNumber(3)).toNumber()] = _module.D5.create_DC9(_459_v41);
          _nw79[(new BigNumber(4)).toNumber()] = _module.D5.create_DC9(_459_v41);
          _nw79[(new BigNumber(5)).toNumber()] = _462_v44;
          _nw79[(new BigNumber(6)).toNumber()] = _462_v44;
          _nw79[(new BigNumber(7)).toNumber()] = _462_v44;
          _nw79[(new BigNumber(8)).toNumber()] = _module.D5.create_DC9(_459_v41);
          _nw79[(new BigNumber(9)).toNumber()] = function (_pat_let4_0) {
            return function (_464_dt__update__tmp_h0) {
              return function (_pat_let5_0) {
                return function (_465_dt__update_hcf13_h0) {
                  return _module.D5.create_DC9(_465_dt__update_hcf13_h0);
                }(_pat_let5_0);
              }(_pat_let_tv2);
            }(_pat_let4_0);
          }(_module.D5.create_DC9(_459_v41));
          _nw79[(new BigNumber(10)).toNumber()] = _462_v44;
          _nw79[(new BigNumber(11)).toNumber()] = _462_v44;
          _nw79[(new BigNumber(12)).toNumber()] = _462_v44;
          _nw79[(new BigNumber(13)).toNumber()] = _462_v44;
          _nw79[(new BigNumber(14)).toNumber()] = _462_v44;
          _nw79[(new BigNumber(15)).toNumber()] = _module.D5.create_DC9(_459_v41);
          _463_v45 = _nw79;
          let _index71 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_463_v45).length));
          (_463_v45)[_index71] = _462_v44;
          let _pat_let_tv3 = _459_v41;
          let _index72 = _module.__default.safeIndex(new BigNumber(440), new BigNumber((_463_v45).length));
          (_463_v45)[_index72] = function (_pat_let6_0) {
            return function (_466_dt__update__tmp_h1) {
              return function (_pat_let7_0) {
                return function (_467_dt__update_hcf13_h1) {
                  return _module.D5.create_DC9(_467_dt__update_hcf13_h1);
                }(_pat_let7_0);
              }(_pat_let_tv3);
            }(_pat_let6_0);
          }(_462_v44);
          let _index73 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_407_v1).length));
          (_407_v1)[_index73] = p0;
          let _index74 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_407_v1).length));
          let _index75 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_459_v41).length));
          let _rhs69 = p0;
          let _rhs70 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(_this.f25, (_459_v41)[_module.__default.safeIndex(new BigNumber(18), new BigNumber((_459_v41).length))]));
          let _rhs71 = p1;
          let _rhs72 = p1;
          let _rhs73 = p0;
          let _lhs65 = _407_v1;
          let _lhs66 = _module.__default.safeIndex(new BigNumber(887), new BigNumber((_407_v1).length));
          let _lhs67 = _459_v41;
          let _lhs68 = _module.__default.safeIndex(new BigNumber(18), new BigNumber((_459_v41).length));
          let _lhs69 = _this;
          let _lhs70 = globalState;
          _lhs65[_lhs66] = _rhs69;
          _lhs67[_lhs68] = _rhs70;
          _lhs69.f25 = _rhs71;
          r0 = _rhs72;
          _lhs70.f18 = _rhs73;
        }
        let _468_v46;
        let _nw80 = Array((new BigNumber(10)).toNumber());
        _nw80[(_dafny.ZERO).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw80[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('e'.codePointAt(0));
        _nw80[(new BigNumber(2)).toNumber()] = (_this).f24;
        _nw80[(new BigNumber(3)).toNumber()] = (_this).f24;
        _nw80[(new BigNumber(4)).toNumber()] = _module.__default.fm10(p0, p0, p0, p0, globalState);
        _nw80[(new BigNumber(5)).toNumber()] = (_this).f24;
        _nw80[(new BigNumber(6)).toNumber()] = (_this).f24;
        _nw80[(new BigNumber(7)).toNumber()] = ((p0) ? ((_this).f24) : ((_this).f24));
        _nw80[(new BigNumber(8)).toNumber()] = (_this).f24;
        _nw80[(new BigNumber(9)).toNumber()] = (_this).f24;
        _468_v46 = _nw80;
        _468_v46 = _468_v46;
        (globalState).f18 = !(true);
        let _469_v47;
        _469_v47 = _module.D7.create_DC14(p0, p1, p1);
        (globalState).f18 = (_469_v47).dtor_cf20;
        let _470_v48;
        _470_v48 = _module.D2.create_DC2(p0);
        _470_v48 = _470_v48;
      } else {
        let _471_v49;
        let _nw81 = new _module.C0();
        _nw81.__ctor();
        _471_v49 = _nw81;
        (globalState).f18 = p0;
        let _472_v50;
        _472_v50 = _dafny.Seq.of(p0);
        let _473_v51;
        _473_v51 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p1),_472_v50);
        let _474_v52;
        _474_v52 = _dafny.Seq.UnicodeFromString("fwhru");
        let _rhs74 = !((_dafny.MultiSet.fromElements(p1, p1)).IsDisjointFrom(_dafny.MultiSet.FromArray(_446_v31)));
        let _rhs75 = new BigNumber((_474_v52).length);
        let _rhs76 = _473_v51;
        let _lhs71 = globalState;
        let _lhs72 = globalState;
        _lhs71.f18 = _rhs74;
        _lhs72.f11 = _rhs75;
        _473_v51 = _rhs76;
        let _475_v53;
        _475_v53 = _dafny.MultiSet.fromElements(!(p0));
        let _rhs77 = (_this.f25).plus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(_this.f25, new BigNumber(-356))).length))).length));
        let _rhs78 = p1;
        let _rhs79 = (_module.__default.fm8(_this.f25, p0, _472_v50, globalState)).IsSubsetOf((_475_v53).Union(_475_v53));
        let _lhs73 = globalState;
        let _lhs74 = globalState;
        let _lhs75 = globalState;
        _lhs73.f22 = _rhs77;
        _lhs74.f16 = _rhs78;
        _lhs75.f18 = _rhs79;
        if (p0) {
          let _476_v54;
          let _nw82 = Array((new BigNumber(22)).toNumber());
          _nw82[(_dafny.ZERO).toNumber()] = _447_v32;
          _nw82[(_dafny.ONE).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_this.f25,p1)).Merge((_447_v32).update((_this).fm6(globalState), p1));
          _nw82[(new BigNumber(2)).toNumber()] = (_447_v32).Merge(_447_v32);
          _nw82[(new BigNumber(3)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(4)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(5)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(6)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(7)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(8)).toNumber()] = (_447_v32).Merge(_447_v32);
          _nw82[(new BigNumber(9)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(10)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(p1,p1)).Merge(_447_v32);
          _nw82[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f25);
          _nw82[(new BigNumber(12)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(13)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(14)).toNumber()] = (_447_v32).Merge(_dafny.Map.Empty.slice().updateUnsafe((_this).fm6(globalState),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f25,_407_v1)).length)));
          _nw82[(new BigNumber(15)).toNumber()] = (_module.__default.fm11(new _dafny.CodePoint('c'.codePointAt(0)), p0, globalState)).Merge(_447_v32);
          _nw82[(new BigNumber(16)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,_this.f25);
          _nw82[(new BigNumber(17)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(18)).toNumber()] = ((true) ? (_447_v32) : ((_447_v32).update(p1, new BigNumber((_dafny.Set.fromElements(p0, p0, p0)).length))));
          _nw82[(new BigNumber(19)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(20)).toNumber()] = _447_v32;
          _nw82[(new BigNumber(21)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          _476_v54 = _nw82;
          let _index76 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_476_v54).length));
          (_476_v54)[_index76] = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _477_v55;
          _477_v55 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,(_this).f24);
          let _478_v56;
          _478_v56 = _dafny.Map.Empty.slice().updateUnsafe(p1,_dafny.Map.Empty.slice().updateUnsafe((_this).f24,true));
          let _index77 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_476_v54).length));
          let _rhs80 = _module.__default.fm11((((_477_v55).contains(p1)) ? ((_477_v55).get(p1)) : (new _dafny.CodePoint('t'.codePointAt(0)))), p0, globalState);
          let _rhs81 = ((p0) ? (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_478_v56).length),(_this).fm6(globalState))).length)) : ((new BigNumber(248)).minus(new BigNumber(281))));
          let _rhs82 = p0;
          let _rhs83 = _471_v49;
          let _lhs76 = _476_v54;
          let _lhs77 = _module.__default.safeIndex(new BigNumber(7), new BigNumber((_476_v54).length));
          let _lhs78 = globalState;
          let _lhs79 = globalState;
          _lhs76[_lhs77] = _rhs80;
          _lhs78.f22 = _rhs81;
          _lhs79.f18 = _rhs82;
          _471_v49 = _rhs83;
          (globalState).f18 = p0;
          (globalState).f18 = p0;
          let _479_v57;
          _479_v57 = _dafny.Map.Empty.slice().updateUnsafe((p1).minus(p1),_406_v0);
          _479_v57 = _479_v57;
          (globalState).f18 = p0;
        } else {
          let _index78 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length));
          (_407_v1)[_index78] = (_this.f25).isLessThanOrEqualTo(p1);
          let _index79 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length));
          (_407_v1)[_index79] = (_472_v50)[_module.__default.safeIndex(new BigNumber(((_module.D3.create_DC5(_472_v50, new BigNumber(-73), new BigNumber(-879))).dtor_cf6).length), new BigNumber((_472_v50).length))];
          (globalState).f0 = _module.__default.safeDivisionInt(((p0) ? (_this.f25) : ((_this).fm6(globalState))), _module.__default.safeDivisionInt(p1, p1));
          (globalState).f11 = _this.f25;
          let _480_v58;
          _480_v58 = _dafny.Seq.of(_445_v30);
          let _481_v59;
          let _nw83 = Array((new BigNumber(20)).toNumber());
          _nw83[(_dafny.ZERO).toNumber()] = _445_v30;
          _nw83[(_dafny.ONE).toNumber()] = (_445_v30).Difference(_445_v30);
          _nw83[(new BigNumber(2)).toNumber()] = (_445_v30).Difference(_dafny.Set.fromElements((_407_v1)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length))], p0));
          _nw83[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(!((_407_v1)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length))]));
          _nw83[(new BigNumber(4)).toNumber()] = _dafny.Set.fromElements((_407_v1)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length))]);
          _nw83[(new BigNumber(5)).toNumber()] = _445_v30;
          _nw83[(new BigNumber(6)).toNumber()] = (_480_v58)[_module.__default.safeIndex(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_446_v31).length), p1)).cardinality()), new BigNumber((_480_v58).length))];
          _nw83[(new BigNumber(7)).toNumber()] = _445_v30;
          _nw83[(new BigNumber(8)).toNumber()] = _445_v30;
          _nw83[(new BigNumber(9)).toNumber()] = _445_v30;
          _nw83[(new BigNumber(10)).toNumber()] = _dafny.Set.fromElements((_407_v1)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length))], (_407_v1)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length))]);
          _nw83[(new BigNumber(11)).toNumber()] = (_dafny.Set.fromElements((_407_v1)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length))], p0)).Intersect(_445_v30);
          _nw83[(new BigNumber(12)).toNumber()] = _445_v30;
          _nw83[(new BigNumber(13)).toNumber()] = _445_v30;
          _nw83[(new BigNumber(14)).toNumber()] = (_445_v30).Union(_445_v30);
          _nw83[(new BigNumber(15)).toNumber()] = _445_v30;
          _nw83[(new BigNumber(16)).toNumber()] = _445_v30;
          _nw83[(new BigNumber(17)).toNumber()] = _445_v30;
          _nw83[(new BigNumber(18)).toNumber()] = _445_v30;
          _nw83[(new BigNumber(19)).toNumber()] = _445_v30;
          _481_v59 = _nw83;
          let _index80 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_481_v59).length));
          (_481_v59)[_index80] = (_445_v30).Difference(_445_v30);
          let _index81 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_481_v59).length));
          let _rhs84 = (_module.__default.fm12((_407_v1)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length))], p0, true, globalState)).Union((_445_v30).Intersect(_445_v30));
          let _rhs85 = (_407_v1)[_module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length))];
          let _lhs80 = _481_v59;
          let _lhs81 = _module.__default.safeIndex(new BigNumber(772), new BigNumber((_481_v59).length));
          let _lhs82 = globalState;
          _lhs80[_lhs81] = _rhs84;
          _lhs82.f18 = _rhs85;
          let _index82 = _module.__default.safeIndex(new BigNumber(118), new BigNumber((_407_v1).length));
          (_407_v1)[_index82] = _module.__default.fm0(_module.__default.safeDivisionInt(_this.f25, new BigNumber((_474_v52).length)), (_dafny.Map.Empty.slice().updateUnsafe(p1,_this.f25)).contains(new BigNumber(658)), !((_445_v30).IsSubsetOf((_481_v59)[_module.__default.safeIndex(new BigNumber(772), new BigNumber((_481_v59).length))])), (_dafny.ZERO).minus(p1), globalState);
        }
      }
      r0 = _module.__default.safeModuloInt(_this.f25, p1);
      r1 = (((true) ? ((_this).fm6(globalState)) : (p1))).minus(((_module.__default.fm0(p1, p0, p0, _this.f25, globalState)) ? (p1) : (_this.f25)));
      r2 = (_dafny.ZERO).minus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p1, _this.f25)));
      r3 = _this.f25;
      return [r0, r1, r2, r3];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this._f27 = _dafny.ZERO;
      this._f28 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f27, f28) {
      let _this = this;
      (_this)._f27 = f27;
      (_this)._f28 = f28;
      return;
    }
    fm4(p0, p1, p2, globalState) {
      let _this = this;
      return (_dafny.MultiSet.fromElements((_this).f28, (_this).f28, (_this).f28, (_this).f28, (_this).f28)).Intersect((_dafny.MultiSet.fromElements((_this).f28, (_this).f28, true, (_this).f28)).Union(_dafny.MultiSet.fromElements((_this).f28)));
    };
    m3(p0, p1, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let _482_v0;
      _482_v0 = new _dafny.CodePoint('m'.codePointAt(0));
      let _483_v1;
      _483_v1 = _dafny.Seq.UnicodeFromString("gxn");
      let _484_v2;
      _484_v2 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(269)), ((_485_v0) => function (_486_i0) {
        return _485_v0;
      })(_482_v0));
      let _487_v3;
      let _nw84 = Array((new BigNumber(12)).toNumber());
      _nw84[(_dafny.ZERO).toNumber()] = _483_v1;
      _nw84[(_dafny.ONE).toNumber()] = _484_v2;
      _nw84[(new BigNumber(2)).toNumber()] = _484_v2;
      _nw84[(new BigNumber(3)).toNumber()] = _483_v1;
      _nw84[(new BigNumber(4)).toNumber()] = _484_v2;
      _nw84[(new BigNumber(5)).toNumber()] = _484_v2;
      _nw84[(new BigNumber(6)).toNumber()] = _484_v2;
      _nw84[(new BigNumber(7)).toNumber()] = _484_v2;
      _nw84[(new BigNumber(8)).toNumber()] = _483_v1;
      _nw84[(new BigNumber(9)).toNumber()] = _484_v2;
      _nw84[(new BigNumber(10)).toNumber()] = _484_v2;
      _nw84[(new BigNumber(11)).toNumber()] = _484_v2;
      _487_v3 = _nw84;
      let _488_v4;
      _488_v4 = _dafny.Map.Empty.slice().updateUnsafe(((_this).f27).plus((_this).f27),_487_v3);
      let _489_v5;
      _489_v5 = _dafny.Set.fromElements((_this).f27, (_this).f27, (_this).f27, new BigNumber(-942));
      let _rhs86 = (_this).f28;
      let _rhs87 = new _dafny.CodePoint('i'.codePointAt(0));
      let _rhs88 = _483_v1;
      let _rhs89 = (((_488_v4).contains(new BigNumber((_489_v5).length))) ? ((_488_v4).get(new BigNumber((_489_v5).length))) : (_487_v3));
      let _lhs83 = globalState;
      let _lhs84 = globalState;
      _lhs83.f18 = _rhs86;
      _482_v0 = _rhs87;
      _lhs84.f3 = _rhs88;
      _487_v3 = _rhs89;
      let _490_v6;
      _490_v6 = _dafny.MultiSet.fromElements(p0, (_this).f28);
      let _491_v7;
      _491_v7 = _module.D3.create_DC5(_dafny.Seq.of((_this).f28), new BigNumber((_490_v6).cardinality()), (_this).f27);
      let _492_v8;
      _492_v8 = (_491_v7).dtor_cf6;
      let _493_v9;
      _493_v9 = _dafny.Seq.of((_this).f28);
      let _source8 = _493_v9;
      let _494___mcc_h0 = _source8;
      let _495_cf0 = _494___mcc_h0;
      if (_dafny.Seq.IsPrefixOf(_483_v1, _dafny.Seq.UnicodeFromString("q"))) {
        let _496_v10;
        let _nw85 = Array((new BigNumber(5)).toNumber()).fill(_dafny.Set.Empty);
        _496_v10 = _nw85;
        _496_v10 = _496_v10;
        let _497_v11;
        _497_v11 = _dafny.Map.Empty.slice().updateUnsafe(_483_v1,(_491_v7).dtor_cf7);
        _497_v11 = (_497_v11).update(_dafny.Seq.Concat(_483_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(59)), function (_498_i1) {
          return new _dafny.CodePoint('p'.codePointAt(0));
        })), (_dafny.ZERO).minus(((_this).f27).multipliedBy((_this).f27)));
        (globalState).f3 = _483_v1;
        let _499_v12;
        let _nw86 = Array((new BigNumber(18)).toNumber()).fill(_dafny.Map.Empty);
        _499_v12 = _nw86;
        let _500_v13;
        _500_v13 = _module.D4.create_DC7(_499_v12);
        _499_v12 = (_500_v13).dtor_cf12;
        let _501_v14;
        _501_v14 = _dafny.Set.fromElements((_this).f28, p0, (!(p1)) && (p1), ((_this).f28) === (p0), (_this).f28);
        (globalState).f22 = new BigNumber((_501_v14).length);
      } else {
        (globalState).f0 = (_this).f27;
        let _502_v15;
        let _nw87 = new _module.C0();
        _nw87.__ctor();
        _502_v15 = _nw87;
        _490_v6 = (_dafny.MultiSet.FromArray(_dafny.Seq.Concat(_493_v9, (_module.D3.create_DC5(_495_cf0, (_this).f27, (_this).f27)).dtor_cf6))).Intersect(_490_v6);
        let _503_v16;
        let _nw88 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
        _503_v16 = _nw88;
        let _504_v17;
        _504_v17 = _module.D5.create_DC9(_503_v16);
        let _505_v18;
        _505_v18 = _dafny.Seq.of((_504_v17).dtor_cf13);
        (globalState).f11 = new BigNumber((_505_v18).length);
        let _506_v19;
        _506_v19 = _module.D5.create_DC10(p1, (_495_cf0)[_module.__default.safeIndex((_this).f27, new BigNumber((_495_cf0).length))]);
        let _507_v20;
        let _nw89 = Array((new BigNumber(19)).toNumber());
        _nw89[(_dafny.ZERO).toNumber()] = p1;
        _nw89[(_dafny.ONE).toNumber()] = !(((_this).f28) || (p0));
        _nw89[(new BigNumber(2)).toNumber()] = !(p0);
        _nw89[(new BigNumber(3)).toNumber()] = (((_495_cf0)[_module.__default.safeIndex((_this).f27, new BigNumber((_495_cf0).length))]) ? ((_this).f28) : (!((_this).f28)));
        _nw89[(new BigNumber(4)).toNumber()] = !((_489_v5).IsDisjointFrom(_dafny.Set.fromElements((_this).f27)));
        _nw89[(new BigNumber(5)).toNumber()] = (_dafny.MultiSet.fromElements(!((_this).f28), (_495_cf0)[_module.__default.safeIndex((_this).f27, new BigNumber((_495_cf0).length))])).equals(_490_v6);
        _nw89[(new BigNumber(6)).toNumber()] = (_this).f28;
        _nw89[(new BigNumber(7)).toNumber()] = !(p1) || (p1);
        _nw89[(new BigNumber(8)).toNumber()] = (_this).f28;
        _nw89[(new BigNumber(9)).toNumber()] = (_506_v19).dtor_cf14;
        _nw89[(new BigNumber(10)).toNumber()] = (_this).f28;
        _nw89[(new BigNumber(11)).toNumber()] = p1;
        _nw89[(new BigNumber(12)).toNumber()] = (_this).f28;
        _nw89[(new BigNumber(13)).toNumber()] = _module.__default.fm0((_this).f27, p1, (_this).f28, (_this).f27, globalState);
        _nw89[(new BigNumber(14)).toNumber()] = p1;
        _nw89[(new BigNumber(15)).toNumber()] = !(p0);
        _nw89[(new BigNumber(16)).toNumber()] = ((_this).f27).isLessThan((_this).f27);
        _nw89[(new BigNumber(17)).toNumber()] = p1;
        _nw89[(new BigNumber(18)).toNumber()] = (_this).f28;
        _507_v20 = _nw89;
        let _index83 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_507_v20).length));
        (_507_v20)[_index83] = ((_this).f27).isLessThan((_this).f27);
        let _index84 = _module.__default.safeIndex(new BigNumber(602), new BigNumber((_507_v20).length));
        (_507_v20)[_index84] = (((_this).f28) ? (p0) : (p0));
      }
      (globalState).f19 = _dafny.Seq.of((_this).f27, (_this).f27);
      let _508_v21;
      let _nw90 = Array((new BigNumber(20)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
      _508_v21 = _nw90;
      let _init13 = ((_509_v0) => function (_510_i2) {
        return _509_v0;
      })(_482_v0);
      let _nw91 = Array((new BigNumber(3)).toNumber());
      for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw91.length); _i0_13++) {
        _nw91[_i0_13] = _init13(new BigNumber(_i0_13));
      }
      _508_v21 = _nw91;
      let _511_v22;
      _511_v22 = _module.D3.create_DC6(new BigNumber(617), (_dafny.ZERO).minus((_this).f27), p0);
      let _512_v23;
      _512_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,p1);
      let _513_v24;
      _513_v24 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber(381));
      let _514_v25;
      _514_v25 = _dafny.MultiSet.fromElements((_511_v22).dtor_cf10, (_this).f27, new BigNumber((_512_v23).length), new BigNumber(((_489_v5).Union(_489_v5)).length), new BigNumber((((_513_v24).update(p1, (_this).f27)).Merge((_513_v24).update(p0, (_this).f27))).length));
      let _515_v26;
      _515_v26 = _dafny.Seq.of(_508_v21);
      let _516_v27;
      _516_v27 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(p0,(_dafny.ZERO).minus((_this).f27)), _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_this).f27), _513_v24);
      let _rhs90 = ((_514_v25).Union(_514_v25)).Union((_514_v25).Intersect(_dafny.MultiSet.fromElements((((_514_v25).contains((_this).f27)) ? ((_514_v25).get((_this).f27)) : ((_this).f27)))));
      let _rhs91 = _dafny.Seq.of(_508_v21);
      let _rhs92 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("o"), _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_516_v27, _dafny.Seq.of(_513_v24, _513_v24))).length), new BigNumber((_dafny.Seq.UnicodeFromString("o")).length)), _482_v0);
      let _rhs93 = !(p1);
      let _lhs85 = globalState;
      let _lhs86 = globalState;
      _514_v25 = _rhs90;
      _515_v26 = _rhs91;
      _lhs85.f3 = _rhs92;
      _lhs86.f18 = _rhs93;
      let _517_v28;
      _517_v28 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
      let _518_v29;
      _518_v29 = _dafny.Set.fromElements(_517_v28);
      if (!(!((_518_v29).IsProperSubsetOf(_518_v29)))) {
        let _519_v30;
        _519_v30 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_this).f27), new BigNumber(-840));
        let _520_v31;
        _520_v31 = _dafny.Seq.of(_489_v5);
        let _521_v32;
        let _nw92 = Array((new BigNumber(24)).toNumber());
        _nw92[(_dafny.ZERO).toNumber()] = new BigNumber((_519_v30).cardinality());
        _nw92[(_dafny.ONE).toNumber()] = new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber(872))).length);
        _nw92[(new BigNumber(2)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(3)).toNumber()] = new BigNumber((_483_v1).length);
        _nw92[(new BigNumber(4)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(5)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("l")).length);
        _nw92[(new BigNumber(6)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus((_this).f27);
        _nw92[(new BigNumber(8)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(9)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(10)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(11)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(12)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(13)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(14)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("xrr")).length);
        _nw92[(new BigNumber(15)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(16)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(17)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(18)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(19)).toNumber()] = new BigNumber((_489_v5).length);
        _nw92[(new BigNumber(20)).toNumber()] = new BigNumber((_520_v31).length);
        _nw92[(new BigNumber(21)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(22)).toNumber()] = (_this).f27;
        _nw92[(new BigNumber(23)).toNumber()] = (_this).f27;
        _521_v32 = _nw92;
        let _522_v33;
        _522_v33 = _module.D5.create_DC9(_521_v32);
        let _source9 = _522_v33;
        if (_source9.is_DC10) {
          let _523___mcc_h1 = (_source9).cf14;
          let _524___mcc_h2 = (_source9).cf15;
          let _525_cf15 = _524___mcc_h2;
          let _526_cf14 = _523___mcc_h1;
          let _527_v34;
          _527_v34 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("ofy"), _483_v1);
          let _528_v35;
          _528_v35 = _module.D3.create_DC6(((p0) ? ((_this).f27) : ((_this).f27)), new BigNumber(((_527_v34).Union(_527_v34)).length), true);
          let _529_v36;
          let _nw93 = Array((new BigNumber(5)).toNumber()).fill([]);
          _529_v36 = _nw93;
          let _530_v37;
          let _init14 = ((_531_v30) => function (_532_i3) {
            return _531_v30;
          })(_519_v30);
          let _nw94 = Array((new BigNumber(27)).toNumber());
          for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw94.length); _i0_14++) {
            _nw94[_i0_14] = _init14(new BigNumber(_i0_14));
          }
          _530_v37 = _nw94;
          let _index85 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_529_v36).length));
          (_529_v36)[_index85] = _530_v37;
          let _index86 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_529_v36).length));
          let _rhs94 = !(p1);
          let _rhs95 = _528_v35;
          let _rhs96 = _530_v37;
          let _lhs87 = globalState;
          let _lhs88 = _529_v36;
          let _lhs89 = _module.__default.safeIndex(new BigNumber(46), new BigNumber((_529_v36).length));
          _lhs87.f18 = _rhs94;
          _528_v35 = _rhs95;
          _lhs88[_lhs89] = _rhs96;
          let _533_v38;
          _533_v38 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,new BigNumber(200));
          (globalState).f18 = ((_519_v30).Intersect(_dafny.MultiSet.fromElements((_this).f27, (_this).f27, (_this).f27, (_this).f27, _module.__default.fm5((_this).f27, (_this).f27, globalState)))).equals(_dafny.MultiSet.fromElements((_this).f27, new BigNumber((_533_v38).length)));
          let _534_v39;
          _534_v39 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,p1);
          let _535_v40;
          _535_v40 = _dafny.Set.fromElements(_534_v39, _dafny.Map.Empty.slice().updateUnsafe((_this).f27,p0), _534_v39, _dafny.Map.Empty.slice().updateUnsafe((((_533_v38).contains((_this).f27)) ? ((_533_v38).get((_this).f27)) : ((_this).f27)),(_this).f28));
          let _536_v41;
          _536_v41 = _module.D3.create_DC4(_535_v40);
          let _537_v42;
          _537_v42 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_536_v41);
          _537_v42 = (_537_v42).update((_this).f27, _module.D3.create_DC4(_535_v40));
          let _538_v43;
          _538_v43 = _dafny.MultiSet.fromElements(_521_v32, _521_v32, _521_v32);
          let _rhs97 = _526_cf14;
          let _rhs98 = !(p1);
          let _rhs99 = p0;
          let _rhs100 = _dafny.MultiSet.fromElements((_this).f27, (_this).f27, new BigNumber(((_538_v43).Intersect(_538_v43)).cardinality()));
          let _lhs90 = globalState;
          let _lhs91 = globalState;
          _lhs90.f18 = _rhs97;
          _526_cf14 = _rhs98;
          _lhs91.f18 = _rhs99;
          _519_v30 = _rhs100;
        } else {
          let _539___mcc_h3 = (_source9).cf13;
          let _540_cf13 = _539___mcc_h3;
          _484_v2 = _484_v2;
          let _541_v44;
          _541_v44 = _dafny.Seq.of(new BigNumber((_493_v9).length));
          let _542_v45;
          _542_v45 = _dafny.Map.Empty.slice().updateUnsafe(_541_v44,((_this).f27).isLessThanOrEqualTo(new BigNumber((_483_v1).length)));
          _542_v45 = (_542_v45).update(_dafny.Seq.Concat(_541_v44, _541_v44), true);
          (globalState).f22 = new BigNumber(409);
          (globalState).f0 = (_this).f27;
        }
        (globalState).f15 = (new BigNumber((_483_v1).length)).plus(((_this).f27).multipliedBy((_this).f27));
        let _543_v46;
        _543_v46 = _dafny.Seq.of((_this).f27, (_this).f27);
        let _544_v47;
        _544_v47 = _module.D2.create_DC2(((p1) ? (p0) : (_module.__default.fm0(new BigNumber((_543_v46).length), p0, p0, (_this).f27, globalState))));
        let _source10 = _544_v47;
        if (_source10.is_DC3) {
          let _545___mcc_h4 = (_source10).cf3;
          let _546___mcc_h5 = (_source10).cf4;
          let _547_cf4 = _546___mcc_h5;
          let _548_cf3 = _545___mcc_h4;
          let _549_v48;
          _549_v48 = _dafny.Map.Empty.slice().updateUnsafe(p1,_521_v32);
          _549_v48 = (_549_v48).update(p1, _521_v32);
          let _550_v49;
          let _nw95 = new _module.C0();
          _nw95.__ctor();
          _550_v49 = _nw95;
          let _index87 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_521_v32).length));
          (_521_v32)[_index87] = (_548_cf3).plus(new BigNumber((_dafny.Set.fromElements(false)).length));
          let _index88 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_521_v32).length));
          (_521_v32)[_index88] = _547_cf4;
          let _551_v50;
          _551_v50 = _dafny.Seq.of(_483_v1, _dafny.Seq.UnicodeFromString("sjtnwy"), _483_v1);
          let _index89 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_521_v32).length));
          let _rhs101 = (_this).f27;
          let _rhs102 = (_521_v32)[_module.__default.safeIndex(new BigNumber(89), new BigNumber((_521_v32).length))];
          let _rhs103 = (_this).f27;
          let _rhs104 = _482_v0;
          let _rhs105 = (_551_v50)[_module.__default.safeIndex(_module.__default.safeModuloInt(new BigNumber((_543_v46).length), new BigNumber(682)), new BigNumber((_551_v50).length))];
          let _lhs92 = globalState;
          let _lhs93 = _521_v32;
          let _lhs94 = _module.__default.safeIndex(new BigNumber(89), new BigNumber((_521_v32).length));
          let _lhs95 = globalState;
          let _lhs96 = globalState;
          _lhs92.f16 = _rhs101;
          _lhs93[_lhs94] = _rhs102;
          _lhs95.f22 = _rhs103;
          _482_v0 = _rhs104;
          _lhs96.f3 = _rhs105;
        } else {
          let _552___mcc_h6 = (_source10).cf2;
          let _553_cf2 = _552___mcc_h6;
          (globalState).f16 = (_this).f27;
          let _554_v51;
          _554_v51 = _dafny.Map.Empty.slice().updateUnsafe(_483_v1,p1);
          (globalState).f18 = (((_554_v51).contains(_dafny.Seq.update(_dafny.Seq.Concat(_483_v1, _483_v1), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("mffkpiqe")).length), new BigNumber((_dafny.Seq.Concat(_483_v1, _483_v1)).length)), _482_v0))) ? ((_554_v51).get(_dafny.Seq.update(_dafny.Seq.Concat(_483_v1, _483_v1), _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("mffkpiqe")).length), new BigNumber((_dafny.Seq.Concat(_483_v1, _483_v1)).length)), _482_v0))) : ((_this).f28));
          let _555_v52;
          let _init15 = function (_556_i4) {
            return !((_this).f27).isEqualTo(new BigNumber(73));
          };
          let _nw96 = Array((new BigNumber(28)).toNumber());
          for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw96.length); _i0_15++) {
            _nw96[_i0_15] = _init15(new BigNumber(_i0_15));
          }
          _555_v52 = _nw96;
          _555_v52 = _555_v52;
          let _557_v53;
          let _nw97 = new _module.C1();
          _nw97.__ctor(new _dafny.CodePoint('a'.codePointAt(0)), (_this).f27);
          _557_v53 = _nw97;
          let _558_v54;
          _558_v54 = _module.D6.create_DC11(_557_v53);
          let _559_v55;
          _559_v55 = _dafny.Seq.of(_557_v53);
          let _560_v56;
          _560_v56 = _dafny.Seq.of(_dafny.Seq.of((_558_v54).dtor_cf16, _557_v53), _559_v55, _559_v55, _559_v55, _559_v55);
          let _561_v57;
          _561_v57 = _dafny.MultiSet.fromElements(_559_v55, _559_v55, _559_v55);
          let _562_v58;
          let _nw98 = Array((new BigNumber(10)).toNumber());
          _nw98[(_dafny.ZERO).toNumber()] = _dafny.MultiSet.FromArray(_560_v56);
          _nw98[(_dafny.ONE).toNumber()] = _561_v57;
          _nw98[(new BigNumber(2)).toNumber()] = _561_v57;
          _nw98[(new BigNumber(3)).toNumber()] = _561_v57;
          _nw98[(new BigNumber(4)).toNumber()] = _561_v57;
          _nw98[(new BigNumber(5)).toNumber()] = _dafny.MultiSet.FromArray(_560_v56);
          _nw98[(new BigNumber(6)).toNumber()] = _561_v57;
          _nw98[(new BigNumber(7)).toNumber()] = ((p1) ? (_561_v57) : (_dafny.MultiSet.fromElements(_dafny.Seq.update(_559_v55, _module.__default.safeIndex(new BigNumber((_518_v29).length), new BigNumber((_559_v55).length)), _557_v53), _559_v55, _dafny.Seq.of(_557_v53, _557_v53), _559_v55, _559_v55)));
          _nw98[(new BigNumber(8)).toNumber()] = (_dafny.MultiSet.FromArray(_560_v56)).update(_559_v55, _module.__default.abs(_557_v53.f25));
          _nw98[(new BigNumber(9)).toNumber()] = _561_v57;
          _562_v58 = _nw98;
          let _563_v59;
          _563_v59 = _559_v55;
          let _index90 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_562_v58).length));
          (_562_v58)[_index90] = _dafny.MultiSet.fromElements(_dafny.Seq.of(_557_v53), (_563_v59), _559_v55);
          let _index91 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_562_v58).length));
          let _rhs106 = _561_v57;
          let _rhs107 = _490_v6;
          let _lhs97 = _562_v58;
          let _lhs98 = _module.__default.safeIndex(new BigNumber(844), new BigNumber((_562_v58).length));
          _lhs97[_lhs98] = _rhs106;
          _490_v6 = _rhs107;
        }
        let _564_v60;
        let _nw99 = new _module.C1();
        _nw99.__ctor(_482_v0, (_this).f27);
        _564_v60 = _nw99;
        let _565_v61;
        _565_v61 = _dafny.Map.Empty.slice().updateUnsafe(_564_v60,p0);
        let _566_v62;
        _566_v62 = _module.D6.create_DC12((((_565_v61).contains(_564_v60)) ? ((_565_v61).get(_564_v60)) : (!((_this).f28))), _module.__default.fm5(new BigNumber(650), (_dafny.ZERO).minus(_module.__default.fm5(_564_v60.f25, _564_v60.f25, globalState)), globalState));
        (globalState).f18 = (_566_v62).dtor_cf17;
        (globalState).f18 = ((_dafny.ZERO).minus(((_this).f27).multipliedBy(_564_v60.f25))).isLessThanOrEqualTo(((_this).f27).minus((_this).f27));
      } else {
        (globalState).f18 = !((((_this).f28) ? (!(_517_v28).equals(_517_v28)) : ((p1) || (!((_this).f28)))));
        (globalState).f18 = !(p1);
        if (!(false) || ((_this).f28)) {
          let _567_v63;
          let _nw100 = new _module.C0();
          _nw100.__ctor();
          _567_v63 = _nw100;
          let _568_v64;
          _568_v64 = _module.D4.create_DC8();
          let _569_v65;
          _569_v65 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_module.__default.fm0(new BigNumber(-11), p0, p1, (_this).f27, globalState));
          let _570_v66;
          _570_v66 = _dafny.MultiSet.fromElements(new BigNumber((_569_v65).length));
          let _571_v67;
          _571_v67 = _dafny.Map.Empty.slice().updateUnsafe((_570_v66).Union(_570_v66),p1);
          let _572_v68;
          _572_v68 = _dafny.Seq.of((_this).f27);
          let _rhs108 = _module.D4.create_DC8();
          let _rhs109 = _571_v67;
          let _rhs110 = (((_this).f27).minus(new BigNumber(337))).plus(_module.__default.safeDivisionInt(new BigNumber((_572_v68).length), (_this).f27));
          let _lhs99 = globalState;
          _568_v64 = _rhs108;
          _571_v67 = _rhs109;
          _lhs99.f15 = _rhs110;
          (globalState).f16 = (_this).f27;
          let _573_v69;
          let _nw101 = Array((new BigNumber(5)).toNumber());
          _nw101[(_dafny.ZERO).toNumber()] = _482_v0;
          _nw101[(_dafny.ONE).toNumber()] = _module.__default.fm10((_this).f28, p0, p0, p1, globalState);
          _nw101[(new BigNumber(2)).toNumber()] = _482_v0;
          _nw101[(new BigNumber(3)).toNumber()] = _482_v0;
          _nw101[(new BigNumber(4)).toNumber()] = _482_v0;
          _573_v69 = _nw101;
          let _574_v70;
          _574_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_573_v69);
          _574_v70 = (_574_v70).update((_this).f27, _573_v69);
          (globalState).f18 = ((_this).f27).isLessThanOrEqualTo((_572_v68)[_module.__default.safeIndex((_this).f27, new BigNumber((_572_v68).length))]);
        } else {
          let _575_v71;
          let _nw102 = new _module.C1();
          _nw102.__ctor(_482_v0, new BigNumber((_490_v6).cardinality()));
          _575_v71 = _nw102;
          let _576_v72;
          _576_v72 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm5((_this).f27, (_this).f27, globalState),_575_v71);
          let _577_v73;
          _577_v73 = _module.D7.create_DC14(p0, new BigNumber((_576_v72).length), _575_v71.f25);
          let _578_v74;
          _578_v74 = _dafny.Map.Empty.slice().updateUnsafe(p1,_module.D3.create_DC6((_this).f27, (_this).f27, (_577_v73).dtor_cf20));
          let _579_v75;
          _579_v75 = _module.D9.create_DC16(_module.__default.fm13(_483_v1, globalState));
          _578_v74 = (_579_v75).dtor_cf24;
          (globalState).f2 = new BigNumber((_483_v1).length);
          let _580_v76;
          let _nw103 = Array((new BigNumber(25)).toNumber()).fill(_dafny.Seq.of());
          _580_v76 = _nw103;
          let _581_v77;
          let _init16 = function (_582_i5) {
            return false;
          };
          let _nw104 = Array((new BigNumber(15)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw104.length); _i0_16++) {
            _nw104[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _581_v77 = _nw104;
          let _583_v78;
          _583_v78 = _dafny.Seq.of(_581_v77);
          let _index92 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_580_v76).length));
          (_580_v76)[_index92] = ((p1) ? (_583_v78) : (_583_v78));
          let _index93 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_580_v76).length));
          let _rhs111 = !(p1);
          let _rhs112 = !(p0);
          let _rhs113 = (_this).f27;
          let _rhs114 = _583_v78;
          let _lhs100 = globalState;
          let _lhs101 = globalState;
          let _lhs102 = globalState;
          let _lhs103 = _580_v76;
          let _lhs104 = _module.__default.safeIndex(new BigNumber(684), new BigNumber((_580_v76).length));
          _lhs100.f18 = _rhs111;
          _lhs101.f18 = _rhs112;
          _lhs102.f22 = _rhs113;
          _lhs103[_lhs104] = _rhs114;
          let _584_v79;
          _584_v79 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,(_this).f28);
          let _585_v80;
          _585_v80 = _dafny.Map.Empty.slice().updateUnsafe((_this).f28,(_584_v79).update(_575_v71.f25, p1));
          let _index94 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_581_v77).length));
          (_581_v77)[_index94] = !(_dafny.Map.Empty.slice().updateUnsafe((_this).f28,_584_v79)).equals(_585_v80);
          let _586_v81;
          _586_v81 = _module.D3.create_DC6(new BigNumber((_584_v79).length), (_this).f27, p0);
          let _index95 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_581_v77).length));
          let _rhs115 = (new BigNumber(-589)).multipliedBy((new BigNumber(615)).minus((_this).f27));
          let _rhs116 = (_586_v81).dtor_cf9;
          let _rhs117 = p0;
          let _rhs118 = (_575_v71.f25).multipliedBy((_this).f27);
          let _lhs105 = _575_v71;
          let _lhs106 = globalState;
          let _lhs107 = _581_v77;
          let _lhs108 = _module.__default.safeIndex(new BigNumber(241), new BigNumber((_581_v77).length));
          let _lhs109 = globalState;
          _lhs105.f25 = _rhs115;
          _lhs106.f0 = _rhs116;
          _lhs107[_lhs108] = _rhs117;
          _lhs109.f22 = _rhs118;
          let _587_v82;
          _587_v82 = _module.D7.create_DC13(_584_v79);
          let _588_v83;
          _588_v83 = _dafny.MultiSet.fromElements(_587_v82, _587_v82);
          let _589_v84;
          let _nw105 = new _module.C1();
          _nw105.__ctor(_482_v0, (((_588_v83).contains(_587_v82)) ? ((_588_v83).get(_587_v82)) : (new BigNumber((_483_v1).length))));
          _589_v84 = _nw105;
        }
        let _590_v85;
        _590_v85 = _dafny.Set.fromElements(p0, p0);
        let _rhs119 = (_this).f28;
        let _rhs120 = p1;
        let _rhs121 = (_dafny.ZERO).minus((_this).f27);
        let _rhs122 = _module.__default.fm0((_dafny.ZERO).minus(((_this).f27).plus((_this).f27)), p1, (_590_v85).IsDisjointFrom((_module.__default.fm14((_this).f28, globalState)).dtor_cf26), (_this).f27, globalState);
        let _lhs110 = globalState;
        let _lhs111 = globalState;
        let _lhs112 = globalState;
        let _lhs113 = globalState;
        _lhs110.f18 = _rhs119;
        _lhs111.f18 = _rhs120;
        _lhs112.f22 = _rhs121;
        _lhs113.f18 = _rhs122;
        let _591_v86;
        _591_v86 = _dafny.MultiSet.fromElements(_482_v0);
        if (((_this).f27).isLessThanOrEqualTo(_module.__default.fm5(new BigNumber(-135), new BigNumber((_591_v86).cardinality()), globalState))) {
          (globalState).f18 = !(p0);
          (globalState).f18 = p1;
          let _index96 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_487_v3).length));
          (_487_v3)[_index96] = _484_v2;
          let _index97 = _module.__default.safeIndex(new BigNumber(398), new BigNumber((_487_v3).length));
          (_487_v3)[_index97] = _484_v2;
          let _592_v87;
          _592_v87 = _module.D5.create_DC10(p1, p1);
          let _593_v88;
          _593_v88 = _dafny.Map.Empty.slice().updateUnsafe((_this).f27,_483_v1);
          let _594_v89;
          let _nw106 = Array((new BigNumber(9)).toNumber());
          _nw106[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_483_v1, _483_v1), _module.__default.safeIndex((_this).f27, new BigNumber((_dafny.Seq.Concat(_483_v1, _483_v1)).length)), _482_v0);
          _nw106[(_dafny.ONE).toNumber()] = _483_v1;
          _nw106[(new BigNumber(2)).toNumber()] = _483_v1;
          _nw106[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("htgefcx");
          _nw106[(new BigNumber(4)).toNumber()] = ((_487_v3)[_module.__default.safeIndex(new BigNumber(398), new BigNumber((_487_v3).length))]);
          _nw106[(new BigNumber(5)).toNumber()] = _483_v1;
          _nw106[(new BigNumber(6)).toNumber()] = _483_v1;
          _nw106[(new BigNumber(7)).toNumber()] = (((_593_v88).contains((_this).f27)) ? ((_593_v88).get((_this).f27)) : (_module.__default.fm15((_this).f27, (_this).f27, (_dafny.ZERO).minus((_this).f27), globalState)));
          _nw106[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_483_v1, _483_v1);
          _594_v89 = _nw106;
          let _index98 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_594_v89).length));
          (_594_v89)[_index98] = _dafny.Seq.Concat(_483_v1, _483_v1);
          let _index99 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_594_v89).length));
          let _rhs123 = function (_pat_let8_0) {
            return function (_595_dt__update__tmp_h2) {
              return function (_pat_let9_0) {
                return function (_596_dt__update_hcf15_h0) {
                  return _module.D5.create_DC10((_595_dt__update__tmp_h2).dtor_cf14, _596_dt__update_hcf15_h0);
                }(_pat_let9_0);
              }((_this).f28);
            }(_pat_let8_0);
          }(_592_v87);
          let _rhs124 = (_this).f27;
          let _rhs125 = p0;
          let _rhs126 = _dafny.Seq.Concat(_483_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-798)), function (_597_i6) {
            return new _dafny.CodePoint('f'.codePointAt(0));
          }));
          let _lhs114 = globalState;
          let _lhs115 = globalState;
          let _lhs116 = _594_v89;
          let _lhs117 = _module.__default.safeIndex(new BigNumber(367), new BigNumber((_594_v89).length));
          _592_v87 = _rhs123;
          _lhs114.f10 = _rhs124;
          _lhs115.f18 = _rhs125;
          _lhs116[_lhs117] = _rhs126;
          let _598_v90;
          _598_v90 = _dafny.Map.Empty.slice().updateUnsafe((p0) || (p1),_dafny.Seq.Concat(_493_v9, _493_v9));
          _598_v90 = (_598_v90).update(false, _493_v9);
        } else {
          let _599_v91;
          let _nw107 = Array((new BigNumber(27)).toNumber());
          _nw107[(_dafny.ZERO).toNumber()] = _dafny.Seq.of(p1, p0, p1);
          _nw107[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_493_v9, _dafny.Seq.of(p1, p1));
          _nw107[(new BigNumber(2)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(3)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(4)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(5)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(6)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(7)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_493_v9, _module.__default.fm7(globalState));
          _nw107[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(p0, (_this).f28, p0, p0);
          _nw107[(new BigNumber(10)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(11)).toNumber()] = ((p1) ? (_493_v9) : (_493_v9));
          _nw107[(new BigNumber(12)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(13)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(14)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(15)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(16)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(17)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(p1), _493_v9);
          _nw107[(new BigNumber(18)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(19)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(20)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(21)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(22)).toNumber()] = _dafny.Seq.Concat(_493_v9, _493_v9);
          _nw107[(new BigNumber(23)).toNumber()] = _dafny.Seq.Concat(_493_v9, _493_v9);
          _nw107[(new BigNumber(24)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(25)).toNumber()] = _493_v9;
          _nw107[(new BigNumber(26)).toNumber()] = _dafny.Seq.of(p0);
          _599_v91 = _nw107;
          let _index100 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_599_v91).length));
          (_599_v91)[_index100] = _493_v9;
          let _index101 = _module.__default.safeIndex(new BigNumber(27), new BigNumber((_599_v91).length));
          (_599_v91)[_index101] = _493_v9;
          (globalState).f18 = p1;
          let _600_v92;
          _600_v92 = _dafny.Map.Empty.slice().updateUnsafe(!(p0),_dafny.Seq.update(_483_v1, _module.__default.safeIndex(new BigNumber(-629), new BigNumber((_483_v1).length)), _482_v0));
          _600_v92 = (_600_v92).update(p1, _dafny.Seq.Concat(_483_v1, _483_v1));
          (globalState).f0 = ((!((_this).f28)) ? (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("html"), _dafny.Seq.UnicodeFromString("i"))).length)) : ((new BigNumber((_dafny.Seq.of(p0, (_this).f28)).length)).multipliedBy(new BigNumber(86))));
          (globalState).f18 = !((_this).f28);
        }
      }
      let _601_v93;
      _601_v93 = _dafny.Seq.of(_dafny.Seq.UnicodeFromString("fyp"), _483_v1);
      let _602_v94;
      _602_v94 = _module.D5.create_DC10(p1, false);
      if (!(!_dafny.areEqual(_dafny.Seq.update(_dafny.Seq.update(_dafny.Seq.Concat(_483_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_603_v0) => function (_604_i7) {
        return _603_v0;
      })(_482_v0))), _module.__default.safeIndex((_this).f27, new BigNumber((_dafny.Seq.Concat(_483_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_605_v0) => function (_606_i7) {
        return _605_v0;
      })(_482_v0)))).length)), _482_v0), _module.__default.safeIndex(new BigNumber(((_601_v93)[_module.__default.safeIndex((_this).f27, new BigNumber((_601_v93).length))]).length), new BigNumber((_dafny.Seq.update(_dafny.Seq.Concat(_483_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_607_v0) => function (_608_i7) {
        return _607_v0;
      })(_482_v0))), _module.__default.safeIndex((_this).f27, new BigNumber((_dafny.Seq.Concat(_483_v1, _dafny.Seq.Create(_module.__default.abs(new BigNumber(280)), ((_609_v0) => function (_610_i7) {
        return _609_v0;
      })(_482_v0)))).length)), _482_v0)).length)), _482_v0), _dafny.Seq.update((((_602_v94).dtor_cf15) ? ((_483_v1)) : (_483_v1)), _module.__default.safeIndex((_this).f27, new BigNumber(((((_602_v94).dtor_cf15) ? ((_483_v1)) : (_483_v1))).length)), _482_v0)))) {
        (globalState).f15 = (_this).f27;
        (globalState).f18 = false;
        let _611_v95;
        _611_v95 = _dafny.Seq.of((_this).f27);
        let _612_v96;
        let _nw108 = Array((new BigNumber(15)).toNumber());
        _nw108[(_dafny.ZERO).toNumber()] = p1;
        _nw108[(_dafny.ONE).toNumber()] = (_this).f28;
        _nw108[(new BigNumber(2)).toNumber()] = _module.__default.fm0((_this).f27, p1, p1, (_this).f27, globalState);
        _nw108[(new BigNumber(3)).toNumber()] = p1;
        _nw108[(new BigNumber(4)).toNumber()] = p0;
        _nw108[(new BigNumber(5)).toNumber()] = !((_this).f28);
        _nw108[(new BigNumber(6)).toNumber()] = p1;
        _nw108[(new BigNumber(7)).toNumber()] = (_this).f28;
        _nw108[(new BigNumber(8)).toNumber()] = p1;
        _nw108[(new BigNumber(9)).toNumber()] = p0;
        _nw108[(new BigNumber(10)).toNumber()] = (_this).f28;
        _nw108[(new BigNumber(11)).toNumber()] = !(p1);
        _nw108[(new BigNumber(12)).toNumber()] = (_this).f28;
        _nw108[(new BigNumber(13)).toNumber()] = p1;
        _nw108[(new BigNumber(14)).toNumber()] = (_602_v94).dtor_cf15;
        _612_v96 = _nw108;
        let _613_v97;
        _613_v97 = _dafny.Map.Empty.slice().updateUnsafe(_612_v96,p1);
        let _614_v98;
        _614_v98 = _dafny.Map.Empty.slice().updateUnsafe(((((_490_v6).contains((_this).f28)) ? ((_490_v6).get((_this).f28)) : ((_this).f27))).minus(new BigNumber((_611_v95).length)),(_613_v97).contains(_612_v96));
        let _615_v99;
        let _nw109 = Array((new BigNumber(20)).toNumber()).fill(_dafny.ZERO);
        _615_v99 = _nw109;
        let _616_v100;
        _616_v100 = _dafny.Map.Empty.slice().updateUnsafe(_615_v99,_module.__default.fm0((_this).f27, !(false), (_this).f28, (_this).f27, globalState));
        _614_v98 = (_614_v98).update((_this).f27, ((_dafny.ZERO).minus((_this).f27)).isLessThan(new BigNumber((_616_v100).length)));
        let _rhs127 = new _dafny.CodePoint('b'.codePointAt(0));
        let _rhs128 = _483_v1;
        let _lhs118 = globalState;
        _482_v0 = _rhs127;
        _lhs118.f3 = _rhs128;
        if (p1) {
          (globalState).f11 = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_module.__default.fm15(new BigNumber((_483_v1).length), (_this).f27, _module.__default.fm5(new BigNumber((_611_v95).length), (_this).f27, globalState), globalState)).length), (((_490_v6).contains(p0)) ? ((_490_v6).get(p0)) : ((_this).f27))));
          let _617_v101;
          let _nw110 = Array((new BigNumber(17)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _617_v101 = _nw110;
          let _index102 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_617_v101).length));
          (_617_v101)[_index102] = _483_v1;
          let _index103 = _module.__default.safeIndex(new BigNumber(63), new BigNumber((_617_v101).length));
          (_617_v101)[_index103] = _483_v1;
          (globalState).f10 = (_dafny.ZERO).minus((_this).f27);
          (globalState).f18 = false;
          let _618_v102;
          _618_v102 = _dafny.MultiSet.fromElements(_492_v8, _493_v9, _492_v8);
          let _619_v103;
          _619_v103 = _dafny.MultiSet.fromElements((_this).f27, new BigNumber((_618_v102).cardinality()));
          _619_v103 = _619_v103;
        } else {
          (globalState).f18 = (p1) === (p1);
          let _index104 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_615_v99).length));
          (_615_v99)[_index104] = new BigNumber(213);
          let _620_v104;
          _620_v104 = _dafny.MultiSet.fromElements((_this).f27);
          let _index105 = _module.__default.safeIndex(new BigNumber(854), new BigNumber((_615_v99).length));
          (_615_v99)[_index105] = (((_this).f28) ? ((_this).f27) : (((p1) ? (new BigNumber((_620_v104).cardinality())) : ((_this).f27))));
          (globalState).f2 = _module.__default.safeDivisionInt(((_615_v99)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_615_v99).length))]).minus((_dafny.ZERO).minus((_615_v99)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_615_v99).length))])), _module.__default.safeModuloInt((_615_v99)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_615_v99).length))], new BigNumber(588)));
          let _621_v105;
          _621_v105 = _module.D3.create_DC6(new BigNumber((_483_v1).length), (_615_v99)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_615_v99).length))], p0);
          let _622_v106;
          _622_v106 = _module.D9.create_DC16(_dafny.Map.Empty.slice().updateUnsafe((_this).f28,_621_v105));
          _622_v106 = _622_v106;
          let _623_v107;
          _623_v107 = _module.D10.create_DC20(p1, _module.__default.fm5((_this).f27, (_615_v99)[_module.__default.safeIndex(new BigNumber(854), new BigNumber((_615_v99).length))], globalState));
          _623_v107 = _623_v107;
        }
      } else {
        _517_v28 = (_517_v28).Merge((_517_v28).Merge(_517_v28));
        let _624_v108;
        _624_v108 = _module.D2.create_DC3((_this).f27, (_this).f27);
        let _625_v109;
        _625_v109 = _dafny.Map.Empty.slice().updateUnsafe(_624_v108,(_dafny.ZERO).minus((_this).f27));
        _625_v109 = (_625_v109).update(_624_v108, new BigNumber(446));
        let _626_v110;
        let _nw111 = Array((new BigNumber(20)).toNumber()).fill(false);
        _626_v110 = _nw111;
        let _index106 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_626_v110).length));
        (_626_v110)[_index106] = p1;
        let _index107 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_626_v110).length));
        (_626_v110)[_index107] = p0;
        let _627_v111;
        _627_v111 = _dafny.Seq.of(new BigNumber((_dafny.Seq.of(!((_this).f28))).length));
        (globalState).f0 = (_627_v111)[_module.__default.safeIndex(new BigNumber((_493_v9).length), new BigNumber((_627_v111).length))];
        let _628_v112;
        _628_v112 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f27);
        (globalState).f15 = (new BigNumber((_628_v112).length)).multipliedBy((_this).f27);
      }
      let _629_v113;
      let _nw112 = new _module.C1();
      _nw112.__ctor(_482_v0, ((_this).f27).minus((_this).f27));
      _629_v113 = _nw112;
      let _630_i8;
      _630_i8 = _dafny.ZERO;
      L2: {
        while (((_490_v6).Difference(((_dafny.MultiSet.fromElements(p1)).update(p1, _module.__default.abs((_this).f27))).update(p0, _module.__default.abs(new BigNumber(103))))).IsDisjointFrom(_490_v6)) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_630_i8)) {
              break L2;
            }
            _630_i8 = (_630_i8).plus(_dafny.ONE);
            (globalState).f18 = p1;
            (globalState).f15 = new BigNumber((_493_v9).length);
            (globalState).f18 = ((_490_v6).Difference(_490_v6)).IsDisjointFrom(_490_v6);
            (globalState).f10 = (_this).f27;
          }
        }
      }
      r0 = function (_source11) {
        if (_source11.is_DC3) {
          let _631___mcc_h7 = (_source11).cf3;
          let _632___mcc_h8 = (_source11).cf4;
          let _633_cf4 = _632___mcc_h8;
          let _634_cf3 = _631___mcc_h7;
          return (_this).f27;
        } else {
          let _635___mcc_h9 = (_source11).cf2;
          let _636_cf2 = _635___mcc_h9;
          return (_dafny.ZERO).minus((_this).f27);
        }
      }(_module.D2.create_DC2((_this).f28));
      let _637_v114;
      _637_v114 = _dafny.MultiSet.fromElements(_482_v0);
      r1 = (_dafny.ZERO).minus((((_637_v114).contains(_482_v0)) ? ((_637_v114).get(_482_v0)) : ((_this).f27)));
      return [r0, r1];
    }
    get f27() {
      let _this = this;
      return _this._f27;
    };
    get f28() {
      let _this = this;
      return _this._f28;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f25 = _dafny.ZERO;
      this._f24 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f31 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    set f25(value) {
      let _this = this;
      _this._f25 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    __ctor(f31, f24, f25) {
      let _this = this;
      (_this).f31 = f31;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("pcxk"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(893)), function (_638_i0) {
        return new _dafny.CodePoint('g'.codePointAt(0));
      }));
    };
    fm18(globalState) {
      let _this = this;
      return !(_this.f31);
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f25 = _dafny.ZERO;
      this._f24 = new _dafny.CodePoint('D'.codePointAt(0));
      this.f29 = _dafny.Map.Empty;
      this._f30 = false;
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    set f25(value) {
      let _this = this;
      _this._f25 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    __ctor(f29, f30, f24, f25) {
      let _this = this;
      (_this).f29 = f29;
      (_this)._f30 = f30;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("if"), _dafny.Seq.UnicodeFromString("kaxkef")), _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("e"), _dafny.Seq.UnicodeFromString("oqanv")));
    };
    m5(p0, p1, p2, globalState) {
      let _this = this;
      (globalState).f11 = p2;
      let _639_i0;
      _639_i0 = _dafny.ZERO;
      L3: {
        while (p1) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_639_i0)) {
              break L3;
            }
            _639_i0 = (_639_i0).plus(_dafny.ONE);
            let _640_v0;
            _640_v0 = _dafny.Map.Empty.slice().updateUnsafe(_module.D2.create_DC2(p1),p2);
            let _641_v1;
            _641_v1 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_dafny.ZERO).minus((_module.__default.fm16(_640_v0, globalState)).dtor_cf29));
            let _642_v2;
            _642_v2 = _dafny.Seq.UnicodeFromString("aqdbq");
            let _643_v3;
            _643_v3 = _dafny.Map.Empty.slice().updateUnsafe(_642_v2,_module.__default.fm0(p2, (_this).f30, p1, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(644)), function (_644_i1) {
              return _this.f25;
            })).length), globalState));
            _641_v1 = (_641_v1).update(new BigNumber(((_643_v3).Merge(_643_v3)).length), p2);
            let _645_v4;
            let _nw113 = Array((new BigNumber(24)).toNumber());
            _645_v4 = _nw113;
            let _646_v5;
            _646_v5 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,_645_v4);
            _645_v4 = (((_646_v5).contains((_this).f30)) ? ((_646_v5).get((_this).f30)) : (_645_v4));
            _641_v1 = (_641_v1).update(_this.f25, _module.__default.safeDivisionInt(_this.f25, (_dafny.ZERO).minus(p2)));
            let _647_v6;
            _647_v6 = _dafny.Seq.of((_this).f30);
            let _648_v7;
            let _init17 = ((_649_v1) => function (_650_i2) {
              return _649_v1;
            })(_641_v1);
            let _nw114 = Array((new BigNumber(16)).toNumber());
            for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw114.length); _i0_17++) {
              _nw114[_i0_17] = _init17(new BigNumber(_i0_17));
            }
            _648_v7 = _nw114;
            let _651_v8;
            _651_v8 = _dafny.MultiSet.fromElements(_648_v7);
            let _652_v9;
            _652_v9 = _dafny.Set.fromElements(!((_this).f30), (_this).f30);
            let _653_v10;
            _653_v10 = _dafny.MultiSet.fromElements(_this.f25);
            let _654_v11;
            _654_v11 = _dafny.Set.fromElements(new BigNumber(489), new BigNumber((_642_v2).length));
            let _655_v12;
            _655_v12 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
            let _656_v13;
            _656_v13 = _dafny.Seq.of(_655_v12);
            let _657_v14;
            _657_v14 = _dafny.Map.Empty.slice().updateUnsafe((_656_v13)[_module.__default.safeIndex(_this.f25, new BigNumber((_656_v13).length))],(_this).f24);
            let _658_v15;
            _658_v15 = _dafny.Seq.of(p2);
            let _659_v16;
            _659_v16 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,(_this).f24);
            let _660_v17;
            let _nw115 = Array((new BigNumber(27)).toNumber());
            _nw115[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(_this.f25, _this.f25);
            _nw115[(_dafny.ONE).toNumber()] = p2;
            _nw115[(new BigNumber(2)).toNumber()] = new BigNumber((_642_v2).length);
            _nw115[(new BigNumber(3)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_647_v6).length), _this.f25);
            _nw115[(new BigNumber(4)).toNumber()] = _this.f25;
            _nw115[(new BigNumber(5)).toNumber()] = new BigNumber((_642_v2).length);
            _nw115[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt((((_651_v8).contains(_648_v7)) ? ((_651_v8).get(_648_v7)) : (new BigNumber((_652_v9).length))), new BigNumber(108));
            _nw115[(new BigNumber(7)).toNumber()] = _this.f25;
            _nw115[(new BigNumber(8)).toNumber()] = (p2).multipliedBy(p2);
            _nw115[(new BigNumber(9)).toNumber()] = (((_653_v10).contains(p2)) ? ((_653_v10).get(p2)) : (p2));
            _nw115[(new BigNumber(10)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_654_v11).length), p2);
            _nw115[(new BigNumber(11)).toNumber()] = new BigNumber(-441);
            _nw115[(new BigNumber(12)).toNumber()] = new BigNumber((_657_v14).length);
            _nw115[(new BigNumber(13)).toNumber()] = p2;
            _nw115[(new BigNumber(14)).toNumber()] = new BigNumber(141);
            _nw115[(new BigNumber(15)).toNumber()] = (_dafny.ZERO).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(592)), function (_661_i3) {
              return new BigNumber(499);
            })).length)));
            _nw115[(new BigNumber(16)).toNumber()] = _this.f25;
            _nw115[(new BigNumber(17)).toNumber()] = p2;
            _nw115[(new BigNumber(18)).toNumber()] = _this.f25;
            _nw115[(new BigNumber(19)).toNumber()] = (_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), function (_662_i4) {
              return new _dafny.CodePoint('m'.codePointAt(0));
            }), _module.__default.safeIndex(_this.f25, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), function (_663_i4) {
              return new _dafny.CodePoint('m'.codePointAt(0));
            })).length)), (_this).f24)).length), new BigNumber(257)));
            _nw115[(new BigNumber(20)).toNumber()] = p2;
            _nw115[(new BigNumber(21)).toNumber()] = (((_this).f30) ? (p2) : (_module.__default.fm17((_this).f30, true, p1, globalState)));
            _nw115[(new BigNumber(22)).toNumber()] = new BigNumber((_658_v15).length);
            _nw115[(new BigNumber(23)).toNumber()] = (_dafny.ZERO).minus(p2);
            _nw115[(new BigNumber(24)).toNumber()] = (new BigNumber((_659_v16).length)).minus(new BigNumber(624));
            _nw115[(new BigNumber(25)).toNumber()] = (new BigNumber(286)).plus(_module.__default.fm17(p1, (_this).f30, p1, globalState));
            _nw115[(new BigNumber(26)).toNumber()] = _this.f25;
            _660_v17 = _nw115;
            let _index108 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_660_v17).length));
            (_660_v17)[_index108] = p2;
            let _index109 = _module.__default.safeIndex(new BigNumber(280), new BigNumber((_660_v17).length));
            (_660_v17)[_index109] = _this.f25;
          }
        }
      }
      let _664_v18;
      _664_v18 = _dafny.Map.Empty.slice().updateUnsafe((_this).f30,p1);
      let _665_v19;
      let _nw116 = new _module.C3();
      _nw116.__ctor(p1, (_this).f24, new BigNumber((_664_v18).length));
      _665_v19 = _nw116;
      let _666_v20;
      _666_v20 = _dafny.Seq.of(_665_v19);
      (globalState).f11 = (_dafny.ZERO).minus((_this.f25).multipliedBy(new BigNumber((_dafny.Seq.of(_666_v20, _666_v20, _666_v20, _666_v20)).length)));
      let _667_v21;
      _667_v21 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(682),p1);
      let _668_v22;
      let _nw117 = Array((new BigNumber(3)).toNumber());
      _nw117[(_dafny.ZERO).toNumber()] = ((((_667_v21).contains(_this.f25)) ? ((_667_v21).get(_this.f25)) : (p1))) && (p1);
      _nw117[(_dafny.ONE).toNumber()] = (((_664_v18).contains((_this).f30)) ? ((_664_v18).get((_this).f30)) : ((_this).f30));
      _nw117[(new BigNumber(2)).toNumber()] = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("uxddo"), _dafny.Seq.UnicodeFromString("teel"));
      _668_v22 = _nw117;
      let _index110 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_668_v22).length));
      (_668_v22)[_index110] = p1;
      let _669_v23;
      _669_v23 = _dafny.Map.Empty.slice().updateUnsafe(p2,p2);
      let _index111 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_668_v22).length));
      (_668_v22)[_index111] = _module.__default.fm0(new BigNumber((_669_v23).length), p1, p1, _665_v19.f25, globalState);
      _667_v21 = (_667_v21).update(_665_v19.f25, ((_dafny.ZERO).minus(_665_v19.f25)).isLessThanOrEqualTo(_this.f25));
      let _670_v24;
      let _nw118 = new _module.C1();
      _nw118.__ctor((_665_v19).f24, _module.__default.safeDivisionInt(_665_v19.f25, new BigNumber((_664_v18).length)));
      _670_v24 = _nw118;
      return;
    }
    get f30() {
      let _this = this;
      return _this._f30;
    };
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f25 = _dafny.ZERO;
      this._f24 = new _dafny.CodePoint('D'.codePointAt(0));
      this._f26 = false;
    }
    _parentTraits() {
      return [_module.T1, _module.T0];
    }
    get f25() {
      let _this = this;
      return _this._f25;
    };
    set f25(value) {
      let _this = this;
      _this._f25 = value;
    };
    get f24() {
      let _this = this;
      return _this._f24;
    };
    __ctor(f26, f24, f25) {
      let _this = this;
      (_this)._f26 = f26;
      (_this)._f24 = f24;
      (_this)._f25 = f25;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("enhmpxpn"), (((_this).f26) ? (_dafny.Seq.UnicodeFromString("bob")) : (_dafny.Seq.UnicodeFromString("ihqo"))));
    };
    m1(p0, p1, p2, globalState) {
      let _this = this;
      let r0 = false;
      let r1 = _dafny.Seq.of();
      let r2 = false;
      let _671_v0;
      let _nw119 = new _module.C2();
      _nw119.__ctor(_this.f25, p1);
      _671_v0 = _nw119;
      let _672_i0;
      _672_i0 = _dafny.ZERO;
      L4: {
        let _pat_let_tv4 = _671_v0;
        let _pat_let_tv5 = _671_v0;
        let _pat_let_tv6 = p1;
        let _pat_let_tv7 = _671_v0;
        let _pat_let_tv8 = _671_v0;
        while (!(function (_source12) {
          if (_source12.is_DC19) {
            let _702___mcc_h0 = (_source12).cf27;
            let _703_cf27 = _702___mcc_h0;
            return !(!(_dafny.Set.fromElements((_pat_let_tv4).f28)).equals(_dafny.Set.fromElements((_pat_let_tv5).f28, _pat_let_tv6)));
          } else if (_source12.is_DC20) {
            let _704___mcc_h1 = (_source12).cf28;
            let _705___mcc_h2 = (_source12).cf29;
            let _706_cf29 = _705___mcc_h2;
            let _707_cf28 = _704___mcc_h1;
            return !(true);
          } else if (_source12.is_DC21) {
            let _708___mcc_h3 = (_source12).cf30;
            let _709___mcc_h4 = (_source12).cf31;
            let _710___mcc_h5 = (_source12).cf32;
            let _711_cf32 = _710___mcc_h5;
            let _712_cf31 = _709___mcc_h4;
            let _713_cf30 = _708___mcc_h3;
            return (_pat_let_tv7).f28;
          } else {
            let _714___mcc_h6 = (_source12).cf26;
            let _715_cf26 = _714___mcc_h6;
            return (_pat_let_tv8).f28;
          }
        }(_module.__default.fm14((_this).f26, globalState)))) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_672_i0)) {
              break L4;
            }
            _672_i0 = (_672_i0).plus(_dafny.ONE);
            let _673_v1;
            let _init18 = function (_674_i1) {
              return (_this).f26;
            };
            let _nw120 = Array((new BigNumber(22)).toNumber());
            for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw120.length); _i0_18++) {
              _nw120[_i0_18] = _init18(new BigNumber(_i0_18));
            }
            _673_v1 = _nw120;
            let _675_v2;
            _675_v2 = _dafny.Seq.of(_673_v1);
            let _676_v3;
            let _nw121 = new _module.C3();
            _nw121.__ctor((_671_v0).f28, (_this).f24, ((_671_v0).f27).minus((_dafny.ZERO).minus(_this.f25)));
            _676_v3 = _nw121;
            let _rhs129 = _676_v3.f25;
            let _rhs130 = !((_671_v0).f28);
            let _rhs131 = p1;
            let _rhs132 = _675_v2;
            let _rhs133 = _676_v3;
            let _lhs119 = globalState;
            let _lhs120 = globalState;
            _lhs119.f0 = _rhs129;
            _lhs120.f18 = _rhs130;
            r0 = _rhs131;
            _675_v2 = _rhs132;
            _676_v3 = _rhs133;
            r0 = false;
            let _677_v4;
            _677_v4 = _dafny.Seq.UnicodeFromString("vfafw");
            (globalState).f3 = _dafny.Seq.Concat(_dafny.Seq.Concat(_677_v4, _dafny.Seq.Create(_module.__default.abs(new BigNumber(-183)), ((_678_v3) => function (_679_i2) {
              return (_678_v3).f24;
            })(_676_v3))), (((_this).f26) ? (_677_v4) : (_677_v4)));
            if ((_671_v0).f28) {
              r2 = (_671_v0).f28;
              (globalState).f18 = !(((_671_v0).f27).isLessThan(((_671_v0).f27).minus((_dafny.ZERO).minus(_676_v3.f25))));
              (globalState).f0 = _module.__default.safeModuloInt((_671_v0).f27, (_671_v0).f27);
              (globalState).f22 = (_676_v3.f25).plus((_dafny.ZERO).minus(_module.__default.safeDivisionInt(new BigNumber(783), (_671_v0).f27)));
              let _680_v5;
              _680_v5 = _dafny.Map.Empty.slice().updateUnsafe((_671_v0).f28,p1);
              let _681_v6;
              let _nw122 = Array((new BigNumber(4)).toNumber());
              _681_v6 = _nw122;
              let _682_v7;
              _682_v7 = _dafny.Map.Empty.slice().updateUnsafe(_680_v5,_681_v6);
              let _683_v8;
              _683_v8 = _dafny.Seq.of(_681_v6);
              let _684_v9;
              let _nw123 = new _module.C0();
              _nw123.__ctor();
              _684_v9 = _nw123;
              let _685_v10;
              _685_v10 = _dafny.Seq.of(_684_v9);
              let _686_v11;
              _686_v11 = _dafny.Map.Empty.slice().updateUnsafe((_671_v0).f27,new BigNumber((_685_v10).length));
              let _687_v12;
              _687_v12 = _module.D11.create_DC22(_681_v6);
              let _688_v13;
              let _nw124 = Array((new BigNumber(25)).toNumber());
              _nw124[(_dafny.ZERO).toNumber()] = (_682_v7).Merge(_682_v7);
              _nw124[(_dafny.ONE).toNumber()] = (_682_v7).Merge(_682_v7);
              _nw124[(new BigNumber(2)).toNumber()] = (_dafny.Map.Empty.slice().updateUnsafe(_680_v5,_681_v6)).update(_680_v5, _681_v6);
              _nw124[(new BigNumber(3)).toNumber()] = ((false) ? (_682_v7) : (_682_v7));
              _nw124[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_680_v5,(_683_v8)[_module.__default.safeIndex(new BigNumber((_686_v11).length), new BigNumber((_683_v8).length))]);
              _nw124[(new BigNumber(5)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(6)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(7)).toNumber()] = (_682_v7).Merge(_682_v7);
              _nw124[(new BigNumber(8)).toNumber()] = (_682_v7).Merge(_dafny.Map.Empty.slice().updateUnsafe(_680_v5,_681_v6));
              _nw124[(new BigNumber(9)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(10)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(11)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(12)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(13)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(14)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(15)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(16)).toNumber()] = (_682_v7).Merge(_682_v7);
              _nw124[(new BigNumber(17)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(18)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(19)).toNumber()] = (_682_v7).update(_680_v5, (_687_v12).dtor_cf33);
              _nw124[(new BigNumber(20)).toNumber()] = (_682_v7).Merge(_682_v7);
              _nw124[(new BigNumber(21)).toNumber()] = (_682_v7).Merge((_682_v7).update(_680_v5, _681_v6));
              _nw124[(new BigNumber(22)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(23)).toNumber()] = _682_v7;
              _nw124[(new BigNumber(24)).toNumber()] = _682_v7;
              _688_v13 = _nw124;
              let _index112 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_688_v13).length));
              (_688_v13)[_index112] = (_dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm19((_671_v0).f28, globalState),_681_v6)).Merge(_682_v7);
              let _index113 = _module.__default.safeIndex(new BigNumber(347), new BigNumber((_688_v13).length));
              (_688_v13)[_index113] = (_682_v7).Merge(_682_v7);
            } else {
              let _689_v14;
              _689_v14 = _module.D4.create_DC8();
              let _690_v15;
              _690_v15 = _dafny.MultiSet.fromElements(_module.D4.create_DC8(), _689_v14);
              let _691_v16;
              _691_v16 = _dafny.Seq.of(_689_v14, _689_v14);
              let _rhs134 = p1;
              let _rhs135 = (_dafny.MultiSet.FromArray(_691_v16)).Union(_dafny.MultiSet.fromElements(_689_v14, _689_v14, _689_v14, _689_v14, _689_v14));
              let _rhs136 = (_this).f26;
              let _rhs137 = p1;
              let _lhs121 = globalState;
              let _lhs122 = globalState;
              _lhs121.f18 = _rhs134;
              _690_v15 = _rhs135;
              r2 = _rhs136;
              _lhs122.f18 = _rhs137;
              let _692_v17;
              let _init19 = ((_693_v0) => function (_694_i3) {
                return _module.__default.safeDivisionInt(_694_i3, (_dafny.ZERO).minus((_693_v0).f27));
              })(_671_v0);
              let _nw125 = Array((new BigNumber(27)).toNumber());
              for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw125.length); _i0_19++) {
                _nw125[_i0_19] = _init19(new BigNumber(_i0_19));
              }
              _692_v17 = _nw125;
              let _nw126 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
              _692_v17 = _nw126;
              let _695_v18;
              _695_v18 = _module.D10.create_DC19((_676_v3).f24);
              let _696_v19;
              _696_v19 = _module.D10.create_DC21(_676_v3.f25, (_this).f26, (_671_v0).f28);
              _695_v18 = ((!((_671_v0).f28)) ? (_695_v18) : ((((_696_v19).dtor_cf32) ? (function (_pat_let10_0) {
                return function (_697_dt__update__tmp_h0) {
                  return function (_pat_let11_0) {
                    return function (_698_dt__update_hcf27_h0) {
                      return _module.D10.create_DC19(_698_dt__update_hcf27_h0);
                    }(_pat_let11_0);
                  }(new _dafny.CodePoint('g'.codePointAt(0)));
                }(_pat_let10_0);
              }(_695_v18)) : (_module.__default.fm20(_676_v3.f25, true, globalState)))));
              let _699_v20;
              let _700_v21;
              let _out8;
              let _out9;
              let _outcollector1 = (_671_v0).m3((_this).f26, (_this).f26, globalState);
              _out8 = _outcollector1[0];
              _out9 = _outcollector1[1];
              _699_v20 = _out8;
              _700_v21 = _out9;
              let _701_v22;
              let _nw127 = new _module.C0();
              _nw127.__ctor();
              _701_v22 = _nw127;
              _701_v22 = _701_v22;
            }
          }
        }
      }
      let _716_v23;
      _716_v23 = _dafny.Seq.of(p1, (_this).f26, p1, p1, p1);
      let _717_v24;
      _717_v24 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,p1);
      let _718_v25;
      _718_v25 = _dafny.Seq.UnicodeFromString("pdxyd");
      let _719_v26;
      let _nw128 = Array((new BigNumber(20)).toNumber());
      _nw128[(_dafny.ZERO).toNumber()] = !(p1);
      _nw128[(_dafny.ONE).toNumber()] = !(((p1) ? ((_671_v0).f28) : (p1)));
      _nw128[(new BigNumber(2)).toNumber()] = (_this).f26;
      _nw128[(new BigNumber(3)).toNumber()] = (_this.f25).isLessThanOrEqualTo((_dafny.ZERO).minus(_this.f25));
      _nw128[(new BigNumber(4)).toNumber()] = !(!((_this.f25).isLessThan(_this.f25)));
      _nw128[(new BigNumber(5)).toNumber()] = (_this).f26;
      _nw128[(new BigNumber(6)).toNumber()] = (_this).f26;
      _nw128[(new BigNumber(7)).toNumber()] = (_this).f26;
      _nw128[(new BigNumber(8)).toNumber()] = (_this).f26;
      _nw128[(new BigNumber(9)).toNumber()] = p1;
      _nw128[(new BigNumber(10)).toNumber()] = p1;
      _nw128[(new BigNumber(11)).toNumber()] = ((_671_v0).f28) === (_module.__default.fm0((_671_v0).f27, p1, (_671_v0).f28, _this.f25, globalState));
      _nw128[(new BigNumber(12)).toNumber()] = ((_this).f26) || (p1);
      _nw128[(new BigNumber(13)).toNumber()] = (_716_v23)[_module.__default.safeIndex((_671_v0).f27, new BigNumber((_716_v23).length))];
      _nw128[(new BigNumber(14)).toNumber()] = !((_this).f26);
      _nw128[(new BigNumber(15)).toNumber()] = (_671_v0).f28;
      _nw128[(new BigNumber(16)).toNumber()] = (((_717_v24).contains(_this.f25)) ? ((_717_v24).get(_this.f25)) : ((_this).f26));
      _nw128[(new BigNumber(17)).toNumber()] = p1;
      _nw128[(new BigNumber(18)).toNumber()] = (new BigNumber((_718_v25).length)).isLessThanOrEqualTo(_this.f25);
      _nw128[(new BigNumber(19)).toNumber()] = (_671_v0).f28;
      _719_v26 = _nw128;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_719_v26).length))) {
        let _720_i4 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_720_i4)) && ((_720_i4).isLessThan(new BigNumber((_719_v26).length))))) {
          (_719_v26)[(_720_i4)] = ((((_671_v0).f28) ? (_dafny.MultiSet.fromElements((_671_v0).f28)) : (_dafny.MultiSet.fromElements((_671_v0).f28, !(p1), (_671_v0).f28)))).IsProperSubsetOf((_dafny.MultiSet.fromElements((_671_v0).f28)).Intersect(_dafny.MultiSet.FromArray(_716_v23)));
        }
      }
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_719_v26).length))) {
        let _721_i5 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_721_i5)) && ((_721_i5).isLessThan(new BigNumber((_719_v26).length))))) {
          (_719_v26)[(_721_i5)] = (_this).f26;
        }
      }
      let _722_v27;
      _722_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this.f25).multipliedBy((_671_v0).f27),_671_v0);
      _671_v0 = (((_722_v27).contains((_671_v0).f27)) ? ((_722_v27).get((_671_v0).f27)) : (_671_v0));
      if (p1) {
        let _723_v28;
        _723_v28 = new _dafny.CodePoint('d'.codePointAt(0));
        _723_v28 = (_this).f24;
        let _724_v29;
        _724_v29 = _module.D7.create_DC14((_this).f26, (_dafny.ZERO).minus(_this.f25), (_671_v0).f27);
        let _725_v30;
        _725_v30 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_716_v23),(_724_v29).dtor_cf22);
        let _726_v31;
        _726_v31 = _dafny.Map.Empty.slice().updateUnsafe((((_717_v24).contains((_671_v0).f27)) ? ((_717_v24).get((_671_v0).f27)) : (true)),!((_671_v0).f28));
        let _727_v32;
        let _nw129 = new _module.C4();
        _nw129.__ctor(_725_v30, (((_726_v31).contains((_this).f26)) ? ((_726_v31).get((_this).f26)) : (p1)), new _dafny.CodePoint('f'.codePointAt(0)), _module.__default.safeDivisionInt(new BigNumber(656), _this.f25));
        _727_v32 = _nw129;
        _727_v32 = _727_v32;
        if (!(false)) {
          (globalState).f11 = _this.f25;
          let _728_v33;
          _728_v33 = _module.D7.create_DC13(_717_v24);
          let _729_v34;
          _729_v34 = _dafny.MultiSet.fromElements((((_this).f26) ? (_728_v33) : (_728_v33)));
          let _730_v35;
          let _init20 = function (_731_i6) {
            return (_731_i6).minus(new BigNumber(-813));
          };
          let _nw130 = Array((new BigNumber(10)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw130.length); _i0_20++) {
            _nw130[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _730_v35 = _nw130;
          let _732_v36;
          _732_v36 = _module.D6.create_DC12(p1, (new BigNumber((_dafny.Seq.update(_dafny.Seq.UnicodeFromString("pshrqd"), _module.__default.safeIndex((_671_v0).f27, new BigNumber((_dafny.Seq.UnicodeFromString("pshrqd")).length)), _723_v28)).length)).plus((_671_v0).f27));
          let _rhs138 = _dafny.MultiSet.FromArray(_module.__default.fm21(_module.__default.safeModuloInt((_671_v0).f27, (_671_v0).f27), _module.__default.safeModuloInt((_671_v0).f27, (_671_v0).f27), globalState));
          let _rhs139 = _730_v35;
          let _rhs140 = _732_v36;
          _729_v34 = _rhs138;
          _730_v35 = _rhs139;
          _732_v36 = _rhs140;
          (globalState).f3 = _dafny.Seq.Concat(_dafny.Seq.Concat(_718_v25, _718_v25), _718_v25);
          let _733_v37;
          _733_v37 = _module.D10.create_DC21(new BigNumber(((_726_v31).update(p1, (_671_v0).f28)).length), (_671_v0).f28, (_671_v0).f28);
          let _734_v38;
          _734_v38 = _dafny.Set.fromElements(_733_v37);
          _734_v38 = _734_v38;
          (globalState).f18 = _dafny.Seq.IsPrefixOf(_716_v23, _module.__default.fm7(globalState));
        } else {
          let _735_v39;
          let _nw131 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _735_v39 = _nw131;
          let _736_v40;
          _736_v40 = _dafny.MultiSet.fromElements((_671_v0).f28);
          let _index114 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_735_v39).length));
          (_735_v39)[_index114] = _module.__default.safeModuloInt(new BigNumber((_736_v40).cardinality()), new BigNumber(515));
          let _index115 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_735_v39).length));
          (_735_v39)[_index115] = new BigNumber(460);
          let _737_v41;
          let _nw132 = new _module.C1();
          _nw132.__ctor((_this).f24, (new BigNumber((_716_v23).length)).multipliedBy((_671_v0).f27));
          _737_v41 = _nw132;
          let _738_v42;
          _738_v42 = _dafny.Map.Empty.slice().updateUnsafe((_671_v0).f27,(_this).f24);
          let _739_v43;
          _739_v43 = _dafny.Set.fromElements(new BigNumber((_718_v25).length), (_671_v0).f27);
          let _nw133 = new _module.C4();
          _nw133.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.FromArray(_dafny.Seq.of((_671_v0).f28, (_671_v0).f28)),_737_v41.f25), (_671_v0).f28, (((_738_v42).contains(new BigNumber((_dafny.Seq.UnicodeFromString("kanvy")).length))) ? ((_738_v42).get(new BigNumber((_dafny.Seq.UnicodeFromString("kanvy")).length))) : (_723_v28)), (new BigNumber((_739_v43).length)).plus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(408)), ((_740_v28) => function (_741_i7) {
            return _740_v28;
          })(_723_v28))).length)));
          _737_v41 = _nw133;
          let _index116 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_735_v39).length));
          (_735_v39)[_index116] = (_671_v0).f27;
          let _742_v44;
          let _init21 = ((_743_v41, _744_v39, _745_v0) => function (_746_i8) {
            return _module.D3.create_DC6(_743_v41.f25, (_744_v39)[_module.__default.safeIndex(new BigNumber(824), new BigNumber((_744_v39).length))], (_745_v0).f28);
          })(_737_v41, _735_v39, _671_v0);
          let _nw134 = Array((_dafny.ONE).toNumber());
          for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw134.length); _i0_21++) {
            _nw134[_i0_21] = _init21(new BigNumber(_i0_21));
          }
          _742_v44 = _nw134;
          let _index117 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_742_v44).length));
          (_742_v44)[_index117] = _module.D3.create_DC6(_module.__default.fm17((_671_v0).f28, (_727_v32).f30, p1, globalState), _this.f25, _module.__default.fm0(_this.f25, !((_727_v32).f30), false, _this.f25, globalState));
          let _747_v45;
          _747_v45 = _dafny.Map.Empty.slice().updateUnsafe((_671_v0).f28,(_735_v39)[_module.__default.safeIndex(new BigNumber(824), new BigNumber((_735_v39).length))]);
          let _748_v46;
          _748_v46 = _module.D3.create_DC6((_671_v0).f27, _737_v41.f25, p1);
          let _index118 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_742_v44).length));
          let _rhs141 = _747_v45;
          let _rhs142 = _748_v46;
          let _lhs123 = globalState;
          let _lhs124 = _742_v44;
          let _lhs125 = _module.__default.safeIndex(new BigNumber(811), new BigNumber((_742_v44).length));
          _lhs123.f13 = _rhs141;
          _lhs124[_lhs125] = _rhs142;
          let _749_v47;
          let _nw135 = Array((new BigNumber(24)).toNumber());
          _749_v47 = _nw135;
          let _750_v48;
          _750_v48 = _module.D11.create_DC22(_749_v47);
          let _751_v49;
          _751_v49 = _dafny.Seq.of(_750_v48, _750_v48);
          let _752_v50;
          _752_v50 = _dafny.Seq.of((_dafny.ZERO).minus((_735_v39)[_module.__default.safeIndex(new BigNumber(824), new BigNumber((_735_v39).length))]));
          let _index119 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_735_v39).length));
          let _rhs143 = _dafny.Seq.Concat(_751_v49, _dafny.Seq.update(_751_v49, _module.__default.safeIndex(new BigNumber((_752_v50).length), new BigNumber((_751_v49).length)), _750_v48));
          let _rhs144 = ((_671_v0).f27).multipliedBy(_737_v41.f25);
          let _rhs145 = _719_v26;
          let _lhs126 = _735_v39;
          let _lhs127 = _module.__default.safeIndex(new BigNumber(824), new BigNumber((_735_v39).length));
          _751_v49 = _rhs143;
          _lhs126[_lhs127] = _rhs144;
          _719_v26 = _rhs145;
        }
        let _753_v51;
        let _nw136 = new _module.C1();
        _nw136.__ctor(new _dafny.CodePoint('v'.codePointAt(0)), new BigNumber((_718_v25).length));
        _753_v51 = _nw136;
        let _754_v53;
        _754_v53 = _dafny.MultiSet.fromElements((_671_v0).f28, (_671_v0).f28, !((_671_v0).f28), !(true), (((_726_v31).contains((_671_v0).f28)) ? ((_726_v31).get((_671_v0).f28)) : (!((_this).f26))));
        let _rhs146 = (_dafny.ZERO).minus((_671_v0).f27);
        let _rhs147 = !(p1);
        let _rhs148 = _719_v26;
        let _rhs149 = ((_671_v0).fm4(function () {
          let _coll34 = new _dafny.Set();
          for (const _compr_34 of _dafny.IntegerRange(new BigNumber(93), new BigNumber(576))) {
            let _755_v52 = _compr_34;
            if (((new BigNumber(93)).isLessThanOrEqualTo(_755_v52)) && ((_755_v52).isLessThan(new BigNumber(576)))) {
              _coll34.add(_module.__default.safeModuloInt(_755_v52, (_671_v0).f27));
            }
          }
          return _coll34;
        }(), new BigNumber((_718_v25).length), (_671_v0).f28, globalState)).IsDisjointFrom(_754_v53);
        let _lhs128 = globalState;
        let _lhs129 = globalState;
        _lhs128.f22 = _rhs146;
        r0 = _rhs147;
        _719_v26 = _rhs148;
        _lhs129.f18 = _rhs149;
      } else {
        r1 = _716_v23;
        (globalState).f2 = (((_671_v0).f27).minus(_module.__default.fm17(!(true), (_this).f26, p1, globalState))).plus((_671_v0).f27);
        let _756_v54;
        _756_v54 = _dafny.Map.Empty.slice().updateUnsafe((_671_v0).f28,new BigNumber((_718_v25).length));
        let _757_v55;
        _757_v55 = _dafny.Map.Empty.slice().updateUnsafe(!(_module.__default.fm0(new BigNumber(-691), (_716_v23)[_module.__default.safeIndex(new BigNumber(-920), new BigNumber((_716_v23).length))], (_this).f26, (((_756_v54).contains(true)) ? ((_756_v54).get(true)) : ((_dafny.ZERO).minus((_671_v0).f27))), globalState)),(_671_v0).f27);
        (globalState).f13 = _756_v54;
        if (p1) {
          let _758_v56;
          _758_v56 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_671_v0).f27);
          let _759_v57;
          _759_v57 = _dafny.MultiSet.fromElements(_716_v23, _716_v23);
          let _760_v58;
          _760_v58 = _dafny.MultiSet.fromElements((_671_v0).f28, (_671_v0).f28, p1, _module.__default.fm0((_671_v0).f27, (_671_v0).f28, (_671_v0).f28, (_671_v0).f27, globalState));
          let _761_v59;
          _761_v59 = _dafny.Map.Empty.slice().updateUnsafe((_671_v0).f27,(_671_v0).f27);
          let _762_v60;
          _762_v60 = _dafny.Set.fromElements((_671_v0).f27);
          let _763_v61;
          _763_v61 = _dafny.Seq.of((((_761_v59).contains(new BigNumber((_762_v60).length))) ? ((_761_v59).get(new BigNumber((_762_v60).length))) : (new BigNumber(125))), new BigNumber((_dafny.Seq.of(true, (_this).f26, (_671_v0).f28, p1, (_671_v0).f28)).length), (_671_v0).f27, (_671_v0).f27, new BigNumber((_761_v59).length));
          let _764_v62;
          let _nw137 = Array((new BigNumber(26)).toNumber());
          _nw137[(_dafny.ZERO).toNumber()] = _module.__default.fm5((_671_v0).f27, (_671_v0).f27, globalState);
          _nw137[(_dafny.ONE).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(2)).toNumber()] = _this.f25;
          _nw137[(new BigNumber(3)).toNumber()] = (((_759_v57).contains(_dafny.Seq.of((_this).f26))) ? ((_759_v57).get(_dafny.Seq.of((_this).f26))) : ((_671_v0).f27));
          _nw137[(new BigNumber(4)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(5)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(6)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(7)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(8)).toNumber()] = _this.f25;
          _nw137[(new BigNumber(9)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(10)).toNumber()] = new BigNumber((p2).length);
          _nw137[(new BigNumber(11)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(12)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(13)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(14)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(15)).toNumber()] = new BigNumber((_760_v58).cardinality());
          _nw137[(new BigNumber(16)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(17)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(18)).toNumber()] = (_671_v0).f27;
          _nw137[(new BigNumber(19)).toNumber()] = new BigNumber((_716_v23).length);
          _nw137[(new BigNumber(20)).toNumber()] = _this.f25;
          _nw137[(new BigNumber(21)).toNumber()] = _this.f25;
          _nw137[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_763_v61).length));
          _nw137[(new BigNumber(23)).toNumber()] = _this.f25;
          _nw137[(new BigNumber(24)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("uxb")).length);
          _nw137[(new BigNumber(25)).toNumber()] = (_671_v0).f27;
          _764_v62 = _nw137;
          let _765_v63;
          let _nw138 = Array((new BigNumber(13)).toNumber());
          _nw138[(_dafny.ZERO).toNumber()] = _764_v62;
          _nw138[(_dafny.ONE).toNumber()] = _764_v62;
          _nw138[(new BigNumber(2)).toNumber()] = _764_v62;
          _nw138[(new BigNumber(3)).toNumber()] = _764_v62;
          _nw138[(new BigNumber(4)).toNumber()] = _764_v62;
          _nw138[(new BigNumber(5)).toNumber()] = _764_v62;
          _nw138[(new BigNumber(6)).toNumber()] = _764_v62;
          _nw138[(new BigNumber(7)).toNumber()] = _764_v62;
          _nw138[(new BigNumber(8)).toNumber()] = _764_v62;
          _nw138[(new BigNumber(9)).toNumber()] = _764_v62;
          _nw138[(new BigNumber(10)).toNumber()] = _764_v62;
          _nw138[(new BigNumber(11)).toNumber()] = _764_v62;
          _nw138[(new BigNumber(12)).toNumber()] = _764_v62;
          _765_v63 = _nw138;
          let _766_v64;
          _766_v64 = _dafny.Seq.UnicodeFromString("bedhejjw");
          let _767_v65;
          _767_v65 = _dafny.Map.Empty.slice().updateUnsafe((_766_v64),(_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_671_v0).f27)).Merge(_758_v56));
          let _rhs150 = _671_v0;
          let _rhs151 = (((_767_v65).contains(_dafny.Seq.Concat(_718_v25, _718_v25))) ? ((_767_v65).get(_dafny.Seq.Concat(_718_v25, _718_v25))) : (_758_v56));
          let _rhs152 = (_671_v0).f27;
          let _rhs153 = _765_v63;
          let _lhs130 = globalState;
          _671_v0 = _rhs150;
          _758_v56 = _rhs151;
          _lhs130.f10 = _rhs152;
          _765_v63 = _rhs153;
          let _index120 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_764_v62).length));
          (_764_v62)[_index120] = (_671_v0).f27;
          let _index121 = _module.__default.safeIndex(new BigNumber(340), new BigNumber((_764_v62).length));
          (_764_v62)[_index121] = (_this.f25).multipliedBy(_this.f25);
          let _768_v66;
          let _769_v67;
          let _out10;
          let _out11;
          let _outcollector2 = (_this).m2(globalState);
          _out10 = _outcollector2[0];
          _out11 = _outcollector2[1];
          _768_v66 = _out10;
          _769_v67 = _out11;
          let _770_v68;
          let _nw139 = new _module.C3();
          _nw139.__ctor(_769_v67, (_this).f24, _this.f25);
          _770_v68 = _nw139;
          _671_v0 = _671_v0;
        } else {
          let _771_v69;
          _771_v69 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_module.__default.fm10((_this).f26, (_671_v0).f28, false, (_671_v0).f28, globalState));
          let _772_v70;
          _772_v70 = _dafny.Seq.of(_771_v69, _771_v69, _771_v69);
          _772_v70 = _772_v70;
          let _773_v71;
          let _nw140 = new _module.C1();
          _nw140.__ctor((_this).f24, (_671_v0).f27);
          _773_v71 = _nw140;
          let _774_v72;
          _774_v72 = _dafny.Map.Empty.slice().updateUnsafe(p1,_773_v71);
          _773_v71 = (((_774_v72).contains((_this).f26)) ? ((_774_v72).get((_this).f26)) : (_773_v71));
          let _775_v73;
          _775_v73 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_716_v23);
          _775_v73 = (_775_v73).update(false, _716_v23);
          let _776_v74;
          _776_v74 = _dafny.Set.fromElements((_this).f26);
          r2 = (_776_v74).IsSubsetOf(_776_v74);
          let _index122 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_719_v26).length));
          (_719_v26)[_index122] = (_this).f26;
          let _index123 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_719_v26).length));
          (_719_v26)[_index123] = (_671_v0).f28;
        }
        let _777_v75;
        let _nw141 = Array((new BigNumber(17)).toNumber()).fill([]);
        _777_v75 = _nw141;
        let _index124 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_777_v75).length));
        (_777_v75)[_index124] = _719_v26;
        let _778_v76;
        let _nw142 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _778_v76 = _nw142;
        let _index125 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_778_v76).length));
        (_778_v76)[_index125] = new BigNumber(-380);
        let _779_v77;
        _779_v77 = _dafny.Seq.of(_719_v26, _719_v26);
        let _780_v78;
        _780_v78 = _module.D3.create_DC5(_716_v23, _this.f25, (_671_v0).f27);
        let _index126 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_777_v75).length));
        let _index127 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_778_v76).length));
        let _rhs154 = (_779_v77)[_module.__default.safeIndex((_671_v0).f27, new BigNumber((_779_v77).length))];
        let _rhs155 = _this.f25;
        let _rhs156 = _module.__default.safeDivisionInt((_this.f25).multipliedBy((_671_v0).f27), (_671_v0).f27);
        let _rhs157 = (_716_v23)[_module.__default.safeIndex((_780_v78).dtor_cf8, new BigNumber((_716_v23).length))];
        let _lhs131 = _777_v75;
        let _lhs132 = _module.__default.safeIndex(new BigNumber(244), new BigNumber((_777_v75).length));
        let _lhs133 = _778_v76;
        let _lhs134 = _module.__default.safeIndex(new BigNumber(26), new BigNumber((_778_v76).length));
        let _lhs135 = globalState;
        _lhs131[_lhs132] = _rhs154;
        _lhs133[_lhs134] = _rhs155;
        _lhs135.f15 = _rhs156;
        r2 = _rhs157;
      }
      r0 = (_this).f26;
      r1 = _dafny.Seq.update(_716_v23, _module.__default.safeIndex((_this.f25).minus(new BigNumber(758)), new BigNumber((_716_v23).length)), (_671_v0).f28);
      r2 = (_671_v0).f28;
      return [r0, r1, r2];
    }
    m2(globalState) {
      let _this = this;
      let r0 = _dafny.Map.Empty;
      let r1 = false;
      let _781_v0;
      _781_v0 = _module.D7.create_DC14((_this).f26, new BigNumber(((_this).fm3((_this).f24, globalState)).length), _this.f25);
      if ((((_this).f26) ? ((_781_v0).dtor_cf20) : (!((new BigNumber(961)).isLessThan((_dafny.ZERO).minus(_this.f25)))))) {
        let _782_v2;
        _782_v2 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('n'.codePointAt(0)),(_this).f26);
        let _783_v3;
        _783_v3 = _dafny.Set.fromElements(_this.f25, new BigNumber((function () {
          let _coll35 = new _dafny.Map();
          for (const _compr_35 of (_782_v2).Keys.Elements) {
            let _784_v1 = _compr_35;
            if ((_782_v2).contains(_784_v1)) {
              _coll35.push([_784_v1,(_dafny.ZERO).minus(_this.f25)]);
            }
          }
          return _coll35;
        }()).length));
        let _785_v4;
        _785_v4 = _module.D11.create_DC23(new BigNumber((_783_v3).length), (_this).f26);
        let _786_v5;
        let _nw143 = new _module.C3();
        _nw143.__ctor(!((_785_v4).dtor_cf35), (_this).f24, (((_this).f26) ? (_this.f25) : (_this.f25)));
        _786_v5 = _nw143;
        let _787_v6;
        _787_v6 = new _dafny.CodePoint('f'.codePointAt(0));
        let _rhs158 = (_module.__default.fm22(new BigNumber(949), !(!((_this).f26)), _786_v5.f31, globalState)).dtor_cf34;
        let _rhs159 = _787_v6;
        let _rhs160 = _786_v5.f31;
        let _rhs161 = (_this).f26;
        let _lhs136 = globalState;
        let _lhs137 = globalState;
        let _lhs138 = globalState;
        _lhs136.f0 = _rhs158;
        _787_v6 = _rhs159;
        _lhs137.f18 = _rhs160;
        _lhs138.f18 = _rhs161;
        let _788_v7;
        _788_v7 = _dafny.MultiSet.fromElements(_this.f25, _this.f25);
        let _789_v8;
        _789_v8 = _dafny.Seq.of(_this.f25);
        (_this).f25 = _module.__default.fm17((_this).f26, (_this).f26, (_dafny.MultiSet.FromArray(_789_v8)).IsProperSubsetOf(_788_v7), globalState);
        let _790_v9;
        _790_v9 = _dafny.Map.Empty.slice().updateUnsafe(_789_v8,_786_v5.f31);
        if ((_790_v9).equals(_790_v9)) {
          let _791_v10;
          _791_v10 = _dafny.Seq.of(_786_v5.f31);
          _791_v10 = _dafny.Seq.Concat(_dafny.Seq.Concat(_791_v10, _791_v10), _791_v10);
          (globalState).f18 = _786_v5.f31;
          let _792_v11;
          _792_v11 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,_dafny.Seq.Concat(_dafny.Seq.update(_791_v10, _module.__default.safeIndex(_this.f25, new BigNumber((_791_v10).length)), _786_v5.f31), _791_v10));
          _792_v11 = (_792_v11).update((_this).f24, _791_v10);
          let _793_v12;
          let _nw144 = Array((new BigNumber(15)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _793_v12 = _nw144;
          _793_v12 = _793_v12;
          (_786_v5).f31 = _786_v5.f31;
        } else {
          let _794_v13;
          _794_v13 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm17(_786_v5.f31, _786_v5.f31, _786_v5.f31, globalState),false);
          _794_v13 = _794_v13;
          _787_v6 = (_this).f24;
          (globalState).f10 = (_this.f25).minus(_this.f25);
          let _795_v14;
          let _init22 = ((_796_v5) => function (_797_i0) {
            return _796_v5.f31;
          })(_786_v5);
          let _nw145 = Array((new BigNumber(20)).toNumber());
          for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw145.length); _i0_22++) {
            _nw145[_i0_22] = _init22(new BigNumber(_i0_22));
          }
          _795_v14 = _nw145;
          let _798_v15;
          _798_v15 = _dafny.Seq.of(_788_v7);
          let _799_v16;
          _799_v16 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.areEqual(_798_v15, _798_v15),_795_v14);
          _795_v14 = (((_799_v16).contains((_this).f26)) ? ((_799_v16).get((_this).f26)) : (_795_v14));
          let _800_v17;
          _800_v17 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(_786_v5.f31),new BigNumber((_783_v3).length));
          let _801_v18;
          _801_v18 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm10((_this).f26, _786_v5.f31, true, (_this).f26, globalState),(_this).f24);
          let _802_v19;
          let _nw146 = new _module.C4();
          _nw146.__ctor(_800_v17, _786_v5.f31, (((_801_v18).contains((_this).f24)) ? ((_801_v18).get((_this).f24)) : (_787_v6)), _this.f25);
          _802_v19 = _nw146;
          let _803_v20;
          _803_v20 = _dafny.Map.Empty.slice().updateUnsafe(_802_v19,(_this).f24);
          _803_v20 = (_803_v20).update(_802_v19, (_802_v19).f24);
        }
        let _804_v21;
        let _nw147 = Array((_dafny.ONE).toNumber());
        _nw147[(_dafny.ZERO).toNumber()] = _786_v5.f31;
        _804_v21 = _nw147;
        let _805_v22;
        let _nw148 = Array((new BigNumber(21)).toNumber()).fill(_module.D11.Default());
        _805_v22 = _nw148;
        let _806_v23;
        let _nw149 = Array((new BigNumber(27)).toNumber());
        _806_v23 = _nw149;
        let _807_v24;
        _807_v24 = _module.D11.create_DC22(_806_v23);
        let _index128 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_805_v22).length));
        (_805_v22)[_index128] = _807_v24;
        let _808_v25;
        let _nw150 = Array((new BigNumber(28)).toNumber()).fill(_dafny.Set.Empty);
        _808_v25 = _nw150;
        let _809_v26;
        let _nw151 = Array((new BigNumber(4)).toNumber());
        _809_v26 = _nw151;
        let _810_v27;
        _810_v27 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_809_v26);
        let _811_v28;
        _811_v28 = _dafny.Seq.of(_804_v21);
        let _812_v29;
        _812_v29 = _dafny.Map.Empty.slice().updateUnsafe((_810_v27).Merge(_810_v27),(_811_v28)[_module.__default.safeIndex(new BigNumber(616), new BigNumber((_811_v28).length))]);
        let _index129 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_805_v22).length));
        let _rhs162 = (((_812_v29).contains(_dafny.Map.Empty.slice().updateUnsafe(false,_809_v26))) ? ((_812_v29).get(_dafny.Map.Empty.slice().updateUnsafe(false,_809_v26))) : (_804_v21));
        let _rhs163 = _807_v24;
        let _rhs164 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_this.f25), _this.f25);
        let _rhs165 = _808_v25;
        let _lhs139 = _805_v22;
        let _lhs140 = _module.__default.safeIndex(new BigNumber(483), new BigNumber((_805_v22).length));
        let _lhs141 = globalState;
        _804_v21 = _rhs162;
        _lhs139[_lhs140] = _rhs163;
        _lhs141.f16 = _rhs164;
        _808_v25 = _rhs165;
      } else {
        let _813_v30;
        let _nw152 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.of());
        _813_v30 = _nw152;
        let _814_v31;
        _814_v31 = _dafny.Seq.of(_module.D10.create_DC21(_this.f25, (_this).f26, (_this).f26));
        let _index130 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_813_v30).length));
        (_813_v30)[_index130] = _dafny.Seq.Concat(_814_v31, _814_v31);
        let _815_v32;
        _815_v32 = _module.D10.create_DC21(new BigNumber(745), (_this).f26, (_this).f26);
        let _index131 = _module.__default.safeIndex(new BigNumber(189), new BigNumber((_813_v30).length));
        (_813_v30)[_index131] = _dafny.Seq.of(_815_v32, _815_v32, _815_v32, _815_v32, _815_v32);
        let _816_v33;
        _816_v33 = _dafny.Seq.UnicodeFromString("qxbymslm");
        let _817_v34;
        _817_v34 = _module.D11.create_DC24((_this).f26, new BigNumber(802));
        let _818_v35;
        _818_v35 = _module.D4.create_DC8();
        let _819_v36;
        _819_v36 = _dafny.Map.Empty.slice().updateUnsafe((_817_v34).dtor_cf36,_818_v35);
        let _820_v37;
        _820_v37 = _dafny.Map.Empty.slice().updateUnsafe(_816_v33,_819_v36);
        _820_v37 = (_820_v37).update(_816_v33, _819_v36);
        let _821_v38;
        _821_v38 = _dafny.Map.Empty.slice().updateUnsafe(_this.f25,(_this).f26);
        let _822_v39;
        _822_v39 = _dafny.Map.Empty.slice().updateUnsafe(false,(_this).f26);
        let _823_v40;
        _823_v40 = _dafny.Seq.of(new BigNumber((_821_v38).length), _this.f25, _this.f25, new BigNumber((_822_v39).length));
        let _824_v41;
        _824_v41 = _dafny.Seq.of((_dafny.ZERO).minus(_this.f25), new BigNumber((_823_v40).length), new BigNumber(-217));
        let _825_v42;
        _825_v42 = _dafny.Set.fromElements((_824_v41)[_module.__default.safeIndex(_this.f25, new BigNumber((_824_v41).length))]);
        let _826_v43;
        _826_v43 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,new BigNumber((_825_v42).length));
        (globalState).f22 = ((_module.__default.fm0(_this.f25, (_this).f26, (_this).f26, _this.f25, globalState)) ? ((((_826_v43).contains((_this).f26)) ? ((_826_v43).get((_this).f26)) : (_this.f25))) : (_this.f25));
        let _827_v44;
        let _nw153 = new _module.C1();
        _nw153.__ctor((_this).f24, _module.__default.safeDivisionInt(_this.f25, _this.f25));
        _827_v44 = _nw153;
        let _828_v45;
        let _nw154 = Array((new BigNumber(8)).toNumber()).fill(_dafny.ZERO);
        _828_v45 = _nw154;
        let _index132 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_828_v45).length));
        (_828_v45)[_index132] = _module.__default.fm5(new BigNumber((_816_v33).length), _this.f25, globalState);
        let _index133 = _module.__default.safeIndex(new BigNumber(102), new BigNumber((_828_v45).length));
        (_828_v45)[_index133] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_this.f25).plus(_module.__default.safeDivisionInt(_this.f25, _this.f25))));
      }
      let _829_v46;
      _829_v46 = _module.D10.create_DC21(_this.f25, !((_this).f26), (_this).f26);
      let _830_v47;
      let _init23 = function (_831_i1) {
        return _dafny.Map.Empty.slice().updateUnsafe(_this.f25,!((_this).f26));
      };
      let _nw155 = Array((new BigNumber(14)).toNumber());
      for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw155.length); _i0_23++) {
        _nw155[_i0_23] = _init23(new BigNumber(_i0_23));
      }
      _830_v47 = _nw155;
      let _source13 = _module.D4.create_DC7((((_829_v46).dtor_cf32) ? (_830_v47) : (_830_v47)));
      if (_source13.is_DC8) {
        let _832_v48;
        let _nw156 = Array((_dafny.ONE).toNumber());
        _nw156[(_dafny.ZERO).toNumber()] = _module.__default.safeDivisionInt(_this.f25, _this.f25);
        _832_v48 = _nw156;
        let _833_v49;
        _833_v49 = _dafny.Seq.UnicodeFromString("v");
        let _834_v50;
        _834_v50 = _dafny.MultiSet.fromElements(_this.f25, new BigNumber((_833_v49).length));
        let _index134 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length));
        (_832_v48)[_index134] = (new BigNumber((_834_v50).cardinality())).minus(_this.f25);
        let _index135 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length));
        (_832_v48)[_index135] = _this.f25;
        let _835_v51;
        _835_v51 = _module.D10.create_DC19((_this).f24);
        let _836_v52;
        let _nw157 = Array((new BigNumber(25)).toNumber());
        _nw157[(_dafny.ZERO).toNumber()] = (_this).f24;
        _nw157[(_dafny.ONE).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(2)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(3)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(4)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(5)).toNumber()] = new _dafny.CodePoint('q'.codePointAt(0));
        _nw157[(new BigNumber(6)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(7)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(8)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(9)).toNumber()] = _module.__default.fm10((_this).f26, (_this).f26, (_this).f26, (_this).f26, globalState);
        _nw157[(new BigNumber(10)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(11)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(12)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(13)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(14)).toNumber()] = _module.__default.fm10((_this).f26, (_this).f26, (_this).f26, (_this).f26, globalState);
        _nw157[(new BigNumber(15)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(16)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(17)).toNumber()] = new _dafny.CodePoint('u'.codePointAt(0));
        _nw157[(new BigNumber(18)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(19)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(20)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(21)).toNumber()] = (_835_v51).dtor_cf27;
        _nw157[(new BigNumber(22)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(23)).toNumber()] = (_this).f24;
        _nw157[(new BigNumber(24)).toNumber()] = (_this).f24;
        _836_v52 = _nw157;
        let _837_v53;
        _837_v53 = _dafny.Seq.of(_833_v49, _833_v49, _833_v49, _dafny.Seq.UnicodeFromString("kucbeyarw"), (_this).fm3((_this).f24, globalState));
        let _838_v54;
        _838_v54 = _dafny.MultiSet.fromElements(_837_v53, _837_v53, _837_v53, _837_v53, _837_v53);
        let _rhs166 = _836_v52;
        let _rhs167 = (_dafny.ZERO).minus((((_838_v54).contains(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(920)), function (_840_i2) {
          return (_this).f24;
        })))) ? ((_838_v54).get(_dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(920)), function (_839_i2) {
          return (_this).f24;
        })))) : (_this.f25)));
        let _rhs168 = (_this).f26;
        let _lhs142 = globalState;
        _836_v52 = _rhs166;
        _lhs142.f16 = _rhs167;
        r1 = _rhs168;
        let _index136 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length));
        (_832_v48)[_index136] = _module.__default.safeModuloInt(new BigNumber(267), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(845)), function (_841_i3) {
          return new _dafny.CodePoint('g'.codePointAt(0));
        }), _833_v49)).length));
        let _842_v55;
        _842_v55 = _module.D3.create_DC6((_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))], (_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))], true);
        if ((_842_v55).dtor_cf11) {
          let _843_v56;
          _843_v56 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm0(new BigNumber((_833_v49).length), (_this).f26, (_this).f26, _this.f25, globalState),_832_v48);
          _832_v48 = (((_843_v56).contains((_this).f26)) ? ((_843_v56).get((_this).f26)) : (_832_v48));
          let _844_v57;
          _844_v57 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,!(_module.__default.fm0(_this.f25, (_this).f26, (_this).f26, _this.f25, globalState)));
          _844_v57 = (_844_v57).update(true, (_this).f26);
          let _845_v58;
          _845_v58 = _dafny.Map.Empty.slice().updateUnsafe((_this).f26,_842_v55);
          let _846_v59;
          _846_v59 = _module.D9.create_DC16((_dafny.Map.Empty.slice().updateUnsafe((_this).f26,_842_v55)).Merge((_845_v58).update(!((_this).f26), _module.__default.fm23(globalState))));
          _846_v59 = _module.__default.fm24((_this).f26, (_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))], _833_v49, globalState);
          let _847_v60;
          _847_v60 = _dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f26);
          let _848_v61;
          _848_v61 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_847_v60).length),_this.f25);
          let _849_v64;
          _849_v64 = _dafny.Set.fromElements((_this).f24, (_this).f24, (_this).f24, (_this).f24);
          let _850_v65;
          _850_v65 = _dafny.Seq.of(_849_v64, _849_v64);
          let _851_v66;
          _851_v66 = _dafny.Map.Empty.slice().updateUnsafe(_module.D10.create_DC19(new _dafny.CodePoint('c'.codePointAt(0))),false);
          let _852_v67;
          _852_v67 = _dafny.Map.Empty.slice().updateUnsafe(!((_this).f26),_851_v66);
          let _index137 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length));
          let _rhs169 = function () {
            let _coll36 = new _dafny.Map();
            for (const _compr_36 of (_dafny.Map.Empty.slice().updateUnsafe(_this.f25,_this.f25)).Keys.Elements) {
              let _853_v62 = _compr_36;
              if ((_dafny.Map.Empty.slice().updateUnsafe(_this.f25,_this.f25)).contains(_853_v62)) {
                _coll36.push([(_853_v62).plus(_this.f25),new BigNumber(((function () {
                  let _coll37 = new _dafny.Map();
                  for (const _compr_37 of (_834_v50).Elements) {
                    let _854_v63 = _compr_37;
                    if ((_834_v50).contains(_854_v63)) {
                      _coll37.push([(_854_v63).plus(_this.f25),(_this).f26]);
                    }
                  }
                  return _coll37;
                }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements((_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))])).length),(_this).f26))).length)]);
              }
            }
            return _coll36;
          }();
          let _rhs170 = _dafny.areEqual(_850_v65, _dafny.Seq.Concat(_850_v65, _850_v65));
          let _rhs171 = _module.__default.safeModuloInt(new BigNumber((_852_v67).length), (_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))]);
          let _lhs143 = globalState;
          let _lhs144 = _832_v48;
          let _lhs145 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length));
          _848_v61 = _rhs169;
          _lhs143.f18 = _rhs170;
          _lhs144[_lhs145] = _rhs171;
          let _855_v68;
          _855_v68 = _dafny.Seq.of((_dafny.ZERO).minus(_this.f25), new BigNumber((_833_v49).length));
          let _856_v69;
          _856_v69 = _dafny.Set.fromElements((_this).f26, (_this).f26);
          let _857_v70;
          _857_v70 = _dafny.MultiSet.fromElements((_this).f26);
          let _858_v71;
          _858_v71 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_this.f25, new BigNumber((_module.__default.fm25(new BigNumber((_856_v69).length), (_this).f26, (_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))], globalState)).length)),_857_v70);
          let _859_v72;
          _859_v72 = _dafny.Seq.of((_this).f26);
          let _860_v73;
          _860_v73 = _dafny.Map.Empty.slice().updateUnsafe((_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))],(_this).f26);
          let _index138 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length));
          let _rhs172 = new BigNumber(33);
          let _rhs173 = (((_848_v61).contains((_dafny.ZERO).minus(((_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))]).multipliedBy(new BigNumber((_855_v68).length))))) ? ((_848_v61).get((_dafny.ZERO).minus(((_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))]).multipliedBy(new BigNumber((_855_v68).length))))) : (new BigNumber((((((_858_v71).contains(_module.__default.fm26(globalState))) ? ((_858_v71).get(_module.__default.fm26(globalState))) : (_857_v70))).Intersect(_dafny.MultiSet.FromArray(_859_v72))).cardinality())));
          let _rhs174 = (((_860_v73).contains((((_857_v70).contains((_this).f26)) ? ((_857_v70).get((_this).f26)) : ((_dafny.ZERO).minus((_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))]))))) ? ((_860_v73).get((((_857_v70).contains((_this).f26)) ? ((_857_v70).get((_this).f26)) : ((_dafny.ZERO).minus((_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))]))))) : (!((((_844_v57).contains((_this).f26)) ? ((_844_v57).get((_this).f26)) : ((_this).f26)))));
          let _lhs146 = globalState;
          let _lhs147 = _832_v48;
          let _lhs148 = _module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length));
          let _lhs149 = globalState;
          _lhs146.f2 = _rhs172;
          _lhs147[_lhs148] = _rhs173;
          _lhs149.f18 = _rhs174;
        } else {
          let _861_v74;
          let _init24 = function (_862_i4) {
            return _dafny.Seq.of(_this.f25, _this.f25);
          };
          let _nw158 = Array((new BigNumber(27)).toNumber());
          for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw158.length); _i0_24++) {
            _nw158[_i0_24] = _init24(new BigNumber(_i0_24));
          }
          _861_v74 = _nw158;
          let _863_v75;
          _863_v75 = _dafny.Map.Empty.slice().updateUnsafe(_861_v74,true);
          (globalState).f18 = (((_863_v75).contains(_861_v74)) ? ((_863_v75).get(_861_v74)) : ((_this).f26));
          let _rhs175 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(913)), ((_864_v49) => function (_865_i5) {
            return _864_v49;
          })(_833_v49));
          let _rhs176 = (_this).f26;
          _837_v53 = _rhs175;
          r1 = _rhs176;
          (globalState).f2 = _this.f25;
          let _866_v76;
          let _nw159 = new _module.C0();
          _nw159.__ctor();
          _866_v76 = _nw159;
          _866_v76 = _866_v76;
          let _867_v77;
          _867_v77 = _dafny.Seq.of(!((_this).f26));
          let _868_v78;
          let _nw160 = Array((new BigNumber(11)).toNumber());
          _nw160[(_dafny.ZERO).toNumber()] = _867_v77;
          _nw160[(_dafny.ONE).toNumber()] = _module.__default.fm7(globalState);
          _nw160[(new BigNumber(2)).toNumber()] = _867_v77;
          _nw160[(new BigNumber(3)).toNumber()] = _867_v77;
          _nw160[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_867_v77, _module.__default.safeIndex((_832_v48)[_module.__default.safeIndex(new BigNumber(676), new BigNumber((_832_v48).length))], new BigNumber((_867_v77).length)), (_this).f26), _867_v77);
          _nw160[(new BigNumber(5)).toNumber()] = _dafny.Seq.of(!((_this).f26), (_this).f26, (_this).f26, (_this).f26, (_this).f26);
          _nw160[(new BigNumber(6)).toNumber()] = _867_v77;
          _nw160[(new BigNumber(7)).toNumber()] = _dafny.Seq.of(false, (_this).f26);
          _nw160[(new BigNumber(8)).toNumber()] = _867_v77;
          _nw160[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.of((_this).f26, (_this).f26), _module.__default.safeIndex(new BigNumber(-115), new BigNumber((_dafny.Seq.of((_this).f26, (_this).f26)).length)), (_this).f26), _867_v77);
          _nw160[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_867_v77, _867_v77);
          _868_v78 = _nw160;
          let _index139 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_868_v78).length));
          (_868_v78)[_index139] = _867_v77;
          let _index140 = _module.__default.safeIndex(new BigNumber(568), new BigNumber((_868_v78).length));
          (_868_v78)[_index140] = _867_v77;
        }
      } else {
        let _869___mcc_h0 = (_source13).cf12;
        let _870_cf12 = _869___mcc_h0;
        (globalState).f19 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(323)), function (_871_i6) {
          return new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("uuvoe"), _dafny.Seq.UnicodeFromString("knnyuorh"))).length);
        });
        let _872_v79;
        let _init25 = function (_873_i7) {
          return _module.__default.safeDivisionInt(_873_i7, new BigNumber((_dafny.Set.fromElements((_this).f26, (_this).f26)).length));
        };
        let _nw161 = Array((new BigNumber(2)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw161.length); _i0_25++) {
          _nw161[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _872_v79 = _nw161;
        _872_v79 = _872_v79;
        r1 = (_this).f26;
        if ((_this).f26) {
          let _874_v80;
          _874_v80 = _dafny.Seq.UnicodeFromString("ksmojox");
          (globalState).f10 = new BigNumber((_874_v80).length);
          let _875_v81;
          _875_v81 = _dafny.Seq.of((_this).f26, !((_this).f26));
          let _876_v82;
          let _nw162 = Array((new BigNumber(10)).toNumber());
          _nw162[(_dafny.ZERO).toNumber()] = (_this.f25).isLessThan(new BigNumber(502));
          _nw162[(_dafny.ONE).toNumber()] = true;
          _nw162[(new BigNumber(2)).toNumber()] = (_this).f26;
          _nw162[(new BigNumber(3)).toNumber()] = (_this).f26;
          _nw162[(new BigNumber(4)).toNumber()] = (_this).f26;
          _nw162[(new BigNumber(5)).toNumber()] = (_this).f26;
          _nw162[(new BigNumber(6)).toNumber()] = true;
          _nw162[(new BigNumber(7)).toNumber()] = (((_this).f26) ? ((_this).f26) : ((_this).f26));
          _nw162[(new BigNumber(8)).toNumber()] = (_this).f26;
          _nw162[(new BigNumber(9)).toNumber()] = _dafny.areEqual(_875_v81, _875_v81);
          _876_v82 = _nw162;
          let _index141 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_876_v82).length));
          (_876_v82)[_index141] = (_this).f26;
          let _index142 = _module.__default.safeIndex(new BigNumber(344), new BigNumber((_876_v82).length));
          (_876_v82)[_index142] = (_this).f26;
          let _877_v83;
          let _nw163 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _877_v83 = _nw163;
          let _index143 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_877_v83).length));
          (_877_v83)[_index143] = _dafny.Seq.Concat(_874_v80, _874_v80);
          let _index144 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_877_v83).length));
          let _rhs177 = _876_v82;
          let _rhs178 = _874_v80;
          let _lhs150 = _877_v83;
          let _lhs151 = _module.__default.safeIndex(new BigNumber(195), new BigNumber((_877_v83).length));
          _876_v82 = _rhs177;
          _lhs150[_lhs151] = _rhs178;
          let _878_v84;
          let _nw164 = new _module.C4();
          _nw164.__ctor(_module.__default.fm27(_this.f25, globalState), _module.__default.fm0(_this.f25, (_this).f26, false, _this.f25, globalState), (_this).f24, _this.f25);
          _878_v84 = _nw164;
          let _879_v85;
          let _nw165 = new _module.C3();
          _nw165.__ctor((_this).f26, _module.__default.fm10((_878_v84).f30, (_878_v84).f30, (_this).f26, (_876_v82)[_module.__default.safeIndex(new BigNumber(344), new BigNumber((_876_v82).length))], globalState), _module.__default.safeModuloInt(new BigNumber((_874_v80).length), _this.f25));
          _879_v85 = _nw165;
        } else {
          let _880_v86;
          _880_v86 = _dafny.Seq.UnicodeFromString("xukawq");
          let _881_v87;
          _881_v87 = _dafny.MultiSet.fromElements(_880_v86, _dafny.Seq.UnicodeFromString("ylh"), _880_v86, _880_v86, _880_v86);
          let _882_v88;
          _882_v88 = _dafny.Seq.UnicodeFromString("qxgte");
          let _883_v89;
          _883_v89 = _dafny.Seq.of((_this).f26, (_this).f26);
          let _884_v90;
          let _nw166 = new _module.C3();
          _nw166.__ctor(!((_881_v87).IsDisjointFrom(_dafny.MultiSet.fromElements((_882_v88)))), (_this).f24, new BigNumber((_883_v89).length));
          _884_v90 = _nw166;
          let _885_v91;
          let _init26 = ((_886_v86) => function (_887_i8) {
            return _886_v86;
          })(_880_v86);
          let _nw167 = Array((new BigNumber(26)).toNumber());
          for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw167.length); _i0_26++) {
            _nw167[_i0_26] = _init26(new BigNumber(_i0_26));
          }
          _885_v91 = _nw167;
          let _index145 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_885_v91).length));
          (_885_v91)[_index145] = _880_v86;
          let _index146 = _module.__default.safeIndex(new BigNumber(692), new BigNumber((_885_v91).length));
          (_885_v91)[_index146] = _dafny.Seq.Concat(_880_v86, _dafny.Seq.Concat(_880_v86, _dafny.Seq.UnicodeFromString("ixxmbm")));
          (globalState).f22 = new BigNumber((_880_v86).length);
          (globalState).f3 = _880_v86;
          (globalState).f18 = (_this).f26;
        }
      }
      let _888_v92;
      let _nw168 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
      _888_v92 = _nw168;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_888_v92).length))) {
        let _889_i9 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_889_i9)) && ((_889_i9).isLessThan(new BigNumber((_888_v92).length))))) {
          (_888_v92)[(_889_i9)] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("gvibnejd"), _dafny.Seq.UnicodeFromString("polnxwjt"));
        }
      }
      let _890_v94;
      let _init27 = function (_891_i10) {
        return (_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f26)).Merge(function () {
          let _coll38 = new _dafny.Map();
          for (const _compr_38 of (_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f26)).Keys.Elements) {
            let _892_v93 = _compr_38;
            if ((_dafny.Map.Empty.slice().updateUnsafe((_this).f24,(_this).f26)).contains(_892_v93)) {
              _coll38.push([_892_v93,(_this).f26]);
            }
          }
          return _coll38;
        }());
      };
      let _nw169 = Array((new BigNumber(7)).toNumber());
      for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw169.length); _i0_27++) {
        _nw169[_i0_27] = _init27(new BigNumber(_i0_27));
      }
      _890_v94 = _nw169;
      _890_v94 = _890_v94;
      let _893_v95;
      let _nw170 = Array((new BigNumber(21)).toNumber()).fill(false);
      _893_v95 = _nw170;
      let _index147 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_893_v95).length));
      (_893_v95)[_index147] = (_this).f26;
      let _894_v96;
      _894_v96 = _dafny.Set.fromElements((_this).f26, (_this).f26);
      let _895_v97;
      _895_v97 = _dafny.Map.Empty.slice().updateUnsafe(_894_v96,_this.f25);
      let _index148 = _module.__default.safeIndex(new BigNumber(908), new BigNumber((_893_v95).length));
      (_893_v95)[_index148] = (((_this).f26) ? (_module.__default.fm0(new BigNumber((_895_v97).length), (_this).f26, false, _this.f25, globalState)) : ((_this).f26));
      let _896_v98;
      _896_v98 = _dafny.Map.Empty.slice().updateUnsafe((_this.f25).multipliedBy(_this.f25),new BigNumber(198));
      _896_v98 = (_896_v98).update(_this.f25, _module.__default.safeModuloInt(_this.f25, _this.f25));
      let _897_v99;
      _897_v99 = _dafny.Map.Empty.slice().updateUnsafe(_893_v95,_dafny.Map.Empty.slice().updateUnsafe(false,true));
      r0 = (_897_v99).Merge((_897_v99).Merge(_897_v99));
      r1 = false;
      return [r0, r1];
    }
    get f26() {
      let _this = this;
      return _this._f26;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this.f23 = _dafny.Map.Empty;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f23) {
      let _this = this;
      (_this).f23 = f23;
      return;
    }
    fm1(p0, globalState) {
      let _this = this;
      return !((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-509))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.UnicodeFromString("clwiitdbt")).length)))).equals((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-196)), function (_898_i0) {
        return new _dafny.CodePoint('k'.codePointAt(0));
      })).length))).length))).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber(-610))));
    };
    fm2(globalState) {
      let _this = this;
      return (new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(339)), function (_899_i0) {
        return new _dafny.CodePoint('y'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("rideqr"))).length)).isLessThanOrEqualTo((new BigNumber(-714)).plus(new BigNumber((function () {
        let _coll39 = new _dafny.Map();
        for (const _compr_39 of (_dafny.MultiSet.fromElements(_this.f23)).Elements) {
          let _900_v0 = _compr_39;
          if ((_dafny.MultiSet.fromElements(_this.f23)).contains(_900_v0)) {
            _coll39.push([_900_v0,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("yf")).length), (_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new _dafny.CodePoint('t'.codePointAt(0)))).length)))).length))]);
          }
        }
        return _coll39;
      }()).length)));
    };
    m0(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.Seq.UnicodeFromString("");
      let _901_v0;
      _901_v0 = true;
      if (_901_v0) {
        (globalState).f2 = _module.__default.safeDivisionInt(_module.__default.safeDivisionInt(p0, p0), new BigNumber(203));
        let _902_v1;
        _902_v1 = _dafny.Seq.of(_901_v0);
        if ((_902_v1)[_module.__default.safeIndex(p0, new BigNumber((_902_v1).length))]) {
          (globalState).f19 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(806)), ((_903_p0) => function (_904_i0) {
            return (_903_p0).multipliedBy(_903_p0);
          })(p0));
          (globalState).f10 = (p0).multipliedBy((p0).minus(p0));
          let _905_v2;
          _905_v2 = _dafny.Map.Empty.slice().updateUnsafe((((_this.f23).contains(p0)) ? ((_this.f23).get(p0)) : (_901_v0)),_901_v0);
          let _906_v3;
          _906_v3 = _module.D2.create_DC2(_901_v0);
          let _907_v4;
          _907_v4 = _dafny.Seq.UnicodeFromString("yxdeimc");
          let _908_v5;
          _908_v5 = new _dafny.CodePoint('n'.codePointAt(0));
          let _909_v6;
          let _nw171 = Array((new BigNumber(15)).toNumber());
          _nw171[(_dafny.ZERO).toNumber()] = true;
          _nw171[(_dafny.ONE).toNumber()] = _901_v0;
          _nw171[(new BigNumber(2)).toNumber()] = (((_905_v2).contains(_901_v0)) ? ((_905_v2).get(_901_v0)) : (_901_v0));
          _nw171[(new BigNumber(3)).toNumber()] = _901_v0;
          _nw171[(new BigNumber(4)).toNumber()] = _901_v0;
          _nw171[(new BigNumber(5)).toNumber()] = _901_v0;
          _nw171[(new BigNumber(6)).toNumber()] = _901_v0;
          _nw171[(new BigNumber(7)).toNumber()] = (_906_v3).dtor_cf2;
          _nw171[(new BigNumber(8)).toNumber()] = !(!(_dafny.areEqual(_dafny.Seq.UnicodeFromString("lkxsn"), _907_v4)));
          _nw171[(new BigNumber(9)).toNumber()] = true;
          _nw171[(new BigNumber(10)).toNumber()] = ((_901_v0) ? (_901_v0) : (true));
          _nw171[(new BigNumber(11)).toNumber()] = !_dafny.Seq.contains(_907_v4, _908_v5);
          _nw171[(new BigNumber(12)).toNumber()] = _901_v0;
          _nw171[(new BigNumber(13)).toNumber()] = _901_v0;
          _nw171[(new BigNumber(14)).toNumber()] = _901_v0;
          _909_v6 = _nw171;
          let _910_v7;
          _910_v7 = _dafny.Seq.of(p0, p0);
          let _index149 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_909_v6).length));
          (_909_v6)[_index149] = !_dafny.areEqual(_910_v7, _910_v7);
          let _index150 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_909_v6).length));
          (_909_v6)[_index150] = (p0).isLessThanOrEqualTo(p0);
          let _911_v9;
          _911_v9 = _dafny.Set.fromElements(((_901_v0) ? (_dafny.Map.Empty.slice().updateUnsafe(p0,(_909_v6)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_909_v6).length))])) : (_this.f23)), _this.f23, _this.f23, function () {
            let _coll40 = new _dafny.Map();
            for (const _compr_40 of _dafny.IntegerRange(new BigNumber(807), new BigNumber(609))) {
              let _912_v8 = _compr_40;
              if (((new BigNumber(807)).isLessThanOrEqualTo(_912_v8)) && ((_912_v8).isLessThan(new BigNumber(609)))) {
                _coll40.push([_module.__default.safeModuloInt(_912_v8, p0),(_909_v6)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_909_v6).length))]]);
              }
            }
            return _coll40;
          }(), _this.f23);
          let _913_v10;
          let _nw172 = Array((new BigNumber(8)).toNumber()).fill([]);
          _913_v10 = _nw172;
          let _914_v11;
          let _init28 = ((_915_v5) => function (_916_i1) {
            return _915_v5;
          })(_908_v5);
          let _nw173 = Array((new BigNumber(18)).toNumber());
          for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw173.length); _i0_28++) {
            _nw173[_i0_28] = _init28(new BigNumber(_i0_28));
          }
          _914_v11 = _nw173;
          let _index151 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_913_v10).length));
          (_913_v10)[_index151] = _914_v11;
          let _917_v12;
          _917_v12 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(218)), function (_918_i2) {
            return new _dafny.CodePoint('v'.codePointAt(0));
          });
          let _919_v13;
          _919_v13 = _dafny.MultiSet.fromElements(_917_v12, _917_v12);
          let _920_v14;
          _920_v14 = _dafny.Seq.of(_dafny.Seq.of((_this).fm1(_919_v13, globalState), _901_v0, false));
          let _921_v15;
          _921_v15 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_920_v14)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_920_v14).length))]);
          let _index152 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_913_v10).length));
          let _index153 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_909_v6).length));
          let _rhs179 = (_911_v9).Intersect((_module.D3.create_DC4(_911_v9)).dtor_cf5);
          let _rhs180 = _914_v11;
          let _rhs181 = !_dafny.areEqual(_910_v7, _910_v7);
          let _rhs182 = _901_v0;
          let _rhs183 = _dafny.Seq.Concat(_902_v1, (((_921_v15).contains(p0)) ? ((_921_v15).get(p0)) : (_dafny.Seq.of((_909_v6)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_909_v6).length))], (_909_v6)[_module.__default.safeIndex(new BigNumber(590), new BigNumber((_909_v6).length))]))));
          let _lhs152 = _913_v10;
          let _lhs153 = _module.__default.safeIndex(new BigNumber(953), new BigNumber((_913_v10).length));
          let _lhs154 = _909_v6;
          let _lhs155 = _module.__default.safeIndex(new BigNumber(590), new BigNumber((_909_v6).length));
          _911_v9 = _rhs179;
          _lhs152[_lhs153] = _rhs180;
          _lhs154[_lhs155] = _rhs181;
          _901_v0 = _rhs182;
          _902_v1 = _rhs183;
          let _rhs184 = (p0).plus(p0);
          let _rhs185 = _910_v7;
          let _lhs156 = globalState;
          _lhs156.f2 = _rhs184;
          _910_v7 = _rhs185;
        } else {
          let _922_v16;
          _922_v16 = _module.D11.create_DC24(_901_v0, p0);
          let _923_v17;
          _923_v17 = _dafny.Map.Empty.slice().updateUnsafe(_922_v16,_901_v0);
          let _924_v18;
          _924_v18 = new _dafny.CodePoint('p'.codePointAt(0));
          let _925_v19;
          _925_v19 = _dafny.Seq.UnicodeFromString("csjg");
          let _926_v20;
          let _nw174 = new _module.C4();
          _nw174.__ctor(_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(false),new BigNumber(707)), (((_923_v17).contains(_922_v16)) ? ((_923_v17).get(_922_v16)) : (_901_v0)), _924_v18, _module.__default.safeModuloInt(new BigNumber((_925_v19).length), p0));
          _926_v20 = _nw174;
          let _927_v21;
          _927_v21 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_dafny.ZERO).minus(p0)),new _dafny.CodePoint('k'.codePointAt(0)));
          _927_v21 = (_927_v21).update((_dafny.ZERO).minus(_module.__default.safeDivisionInt(p0, p0)), new _dafny.CodePoint('e'.codePointAt(0)));
          let _928_v22;
          _928_v22 = _dafny.Set.fromElements(p0);
          let _929_v23;
          _929_v23 = _module.D12.create_DC26(_928_v22);
          r2 = _module.__default.fm15(p0, (_dafny.ZERO).minus(p0), new BigNumber(((_929_v23).dtor_cf39).length), globalState);
          let _930_v24;
          let _nw175 = Array((new BigNumber(2)).toNumber()).fill([]);
          _930_v24 = _nw175;
          _930_v24 = _930_v24;
          let _931_v25;
          _931_v25 = _module.D7.create_DC13(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(703),_901_v0));
          let _932_v26;
          _932_v26 = _dafny.Map.Empty.slice().updateUnsafe(_931_v25,_902_v1);
          let _933_v27;
          _933_v27 = _dafny.Map.Empty.slice().updateUnsafe((_926_v20).f30,(_932_v26).Merge(_dafny.Map.Empty.slice().updateUnsafe(_931_v25,_902_v1)));
          _933_v27 = (_933_v27).update((_926_v20).f30, (_932_v26).update(_931_v25, _902_v1));
        }
        let _934_v28;
        let _init29 = ((_935_v1) => function (_936_i3) {
          return _935_v1;
        })(_902_v1);
        let _nw176 = Array((new BigNumber(16)).toNumber());
        for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw176.length); _i0_29++) {
          _nw176[_i0_29] = _init29(new BigNumber(_i0_29));
        }
        _934_v28 = _nw176;
        let _937_v29;
        _937_v29 = _dafny.Seq.UnicodeFromString("telw");
        let _938_v30;
        _938_v30 = _dafny.Seq.of(_902_v1);
        let _nw177 = Array((new BigNumber(22)).toNumber());
        _nw177[(_dafny.ZERO).toNumber()] = _902_v1;
        _nw177[(_dafny.ONE).toNumber()] = _902_v1;
        _nw177[(new BigNumber(2)).toNumber()] = _dafny.Seq.of(_901_v0, _901_v0, _901_v0);
        _nw177[(new BigNumber(3)).toNumber()] = _902_v1;
        _nw177[(new BigNumber(4)).toNumber()] = _902_v1;
        _nw177[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_902_v1, _dafny.Seq.update(_902_v1, _module.__default.safeIndex(new BigNumber((_937_v29).length), new BigNumber((_902_v1).length)), _901_v0));
        _nw177[(new BigNumber(6)).toNumber()] = _902_v1;
        _nw177[(new BigNumber(7)).toNumber()] = _902_v1;
        _nw177[(new BigNumber(8)).toNumber()] = _902_v1;
        _nw177[(new BigNumber(9)).toNumber()] = _dafny.Seq.of(true, _901_v0, _901_v0, !((((_this.f23).contains(new BigNumber((_902_v1).length))) ? ((_this.f23).get(new BigNumber((_902_v1).length))) : (_901_v0))));
        _nw177[(new BigNumber(10)).toNumber()] = _902_v1;
        _nw177[(new BigNumber(11)).toNumber()] = _902_v1;
        _nw177[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_902_v1, _902_v1);
        _nw177[(new BigNumber(13)).toNumber()] = _902_v1;
        _nw177[(new BigNumber(14)).toNumber()] = _dafny.Seq.Concat((_938_v30)[_module.__default.safeIndex(p0, new BigNumber((_938_v30).length))], _902_v1);
        _nw177[(new BigNumber(15)).toNumber()] = _902_v1;
        _nw177[(new BigNumber(16)).toNumber()] = _dafny.Seq.Concat(_902_v1, _902_v1);
        _nw177[(new BigNumber(17)).toNumber()] = _dafny.Seq.of(_901_v0);
        _nw177[(new BigNumber(18)).toNumber()] = _dafny.Seq.Concat(_902_v1, _902_v1);
        _nw177[(new BigNumber(19)).toNumber()] = _module.__default.fm7(globalState);
        _nw177[(new BigNumber(20)).toNumber()] = _902_v1;
        _nw177[(new BigNumber(21)).toNumber()] = _902_v1;
        _934_v28 = _nw177;
        let _939_v31;
        _939_v31 = _dafny.Seq.of((_dafny.ZERO).minus(p0), new BigNumber((_dafny.Seq.of(_901_v0, false)).length));
        let _rhs186 = _module.__default.safeDivisionInt((_939_v31)[_module.__default.safeIndex(p0, new BigNumber((_939_v31).length))], p0);
        let _rhs187 = _901_v0;
        let _rhs188 = _901_v0;
        let _lhs157 = globalState;
        let _lhs158 = globalState;
        _lhs157.f10 = _rhs186;
        _lhs158.f18 = _rhs187;
        _901_v0 = _rhs188;
        r1 = p0;
      } else {
        let _940_v32;
        _940_v32 = _dafny.Seq.UnicodeFromString("kysw");
        let _941_v33;
        _941_v33 = _module.D9.create_DC17(new BigNumber((_940_v32).length));
        let _942_v34;
        _942_v34 = _dafny.Seq.of(p0);
        let _943_v35;
        _943_v35 = _dafny.Map.Empty.slice().updateUnsafe(_941_v33,_dafny.Seq.contains(_942_v34, p0));
        _943_v35 = ((_dafny.Map.Empty.slice().updateUnsafe(_module.D9.create_DC17(p0),_901_v0)).Merge(_943_v35)).Merge(_module.__default.fm28(p0, !(_901_v0), _901_v0, globalState));
        _901_v0 = false;
        let _944_v36;
        _944_v36 = _dafny.MultiSet.fromElements(_901_v0, _901_v0, true, _901_v0);
        let _945_v37;
        _945_v37 = _dafny.Map.Empty.slice().updateUnsafe(_944_v36,p0);
        let _946_v38;
        _946_v38 = new _dafny.CodePoint('u'.codePointAt(0));
        let _947_v39;
        let _nw178 = new _module.C4();
        _nw178.__ctor(_945_v37, _901_v0, _946_v38, p0);
        _947_v39 = _nw178;
        let _948_v40;
        _948_v40 = _dafny.Seq.of(_947_v39, _947_v39, _947_v39, _947_v39);
        let _949_v41;
        _949_v41 = _948_v40;
        let _950_v42;
        _950_v42 = _dafny.Set.fromElements(_949_v41);
        (globalState).f18 = !(!((_950_v42).IsDisjointFrom(_950_v42)));
        let _951_v44;
        _951_v44 = _dafny.Seq.of(_944_v36);
        let _952_v45;
        let _nw179 = new _module.C4();
        _nw179.__ctor((function () {
          let _coll41 = new _dafny.Map();
          for (const _compr_41 of (_951_v44).Elements) {
            let _953_v43 = _compr_41;
            if (_dafny.Seq.contains(_951_v44, _953_v43)) {
              _coll41.push([_953_v43,new BigNumber(943)]);
            }
          }
          return _coll41;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_944_v36,p0)), _901_v0, _module.__default.fm10(!(_901_v0), _901_v0, _901_v0, _901_v0, globalState), p0);
        _952_v45 = _nw179;
        let _954_v46;
        let _nw180 = new _module.C0();
        _nw180.__ctor();
        _954_v46 = _nw180;
        let _955_v47;
        _955_v47 = _dafny.Map.Empty.slice().updateUnsafe((_952_v45).f30,_954_v46);
        let _rhs189 = p0;
        let _rhs190 = (_947_v39).f24;
        let _rhs191 = _952_v45;
        let _rhs192 = (((_955_v47).contains(_901_v0)) ? ((_955_v47).get(_901_v0)) : (_954_v46));
        let _lhs159 = globalState;
        _lhs159.f11 = _rhs189;
        _946_v38 = _rhs190;
        _952_v45 = _rhs191;
        _954_v46 = _rhs192;
        (globalState).f10 = new BigNumber((_dafny.Seq.Concat(_940_v32, _dafny.Seq.Concat(_940_v32, _dafny.Seq.Create(_module.__default.abs(new BigNumber(564)), ((_956_v38) => function (_957_i4) {
          return _956_v38;
        })(_946_v38))))).length);
      }
      if (_901_v0) {
        let _958_v48;
        _958_v48 = _dafny.Seq.UnicodeFromString("sepaclhxe");
        let _959_v49;
        _959_v49 = new _dafny.CodePoint('g'.codePointAt(0));
        let _960_v50;
        _960_v50 = _dafny.Map.Empty.slice().updateUnsafe(p0,_959_v49);
        let _961_v51;
        let _nw181 = new _module.C5();
        _nw181.__ctor(_module.__default.fm0(new BigNumber((_958_v48).length), !(_901_v0), _901_v0, (_dafny.ZERO).minus(p0), globalState), (((_960_v50).contains(p0)) ? ((_960_v50).get(p0)) : (_959_v49)), p0);
        _961_v51 = _nw181;
        let _962_v52;
        _962_v52 = _dafny.Set.fromElements((p0).isLessThan((_dafny.ZERO).minus(p0)));
        _962_v52 = _962_v52;
        let _963_v53;
        let _nw182 = Array((new BigNumber(18)).toNumber()).fill(_module.D9.Default());
        _963_v53 = _nw182;
        let _964_v54;
        _964_v54 = _dafny.Map.Empty.slice().updateUnsafe((_961_v51).f26,_module.D3.create_DC6(p0, p0, false));
        let _965_v55;
        _965_v55 = _module.D9.create_DC16(_964_v54);
        let _index154 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_963_v53).length));
        (_963_v53)[_index154] = _965_v55;
        let _index155 = _module.__default.safeIndex(new BigNumber(705), new BigNumber((_963_v53).length));
        (_963_v53)[_index155] = _module.D9.create_DC16((_module.D9.create_DC16(_964_v54)).dtor_cf24);
        (globalState).f18 = !(!(_901_v0));
        let _966_v56;
        let _nw183 = Array((new BigNumber(2)).toNumber());
        _nw183[(_dafny.ZERO).toNumber()] = _959_v49;
        _nw183[(_dafny.ONE).toNumber()] = _959_v49;
        _966_v56 = _nw183;
        let _index156 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_966_v56).length));
        (_966_v56)[_index156] = _959_v49;
        let _index157 = _module.__default.safeIndex(new BigNumber(842), new BigNumber((_966_v56).length));
        (_966_v56)[_index157] = _959_v49;
      } else {
        let _967_v57;
        _967_v57 = _dafny.MultiSet.fromElements(p0, p0);
        let _968_v58;
        _968_v58 = _module.D7.create_DC14(_901_v0, p0, p0);
        let _source14 = ((!((_967_v57).equals(_967_v57))) ? (_968_v58) : (_968_v58));
        if (_source14.is_DC14) {
          let _969___mcc_h0 = (_source14).cf20;
          let _970___mcc_h1 = (_source14).cf21;
          let _971___mcc_h2 = (_source14).cf22;
          let _972_cf22 = _971___mcc_h2;
          let _973_cf21 = _970___mcc_h1;
          let _974_cf20 = _969___mcc_h0;
          let _975_v59;
          _975_v59 = new _dafny.CodePoint('w'.codePointAt(0));
          _975_v59 = (((_974_cf20) || (_901_v0)) ? (_975_v59) : (_975_v59));
          _975_v59 = ((false) ? (_975_v59) : (_975_v59));
          (globalState).f18 = !(false) || ((_974_cf20) && (_901_v0));
          let _976_v60;
          let _nw184 = Array((new BigNumber(15)).toNumber()).fill([]);
          _976_v60 = _nw184;
          let _977_v61;
          _977_v61 = _dafny.MultiSet.fromElements(_974_cf20);
          let _978_v62;
          _978_v62 = _dafny.Map.Empty.slice().updateUnsafe(_977_v61,_972_cf22);
          let _979_v63;
          let _nw185 = new _module.C4();
          _nw185.__ctor(_978_v62, _901_v0, _975_v59, _972_cf22);
          _979_v63 = _nw185;
          let _980_v64;
          _980_v64 = _dafny.Map.Empty.slice().updateUnsafe(_979_v63,(_979_v63).f30);
          let _981_v65;
          _981_v65 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(596),_980_v64);
          let _982_v66;
          _982_v66 = _dafny.Seq.UnicodeFromString("x");
          let _983_v67;
          _983_v67 = _dafny.Set.fromElements(new BigNumber(440));
          let _984_v68;
          _984_v68 = _dafny.Seq.of((_979_v63).f30, true, (_979_v63).f30, _901_v0, _974_cf20);
          let _985_v69;
          let _nw186 = Array((new BigNumber(8)).toNumber());
          _nw186[(_dafny.ZERO).toNumber()] = new BigNumber((_dafny.Seq.of(false, _974_cf20, _974_cf20)).length);
          _nw186[(_dafny.ONE).toNumber()] = p0;
          _nw186[(new BigNumber(2)).toNumber()] = new BigNumber((_dafny.MultiSet.fromElements(_974_cf20)).cardinality());
          _nw186[(new BigNumber(3)).toNumber()] = new BigNumber((_981_v65).length);
          _nw186[(new BigNumber(4)).toNumber()] = new BigNumber((_982_v66).length);
          _nw186[(new BigNumber(5)).toNumber()] = new BigNumber((_983_v67).length);
          _nw186[(new BigNumber(6)).toNumber()] = _module.__default.fm17(_901_v0, (_979_v63).f30, _module.__default.fm0(new BigNumber((_984_v68).length), _974_cf20, _974_cf20, p0, globalState), globalState);
          _nw186[(new BigNumber(7)).toNumber()] = new BigNumber((_984_v68).length);
          _985_v69 = _nw186;
          let _index158 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_976_v60).length));
          (_976_v60)[_index158] = _985_v69;
          let _986_v70;
          _986_v70 = _module.D5.create_DC9(_985_v69);
          let _pat_let_tv9 = _985_v69;
          let _index159 = _module.__default.safeIndex(new BigNumber(669), new BigNumber((_976_v60).length));
          (_976_v60)[_index159] = (function (_pat_let12_0) {
            return function (_987_dt__update__tmp_h0) {
              return function (_pat_let13_0) {
                return function (_988_dt__update_hcf13_h0) {
                  return _module.D5.create_DC9(_988_dt__update_hcf13_h0);
                }(_pat_let13_0);
              }(_pat_let_tv9);
            }(_pat_let12_0);
          }(_986_v70)).dtor_cf13;
        } else {
          let _989___mcc_h3 = (_source14).cf19;
          let _990_cf19 = _989___mcc_h3;
          let _991_v71;
          let _nw187 = Array((new BigNumber(4)).toNumber());
          _nw187[(_dafny.ZERO).toNumber()] = _901_v0;
          _nw187[(_dafny.ONE).toNumber()] = _901_v0;
          _nw187[(new BigNumber(2)).toNumber()] = _901_v0;
          _nw187[(new BigNumber(3)).toNumber()] = _901_v0;
          _991_v71 = _nw187;
          let _992_v72;
          _992_v72 = _dafny.Map.Empty.slice().updateUnsafe(_901_v0,p0);
          let _993_v73;
          _993_v73 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
          let _994_v74;
          _994_v74 = _dafny.Map.Empty.slice().updateUnsafe(_991_v71,(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_992_v72).length),new BigNumber((_993_v73).length))).Merge(_993_v73));
          let _995_v75;
          let _nw188 = Array((new BigNumber(19)).toNumber()).fill([]);
          _995_v75 = _nw188;
          let _996_v76;
          let _init30 = function (_997_i5) {
            return new _dafny.CodePoint('q'.codePointAt(0));
          };
          let _nw189 = Array((new BigNumber(14)).toNumber());
          for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw189.length); _i0_30++) {
            _nw189[_i0_30] = _init30(new BigNumber(_i0_30));
          }
          _996_v76 = _nw189;
          let _index160 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_995_v75).length));
          (_995_v75)[_index160] = _996_v76;
          let _998_v77;
          _998_v77 = _dafny.Seq.of(_994_v74, _994_v74, _dafny.Map.Empty.slice().updateUnsafe(_991_v71,(_993_v73).update(p0, (_module.D11.create_DC24(_module.__default.fm0(p0, _901_v0, _901_v0, p0, globalState), p0)).dtor_cf37)), _994_v74);
          let _999_v78;
          _999_v78 = _dafny.Seq.of((_994_v74).Merge(_994_v74), (_994_v74).Merge(_994_v74), (_998_v77)[_module.__default.safeIndex((_dafny.ZERO).minus(p0), new BigNumber((_998_v77).length))], _994_v74);
          let _1000_v79;
          let _nw190 = Array((new BigNumber(12)).toNumber());
          _nw190[(_dafny.ZERO).toNumber()] = p0;
          _nw190[(_dafny.ONE).toNumber()] = p0;
          _nw190[(new BigNumber(2)).toNumber()] = p0;
          _nw190[(new BigNumber(3)).toNumber()] = p0;
          _nw190[(new BigNumber(4)).toNumber()] = p0;
          _nw190[(new BigNumber(5)).toNumber()] = new BigNumber(838);
          _nw190[(new BigNumber(6)).toNumber()] = p0;
          _nw190[(new BigNumber(7)).toNumber()] = p0;
          _nw190[(new BigNumber(8)).toNumber()] = p0;
          _nw190[(new BigNumber(9)).toNumber()] = p0;
          _nw190[(new BigNumber(10)).toNumber()] = p0;
          _nw190[(new BigNumber(11)).toNumber()] = p0;
          _1000_v79 = _nw190;
          let _1001_v80;
          _1001_v80 = _module.D13.create_DC28(_dafny.Set.fromElements(_1000_v79));
          let _1002_v81;
          _1002_v81 = _module.D5.create_DC9(_1000_v79);
          let _1003_v82;
          _1003_v82 = _dafny.Set.fromElements(_1000_v79, (_1002_v81).dtor_cf13, _1000_v79, _1000_v79);
          let _1004_v83;
          _1004_v83 = _dafny.MultiSet.fromElements(_901_v0);
          let _1005_v84;
          _1005_v84 = _dafny.Seq.UnicodeFromString("sc");
          let _1006_v85;
          _1006_v85 = _module.D12.create_DC27((((_967_v57).contains(p0)) ? ((_967_v57).get(p0)) : ((((_1004_v83).contains(_901_v0)) ? ((_1004_v83).get(_901_v0)) : (new BigNumber((_1005_v84).length))))), _901_v0);
          let _1007_v86;
          _1007_v86 = _dafny.Map.Empty.slice().updateUnsafe(false,_901_v0);
          let _1008_v87;
          _1008_v87 = _dafny.Map.Empty.slice().updateUnsafe(p0,_1007_v86);
          let _index161 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_995_v75).length));
          let _rhs193 = (_999_v78)[_module.__default.safeIndex(p0, new BigNumber((_999_v78).length))];
          let _rhs194 = _996_v76;
          let _rhs195 = p0;
          let _rhs196 = ((_1001_v80).dtor_cf42).IsDisjointFrom(_1003_v82);
          let _rhs197 = (_module.D7.create_DC14((_1006_v85).dtor_cf41, p0, new BigNumber(((((_1008_v87).contains(p0)) ? ((_1008_v87).get(p0)) : (_1007_v86))).length))).dtor_cf20;
          let _lhs160 = _995_v75;
          let _lhs161 = _module.__default.safeIndex(new BigNumber(903), new BigNumber((_995_v75).length));
          let _lhs162 = globalState;
          let _lhs163 = globalState;
          let _lhs164 = globalState;
          _994_v74 = _rhs193;
          _lhs160[_lhs161] = _rhs194;
          _lhs162.f16 = _rhs195;
          _lhs163.f18 = _rhs196;
          _lhs164.f18 = _rhs197;
          let _1009_v88;
          _1009_v88 = new _dafny.CodePoint('k'.codePointAt(0));
          _1009_v88 = _module.__default.fm10(_901_v0, (false) || (_901_v0), _901_v0, _901_v0, globalState);
          let _1010_v89;
          _1010_v89 = _dafny.MultiSet.fromElements(_module.__default.fm29(p0, globalState));
          _993_v73 = _module.__default.fm11(_1009_v88, (_this).fm1(_1010_v89, globalState), globalState);
          let _index162 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_991_v71).length));
          (_991_v71)[_index162] = !(_901_v0);
          let _index163 = _module.__default.safeIndex(new BigNumber(901), new BigNumber((_991_v71).length));
          (_991_v71)[_index163] = !(_901_v0);
        }
        let _1011_v90;
        _1011_v90 = _dafny.Seq.UnicodeFromString("sc");
        let _1012_v91;
        _1012_v91 = _dafny.Seq.of(new BigNumber((_1011_v90).length));
        (globalState).f11 = _module.__default.safeDivisionInt(new BigNumber((_1012_v91).length), p0);
        let _1013_v92;
        _1013_v92 = _module.D11.create_DC24(_901_v0, p0);
        let _1014_v93;
        _1014_v93 = _dafny.Seq.of((_1013_v92).dtor_cf36);
        if (_dafny.Seq.IsProperPrefixOf(_1014_v93, _1014_v93)) {
          (globalState).f11 = p0;
          let _1015_v94;
          let _nw191 = new _module.C2();
          _nw191.__ctor(p0, _901_v0);
          _1015_v94 = _nw191;
          let _1016_v95;
          _1016_v95 = _dafny.Map.Empty.slice().updateUnsafe(_901_v0,_1015_v94);
          let _1017_v96;
          _1017_v96 = _dafny.Map.Empty.slice().updateUnsafe(_901_v0,_901_v0);
          let _1018_v97;
          _1018_v97 = _dafny.Set.fromElements(_1011_v90);
          let _1019_v98;
          _1019_v98 = _dafny.MultiSet.fromElements(_901_v0);
          let _1020_v99;
          _1020_v99 = _dafny.Seq.UnicodeFromString("tkymklej");
          let _1021_v100;
          _1021_v100 = _dafny.MultiSet.fromElements(_1020_v99, _1011_v90, _1011_v90, _1020_v99, _1020_v99);
          let _1022_v101;
          let _nw192 = Array((new BigNumber(20)).toNumber());
          _nw192[(_dafny.ZERO).toNumber()] = !(new BigNumber((_1016_v95).length)).isEqualTo((_1015_v94).f27);
          _nw192[(_dafny.ONE).toNumber()] = (_1015_v94).f28;
          _nw192[(new BigNumber(2)).toNumber()] = (((_1017_v96).contains(true)) ? ((_1017_v96).get(true)) : ((_1015_v94).f28));
          _nw192[(new BigNumber(3)).toNumber()] = (_1015_v94).f28;
          _nw192[(new BigNumber(4)).toNumber()] = (_901_v0) || ((_1015_v94).f28);
          _nw192[(new BigNumber(5)).toNumber()] = (_1015_v94).f28;
          _nw192[(new BigNumber(6)).toNumber()] = _901_v0;
          _nw192[(new BigNumber(7)).toNumber()] = !_dafny.Seq.contains(_1014_v93, (_1015_v94).f28);
          _nw192[(new BigNumber(8)).toNumber()] = (_1015_v94).f28;
          _nw192[(new BigNumber(9)).toNumber()] = (_1015_v94).f28;
          _nw192[(new BigNumber(10)).toNumber()] = _901_v0;
          _nw192[(new BigNumber(11)).toNumber()] = (_1018_v97).IsDisjointFrom(_1018_v97);
          _nw192[(new BigNumber(12)).toNumber()] = (_1015_v94).f28;
          _nw192[(new BigNumber(13)).toNumber()] = _module.__default.fm0(p0, (_1015_v94).f28, _901_v0, p0, globalState);
          _nw192[(new BigNumber(14)).toNumber()] = (_1019_v98).IsDisjointFrom(_1019_v98);
          _nw192[(new BigNumber(15)).toNumber()] = _901_v0;
          _nw192[(new BigNumber(16)).toNumber()] = !((_this).fm1(_1021_v100, globalState));
          _nw192[(new BigNumber(17)).toNumber()] = _901_v0;
          _nw192[(new BigNumber(18)).toNumber()] = (p0).isLessThan(new BigNumber(348));
          _nw192[(new BigNumber(19)).toNumber()] = !(((_1015_v94).f28) && (false));
          _1022_v101 = _nw192;
          let _index164 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_1022_v101).length));
          (_1022_v101)[_index164] = _901_v0;
          let _index165 = _module.__default.safeIndex(new BigNumber(517), new BigNumber((_1022_v101).length));
          (_1022_v101)[_index165] = !(_901_v0) || (false);
          let _1023_v102;
          _1023_v102 = _module.D9.create_DC17(new BigNumber(338));
          let _1024_v103;
          let _init31 = ((_1025_v94, _1026_p0, _1027_v91) => function (_1028_i6) {
            return _dafny.Seq.Concat(_dafny.Seq.of((_1025_v94).f27, _1026_p0), _1027_v91);
          })(_1015_v94, p0, _1012_v91);
          let _nw193 = Array((new BigNumber(23)).toNumber());
          for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw193.length); _i0_31++) {
            _nw193[_i0_31] = _init31(new BigNumber(_i0_31));
          }
          _1024_v103 = _nw193;
          let _1029_v104;
          _1029_v104 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(820)), ((_1030_v94) => function (_1031_i7) {
            return (_1030_v94).f27;
          })(_1015_v94));
          let _index166 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_1024_v103).length));
          (_1024_v103)[_index166] = (_1029_v104);
          let _1032_v105;
          _1032_v105 = _module.D13.create_DC29((_1022_v101)[_module.__default.safeIndex(new BigNumber(517), new BigNumber((_1022_v101).length))], _901_v0, _1022_v101);
          let _1033_v106;
          _1033_v106 = _dafny.Map.Empty.slice().updateUnsafe((_1032_v105).dtor_cf45,(_1015_v94).f28);
          let _1034_v107;
          _1034_v107 = _dafny.Seq.of(_1019_v98);
          let _index167 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_1024_v103).length));
          let _rhs198 = (((((_1033_v106).contains(_1022_v101)) ? ((_1033_v106).get(_1022_v101)) : ((_1015_v94).f28))) ? (_1023_v102) : (_module.D9.create_DC17(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_1022_v101)[_module.__default.safeIndex(new BigNumber(517), new BigNumber((_1022_v101).length))],(_1015_v94).f28)).length))));
          let _rhs199 = _dafny.Seq.of(p0);
          let _rhs200 = ((((_1015_v94).f28) ? (new BigNumber((_1034_v107).length)) : ((_1015_v94).f27))).minus(p0);
          let _rhs201 = _1022_v101;
          let _lhs165 = _1024_v103;
          let _lhs166 = _module.__default.safeIndex(new BigNumber(24), new BigNumber((_1024_v103).length));
          let _lhs167 = globalState;
          _1023_v102 = _rhs198;
          _lhs165[_lhs166] = _rhs199;
          _lhs167.f16 = _rhs200;
          _1022_v101 = _rhs201;
          (globalState).f18 = (p0).isLessThan(p0);
          (globalState).f0 = (_dafny.ZERO).minus((_1015_v94).f27);
        } else {
          let _1035_v108;
          let _nw194 = Array((new BigNumber(21)).toNumber());
          _nw194[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_1012_v91, _1012_v91);
          _nw194[(_dafny.ONE).toNumber()] = _dafny.Seq.Concat(_1012_v91, _1012_v91);
          _nw194[(new BigNumber(2)).toNumber()] = _dafny.Seq.of((_dafny.ZERO).minus(p0));
          _nw194[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_1012_v91, _1012_v91);
          _nw194[(new BigNumber(4)).toNumber()] = _dafny.Seq.update(_dafny.Seq.of(p0), _module.__default.safeIndex(_module.__default.fm5(new BigNumber((_967_v57).cardinality()), p0, globalState), new BigNumber((_dafny.Seq.of(p0)).length)), new BigNumber((_1012_v91).length));
          _nw194[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_1012_v91, _1012_v91);
          _nw194[(new BigNumber(6)).toNumber()] = _1012_v91;
          _nw194[(new BigNumber(7)).toNumber()] = _1012_v91;
          _nw194[(new BigNumber(8)).toNumber()] = _1012_v91;
          _nw194[(new BigNumber(9)).toNumber()] = _1012_v91;
          _nw194[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_1012_v91, _dafny.Seq.Create(_module.__default.abs(new BigNumber(229)), ((_1036_p0) => function (_1037_i8) {
            return _1036_p0;
          })(p0)));
          _nw194[(new BigNumber(11)).toNumber()] = _1012_v91;
          _nw194[(new BigNumber(12)).toNumber()] = _1012_v91;
          _nw194[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_dafny.Seq.Concat(_1012_v91, _1012_v91), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.Concat(_1012_v91, _1012_v91)).length)), p0);
          _nw194[(new BigNumber(14)).toNumber()] = _1012_v91;
          _nw194[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(p0);
          _nw194[(new BigNumber(16)).toNumber()] = _1012_v91;
          _nw194[(new BigNumber(17)).toNumber()] = _1012_v91;
          _nw194[(new BigNumber(18)).toNumber()] = _1012_v91;
          _nw194[(new BigNumber(19)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(370)), ((_1038_v91, _1039_p0) => function (_1040_i9) {
            return (_1038_v91)[_module.__default.safeIndex(_1039_p0, new BigNumber((_1038_v91).length))];
          })(_1012_v91, p0));
          _nw194[(new BigNumber(20)).toNumber()] = _1012_v91;
          _1035_v108 = _nw194;
          let _index168 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1035_v108).length));
          (_1035_v108)[_index168] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(483)), ((_1041_p0) => function (_1042_i10) {
            return _1041_p0;
          })(p0));
          let _index169 = _module.__default.safeIndex(new BigNumber(306), new BigNumber((_1035_v108).length));
          (_1035_v108)[_index169] = _1012_v91;
          _901_v0 = _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("rbg"), ((_901_v0) ? (_1011_v90) : (_1011_v90)));
          let _1043_v109;
          let _nw195 = new _module.C5();
          _nw195.__ctor(_901_v0, (_1011_v90)[_module.__default.safeIndex(p0, new BigNumber((_1011_v90).length))], p0);
          _1043_v109 = _nw195;
          let _1044_v110;
          let _nw196 = new _module.C0();
          _nw196.__ctor();
          _1044_v110 = _nw196;
          let _1045_v111;
          _1045_v111 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("ocqonr"),(_1043_v109).f26);
          _1045_v111 = (_1045_v111).Merge(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(928)), function (_1046_i11) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          }),!(false)));
        }
        let _1047_v112;
        _1047_v112 = _1011_v90;
        let _1048_v113;
        _1048_v113 = _dafny.Map.Empty.slice().updateUnsafe(_901_v0,_1047_v112);
        let _1049_v114;
        _1049_v114 = _dafny.Map.Empty.slice().updateUnsafe(_1048_v113,_901_v0);
        (globalState).f0 = new BigNumber(((_1049_v114).Merge((_1049_v114).Merge(_module.__default.fm30(new BigNumber(746), p0, _901_v0, _901_v0, globalState)))).length);
        (globalState).f2 = p0;
      }
      let _1050_v115;
      _1050_v115 = new _dafny.CodePoint('p'.codePointAt(0));
      let _1051_v116;
      _1051_v116 = _module.D10.create_DC19(_1050_v115);
      let _1052_v117;
      _1052_v117 = _dafny.Seq.of(_1051_v116);
      let _1053_v118;
      _1053_v118 = _dafny.Seq.UnicodeFromString("okgjhwo");
      let _1054_v119;
      _1054_v119 = _dafny.Map.Empty.slice().updateUnsafe(_1050_v115,_1053_v118);
      let _1055_v120;
      _1055_v120 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.IsProperPrefixOf(_1052_v117, _1052_v117),(_1054_v119).Merge(_1054_v119));
      _1055_v120 = (_1055_v120).update(_901_v0, _1054_v119);
      let _1056_v121;
      _1056_v121 = _dafny.Set.fromElements(_module.__default.fm17(!(false), true, false, globalState));
      let _1057_v122;
      _1057_v122 = _dafny.Map.Empty.slice().updateUnsafe(_901_v0,_901_v0);
      let _1058_v123;
      _1058_v123 = _dafny.MultiSet.fromElements(_module.__default.fm5(new BigNumber((_1057_v122).length), p0, globalState));
      _1056_v121 = function () {
        let _coll42 = new _dafny.Set();
        for (const _compr_42 of (_1058_v123).Elements) {
          let _1059_v124 = _compr_42;
          if ((_1058_v123).contains(_1059_v124)) {
            _coll42.add(_module.__default.safeDivisionInt(_1059_v124, (new BigNumber((_dafny.Seq.UnicodeFromString("e")).length)).plus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_module.D3.create_DC5(_dafny.Seq.of(false), new BigNumber(492), new BigNumber((function () {
  let _coll43 = new _dafny.Set();
  for (const _compr_43 of (_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-941)), new BigNumber((_dafny.Seq.of(!(false))).length), new BigNumber(774), new BigNumber(187))).length))).Elements) {
    let _1060_v125 = _compr_43;
    if ((_dafny.Set.fromElements(new BigNumber((_dafny.Set.fromElements((_dafny.ZERO).minus(new BigNumber(-941)), new BigNumber((_dafny.Seq.of(!(false))).length), new BigNumber(774), new BigNumber(187))).length))).contains(_1060_v125)) {
      _coll43.add((_1060_v125).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("xn")).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(780)), function (_1061_i12) {
        return new _dafny.CodePoint('h'.codePointAt(0));
      })).length))).length)));
    }
  }
  return _coll43;
}()).length)),!(!(false)))).length))));
          }
        }
        return _coll42;
      }();
      let _1062_v126;
      _1062_v126 = _dafny.Seq.of(_901_v0, !(_901_v0) || (_901_v0), _901_v0, _901_v0);
      if ((_1062_v126)[_module.__default.safeIndex((p0).plus(p0), new BigNumber((_1062_v126).length))]) {
        (globalState).f0 = new BigNumber(330);
        let _1063_v127;
        let _nw197 = new _module.C0();
        _nw197.__ctor();
        _1063_v127 = _nw197;
        r2 = _dafny.Seq.Concat(_1053_v118, _1053_v118);
        _1058_v123 = _1058_v123;
        let _1064_v128;
        _1064_v128 = _dafny.Map.Empty.slice().updateUnsafe(_901_v0,p0);
        (globalState).f15 = ((((_1064_v128).contains(_901_v0)) ? ((_1064_v128).get(_901_v0)) : (new BigNumber((_dafny.Seq.update(_1053_v118, _module.__default.safeIndex(p0, new BigNumber((_1053_v118).length)), _1050_v115)).length)))).plus(p0);
      } else {
        let _1065_v129;
        let _nw198 = new _module.C3();
        _nw198.__ctor(true, new _dafny.CodePoint('r'.codePointAt(0)), p0);
        _1065_v129 = _nw198;
        let _1066_v130;
        _1066_v130 = _dafny.Map.Empty.slice().updateUnsafe(_1065_v129,(_1065_v129).fm18(globalState));
        (globalState).f18 = !(new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(_1065_v129,true)).Merge(_1066_v130)).length)).isEqualTo(p0);
        let _1067_v131;
        _1067_v131 = _dafny.Seq.of(p0, p0);
        let _1068_v132;
        let _nw199 = new _module.C1();
        _nw199.__ctor(_1050_v115, new BigNumber((_1067_v131).length));
        _1068_v132 = _nw199;
        let _1069_v133;
        _1069_v133 = _dafny.Map.Empty.slice().updateUnsafe(true,_1068_v132);
        let _1070_v134;
        _1070_v134 = _dafny.Set.fromElements((((_1069_v133).contains(_1065_v129.f31)) ? ((_1069_v133).get(_1065_v129.f31)) : (_1068_v132)), _1068_v132);
        let _1071_v135;
        _1071_v135 = _module.D7.create_DC13((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1070_v134).length),_901_v0)).update((_dafny.ZERO).minus(p0), _1065_v129.f31));
        let _pat_let_tv10 = _1054_v119;
        let _pat_let_tv11 = _901_v0;
        let _rhs202 = function (_pat_let14_0) {
          return function (_1072_dt__update__tmp_h1) {
            return function (_pat_let15_0) {
              return function (_1073_dt__update_hcf19_h0) {
                return _module.D7.create_DC13(_1073_dt__update_hcf19_h0);
              }(_pat_let15_0);
            }(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_pat_let_tv10).length),_pat_let_tv11));
          }(_pat_let14_0);
        }(_1071_v135);
        let _rhs203 = _1065_v129.f31;
        let _lhs168 = _1065_v129;
        _1071_v135 = _rhs202;
        _lhs168.f31 = _rhs203;
        _901_v0 = !(_1065_v129.f31) || (((_1065_v129.f31) ? (_1065_v129.f31) : (_1065_v129.f31)));
        (globalState).f18 = _1065_v129.f31;
        let _1074_v136;
        let _nw200 = new _module.C1();
        _nw200.__ctor(_1050_v115, p0);
        _1074_v136 = _nw200;
      }
      let _1075_v137;
      _1075_v137 = _module.D9.create_DC16(_dafny.Map.Empty.slice().updateUnsafe(_901_v0,_module.D3.create_DC6(p0, p0, !(_901_v0))));
      let _pat_let_tv12 = _1056_v121;
      let _pat_let_tv13 = _901_v0;
      if (function (_source15) {
        if (_source15.is_DC17) {
          let _1076___mcc_h5 = (_source15).cf25;
          let _1077_cf25 = _1076___mcc_h5;
          return (new BigNumber(619)).isLessThan(new BigNumber((_pat_let_tv12).length));
        } else {
          let _1078___mcc_h6 = (_source15).cf24;
          let _1079_cf24 = _1078___mcc_h6;
          return _pat_let_tv13;
        }
      }(_1075_v137)) {
        let _1080_v138;
        _1080_v138 = _1053_v118;
        let _source16 = _1080_v138;
        let _1081___mcc_h4 = _source16;
        let _1082_cf1 = _1081___mcc_h4;
        let _1083_v139;
        let _nw201 = Array((new BigNumber(14)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
        _1083_v139 = _nw201;
        let _1084_v140;
        _1084_v140 = _dafny.Seq.of(_module.__default.fm5(p0, p0, globalState));
        let _1085_v141;
        _1085_v141 = _dafny.Map.Empty.slice().updateUnsafe(_1083_v139,((true) ? (new BigNumber((_1084_v140).length)) : (p0)));
        r1 = new BigNumber((_1085_v141).length);
        _1058_v123 = (_1058_v123).Difference(_1058_v123);
        (_this).f23 = (_this.f23).update(new BigNumber((_1082_cf1).length), _901_v0);
        r2 = _1053_v118;
        r1 = (p0).plus(_module.__default.safeModuloInt(p0, (_dafny.ZERO).minus(new BigNumber(-498))));
        (globalState).f18 = !(!(true));
        let _1086_v142;
        let _nw202 = Array((new BigNumber(4)).toNumber());
        _nw202[(_dafny.ZERO).toNumber()] = _901_v0;
        _nw202[(_dafny.ONE).toNumber()] = true;
        _nw202[(new BigNumber(2)).toNumber()] = _901_v0;
        _nw202[(new BigNumber(3)).toNumber()] = _901_v0;
        _1086_v142 = _nw202;
        let _1087_v143;
        _1087_v143 = _dafny.Seq.of(_1086_v142);
        let _1088_v144;
        _1088_v144 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1087_v143).length),_1050_v115);
        _1088_v144 = (_1088_v144).update(new BigNumber((_1053_v118).length), _module.__default.fm10(_901_v0, _901_v0, false, _901_v0, globalState));
        let _1089_v145;
        let _nw203 = Array((new BigNumber(9)).toNumber()).fill(_dafny.ZERO);
        _1089_v145 = _nw203;
        let _1090_v146;
        _1090_v146 = _dafny.Map.Empty.slice().updateUnsafe(_901_v0,_1089_v145);
        if (!((_1090_v146).Merge(_1090_v146)).equals((_1090_v146).Merge((_1090_v146).update(false, _1089_v145)))) {
          let _1091_v147;
          _1091_v147 = _dafny.MultiSet.fromElements(_901_v0, _901_v0);
          _1091_v147 = (_1091_v147).Difference(_1091_v147);
          (globalState).f18 = _901_v0;
          let _1092_v148;
          let _init32 = ((_1093_v118) => function (_1094_i13) {
            return _1093_v118;
          })(_1053_v118);
          let _nw204 = Array((new BigNumber(16)).toNumber());
          for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw204.length); _i0_32++) {
            _nw204[_i0_32] = _init32(new BigNumber(_i0_32));
          }
          _1092_v148 = _nw204;
          let _index170 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1092_v148).length));
          (_1092_v148)[_index170] = _1053_v118;
          let _index171 = _module.__default.safeIndex(new BigNumber(110), new BigNumber((_1092_v148).length));
          (_1092_v148)[_index171] = _1053_v118;
          let _1095_v149;
          let _nw205 = Array((new BigNumber(8)).toNumber()).fill(_dafny.Set.Empty);
          _1095_v149 = _nw205;
          let _1096_v150;
          _1096_v150 = _dafny.Seq.of(_1056_v121, _dafny.Set.fromElements((_dafny.ZERO).minus(p0)));
          let _1097_v151;
          _1097_v151 = _dafny.Seq.of(p0, p0);
          let _1098_v152;
          _1098_v152 = _dafny.Seq.of(_1097_v151);
          let _1099_v153;
          _1099_v153 = _dafny.Seq.of((_1098_v152)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1057_v122).length)), new BigNumber((_1098_v152).length))], _dafny.Seq.of(p0));
          let _1100_v154;
          _1100_v154 = _dafny.Seq.of(_1056_v121, _1056_v121, (_1096_v150)[_module.__default.safeIndex(new BigNumber(433), new BigNumber((_1096_v150).length))], _dafny.Set.fromElements(new BigNumber((_1098_v152).length), p0), _dafny.Set.fromElements(p0, p0));
          let _1101_v155;
          _1101_v155 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(_901_v0)).cardinality()),p0);
          let _index172 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1095_v149).length));
          (_1095_v149)[_index172] = (_1096_v150)[_module.__default.safeIndex(new BigNumber((_1101_v155).length), new BigNumber((_1096_v150).length))];
          let _index173 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1086_v142).length));
          (_1086_v142)[_index173] = _901_v0;
          let _1102_v156;
          _1102_v156 = _dafny.Seq.of(_this.f23, _this.f23, _this.f23);
          let _index174 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1095_v149).length));
          let _index175 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1086_v142).length));
          let _rhs204 = (_1056_v121).Intersect((_1056_v121).Union(_dafny.Set.fromElements(p0, p0)));
          let _rhs205 = new BigNumber(374);
          let _rhs206 = _dafny.Seq.contains(_1102_v156, (_this.f23).Merge(_this.f23));
          let _lhs169 = _1095_v149;
          let _lhs170 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((_1095_v149).length));
          let _lhs171 = globalState;
          let _lhs172 = _1086_v142;
          let _lhs173 = _module.__default.safeIndex(new BigNumber(802), new BigNumber((_1086_v142).length));
          _lhs169[_lhs170] = _rhs204;
          _lhs171.f10 = _rhs205;
          _lhs172[_lhs173] = _rhs206;
          (globalState).f10 = new BigNumber(((((_1086_v142)[_module.__default.safeIndex(new BigNumber(802), new BigNumber((_1086_v142).length))]) ? (_1091_v147) : (_1091_v147))).cardinality());
        } else {
          let _index176 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_1086_v142).length));
          (_1086_v142)[_index176] = _901_v0;
          let _index177 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_1086_v142).length));
          (_1086_v142)[_index177] = !(_901_v0);
          let _1103_v157;
          _1103_v157 = _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('g'.codePointAt(0)),p0);
          let _1104_v158;
          _1104_v158 = _dafny.Seq.of(_1103_v157);
          let _1105_v159;
          _1105_v159 = _dafny.Seq.of(new BigNumber(794), new BigNumber(((_1103_v157).Merge((_1104_v158)[_module.__default.safeIndex(p0, new BigNumber((_1104_v158).length))])).length), p0, p0);
          (globalState).f15 = (_1105_v159)[_module.__default.safeIndex(p0, new BigNumber((_1105_v159).length))];
          let _1106_v160;
          let _nw206 = new _module.C0();
          _nw206.__ctor();
          _1106_v160 = _nw206;
          let _index178 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_1086_v142).length));
          let _rhs207 = _901_v0;
          let _rhs208 = _dafny.Seq.Concat(_1105_v159, _1105_v159);
          let _rhs209 = _1106_v160;
          let _rhs210 = (_1086_v142)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_1086_v142).length))];
          let _rhs211 = (_1058_v123).IsDisjointFrom((_1058_v123).update(p0, _module.__default.abs((_dafny.ZERO).minus(new BigNumber((_module.__default.fm31(new BigNumber((_1105_v159).length), (_1086_v142)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_1086_v142).length))], _901_v0, new BigNumber((_1053_v118).length), globalState)).length)))));
          let _lhs174 = globalState;
          let _lhs175 = _1086_v142;
          let _lhs176 = _module.__default.safeIndex(new BigNumber(262), new BigNumber((_1086_v142).length));
          let _lhs177 = globalState;
          _901_v0 = _rhs207;
          _lhs174.f19 = _rhs208;
          _1106_v160 = _rhs209;
          _lhs175[_lhs176] = _rhs210;
          _lhs177.f18 = _rhs211;
          let _1107_v161;
          let _nw207 = Array((new BigNumber(6)).toNumber());
          _nw207[(_dafny.ZERO).toNumber()] = (_1086_v142)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_1086_v142).length))];
          _nw207[(_dafny.ONE).toNumber()] = _901_v0;
          _nw207[(new BigNumber(2)).toNumber()] = (_1086_v142)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_1086_v142).length))];
          _nw207[(new BigNumber(3)).toNumber()] = _901_v0;
          _nw207[(new BigNumber(4)).toNumber()] = false;
          _nw207[(new BigNumber(5)).toNumber()] = false;
          _1107_v161 = _nw207;
          let _1108_v162;
          _1108_v162 = _dafny.Map.Empty.slice().updateUnsafe(_1107_v161,new BigNumber((_1058_v123).cardinality()));
          _1108_v162 = (_1108_v162).update(_1107_v161, (p0).minus((_dafny.ZERO).minus(p0)));
          let _1109_v163;
          let _nw208 = Array((new BigNumber(6)).toNumber()).fill(_dafny.Map.Empty);
          _1109_v163 = _nw208;
          let _1110_v164;
          _1110_v164 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.UnicodeFromString("oo"),(_1086_v142)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_1086_v142).length))]);
          let _index179 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_1109_v163).length));
          (_1109_v163)[_index179] = _1110_v164;
          let _index180 = _module.__default.safeIndex(new BigNumber(958), new BigNumber((_1109_v163).length));
          (_1109_v163)[_index180] = ((_1110_v164).update(_1053_v118, (_1086_v142)[_module.__default.safeIndex(new BigNumber(262), new BigNumber((_1086_v142).length))])).Merge(_1110_v164);
        }
      } else {
        let _1111_v165;
        let _nw209 = Array((new BigNumber(29)).toNumber()).fill(false);
        _1111_v165 = _nw209;
        let _index181 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_1111_v165).length));
        (_1111_v165)[_index181] = _901_v0;
        let _index182 = _module.__default.safeIndex(new BigNumber(172), new BigNumber((_1111_v165).length));
        (_1111_v165)[_index182] = !(_901_v0);
        (globalState).f16 = _module.__default.safeModuloInt(new BigNumber((_1057_v122).length), new BigNumber((_dafny.Seq.of((_1111_v165)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1111_v165).length))])).length));
        (globalState).f18 = !(_module.__default.fm0(p0, (p0).isLessThan(p0), false, p0, globalState));
        _1053_v118 = (((_1111_v165)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1111_v165).length))]) ? (_1053_v118) : (_dafny.Seq.Concat(_1053_v118, _1053_v118)));
        let _1112_v166;
        let _nw210 = new _module.C3();
        _nw210.__ctor(_901_v0, _1050_v115, p0);
        _1112_v166 = _nw210;
        _1112_v166 = (((_1111_v165)[_module.__default.safeIndex(new BigNumber(172), new BigNumber((_1111_v165).length))]) ? (_1112_v166) : (_1112_v166));
      }
      r0 = p0;
      r1 = _module.__default.fm17(_901_v0, _901_v0, (_1062_v126)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.update(_1053_v118, _module.__default.safeIndex(new BigNumber(611), new BigNumber((_1053_v118).length)), _module.__default.fm10(_901_v0, _901_v0, _901_v0, _901_v0, globalState))).length), new BigNumber((_1062_v126).length))], globalState);
      r2 = _1053_v118;
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
