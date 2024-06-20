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
    static fm0(p0, p1, p2, globalState) {
      let _source0 = _dafny.Map.Empty.slice().updateUnsafe(true,false);
      let _0___mcc_h0 = _source0;
      let _1_cf18 = _0___mcc_h0;
      return function () {
        let _coll0 = new _dafny.Map();
        for (const _compr_0 of _dafny.IntegerRange(new BigNumber(538), new BigNumber(46))) {
          let _2_v0 = _compr_0;
          if (((new BigNumber(538)).isLessThanOrEqualTo(_2_v0)) && ((_2_v0).isLessThan(new BigNumber(46)))) {
            _coll0.push([(_2_v0).plus(new BigNumber((_dafny.Seq.UnicodeFromString("kdylahw")).length)),_dafny.Set.fromElements(true)]);
          }
        }
        return _coll0;
      }();
    };
    static fm1(p0, p1, p2, p3, globalState) {
      if (false) {
        return (_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(true,!(false)), _dafny.Map.Empty.slice().updateUnsafe(true,false))).IsProperSubsetOf(_dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe(false,false)));
      } else {
        return !(true) || (false);
      }
    };
    static fm2(p0, globalState) {
      return _dafny.Seq.UnicodeFromString("jkyhcsar");
    };
    static fm4(p0, p1, p2, globalState) {
      let _source1 = _module.D17.create_DC41(new BigNumber(-267), !(true));
      if (_source1.is_DC41) {
        let _3___mcc_h0 = (_source1).cf77;
        let _4___mcc_h1 = (_source1).cf78;
        let _5_cf78 = _4___mcc_h1;
        let _6_cf77 = _3___mcc_h0;
        return new _dafny.CodePoint('b'.codePointAt(0));
      } else if (_source1.is_DC40) {
        let _7___mcc_h2 = (_source1).cf76;
        let _8_cf76 = _7___mcc_h2;
        return new _dafny.CodePoint('u'.codePointAt(0));
      } else {
        let _9___mcc_h3 = (_source1).cf79;
        let _10_cf79 = _9___mcc_h3;
        return new _dafny.CodePoint('y'.codePointAt(0));
      }
    };
    static fm7(globalState) {
      return _dafny.Seq.of(new BigNumber(548), new BigNumber(-508));
    };
    static fm8(p0, p1, p2, p3, globalState) {
      let _source2 = _module.D14.create_DC36(true, new BigNumber(110), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Set.fromElements(true, !(!(false)))).length))).length))).cardinality())), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(888),true));
      if (_source2.is_DC36) {
        let _11___mcc_h0 = (_source2).cf68;
        let _12___mcc_h1 = (_source2).cf69;
        let _13___mcc_h2 = (_source2).cf70;
        let _14___mcc_h3 = (_source2).cf71;
        let _15_cf71 = _14___mcc_h3;
        let _16_cf70 = _13___mcc_h2;
        let _17_cf69 = _12___mcc_h1;
        let _18_cf68 = _11___mcc_h0;
        return new _dafny.CodePoint('q'.codePointAt(0));
      } else {
        let _19___mcc_h4 = (_source2).cf67;
        let _20_cf67 = _19___mcc_h4;
        return new _dafny.CodePoint('h'.codePointAt(0));
      }
    };
    static fm9(globalState) {
      let _source3 = _module.D1.create_DC5();
      if (_source3.is_DC5) {
        return (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(!(true), !(false)), _dafny.MultiSet.fromElements(true, false))).Difference(function () {
          let _coll1 = new _dafny.Set();
          for (const _compr_1 of (_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),new BigNumber(213))).Keys.Elements) {
            let _21_v0 = _compr_1;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements(true),new BigNumber(213))).contains(_21_v0)) {
              _coll1.add(_21_v0);
            }
          }
          return _coll1;
        }());
      } else {
        let _22___mcc_h0 = (_source3).cf11;
        let _23_cf11 = _22___mcc_h0;
        return (_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true, true, false, true, false))).Difference(_dafny.Set.fromElements(_dafny.MultiSet.fromElements(true), _dafny.MultiSet.fromElements(false)));
      }
    };
    static fm10(p0, p1, p2, globalState) {
      return ((_dafny.Set.fromElements(new BigNumber(-989))).Union(_dafny.Set.fromElements(new BigNumber(922)))).Difference((_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of(true, true)).length)))).length))).Union(_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber(17), new BigNumber((_dafny.Seq.UnicodeFromString("bbpgrxd")).length))).length),false)).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(791),!(!(true)))).length),new BigNumber(854))).length), new BigNumber(439))));
    };
    static fm12(globalState) {
      let _source4 = _module.D1.create_DC4(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(false)).length)));
      if (_source4.is_DC5) {
        if (true) {
          return new _dafny.CodePoint('u'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('u'.codePointAt(0));
        }
      } else {
        let _24___mcc_h0 = (_source4).cf11;
        let _25_cf11 = _24___mcc_h0;
        if (!(!(false))) {
          return new _dafny.CodePoint('r'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('b'.codePointAt(0));
        }
      }
    };
    static fm13(p0, p1, globalState) {
      let _source5 = _module.D5.create_DC13(true, !(!(false)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-59)), function (_26_i0) {
  return new _dafny.CodePoint('i'.codePointAt(0));
}), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(921)), function (_27_i1) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})).length));
      if (_source5.is_DC12) {
        let _28___mcc_h0 = (_source5).cf20;
        let _29_cf20 = _28___mcc_h0;
        return _module.D0.create_DC1(_29_cf20, false, _29_cf20, !(true));
      } else if (_source5.is_DC13) {
        let _30___mcc_h1 = (_source5).cf21;
        let _31___mcc_h2 = (_source5).cf22;
        let _32___mcc_h3 = (_source5).cf23;
        let _33___mcc_h4 = (_source5).cf24;
        let _34_cf24 = _33___mcc_h4;
        let _35_cf23 = _32___mcc_h3;
        let _36_cf22 = _31___mcc_h2;
        let _37_cf21 = _30___mcc_h1;
        return _module.D0.create_DC1(_34_cf24, _36_cf22, new BigNumber(842), _36_cf22);
      } else if (_source5.is_DC11) {
        let _38___mcc_h5 = (_source5).cf19;
        let _39_cf19 = _38___mcc_h5;
        return _module.D0.create_DC1(new BigNumber(651), !(false), new BigNumber(-746), false);
      } else {
        let _40___mcc_h6 = (_source5).cf25;
        let _41_cf25 = _40___mcc_h6;
        return _module.D0.create_DC1(new BigNumber(114), true, new BigNumber(-419), !(true));
      }
    };
    static fm14(p0, p1, p2, globalState) {
      if (!(!(((true) ? (true) : (true))))) {
        return _dafny.Map.Empty.slice().updateUnsafe((_module.D17.create_DC41(new BigNumber(377), true)).dtor_cf78,true);
      } else {
        return (_dafny.Map.Empty.slice().updateUnsafe(true,false)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true));
      }
    };
    static fm15(p0, p1, globalState) {
      return (_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(false, false)).length), new BigNumber(183))).minus((new BigNumber(879)).minus(new BigNumber((_dafny.Seq.UnicodeFromString("raqgm")).length)));
    };
    static fm19(p0, p1, p2, p3, globalState) {
      return _module.D0.create_DC2(new BigNumber(((_dafny.Set.fromElements(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-412))).cardinality()))).Intersect(function () {
  let _coll2 = new _dafny.Set();
  for (const _compr_2 of _dafny.IntegerRange(new BigNumber(180), new BigNumber(64))) {
    let _42_v0 = _compr_2;
    if (((new BigNumber(180)).isLessThanOrEqualTo(_42_v0)) && ((_42_v0).isLessThan(new BigNumber(64)))) {
      _coll2.add((_42_v0).multipliedBy(new BigNumber(170)));
    }
  }
  return _coll2;
}())).length), new _dafny.CodePoint('f'.codePointAt(0)), new BigNumber((_dafny.Seq.UnicodeFromString("li")).length), new BigNumber((_dafny.Seq.UnicodeFromString("mhhv")).length), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(true)).length), new BigNumber(938)), _dafny.Seq.Create(_module.__default.abs(new BigNumber(288)), function (_43_i0) {
  return new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(false, false, true, true)).length))).length);
}))).length));
    };
    static fm21(p0, globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(!(true))), _dafny.Seq.of(!(!(false))));
    };
    static fm22(globalState) {
      return function () {
        let _coll3 = new _dafny.Set();
        for (const _compr_3 of (_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-362),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ksdltj")).length)))).Keys.Elements) {
          let _44_v0 = _compr_3;
          if ((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-362),_dafny.MultiSet.fromElements(new BigNumber((_dafny.Seq.UnicodeFromString("ksdltj")).length)))).contains(_44_v0)) {
            _coll3.add((_44_v0).minus(_module.__default.safeModuloInt(new BigNumber(428), new BigNumber((_dafny.Seq.of(_dafny.Seq.UnicodeFromString("bdnraex"), _dafny.Seq.UnicodeFromString("lktqteah"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(589)), function (_45_i0) {
              return new _dafny.CodePoint('d'.codePointAt(0));
            }))).length))));
          }
        }
        return _coll3;
      }();
    };
    static fm24(p0, globalState) {
      if ((new BigNumber(814)).isLessThan(new BigNumber((_dafny.Seq.UnicodeFromString("y")).length))) {
        return _module.__default.safeDivisionInt(new BigNumber(-666), (_module.D17.create_DC41(new BigNumber(94), true)).dtor_cf77);
      } else {
        return _module.__default.safeModuloInt((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("hhew")).length)), new BigNumber(885));
      }
    };
    static fm25(p0, p1, p2, globalState) {
      return _dafny.Seq.of(new BigNumber(-408), _module.__default.safeDivisionInt(new BigNumber((_dafny.Seq.of(!(true))).length), new BigNumber(572)), (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(false))).length)), new BigNumber(547));
    };
    static fm26(p0, globalState) {
      return (_dafny.Set.fromElements(false)).Intersect(_dafny.Set.fromElements(false, true, false, true));
    };
    static fm27(globalState) {
      return _dafny.Seq.of(_dafny.Seq.of(!(false), false), _dafny.Seq.of(true, false), _dafny.Seq.of(false), _dafny.Seq.Concat(_dafny.Seq.of(true), _dafny.Seq.of(false)));
    };
    static fm28(p0, p1, p2, globalState) {
      if (!(_dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC19(_dafny.MultiSet.fromElements(false, !(!(false))), new BigNumber(330)),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,!(!(true)))).length))).equals(_dafny.Map.Empty.slice().updateUnsafe(_module.D7.create_DC19(_dafny.MultiSet.fromElements(!(true)), new BigNumber((function () {
  let _coll4 = new _dafny.Set();
  for (const _compr_4 of _dafny.IntegerRange(new BigNumber(-713), new BigNumber(-163))) {
    let _46_v0 = _compr_4;
    if (((new BigNumber(-713)).isLessThanOrEqualTo(_46_v0)) && ((_46_v0).isLessThan(new BigNumber(-163)))) {
      _coll4.add(_module.__default.safeDivisionInt(_46_v0, new BigNumber((_dafny.Seq.of(false)).length)));
    }
  }
  return _coll4;
}()).length)),new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(26), new BigNumber(942))).cardinality())))) {
        if (false) {
          return new _dafny.CodePoint('x'.codePointAt(0));
        } else {
          return new _dafny.CodePoint('r'.codePointAt(0));
        }
      } else if (true) {
        return new _dafny.CodePoint('q'.codePointAt(0));
      } else {
        return new _dafny.CodePoint('p'.codePointAt(0));
      }
    };
    static fm29(p0, p1, globalState) {
      let _source6 = _module.D13.create_DC31(function () {
  let _coll5 = new _dafny.Map();
  for (const _compr_5 of (_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality()))).Elements) {
    let _47_v0 = _compr_5;
    if (_dafny.Seq.contains(_dafny.Seq.of(new BigNumber((_dafny.MultiSet.FromArray(_dafny.Seq.of(false))).cardinality())), _47_v0)) {
      _coll5.push([(_47_v0).multipliedBy(new BigNumber(-611)),false]);
    }
  }
  return _coll5;
}());
      if (_source6.is_DC32) {
        let _48___mcc_h0 = (_source6).cf56;
        let _49___mcc_h1 = (_source6).cf57;
        let _50___mcc_h2 = (_source6).cf58;
        let _51___mcc_h3 = (_source6).cf59;
        let _52___mcc_h4 = (_source6).cf60;
        let _53_cf60 = _52___mcc_h4;
        let _54_cf59 = _51___mcc_h3;
        let _55_cf58 = _50___mcc_h2;
        let _56_cf57 = _49___mcc_h1;
        let _57_cf56 = _48___mcc_h0;
        return _module.D6.create_DC16((_dafny.ZERO).minus(_57_cf56));
      } else if (_source6.is_DC33) {
        let _58___mcc_h5 = (_source6).cf61;
        let _59___mcc_h6 = (_source6).cf62;
        let _60___mcc_h7 = (_source6).cf63;
        let _61___mcc_h8 = (_source6).cf64;
        let _62___mcc_h9 = (_source6).cf65;
        let _63_cf65 = _62___mcc_h9;
        let _64_cf64 = _61___mcc_h8;
        let _65_cf63 = _60___mcc_h7;
        let _66_cf62 = _59___mcc_h6;
        let _67_cf61 = _58___mcc_h5;
        return _module.D6.create_DC16((_63_cf65).f9);
      } else if (_source6.is_DC31) {
        let _68___mcc_h10 = (_source6).cf55;
        let _69_cf55 = _68___mcc_h10;
        return _module.D6.create_DC16((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-504)), function (_70_i0) {
  return new _dafny.CodePoint('f'.codePointAt(0));
})).length)));
      } else {
        let _71___mcc_h11 = (_source6).cf66;
        let _72_cf66 = _71___mcc_h11;
        if (false) {
          return _module.D6.create_DC16(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(33),false)).length));
        } else {
          return _module.D6.create_DC16(new BigNumber(746));
        }
      }
    };
    static fm30(p0, p1, p2, globalState) {
      return _dafny.MultiSet.fromElements(true);
    };
    static fm31(p0, p1, p2, p3, globalState) {
      return _module.D6.create_DC15((_dafny.Set.fromElements(true)).Union(_dafny.Set.fromElements(true)));
    };
    static fm32(globalState) {
      return _module.D0.create_DC3(_module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.MultiSet.fromElements(new BigNumber(-345))).cardinality()),_dafny.Set.fromElements(true))));
    };
    static fm33(globalState) {
      let _source7 = _module.D17.create_DC41(new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), false);
      if (_source7.is_DC41) {
        let _73___mcc_h0 = (_source7).cf77;
        let _74___mcc_h1 = (_source7).cf78;
        let _75_cf78 = _74___mcc_h1;
        let _76_cf77 = _73___mcc_h0;
        if (_75_cf78) {
          return _module.D1.create_DC4(_dafny.Seq.of(_76_cf77, new BigNumber((_dafny.Seq.of(false)).length)));
        } else {
          return _module.D1.create_DC4(_dafny.Seq.of(_76_cf77, _76_cf77, _76_cf77));
        }
      } else if (_source7.is_DC40) {
        let _77___mcc_h2 = (_source7).cf76;
        let _78_cf76 = _77___mcc_h2;
        return _module.D1.create_DC4(_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(458)), function (_79_i0) {
  return new _dafny.CodePoint('n'.codePointAt(0));
})).length), (_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(false, true, true, false, false)).cardinality()))));
      } else {
        let _80___mcc_h3 = (_source7).cf79;
        let _81_cf79 = _80___mcc_h3;
        return _module.D1.create_DC4(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber(128), new BigNumber((_dafny.Seq.of(false)).length), new BigNumber(321))).length), new BigNumber(658), new BigNumber(195), new BigNumber((_dafny.Seq.UnicodeFromString("vhqc")).length), new BigNumber(888)));
      }
    };
    static fm34(globalState) {
      return _dafny.Map.Empty.slice().updateUnsafe(!(function () {
        let _coll6 = new _dafny.Map();
        for (const _compr_6 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(394),true), function () {
          let _coll7 = new _dafny.Map();
          for (const _compr_7 of _dafny.IntegerRange(new BigNumber(246), new BigNumber(118))) {
            let _82_v1 = _compr_7;
            if (((new BigNumber(246)).isLessThanOrEqualTo(_82_v1)) && ((_82_v1).isLessThan(new BigNumber(118)))) {
              _coll7.push([(_82_v1).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)),!(true)]);
            }
          }
          return _coll7;
        }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(70),!(false)))).Elements) {
          let _83_v0 = _compr_6;
          if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(394),true), function () {
            let _coll8 = new _dafny.Map();
            for (const _compr_8 of _dafny.IntegerRange(new BigNumber(246), new BigNumber(118))) {
              let _82_v1 = _compr_8;
              if (((new BigNumber(246)).isLessThanOrEqualTo(_82_v1)) && ((_82_v1).isLessThan(new BigNumber(118)))) {
                _coll8.push([(_82_v1).multipliedBy(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(false,true)).length)),!(true)]);
              }
            }
            return _coll8;
          }(), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(70),!(false))), _83_v0)) {
            _coll6.push([_83_v0,new BigNumber(999)]);
          }
        }
        return _coll6;
      }()).contains((_module.D13.create_DC31(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(_dafny.Seq.of(false), _dafny.Seq.of(true, false))).length),true))).dtor_cf55),(_dafny.ZERO).minus(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-291)), function (_84_i0) {
        return new BigNumber(720);
      })).length),new BigNumber((_dafny.Seq.of(false)).length))).length), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(false),new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Set.fromElements(new BigNumber(-82), new BigNumber((function () {
        let _coll9 = new _dafny.Map();
        for (const _compr_9 of _dafny.IntegerRange(new BigNumber(133), new BigNumber(439))) {
          let _85_v2 = _compr_9;
          if (((new BigNumber(133)).isLessThanOrEqualTo(_85_v2)) && ((_85_v2).isLessThan(new BigNumber(439)))) {
            _coll9.push([(_85_v2).minus(new BigNumber(214)),new BigNumber((_dafny.Set.fromElements(false)).length)]);
          }
        }
        return _coll9;
      }()).length))).length),!(true))).length))).length))));
    };
    static fm35(p0, globalState) {
      return ((function () {
        let _coll10 = new _dafny.Map();
        for (const _compr_10 of _dafny.IntegerRange(new BigNumber(735), new BigNumber(-829))) {
          let _86_v0 = _compr_10;
          if (((new BigNumber(735)).isLessThanOrEqualTo(_86_v0)) && ((_86_v0).isLessThan(new BigNumber(-829)))) {
            _coll10.push([_module.__default.safeModuloInt(_86_v0, new BigNumber(520)),false]);
          }
        }
        return _coll10;
      }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length))).length),false))).Merge((_dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(679)), function (_87_i0) {
        return new _dafny.CodePoint('c'.codePointAt(0));
      })).length)),true)).Merge(function () {
        let _coll11 = new _dafny.Map();
        for (const _compr_11 of _dafny.IntegerRange(new BigNumber(120), new BigNumber(-129))) {
          let _88_v1 = _compr_11;
          if (((new BigNumber(120)).isLessThanOrEqualTo(_88_v1)) && ((_88_v1).isLessThan(new BigNumber(-129)))) {
            _coll11.push([(_88_v1).minus(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.of(true)).length))).length)),false]);
          }
        }
        return _coll11;
      }()));
    };
    static fm36(p0, p1, p2, p3, globalState) {
      return function () {
        let _coll12 = new _dafny.Map();
        for (const _compr_12 of (_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))).Elements) {
          let _89_v0 = _compr_12;
          if ((_dafny.Set.fromElements(new _dafny.CodePoint('u'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('k'.codePointAt(0)))).contains(_89_v0)) {
            _coll12.push([_89_v0,(_dafny.Set.fromElements(true, true)).contains(!(true))]);
          }
        }
        return _coll12;
      }();
    };
    static fm37(p0, p1, p2, globalState) {
      return _module.D12.create_DC28(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(-806)));
    };
    static fm38(p0, globalState) {
      return ((_dafny.Map.Empty.slice().updateUnsafe(false,_dafny.Seq.of(false))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(false)))).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Seq.of(false)));
    };
    static fm39(p0, globalState) {
      return (_module.D16.create_DC39(new BigNumber((_dafny.Seq.UnicodeFromString("ok")).length), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.of(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(57)), function (_90_i0) {
  return new _dafny.CodePoint('y'.codePointAt(0));
})).length), new BigNumber((_dafny.MultiSet.fromElements(true)).cardinality()), new BigNumber(-821))).length),new BigNumber(181)))).dtor_cf75;
    };
    static fm40(globalState) {
      return _dafny.Seq.Concat(_dafny.Seq.of(_dafny.Seq.of(new BigNumber(-31), new BigNumber(889))), _dafny.Seq.of(_dafny.Seq.of(new BigNumber((_dafny.Seq.of(!(true))).length)), _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(677)),false)).length))));
    };
    static m7(p0, globalState) {
      let _91_v0;
      _91_v0 = _dafny.Seq.of(p0);
      let _92_v1;
      let _nw0 = new _module.C3();
      _nw0.__ctor(new BigNumber(823), _91_v0, (_dafny.ZERO).minus(p0));
      _92_v1 = _nw0;
      let _93_v2;
      _93_v2 = false;
      let _94_v3;
      _94_v3 = _dafny.Map.Empty.slice().updateUnsafe(_93_v2,(_92_v1).f9);
      let _95_v4;
      _95_v4 = _dafny.Map.Empty.slice().updateUnsafe(false,_94_v3);
      let _96_v5;
      _96_v5 = _dafny.Seq.of(_93_v2, _93_v2);
      let _97_v6;
      _97_v6 = _dafny.Seq.of(!((_96_v5)[_module.__default.safeIndex(p0, new BigNumber((_96_v5).length))]));
      let _98_v7;
      _98_v7 = _module.D13.create_DC33(_93_v2, new BigNumber((_95_v4).length), new BigNumber((_96_v5).length), _93_v2, _92_v1);
      let _99_v8;
      let _nw1 = Array((new BigNumber(10)).toNumber());
      _nw1[(_dafny.ZERO).toNumber()] = _92_v1;
      _nw1[(_dafny.ONE).toNumber()] = _92_v1;
      _nw1[(new BigNumber(2)).toNumber()] = ((_93_v2) ? (_92_v1) : (_92_v1));
      _nw1[(new BigNumber(3)).toNumber()] = (_98_v7).dtor_cf65;
      _nw1[(new BigNumber(4)).toNumber()] = _92_v1;
      _nw1[(new BigNumber(5)).toNumber()] = _92_v1;
      _nw1[(new BigNumber(6)).toNumber()] = _92_v1;
      _nw1[(new BigNumber(7)).toNumber()] = _92_v1;
      _nw1[(new BigNumber(8)).toNumber()] = _92_v1;
      _nw1[(new BigNumber(9)).toNumber()] = _92_v1;
      _99_v8 = _nw1;
      let _index0 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_99_v8).length));
      (_99_v8)[_index0] = _92_v1;
      let _index1 = _module.__default.safeIndex(new BigNumber(551), new BigNumber((_99_v8).length));
      (_99_v8)[_index1] = _92_v1;
      _97_v6 = _97_v6;
      let _100_v9;
      let _nw2 = Array((new BigNumber(26)).toNumber());
      _100_v9 = _nw2;
      let _101_v10;
      let _nw3 = new _module.C0();
      _nw3.__ctor(!(!(_93_v2)));
      _101_v10 = _nw3;
      let _index2 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_100_v9).length));
      (_100_v9)[_index2] = _101_v10;
      let _index3 = _module.__default.safeIndex(new BigNumber(796), new BigNumber((_100_v9).length));
      let _nw4 = new _module.C0();
      _nw4.__ctor((_101_v10).f13);
      (_100_v9)[_index3] = _nw4;
      let _102_v11;
      let _nw5 = new _module.C0();
      _nw5.__ctor((_101_v10).f13);
      _102_v11 = _nw5;
      let _103_v12;
      let _init0 = ((_104_v10) => function (_105_i0) {
        return (_104_v10).f13;
      })(_101_v10);
      let _nw6 = Array((new BigNumber(5)).toNumber());
      for (let _i0_0 = 0; _i0_0 < new BigNumber(_nw6.length); _i0_0++) {
        _nw6[_i0_0] = _init0(new BigNumber(_i0_0));
      }
      _103_v12 = _nw6;
      let _index4 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_103_v12).length));
      (_103_v12)[_index4] = _93_v2;
      let _106_v13;
      _106_v13 = _module.D5.create_DC12(p0);
      let _pat_let_tv0 = _92_v1;
      let _index5 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_103_v12).length));
      let _rhs0 = (_102_v11).f13;
      let _rhs1 = !(_93_v2) || (_93_v2);
      let _rhs2 = function (_pat_let0_0) {
        return function (_107_dt__update__tmp_h0) {
          return function (_pat_let1_0) {
            return function (_108_dt__update_hcf20_h0) {
              return _module.D5.create_DC12(_108_dt__update_hcf20_h0);
            }(_pat_let1_0);
          }((_pat_let_tv0).f9);
        }(_pat_let0_0);
      }(_106_v13);
      let _lhs0 = _103_v12;
      let _lhs1 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_103_v12).length));
      let _lhs2 = globalState;
      _lhs0[_lhs1] = _rhs0;
      _lhs2.f5 = _rhs1;
      _106_v13 = _rhs2;
      let _hi0 = (_92_v1).f9;
      for (let _109_i1 = p0; _109_i1.isLessThan(_hi0); _109_i1 = _109_i1.plus(_dafny.ONE)) {
        let _110_v14;
        _110_v14 = _dafny.Seq.UnicodeFromString("ode");
        let _111_v15;
        _111_v15 = _dafny.Map.Empty.slice().updateUnsafe((_92_v1).f9,(((_103_v12)[_module.__default.safeIndex(new BigNumber(177), new BigNumber((_103_v12).length))]) ? (_110_v14) : (_110_v14)));
        _111_v15 = (_111_v15).update(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(479)), ((_112_v1) => function (_113_i2) {
          return (_112_v1).f9;
        })(_92_v1))).length), (_92_v1).f9), _110_v14);
        let _index6 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_103_v12).length));
        (_103_v12)[_index6] = (_96_v5)[_module.__default.safeIndex(_109_i1, new BigNumber((_96_v5).length))];
        if (!(_93_v2)) {
          let _114_v16;
          _114_v16 = _dafny.Map.Empty.slice().updateUnsafe(_103_v12,(_92_v1).fm16(_dafny.MultiSet.fromElements((_101_v10).f13), globalState));
          _114_v16 = _114_v16;
          _93_v2 = !(_93_v2);
          let _index7 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_103_v12).length));
          (_103_v12)[_index7] = (_101_v10).f13;
          let _115_v17;
          let _nw7 = new _module.C1();
          _nw7.__ctor();
          _115_v17 = _nw7;
          let _116_v18;
          _116_v18 = _dafny.MultiSet.fromElements(new BigNumber(442), (_dafny.ZERO).minus((_dafny.ZERO).minus((_92_v1).f9)), (_dafny.ZERO).minus(p0));
          let _117_v19;
          _117_v19 = _dafny.Seq.of(_116_v18, _dafny.MultiSet.FromArray((_92_v1).f10));
          let _118_v20;
          _118_v20 = _dafny.Set.fromElements((_102_v11).f13, (_102_v11).f13);
          let _119_v21;
          _119_v21 = _dafny.Map.Empty.slice().updateUnsafe(((_117_v19)[_module.__default.safeIndex(new BigNumber((_118_v20).length), new BigNumber((_117_v19).length))]).IsProperSubsetOf(_116_v18),(_101_v10).f13);
          let _120_v22;
          _120_v22 = _119_v21;
          _119_v21 = (_119_v21).Merge(((_120_v22)));
        } else {
          let _121_v23;
          let _nw8 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _121_v23 = _nw8;
          let _index8 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_121_v23).length));
          (_121_v23)[_index8] = p0;
          let _index9 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_121_v23).length));
          (_121_v23)[_index9] = _109_i1;
          let _index10 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_103_v12).length));
          (_103_v12)[_index10] = (_101_v10).f13;
          let _122_v24;
          _122_v24 = new _dafny.CodePoint('w'.codePointAt(0));
          let _123_v25;
          _123_v25 = _module.D0.create_DC2((_dafny.ZERO).minus((_dafny.ZERO).minus(_109_i1)), _122_v24, (_dafny.ZERO).minus(new BigNumber((_dafny.Seq.update(_97_v6, _module.__default.safeIndex((_121_v23)[_module.__default.safeIndex(new BigNumber(532), new BigNumber((_121_v23).length))], new BigNumber((_97_v6).length)), (_101_v10).f13)).length)), _109_i1, (_92_v1).f9);
          _122_v24 = (_123_v25).dtor_cf6;
          let _124_v26;
          _124_v26 = _dafny.Set.fromElements(!((_101_v10).f13) || ((_101_v10).f13), _dafny.Seq.IsPrefixOf(_110_v14, _110_v14), (_102_v11).f13);
          let _index11 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_103_v12).length));
          let _rhs3 = _103_v12;
          let _rhs4 = (_102_v11).f13;
          let _rhs5 = (_124_v26).Difference((_124_v26).Intersect(_124_v26));
          let _lhs3 = _103_v12;
          let _lhs4 = _module.__default.safeIndex(new BigNumber(177), new BigNumber((_103_v12).length));
          _103_v12 = _rhs3;
          _lhs3[_lhs4] = _rhs4;
          _124_v26 = _rhs5;
          let _index12 = _module.__default.safeIndex(new BigNumber(532), new BigNumber((_121_v23).length));
          (_121_v23)[_index12] = (_92_v1).f9;
        }
        let _125_v27;
        _125_v27 = _dafny.MultiSet.fromElements(_103_v12, _103_v12, _103_v12, _103_v12, _103_v12);
        (globalState).f1 = !(((_125_v27).Union(_125_v27)).IsSubsetOf(_125_v27));
      }
      return;
    }
    static Main(__noArgsParameter) {
      let _126_v0;
      _126_v0 = _dafny.Seq.UnicodeFromString("oeg");
      let _127_globalState;
      let _nw9 = new _module.GlobalState();
      _nw9.__ctor(new BigNumber(-971), false, new BigNumber(560), true, _126_v0, false);
      _127_globalState = _nw9;
      let _128_v1;
      _128_v1 = false;
      let _129_v2;
      _129_v2 = _dafny.Set.fromElements(_128_v1);
      let _130_v3;
      _130_v3 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_129_v2);
      let _131_v5;
      _131_v5 = new BigNumber(371);
      let _132_v7;
      _132_v7 = new _dafny.CodePoint('g'.codePointAt(0));
      let _133_v8;
      _133_v8 = _dafny.MultiSet.fromElements(_131_v5);
      let _134_v9;
      _134_v9 = _module.D0.create_DC2(_131_v5, _132_v7, _131_v5, _131_v5, new BigNumber((_133_v8).cardinality()));
      let _pat_let_tv1 = _131_v5;
      let _135_v10;
      let _nw10 = Array((new BigNumber(17)).toNumber());
      _nw10[(_dafny.ZERO).toNumber()] = ((true) ? (_130_v3) : (_130_v3));
      _nw10[(_dafny.ONE).toNumber()] = _130_v3;
      _nw10[(new BigNumber(2)).toNumber()] = _130_v3;
      _nw10[(new BigNumber(3)).toNumber()] = _130_v3;
      _nw10[(new BigNumber(4)).toNumber()] = _130_v3;
      _nw10[(new BigNumber(5)).toNumber()] = (_130_v3).Merge(_130_v3);
      _nw10[(new BigNumber(6)).toNumber()] = _130_v3;
      _nw10[(new BigNumber(7)).toNumber()] = _130_v3;
      _nw10[(new BigNumber(8)).toNumber()] = _130_v3;
      _nw10[(new BigNumber(9)).toNumber()] = _130_v3;
      _nw10[(new BigNumber(10)).toNumber()] = function () {
        let _coll13 = new _dafny.Map();
        for (const _compr_13 of (_dafny.MultiSet.fromElements(_131_v5)).Elements) {
          let _136_v4 = _compr_13;
          if ((_dafny.MultiSet.fromElements(_131_v5)).contains(_136_v4)) {
            _coll13.push([(_136_v4).plus(_131_v5),_129_v2]);
          }
        }
        return _coll13;
      }();
      _nw10[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,_129_v2);
      _nw10[(new BigNumber(12)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,_dafny.Set.fromElements(_128_v1, _128_v1));
      _nw10[(new BigNumber(13)).toNumber()] = (_130_v3).Merge(_130_v3);
      _nw10[(new BigNumber(14)).toNumber()] = (_module.D0.create_DC0(function () {
  let _coll14 = new _dafny.Map();
  for (const _compr_14 of _dafny.IntegerRange(new BigNumber(181), new BigNumber(229))) {
    let _137_v6 = _compr_14;
    if (((new BigNumber(181)).isLessThanOrEqualTo(_137_v6)) && ((_137_v6).isLessThan(new BigNumber(229)))) {
      _coll14.push([_module.__default.safeDivisionInt(_137_v6, _131_v5),_129_v2]);
    }
  }
  return _coll14;
}())).dtor_cf0;
      _nw10[(new BigNumber(15)).toNumber()] = _130_v3;
      _nw10[(new BigNumber(16)).toNumber()] = _module.__default.fm0(_128_v1, _module.__default.fm1(_module.D0.create_DC0(_130_v3), _128_v1, _131_v5, _128_v1, _127_globalState), function (_pat_let2_0) {
        return function (_138_dt__update__tmp_h0) {
          return function (_pat_let3_0) {
            return function (_139_dt__update_hcf9_h0) {
              return _module.D0.create_DC2((_138_dt__update__tmp_h0).dtor_cf5, (_138_dt__update__tmp_h0).dtor_cf6, (_138_dt__update__tmp_h0).dtor_cf7, (_138_dt__update__tmp_h0).dtor_cf8, _139_dt__update_hcf9_h0);
            }(_pat_let3_0);
          }(_pat_let_tv1);
        }(_pat_let2_0);
      }(_134_v9), _127_globalState);
      _135_v10 = _nw10;
      let _index13 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_135_v10).length));
      (_135_v10)[_index13] = _130_v3;
      let _140_v11;
      _140_v11 = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,_130_v3);
      let _141_v12;
      _141_v12 = _module.D0.create_DC0(_130_v3);
      let _index14 = _module.__default.safeIndex(new BigNumber(207), new BigNumber((_135_v10).length));
      (_135_v10)[_index14] = (((_140_v11).contains(_131_v5)) ? ((_140_v11).get(_131_v5)) : (_module.__default.fm0(_module.__default.fm1(_141_v12, _128_v1, (_dafny.ZERO).minus(new BigNumber((_129_v2).length)), _128_v1, _127_globalState), _128_v1, _134_v9, _127_globalState)));
      if (_128_v1) {
        let _142_v13;
        _142_v13 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(_131_v5, _131_v5),_131_v5);
        _142_v13 = _142_v13;
        let _143_v14;
        let _nw11 = Array((new BigNumber(26)).toNumber()).fill(false);
        _143_v14 = _nw11;
        let _index15 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length));
        (_143_v14)[_index15] = _128_v1;
        let _index16 = _module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length));
        (_143_v14)[_index16] = _128_v1;
        let _144_v15;
        _144_v15 = _dafny.Map.Empty.slice().updateUnsafe((_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))],_131_v5);
        let _145_v16;
        _145_v16 = _dafny.Map.Empty.slice().updateUnsafe(_126_v0,new BigNumber((_144_v15).length));
        let _146_v17;
        _146_v17 = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,_dafny.Seq.UnicodeFromString("e"));
        let _147_v19;
        _147_v19 = _dafny.Seq.of(new BigNumber((function () {
          let _coll15 = new _dafny.Set();
          for (const _compr_15 of ((_146_v17).update(_131_v5, _dafny.Seq.UnicodeFromString("k"))).Keys.Elements) {
            let _148_v18 = _compr_15;
            if (((_146_v17).update(_131_v5, _dafny.Seq.UnicodeFromString("k"))).contains(_148_v18)) {
              _coll15.add(_module.__default.safeDivisionInt(_148_v18, new BigNumber((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(511),new BigNumber(38))).length))).length), new BigNumber((_dafny.Seq.UnicodeFromString("edejin")).length))).cardinality())));
            }
          }
          return _coll15;
        }()).length));
        _145_v16 = (_145_v16).update(_dafny.Seq.Concat(_126_v0, _126_v0), new BigNumber((_147_v19).length));
        _126_v0 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-612)), function (_149_i0) {
          return new _dafny.CodePoint('v'.codePointAt(0));
        }), _module.__default.fm2(_126_v0, _127_globalState));
        let _150_v20;
        _150_v20 = _module.D0.create_DC1(new BigNumber((_126_v0).length), (_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))], _131_v5, (_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))]);
        let _151_v21;
        _151_v21 = _dafny.Seq.of(_128_v1);
        let _152_v22;
        _152_v22 = _dafny.MultiSet.fromElements((_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))], _128_v1, false);
        let _nw12 = Array((new BigNumber(15)).toNumber());
        _nw12[(_dafny.ZERO).toNumber()] = !((_150_v20).dtor_cf2);
        _nw12[(_dafny.ONE).toNumber()] = (_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))];
        _nw12[(new BigNumber(2)).toNumber()] = (_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))];
        _nw12[(new BigNumber(3)).toNumber()] = _module.__default.fm1(_141_v12, _128_v1, _131_v5, (_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))], _127_globalState);
        _nw12[(new BigNumber(4)).toNumber()] = _128_v1;
        _nw12[(new BigNumber(5)).toNumber()] = (_131_v5).isLessThan(_131_v5);
        _nw12[(new BigNumber(6)).toNumber()] = !((_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))]) || ((_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))]);
        _nw12[(new BigNumber(7)).toNumber()] = _128_v1;
        _nw12[(new BigNumber(8)).toNumber()] = (_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))];
        _nw12[(new BigNumber(9)).toNumber()] = _dafny.Seq.IsPrefixOf(_151_v21, _151_v21);
        _nw12[(new BigNumber(10)).toNumber()] = (_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))];
        _nw12[(new BigNumber(11)).toNumber()] = _128_v1;
        _nw12[(new BigNumber(12)).toNumber()] = _128_v1;
        _nw12[(new BigNumber(13)).toNumber()] = (_143_v14)[_module.__default.safeIndex(new BigNumber(28), new BigNumber((_143_v14).length))];
        _nw12[(new BigNumber(14)).toNumber()] = (_152_v22).IsSubsetOf(_152_v22);
        _143_v14 = _nw12;
      } else {
        (_127_globalState).f5 = _module.__default.fm1(_module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_126_v0).length),_129_v2)), _128_v1, ((_128_v1) ? (_131_v5) : (_131_v5)), _128_v1, _127_globalState);
        if (_128_v1) {
          let _153_v23;
          let _nw13 = new _module.C0();
          _nw13.__ctor(((true) ? (_128_v1) : (_128_v1)));
          _153_v23 = _nw13;
          let _154_v24;
          _154_v24 = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,_131_v5);
          let _rhs6 = (_131_v5).minus((((_154_v24).contains(_131_v5)) ? ((_154_v24).get(_131_v5)) : (new BigNumber((_126_v0).length))));
          let _rhs7 = (_153_v23).f13;
          let _rhs8 = _131_v5;
          let _lhs5 = _127_globalState;
          let _lhs6 = _127_globalState;
          let _lhs7 = _127_globalState;
          _lhs5.f2 = _rhs6;
          _lhs6.f1 = _rhs7;
          _lhs7.f2 = _rhs8;
          let _pat_let_tv2 = _131_v5;
          let _pat_let_tv3 = _131_v5;
          let _pat_let_tv4 = _129_v2;
          let _pat_let_tv5 = _131_v5;
          let _pat_let_tv6 = _131_v5;
          let _pat_let_tv7 = _129_v2;
          let _155_v27;
          _155_v27 = _module.D13.create_DC32(_131_v5, true, _module.__default.fm1(function (_pat_let4_0) {
  return function (_156_dt__update__tmp_h1) {
    return function (_pat_let5_0) {
      return function (_160_dt__update_hcf0_h0) {
        return _module.D0.create_DC0(_160_dt__update_hcf0_h0);
      }(_pat_let5_0);
    }((function () {
      let _coll16 = new _dafny.Map();
      for (const _compr_16 of (function () {
        let _coll17 = new _dafny.Set();
        for (const _compr_17 of _dafny.IntegerRange(new BigNumber(598), new BigNumber(884))) {
          let _157_v26 = _compr_17;
          if (((new BigNumber(598)).isLessThanOrEqualTo(_157_v26)) && ((_157_v26).isLessThan(new BigNumber(884)))) {
            _coll17.add((_157_v26).plus(_pat_let_tv2));
          }
        }
        return _coll17;
      }()).Elements) {
        let _158_v25 = _compr_16;
        if ((function () {
          let _coll18 = new _dafny.Set();
          for (const _compr_18 of _dafny.IntegerRange(new BigNumber(598), new BigNumber(884))) {
            let _159_v26 = _compr_18;
            if (((new BigNumber(598)).isLessThanOrEqualTo(_159_v26)) && ((_159_v26).isLessThan(new BigNumber(884)))) {
              _coll18.add((_159_v26).plus(_pat_let_tv3));
            }
          }
          return _coll18;
        }()).contains(_158_v25)) {
          _coll16.push([_module.__default.safeModuloInt(_158_v25, _pat_let_tv5),_pat_let_tv4]);
        }
      }
      return _coll16;
    }()).update(_pat_let_tv6, _pat_let_tv7));
  }(_pat_let4_0);
}(_141_v12), _128_v1, _131_v5, _128_v1, _127_globalState), (_153_v23).f13, _128_v1);
          let _pat_let_tv8 = _128_v1;
          let _pat_let_tv9 = _153_v23;
          let _pat_let_tv10 = _153_v23;
          let _pat_let_tv11 = _128_v1;
          let _161_v28;
          _161_v28 = _dafny.Map.Empty.slice().updateUnsafe((function (_pat_let6_0) {
            return function (_166_dt__update__tmp_h3) {
              return function (_pat_let11_0) {
                return function (_167_dt__update_hcf58_h1) {
                  return _module.D13.create_DC32((_166_dt__update__tmp_h3).dtor_cf56, (_166_dt__update__tmp_h3).dtor_cf57, _167_dt__update_hcf58_h1, (_166_dt__update__tmp_h3).dtor_cf59, (_166_dt__update__tmp_h3).dtor_cf60);
                }(_pat_let11_0);
              }(_pat_let_tv11);
            }(_pat_let6_0);
          }(function (_pat_let7_0) {
            return function (_162_dt__update__tmp_h2) {
              return function (_pat_let8_0) {
                return function (_163_dt__update_hcf59_h0) {
                  return function (_pat_let9_0) {
                    return function (_164_dt__update_hcf58_h0) {
                      return function (_pat_let10_0) {
                        return function (_165_dt__update_hcf57_h0) {
                          return _module.D13.create_DC32((_162_dt__update__tmp_h2).dtor_cf56, _165_dt__update_hcf57_h0, _164_dt__update_hcf58_h0, _163_dt__update_hcf59_h0, (_162_dt__update__tmp_h2).dtor_cf60);
                        }(_pat_let10_0);
                      }((_pat_let_tv10).f13);
                    }(_pat_let9_0);
                  }((_pat_let_tv9).f13);
                }(_pat_let8_0);
              }(_pat_let_tv8);
            }(_pat_let7_0);
          }(_155_v27))).dtor_cf58,new BigNumber((_dafny.Set.fromElements((_153_v23).f13)).length));
          _161_v28 = (_161_v28).Merge((_161_v28).Merge(_dafny.Map.Empty.slice().updateUnsafe(!(_128_v1),_131_v5)));
          _126_v0 = _dafny.Seq.Concat(_module.__default.fm2(_dafny.Seq.UnicodeFromString("ioebdfr"), _127_globalState), _126_v0);
          _131_v5 = (_134_v9).dtor_cf5;
        } else {
          _module.__default.m7(_131_v5, _127_globalState);
          let _168_v29;
          let _nw14 = Array((new BigNumber(17)).toNumber()).fill(_dafny.MultiSet.Empty);
          _168_v29 = _nw14;
          _168_v29 = ((((!(_128_v1)) ? (!(_128_v1)) : (_128_v1))) ? (_168_v29) : (_168_v29));
          let _169_v30;
          _169_v30 = _dafny.Seq.of(_128_v1, _128_v1);
          let _170_v31;
          _170_v31 = _dafny.Seq.of(_131_v5, _131_v5);
          let _171_v32;
          let _nw15 = new _module.C4();
          _nw15.__ctor(new BigNumber((_170_v31).length));
          _171_v32 = _nw15;
          let _172_v33;
          _172_v33 = _dafny.Seq.of(_171_v32, _171_v32);
          let _173_v34;
          let _nw16 = Array((new BigNumber(24)).toNumber());
          _nw16[(_dafny.ZERO).toNumber()] = _131_v5;
          _nw16[(_dafny.ONE).toNumber()] = _131_v5;
          _nw16[(new BigNumber(2)).toNumber()] = _131_v5;
          _nw16[(new BigNumber(3)).toNumber()] = _131_v5;
          _nw16[(new BigNumber(4)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber((_169_v30).length), _131_v5);
          _nw16[(new BigNumber(5)).toNumber()] = _module.__default.safeModuloInt(new BigNumber(806), _131_v5);
          _nw16[(new BigNumber(6)).toNumber()] = _131_v5;
          _nw16[(new BigNumber(7)).toNumber()] = (_131_v5).multipliedBy(new BigNumber(442));
          _nw16[(new BigNumber(8)).toNumber()] = (new BigNumber(694)).multipliedBy(new BigNumber((_172_v33).length));
          _nw16[(new BigNumber(9)).toNumber()] = _131_v5;
          _nw16[(new BigNumber(10)).toNumber()] = _131_v5;
          _nw16[(new BigNumber(11)).toNumber()] = new BigNumber(660);
          _nw16[(new BigNumber(12)).toNumber()] = _131_v5;
          _nw16[(new BigNumber(13)).toNumber()] = (_131_v5).multipliedBy(new BigNumber((_169_v30).length));
          _nw16[(new BigNumber(14)).toNumber()] = _131_v5;
          _nw16[(new BigNumber(15)).toNumber()] = _131_v5;
          _nw16[(new BigNumber(16)).toNumber()] = _131_v5;
          _nw16[(new BigNumber(17)).toNumber()] = _module.__default.fm24(_131_v5, _127_globalState);
          _nw16[(new BigNumber(18)).toNumber()] = (_dafny.ZERO).minus(_131_v5);
          _nw16[(new BigNumber(19)).toNumber()] = (_131_v5).multipliedBy(_131_v5);
          _nw16[(new BigNumber(20)).toNumber()] = _131_v5;
          _nw16[(new BigNumber(21)).toNumber()] = new BigNumber((_dafny.Seq.of(_128_v1, _128_v1)).length);
          _nw16[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus(new BigNumber(-329));
          _nw16[(new BigNumber(23)).toNumber()] = _131_v5;
          _173_v34 = _nw16;
          let _index17 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_173_v34).length));
          (_173_v34)[_index17] = _module.__default.safeDivisionInt(_131_v5, _131_v5);
          let _174_v35;
          _174_v35 = _dafny.MultiSet.fromElements(_173_v34, _173_v34);
          let _175_v36;
          _175_v36 = _dafny.Map.Empty.slice().updateUnsafe(_128_v1,_131_v5);
          let _176_v37;
          _176_v37 = _dafny.MultiSet.fromElements(!(_131_v5).isEqualTo((((_174_v35).contains(_173_v34)) ? ((_174_v35).get(_173_v34)) : ((_dafny.ZERO).minus(new BigNumber((_175_v36).length))))), _module.__default.fm1(_141_v12, _128_v1, (_dafny.ZERO).minus(_131_v5), _128_v1, _127_globalState), false);
          let _index18 = _module.__default.safeIndex(new BigNumber(597), new BigNumber((_173_v34).length));
          (_173_v34)[_index18] = (((_176_v37).contains((_128_v1) || (_128_v1))) ? ((_176_v37).get((_128_v1) || (_128_v1))) : (new BigNumber(83)));
          (_127_globalState).f2 = (((_129_v2).IsSubsetOf(_129_v2)) ? (_131_v5) : (_131_v5));
          (_127_globalState).f0 = ((!((new BigNumber((_129_v2).length)).isEqualTo((_173_v34)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_173_v34).length))]))) ? ((_173_v34)[_module.__default.safeIndex(new BigNumber(597), new BigNumber((_173_v34).length))]) : (_module.__default.fm15(_131_v5, _128_v1, _127_globalState)));
        }
        let _177_v38;
        _177_v38 = _dafny.MultiSet.fromElements(_128_v1, _128_v1);
        (_127_globalState).f2 = (_131_v5).plus(new BigNumber((_177_v38).cardinality()));
        (_127_globalState).f0 = ((_131_v5).plus((_dafny.ZERO).minus(_131_v5))).multipliedBy(_131_v5);
        if (_128_v1) {
          _module.__default.m7(_131_v5, _127_globalState);
          let _178_v39;
          let _init1 = ((_179_v1) => function (_180_i1) {
            return _179_v1;
          })(_128_v1);
          let _nw17 = Array((new BigNumber(25)).toNumber());
          for (let _i0_1 = 0; _i0_1 < new BigNumber(_nw17.length); _i0_1++) {
            _nw17[_i0_1] = _init1(new BigNumber(_i0_1));
          }
          _178_v39 = _nw17;
          let _index19 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_178_v39).length));
          (_178_v39)[_index19] = !(_128_v1);
          let _181_v40;
          _181_v40 = _dafny.Seq.of(_131_v5, _131_v5, _131_v5);
          let _182_v41;
          _182_v41 = _dafny.Seq.of(!(true), _128_v1, _module.__default.fm1(_module.D0.create_DC0(_130_v3), _128_v1, (_181_v40)[_module.__default.safeIndex(new BigNumber(900), new BigNumber((_181_v40).length))], _128_v1, _127_globalState), _128_v1, _128_v1);
          let _183_v42;
          _183_v42 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_126_v0).length),_128_v1);
          let _184_v43;
          _184_v43 = _module.D14.create_DC36(_128_v1, (_dafny.ZERO).minus(_131_v5), new BigNumber((_dafny.Seq.UnicodeFromString("ncd")).length), _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(-972),(((_183_v42).contains(_131_v5)) ? ((_183_v42).get(_131_v5)) : (_128_v1))));
          let _185_v44;
          _185_v44 = _dafny.Map.Empty.slice().updateUnsafe(_132_v7,_184_v43);
          let _index20 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_178_v39).length));
          let _rhs9 = _131_v5;
          let _rhs10 = _128_v1;
          let _rhs11 = (_182_v41)[_module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_185_v44).length)), new BigNumber((_182_v41).length))];
          let _lhs8 = _127_globalState;
          let _lhs9 = _178_v39;
          let _lhs10 = _module.__default.safeIndex(new BigNumber(957), new BigNumber((_178_v39).length));
          let _lhs11 = _127_globalState;
          _lhs8.f2 = _rhs9;
          _lhs9[_lhs10] = _rhs10;
          _lhs11.f1 = _rhs11;
          (_127_globalState).f5 = !((_131_v5).isLessThan(_131_v5));
          _181_v40 = _181_v40;
          let _186_v45;
          _186_v45 = _dafny.Map.Empty.slice().updateUnsafe(_128_v1,_dafny.Seq.of(false, (_178_v39)[_module.__default.safeIndex(new BigNumber(957), new BigNumber((_178_v39).length))], (_178_v39)[_module.__default.safeIndex(new BigNumber(957), new BigNumber((_178_v39).length))]));
          (_127_globalState).f1 = !((_186_v45).equals((_module.__default.fm38((_178_v39)[_module.__default.safeIndex(new BigNumber(957), new BigNumber((_178_v39).length))], _127_globalState)).Merge(_186_v45)));
        } else {
          let _187_v46;
          _187_v46 = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,((true) ? (_128_v1) : (_128_v1)));
          let _188_v47;
          _188_v47 = _dafny.Map.Empty.slice().updateUnsafe(_132_v7,_131_v5);
          _187_v46 = (_187_v46).update(_131_v5, (_133_v8).IsSubsetOf(_dafny.MultiSet.FromArray(_module.__default.fm25(_131_v5, _188_v47, _132_v7, _127_globalState))));
          _126_v0 = _dafny.Seq.UnicodeFromString("dcti");
          let _189_v48;
          let _nw18 = new _module.C6();
          _nw18.__ctor(_131_v5);
          _189_v48 = _nw18;
          let _190_v49;
          let _nw19 = new _module.C1();
          _nw19.__ctor();
          _190_v49 = _nw19;
          let _191_v50;
          _191_v50 = _dafny.MultiSet.fromElements(_190_v49, _190_v49, _190_v49, _190_v49);
          let _192_v51;
          _192_v51 = _module.D11.create_DC26(_189_v48, _191_v50, _131_v5, _126_v0);
          let _193_v52;
          _193_v52 = _dafny.Set.fromElements(_189_v48);
          let _194_v53;
          let _nw20 = Array((new BigNumber(5)).toNumber());
          _nw20[(_dafny.ZERO).toNumber()] = _128_v1;
          _nw20[(_dafny.ONE).toNumber()] = !(_131_v5).isEqualTo(_131_v5);
          _nw20[(new BigNumber(2)).toNumber()] = (_dafny.Set.fromElements((_192_v51).dtor_cf45, _189_v48)).IsDisjointFrom(_193_v52);
          _nw20[(new BigNumber(3)).toNumber()] = _128_v1;
          _nw20[(new BigNumber(4)).toNumber()] = _128_v1;
          _194_v53 = _nw20;
          let _index21 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_194_v53).length));
          (_194_v53)[_index21] = _128_v1;
          let _195_v54;
          let _nw21 = new _module.C0();
          _nw21.__ctor(_128_v1);
          _195_v54 = _nw21;
          let _196_v55;
          _196_v55 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(516),_195_v54);
          let _197_v56;
          _197_v56 = _dafny.Set.fromElements(new BigNumber((_dafny.Seq.of(new BigNumber((_196_v55).length), _131_v5)).length));
          let _198_v57;
          _198_v57 = _dafny.Seq.of((_189_v48).f6, _131_v5);
          let _index22 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_194_v53).length));
          (_194_v53)[_index22] = (_module.__default.fm10(new BigNumber((_198_v57).length), (_195_v54).f13, _128_v1, _127_globalState)).IsSubsetOf(_197_v56);
          _128_v1 = _128_v1;
          let _199_v58;
          _199_v58 = _dafny.Map.Empty.slice().updateUnsafe((_195_v54).f13,new BigNumber((_198_v57).length));
          (_127_globalState).f2 = ((_128_v1) ? ((_dafny.ZERO).minus((new BigNumber((_199_v58).length)).multipliedBy(new BigNumber(826)))) : (_module.__default.safeModuloInt(_131_v5, new BigNumber(964))));
        }
      }
      let _200_v59;
      let _nw22 = new _module.C0();
      _nw22.__ctor(_128_v1);
      _200_v59 = _nw22;
      let _source8 = _200_v59;
      let _201___mcc_h0 = _source8;
      let _202_cf72 = _201___mcc_h0;
      let _203_v60;
      let _nw23 = Array((new BigNumber(6)).toNumber()).fill(false);
      _203_v60 = _nw23;
      let _index23 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_203_v60).length));
      (_203_v60)[_index23] = _128_v1;
      let _204_v61;
      _204_v61 = _dafny.Set.fromElements(_131_v5);
      let _205_v62;
      _205_v62 = _dafny.MultiSet.fromElements(_204_v61);
      let _index24 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((_203_v60).length));
      (_203_v60)[_index24] = (_205_v62).contains(_204_v61);
      let _206_v63;
      _206_v63 = _dafny.Seq.of(_131_v5);
      _133_v8 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(((_128_v1) ? (new BigNumber((_206_v63).length)) : (_module.__default.fm15(_131_v5, _128_v1, _127_globalState)))), (_131_v5).plus(new BigNumber((_204_v61).length)), _131_v5, _131_v5);
      (_127_globalState).f1 = !((_200_v59).f13);
      if ((((_200_v59).f13) ? ((_200_v59).f13) : (!((_202_cf72).f13) || (_128_v1)))) {
        let _207_v64;
        _207_v64 = _dafny.Seq.of((_202_cf72).f13);
        _207_v64 = _dafny.Seq.Concat(_207_v64, _dafny.Seq.of((_200_v59).f13));
        let _208_v65;
        let _nw24 = new _module.C2();
        _nw24.__ctor(_126_v0, (_202_cf72).f13);
        _208_v65 = _nw24;
        let _209_v66;
        _209_v66 = _dafny.Map.Empty.slice().updateUnsafe(_204_v61,_207_v64);
        (_208_v65).f12 = !(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(new BigNumber((_204_v61).length)),_module.__default.fm21((_208_v65).f11, _127_globalState))).equals(_209_v66);
        let _210_v67;
        let _init2 = ((_211_v5) => function (_212_i2) {
          return (_212_i2).multipliedBy(_211_v5);
        })(_131_v5);
        let _nw25 = Array((new BigNumber(16)).toNumber());
        for (let _i0_2 = 0; _i0_2 < new BigNumber(_nw25.length); _i0_2++) {
          _nw25[_i0_2] = _init2(new BigNumber(_i0_2));
        }
        _210_v67 = _nw25;
        let _init3 = ((_213_v1) => function (_214_i3) {
          return (_214_i3).multipliedBy((_dafny.ZERO).minus(new BigNumber((_dafny.MultiSet.fromElements(_213_v1)).cardinality())));
        })(_128_v1);
        let _nw26 = Array((new BigNumber(26)).toNumber());
        for (let _i0_3 = 0; _i0_3 < new BigNumber(_nw26.length); _i0_3++) {
          _nw26[_i0_3] = _init3(new BigNumber(_i0_3));
        }
        _210_v67 = _nw26;
        let _215_v68;
        _215_v68 = _dafny.Map.Empty.slice().updateUnsafe(_132_v7,_module.__default.fm1(_141_v12, _208_v65.f12, (_dafny.ZERO).minus(_131_v5), (_202_cf72).f13, _127_globalState));
        let _216_v69;
        _216_v69 = _dafny.Seq.of(_215_v68, _215_v68, (_215_v68).Merge(_215_v68));
        _216_v69 = _216_v69;
      } else {
        _module.__default.m7(_131_v5, _127_globalState);
        (_127_globalState).f5 = (_200_v59).f13;
        _126_v0 = _126_v0;
        (_127_globalState).f5 = true;
        let _217_v70;
        _217_v70 = _dafny.Seq.of((_200_v59).f13, (_200_v59).f13);
        let _218_v71;
        _218_v71 = _dafny.Map.Empty.slice().updateUnsafe(!(!((new BigNumber((_217_v70).length)).isEqualTo(_131_v5))),(((_202_cf72).f13) ? ((_200_v59).f13) : ((_202_cf72).f13)));
        _218_v71 = (_218_v71).update((((_218_v71).contains(true)) ? ((_218_v71).get(true)) : (_128_v1)), (_133_v8).IsProperSubsetOf(_133_v8));
      }
      let _219_v72;
      _219_v72 = _dafny.Map.Empty.slice().updateUnsafe(_128_v1,_131_v5);
      let _220_v73;
      _220_v73 = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,(_200_v59).f13);
      let _221_v74;
      _221_v74 = _dafny.Set.fromElements(_module.__default.safeModuloInt(_131_v5, new BigNumber((_dafny.Seq.update(_dafny.Seq.of(_219_v72, _219_v72), _module.__default.safeIndex(_131_v5, new BigNumber((_dafny.Seq.of(_219_v72, _219_v72)).length)), _dafny.Map.Empty.slice().updateUnsafe(_128_v1,new BigNumber((_220_v73).length)))).length)));
      _221_v74 = _221_v74;
      _128_v1 = false;
      (_127_globalState).f0 = _131_v5;
      let _rhs12 = _131_v5;
      let _rhs13 = _128_v1;
      let _lhs12 = _127_globalState;
      let _lhs13 = _127_globalState;
      _lhs12.f0 = _rhs12;
      _lhs13.f5 = _rhs13;
      let _222_v75;
      _222_v75 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_module.__default.fm39(_131_v5, _127_globalState)).length),_131_v5);
      let _223_v76;
      _223_v76 = _module.D16.create_DC39((_dafny.ZERO).minus(new BigNumber((_126_v0).length)), _222_v75);
      let _pat_let_tv12 = _222_v75;
      let _source9 = function (_pat_let12_0) {
        return function (_224_dt__update__tmp_h4) {
          return function (_pat_let13_0) {
            return function (_225_dt__update_hcf75_h0) {
              return _module.D16.create_DC39((_224_dt__update__tmp_h4).dtor_cf74, _225_dt__update_hcf75_h0);
            }(_pat_let13_0);
          }(_pat_let_tv12);
        }(_pat_let12_0);
      }(_223_v76);
      if (_source9.is_DC39) {
        let _226___mcc_h1 = (_source9).cf74;
        let _227___mcc_h2 = (_source9).cf75;
        let _228_cf75 = _227___mcc_h2;
        let _229_cf74 = _226___mcc_h1;
        let _230_v77;
        _230_v77 = _module.D1.create_DC5();
        let _source10 = _230_v77;
        if (_source10.is_DC5) {
          let _231_v78;
          _231_v78 = _dafny.MultiSet.fromElements((_200_v59).f13);
          let _232_v79;
          _232_v79 = _dafny.Seq.of(_229_cf74);
          let _233_v80;
          let _nw27 = new _module.C3();
          _nw27.__ctor(new BigNumber((_231_v78).cardinality()), _232_v79, _229_cf74);
          _233_v80 = _nw27;
          let _234_v81;
          _234_v81 = _module.D13.create_DC33((_200_v59).f13, new BigNumber((_129_v2).length), new BigNumber(11), (_200_v59).f13, _233_v80);
          let _235_v82;
          let _nw28 = Array((new BigNumber(6)).toNumber());
          _nw28[(_dafny.ZERO).toNumber()] = _229_cf74;
          _nw28[(_dafny.ONE).toNumber()] = new BigNumber((_222_v75).length);
          _nw28[(new BigNumber(2)).toNumber()] = _module.__default.safeDivisionInt(_131_v5, _229_cf74);
          _nw28[(new BigNumber(3)).toNumber()] = _131_v5;
          _nw28[(new BigNumber(4)).toNumber()] = ((!(_module.__default.fm1(_141_v12, (_234_v81).dtor_cf64, _131_v5, !((_200_v59).f13), _127_globalState))) ? (new BigNumber((_221_v74).length)) : (_131_v5));
          _nw28[(new BigNumber(5)).toNumber()] = _131_v5;
          _235_v82 = _nw28;
          let _index25 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_235_v82).length));
          (_235_v82)[_index25] = (_dafny.ZERO).minus((_dafny.ZERO).minus((_233_v80).f9));
          let _index26 = _module.__default.safeIndex(new BigNumber(649), new BigNumber((_235_v82).length));
          (_235_v82)[_index26] = (_233_v80).f9;
          let _236_v83;
          let _nw29 = new _module.C5();
          _nw29.__ctor((_200_v59).f13, _132_v7);
          _236_v83 = _nw29;
          _236_v83 = _236_v83;
          let _237_v84;
          let _nw30 = Array((new BigNumber(12)).toNumber()).fill(_dafny.Seq.of());
          _237_v84 = _nw30;
          let _238_v85;
          let _init4 = ((_239_v83) => function (_240_i4) {
            return (_239_v83).f7;
          })(_236_v83);
          let _nw31 = Array((new BigNumber(26)).toNumber());
          for (let _i0_4 = 0; _i0_4 < new BigNumber(_nw31.length); _i0_4++) {
            _nw31[_i0_4] = _init4(new BigNumber(_i0_4));
          }
          _238_v85 = _nw31;
          let _241_v86;
          _241_v86 = _dafny.Seq.of(_238_v85, _238_v85);
          let _index27 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_237_v84).length));
          (_237_v84)[_index27] = _241_v86;
          let _index28 = _module.__default.safeIndex(new BigNumber(360), new BigNumber((_237_v84).length));
          (_237_v84)[_index28] = _241_v86;
          (_127_globalState).f2 = new BigNumber((_231_v78).cardinality());
        } else {
          let _242___mcc_h4 = (_source10).cf11;
          let _243_cf11 = _242___mcc_h4;
          let _244_v87;
          let _nw32 = Array((new BigNumber(26)).toNumber()).fill(false);
          _244_v87 = _nw32;
          let _index29 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_244_v87).length));
          (_244_v87)[_index29] = (_200_v59).f13;
          let _index30 = _module.__default.safeIndex(new BigNumber(671), new BigNumber((_244_v87).length));
          (_244_v87)[_index30] = _128_v1;
          (_127_globalState).f1 = (_200_v59).f13;
          let _245_v88;
          let _init5 = ((_246_v5) => function (_247_i5) {
            return (_247_i5).plus(_246_v5);
          })(_131_v5);
          let _nw33 = Array((new BigNumber(2)).toNumber());
          for (let _i0_5 = 0; _i0_5 < new BigNumber(_nw33.length); _i0_5++) {
            _nw33[_i0_5] = _init5(new BigNumber(_i0_5));
          }
          _245_v88 = _nw33;
          _245_v88 = _245_v88;
          _module.__default.m7(_229_cf74, _127_globalState);
        }
        if ((_200_v59).f13) {
          let _248_v89;
          let _nw34 = Array((new BigNumber(10)).toNumber()).fill(false);
          _248_v89 = _nw34;
          _248_v89 = _248_v89;
          let _249_v90;
          _249_v90 = _dafny.Map.Empty.slice().updateUnsafe(_248_v89,_128_v1);
          (_127_globalState).f5 = (((_249_v90).contains(_248_v89)) ? (false) : ((_200_v59).f13));
          let _250_v91;
          _250_v91 = _module.D0.create_DC1(_131_v5, _128_v1, _229_cf74, (_200_v59).f13);
          let _251_v92;
          _251_v92 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_141_v12, !((_200_v59).f13), new BigNumber((_126_v0).length), _module.__default.fm1(_141_v12, (_200_v59).f13, new BigNumber(-926), (_250_v91).dtor_cf4, _127_globalState), _127_globalState),_126_v0);
          _251_v92 = (_251_v92).update((_200_v59).f13, _126_v0);
          let _252_v93;
          let _nw35 = Array((new BigNumber(19)).toNumber());
          _nw35[(_dafny.ZERO).toNumber()] = _248_v89;
          _nw35[(_dafny.ONE).toNumber()] = _248_v89;
          _nw35[(new BigNumber(2)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(3)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(4)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(5)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(6)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(7)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(8)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(9)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(10)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(11)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(12)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(13)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(14)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(15)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(16)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(17)).toNumber()] = _248_v89;
          _nw35[(new BigNumber(18)).toNumber()] = _248_v89;
          _252_v93 = _nw35;
          let _index31 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_252_v93).length));
          (_252_v93)[_index31] = _248_v89;
          let _253_v94;
          _253_v94 = _248_v89;
          let _index32 = _module.__default.safeIndex(new BigNumber(944), new BigNumber((_252_v93).length));
          (_252_v93)[_index32] = ((_253_v94));
          let _254_v95;
          let _nw36 = new _module.C0();
          _nw36.__ctor((_229_cf74).isEqualTo(_229_cf74));
          _254_v95 = _nw36;
        } else {
          let _255_v96;
          let _nw37 = new _module.C4();
          _nw37.__ctor(_229_cf74);
          _255_v96 = _nw37;
          let _256_v97;
          _256_v97 = _dafny.Seq.of(_255_v96, _255_v96);
          let _257_v98;
          _257_v98 = _dafny.Map.Empty.slice().updateUnsafe((_200_v59).f13,_dafny.MultiSet.fromElements(_256_v97));
          let _258_v99;
          _258_v99 = _dafny.Seq.of(_256_v97);
          _257_v98 = (_257_v98).update((_200_v59).f13, _dafny.MultiSet.FromArray(_258_v99));
          let _259_v100;
          _259_v100 = _dafny.Seq.of(_229_cf74, (_255_v96).f6);
          _259_v100 = _259_v100;
          (_127_globalState).f5 = ((_200_v59).f13) === ((_200_v59).f13);
          let _260_v101;
          _260_v101 = _module.D8.create_DC21((_200_v59).f13, _module.__default.fm1(_141_v12, (_200_v59).f13, (_255_v96).f6, false, _127_globalState), _131_v5);
          (_127_globalState).f5 = !((_260_v101).dtor_cf37);
          let _261_v102;
          _261_v102 = _dafny.Seq.of((_200_v59).f13, _128_v1, false);
          (_127_globalState).f0 = (new BigNumber((_261_v102).length)).minus(_131_v5);
        }
        let _262_v103;
        let _nw38 = Array((new BigNumber(20)).toNumber()).fill(_dafny.Map.Empty);
        _262_v103 = _nw38;
        let _263_v104;
        _263_v104 = _module.D14.create_DC35(_262_v103);
        let _source11 = _263_v104;
        if (_source11.is_DC36) {
          let _264___mcc_h5 = (_source11).cf68;
          let _265___mcc_h6 = (_source11).cf69;
          let _266___mcc_h7 = (_source11).cf70;
          let _267___mcc_h8 = (_source11).cf71;
          let _268_cf71 = _267___mcc_h8;
          let _269_cf70 = _266___mcc_h7;
          let _270_cf69 = _265___mcc_h6;
          let _271_cf68 = _264___mcc_h5;
          (_127_globalState).f1 = _module.__default.fm1(_141_v12, !(_271_cf68), _269_cf70, (_200_v59).f13, _127_globalState);
          _128_v1 = !((_269_cf70).multipliedBy(new BigNumber((_126_v0).length))).isEqualTo((new BigNumber((_126_v0).length)).minus(_229_cf74));
          let _272_v105;
          _272_v105 = _dafny.Seq.of(_module.D1.create_DC5(), _230_v77, ((_271_cf68) ? (_module.D1.create_DC5()) : (_230_v77)), _module.D1.create_DC5(), _230_v77);
          _272_v105 = ((_128_v1) ? (_272_v105) : (_272_v105));
          _module.__default.m7(_131_v5, _127_globalState);
        } else {
          let _273___mcc_h9 = (_source11).cf67;
          let _274_cf67 = _273___mcc_h9;
          let _275_v106;
          let _nw39 = Array((new BigNumber(3)).toNumber());
          _nw39[(_dafny.ZERO).toNumber()] = (_200_v59).f13;
          _nw39[(_dafny.ONE).toNumber()] = (_200_v59).f13;
          _nw39[(new BigNumber(2)).toNumber()] = (_200_v59).f13;
          _275_v106 = _nw39;
          let _index33 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_275_v106).length));
          (_275_v106)[_index33] = (_200_v59).f13;
          let _276_v107;
          _276_v107 = _dafny.Set.fromElements(_dafny.Seq.update(_126_v0, _module.__default.safeIndex(_131_v5, new BigNumber((_126_v0).length)), _132_v7));
          let _277_v108;
          _277_v108 = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,_module.__default.fm8(_132_v7, _128_v1, new BigNumber(-237), _module.__default.fm24(_131_v5, _127_globalState), _127_globalState));
          let _index34 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_275_v106).length));
          let _rhs14 = !((new BigNumber(((_129_v2).Difference(_129_v2)).length)).isLessThan(new BigNumber(285)));
          let _rhs15 = (((_277_v108).contains(_131_v5)) ? ((_277_v108).get(_131_v5)) : (_132_v7));
          let _rhs16 = (_200_v59).f13;
          let _rhs17 = _dafny.Set.fromElements(_126_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(569)), ((_278_v7) => function (_279_i6) {
            return _278_v7;
          })(_132_v7)), _dafny.Seq.Concat(_126_v0, _126_v0));
          let _lhs14 = _127_globalState;
          let _lhs15 = _275_v106;
          let _lhs16 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_275_v106).length));
          _lhs14.f1 = _rhs14;
          _132_v7 = _rhs15;
          _lhs15[_lhs16] = _rhs16;
          _276_v107 = _rhs17;
          let _index35 = _module.__default.safeIndex(new BigNumber(482), new BigNumber((_275_v106).length));
          (_275_v106)[_index35] = _module.__default.fm1(_141_v12, (new BigNumber(695)).isLessThanOrEqualTo(_131_v5), ((((_133_v8).contains(_229_cf74)) ? ((_133_v8).get(_229_cf74)) : (_131_v5))).multipliedBy(new BigNumber(-56)), (_200_v59).f13, _127_globalState);
          let _280_v109;
          let _nw40 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _280_v109 = _nw40;
          let _rhs18 = _280_v109;
          let _rhs19 = _module.__default.fm24(_131_v5, _127_globalState);
          let _rhs20 = (new BigNumber((_dafny.Seq.Concat(_126_v0, _126_v0)).length)).isLessThan(_131_v5);
          _280_v109 = _rhs18;
          _131_v5 = _rhs19;
          _128_v1 = _rhs20;
          let _281_v110;
          _281_v110 = _module.D8.create_DC21((_275_v106)[_module.__default.safeIndex(new BigNumber(482), new BigNumber((_275_v106).length))], (_200_v59).f13, _229_cf74);
          (_127_globalState).f2 = (_281_v110).dtor_cf38;
        }
        let _282_v111;
        _282_v111 = _dafny.Map.Empty.slice().updateUnsafe(_230_v77,(_200_v59).f13);
        let _rhs21 = _module.__default.safeModuloInt(_131_v5, _229_cf74);
        let _rhs22 = _dafny.Map.Empty.slice().updateUnsafe(_230_v77,false);
        _131_v5 = _rhs21;
        _282_v111 = _rhs22;
      } else {
        let _283___mcc_h3 = (_source9).cf73;
        let _284_cf73 = _283___mcc_h3;
        _126_v0 = _126_v0;
        let _285_v112;
        let _init6 = ((_286_v59) => function (_287_i7) {
          return (_286_v59).f13;
        })(_200_v59);
        let _nw41 = Array((new BigNumber(22)).toNumber());
        for (let _i0_6 = 0; _i0_6 < new BigNumber(_nw41.length); _i0_6++) {
          _nw41[_i0_6] = _init6(new BigNumber(_i0_6));
        }
        _285_v112 = _nw41;
        let _288_v113;
        _288_v113 = _dafny.Seq.of(_dafny.Seq.of(new BigNumber(-204), _131_v5, _131_v5, _131_v5, _131_v5));
        let _index36 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_285_v112).length));
        (_285_v112)[_index36] = (_200_v59).f13;
        let _289_v114;
        _289_v114 = _dafny.Seq.of(_129_v2, _129_v2, _129_v2);
        let _index37 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_285_v112).length));
        let _rhs23 = (_dafny.ZERO).minus(new BigNumber((((_289_v114)[_module.__default.safeIndex(_131_v5, new BigNumber((_289_v114).length))]).Difference(_dafny.Set.fromElements((_200_v59).f13))).length));
        let _rhs24 = ((_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_126_v0).length))).Merge(_219_v72)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_200_v59).f13,new BigNumber((_126_v0).length)));
        let _rhs25 = _285_v112;
        let _rhs26 = _module.__default.fm40(_127_globalState);
        let _rhs27 = _128_v1;
        let _lhs17 = _127_globalState;
        let _lhs18 = _285_v112;
        let _lhs19 = _module.__default.safeIndex(new BigNumber(421), new BigNumber((_285_v112).length));
        _lhs17.f2 = _rhs23;
        _219_v72 = _rhs24;
        _285_v112 = _rhs25;
        _288_v113 = _rhs26;
        _lhs18[_lhs19] = _rhs27;
        let _290_v115;
        let _nw42 = new _module.C1();
        _nw42.__ctor();
        _290_v115 = _nw42;
        let _291_v116;
        _291_v116 = _dafny.Map.Empty.slice().updateUnsafe(_290_v115,(_285_v112)[_module.__default.safeIndex(new BigNumber(421), new BigNumber((_285_v112).length))]);
        let _rhs28 = _291_v116;
        let _rhs29 = _129_v2;
        let _rhs30 = _200_v59;
        let _rhs31 = (_dafny.ZERO).minus((_131_v5).plus((_dafny.ZERO).minus(new BigNumber((_219_v72).length))));
        let _rhs32 = (((_285_v112)[_module.__default.safeIndex(new BigNumber(421), new BigNumber((_285_v112).length))]) ? ((function () {
          let _coll19 = new _dafny.Set();
          for (const _compr_19 of (_222_v75).Keys.Elements) {
            let _292_v117 = _compr_19;
            if ((_222_v75).contains(_292_v117)) {
              _coll19.add((_292_v117).plus(new BigNumber(734)));
            }
          }
          return _coll19;
        }()).Difference(_221_v74)) : (function () {
          let _coll20 = new _dafny.Set();
          for (const _compr_20 of _dafny.IntegerRange(new BigNumber(-261), new BigNumber(593))) {
            let _293_v118 = _compr_20;
            if (((new BigNumber(-261)).isLessThanOrEqualTo(_293_v118)) && ((_293_v118).isLessThan(new BigNumber(593)))) {
              _coll20.add((_293_v118).multipliedBy(new BigNumber((_126_v0).length)));
            }
          }
          return _coll20;
        }()));
        let _lhs20 = _127_globalState;
        _291_v116 = _rhs28;
        _129_v2 = _rhs29;
        _200_v59 = _rhs30;
        _lhs20.f0 = _rhs31;
        _221_v74 = _rhs32;
        let _294_v119;
        let _nw43 = new _module.C2();
        _nw43.__ctor(_126_v0, true);
        _294_v119 = _nw43;
        _294_v119 = _294_v119;
      }
      let _295_v120;
      _295_v120 = _dafny.Set.fromElements(_200_v59);
      let _296_v121;
      _296_v121 = _200_v59;
      _295_v120 = _dafny.Set.fromElements(_200_v59, (_200_v59), _200_v59, _200_v59);
      let _297_v122;
      _297_v122 = _dafny.MultiSet.fromElements(_128_v1, _128_v1, _128_v1, _128_v1, false);
      let _298_v123;
      _298_v123 = _dafny.Map.Empty.slice().updateUnsafe(!(new BigNumber((_297_v122).cardinality())).isEqualTo(new BigNumber((_129_v2).length)),_129_v2);
      _298_v123 = (_298_v123).Merge((_298_v123).Merge(_298_v123));
      let _hi1 = (_131_v5).minus(_131_v5);
      for (let _299_i8 = new BigNumber((_126_v0).length); _299_i8.isLessThan(_hi1); _299_i8 = _299_i8.plus(_dafny.ONE)) {
        let _rhs33 = _132_v7;
        let _rhs34 = (_dafny.ZERO).minus(_299_i8);
        let _lhs21 = _127_globalState;
        _132_v7 = _rhs33;
        _lhs21.f2 = _rhs34;
        let _300_v124;
        let _init7 = ((_301_v59) => function (_302_i9) {
          return (_301_v59).f13;
        })(_200_v59);
        let _nw44 = Array((new BigNumber(29)).toNumber());
        for (let _i0_7 = 0; _i0_7 < new BigNumber(_nw44.length); _i0_7++) {
          _nw44[_i0_7] = _init7(new BigNumber(_i0_7));
        }
        _300_v124 = _nw44;
        let _index38 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_300_v124).length));
        (_300_v124)[_index38] = _128_v1;
        let _303_v125;
        _303_v125 = _dafny.Seq.of(_299_i8, _131_v5);
        let _304_v126;
        _304_v126 = _dafny.Map.Empty.slice().updateUnsafe((_200_v59).f13,_module.__default.fm8(_132_v7, (_200_v59).f13, _131_v5, _131_v5, _127_globalState));
        let _index39 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_300_v124).length));
        (_300_v124)[_index39] = _module.__default.fm1(_141_v12, true, (_303_v125)[_module.__default.safeIndex(_131_v5, new BigNumber((_303_v125).length))], (new BigNumber((_304_v126).length)).isLessThan(_299_i8), _127_globalState);
        (_127_globalState).f1 = true;
        let _index40 = _module.__default.safeIndex(new BigNumber(752), new BigNumber((_300_v124).length));
        (_300_v124)[_index40] = _128_v1;
      }
      _module.__default.m7(_131_v5, _127_globalState);
      let _305_v127;
      _305_v127 = _dafny.Seq.of(_200_v59);
      let _306_v128;
      let _nw45 = Array((new BigNumber(13)).toNumber());
      _nw45[(_dafny.ZERO).toNumber()] = (new BigNumber(773)).plus(_131_v5);
      _nw45[(_dafny.ONE).toNumber()] = _module.__default.fm15((_dafny.ZERO).minus(_131_v5), _128_v1, _127_globalState);
      _nw45[(new BigNumber(2)).toNumber()] = (_131_v5).multipliedBy((_134_v9).dtor_cf9);
      _nw45[(new BigNumber(3)).toNumber()] = _131_v5;
      _nw45[(new BigNumber(4)).toNumber()] = _131_v5;
      _nw45[(new BigNumber(5)).toNumber()] = (_131_v5).plus(new BigNumber((_305_v127).length));
      _nw45[(new BigNumber(6)).toNumber()] = (_223_v76).dtor_cf74;
      _nw45[(new BigNumber(7)).toNumber()] = _131_v5;
      _nw45[(new BigNumber(8)).toNumber()] = _131_v5;
      _nw45[(new BigNumber(9)).toNumber()] = _131_v5;
      _nw45[(new BigNumber(10)).toNumber()] = _131_v5;
      _nw45[(new BigNumber(11)).toNumber()] = (_131_v5).minus(_131_v5);
      _nw45[(new BigNumber(12)).toNumber()] = (_dafny.ZERO).minus((((_200_v59).f13) ? (_131_v5) : (new BigNumber((_dafny.Seq.of(!((_200_v59).f13), _128_v1, _128_v1)).length))));
      _306_v128 = _nw45;
      for (const _guard_loop_0 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_306_v128).length))) {
        let _307_i10 = _guard_loop_0;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_307_i10)) && ((_307_i10).isLessThan(new BigNumber((_306_v128).length))))) {
          (_306_v128)[(_307_i10)] = _module.__default.safeModuloInt(_307_i10, _131_v5);
        }
      }
      let _hi2 = (_131_v5).plus(_131_v5);
      for (let _308_i11 = _131_v5; _308_i11.isLessThan(_hi2); _308_i11 = _308_i11.plus(_dafny.ONE)) {
        let _309_v129;
        let _nw46 = new _module.C1();
        _nw46.__ctor();
        _309_v129 = _nw46;
        let _310_v130;
        let _nw47 = new _module.C6();
        _nw47.__ctor(_131_v5);
        _310_v130 = _nw47;
        let _311_v131;
        _311_v131 = _dafny.Map.Empty.slice().updateUnsafe(_309_v129,_310_v130);
        _311_v131 = _311_v131;
        (_127_globalState).f5 = (_200_v59).f13;
        let _312_v132;
        let _init8 = ((_313_v59) => function (_314_i12) {
          return (_313_v59).f13;
        })(_200_v59);
        let _nw48 = Array((new BigNumber(25)).toNumber());
        for (let _i0_8 = 0; _i0_8 < new BigNumber(_nw48.length); _i0_8++) {
          _nw48[_i0_8] = _init8(new BigNumber(_i0_8));
        }
        _312_v132 = _nw48;
        let _index41 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_312_v132).length));
        (_312_v132)[_index41] = _dafny.Seq.IsPrefixOf(_126_v0, _126_v0);
        let _index42 = _module.__default.safeIndex(new BigNumber(699), new BigNumber((_312_v132).length));
        (_312_v132)[_index42] = (_200_v59).f13;
        _126_v0 = _126_v0;
      }
      if ((_200_v59).f13) {
        if ((_200_v59).f13) {
          let _315_v133;
          _315_v133 = _dafny.Map.Empty.slice().updateUnsafe((_200_v59).fm23(_131_v5, _131_v5, _132_v7, (_200_v59).f13, _127_globalState),_dafny.Seq.update(_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-126)), function (_316_i13) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("vky")).length);
          }), _dafny.Seq.of(_131_v5)), _module.__default.safeIndex(_131_v5, new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-126)), function (_317_i13) {
            return new BigNumber((_dafny.Seq.UnicodeFromString("vky")).length);
          }), _dafny.Seq.of(_131_v5))).length)), new BigNumber(224)));
          let _318_v134;
          _318_v134 = _dafny.Seq.of(_131_v5);
          _315_v133 = (_315_v133).update(_131_v5, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(860)), ((_319_v0, _320_v59) => function (_321_i14) {
            return new BigNumber(((_module.D12.create_DC29(_319_v0, (_320_v59).f13)).dtor_cf52).length);
          })(_126_v0, _200_v59)), _318_v134));
          let _322_v135;
          let _nw49 = Array((new BigNumber(19)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _322_v135 = _nw49;
          _322_v135 = _322_v135;
          (_127_globalState).f0 = _module.__default.safeDivisionInt(_131_v5, _131_v5);
          (_127_globalState).f0 = ((_318_v134)[_module.__default.safeIndex((_200_v59).fm23(_131_v5, new BigNumber((_222_v75).length), _132_v7, _128_v1, _127_globalState), new BigNumber((_318_v134).length))]).minus(new BigNumber(((_129_v2).Intersect(_dafny.Set.fromElements(!((_200_v59).f13)))).length));
          let _323_v136;
          _323_v136 = _dafny.Seq.of(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-864)), ((_324_v7) => function (_325_i15) {
            return _324_v7;
          })(_132_v7)), _dafny.Seq.UnicodeFromString("c"), _126_v0, _126_v0);
          _module.__default.m7(_module.__default.safeModuloInt(_131_v5, new BigNumber((_323_v136).length)), _127_globalState);
        } else {
          let _326_v137;
          let _init9 = ((_327_v1) => function (_328_i16) {
            return _327_v1;
          })(_128_v1);
          let _nw50 = Array((new BigNumber(20)).toNumber());
          for (let _i0_9 = 0; _i0_9 < new BigNumber(_nw50.length); _i0_9++) {
            _nw50[_i0_9] = _init9(new BigNumber(_i0_9));
          }
          _326_v137 = _nw50;
          let _index43 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_326_v137).length));
          (_326_v137)[_index43] = !(_128_v1);
          let _index44 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_306_v128).length));
          (_306_v128)[_index44] = _131_v5;
          let _329_v138;
          _329_v138 = _dafny.Seq.of(_131_v5);
          let _330_v139;
          _330_v139 = _dafny.Map.Empty.slice().updateUnsafe(_329_v138,_306_v128);
          let _331_v140;
          _331_v140 = _module.D8.create_DC20(_330_v139);
          let _332_v141;
          _332_v141 = _dafny.Set.fromElements(_331_v140, _331_v140);
          let _index45 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_326_v137).length));
          let _index46 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_306_v128).length));
          let _rhs35 = (_dafny.ZERO).minus((_329_v138)[_module.__default.safeIndex((_131_v5).multipliedBy(_131_v5), new BigNumber((_329_v138).length))]);
          let _rhs36 = (_133_v8).IsSubsetOf(_133_v8);
          let _rhs37 = !(((_332_v141).Intersect(_332_v141)).IsSubsetOf(_332_v141));
          let _rhs38 = (_200_v59).f13;
          let _rhs39 = _module.__default.safeModuloInt(new BigNumber((_126_v0).length), _module.__default.safeModuloInt(_131_v5, _131_v5));
          let _lhs22 = _127_globalState;
          let _lhs23 = _326_v137;
          let _lhs24 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_326_v137).length));
          let _lhs25 = _127_globalState;
          let _lhs26 = _127_globalState;
          let _lhs27 = _306_v128;
          let _lhs28 = _module.__default.safeIndex(new BigNumber(785), new BigNumber((_306_v128).length));
          _lhs22.f0 = _rhs35;
          _lhs23[_lhs24] = _rhs36;
          _lhs25.f1 = _rhs37;
          _lhs26.f5 = _rhs38;
          _lhs27[_lhs28] = _rhs39;
          let _333_v142;
          _333_v142 = _module.D5.create_DC13(_128_v1, _128_v1, _126_v0, new BigNumber(-978));
          let _334_v143;
          _334_v143 = _module.D3.create_DC7((_333_v142).dtor_cf23);
          let _index47 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_326_v137).length));
          let _rhs40 = _334_v143;
          let _rhs41 = (new BigNumber((_126_v0).length)).plus(((true) ? ((_306_v128)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_306_v128).length))]) : (_131_v5)));
          let _rhs42 = (_200_v59).f13;
          let _lhs29 = _127_globalState;
          let _lhs30 = _326_v137;
          let _lhs31 = _module.__default.safeIndex(new BigNumber(331), new BigNumber((_326_v137).length));
          _334_v143 = _rhs40;
          _lhs29.f2 = _rhs41;
          _lhs30[_lhs31] = _rhs42;
          let _335_v144;
          _335_v144 = _module.D8.create_DC21((_200_v59).f13, false, (new BigNumber((_126_v0).length)).multipliedBy(new BigNumber((_dafny.MultiSet.fromElements(_128_v1, _128_v1, (_200_v59).f13, _128_v1)).cardinality())));
          _335_v144 = _335_v144;
          let _336_v145;
          let _nw51 = new _module.C4();
          _nw51.__ctor((_306_v128)[_module.__default.safeIndex(new BigNumber(785), new BigNumber((_306_v128).length))]);
          _336_v145 = _nw51;
          let _337_v146;
          let _init10 = ((_338_v5) => function (_339_i17) {
            return _module.__default.safeDivisionInt(_339_i17, _338_v5);
          })(_131_v5);
          let _nw52 = Array((new BigNumber(15)).toNumber());
          for (let _i0_10 = 0; _i0_10 < new BigNumber(_nw52.length); _i0_10++) {
            _nw52[_i0_10] = _init10(new BigNumber(_i0_10));
          }
          _337_v146 = _nw52;
          let _340_v147;
          _340_v147 = _dafny.Map.Empty.slice().updateUnsafe(_337_v146,(_dafny.ZERO).minus(_131_v5));
          (_127_globalState).f5 = (_340_v147).contains(_337_v146);
        }
        (_127_globalState).f5 = !((_200_v59).f13);
        let _341_v148;
        let _nw53 = Array((new BigNumber(15)).toNumber()).fill([]);
        _341_v148 = _nw53;
        let _342_v149;
        let _nw54 = Array((new BigNumber(2)).toNumber()).fill([]);
        _342_v149 = _nw54;
        let _index48 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_341_v148).length));
        (_341_v148)[_index48] = _342_v149;
        let _343_v150;
        let _nw55 = new _module.C1();
        _nw55.__ctor();
        _343_v150 = _nw55;
        let _344_v151;
        let _nw56 = Array((new BigNumber(22)).toNumber());
        _nw56[(_dafny.ZERO).toNumber()] = (_200_v59).f13;
        _nw56[(_dafny.ONE).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(2)).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(3)).toNumber()] = _128_v1;
        _nw56[(new BigNumber(4)).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(5)).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(6)).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(7)).toNumber()] = true;
        _nw56[(new BigNumber(8)).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(9)).toNumber()] = !(_128_v1);
        _nw56[(new BigNumber(10)).toNumber()] = _128_v1;
        _nw56[(new BigNumber(11)).toNumber()] = !((_200_v59).f13);
        _nw56[(new BigNumber(12)).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(13)).toNumber()] = false;
        _nw56[(new BigNumber(14)).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(15)).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(16)).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(17)).toNumber()] = (_200_v59).f13;
        _nw56[(new BigNumber(18)).toNumber()] = !((_200_v59).f13);
        _nw56[(new BigNumber(19)).toNumber()] = !((_200_v59).f13);
        _nw56[(new BigNumber(20)).toNumber()] = !(true);
        _nw56[(new BigNumber(21)).toNumber()] = _128_v1;
        _344_v151 = _nw56;
        let _345_v152;
        _345_v152 = _module.D11.create_DC25(_344_v151, _343_v150, new BigNumber((_221_v74).length));
        let _346_v153;
        let _nw57 = Array((new BigNumber(27)).toNumber());
        _nw57[(_dafny.ZERO).toNumber()] = _343_v150;
        _nw57[(_dafny.ONE).toNumber()] = _343_v150;
        _nw57[(new BigNumber(2)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(3)).toNumber()] = (((_200_v59).f13) ? (_343_v150) : (_343_v150));
        _nw57[(new BigNumber(4)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(5)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(6)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(7)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(8)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(9)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(10)).toNumber()] = (_345_v152).dtor_cf43;
        _nw57[(new BigNumber(11)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(12)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(13)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(14)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(15)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(16)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(17)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(18)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(19)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(20)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(21)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(22)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(23)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(24)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(25)).toNumber()] = _343_v150;
        _nw57[(new BigNumber(26)).toNumber()] = _343_v150;
        _346_v153 = _nw57;
        let _index49 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_346_v153).length));
        (_346_v153)[_index49] = _343_v150;
        let _index50 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_341_v148).length));
        let _index51 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_346_v153).length));
        let _rhs43 = _342_v149;
        let _rhs44 = _343_v150;
        let _lhs32 = _341_v148;
        let _lhs33 = _module.__default.safeIndex(new BigNumber(896), new BigNumber((_341_v148).length));
        let _lhs34 = _346_v153;
        let _lhs35 = _module.__default.safeIndex(new BigNumber(892), new BigNumber((_346_v153).length));
        _lhs32[_lhs33] = _rhs43;
        _lhs34[_lhs35] = _rhs44;
        _221_v74 = function () {
          let _coll21 = new _dafny.Set();
          for (const _compr_21 of _dafny.IntegerRange(new BigNumber(306), new BigNumber(-263))) {
            let _347_v154 = _compr_21;
            if (((new BigNumber(306)).isLessThanOrEqualTo(_347_v154)) && ((_347_v154).isLessThan(new BigNumber(-263)))) {
              _coll21.add((_347_v154).minus(_131_v5));
            }
          }
          return _coll21;
        }();
        if (!(_128_v1) || ((((_200_v59).f13) ? (_module.__default.fm1(_141_v12, (_200_v59).f13, _131_v5, (_200_v59).f13, _127_globalState)) : ((_200_v59).f13)))) {
          let _348_v155;
          let _nw58 = new _module.C5();
          _nw58.__ctor(!(!(_128_v1)), _132_v7);
          _348_v155 = _nw58;
          let _349_v156;
          let _nw59 = new _module.C1();
          _nw59.__ctor();
          _349_v156 = _nw59;
          let _350_v157;
          _350_v157 = _module.D18.create_DC43(_349_v156);
          let _351_v158;
          _351_v158 = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,_349_v156);
          let _352_v159;
          _352_v159 = _dafny.Seq.of(_131_v5, _131_v5);
          let _353_v160;
          let _nw60 = new _module.C3();
          _nw60.__ctor(_131_v5, _352_v159, new BigNumber(507));
          _353_v160 = _nw60;
          let _354_v161;
          _354_v161 = _dafny.Map.Empty.slice().updateUnsafe(_353_v160,new _dafny.CodePoint('k'.codePointAt(0)));
          let _355_v162;
          _355_v162 = _dafny.MultiSet.fromElements(_354_v161, _354_v161);
          let _356_v163;
          let _nw61 = Array((new BigNumber(24)).toNumber());
          _nw61[(_dafny.ZERO).toNumber()] = _349_v156;
          _nw61[(_dafny.ONE).toNumber()] = _349_v156;
          _nw61[(new BigNumber(2)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(3)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(4)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(5)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(6)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(7)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(8)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(9)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(10)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(11)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(12)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(13)).toNumber()] = (_350_v157).dtor_cf80;
          _nw61[(new BigNumber(14)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(15)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(16)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(17)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(18)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(19)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(20)).toNumber()] = (((_351_v158).contains(new BigNumber((_355_v162).cardinality()))) ? ((_351_v158).get(new BigNumber((_355_v162).cardinality()))) : (_349_v156));
          _nw61[(new BigNumber(21)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(22)).toNumber()] = _349_v156;
          _nw61[(new BigNumber(23)).toNumber()] = _349_v156;
          _356_v163 = _nw61;
          let _index52 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_356_v163).length));
          (_356_v163)[_index52] = _349_v156;
          let _index53 = _module.__default.safeIndex(new BigNumber(305), new BigNumber((_356_v163).length));
          (_356_v163)[_index53] = _349_v156;
          (_127_globalState).f5 = (_200_v59).f13;
          _132_v7 = _132_v7;
          let _357_v164;
          let _nw62 = Array((new BigNumber(17)).toNumber()).fill(new _dafny.CodePoint('D'.codePointAt(0)));
          _357_v164 = _nw62;
          let _358_v165;
          _358_v165 = _module.D19.create_DC45(_357_v164);
          _357_v164 = (_358_v165).dtor_cf86;
        } else {
          let _359_v166;
          _359_v166 = _dafny.Map.Empty.slice().updateUnsafe((_346_v153)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_346_v153).length))],_128_v1);
          let _360_v167;
          _360_v167 = _module.D13.create_DC32(_131_v5, false, (_200_v59).f13, (_200_v59).f13, (_200_v59).f13);
          _359_v166 = (_359_v166).update((_346_v153)[_module.__default.safeIndex(new BigNumber(892), new BigNumber((_346_v153).length))], (_360_v167).dtor_cf57);
          let _361_v168;
          _361_v168 = _dafny.Seq.of(_128_v1);
          (_127_globalState).f1 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.Concat(_361_v168, _361_v168), _dafny.Seq.Concat(_361_v168, _361_v168));
          let _362_v169;
          let _init11 = ((_363_v7) => function (_364_i18) {
            return _363_v7;
          })(_132_v7);
          let _nw63 = Array((new BigNumber(29)).toNumber());
          for (let _i0_11 = 0; _i0_11 < new BigNumber(_nw63.length); _i0_11++) {
            _nw63[_i0_11] = _init11(new BigNumber(_i0_11));
          }
          _362_v169 = _nw63;
          let _index54 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_362_v169).length));
          (_362_v169)[_index54] = _132_v7;
          let _365_v170;
          _365_v170 = _dafny.Map.Empty.slice().updateUnsafe((_200_v59).f13,(_200_v59).f13);
          let _index55 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_344_v151).length));
          (_344_v151)[_index55] = (((_365_v170).contains(_128_v1)) ? ((_365_v170).get(_128_v1)) : (_128_v1));
          let _index56 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_362_v169).length));
          let _index57 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_344_v151).length));
          let _rhs45 = (_module.__default.fm19((_200_v59).f13, _132_v7, _dafny.Map.Empty.slice().updateUnsafe((((_219_v72).contains(_128_v1)) ? ((_219_v72).get(_128_v1)) : (_131_v5)),true), _128_v1, _127_globalState)).dtor_cf6;
          let _rhs46 = !((_200_v59).f13) || (!(_dafny.Seq.IsProperPrefixOf(_126_v0, _dafny.Seq.update(_126_v0, _module.__default.safeIndex(_131_v5, new BigNumber((_126_v0).length)), _132_v7))));
          let _lhs36 = _362_v169;
          let _lhs37 = _module.__default.safeIndex(new BigNumber(492), new BigNumber((_362_v169).length));
          let _lhs38 = _344_v151;
          let _lhs39 = _module.__default.safeIndex(new BigNumber(904), new BigNumber((_344_v151).length));
          _lhs36[_lhs37] = _rhs45;
          _lhs38[_lhs39] = _rhs46;
          _module.__default.m7(_131_v5, _127_globalState);
          (_127_globalState).f0 = (_dafny.ZERO).minus((_module.__default.safeDivisionInt(_131_v5, new BigNumber((_126_v0).length))).multipliedBy((_200_v59).fm23(_131_v5, _131_v5, new _dafny.CodePoint('n'.codePointAt(0)), (_module.D8.create_DC21(_128_v1, (_200_v59).f13, _131_v5)).dtor_cf37, _127_globalState)));
        }
      } else {
        let _366_v171;
        let _nw64 = new _module.C4();
        _nw64.__ctor(_131_v5);
        _366_v171 = _nw64;
        let _367_v172;
        _367_v172 = _dafny.Seq.of(new BigNumber(771), _131_v5, _131_v5);
        let _368_v173;
        _368_v173 = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,_126_v0);
        let _369_v174;
        _369_v174 = _dafny.MultiSet.fromElements(_367_v172, _dafny.Seq.of(new BigNumber(((((_368_v173).contains(_131_v5)) ? ((_368_v173).get(_131_v5)) : (_126_v0))).length), new BigNumber((_module.__default.fm21(_126_v0, _127_globalState)).length)), _367_v172);
        let _370_v175;
        _370_v175 = _dafny.Seq.of(_306_v128, _306_v128);
        let _371_v176;
        _371_v176 = _dafny.Map.Empty.slice().updateUnsafe((((_200_v59).f13) ? (_dafny.MultiSet.FromArray(_module.__default.fm40(_127_globalState))) : (_369_v174)),_module.__default.fm15(_module.__default.fm15((_dafny.ZERO).minus(new BigNumber((_370_v175).length)), (_200_v59).f13, _127_globalState), (_200_v59).f13, _127_globalState));
        _371_v176 = (_371_v176).update(_369_v174, _131_v5);
        let _372_v177;
        let _nw65 = Array((new BigNumber(21)).toNumber()).fill(_module.D14.Default());
        _372_v177 = _nw65;
        _372_v177 = _372_v177;
        let _373_v178;
        _373_v178 = _dafny.Map.Empty.slice().updateUnsafe(_131_v5,_132_v7);
        let _374_v179;
        let _nw66 = new _module.C3();
        _nw66.__ctor(_module.__default.safeModuloInt(_131_v5, new BigNumber((_dafny.Seq.of((_200_v59).f13, (_200_v59).f13, _128_v1, _128_v1)).length)), _dafny.Seq.update(_dafny.Seq.of(new BigNumber((_373_v178).length)), _module.__default.safeIndex(new BigNumber((_module.__default.fm2(_126_v0, _127_globalState)).length), new BigNumber((_dafny.Seq.of(new BigNumber((_373_v178).length))).length)), _131_v5), _131_v5);
        _374_v179 = _nw66;
        _128_v1 = !(_128_v1) || (false);
      }
      (_127_globalState).f0 = new BigNumber((_dafny.Seq.Concat(_dafny.Seq.Concat(_126_v0, _dafny.Seq.Create(_module.__default.abs(new BigNumber(788)), function (_375_i19) {
        return new _dafny.CodePoint('l'.codePointAt(0));
      })), _dafny.Seq.UnicodeFromString("oobul"))).length);
      process.stdout.write((_126_v0).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_127_globalState.f0));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_127_globalState.f1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_127_globalState.f2));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_127_globalState).f3));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(((_127_globalState).f4).toVerbatimString(false));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_127_globalState.f5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_128_v1));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_129_v2).equals(_dafny.Set.fromElements(false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_130_v3).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_131_v5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(_132_v7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_133_v8).equals(_dafny.MultiSet.fromElements(new BigNumber(872), new BigNumber(372), new BigNumber(371), new BigNumber(371)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v9).dtor_cf5));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v9).dtor_cf6));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v9).dtor_cf7));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v9).dtor_cf8));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_134_v9).dtor_cf9));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[_dafny.ZERO]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[_dafny.ONE]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(2)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(3)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(4)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(5)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(6)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(7)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(8)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(9)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(10)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(742),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(11)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(371),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(12)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(371),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(13)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(14)]).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ZERO,_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(15)]).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_135_v10)[new BigNumber(16)]).equals(_dafny.Map.Empty.slice())));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_140_v11).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(371),_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false))))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_141_v12).dtor_cf0).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(277),_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_200_v59).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_219_v72).equals(_dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber(371)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_220_v73).equals(_dafny.Map.Empty.slice().updateUnsafe(new BigNumber(371),false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_221_v74).equals(_dafny.Set.fromElements(_dafny.ONE))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_222_v75).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(371)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_223_v76).dtor_cf74));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_223_v76).dtor_cf75).equals(_dafny.Map.Empty.slice().updateUnsafe(_dafny.ONE,new BigNumber(371)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_295_v120).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(((_296_v121)).f13));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_297_v122).equals(_dafny.MultiSet.fromElements(true, true, true, true, false))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_298_v123).equals(_dafny.Map.Empty.slice().updateUnsafe(true,_dafny.Set.fromElements(false)))));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString(new BigNumber((_305_v127).length)));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[_dafny.ZERO]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[_dafny.ONE]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(2)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(3)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(4)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(5)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(6)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(7)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(8)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(9)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(10)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(11)]));
      process.stdout.write((_dafny.Seq.UnicodeFromString("\n")).toVerbatimString(false));
      process.stdout.write(_dafny.toString((_306_v128)[new BigNumber(12)]));
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
    static create_DC2(cf5, cf6, cf7, cf8, cf9) {
      let $dt = new D0(1);
      $dt.cf5 = cf5;
      $dt.cf6 = cf6;
      $dt.cf7 = cf7;
      $dt.cf8 = cf8;
      $dt.cf9 = cf9;
      return $dt;
    }
    static create_DC0(cf0) {
      let $dt = new D0(2);
      $dt.cf0 = cf0;
      return $dt;
    }
    static create_DC3(cf10) {
      let $dt = new D0(3);
      $dt.cf10 = cf10;
      return $dt;
    }
    get is_DC1() { return this.$tag === 0; }
    get is_DC2() { return this.$tag === 1; }
    get is_DC0() { return this.$tag === 2; }
    get is_DC3() { return this.$tag === 3; }
    get dtor_cf1() { return this.cf1; }
    get dtor_cf2() { return this.cf2; }
    get dtor_cf3() { return this.cf3; }
    get dtor_cf4() { return this.cf4; }
    get dtor_cf5() { return this.cf5; }
    get dtor_cf6() { return this.cf6; }
    get dtor_cf7() { return this.cf7; }
    get dtor_cf8() { return this.cf8; }
    get dtor_cf9() { return this.cf9; }
    get dtor_cf0() { return this.cf0; }
    get dtor_cf10() { return this.cf10; }
    toString() {
      if (this.$tag === 0) {
        return "D0.DC1" + "(" + _dafny.toString(this.cf1) + ", " + _dafny.toString(this.cf2) + ", " + _dafny.toString(this.cf3) + ", " + _dafny.toString(this.cf4) + ")";
      } else if (this.$tag === 1) {
        return "D0.DC2" + "(" + _dafny.toString(this.cf5) + ", " + _dafny.toString(this.cf6) + ", " + _dafny.toString(this.cf7) + ", " + _dafny.toString(this.cf8) + ", " + _dafny.toString(this.cf9) + ")";
      } else if (this.$tag === 2) {
        return "D0.DC0" + "(" + _dafny.toString(this.cf0) + ")";
      } else if (this.$tag === 3) {
        return "D0.DC3" + "(" + _dafny.toString(this.cf10) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf1, other.cf1) && this.cf2 === other.cf2 && _dafny.areEqual(this.cf3, other.cf3) && this.cf4 === other.cf4;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf5, other.cf5) && _dafny.areEqual(this.cf6, other.cf6) && _dafny.areEqual(this.cf7, other.cf7) && _dafny.areEqual(this.cf8, other.cf8) && _dafny.areEqual(this.cf9, other.cf9);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf0, other.cf0);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf10, other.cf10);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D0.create_DC1(_dafny.ZERO, false, _dafny.ZERO, false);
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
    static create_DC5() {
      let $dt = new D1(0);
      return $dt;
    }
    static create_DC4(cf11) {
      let $dt = new D1(1);
      $dt.cf11 = cf11;
      return $dt;
    }
    get is_DC5() { return this.$tag === 0; }
    get is_DC4() { return this.$tag === 1; }
    get dtor_cf11() { return this.cf11; }
    toString() {
      if (this.$tag === 0) {
        return "D1.DC5";
      } else if (this.$tag === 1) {
        return "D1.DC4" + "(" + _dafny.toString(this.cf11) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf11, other.cf11);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D1.create_DC5();
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
    static create_DC6(cf12) {
      let $dt = new D2(0);
      $dt.cf12 = cf12;
      return $dt;
    }
    get is_DC6() { return this.$tag === 0; }
    get dtor_cf12() { return this.cf12; }
    toString() {
      if (this.$tag === 0) {
        return "D2.DC6" + "(" + _dafny.toString(this.cf12) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf12 === other.cf12;
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
    static create_DC8(cf14, cf15, cf16) {
      let $dt = new D3(0);
      $dt.cf14 = cf14;
      $dt.cf15 = cf15;
      $dt.cf16 = cf16;
      return $dt;
    }
    static create_DC7(cf13) {
      let $dt = new D3(1);
      $dt.cf13 = cf13;
      return $dt;
    }
    static create_DC9(cf17) {
      let $dt = new D3(2);
      $dt.cf17 = cf17;
      return $dt;
    }
    get is_DC8() { return this.$tag === 0; }
    get is_DC7() { return this.$tag === 1; }
    get is_DC9() { return this.$tag === 2; }
    get dtor_cf14() { return this.cf14; }
    get dtor_cf15() { return this.cf15; }
    get dtor_cf16() { return this.cf16; }
    get dtor_cf13() { return this.cf13; }
    get dtor_cf17() { return this.cf17; }
    toString() {
      if (this.$tag === 0) {
        return "D3.DC8" + "(" + _dafny.toString(this.cf14) + ", " + _dafny.toString(this.cf15) + ", " + _dafny.toString(this.cf16) + ")";
      } else if (this.$tag === 1) {
        return "D3.DC7" + "(" + this.cf13.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D3.DC9" + "(" + _dafny.toString(this.cf17) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf14, other.cf14) && this.cf15 === other.cf15 && _dafny.areEqual(this.cf16, other.cf16);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf13, other.cf13);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf17, other.cf17);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D3.create_DC8(_module.D1.Default(), [], _dafny.ZERO);
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
    static create_DC10(cf18) {
      let $dt = new D4(0);
      $dt.cf18 = cf18;
      return $dt;
    }
    get is_DC10() { return this.$tag === 0; }
    get dtor_cf18() { return this.cf18; }
    toString() {
      if (this.$tag === 0) {
        return "D4.DC10" + "(" + _dafny.toString(this.cf18) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf18, other.cf18);
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
          return D4.Default();
        }
      };
    }
  }

  $module.D5 = class D5 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC12(cf20) {
      let $dt = new D5(0);
      $dt.cf20 = cf20;
      return $dt;
    }
    static create_DC13(cf21, cf22, cf23, cf24) {
      let $dt = new D5(1);
      $dt.cf21 = cf21;
      $dt.cf22 = cf22;
      $dt.cf23 = cf23;
      $dt.cf24 = cf24;
      return $dt;
    }
    static create_DC11(cf19) {
      let $dt = new D5(2);
      $dt.cf19 = cf19;
      return $dt;
    }
    static create_DC14(cf25) {
      let $dt = new D5(3);
      $dt.cf25 = cf25;
      return $dt;
    }
    get is_DC12() { return this.$tag === 0; }
    get is_DC13() { return this.$tag === 1; }
    get is_DC11() { return this.$tag === 2; }
    get is_DC14() { return this.$tag === 3; }
    get dtor_cf20() { return this.cf20; }
    get dtor_cf21() { return this.cf21; }
    get dtor_cf22() { return this.cf22; }
    get dtor_cf23() { return this.cf23; }
    get dtor_cf24() { return this.cf24; }
    get dtor_cf19() { return this.cf19; }
    get dtor_cf25() { return this.cf25; }
    toString() {
      if (this.$tag === 0) {
        return "D5.DC12" + "(" + _dafny.toString(this.cf20) + ")";
      } else if (this.$tag === 1) {
        return "D5.DC13" + "(" + _dafny.toString(this.cf21) + ", " + _dafny.toString(this.cf22) + ", " + this.cf23.toVerbatimString(true) + ", " + _dafny.toString(this.cf24) + ")";
      } else if (this.$tag === 2) {
        return "D5.DC11" + "(" + _dafny.toString(this.cf19) + ")";
      } else if (this.$tag === 3) {
        return "D5.DC14" + "(" + _dafny.toString(this.cf25) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf20, other.cf20);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf21 === other.cf21 && this.cf22 === other.cf22 && _dafny.areEqual(this.cf23, other.cf23) && _dafny.areEqual(this.cf24, other.cf24);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf19, other.cf19);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf25, other.cf25);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D5.create_DC12(_dafny.ZERO);
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
    static create_DC16(cf27) {
      let $dt = new D6(0);
      $dt.cf27 = cf27;
      return $dt;
    }
    static create_DC15(cf26) {
      let $dt = new D6(1);
      $dt.cf26 = cf26;
      return $dt;
    }
    get is_DC16() { return this.$tag === 0; }
    get is_DC15() { return this.$tag === 1; }
    get dtor_cf27() { return this.cf27; }
    get dtor_cf26() { return this.cf26; }
    toString() {
      if (this.$tag === 0) {
        return "D6.DC16" + "(" + _dafny.toString(this.cf27) + ")";
      } else if (this.$tag === 1) {
        return "D6.DC15" + "(" + _dafny.toString(this.cf26) + ")";
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
        return other.$tag === 1 && _dafny.areEqual(this.cf26, other.cf26);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D6.create_DC16(_dafny.ZERO);
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
    static create_DC18(cf29, cf30, cf31, cf32) {
      let $dt = new D7(0);
      $dt.cf29 = cf29;
      $dt.cf30 = cf30;
      $dt.cf31 = cf31;
      $dt.cf32 = cf32;
      return $dt;
    }
    static create_DC19(cf33, cf34) {
      let $dt = new D7(1);
      $dt.cf33 = cf33;
      $dt.cf34 = cf34;
      return $dt;
    }
    static create_DC17(cf28) {
      let $dt = new D7(2);
      $dt.cf28 = cf28;
      return $dt;
    }
    get is_DC18() { return this.$tag === 0; }
    get is_DC19() { return this.$tag === 1; }
    get is_DC17() { return this.$tag === 2; }
    get dtor_cf29() { return this.cf29; }
    get dtor_cf30() { return this.cf30; }
    get dtor_cf31() { return this.cf31; }
    get dtor_cf32() { return this.cf32; }
    get dtor_cf33() { return this.cf33; }
    get dtor_cf34() { return this.cf34; }
    get dtor_cf28() { return this.cf28; }
    toString() {
      if (this.$tag === 0) {
        return "D7.DC18" + "(" + _dafny.toString(this.cf29) + ", " + _dafny.toString(this.cf30) + ", " + _dafny.toString(this.cf31) + ", " + _dafny.toString(this.cf32) + ")";
      } else if (this.$tag === 1) {
        return "D7.DC19" + "(" + _dafny.toString(this.cf33) + ", " + _dafny.toString(this.cf34) + ")";
      } else if (this.$tag === 2) {
        return "D7.DC17" + "(" + _dafny.toString(this.cf28) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf29 === other.cf29 && this.cf30 === other.cf30 && _dafny.areEqual(this.cf31, other.cf31) && this.cf32 === other.cf32;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf33, other.cf33) && _dafny.areEqual(this.cf34, other.cf34);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf28, other.cf28);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D7.create_DC18(false, false, _dafny.ZERO, false);
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
    static create_DC21(cf36, cf37, cf38) {
      let $dt = new D8(0);
      $dt.cf36 = cf36;
      $dt.cf37 = cf37;
      $dt.cf38 = cf38;
      return $dt;
    }
    static create_DC20(cf35) {
      let $dt = new D8(1);
      $dt.cf35 = cf35;
      return $dt;
    }
    get is_DC21() { return this.$tag === 0; }
    get is_DC20() { return this.$tag === 1; }
    get dtor_cf36() { return this.cf36; }
    get dtor_cf37() { return this.cf37; }
    get dtor_cf38() { return this.cf38; }
    get dtor_cf35() { return this.cf35; }
    toString() {
      if (this.$tag === 0) {
        return "D8.DC21" + "(" + _dafny.toString(this.cf36) + ", " + _dafny.toString(this.cf37) + ", " + _dafny.toString(this.cf38) + ")";
      } else if (this.$tag === 1) {
        return "D8.DC20" + "(" + _dafny.toString(this.cf35) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf36 === other.cf36 && this.cf37 === other.cf37 && _dafny.areEqual(this.cf38, other.cf38);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf35, other.cf35);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D8.create_DC21(false, false, _dafny.ZERO);
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
    static create_DC22(cf39) {
      let $dt = new D9(0);
      $dt.cf39 = cf39;
      return $dt;
    }
    get is_DC22() { return this.$tag === 0; }
    get dtor_cf39() { return this.cf39; }
    toString() {
      if (this.$tag === 0) {
        return "D9.DC22" + "(" + _dafny.toString(this.cf39) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf39, other.cf39);
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
          return D9.Default();
        }
      };
    }
  }

  $module.D10 = class D10 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC23(cf40) {
      let $dt = new D10(0);
      $dt.cf40 = cf40;
      return $dt;
    }
    get is_DC23() { return this.$tag === 0; }
    get dtor_cf40() { return this.cf40; }
    toString() {
      if (this.$tag === 0) {
        return "D10.DC23" + "(" + _dafny.toString(this.cf40) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf40 === other.cf40;
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
          return D10.Default();
        }
      };
    }
  }

  $module.D11 = class D11 {
    constructor(tag) {
      this.$tag = tag;
    }
    static create_DC25(cf42, cf43, cf44) {
      let $dt = new D11(0);
      $dt.cf42 = cf42;
      $dt.cf43 = cf43;
      $dt.cf44 = cf44;
      return $dt;
    }
    static create_DC26(cf45, cf46, cf47, cf48) {
      let $dt = new D11(1);
      $dt.cf45 = cf45;
      $dt.cf46 = cf46;
      $dt.cf47 = cf47;
      $dt.cf48 = cf48;
      return $dt;
    }
    static create_DC27(cf49, cf50) {
      let $dt = new D11(2);
      $dt.cf49 = cf49;
      $dt.cf50 = cf50;
      return $dt;
    }
    static create_DC24(cf41) {
      let $dt = new D11(3);
      $dt.cf41 = cf41;
      return $dt;
    }
    get is_DC25() { return this.$tag === 0; }
    get is_DC26() { return this.$tag === 1; }
    get is_DC27() { return this.$tag === 2; }
    get is_DC24() { return this.$tag === 3; }
    get dtor_cf42() { return this.cf42; }
    get dtor_cf43() { return this.cf43; }
    get dtor_cf44() { return this.cf44; }
    get dtor_cf45() { return this.cf45; }
    get dtor_cf46() { return this.cf46; }
    get dtor_cf47() { return this.cf47; }
    get dtor_cf48() { return this.cf48; }
    get dtor_cf49() { return this.cf49; }
    get dtor_cf50() { return this.cf50; }
    get dtor_cf41() { return this.cf41; }
    toString() {
      if (this.$tag === 0) {
        return "D11.DC25" + "(" + _dafny.toString(this.cf42) + ", " + _dafny.toString(this.cf43) + ", " + _dafny.toString(this.cf44) + ")";
      } else if (this.$tag === 1) {
        return "D11.DC26" + "(" + _dafny.toString(this.cf45) + ", " + _dafny.toString(this.cf46) + ", " + _dafny.toString(this.cf47) + ", " + this.cf48.toVerbatimString(true) + ")";
      } else if (this.$tag === 2) {
        return "D11.DC27" + "(" + _dafny.toString(this.cf49) + ", " + _dafny.toString(this.cf50) + ")";
      } else if (this.$tag === 3) {
        return "D11.DC24" + "(" + _dafny.toString(this.cf41) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf42 === other.cf42 && this.cf43 === other.cf43 && _dafny.areEqual(this.cf44, other.cf44);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf45 === other.cf45 && _dafny.areEqual(this.cf46, other.cf46) && _dafny.areEqual(this.cf47, other.cf47) && _dafny.areEqual(this.cf48, other.cf48);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf49, other.cf49) && _dafny.areEqual(this.cf50, other.cf50);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf41, other.cf41);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D11.create_DC25([], null, _dafny.ZERO);
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
    static create_DC29(cf52, cf53) {
      let $dt = new D12(0);
      $dt.cf52 = cf52;
      $dt.cf53 = cf53;
      return $dt;
    }
    static create_DC28(cf51) {
      let $dt = new D12(1);
      $dt.cf51 = cf51;
      return $dt;
    }
    static create_DC30(cf54) {
      let $dt = new D12(2);
      $dt.cf54 = cf54;
      return $dt;
    }
    get is_DC29() { return this.$tag === 0; }
    get is_DC28() { return this.$tag === 1; }
    get is_DC30() { return this.$tag === 2; }
    get dtor_cf52() { return this.cf52; }
    get dtor_cf53() { return this.cf53; }
    get dtor_cf51() { return this.cf51; }
    get dtor_cf54() { return this.cf54; }
    toString() {
      if (this.$tag === 0) {
        return "D12.DC29" + "(" + this.cf52.toVerbatimString(true) + ", " + _dafny.toString(this.cf53) + ")";
      } else if (this.$tag === 1) {
        return "D12.DC28" + "(" + _dafny.toString(this.cf51) + ")";
      } else if (this.$tag === 2) {
        return "D12.DC30" + "(" + _dafny.toString(this.cf54) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf52, other.cf52) && this.cf53 === other.cf53;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf51, other.cf51);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf54, other.cf54);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D12.create_DC29(_dafny.Seq.UnicodeFromString(""), false);
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
    static create_DC32(cf56, cf57, cf58, cf59, cf60) {
      let $dt = new D13(0);
      $dt.cf56 = cf56;
      $dt.cf57 = cf57;
      $dt.cf58 = cf58;
      $dt.cf59 = cf59;
      $dt.cf60 = cf60;
      return $dt;
    }
    static create_DC33(cf61, cf62, cf63, cf64, cf65) {
      let $dt = new D13(1);
      $dt.cf61 = cf61;
      $dt.cf62 = cf62;
      $dt.cf63 = cf63;
      $dt.cf64 = cf64;
      $dt.cf65 = cf65;
      return $dt;
    }
    static create_DC31(cf55) {
      let $dt = new D13(2);
      $dt.cf55 = cf55;
      return $dt;
    }
    static create_DC34(cf66) {
      let $dt = new D13(3);
      $dt.cf66 = cf66;
      return $dt;
    }
    get is_DC32() { return this.$tag === 0; }
    get is_DC33() { return this.$tag === 1; }
    get is_DC31() { return this.$tag === 2; }
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
    get dtor_cf55() { return this.cf55; }
    get dtor_cf66() { return this.cf66; }
    toString() {
      if (this.$tag === 0) {
        return "D13.DC32" + "(" + _dafny.toString(this.cf56) + ", " + _dafny.toString(this.cf57) + ", " + _dafny.toString(this.cf58) + ", " + _dafny.toString(this.cf59) + ", " + _dafny.toString(this.cf60) + ")";
      } else if (this.$tag === 1) {
        return "D13.DC33" + "(" + _dafny.toString(this.cf61) + ", " + _dafny.toString(this.cf62) + ", " + _dafny.toString(this.cf63) + ", " + _dafny.toString(this.cf64) + ", " + _dafny.toString(this.cf65) + ")";
      } else if (this.$tag === 2) {
        return "D13.DC31" + "(" + _dafny.toString(this.cf55) + ")";
      } else if (this.$tag === 3) {
        return "D13.DC34" + "(" + _dafny.toString(this.cf66) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf56, other.cf56) && this.cf57 === other.cf57 && this.cf58 === other.cf58 && this.cf59 === other.cf59 && this.cf60 === other.cf60;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf61 === other.cf61 && _dafny.areEqual(this.cf62, other.cf62) && _dafny.areEqual(this.cf63, other.cf63) && this.cf64 === other.cf64 && this.cf65 === other.cf65;
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf55, other.cf55);
      } else if (this.$tag === 3) {
        return other.$tag === 3 && _dafny.areEqual(this.cf66, other.cf66);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D13.create_DC32(_dafny.ZERO, false, false, false, false);
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
    static create_DC36(cf68, cf69, cf70, cf71) {
      let $dt = new D14(0);
      $dt.cf68 = cf68;
      $dt.cf69 = cf69;
      $dt.cf70 = cf70;
      $dt.cf71 = cf71;
      return $dt;
    }
    static create_DC35(cf67) {
      let $dt = new D14(1);
      $dt.cf67 = cf67;
      return $dt;
    }
    get is_DC36() { return this.$tag === 0; }
    get is_DC35() { return this.$tag === 1; }
    get dtor_cf68() { return this.cf68; }
    get dtor_cf69() { return this.cf69; }
    get dtor_cf70() { return this.cf70; }
    get dtor_cf71() { return this.cf71; }
    get dtor_cf67() { return this.cf67; }
    toString() {
      if (this.$tag === 0) {
        return "D14.DC36" + "(" + _dafny.toString(this.cf68) + ", " + _dafny.toString(this.cf69) + ", " + _dafny.toString(this.cf70) + ", " + _dafny.toString(this.cf71) + ")";
      } else if (this.$tag === 1) {
        return "D14.DC35" + "(" + _dafny.toString(this.cf67) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf68 === other.cf68 && _dafny.areEqual(this.cf69, other.cf69) && _dafny.areEqual(this.cf70, other.cf70) && _dafny.areEqual(this.cf71, other.cf71);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf67 === other.cf67;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D14.create_DC36(false, _dafny.ZERO, _dafny.ZERO, _dafny.Map.Empty);
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
    static create_DC37(cf72) {
      let $dt = new D15(0);
      $dt.cf72 = cf72;
      return $dt;
    }
    get is_DC37() { return this.$tag === 0; }
    get dtor_cf72() { return this.cf72; }
    toString() {
      if (this.$tag === 0) {
        return "D15.DC37" + "(" + _dafny.toString(this.cf72) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && this.cf72 === other.cf72;
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
    static create_DC39(cf74, cf75) {
      let $dt = new D16(0);
      $dt.cf74 = cf74;
      $dt.cf75 = cf75;
      return $dt;
    }
    static create_DC38(cf73) {
      let $dt = new D16(1);
      $dt.cf73 = cf73;
      return $dt;
    }
    get is_DC39() { return this.$tag === 0; }
    get is_DC38() { return this.$tag === 1; }
    get dtor_cf74() { return this.cf74; }
    get dtor_cf75() { return this.cf75; }
    get dtor_cf73() { return this.cf73; }
    toString() {
      if (this.$tag === 0) {
        return "D16.DC39" + "(" + _dafny.toString(this.cf74) + ", " + _dafny.toString(this.cf75) + ")";
      } else if (this.$tag === 1) {
        return "D16.DC38" + "(" + _dafny.toString(this.cf73) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf74, other.cf74) && _dafny.areEqual(this.cf75, other.cf75);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf73, other.cf73);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D16.create_DC39(_dafny.ZERO, _dafny.Map.Empty);
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
    static create_DC41(cf77, cf78) {
      let $dt = new D17(0);
      $dt.cf77 = cf77;
      $dt.cf78 = cf78;
      return $dt;
    }
    static create_DC40(cf76) {
      let $dt = new D17(1);
      $dt.cf76 = cf76;
      return $dt;
    }
    static create_DC42(cf79) {
      let $dt = new D17(2);
      $dt.cf79 = cf79;
      return $dt;
    }
    get is_DC41() { return this.$tag === 0; }
    get is_DC40() { return this.$tag === 1; }
    get is_DC42() { return this.$tag === 2; }
    get dtor_cf77() { return this.cf77; }
    get dtor_cf78() { return this.cf78; }
    get dtor_cf76() { return this.cf76; }
    get dtor_cf79() { return this.cf79; }
    toString() {
      if (this.$tag === 0) {
        return "D17.DC41" + "(" + _dafny.toString(this.cf77) + ", " + _dafny.toString(this.cf78) + ")";
      } else if (this.$tag === 1) {
        return "D17.DC40" + "(" + _dafny.toString(this.cf76) + ")";
      } else if (this.$tag === 2) {
        return "D17.DC42" + "(" + _dafny.toString(this.cf79) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf77, other.cf77) && this.cf78 === other.cf78;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && _dafny.areEqual(this.cf76, other.cf76);
      } else if (this.$tag === 2) {
        return other.$tag === 2 && _dafny.areEqual(this.cf79, other.cf79);
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D17.create_DC41(_dafny.ZERO, false);
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
    static create_DC44(cf81, cf82, cf83, cf84, cf85) {
      let $dt = new D18(0);
      $dt.cf81 = cf81;
      $dt.cf82 = cf82;
      $dt.cf83 = cf83;
      $dt.cf84 = cf84;
      $dt.cf85 = cf85;
      return $dt;
    }
    static create_DC43(cf80) {
      let $dt = new D18(1);
      $dt.cf80 = cf80;
      return $dt;
    }
    get is_DC44() { return this.$tag === 0; }
    get is_DC43() { return this.$tag === 1; }
    get dtor_cf81() { return this.cf81; }
    get dtor_cf82() { return this.cf82; }
    get dtor_cf83() { return this.cf83; }
    get dtor_cf84() { return this.cf84; }
    get dtor_cf85() { return this.cf85; }
    get dtor_cf80() { return this.cf80; }
    toString() {
      if (this.$tag === 0) {
        return "D18.DC44" + "(" + _dafny.toString(this.cf81) + ", " + _dafny.toString(this.cf82) + ", " + _dafny.toString(this.cf83) + ", " + _dafny.toString(this.cf84) + ", " + _dafny.toString(this.cf85) + ")";
      } else if (this.$tag === 1) {
        return "D18.DC43" + "(" + _dafny.toString(this.cf80) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf81, other.cf81) && _dafny.areEqual(this.cf82, other.cf82) && _dafny.areEqual(this.cf83, other.cf83) && this.cf84 === other.cf84 && this.cf85 === other.cf85;
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf80 === other.cf80;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D18.create_DC44(_dafny.Set.Empty, _dafny.ZERO, _dafny.ZERO, false, false);
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
    static create_DC46(cf87, cf88, cf89, cf90, cf91) {
      let $dt = new D19(0);
      $dt.cf87 = cf87;
      $dt.cf88 = cf88;
      $dt.cf89 = cf89;
      $dt.cf90 = cf90;
      $dt.cf91 = cf91;
      return $dt;
    }
    static create_DC45(cf86) {
      let $dt = new D19(1);
      $dt.cf86 = cf86;
      return $dt;
    }
    get is_DC46() { return this.$tag === 0; }
    get is_DC45() { return this.$tag === 1; }
    get dtor_cf87() { return this.cf87; }
    get dtor_cf88() { return this.cf88; }
    get dtor_cf89() { return this.cf89; }
    get dtor_cf90() { return this.cf90; }
    get dtor_cf91() { return this.cf91; }
    get dtor_cf86() { return this.cf86; }
    toString() {
      if (this.$tag === 0) {
        return "D19.DC46" + "(" + this.cf87.toVerbatimString(true) + ", " + _dafny.toString(this.cf88) + ", " + _dafny.toString(this.cf89) + ", " + _dafny.toString(this.cf90) + ", " + _dafny.toString(this.cf91) + ")";
      } else if (this.$tag === 1) {
        return "D19.DC45" + "(" + _dafny.toString(this.cf86) + ")";
      } else  {
        return "<unexpected>";
      }
    }
    equals(other) {
      if (this === other) {
        return true;
      } else if (this.$tag === 0) {
        return other.$tag === 0 && _dafny.areEqual(this.cf87, other.cf87) && _dafny.areEqual(this.cf88, other.cf88) && _dafny.areEqual(this.cf89, other.cf89) && _dafny.areEqual(this.cf90, other.cf90) && _dafny.areEqual(this.cf91, other.cf91);
      } else if (this.$tag === 1) {
        return other.$tag === 1 && this.cf86 === other.cf86;
      } else  {
        return false; // unexpected
      }
    }
    static Default() {
      return _module.D19.create_DC46(_dafny.Seq.UnicodeFromString(""), _dafny.ZERO, _dafny.ZERO, _dafny.ZERO, _dafny.Set.Empty);
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
      this.f0 = _dafny.ZERO;
      this.f1 = false;
      this.f2 = _dafny.ZERO;
      this.f5 = false;
      this._f3 = false;
      this._f4 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f0, f1, f2, f3, f4, f5) {
      let _this = this;
      (_this).f0 = f0;
      (_this).f1 = f1;
      (_this).f2 = f2;
      (_this)._f3 = f3;
      (_this)._f4 = f4;
      (_this).f5 = f5;
      return;
    }
    get f3() {
      let _this = this;
      return _this._f3;
    };
    get f4() {
      let _this = this;
      return _this._f4;
    };
  };

  $module.C0 = class C0 {
    constructor () {
      this._tname = "_module.C0";
      this._f13 = false;
    }
    _parentTraits() {
      return [];
    }
    __ctor(f13) {
      let _this = this;
      (_this)._f13 = f13;
      return;
    }
    fm23(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(new BigNumber(-986), new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("sicrkla"), _dafny.Seq.UnicodeFromString("nyetvlcd"))).length));
    };
    get f13() {
      let _this = this;
      return _this._f13;
    };
  };

  $module.C1 = class C1 {
    constructor () {
      this._tname = "_module.C1";
    }
    _parentTraits() {
      return [_module.T1];
    }
    __ctor() {
      let _this = this;
      return;
    }
    fm16(p0, globalState) {
      let _this = this;
      return !(!(_dafny.Seq.contains(_dafny.Seq.UnicodeFromString("lfuq"), new _dafny.CodePoint('d'.codePointAt(0)))) || (_dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(new _dafny.CodePoint('k'.codePointAt(0)), new _dafny.CodePoint('g'.codePointAt(0)), new _dafny.CodePoint('x'.codePointAt(0)), new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('l'.codePointAt(0))), _dafny.Seq.of(new _dafny.CodePoint('j'.codePointAt(0)), new _dafny.CodePoint('i'.codePointAt(0))))));
    };
    fm17(p0, p1, globalState) {
      let _this = this;
      return ((_dafny.ZERO).minus(((true) ? (new BigNumber((_dafny.Seq.UnicodeFromString("igmm")).length)) : (new BigNumber(913))))).plus(new BigNumber(828));
    };
    fm20(p0, p1, p2, p3, globalState) {
      let _this = this;
      return _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(699)), function (_376_i0) {
        return new _dafny.CodePoint('a'.codePointAt(0));
      }), _dafny.Seq.UnicodeFromString("qodmkq"));
    };
    m6(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let _377_v0;
      _377_v0 = new BigNumber(477);
      (globalState).f2 = _377_v0;
      let _hi3 = _377_v0;
      for (let _378_i0 = _377_v0; _378_i0.isLessThan(_hi3); _378_i0 = _378_i0.plus(_dafny.ONE)) {
        let _379_v1;
        _379_v1 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(65),_378_i0);
        let _380_v2;
        _380_v2 = true;
        let _381_v4;
        _381_v4 = _dafny.Seq.of(_dafny.Set.fromElements(_377_v0));
        let _382_v5;
        _382_v5 = _dafny.Seq.UnicodeFromString("a");
        let _383_v6;
        _383_v6 = new _dafny.CodePoint('t'.codePointAt(0));
        let _384_v7;
        _384_v7 = _dafny.Set.fromElements((_this).fm17(new BigNumber(-88), _dafny.Set.fromElements(new BigNumber((_dafny.Seq.update(_382_v5, _module.__default.safeIndex(_377_v0, new BigNumber((_382_v5).length)), _383_v6)).length)), globalState));
        let _385_v9;
        _385_v9 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(17),_380_v2);
        let _386_v10;
        _386_v10 = _dafny.MultiSet.fromElements(new BigNumber((_385_v9).length), _377_v0);
        let _387_v12;
        let _nw67 = Array((new BigNumber(14)).toNumber());
        _nw67[(_dafny.ZERO).toNumber()] = _379_v1;
        _nw67[(_dafny.ONE).toNumber()] = ((_380_v2) ? (_379_v1) : (function () {
          let _coll22 = new _dafny.Map();
          for (const _compr_22 of ((_381_v4)[_module.__default.safeIndex(_377_v0, new BigNumber((_381_v4).length))]).Elements) {
            let _388_v3 = _compr_22;
            if (((_381_v4)[_module.__default.safeIndex(_377_v0, new BigNumber((_381_v4).length))]).contains(_388_v3)) {
              _coll22.push([_module.__default.safeModuloInt(_388_v3, _378_i0),_378_i0]);
            }
          }
          return _coll22;
        }()));
        _nw67[(new BigNumber(2)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_377_v0,_377_v0);
        _nw67[(new BigNumber(3)).toNumber()] = _379_v1;
        _nw67[(new BigNumber(4)).toNumber()] = (_379_v1).Merge(_dafny.Map.Empty.slice().updateUnsafe(_377_v0,new BigNumber((_384_v7).length)));
        _nw67[(new BigNumber(5)).toNumber()] = _379_v1;
        _nw67[(new BigNumber(6)).toNumber()] = function () {
          let _coll23 = new _dafny.Map();
          for (const _compr_23 of (_386_v10).Elements) {
            let _389_v8 = _compr_23;
            if ((_386_v10).contains(_389_v8)) {
              _coll23.push([(_389_v8).multipliedBy(_378_i0),new BigNumber((_dafny.Seq.of(_378_i0)).length)]);
            }
          }
          return _coll23;
        }();
        _nw67[(new BigNumber(7)).toNumber()] = _379_v1;
        _nw67[(new BigNumber(8)).toNumber()] = (_379_v1).update(_377_v0, _377_v0);
        _nw67[(new BigNumber(9)).toNumber()] = _379_v1;
        _nw67[(new BigNumber(10)).toNumber()] = _379_v1;
        _nw67[(new BigNumber(11)).toNumber()] = _379_v1;
        _nw67[(new BigNumber(12)).toNumber()] = _379_v1;
        _nw67[(new BigNumber(13)).toNumber()] = function () {
          let _coll24 = new _dafny.Map();
          for (const _compr_24 of _dafny.IntegerRange(new BigNumber(318), new BigNumber(205))) {
            let _390_v11 = _compr_24;
            if (((new BigNumber(318)).isLessThanOrEqualTo(_390_v11)) && ((_390_v11).isLessThan(new BigNumber(205)))) {
              _coll24.push([(_390_v11).plus(_378_i0),_377_v0]);
            }
          }
          return _coll24;
        }();
        _387_v12 = _nw67;
        let _index58 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_387_v12).length));
        (_387_v12)[_index58] = _379_v1;
        let _index59 = _module.__default.safeIndex(new BigNumber(989), new BigNumber((_387_v12).length));
        (_387_v12)[_index59] = _379_v1;
        let _391_v13;
        _391_v13 = _dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber(754));
        _391_v13 = (_391_v13).update(!(_380_v2), _377_v0);
        let _392_v14;
        let _nw68 = Array((new BigNumber(8)).toNumber()).fill(false);
        _392_v14 = _nw68;
        let _393_v15;
        _393_v15 = _dafny.Map.Empty.slice().updateUnsafe(_380_v2,_392_v14);
        let _394_v16;
        _394_v16 = _module.D0.create_DC2(_377_v0, _383_v6, new BigNumber((_393_v15).length), _378_i0, (_dafny.ZERO).minus(new BigNumber((_module.__default.fm21(_382_v5, globalState)).length)));
        _383_v6 = (_394_v16).dtor_cf6;
        _384_v7 = ((_384_v7).Difference(_module.__default.fm22(globalState))).Difference(function () {
          let _coll25 = new _dafny.Set();
          for (const _compr_25 of _dafny.IntegerRange(new BigNumber(873), new BigNumber(595))) {
            let _395_v17 = _compr_25;
            if (((new BigNumber(873)).isLessThanOrEqualTo(_395_v17)) && ((_395_v17).isLessThan(new BigNumber(595)))) {
              _coll25.add(_module.__default.safeDivisionInt(_395_v17, _377_v0));
            }
          }
          return _coll25;
        }());
      }
      let _396_v18;
      let _nw69 = Array((new BigNumber(25)).toNumber()).fill(false);
      _396_v18 = _nw69;
      let _nw70 = Array((new BigNumber(19)).toNumber()).fill(false);
      _396_v18 = _nw70;
      let _397_i1;
      _397_i1 = _dafny.ZERO;
      L0: {
        while (!(false)) {
          C0: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_397_i1)) {
              break L0;
            }
            _397_i1 = (_397_i1).plus(_dafny.ONE);
            let _398_v19;
            _398_v19 = true;
            let _399_v20;
            _399_v20 = _dafny.MultiSet.fromElements(_398_v19, _398_v19);
            let _400_v21;
            _400_v21 = _dafny.Seq.of(_398_v19, _398_v19, (_this).fm16(_399_v20, globalState));
            let _401_v22;
            _401_v22 = new _dafny.CodePoint('b'.codePointAt(0));
            let _402_v23;
            _402_v23 = _dafny.Map.Empty.slice().updateUnsafe(_401_v22,_401_v22);
            (globalState).f5 = _dafny.areEqual(_400_v21, _dafny.Seq.update(_400_v21, _module.__default.safeIndex(new BigNumber((_402_v23).length), new BigNumber((_400_v21).length)), true));
            let _403_v24;
            _403_v24 = _dafny.Seq.UnicodeFromString("whpm");
            _403_v24 = _403_v24;
            let _404_v25;
            _404_v25 = _dafny.MultiSet.fromElements(_module.__default.safeModuloInt(_377_v0, _377_v0));
            _404_v25 = _404_v25;
            let _405_v26;
            _405_v26 = _module.D0.create_DC2(new BigNumber(955), new _dafny.CodePoint('y'.codePointAt(0)), _377_v0, _377_v0, _377_v0);
            let _406_v27;
            _406_v27 = _dafny.Map.Empty.slice().updateUnsafe(_377_v0,_405_v26);
            let _407_v28;
            _407_v28 = _dafny.Map.Empty.slice().updateUnsafe(_398_v19,new BigNumber((_406_v27).length));
            let _408_v29;
            let _init12 = ((_409_v0) => function (_410_i2) {
              return (_410_i2).plus(_409_v0);
            })(_377_v0);
            let _nw71 = Array((new BigNumber(20)).toNumber());
            for (let _i0_12 = 0; _i0_12 < new BigNumber(_nw71.length); _i0_12++) {
              _nw71[_i0_12] = _init12(new BigNumber(_i0_12));
            }
            _408_v29 = _nw71;
            let _411_v30;
            _411_v30 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.MultiSet.fromElements((_dafny.ZERO).minus(new BigNumber((_407_v28).length))),_408_v29);
            let _412_v31;
            let _nw72 = new _module.C0();
            _nw72.__ctor((_411_v30).equals(_411_v30));
            _412_v31 = _nw72;
          }
        }
      }
      let _413_v32;
      _413_v32 = false;
      let _414_v33;
      _414_v33 = _dafny.Set.fromElements(_413_v32);
      let _415_v34;
      _415_v34 = _module.D6.create_DC15(_414_v33);
      let _416_v35;
      let _nw73 = new _module.C0();
      _nw73.__ctor(((_415_v34).dtor_cf26).IsSubsetOf(_414_v33));
      _416_v35 = _nw73;
      let _417_v36;
      _417_v36 = _dafny.Seq.UnicodeFromString("le");
      let _418_v37;
      _418_v37 = _dafny.Map.Empty.slice().updateUnsafe(_413_v32,_dafny.Seq.Concat(_417_v36, _module.__default.fm2(_417_v36, globalState)));
      _418_v37 = (_418_v37).update(!(_413_v32), _417_v36);
      r0 = _377_v0;
      let _419_v38;
      _419_v38 = new _dafny.CodePoint('q'.codePointAt(0));
      let _420_v39;
      _420_v39 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(237),_dafny.Set.fromElements(true));
      let _421_v40;
      _421_v40 = _module.D0.create_DC0(_420_v39);
      let _422_v41;
      _422_v41 = _dafny.Seq.of(_413_v32, (_416_v35).f13);
      let _423_v43;
      let _nw74 = Array((new BigNumber(19)).toNumber());
      _nw74[(_dafny.ZERO).toNumber()] = _377_v0;
      _nw74[(_dafny.ONE).toNumber()] = _377_v0;
      _nw74[(new BigNumber(2)).toNumber()] = (_this).fm17(new BigNumber((_417_v36).length), function () {
        let _coll26 = new _dafny.Set();
        for (const _compr_26 of _dafny.IntegerRange(new BigNumber(381), new BigNumber(786))) {
          let _424_v42 = _compr_26;
          if (((new BigNumber(381)).isLessThanOrEqualTo(_424_v42)) && ((_424_v42).isLessThan(new BigNumber(786)))) {
            _coll26.add((_424_v42).plus(new BigNumber(330)));
          }
        }
        return _coll26;
      }(), globalState);
      _nw74[(new BigNumber(3)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(4)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(5)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(6)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(7)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(8)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(9)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(10)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(11)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(12)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(13)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(14)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(15)).toNumber()] = new BigNumber((_414_v33).length);
      _nw74[(new BigNumber(16)).toNumber()] = new BigNumber((_414_v33).length);
      _nw74[(new BigNumber(17)).toNumber()] = _377_v0;
      _nw74[(new BigNumber(18)).toNumber()] = new BigNumber((_417_v36).length);
      _423_v43 = _nw74;
      let _425_v44;
      _425_v44 = _dafny.Seq.of(_423_v43, _423_v43, _423_v43, _423_v43);
      r1 = ((_dafny.ZERO).minus((((_416_v35).f13) ? (new BigNumber((_417_v36).length)) : (new BigNumber(231))))).plus((_416_v35).fm23(_377_v0, _377_v0, _419_v38, _module.__default.fm1(_421_v40, (_416_v35).f13, _377_v0, (_422_v41)[_module.__default.safeIndex(new BigNumber((_425_v44).length), new BigNumber((_422_v41).length))], globalState), globalState));
      r2 = !((_416_v35).f13) || ((_422_v41)[_module.__default.safeIndex(_377_v0, new BigNumber((_422_v41).length))]);
      return [r0, r1, r2];
    }
  };

  $module.C2 = class C2 {
    constructor () {
      this._tname = "_module.C2";
      this.f12 = false;
      this._f11 = _dafny.Seq.UnicodeFromString("");
    }
    _parentTraits() {
      return [];
    }
    __ctor(f11, f12) {
      let _this = this;
      (_this)._f11 = f11;
      (_this).f12 = f12;
      return;
    }
    m5(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      if ((p1) && (((false) ? (!(p1)) : (p1)))) {
        let _426_v0;
        _426_v0 = new BigNumber(-405);
        let _427_v1;
        _427_v1 = _dafny.MultiSet.fromElements((new BigNumber(803)).minus(_426_v0));
        let _428_v2;
        _428_v2 = _dafny.Seq.of(_426_v0, new BigNumber(432));
        let _429_v3;
        _429_v3 = _dafny.Seq.of(_427_v1);
        _427_v1 = ((_dafny.MultiSet.FromArray(_428_v2)).Intersect(_427_v1)).Union((_429_v3)[_module.__default.safeIndex(_426_v0, new BigNumber((_429_v3).length))]);
        let _430_v4;
        _430_v4 = _dafny.MultiSet.fromElements(_this.f12, p0);
        let _431_v5;
        _431_v5 = _dafny.Set.fromElements(_430_v4, _430_v4);
        if ((_426_v0).isLessThanOrEqualTo(new BigNumber((_431_v5).length))) {
          let _index60 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((p2).length));
          (p2)[_index60] = p1;
          let _index61 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((p2).length));
          let _rhs47 = p0;
          let _rhs48 = p0;
          let _rhs49 = ((((_this.f12) ? (p1) : (_this.f12))) ? (true) : (true));
          let _lhs40 = _this;
          let _lhs41 = p2;
          let _lhs42 = _module.__default.safeIndex(new BigNumber(782), new BigNumber((p2).length));
          let _lhs43 = _this;
          _lhs40.f12 = _rhs47;
          _lhs41[_lhs42] = _rhs48;
          _lhs43.f12 = _rhs49;
          let _432_v6;
          _432_v6 = _dafny.Set.fromElements(_426_v0, new BigNumber((_427_v1).cardinality()));
          (globalState).f2 = new BigNumber(((_432_v6).Union((_432_v6).Intersect(_432_v6))).length);
          (globalState).f0 = new BigNumber(409);
          let _433_v7;
          _433_v7 = _dafny.Map.Empty.slice().updateUnsafe(false,_426_v0);
          let _434_v8;
          _434_v8 = _dafny.Map.Empty.slice().updateUnsafe(true,(p2)[_module.__default.safeIndex(new BigNumber(782), new BigNumber((p2).length))]);
          let _rhs50 = ((_433_v7).Merge(_dafny.Map.Empty.slice().updateUnsafe(p1,_426_v0))).Merge((_dafny.Map.Empty.slice().updateUnsafe(p1,new BigNumber((_434_v8).length))).Merge(_433_v7));
          let _rhs51 = p0;
          _433_v7 = _rhs50;
          r0 = _rhs51;
          _426_v0 = _426_v0;
        } else {
          let _435_v9;
          let _nw75 = new _module.C0();
          _nw75.__ctor(!(p0));
          _435_v9 = _nw75;
          let _436_v10;
          _436_v10 = _module.D5.create_DC12(_426_v0);
          (globalState).f2 = (_436_v10).dtor_cf20;
          (globalState).f0 = _426_v0;
          let _437_v11;
          _437_v11 = _dafny.Seq.UnicodeFromString("ts");
          _437_v11 = (_this).f11;
          (globalState).f0 = _426_v0;
        }
        if (_this.f12) {
          let _438_v12;
          let _nw76 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
          _438_v12 = _nw76;
          let _index62 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_438_v12).length));
          (_438_v12)[_index62] = _module.__default.fm24(_426_v0, globalState);
          let _index63 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_438_v12).length));
          (_438_v12)[_index63] = new BigNumber(-795);
          (globalState).f0 = _module.__default.safeDivisionInt((_426_v0).minus(_426_v0), _426_v0);
          (globalState).f0 = _426_v0;
          let _439_v13;
          _439_v13 = _dafny.Set.fromElements((((_430_v4).contains(p1)) ? ((_430_v4).get(p1)) : (_426_v0)), (_438_v12)[_module.__default.safeIndex(new BigNumber(714), new BigNumber((_438_v12).length))], (_438_v12)[_module.__default.safeIndex(new BigNumber(714), new BigNumber((_438_v12).length))]);
          let _440_v14;
          _440_v14 = _dafny.Map.Empty.slice().updateUnsafe(_430_v4,_426_v0);
          let _441_v15;
          _441_v15 = new _dafny.CodePoint('m'.codePointAt(0));
          let _442_v16;
          _442_v16 = _dafny.Map.Empty.slice().updateUnsafe(_441_v15,(_438_v12)[_module.__default.safeIndex(new BigNumber(714), new BigNumber((_438_v12).length))]);
          let _index64 = _module.__default.safeIndex(new BigNumber(714), new BigNumber((_438_v12).length));
          (_438_v12)[_index64] = (new BigNumber(((_439_v13).Union(_439_v13)).length)).multipliedBy(new BigNumber((((!(p0)) ? (_module.__default.fm25(new BigNumber((_440_v14).length), (_442_v16).update(_441_v15, _426_v0), _441_v15, globalState)) : (_428_v2))).length));
          (globalState).f1 = !((_438_v12)[_module.__default.safeIndex(new BigNumber(714), new BigNumber((_438_v12).length))]).isEqualTo((_dafny.ZERO).minus((_438_v12)[_module.__default.safeIndex(new BigNumber(714), new BigNumber((_438_v12).length))]));
        } else {
          let _443_v17;
          _443_v17 = _module.D1.create_DC5();
          let _444_v18;
          _444_v18 = new _dafny.CodePoint('o'.codePointAt(0));
          let _445_v19;
          let _nw77 = Array((new BigNumber(10)).toNumber());
          _nw77[(_dafny.ZERO).toNumber()] = _444_v18;
          _nw77[(_dafny.ONE).toNumber()] = ((_this).f11)[_module.__default.safeIndex(_426_v0, new BigNumber(((_this).f11).length))];
          _nw77[(new BigNumber(2)).toNumber()] = _444_v18;
          _nw77[(new BigNumber(3)).toNumber()] = _444_v18;
          _nw77[(new BigNumber(4)).toNumber()] = _444_v18;
          _nw77[(new BigNumber(5)).toNumber()] = _444_v18;
          _nw77[(new BigNumber(6)).toNumber()] = _444_v18;
          _nw77[(new BigNumber(7)).toNumber()] = _444_v18;
          _nw77[(new BigNumber(8)).toNumber()] = _444_v18;
          _nw77[(new BigNumber(9)).toNumber()] = _444_v18;
          _445_v19 = _nw77;
          let _index65 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_445_v19).length));
          (_445_v19)[_index65] = _444_v18;
          let _446_v20;
          _446_v20 = _dafny.Seq.of((_430_v4).update(false, _module.__default.abs(_426_v0)));
          let _index66 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_445_v19).length));
          let _rhs52 = _module.D1.create_DC5();
          let _rhs53 = _444_v18;
          let _rhs54 = ((_this).f11)[_module.__default.safeIndex((_dafny.ZERO).minus(_426_v0), new BigNumber(((_this).f11).length))];
          let _rhs55 = _426_v0;
          let _rhs56 = (new BigNumber((_446_v20).length)).isLessThan(_426_v0);
          let _lhs44 = _445_v19;
          let _lhs45 = _module.__default.safeIndex(new BigNumber(458), new BigNumber((_445_v19).length));
          let _lhs46 = globalState;
          let _lhs47 = _this;
          _443_v17 = _rhs52;
          _444_v18 = _rhs53;
          _lhs44[_lhs45] = _rhs54;
          _lhs46.f2 = _rhs55;
          _lhs47.f12 = _rhs56;
          let _447_v21;
          _447_v21 = _dafny.Seq.of(p1);
          let _rhs57 = _426_v0;
          let _rhs58 = _447_v21;
          let _lhs48 = globalState;
          _lhs48.f2 = _rhs57;
          _447_v21 = _rhs58;
          let _448_v22;
          let _nw78 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
          _448_v22 = _nw78;
          _448_v22 = _448_v22;
          (globalState).f1 = false;
          _428_v2 = _428_v2;
        }
        let _449_v23;
        _449_v23 = new _dafny.CodePoint('s'.codePointAt(0));
        let _450_v24;
        _450_v24 = _dafny.Set.fromElements(_426_v0);
        let _451_v25;
        _451_v25 = _module.D0.create_DC2(new BigNumber((_dafny.Set.fromElements(new BigNumber((_430_v4).cardinality()))).length), _449_v23, _426_v0, _426_v0, new BigNumber((_450_v24).length));
        let _pat_let_tv13 = _426_v0;
        let _452_v26;
        _452_v26 = _module.D0.create_DC0(_module.__default.fm0(_this.f12, p1, function (_pat_let14_0) {
  return function (_453_dt__update__tmp_h0) {
    return function (_pat_let15_0) {
      return function (_454_dt__update_hcf5_h0) {
        return _module.D0.create_DC2(_454_dt__update_hcf5_h0, (_453_dt__update__tmp_h0).dtor_cf6, (_453_dt__update__tmp_h0).dtor_cf7, (_453_dt__update__tmp_h0).dtor_cf8, (_453_dt__update__tmp_h0).dtor_cf9);
      }(_pat_let15_0);
    }(_pat_let_tv13);
  }(_pat_let14_0);
}(_451_v25), globalState));
        let _455_v27;
        _455_v27 = _dafny.Set.fromElements(_this.f12, true, p0);
        let _456_v28;
        _456_v28 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_426_v0,new BigNumber(168))).length),_455_v27);
        _452_v26 = _module.D0.create_DC0((_456_v28).Merge(_456_v28));
        let _457_v29;
        _457_v29 = _dafny.Map.Empty.slice().updateUnsafe(p0,new BigNumber((_428_v2).length));
        _457_v29 = (_457_v29).update(!(p0), _426_v0);
      } else {
        let _458_v30;
        _458_v30 = new BigNumber(12);
        let _459_v31;
        _459_v31 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f12,p0)).length),_458_v30);
        let _460_v32;
        let _nw79 = Array((new BigNumber(7)).toNumber());
        _nw79[(_dafny.ZERO).toNumber()] = new BigNumber(914);
        _nw79[(_dafny.ONE).toNumber()] = new BigNumber((_459_v31).length);
        _nw79[(new BigNumber(2)).toNumber()] = _458_v30;
        _nw79[(new BigNumber(3)).toNumber()] = new BigNumber(((_this).f11).length);
        _nw79[(new BigNumber(4)).toNumber()] = _458_v30;
        _nw79[(new BigNumber(5)).toNumber()] = _458_v30;
        _nw79[(new BigNumber(6)).toNumber()] = _458_v30;
        _460_v32 = _nw79;
        let _461_v33;
        _461_v33 = _dafny.Set.fromElements(_460_v32);
        let _462_v34;
        _462_v34 = _module.D5.create_DC11(_461_v33);
        let _463_v35;
        _463_v35 = _dafny.Seq.UnicodeFromString("lwclh");
        let _464_v36;
        _464_v36 = _dafny.Seq.of(_this.f12);
        let _465_v37;
        _465_v37 = _dafny.Map.Empty.slice().updateUnsafe(_458_v30,_module.__default.fm26(p1, globalState));
        let _466_v38;
        _466_v38 = _module.D0.create_DC0(_465_v37);
        let _rhs59 = _462_v34;
        let _rhs60 = _module.__default.fm1((((_464_v36)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("losi")).length), new BigNumber((_464_v36).length))]) ? (_466_v38) : (_466_v38)), (false) || (p0), _458_v30, p1, globalState);
        let _rhs61 = _module.__default.fm2((_this).f11, globalState);
        let _rhs62 = p0;
        let _rhs63 = !((p3).dtor_cf2);
        let _lhs49 = globalState;
        let _lhs50 = globalState;
        _462_v34 = _rhs59;
        _lhs49.f1 = _rhs60;
        _463_v35 = _rhs61;
        r0 = _rhs62;
        _lhs50.f1 = _rhs63;
        (globalState).f2 = _module.__default.safeModuloInt(_458_v30, _458_v30);
        let _467_v39;
        _467_v39 = _dafny.Seq.of(new BigNumber((_463_v35).length));
        let _468_v40;
        _468_v40 = _module.D3.create_DC8(_module.D1.create_DC4(_467_v39), _460_v32, new BigNumber(((_this).f11).length));
        let _469_v41;
        _469_v41 = _module.D3.create_DC9(((p0) ? (_468_v40) : (_468_v40)));
        let _470_v42;
        _470_v42 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm27(globalState),_467_v39);
        let _471_v43;
        _471_v43 = _dafny.Seq.of(_464_v36, _464_v36, _dafny.Seq.of(_module.__default.fm1(_module.D0.create_DC0(_465_v37), false, new BigNumber(842), p1, globalState), p1));
        let _472_v44;
        _472_v44 = _dafny.Map.Empty.slice().updateUnsafe(p1,_this.f12);
        let _rhs64 = _469_v41;
        let _rhs65 = (((_470_v42).contains(_471_v43)) ? ((_470_v42).get(_471_v43)) : (_dafny.Seq.of(new BigNumber((_472_v44).length))));
        let _rhs66 = (_458_v30).minus((_dafny.ZERO).minus((_dafny.ZERO).minus(((_this.f12) ? (_458_v30) : (_458_v30)))));
        _469_v41 = _rhs64;
        _467_v39 = _rhs65;
        _458_v30 = _rhs66;
        r0 = _this.f12;
        _459_v31 = (_459_v31).update(_458_v30, _458_v30);
      }
      let _473_v45;
      _473_v45 = new BigNumber(997);
      let _474_v46;
      _474_v46 = _dafny.Seq.of(_473_v45);
      _474_v46 = _474_v46;
      if (p1) {
        let _475_v47;
        _475_v47 = _dafny.Seq.of(!(_this.f12));
        (globalState).f0 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_473_v45).minus(new BigNumber((_475_v47).length))), new BigNumber(921));
        r0 = _this.f12;
        let _476_v48;
        let _nw80 = Array((new BigNumber(23)).toNumber());
        _476_v48 = _nw80;
        let _477_v49;
        _477_v49 = _dafny.Seq.of(_476_v48, _476_v48);
        (globalState).f2 = new BigNumber((_477_v49).length);
        (globalState).f1 = !(p0);
        let _478_v50;
        let _nw81 = new _module.C1();
        _nw81.__ctor();
        _478_v50 = _nw81;
      } else {
        let _479_v51;
        let _nw82 = new _module.C0();
        _nw82.__ctor(p0);
        _479_v51 = _nw82;
        if (!_dafny.areEqual((_this).f11, _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("cyofsjs"), (_this).f11))) {
          let _index67 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((p2).length));
          (p2)[_index67] = (((_479_v51).f13) ? (p1) : (p1));
          let _index68 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((p2).length));
          (p2)[_index68] = p0;
          let _480_v52;
          _480_v52 = _dafny.MultiSet.fromElements(p0);
          let _481_v53;
          _481_v53 = _dafny.Set.fromElements(p1);
          let _482_v54;
          _482_v54 = _module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe((((_480_v52).contains((_479_v51).f13)) ? ((_480_v52).get((_479_v51).f13)) : (new BigNumber((_474_v46).length))),_481_v53));
          let _483_v55;
          _483_v55 = _dafny.Map.Empty.slice().updateUnsafe((_479_v51).f13,new BigNumber(-292));
          (globalState).f5 = _module.__default.fm1(_482_v54, (_479_v51).f13, (new BigNumber((_483_v55).length)).multipliedBy(_473_v45), (_479_v51).f13, globalState);
          let _484_v56;
          let _nw83 = Array((new BigNumber(6)).toNumber());
          _nw83[(_dafny.ZERO).toNumber()] = p1;
          _nw83[(_dafny.ONE).toNumber()] = (_480_v52).contains((_479_v51).f13);
          _nw83[(new BigNumber(2)).toNumber()] = (_479_v51).f13;
          _nw83[(new BigNumber(3)).toNumber()] = (_479_v51).f13;
          _nw83[(new BigNumber(4)).toNumber()] = (((_479_v51).f13) ? (p0) : (_module.__default.fm1(_482_v54, _this.f12, _473_v45, p1, globalState)));
          _nw83[(new BigNumber(5)).toNumber()] = _this.f12;
          _484_v56 = _nw83;
          let _485_v57;
          _485_v57 = _484_v56;
          let _486_v58;
          _486_v58 = _dafny.Seq.of((_485_v57), _484_v56);
          _484_v56 = (_486_v58)[_module.__default.safeIndex(_473_v45, new BigNumber((_486_v58).length))];
          let _index69 = _module.__default.safeIndex(new BigNumber(889), new BigNumber((p2).length));
          (p2)[_index69] = ((_this.f12) ? ((((_479_v51).f13) ? (_this.f12) : ((p2)[_module.__default.safeIndex(new BigNumber(889), new BigNumber((p2).length))]))) : ((_479_v51).f13));
          let _487_v59;
          _487_v59 = _dafny.Map.Empty.slice().updateUnsafe(_473_v45,_dafny.Set.fromElements(p0, (p2)[_module.__default.safeIndex(new BigNumber(889), new BigNumber((p2).length))], p1, p0));
          (globalState).f1 = ((_module.__default.fm1(_module.D0.create_DC0(_487_v59), _this.f12, new BigNumber((_481_v53).length), (_479_v51).f13, globalState)) ? (p0) : ((_479_v51).f13));
        } else {
          let _488_v61;
          _488_v61 = _dafny.Map.Empty.slice().updateUnsafe(_this.f12,_473_v45);
          let _489_v62;
          let _nw84 = Array((new BigNumber(20)).toNumber());
          _nw84[(_dafny.ZERO).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.of(_473_v45), _474_v46);
          _nw84[(_dafny.ONE).toNumber()] = _474_v46;
          _nw84[(new BigNumber(2)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(637)), ((_490_v45) => function (_491_i0) {
            return _490_v45;
          })(_473_v45)), _474_v46);
          _nw84[(new BigNumber(3)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(906)), ((_492_v45) => function (_493_i1) {
            return _492_v45;
          })(_473_v45)), _474_v46);
          _nw84[(new BigNumber(4)).toNumber()] = _474_v46;
          _nw84[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_474_v46, _474_v46);
          _nw84[(new BigNumber(6)).toNumber()] = (_module.D1.create_DC4(_474_v46)).dtor_cf11;
          _nw84[(new BigNumber(7)).toNumber()] = _dafny.Seq.update(_474_v46, _module.__default.safeIndex(_473_v45, new BigNumber((_474_v46).length)), (_474_v46)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(function () {
            let _coll27 = new _dafny.Set();
            for (const _compr_27 of (_474_v46).Elements) {
              let _494_v60 = _compr_27;
              if (_dafny.Seq.contains(_474_v46, _494_v60)) {
                _coll27.add((_494_v60).minus((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-190)), function (_495_i2) {
                  return new _dafny.CodePoint('g'.codePointAt(0));
                })).length))));
              }
            }
            return _coll27;
          }())).length), new BigNumber((_474_v46).length))]);
          _nw84[(new BigNumber(8)).toNumber()] = _dafny.Seq.Concat(_474_v46, _dafny.Seq.of(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_488_v61,(_479_v51).f13)).length), _473_v45, _473_v45));
          _nw84[(new BigNumber(9)).toNumber()] = _474_v46;
          _nw84[(new BigNumber(10)).toNumber()] = _474_v46;
          _nw84[(new BigNumber(11)).toNumber()] = _dafny.Seq.Concat(_474_v46, _474_v46);
          _nw84[(new BigNumber(12)).toNumber()] = _474_v46;
          _nw84[(new BigNumber(13)).toNumber()] = _dafny.Seq.update(_474_v46, _module.__default.safeIndex(new BigNumber((_dafny.Seq.UnicodeFromString("usfonr")).length), new BigNumber((_474_v46).length)), _473_v45);
          _nw84[(new BigNumber(14)).toNumber()] = (((_479_v51).f13) ? (_474_v46) : (_474_v46));
          _nw84[(new BigNumber(15)).toNumber()] = _dafny.Seq.of(_473_v45);
          _nw84[(new BigNumber(16)).toNumber()] = _474_v46;
          _nw84[(new BigNumber(17)).toNumber()] = _474_v46;
          _nw84[(new BigNumber(18)).toNumber()] = _474_v46;
          _nw84[(new BigNumber(19)).toNumber()] = (((_479_v51).f13) ? (_474_v46) : (_474_v46));
          _489_v62 = _nw84;
          let _index70 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_489_v62).length));
          (_489_v62)[_index70] = _dafny.Seq.Concat(_474_v46, _474_v46);
          let _index71 = _module.__default.safeIndex(new BigNumber(178), new BigNumber((_489_v62).length));
          (_489_v62)[_index71] = _474_v46;
          let _496_v63;
          _496_v63 = _module.D3.create_DC7(_dafny.Seq.UnicodeFromString("x"));
          let _497_v64;
          let _nw85 = Array((new BigNumber(15)).toNumber());
          _nw85[(_dafny.ZERO).toNumber()] = (_this).f11;
          _nw85[(_dafny.ONE).toNumber()] = _module.__default.fm2((_this).f11, globalState);
          _nw85[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(728)), function (_498_i3) {
            return new _dafny.CodePoint('r'.codePointAt(0));
          });
          _nw85[(new BigNumber(3)).toNumber()] = (_this).f11;
          _nw85[(new BigNumber(4)).toNumber()] = (_this).f11;
          _nw85[(new BigNumber(5)).toNumber()] = (_this).f11;
          _nw85[(new BigNumber(6)).toNumber()] = (_this).f11;
          _nw85[(new BigNumber(7)).toNumber()] = _module.__default.fm2((_this).f11, globalState);
          _nw85[(new BigNumber(8)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(382)), function (_499_i4) {
            return new _dafny.CodePoint('t'.codePointAt(0));
          });
          _nw85[(new BigNumber(9)).toNumber()] = _dafny.Seq.Concat((_this).f11, (_this).f11);
          _nw85[(new BigNumber(10)).toNumber()] = (_this).f11;
          _nw85[(new BigNumber(11)).toNumber()] = (_this).f11;
          _nw85[(new BigNumber(12)).toNumber()] = (_496_v63).dtor_cf13;
          _nw85[(new BigNumber(13)).toNumber()] = (_this).f11;
          _nw85[(new BigNumber(14)).toNumber()] = (_this).f11;
          _497_v64 = _nw85;
          _497_v64 = _497_v64;
          let _500_v67;
          _500_v67 = _dafny.Seq.of(true);
          let _501_v68;
          _501_v68 = _dafny.Set.fromElements(_473_v45, _473_v45, new BigNumber(162), new BigNumber((_500_v67).length));
          let _502_v69;
          _502_v69 = _dafny.Seq.of(_501_v68, _501_v68, _501_v68, _501_v68, _501_v68);
          let _503_v70;
          _503_v70 = new _dafny.CodePoint('p'.codePointAt(0));
          let _504_v71;
          _504_v71 = _dafny.Map.Empty.slice().updateUnsafe(_503_v70,_473_v45);
          let _505_v72;
          _505_v72 = _module.D5.create_DC12(_473_v45);
          let _506_v73;
          _506_v73 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_473_v45, new BigNumber(((_this).f11).length), new BigNumber((_504_v71).length), _473_v45, _473_v45),_505_v72);
          let _507_v74;
          _507_v74 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
          let _508_v76;
          _508_v76 = _dafny.Map.Empty.slice().updateUnsafe(_501_v68,_473_v45);
          let _509_v77;
          _509_v77 = _dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(_473_v45, new BigNumber((_507_v74).length)),_505_v72), function () {
            let _coll28 = new _dafny.Map();
            for (const _compr_28 of (_508_v76).Keys.Elements) {
              let _510_v75 = _compr_28;
              if ((_508_v76).contains(_510_v75)) {
                _coll28.push([_510_v75,_505_v72]);
              }
            }
            return _coll28;
          }(), _dafny.Map.Empty.slice().updateUnsafe((_502_v69)[_module.__default.safeIndex(_473_v45, new BigNumber((_502_v69).length))],_505_v72), _dafny.Map.Empty.slice().updateUnsafe(_501_v68,_505_v72));
          let _511_v79;
          _511_v79 = _dafny.Set.fromElements(_501_v68, _501_v68);
          let _pat_let_tv14 = _473_v45;
          let _512_v80;
          let _nw86 = Array((new BigNumber(17)).toNumber());
          _nw86[(_dafny.ZERO).toNumber()] = function () {
            let _coll29 = new _dafny.Map();
            for (const _compr_29 of (function () {
              let _coll30 = new _dafny.Map();
              for (const _compr_30 of (_502_v69).Elements) {
                let _513_v66 = _compr_30;
                if (_dafny.Seq.contains(_502_v69, _513_v66)) {
                  _coll30.push([_513_v66,p0]);
                }
              }
              return _coll30;
            }()).Keys.Elements) {
              let _514_v65 = _compr_29;
              if ((function () {
                let _coll31 = new _dafny.Map();
                for (const _compr_31 of (_502_v69).Elements) {
                  let _513_v66 = _compr_31;
                  if (_dafny.Seq.contains(_502_v69, _513_v66)) {
                    _coll31.push([_513_v66,p0]);
                  }
                }
                return _coll31;
              }()).contains(_514_v65)) {
                _coll29.push([_514_v65,_module.D5.create_DC12(_473_v45)]);
              }
            }
            return _coll29;
          }();
          _nw86[(_dafny.ONE).toNumber()] = _506_v73;
          _nw86[(new BigNumber(2)).toNumber()] = (_509_v77)[_module.__default.safeIndex((_dafny.ZERO).minus(_473_v45), new BigNumber((_509_v77).length))];
          _nw86[(new BigNumber(3)).toNumber()] = function () {
            let _coll32 = new _dafny.Map();
            for (const _compr_32 of (_511_v79).Elements) {
              let _515_v78 = _compr_32;
              if ((_511_v79).contains(_515_v78)) {
                _coll32.push([_515_v78,_505_v72]);
              }
            }
            return _coll32;
          }();
          _nw86[(new BigNumber(4)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(5)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(6)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(7)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(8)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(9)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(10)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(11)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(12)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(13)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(14)).toNumber()] = _506_v73;
          _nw86[(new BigNumber(15)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_501_v68,function (_pat_let16_0) {
            return function (_516_dt__update__tmp_h1) {
              return function (_pat_let17_0) {
                return function (_517_dt__update_hcf20_h0) {
                  return _module.D5.create_DC12(_517_dt__update_hcf20_h0);
                }(_pat_let17_0);
              }(_pat_let_tv14);
            }(_pat_let16_0);
          }(_505_v72));
          _nw86[(new BigNumber(16)).toNumber()] = _506_v73;
          _512_v80 = _nw86;
          let _index72 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_512_v80).length));
          (_512_v80)[_index72] = _506_v73;
          let _index73 = _module.__default.safeIndex(new BigNumber(511), new BigNumber((_512_v80).length));
          (_512_v80)[_index73] = (_506_v73).Merge(_dafny.Map.Empty.slice().updateUnsafe(_501_v68,_505_v72));
          let _518_v81;
          _518_v81 = _dafny.Map.Empty.slice().updateUnsafe(_473_v45,_473_v45);
          let _519_v82;
          _519_v82 = _module.D1.create_DC4(_dafny.Seq.of(new BigNumber((_518_v81).length), _473_v45, new BigNumber((_500_v67).length)));
          let _520_v83;
          let _init13 = ((_521_v45) => function (_522_i5) {
            return _module.__default.safeModuloInt(_522_i5, _521_v45);
          })(_473_v45);
          let _nw87 = Array((new BigNumber(8)).toNumber());
          for (let _i0_13 = 0; _i0_13 < new BigNumber(_nw87.length); _i0_13++) {
            _nw87[_i0_13] = _init13(new BigNumber(_i0_13));
          }
          _520_v83 = _nw87;
          let _523_v84;
          _523_v84 = _module.D3.create_DC8(_519_v82, _520_v83, new BigNumber(((_this).f11).length));
          (globalState).f2 = (((p0) ? (_module.__default.fm24((_523_v84).dtor_cf16, globalState)) : (_473_v45))).multipliedBy(_473_v45);
          let _524_v85;
          _524_v85 = _dafny.Seq.of((_this).f11);
          _524_v85 = _dafny.Seq.Concat(_dafny.Seq.update(_dafny.Seq.Concat(_524_v85, _524_v85), _module.__default.safeIndex(_473_v45, new BigNumber((_dafny.Seq.Concat(_524_v85, _524_v85)).length)), (_this).f11), _524_v85);
        }
        let _525_v86;
        _525_v86 = new _dafny.CodePoint('p'.codePointAt(0));
        _525_v86 = _525_v86;
        let _526_v87;
        _526_v87 = _dafny.Seq.UnicodeFromString("bdx");
        let _527_v88;
        _527_v88 = _module.D3.create_DC7(_526_v87);
        _526_v87 = (_527_v88).dtor_cf13;
        let _528_v90;
        _528_v90 = _module.D0.create_DC0(function () {
  let _coll33 = new _dafny.Map();
  for (const _compr_33 of _dafny.IntegerRange(new BigNumber(-760), new BigNumber(-629))) {
    let _529_v89 = _compr_33;
    if (((new BigNumber(-760)).isLessThanOrEqualTo(_529_v89)) && ((_529_v89).isLessThan(new BigNumber(-629)))) {
      _coll33.push([_module.__default.safeModuloInt(_529_v89, new BigNumber((_dafny.Set.fromElements(false, (_479_v51).f13)).length)),_dafny.Set.fromElements(false)]);
    }
  }
  return _coll33;
}());
        let _530_v91;
        _530_v91 = _dafny.MultiSet.fromElements(_module.__default.fm1(_528_v90, !((_479_v51).f13), _473_v45, !((_479_v51).f13), globalState));
        let _531_v92;
        let _nw88 = new _module.C1();
        _nw88.__ctor();
        _531_v92 = _nw88;
        let _532_v93;
        let _nw89 = Array((new BigNumber(10)).toNumber());
        _nw89[(_dafny.ZERO).toNumber()] = _473_v45;
        _nw89[(_dafny.ONE).toNumber()] = _473_v45;
        _nw89[(new BigNumber(2)).toNumber()] = new BigNumber(((_dafny.MultiSet.fromElements(p0)).Difference(_530_v91)).cardinality());
        _nw89[(new BigNumber(3)).toNumber()] = _module.__default.safeDivisionInt(_473_v45, new BigNumber(((_this).f11).length));
        _nw89[(new BigNumber(4)).toNumber()] = _473_v45;
        _nw89[(new BigNumber(5)).toNumber()] = _473_v45;
        _nw89[(new BigNumber(6)).toNumber()] = new BigNumber((_526_v87).length);
        _nw89[(new BigNumber(7)).toNumber()] = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_dafny.ZERO).minus(_473_v45)), _473_v45);
        _nw89[(new BigNumber(8)).toNumber()] = _module.__default.safeModuloInt(new BigNumber((_dafny.Set.fromElements(_531_v92, _531_v92, _531_v92, _531_v92, _531_v92)).length), _473_v45);
        _nw89[(new BigNumber(9)).toNumber()] = new BigNumber(((_527_v88).dtor_cf13).length);
        _532_v93 = _nw89;
        let _index74 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_532_v93).length));
        (_532_v93)[_index74] = _473_v45;
        let _533_v94;
        _533_v94 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus(_473_v45));
        let _534_v95;
        _534_v95 = _module.D3.create_DC8(_module.D1.create_DC4(_474_v46), _532_v93, new BigNumber((_533_v94).cardinality()));
        let _index75 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_532_v93).length));
        let _rhs67 = (_534_v95).dtor_cf16;
        let _rhs68 = p1;
        let _lhs51 = _532_v93;
        let _lhs52 = _module.__default.safeIndex(new BigNumber(792), new BigNumber((_532_v93).length));
        let _lhs53 = globalState;
        _lhs51[_lhs52] = _rhs67;
        _lhs53.f5 = _rhs68;
      }
      let _535_v96;
      let _init14 = ((_536_p0) => function (_537_i7) {
        return _module.__default.safeModuloInt(_537_i7, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_this.f12,_536_p0)).length));
      })(p0);
      let _nw90 = Array((new BigNumber(20)).toNumber());
      for (let _i0_14 = 0; _i0_14 < new BigNumber(_nw90.length); _i0_14++) {
        _nw90[_i0_14] = _init14(new BigNumber(_i0_14));
      }
      _535_v96 = _nw90;
      for (const _guard_loop_1 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_535_v96).length))) {
        let _538_i6 = _guard_loop_1;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_538_i6)) && ((_538_i6).isLessThan(new BigNumber((_535_v96).length))))) {
          (_535_v96)[(_538_i6)] = (_538_i6).plus(new BigNumber((_dafny.Set.fromElements(_473_v45, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(249)), function (_539_i8) {
            return new _dafny.CodePoint('o'.codePointAt(0));
          })).length))).length));
        }
      }
      let _540_v97;
      _540_v97 = _dafny.Seq.of(_this.f12, p0, _this.f12, false);
      let _541_v98;
      _541_v98 = _dafny.Set.fromElements(_540_v97, _540_v97);
      if ((_541_v98).IsSubsetOf(_541_v98)) {
        let _542_v99;
        _542_v99 = new _dafny.CodePoint('t'.codePointAt(0));
        let _543_v100;
        _543_v100 = _dafny.Map.Empty.slice().updateUnsafe(p0,_473_v45);
        let _544_v101;
        _544_v101 = _module.D0.create_DC2(_473_v45, _542_v99, _473_v45, _473_v45, (((_543_v100).contains(p0)) ? ((_543_v100).get(p0)) : (_473_v45)));
        let _545_v102;
        _545_v102 = _module.D0.create_DC3(_544_v101);
        let _546_v103;
        _546_v103 = _dafny.Set.fromElements(!(!(_this.f12)), p1);
        _545_v102 = (((_546_v103).IsSubsetOf(_546_v103)) ? (_545_v102) : (_545_v102));
        let _547_v104;
        let _nw91 = new _module.C0();
        _nw91.__ctor(_this.f12);
        _547_v104 = _nw91;
        let _548_v105;
        let _nw92 = new _module.C1();
        _nw92.__ctor();
        _548_v105 = _nw92;
        (_this).f12 = _this.f12;
        _542_v99 = _542_v99;
      } else {
        let _549_v106;
        _549_v106 = _dafny.Set.fromElements(_473_v45);
        let _550_v107;
        _550_v107 = _module.D7.create_DC17(_540_v97);
        let _551_v108;
        _551_v108 = _dafny.Map.Empty.slice().updateUnsafe(_473_v45,_549_v106);
        let _552_v112;
        let _nw93 = Array((new BigNumber(27)).toNumber());
        _nw93[(_dafny.ZERO).toNumber()] = _549_v106;
        _nw93[(_dafny.ONE).toNumber()] = (_549_v106).Union(_549_v106);
        _nw93[(new BigNumber(2)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(3)).toNumber()] = _dafny.Set.fromElements(new BigNumber(((_550_v107).dtor_cf28).length));
        _nw93[(new BigNumber(4)).toNumber()] = (((_551_v108).contains(_473_v45)) ? ((_551_v108).get(_473_v45)) : (function () {
          let _coll34 = new _dafny.Set();
          for (const _compr_34 of _dafny.IntegerRange(new BigNumber(170), new BigNumber(858))) {
            let _553_v109 = _compr_34;
            if (((new BigNumber(170)).isLessThanOrEqualTo(_553_v109)) && ((_553_v109).isLessThan(new BigNumber(858)))) {
              _coll34.add(_module.__default.safeModuloInt(_553_v109, _473_v45));
            }
          }
          return _coll34;
        }()));
        _nw93[(new BigNumber(5)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(6)).toNumber()] = (_549_v106).Union(_549_v106);
        _nw93[(new BigNumber(7)).toNumber()] = _dafny.Set.fromElements(_473_v45);
        _nw93[(new BigNumber(8)).toNumber()] = (_549_v106).Intersect(_549_v106);
        _nw93[(new BigNumber(9)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(10)).toNumber()] = (_dafny.Set.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_473_v45,p0)).length))).Intersect(_dafny.Set.fromElements(new BigNumber(-132)));
        _nw93[(new BigNumber(11)).toNumber()] = function () {
          let _coll35 = new _dafny.Set();
          for (const _compr_35 of _dafny.IntegerRange(new BigNumber(790), new BigNumber(697))) {
            let _554_v110 = _compr_35;
            if (((new BigNumber(790)).isLessThanOrEqualTo(_554_v110)) && ((_554_v110).isLessThan(new BigNumber(697)))) {
              _coll35.add((_554_v110).plus((_474_v46)[_module.__default.safeIndex(_473_v45, new BigNumber((_474_v46).length))]));
            }
          }
          return _coll35;
        }();
        _nw93[(new BigNumber(12)).toNumber()] = (_549_v106).Union(_549_v106);
        _nw93[(new BigNumber(13)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(14)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(15)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(16)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(17)).toNumber()] = ((p0) ? (function () {
          let _coll36 = new _dafny.Set();
          for (const _compr_36 of _dafny.IntegerRange(new BigNumber(-491), new BigNumber(281))) {
            let _555_v111 = _compr_36;
            if (((new BigNumber(-491)).isLessThanOrEqualTo(_555_v111)) && ((_555_v111).isLessThan(new BigNumber(281)))) {
              _coll36.add(_module.__default.safeDivisionInt(_555_v111, new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_474_v46).length),_473_v45)).length)));
            }
          }
          return _coll36;
        }()) : (_549_v106));
        _nw93[(new BigNumber(18)).toNumber()] = (_module.__default.fm22(globalState)).Union(_549_v106);
        _nw93[(new BigNumber(19)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(20)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(21)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(22)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(23)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(24)).toNumber()] = _549_v106;
        _nw93[(new BigNumber(25)).toNumber()] = _dafny.Set.fromElements(_473_v45, new BigNumber((_540_v97).length), _473_v45, _473_v45);
        _nw93[(new BigNumber(26)).toNumber()] = _549_v106;
        _552_v112 = _nw93;
        let _index76 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_552_v112).length));
        (_552_v112)[_index76] = _549_v106;
        let _index77 = _module.__default.safeIndex(new BigNumber(740), new BigNumber((_552_v112).length));
        (_552_v112)[_index77] = (_549_v106).Union(function () {
          let _coll37 = new _dafny.Set();
          for (const _compr_37 of (_module.__default.fm22(globalState)).Elements) {
            let _556_v113 = _compr_37;
            if ((_module.__default.fm22(globalState)).contains(_556_v113)) {
              _coll37.add((_556_v113).plus((_module.D0.create_DC1(new BigNumber(-491), !(true), new BigNumber((_dafny.Seq.of(true)).length), false)).dtor_cf3));
            }
          }
          return _coll37;
        }());
        (globalState).f1 = p0;
        let _index78 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_535_v96).length));
        (_535_v96)[_index78] = new BigNumber((_dafny.Seq.Concat(_540_v97, _540_v97)).length);
        let _557_v114;
        _557_v114 = _dafny.Seq.of(_dafny.Seq.Concat((_this).f11, (_this).f11));
        let _index79 = _module.__default.safeIndex(new BigNumber(445), new BigNumber((_535_v96).length));
        (_535_v96)[_index79] = new BigNumber((_dafny.Seq.update((_557_v114)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_473_v45, _473_v45), new BigNumber((_557_v114).length))], _module.__default.safeIndex(_473_v45, new BigNumber(((_557_v114)[_module.__default.safeIndex(_module.__default.safeDivisionInt(_473_v45, _473_v45), new BigNumber((_557_v114).length))]).length)), _module.__default.fm28((_474_v46)[_module.__default.safeIndex(_473_v45, new BigNumber((_474_v46).length))], (_dafny.ZERO).minus(_module.__default.fm24(new BigNumber((_dafny.Seq.UnicodeFromString("xvcevpcui")).length), globalState)), _this.f12, globalState))).length);
        let _558_v115;
        let _nw94 = new _module.C0();
        _nw94.__ctor(false);
        _558_v115 = _nw94;
        let _index80 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((p2).length));
        (p2)[_index80] = _dafny.Seq.IsPrefixOf(_540_v97, _540_v97);
        let _index81 = _module.__default.safeIndex(new BigNumber(251), new BigNumber((p2).length));
        (p2)[_index81] = (_558_v115).f13;
      }
      let _559_v116;
      _559_v116 = _module.D6.create_DC16(_473_v45);
      (globalState).f2 = (((_this.f12) ? (_module.__default.fm29(p0, _473_v45, globalState)) : (_559_v116))).dtor_cf27;
      r0 = p1;
      return r0;
    }
    get f11() {
      let _this = this;
      return _this._f11;
    };
  };

  $module.C3 = class C3 {
    constructor () {
      this._tname = "_module.C3";
      this._f6 = _dafny.ZERO;
      this._f9 = _dafny.ZERO;
      this._f10 = _dafny.Seq.of();
    }
    _parentTraits() {
      return [_module.T0, _module.T1];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f9, f10, f6) {
      let _this = this;
      (_this)._f9 = f9;
      (_this)._f10 = f10;
      (_this)._f6 = f6;
      return;
    }
    fm16(p0, globalState) {
      let _this = this;
      return !((_this).f6).isEqualTo((_this).f6);
    };
    fm17(p0, p1, globalState) {
      let _this = this;
      return _module.__default.safeModuloInt((_this).f9, new BigNumber(((_dafny.Set.fromElements(!(true))).Union(_dafny.Set.fromElements(false))).length));
    };
    fm18(p0, p1, p2, globalState) {
      let _this = this;
      return ((_dafny.Map.Empty.slice().updateUnsafe(true,true)).Merge(_dafny.Map.Empty.slice().updateUnsafe(true,true))).Merge(_dafny.Map.Empty.slice().updateUnsafe(false,true));
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let _560_v0;
      let _nw95 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
      _560_v0 = _nw95;
      for (const _guard_loop_2 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_560_v0).length))) {
        let _561_i0 = _guard_loop_2;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_561_i0)) && ((_561_i0).isLessThan(new BigNumber((_560_v0).length))))) {
          (_560_v0)[(_561_i0)] = (_561_i0).minus(_module.__default.safeModuloInt(new BigNumber(-943), (_this).f6));
        }
      }
      let _562_v1;
      _562_v1 = _module.D5.create_DC11(_dafny.Set.fromElements(_560_v0, _560_v0, _560_v0));
      let _563_v2;
      _563_v2 = _dafny.Set.fromElements(_560_v0, _560_v0, _560_v0);
      let _564_v3;
      _564_v3 = _dafny.Seq.of(true, ((_this).f6).isLessThan((_this).f9), (_563_v2).IsProperSubsetOf((_562_v1).dtor_cf19));
      _564_v3 = _564_v3;
      let _565_v4;
      _565_v4 = true;
      let _566_v5;
      _566_v5 = new _dafny.CodePoint('v'.codePointAt(0));
      let _567_v6;
      _567_v6 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(27),_566_v5);
      let _568_v7;
      _568_v7 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_565_v4);
      let _569_i1;
      _569_i1 = _dafny.ZERO;
      L1: {
        let _pat_let_tv15 = _565_v4;
        let _pat_let_tv16 = _564_v3;
        let _pat_let_tv17 = _564_v3;
        let _pat_let_tv18 = _565_v4;
        let _pat_let_tv19 = _565_v4;
        let _pat_let_tv20 = _565_v4;
        while (function (_source12) {
          if (_source12.is_DC1) {
            let _575___mcc_h0 = (_source12).cf1;
            let _576___mcc_h1 = (_source12).cf2;
            let _577___mcc_h2 = (_source12).cf3;
            let _578___mcc_h3 = (_source12).cf4;
            let _579_cf4 = _578___mcc_h3;
            let _580_cf3 = _577___mcc_h2;
            let _581_cf2 = _576___mcc_h1;
            let _582_cf1 = _575___mcc_h0;
            return _581_cf2;
          } else if (_source12.is_DC2) {
            let _583___mcc_h4 = (_source12).cf5;
            let _584___mcc_h5 = (_source12).cf6;
            let _585___mcc_h6 = (_source12).cf7;
            let _586___mcc_h7 = (_source12).cf8;
            let _587___mcc_h8 = (_source12).cf9;
            let _588_cf9 = _587___mcc_h8;
            let _589_cf8 = _586___mcc_h7;
            let _590_cf7 = _585___mcc_h6;
            let _591_cf6 = _584___mcc_h5;
            let _592_cf5 = _583___mcc_h4;
            return _pat_let_tv15;
          } else if (_source12.is_DC0) {
            let _593___mcc_h9 = (_source12).cf0;
            let _594_cf0 = _593___mcc_h9;
            return !_dafny.areEqual(_pat_let_tv16, _pat_let_tv17);
          } else {
            let _595___mcc_h10 = (_source12).cf10;
            let _596_cf10 = _595___mcc_h10;
            if (_pat_let_tv18) {
              return _pat_let_tv19;
            } else {
              return _pat_let_tv20;
            }
          }
        }(_module.__default.fm19(_565_v4, (((_567_v6).contains((_this).f6)) ? ((_567_v6).get((_this).f6)) : (_566_v5)), _568_v7, true, globalState))) {
          C1: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_569_i1)) {
              break L1;
            }
            _569_i1 = (_569_i1).plus(_dafny.ONE);
            let _570_v8;
            _570_v8 = _dafny.Map.Empty.slice().updateUnsafe(_565_v4,new BigNumber(969));
            let _571_v9;
            _571_v9 = _dafny.Seq.of(_570_v8);
            _571_v9 = _571_v9;
            let _572_v10;
            let _nw96 = new _module.C1();
            _nw96.__ctor();
            _572_v10 = _nw96;
            let _573_v11;
            _573_v11 = _module.D7.create_DC18(!(_565_v4), !(_565_v4), (_this).f6, true);
            (globalState).f2 = (_dafny.ZERO).minus((_573_v11).dtor_cf31);
            let _574_v12;
            let _nw97 = new _module.C1();
            _nw97.__ctor();
            _574_v12 = _nw97;
          }
        }
      }
      let _597_v13;
      let _nw98 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Seq.of());
      _597_v13 = _nw98;
      let _598_v14;
      _598_v14 = _dafny.Seq.UnicodeFromString("ftx");
      let _599_v15;
      _599_v15 = _dafny.Seq.of(_598_v14);
      let _index82 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_597_v13).length));
      (_597_v13)[_index82] = _dafny.Seq.update(_599_v15, _module.__default.safeIndex((_this).f9, new BigNumber((_599_v15).length)), _598_v14);
      let _index83 = _module.__default.safeIndex(new BigNumber(672), new BigNumber((_597_v13).length));
      (_597_v13)[_index83] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(359)), ((_600_v14) => function (_601_i2) {
        return _600_v14;
      })(_598_v14));
      let _602_v16;
      let _nw99 = new _module.C0();
      _nw99.__ctor(_565_v4);
      _602_v16 = _nw99;
      let _hi4 = ((_this).f9).plus((_this).f9);
      for (let _603_i3 = (_this).f9; _603_i3.isLessThan(_hi4); _603_i3 = _603_i3.plus(_dafny.ONE)) {
        _598_v14 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ubvcc"), _598_v14);
        r0 = (_this).f6;
        let _604_v17;
        let _init15 = ((_605_v5) => function (_606_i4) {
          return _dafny.Seq.Create(_module.__default.abs(new BigNumber(-871)), ((_607_v5) => function (_608_i5) {
            return _607_v5;
          })(_605_v5));
        })(_566_v5);
        let _nw100 = Array((new BigNumber(29)).toNumber());
        for (let _i0_15 = 0; _i0_15 < new BigNumber(_nw100.length); _i0_15++) {
          _nw100[_i0_15] = _init15(new BigNumber(_i0_15));
        }
        _604_v17 = _nw100;
        let _index84 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_604_v17).length));
        (_604_v17)[_index84] = _598_v14;
        let _index85 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_604_v17).length));
        (_604_v17)[_index85] = _dafny.Seq.Concat(_598_v14, _dafny.Seq.Create(_module.__default.abs(new BigNumber(190)), ((_609_v5) => function (_610_i6) {
          return _609_v5;
        })(_566_v5)));
        let _611_v18;
        _611_v18 = _dafny.Set.fromElements(_603_i3, _603_i3, _module.__default.safeDivisionInt(_603_i3, _603_i3));
        let _rhs69 = (((_602_v16).f13) ? (!((_602_v16).f13)) : ((_602_v16).f13));
        let _rhs70 = _611_v18;
        let _rhs71 = _560_v0;
        let _lhs54 = globalState;
        _lhs54.f1 = _rhs69;
        _611_v18 = _rhs70;
        _560_v0 = _rhs71;
      }
      r0 = ((_565_v4) ? ((_this).f6) : ((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.UnicodeFromString("jlx")).length))));
      let _612_v20;
      _612_v20 = _dafny.Set.fromElements(new _dafny.CodePoint('g'.codePointAt(0)), _566_v5);
      let _613_v21;
      _613_v21 = _dafny.Seq.of(_612_v20, _612_v20);
      r1 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Create(_module.__default.abs(new BigNumber(712)), ((_614_v5) => function (_615_i7) {
        return function () {
          let _coll38 = new _dafny.Set();
          for (const _compr_38 of (_dafny.Map.Empty.slice().updateUnsafe(_614_v5,(_this).f6)).Keys.Elements) {
            let _616_v19 = _compr_38;
            if ((_dafny.Map.Empty.slice().updateUnsafe(_614_v5,(_this).f6)).contains(_616_v19)) {
              _coll38.add(_616_v19);
            }
          }
          return _coll38;
        }();
      })(_566_v5)), _613_v21);
      let _617_v22;
      _617_v22 = _dafny.Set.fromElements(true);
      let _618_v23;
      _618_v23 = _dafny.Map.Empty.slice().updateUnsafe((_this).f9,_617_v22);
      r2 = (((_this).f9).isLessThanOrEqualTo(new BigNumber((_567_v6).length))) === (_module.__default.fm1(_module.D0.create_DC0(_618_v23), (_602_v16).f13, (_this).f9, (_602_v16).f13, globalState));
      return [r0, r1, r2];
    }
    get f9() {
      let _this = this;
      return _this._f9;
    };
    get f10() {
      let _this = this;
      return _this._f10;
    };
  };

  $module.C4 = class C4 {
    constructor () {
      this._tname = "_module.C4";
      this._f6 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm11(globalState) {
      let _this = this;
      return ((_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("tp"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(-574)), function (_619_i0) {
        return new _dafny.CodePoint('w'.codePointAt(0));
      }))).Intersect(_dafny.Set.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(622)), function (_620_i1) {
        return new _dafny.CodePoint('i'.codePointAt(0));
      })))).IsDisjointFrom((function () {
        let _coll39 = new _dafny.Set();
        for (const _compr_39 of (_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("wcshiaqcw"))).Elements) {
          let _621_v0 = _compr_39;
          if ((_dafny.MultiSet.fromElements(_dafny.Seq.UnicodeFromString("wcshiaqcw"))).contains(_621_v0)) {
            _coll39.add(_621_v0);
          }
        }
        return _coll39;
      }()).Union(_dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("qhbibqr"), _dafny.Seq.UnicodeFromString("jiutwlf"), _dafny.Seq.UnicodeFromString("tmcd"))));
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let _622_v0;
      let _nw101 = Array((new BigNumber(15)).toNumber()).fill(false);
      _622_v0 = _nw101;
      let _index86 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length));
      (_622_v0)[_index86] = false;
      let _index87 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length));
      (_622_v0)[_index87] = false;
      let _623_v1;
      _623_v1 = new _dafny.CodePoint('c'.codePointAt(0));
      let _624_v2;
      _624_v2 = _dafny.Seq.of(_623_v1, new _dafny.CodePoint('a'.codePointAt(0)));
      let _625_v3;
      let _nw102 = Array((new BigNumber(28)).toNumber());
      _nw102[(_dafny.ZERO).toNumber()] = (((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]) ? (_623_v1) : (_623_v1));
      _nw102[(_dafny.ONE).toNumber()] = (_624_v2)[_module.__default.safeIndex(new BigNumber(-956), new BigNumber((_624_v2).length))];
      _nw102[(new BigNumber(2)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(3)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(4)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(5)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(6)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(7)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(8)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(9)).toNumber()] = (_624_v2)[_module.__default.safeIndex(new BigNumber(953), new BigNumber((_624_v2).length))];
      _nw102[(new BigNumber(10)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(11)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(12)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(13)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(14)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(15)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(16)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(17)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(18)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(19)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(20)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(21)).toNumber()] = (((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]) ? (new _dafny.CodePoint('y'.codePointAt(0))) : (_623_v1));
      _nw102[(new BigNumber(22)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(23)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(24)).toNumber()] = _module.__default.fm12(globalState);
      _nw102[(new BigNumber(25)).toNumber()] = _623_v1;
      _nw102[(new BigNumber(26)).toNumber()] = new _dafny.CodePoint('s'.codePointAt(0));
      _nw102[(new BigNumber(27)).toNumber()] = _623_v1;
      _625_v3 = _nw102;
      let _626_v4;
      _626_v4 = _dafny.Set.fromElements((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]);
      let _627_v5;
      _627_v5 = _module.D0.create_DC1((_this).f6, ((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]) && ((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]), (_this).f6, (_626_v4).IsProperSubsetOf(_626_v4));
      let _628_v6;
      _628_v6 = _module.D3.create_DC7(_624_v2);
      let _pat_let_tv21 = _624_v2;
      let _629_v7;
      _629_v7 = _module.D0.create_DC2(new BigNumber(542), _623_v1, (_this).f6, (_this).f6, new BigNumber(((function (_pat_let18_0) {
  return function (_630_dt__update__tmp_h0) {
    return function (_pat_let19_0) {
      return function (_631_dt__update_hcf13_h0) {
        return _module.D3.create_DC7(_631_dt__update_hcf13_h0);
      }(_pat_let19_0);
    }(_pat_let_tv21);
  }(_pat_let18_0);
}(_628_v6)).dtor_cf13).length));
      let _rhs72 = (((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]) ? (_625_v3) : (_625_v3));
      let _rhs73 = (_629_v7).dtor_cf8;
      let _rhs74 = _module.__default.fm13(_623_v1, _624_v2, globalState);
      let _lhs55 = globalState;
      _625_v3 = _rhs72;
      _lhs55.f2 = _rhs73;
      _627_v5 = _rhs74;
      let _index88 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length));
      (_622_v0)[_index88] = (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))];
      let _hi5 = (_this).f6;
      for (let _632_i0 = (_this).f6; _632_i0.isLessThan(_hi5); _632_i0 = _632_i0.plus(_dafny.ONE)) {
        if ((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]) {
          let _633_v8;
          _633_v8 = _module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(_632_i0,_626_v4));
          r2 = !(!(_module.__default.fm1(_633_v8, (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))], _632_i0, !((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]), globalState)) || ((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))])) || (!((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]));
          let _634_v9;
          let _init16 = ((_635_i0) => function (_636_i1) {
            return (_636_i1).plus(_635_i0);
          })(_632_i0);
          let _nw103 = Array((new BigNumber(5)).toNumber());
          for (let _i0_16 = 0; _i0_16 < new BigNumber(_nw103.length); _i0_16++) {
            _nw103[_i0_16] = _init16(new BigNumber(_i0_16));
          }
          _634_v9 = _nw103;
          let _init17 = function (_637_i2) {
            return (_637_i2).minus((_this).f6);
          };
          let _nw104 = Array((new BigNumber(14)).toNumber());
          for (let _i0_17 = 0; _i0_17 < new BigNumber(_nw104.length); _i0_17++) {
            _nw104[_i0_17] = _init17(new BigNumber(_i0_17));
          }
          _634_v9 = _nw104;
          _623_v1 = (_629_v7).dtor_cf6;
          let _638_v11;
          let _init18 = ((_639_v0) => function (_640_i3) {
            return function () {
              let _coll40 = new _dafny.Set();
              for (const _compr_40 of (_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_639_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_639_v0).length))],(_639_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_639_v0).length))]))).Elements) {
                let _641_v10 = _compr_40;
                if (_dafny.Seq.contains(_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe((_639_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_639_v0).length))],(_639_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_639_v0).length))])), _641_v10)) {
                  _coll40.add(_641_v10);
                }
              }
              return _coll40;
            }();
          })(_622_v0);
          let _nw105 = Array((new BigNumber(22)).toNumber());
          for (let _i0_18 = 0; _i0_18 < new BigNumber(_nw105.length); _i0_18++) {
            _nw105[_i0_18] = _init18(new BigNumber(_i0_18));
          }
          _638_v11 = _nw105;
          let _642_v12;
          _642_v12 = _dafny.Map.Empty.slice().updateUnsafe((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))],(_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]);
          let _643_v13;
          _643_v13 = _dafny.Seq.of(_dafny.Set.fromElements(_642_v12, _642_v12), _dafny.Set.fromElements(_dafny.Map.Empty.slice().updateUnsafe((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))],true)));
          let _644_v14;
          _644_v14 = _dafny.MultiSet.fromElements((_this).f6);
          let _645_v15;
          _645_v15 = _dafny.Seq.of(new BigNumber((_644_v14).cardinality()), _632_i0);
          let _index89 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_638_v11).length));
          (_638_v11)[_index89] = ((_643_v13)[_module.__default.safeIndex((_this).f6, new BigNumber((_643_v13).length))]).Union(_dafny.Set.fromElements(_module.__default.fm14(new BigNumber(702), _632_i0, new BigNumber((_645_v15).length), globalState)));
          let _646_v16;
          _646_v16 = _dafny.Map.Empty.slice().updateUnsafe((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))],_642_v12);
          let _647_v17;
          _647_v17 = _dafny.Map.Empty.slice().updateUnsafe((((_646_v16).contains((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))])) ? ((_646_v16).get((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))])) : (_642_v12)),_632_i0);
          let _index90 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_638_v11).length));
          let _index91 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length));
          let _index92 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length));
          let _rhs75 = function () {
            let _coll41 = new _dafny.Set();
            for (const _compr_41 of ((_647_v17).Merge(_647_v17)).Keys.Elements) {
              let _648_v18 = _compr_41;
              if (((_647_v17).Merge(_647_v17)).contains(_648_v18)) {
                _coll41.add(_648_v18);
              }
            }
            return _coll41;
          }();
          let _rhs76 = (_this).fm11(globalState);
          let _rhs77 = (_627_v5).dtor_cf2;
          let _rhs78 = (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))];
          let _lhs56 = _638_v11;
          let _lhs57 = _module.__default.safeIndex(new BigNumber(825), new BigNumber((_638_v11).length));
          let _lhs58 = _622_v0;
          let _lhs59 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length));
          let _lhs60 = _622_v0;
          let _lhs61 = _module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length));
          let _lhs62 = globalState;
          _lhs56[_lhs57] = _rhs75;
          _lhs58[_lhs59] = _rhs76;
          _lhs60[_lhs61] = _rhs77;
          _lhs62.f1 = _rhs78;
          let _649_v19;
          let _nw106 = Array((new BigNumber(22)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _649_v19 = _nw106;
          let _index93 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_649_v19).length));
          (_649_v19)[_index93] = _624_v2;
          let _index94 = _module.__default.safeIndex(new BigNumber(860), new BigNumber((_649_v19).length));
          (_649_v19)[_index94] = _dafny.Seq.Concat(_624_v2, _624_v2);
        } else {
          let _650_v20;
          _650_v20 = _dafny.Seq.of(false, (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]);
          let _651_v21;
          _651_v21 = _dafny.Seq.of(_module.__default.fm15(_632_i0, true, globalState), new BigNumber((_650_v20).length), (_this).f6, _632_i0, (_this).f6);
          let _652_v22;
          let _653_v23;
          let _654_v24;
          let _655_v25;
          let _out0;
          let _out1;
          let _out2;
          let _out3;
          let _outcollector0 = (_this).m3(new BigNumber((_651_v21).length), (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))], _632_i0, (_632_i0).isEqualTo(new BigNumber(231)), globalState);
          _out0 = _outcollector0[0];
          _out1 = _outcollector0[1];
          _out2 = _outcollector0[2];
          _out3 = _outcollector0[3];
          _652_v22 = _out0;
          _653_v23 = _out1;
          _654_v24 = _out2;
          _655_v25 = _out3;
          let _656_v26;
          _656_v26 = _dafny.Map.Empty.slice().updateUnsafe(_654_v24,(_this).f6);
          let _657_v27;
          _657_v27 = _dafny.Set.fromElements((((_656_v26).contains(new BigNumber(-238))) ? ((_656_v26).get(new BigNumber(-238))) : (new BigNumber((_651_v21).length))));
          let _rhs79 = new BigNumber(((((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]) ? ((_657_v27).Union(_dafny.Set.fromElements(new BigNumber(562)))) : (_dafny.Set.fromElements((_this).f6, _654_v24)))).length);
          let _rhs80 = (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))];
          let _rhs81 = (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))];
          let _rhs82 = new BigNumber(370);
          let _lhs63 = globalState;
          let _lhs64 = globalState;
          let _lhs65 = globalState;
          r0 = _rhs79;
          _lhs63.f1 = _rhs80;
          _lhs64.f1 = _rhs81;
          _lhs65.f2 = _rhs82;
          let _658_v28;
          _658_v28 = _dafny.MultiSet.fromElements((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))], (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]);
          let _659_v29;
          let _nw107 = Array((new BigNumber(12)).toNumber());
          _nw107[(_dafny.ZERO).toNumber()] = (_dafny.ZERO).minus(_654_v24);
          _nw107[(_dafny.ONE).toNumber()] = _module.__default.fm15(_654_v24, (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))], globalState);
          _nw107[(new BigNumber(2)).toNumber()] = new BigNumber((_657_v27).length);
          _nw107[(new BigNumber(3)).toNumber()] = (_this).f6;
          _nw107[(new BigNumber(4)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_651_v21, _651_v21)).length);
          _nw107[(new BigNumber(5)).toNumber()] = (((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]) ? (_652_v22) : ((_this).f6));
          _nw107[(new BigNumber(6)).toNumber()] = _module.__default.safeModuloInt(_654_v24, _632_i0);
          _nw107[(new BigNumber(7)).toNumber()] = (((_658_v28).contains((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))])) ? ((_658_v28).get((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))])) : (new BigNumber(-667)));
          _nw107[(new BigNumber(8)).toNumber()] = new BigNumber(962);
          _nw107[(new BigNumber(9)).toNumber()] = _654_v24;
          _nw107[(new BigNumber(10)).toNumber()] = _652_v22;
          _nw107[(new BigNumber(11)).toNumber()] = (_dafny.ZERO).minus(_632_i0);
          _659_v29 = _nw107;
          let _index95 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_659_v29).length));
          (_659_v29)[_index95] = _module.__default.safeModuloInt((_this).f6, _632_i0);
          let _index96 = _module.__default.safeIndex(new BigNumber(220), new BigNumber((_659_v29).length));
          (_659_v29)[_index96] = _module.__default.safeDivisionInt((_652_v22).minus((_this).f6), (_dafny.ZERO).minus((_627_v5).dtor_cf1));
          _623_v1 = (_624_v2)[_module.__default.safeIndex(_632_i0, new BigNumber((_624_v2).length))];
          (globalState).f5 = (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))];
        }
        let _660_v30;
        _660_v30 = _dafny.Map.Empty.slice().updateUnsafe((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))],(_this).f6);
        let _661_v31;
        _661_v31 = _dafny.Map.Empty.slice().updateUnsafe(!((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]),(_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]);
        let _662_v32;
        _662_v32 = _661_v31;
        _660_v30 = (_660_v30).update((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))], (_dafny.ZERO).minus(new BigNumber(((_661_v31)).length)));
        _623_v1 = _623_v1;
        let _663_v33;
        _663_v33 = _dafny.Map.Empty.slice().updateUnsafe(_627_v5,new BigNumber(-138));
        _663_v33 = (_663_v33).update(_627_v5, (_629_v7).dtor_cf7);
      }
      let _664_v34;
      _664_v34 = _dafny.Map.Empty.slice().updateUnsafe((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))],(_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]);
      let _665_v35;
      _665_v35 = _dafny.Map.Empty.slice().updateUnsafe(_664_v34,(_this).fm11(globalState));
      _665_v35 = (_665_v35).update((_664_v34).update((_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))], (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]), (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]);
      let _666_v36;
      let _nw108 = new _module.C2();
      _nw108.__ctor(_dafny.Seq.UnicodeFromString("iyanqn"), (_622_v0)[_module.__default.safeIndex(new BigNumber(880), new BigNumber((_622_v0).length))]);
      _666_v36 = _nw108;
      r0 = (_this).f6;
      r1 = true;
      let _667_v37;
      _667_v37 = _dafny.MultiSet.fromElements(_666_v36.f12);
      let _668_v38;
      _668_v38 = _module.D7.create_DC19(_667_v37, (_this).f6);
      r2 = (_667_v37).equals((_668_v38).dtor_cf33);
      return [r0, r1, r2];
    }
    m3(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _module.D1.Default();
      let r2 = _dafny.ZERO;
      let r3 = _dafny.Set.Empty;
      (globalState).f5 = p3;
      let _hi6 = (_this).f6;
      for (let _669_i0 = p0; _669_i0.isLessThan(_hi6); _669_i0 = _669_i0.plus(_dafny.ONE)) {
        let _670_v0;
        let _nw109 = Array((new BigNumber(14)).toNumber()).fill(_dafny.ZERO);
        _670_v0 = _nw109;
        let _671_v1;
        _671_v1 = _dafny.Seq.UnicodeFromString("gpng");
        let _index97 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_670_v0).length));
        (_670_v0)[_index97] = _module.__default.safeDivisionInt((_this).f6, new BigNumber((_671_v1).length));
        let _index98 = _module.__default.safeIndex(new BigNumber(959), new BigNumber((_670_v0).length));
        (_670_v0)[_index98] = (_dafny.ZERO).minus(_669_i0);
        let _672_v2;
        _672_v2 = new _dafny.CodePoint('p'.codePointAt(0));
        let _673_v3;
        _673_v3 = _dafny.Map.Empty.slice().updateUnsafe(_672_v2,new BigNumber((_dafny.Set.fromElements(p3)).length));
        let _674_v4;
        _674_v4 = _dafny.Seq.of(p3, false);
        let _675_v5;
        _675_v5 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_674_v4).length),new BigNumber(-529));
        let _676_v6;
        _676_v6 = _module.D7.create_DC18(p3, p1, new BigNumber((_675_v5).length), true);
        let _677_v7;
        _677_v7 = _module.D3.create_DC7(_671_v1);
        let _678_v8;
        _678_v8 = _module.D5.create_DC13((_676_v6).dtor_cf30, p3, (_677_v7).dtor_cf13, p0);
        let _679_v9;
        let _nw110 = new _module.C3();
        _nw110.__ctor(new BigNumber((_dafny.Seq.UnicodeFromString("poee")).length), _module.__default.fm25(p2, _673_v3, _672_v2, globalState), (_678_v8).dtor_cf24);
        _679_v9 = _nw110;
        let _680_v10;
        _680_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(688),_674_v4);
        _680_v10 = (_680_v10).update((_670_v0)[_module.__default.safeIndex(new BigNumber(959), new BigNumber((_670_v0).length))], _674_v4);
        let _681_v11;
        _681_v11 = _dafny.MultiSet.fromElements(p3, p3, p1, p3);
        let _682_v12;
        _682_v12 = _dafny.Seq.of(_681_v11);
        let _683_v13;
        let _nw111 = Array((new BigNumber(11)).toNumber());
        _nw111[(_dafny.ZERO).toNumber()] = true;
        _nw111[(_dafny.ONE).toNumber()] = ((_682_v12)[_module.__default.safeIndex((_this).f6, new BigNumber((_682_v12).length))]).IsDisjointFrom(_dafny.MultiSet.FromArray(_674_v4));
        _nw111[(new BigNumber(2)).toNumber()] = p3;
        _nw111[(new BigNumber(3)).toNumber()] = p1;
        _nw111[(new BigNumber(4)).toNumber()] = p1;
        _nw111[(new BigNumber(5)).toNumber()] = (p1) || (p3);
        _nw111[(new BigNumber(6)).toNumber()] = p1;
        _nw111[(new BigNumber(7)).toNumber()] = p3;
        _nw111[(new BigNumber(8)).toNumber()] = p3;
        _nw111[(new BigNumber(9)).toNumber()] = ((p1) ? (!(false)) : (p3));
        _nw111[(new BigNumber(10)).toNumber()] = _dafny.areEqual(_671_v1, _671_v1);
        _683_v13 = _nw111;
        let _index99 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_683_v13).length));
        (_683_v13)[_index99] = p1;
        let _index100 = _module.__default.safeIndex(new BigNumber(749), new BigNumber((_683_v13).length));
        (_683_v13)[_index100] = p1;
      }
      if (true) {
        (globalState).f5 = p1;
        let _684_v14;
        let _nw112 = Array((new BigNumber(22)).toNumber()).fill(false);
        _684_v14 = _nw112;
        let _685_v15;
        _685_v15 = _dafny.Seq.of(_684_v14, _684_v14, _684_v14);
        _685_v15 = _dafny.Seq.Concat(_685_v15, _685_v15);
        let _686_v16;
        _686_v16 = _module.D0.create_DC1(_module.__default.safeModuloInt(p0, new BigNumber(251)), p1, new BigNumber((_dafny.Seq.UnicodeFromString("fr")).length), p1);
        let _pat_let_tv22 = p0;
        _686_v16 = function (_pat_let20_0) {
          return function (_687_dt__update__tmp_h0) {
            return function (_pat_let21_0) {
              return function (_688_dt__update_hcf1_h0) {
                return function (_pat_let22_0) {
                  return function (_689_dt__update_hcf4_h0) {
                    return function (_pat_let23_0) {
                      return function (_690_dt__update_hcf3_h0) {
                        return _module.D0.create_DC1(_688_dt__update_hcf1_h0, (_687_dt__update__tmp_h0).dtor_cf2, _690_dt__update_hcf3_h0, _689_dt__update_hcf4_h0);
                      }(_pat_let23_0);
                    }((_this).f6);
                  }(_pat_let22_0);
                }(true);
              }(_pat_let21_0);
            }(_pat_let_tv22);
          }(_pat_let20_0);
        }(_686_v16);
        let _691_v17;
        let _nw113 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _691_v17 = _nw113;
        let _index101 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_691_v17).length));
        (_691_v17)[_index101] = (new BigNumber(-761)).multipliedBy((_this).f6);
        let _692_v18;
        _692_v18 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(p2),_684_v14);
        let _index102 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_691_v17).length));
        let _rhs83 = _module.__default.safeDivisionInt(new BigNumber(362), (p2).minus(new BigNumber(667)));
        let _rhs84 = (((_692_v18).contains((_this).f6)) ? ((_692_v18).get((_this).f6)) : (_684_v14));
        let _rhs85 = _684_v14;
        let _rhs86 = (_this).f6;
        let _lhs66 = _691_v17;
        let _lhs67 = _module.__default.safeIndex(new BigNumber(410), new BigNumber((_691_v17).length));
        let _lhs68 = globalState;
        _lhs66[_lhs67] = _rhs83;
        _684_v14 = _rhs84;
        _684_v14 = _rhs85;
        _lhs68.f2 = _rhs86;
        let _693_v19;
        _693_v19 = _dafny.Seq.of(p0);
        let _694_v20;
        _694_v20 = _dafny.Map.Empty.slice().updateUnsafe(p3,p3);
        let _695_v21;
        let _nw114 = new _module.C3();
        _nw114.__ctor(p0, _693_v19, (new BigNumber(((_module.D1.create_DC4(_693_v19)).dtor_cf11).length)).plus(new BigNumber((_694_v20).length)));
        _695_v21 = _nw114;
      } else {
        let _696_v22;
        _696_v22 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        if (!((_696_v22).Merge(_696_v22)).equals(_696_v22)) {
          let _697_v23;
          _697_v23 = _dafny.Seq.of(p0);
          let _698_v24;
          _698_v24 = _dafny.Seq.UnicodeFromString("aofqlqta");
          let _699_v25;
          let _nw115 = new _module.C3();
          _nw115.__ctor((_this).f6, _697_v23, new BigNumber((_698_v24).length));
          _699_v25 = _nw115;
          _699_v25 = _699_v25;
          let _700_v26;
          _700_v26 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(new BigNumber((_696_v22).length), new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(256)), ((_701_p0, _702_p3) => function (_703_i1) {
            return _dafny.Map.Empty.slice().updateUnsafe(_701_p0,_702_p3);
          })(p0, p3))).length))).length), (_699_v25).f6),p0);
          _700_v26 = (_700_v26).Merge(_700_v26);
          let _704_v27;
          _704_v27 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeDivisionInt(p0, (_699_v25).f6),!(p3) || (p3));
          (globalState).f5 = (((_704_v27).contains((_dafny.ZERO).minus(new BigNumber((((!(p3)) ? (_698_v24) : (_698_v24))).length)))) ? ((_704_v27).get((_dafny.ZERO).minus(new BigNumber((((!(p3)) ? (_698_v24) : (_698_v24))).length)))) : ((new BigNumber((_698_v24).length)).isLessThan(p2)));
          let _705_v28;
          _705_v28 = _dafny.Set.fromElements(_dafny.Seq.UnicodeFromString("cdiwerm"));
          let _706_v29;
          _706_v29 = _dafny.Seq.of(_698_v24, _698_v24);
          let _707_v30;
          _707_v30 = _dafny.Set.fromElements(new BigNumber(-403), (_this).f6, (new BigNumber((_705_v28).length)).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber((_706_v29).length), _module.__default.fm24((_699_v25).f6, globalState))).length)));
          _707_v30 = _707_v30;
          (globalState).f5 = p1;
        } else {
          let _708_v31;
          let _nw116 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          _708_v31 = _nw116;
          let _709_v32;
          _709_v32 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of(_module.__default.fm24(new BigNumber(811), globalState)),_708_v31);
          let _710_v33;
          _710_v33 = _dafny.Seq.UnicodeFromString("dmvccnxx");
          let _711_v34;
          _711_v34 = _dafny.Seq.of((_this).f6, new BigNumber((_710_v33).length));
          let _712_v35;
          _712_v35 = _module.D8.create_DC20(_dafny.Map.Empty.slice().updateUnsafe(_711_v34,_708_v31));
          let _713_v36;
          _713_v36 = _module.D1.create_DC4(_711_v34);
          let _714_v37;
          let _nw117 = Array((new BigNumber(14)).toNumber()).fill([]);
          _714_v37 = _nw117;
          let _715_v38;
          _715_v38 = _dafny.Map.Empty.slice().updateUnsafe(_714_v37,_709_v32);
          let _pat_let_tv23 = p2;
          let _pat_let_tv24 = _708_v31;
          let _716_v39;
          let _nw118 = Array((new BigNumber(24)).toNumber());
          _nw118[(_dafny.ZERO).toNumber()] = _709_v32;
          _nw118[(_dafny.ONE).toNumber()] = _709_v32;
          _nw118[(new BigNumber(2)).toNumber()] = (_709_v32).Merge(_709_v32);
          _nw118[(new BigNumber(3)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(4)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f6),_708_v31);
          _nw118[(new BigNumber(5)).toNumber()] = (function (_pat_let24_0) {
            return function (_717_dt__update__tmp_h1) {
              return function (_pat_let25_0) {
                return function (_720_dt__update_hcf35_h0) {
                  return _module.D8.create_DC20(_720_dt__update_hcf35_h0);
                }(_pat_let25_0);
              }(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-35)), ((_718_p2) => function (_719_i2) {
                return _718_p2;
              })(_pat_let_tv23)),_pat_let_tv24));
            }(_pat_let24_0);
          }(_712_v35)).dtor_cf35;
          _nw118[(new BigNumber(6)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(7)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(8)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(9)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(10)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(11)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((_713_v36).dtor_cf11, _module.__default.safeIndex(p2, new BigNumber(((_713_v36).dtor_cf11).length)), p0),_708_v31);
          _nw118[(new BigNumber(12)).toNumber()] = (((_715_v38).contains(_714_v37)) ? ((_715_v38).get(_714_v37)) : (_dafny.Map.Empty.slice().updateUnsafe(_711_v34,_708_v31)));
          _nw118[(new BigNumber(13)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(14)).toNumber()] = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_this).f6),_708_v31);
          _nw118[(new BigNumber(15)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(16)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(17)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(18)).toNumber()] = (_709_v32).Merge(_709_v32);
          _nw118[(new BigNumber(19)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(20)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(21)).toNumber()] = (_712_v35).dtor_cf35;
          _nw118[(new BigNumber(22)).toNumber()] = _709_v32;
          _nw118[(new BigNumber(23)).toNumber()] = _709_v32;
          _716_v39 = _nw118;
          let _index103 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_716_v39).length));
          (_716_v39)[_index103] = _709_v32;
          let _721_v40;
          _721_v40 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm12(globalState),_709_v32);
          let _722_v41;
          _722_v41 = new _dafny.CodePoint('v'.codePointAt(0));
          let _index104 = _module.__default.safeIndex(new BigNumber(778), new BigNumber((_716_v39).length));
          (_716_v39)[_index104] = ((((_721_v40).contains(_722_v41)) ? ((_721_v40).get(_722_v41)) : (_709_v32))).update(_dafny.Seq.of((_this).f6), _708_v31);
          let _723_v42;
          _723_v42 = _dafny.Seq.of(p3);
          (globalState).f0 = _module.__default.safeModuloInt(new BigNumber((((true) ? (_723_v42) : (_dafny.Seq.of(p1)))).length), new BigNumber(435));
          (globalState).f1 = p3;
          (globalState).f0 = p0;
          r0 = p2;
        }
        r2 = p0;
        let _724_v43;
        _724_v43 = _dafny.MultiSet.fromElements(p1, p3);
        let _725_v44;
        _725_v44 = _dafny.Seq.UnicodeFromString("prilxjot");
        let _726_v45;
        _726_v45 = _dafny.Map.Empty.slice().updateUnsafe(_725_v44,p2);
        _724_v43 = _module.__default.fm30(_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xadt"), _725_v44), _726_v45, p3, globalState);
        let _727_v46;
        let _init19 = function (_728_i3) {
          return _module.__default.safeDivisionInt(_728_i3, (_this).f6);
        };
        let _nw119 = Array((new BigNumber(13)).toNumber());
        for (let _i0_19 = 0; _i0_19 < new BigNumber(_nw119.length); _i0_19++) {
          _nw119[_i0_19] = _init19(new BigNumber(_i0_19));
        }
        _727_v46 = _nw119;
        let _index105 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_727_v46).length));
        (_727_v46)[_index105] = (_this).f6;
        let _729_v47;
        _729_v47 = _dafny.MultiSet.fromElements((_this).f6);
        let _730_v48;
        _730_v48 = _dafny.Map.Empty.slice().updateUnsafe(p0,p3);
        let _731_v49;
        _731_v49 = _dafny.Map.Empty.slice().updateUnsafe((((_729_v47).contains(new BigNumber((_730_v48).length))) ? ((_729_v47).get(new BigNumber((_730_v48).length))) : ((_dafny.ZERO).minus(p0))),(_this).f6);
        let _index106 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_727_v46).length));
        (_727_v46)[_index106] = (((_731_v49).contains((_this).f6)) ? ((_731_v49).get((_this).f6)) : (p2));
        if (!(!(p1) || (p3))) {
          let _732_v50;
          let _nw120 = new _module.C1();
          _nw120.__ctor();
          _732_v50 = _nw120;
          let _733_v51;
          _733_v51 = _dafny.MultiSet.fromElements(_732_v50);
          let _734_v52;
          _734_v52 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(((_733_v51)).cardinality()),_730_v48);
          let _735_v53;
          _735_v53 = _module.D5.create_DC13(p1, !(p3), _dafny.Seq.Create(_module.__default.abs(new BigNumber(478)), function (_736_i4) {
  return new _dafny.CodePoint('s'.codePointAt(0));
}), (_727_v46)[_module.__default.safeIndex(new BigNumber(823), new BigNumber((_727_v46).length))]);
          let _737_v54;
          _737_v54 = _dafny.Seq.of(_735_v53);
          let _738_v56;
          _738_v56 = _dafny.Seq.of(p2);
          let _739_v57;
          _739_v57 = _dafny.Seq.of(new BigNumber((_738_v56).length), new BigNumber(567));
          _734_v52 = (_734_v52).update(new BigNumber((_dafny.Set.fromElements(_dafny.Seq.of(_735_v53, _module.D5.create_DC13(p1, p1, _725_v44, (_this).f6)), _737_v54)).length), ((p1) ? (function () {
            let _coll42 = new _dafny.Map();
            for (const _compr_42 of (_739_v57).Elements) {
              let _740_v55 = _compr_42;
              if (_dafny.Seq.contains(_739_v57, _740_v55)) {
                _coll42.push([(_740_v55).multipliedBy(p2),false]);
              }
            }
            return _coll42;
          }()) : (function () {
            let _coll43 = new _dafny.Map();
            for (const _compr_43 of _dafny.IntegerRange(new BigNumber(476), new BigNumber(297))) {
              let _741_v58 = _compr_43;
              if (((new BigNumber(476)).isLessThanOrEqualTo(_741_v58)) && ((_741_v58).isLessThan(new BigNumber(297)))) {
                _coll43.push([(_741_v58).minus(p0),p1]);
              }
            }
            return _coll43;
          }())));
          let _742_v59;
          _742_v59 = _dafny.Set.fromElements(p1);
          _731_v49 = (_731_v49).update(new BigNumber((_742_v59).length), p0);
          let _743_v60;
          _743_v60 = _module.D6.create_DC15(((p3) ? (_742_v59) : (_dafny.Set.fromElements(true, p3))));
          let _744_v61;
          _744_v61 = _dafny.Seq.of(p3);
          _743_v60 = _module.__default.fm31((_dafny.ZERO).minus((_727_v46)[_module.__default.safeIndex(new BigNumber(823), new BigNumber((_727_v46).length))]), p1, p3, _dafny.MultiSet.fromElements(_744_v61, _744_v61, _744_v61, _744_v61), globalState);
          let _745_v62;
          let _nw121 = Array((new BigNumber(27)).toNumber()).fill(false);
          _745_v62 = _nw121;
          let _index107 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_745_v62).length));
          (_745_v62)[_index107] = p1;
          let _index108 = _module.__default.safeIndex(new BigNumber(166), new BigNumber((_745_v62).length));
          (_745_v62)[_index108] = p1;
          let _746_v63;
          let _init20 = ((_747_v53) => function (_748_i5) {
            return _747_v53;
          })(_735_v53);
          let _nw122 = Array((new BigNumber(9)).toNumber());
          for (let _i0_20 = 0; _i0_20 < new BigNumber(_nw122.length); _i0_20++) {
            _nw122[_i0_20] = _init20(new BigNumber(_i0_20));
          }
          _746_v63 = _nw122;
          let _index109 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_746_v63).length));
          (_746_v63)[_index109] = _735_v53;
          let _749_v64;
          _749_v64 = new _dafny.CodePoint('r'.codePointAt(0));
          let _pat_let_tv25 = _725_v44;
          let _pat_let_tv26 = _725_v44;
          let _pat_let_tv27 = _749_v64;
          let _index110 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_746_v63).length));
          let _rhs87 = ((p1) ? (p2) : (new BigNumber((_dafny.Set.fromElements((_745_v62)[_module.__default.safeIndex(new BigNumber(166), new BigNumber((_745_v62).length))], !(p1))).length)));
          let _rhs88 = (_745_v62)[_module.__default.safeIndex(new BigNumber(166), new BigNumber((_745_v62).length))];
          let _rhs89 = function (_pat_let26_0) {
            return function (_750_dt__update__tmp_h2) {
              return function (_pat_let27_0) {
                return function (_751_dt__update_hcf23_h0) {
                  return _module.D5.create_DC13((_750_dt__update__tmp_h2).dtor_cf21, (_750_dt__update__tmp_h2).dtor_cf22, _751_dt__update_hcf23_h0, (_750_dt__update__tmp_h2).dtor_cf24);
                }(_pat_let27_0);
              }(_dafny.Seq.update(_pat_let_tv25, _module.__default.safeIndex(new BigNumber(211), new BigNumber((_pat_let_tv26).length)), _pat_let_tv27));
            }(_pat_let26_0);
          }(_735_v53);
          let _lhs69 = globalState;
          let _lhs70 = globalState;
          let _lhs71 = _746_v63;
          let _lhs72 = _module.__default.safeIndex(new BigNumber(35), new BigNumber((_746_v63).length));
          _lhs69.f0 = _rhs87;
          _lhs70.f5 = _rhs88;
          _lhs71[_lhs72] = _rhs89;
        } else {
          _725_v44 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(533)), function (_752_i6) {
            return new _dafny.CodePoint('m'.codePointAt(0));
          });
          let _753_v65;
          _753_v65 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm15(p0, p1, globalState),_dafny.Set.fromElements(p3, !(p3), p1, p3));
          let _754_v66;
          _754_v66 = _module.D0.create_DC0(_753_v65);
          (globalState).f5 = _module.__default.fm1(_754_v66, ((p3) ? (p1) : (!(p1))), p0, ((false) ? (p1) : (!(p3))), globalState);
          let _index111 = _module.__default.safeIndex(new BigNumber(823), new BigNumber((_727_v46).length));
          (_727_v46)[_index111] = _module.__default.safeDivisionInt((_this).f6, (new BigNumber(352)).minus((_727_v46)[_module.__default.safeIndex(new BigNumber(823), new BigNumber((_727_v46).length))]));
          let _755_v67;
          _755_v67 = new _dafny.CodePoint('w'.codePointAt(0));
          let _756_v68;
          _756_v68 = _module.D0.create_DC2(new BigNumber(5), _755_v67, (_this).f6, p0, p2);
          let _757_v69;
          _757_v69 = _module.D0.create_DC3(_756_v68);
          _757_v69 = _module.__default.fm32(globalState);
          let _758_v70;
          _758_v70 = _dafny.Seq.of((_727_v46)[_module.__default.safeIndex(new BigNumber(823), new BigNumber((_727_v46).length))]);
          _758_v70 = (_module.__default.fm33(globalState)).dtor_cf11;
        }
      }
      let _759_v71;
      _759_v71 = _dafny.Seq.UnicodeFromString("tvo");
      _759_v71 = _759_v71;
      let _hi7 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(421)), function (_760_i8) {
        return new _dafny.CodePoint('v'.codePointAt(0));
      })).length)).plus(p0);
      for (let _761_i7 = _module.__default.fm24((_this).f6, globalState); _761_i7.isLessThan(_hi7); _761_i7 = _761_i7.plus(_dafny.ONE)) {
        let _762_v72;
        _762_v72 = _dafny.Seq.of(p2);
        let _763_v73;
        let _nw123 = Array((new BigNumber(14)).toNumber());
        _nw123[(_dafny.ZERO).toNumber()] = _761_i7;
        _nw123[(_dafny.ONE).toNumber()] = (_dafny.ZERO).minus(((_dafny.ZERO).minus(p0)).multipliedBy((_762_v72)[_module.__default.safeIndex(p0, new BigNumber((_762_v72).length))]));
        _nw123[(new BigNumber(2)).toNumber()] = _761_i7;
        _nw123[(new BigNumber(3)).toNumber()] = (_this).f6;
        _nw123[(new BigNumber(4)).toNumber()] = p0;
        _nw123[(new BigNumber(5)).toNumber()] = (_762_v72)[_module.__default.safeIndex(p0, new BigNumber((_762_v72).length))];
        _nw123[(new BigNumber(6)).toNumber()] = (_this).f6;
        _nw123[(new BigNumber(7)).toNumber()] = (_dafny.ZERO).minus(new BigNumber((_module.__default.fm2(_759_v71, globalState)).length));
        _nw123[(new BigNumber(8)).toNumber()] = p2;
        _nw123[(new BigNumber(9)).toNumber()] = new BigNumber((_dafny.Seq.UnicodeFromString("qqajrefo")).length);
        _nw123[(new BigNumber(10)).toNumber()] = new BigNumber((_759_v71).length);
        _nw123[(new BigNumber(11)).toNumber()] = new BigNumber((_759_v71).length);
        _nw123[(new BigNumber(12)).toNumber()] = _module.__default.safeModuloInt(_module.__default.fm15(p2, p1, globalState), (_dafny.ZERO).minus(_761_i7));
        _nw123[(new BigNumber(13)).toNumber()] = new BigNumber(898);
        _763_v73 = _nw123;
        let _rhs90 = p2;
        let _rhs91 = p2;
        let _rhs92 = _763_v73;
        let _lhs73 = globalState;
        let _lhs74 = globalState;
        _lhs73.f2 = _rhs90;
        _lhs74.f2 = _rhs91;
        _763_v73 = _rhs92;
        let _764_v74;
        let _nw124 = Array((new BigNumber(3)).toNumber()).fill(_dafny.Map.Empty);
        _764_v74 = _nw124;
        let _765_v75;
        _765_v75 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Set.fromElements(p0, _761_i7, _761_i7, _761_i7, (_dafny.ZERO).minus(p0)),p1);
        let _index112 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_764_v74).length));
        (_764_v74)[_index112] = _765_v75;
        let _766_v76;
        let _nw125 = Array((new BigNumber(12)).toNumber()).fill(false);
        _766_v76 = _nw125;
        let _767_v77;
        _767_v77 = _dafny.Set.fromElements(_766_v76);
        let _768_v78;
        _768_v78 = _module.D0.create_DC1(new BigNumber(117), p3, (_this).f6, p1);
        let _index113 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_764_v74).length));
        let _rhs93 = (_767_v77).IsSubsetOf(_767_v77);
        let _rhs94 = _765_v75;
        let _rhs95 = (_768_v78).dtor_cf2;
        let _rhs96 = (_dafny.ZERO).minus(p0);
        let _rhs97 = p1;
        let _lhs75 = globalState;
        let _lhs76 = _764_v74;
        let _lhs77 = _module.__default.safeIndex(new BigNumber(539), new BigNumber((_764_v74).length));
        let _lhs78 = globalState;
        let _lhs79 = globalState;
        _lhs75.f5 = _rhs93;
        _lhs76[_lhs77] = _rhs94;
        _lhs78.f5 = _rhs95;
        r0 = _rhs96;
        _lhs79.f1 = _rhs97;
        let _769_v79;
        let _nw126 = new _module.C3();
        _nw126.__ctor(p2, _762_v72, (_this).f6);
        _769_v79 = _nw126;
        let _770_v80;
        _770_v80 = _dafny.MultiSet.fromElements(_769_v79);
        let _rhs98 = new BigNumber(((_770_v80).update(_769_v79, _module.__default.abs((_dafny.ZERO).minus(_761_i7)))).cardinality());
        let _rhs99 = (_module.__default.fm24((_this).f6, globalState)).minus(p2);
        let _rhs100 = new BigNumber(600);
        let _rhs101 = (_769_v79).f6;
        let _lhs80 = globalState;
        let _lhs81 = globalState;
        let _lhs82 = globalState;
        _lhs80.f2 = _rhs98;
        _lhs81.f0 = _rhs99;
        _lhs82.f2 = _rhs100;
        r2 = _rhs101;
        let _771_v81;
        _771_v81 = new _dafny.CodePoint('o'.codePointAt(0));
        let _772_v82;
        _772_v82 = _dafny.Set.fromElements(_771_v81, _771_v81);
        let _773_v83;
        _773_v83 = _dafny.Map.Empty.slice().updateUnsafe(p1,p0);
        let _774_v84;
        _774_v84 = _dafny.Set.fromElements(new BigNumber((_772_v82).length), new BigNumber((_759_v71).length), p0, new BigNumber((_773_v83).length));
        if ((new BigNumber((_dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("l"), _dafny.Seq.UnicodeFromString("ngk"))).length)).isLessThanOrEqualTo(new BigNumber((_774_v84).length))) {
          let _775_v85;
          _775_v85 = _dafny.Map.Empty.slice().updateUnsafe(p2,p1);
          let _776_v86;
          _776_v86 = _dafny.Map.Empty.slice().updateUnsafe(_775_v85,_766_v76);
          let _777_v87;
          _777_v87 = _dafny.Map.Empty.slice().updateUnsafe((_769_v79).f6,(_769_v79).f6);
          let _778_v88;
          let _nw127 = new _module.C3();
          _nw127.__ctor((_dafny.ZERO).minus((_769_v79).f6), _dafny.Seq.of(p0, _761_i7, (_769_v79).f6), (((_777_v87).contains(p2)) ? ((_777_v87).get(p2)) : (p0)));
          _778_v88 = _nw127;
          let _index114 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_763_v73).length));
          (_763_v73)[_index114] = new BigNumber((_773_v83).length);
          let _index115 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_763_v73).length));
          let _rhs102 = (_776_v86).Merge(_776_v86);
          let _rhs103 = _763_v73;
          let _rhs104 = _778_v88;
          let _rhs105 = _module.__default.safeModuloInt(((_dafny.ZERO).minus(new BigNumber((_775_v85).length))).multipliedBy((_769_v79).f6), ((_this).f6).minus(p0));
          let _rhs106 = _module.__default.safeModuloInt((_dafny.ZERO).minus(p2), ((_778_v88).f10)[_module.__default.safeIndex((_769_v79).f6, new BigNumber(((_778_v88).f10).length))]);
          let _lhs83 = _763_v73;
          let _lhs84 = _module.__default.safeIndex(new BigNumber(895), new BigNumber((_763_v73).length));
          _776_v86 = _rhs102;
          _763_v73 = _rhs103;
          _778_v88 = _rhs104;
          _lhs83[_lhs84] = _rhs105;
          r2 = _rhs106;
          let _779_v89;
          _779_v89 = _dafny.MultiSet.fromElements(p3, p3);
          (globalState).f2 = _module.__default.safeModuloInt((_dafny.ZERO).minus(_761_i7), (new BigNumber((_779_v89).cardinality())).minus(_761_i7));
          let _780_v90;
          _780_v90 = _module.D8.create_DC21(true, p3, (_769_v79).f6);
          let _781_v91;
          let _nw128 = new _module.C0();
          _nw128.__ctor((_780_v90).dtor_cf37);
          _781_v91 = _nw128;
          _781_v91 = ((((p1) ? (p1) : (p1))) ? (_781_v91) : (_781_v91));
          let _782_v92;
          let _nw129 = new _module.C0();
          _nw129.__ctor(p1);
          _782_v92 = _nw129;
          let _783_v93;
          let _nw130 = new _module.C3();
          _nw130.__ctor(new BigNumber(-724), _762_v72, (_this).f6);
          _783_v93 = _nw130;
        } else {
          _763_v73 = _763_v73;
          let _784_v94;
          let _nw131 = Array((new BigNumber(27)).toNumber()).fill(_dafny.MultiSet.Empty);
          _784_v94 = _nw131;
          let _785_v95;
          _785_v95 = _dafny.MultiSet.fromElements((_dafny.ZERO).minus((_769_v79).f6));
          let _index116 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_784_v94).length));
          (_784_v94)[_index116] = _785_v95;
          let _index117 = _module.__default.safeIndex(new BigNumber(857), new BigNumber((_784_v94).length));
          (_784_v94)[_index117] = _785_v95;
          let _786_v96;
          _786_v96 = _dafny.Set.fromElements((_769_v79));
          let _787_v97;
          _787_v97 = _dafny.MultiSet.fromElements(p1);
          r3 = _dafny.Set.fromElements((_786_v96).IsProperSubsetOf(_786_v96), (_787_v97).IsDisjointFrom(_787_v97));
          let _788_v98;
          _788_v98 = _module.D7.create_DC18(!(p1), p1, _761_i7, p3);
          let _789_v99;
          _789_v99 = _dafny.Map.Empty.slice().updateUnsafe(_788_v98,_766_v76);
          let _790_v100;
          _790_v100 = _module.D11.create_DC24(_789_v99);
          _789_v99 = ((_790_v100).dtor_cf41).Merge((_789_v99).Merge(_789_v99));
          let _791_v101;
          _791_v101 = _dafny.Set.fromElements(p1);
          let _792_v102;
          _792_v102 = _module.D0.create_DC0(_dafny.Map.Empty.slice().updateUnsafe(p2,_791_v101));
          let _793_v103;
          _793_v103 = _dafny.Seq.of(p3);
          let _794_v104;
          _794_v104 = _module.D3.create_DC7(_759_v71);
          let _795_v105;
          _795_v105 = _module.D12.create_DC28(_module.__default.fm34(globalState));
          let _rhs107 = _module.__default.fm1(_792_v102, p1, _module.__default.fm15(new BigNumber(388), (_793_v103)[_module.__default.safeIndex(p2, new BigNumber((_793_v103).length))], globalState), (!(p3)) || (true), globalState);
          let _rhs108 = !_dafny.Seq.contains((_794_v104).dtor_cf13, _771_v81);
          let _rhs109 = _dafny.Seq.Create(_module.__default.abs(new BigNumber(601)), ((_796_v81) => function (_797_i9) {
            return _796_v81;
          })(_771_v81));
          let _rhs110 = (_793_v103)[_module.__default.safeIndex(p0, new BigNumber((_793_v103).length))];
          let _rhs111 = (_795_v105).dtor_cf51;
          let _lhs85 = globalState;
          let _lhs86 = globalState;
          let _lhs87 = globalState;
          _lhs85.f5 = _rhs107;
          _lhs86.f1 = _rhs108;
          _759_v71 = _rhs109;
          _lhs87.f5 = _rhs110;
          _773_v83 = _rhs111;
        }
      }
      let _798_v106;
      _798_v106 = _dafny.Set.fromElements(p1, p3);
      let _799_v107;
      _799_v107 = _dafny.Map.Empty.slice().updateUnsafe(p2,_798_v106);
      let _800_v108;
      _800_v108 = _module.D0.create_DC0(_799_v107);
      let _801_v109;
      _801_v109 = _dafny.Set.fromElements(_module.__default.fm2(_dafny.Seq.UnicodeFromString("et"), globalState), _759_v71);
      let _802_v110;
      _802_v110 = _dafny.Seq.of(p1);
      let _803_v111;
      _803_v111 = _dafny.Seq.of(_dafny.Seq.update(_dafny.Seq.update(_802_v110, _module.__default.safeIndex(p0, new BigNumber((_802_v110).length)), false), _module.__default.safeIndex(p0, new BigNumber((_dafny.Seq.update(_802_v110, _module.__default.safeIndex(p0, new BigNumber((_802_v110).length)), false)).length)), p3));
      let _rhs112 = _800_v108;
      let _rhs113 = _dafny.Seq.IsProperPrefixOf(_803_v111, _dafny.Seq.of(_802_v110));
      let _rhs114 = (function () {
        let _coll44 = new _dafny.Set();
        for (const _compr_44 of (_801_v109).Elements) {
          let _804_v112 = _compr_44;
          if ((_801_v109).contains(_804_v112)) {
            _coll44.add(_804_v112);
          }
        }
        return _coll44;
      }()).Difference((_801_v109).Difference(_801_v109));
      let _rhs115 = p3;
      let _lhs88 = globalState;
      let _lhs89 = globalState;
      _800_v108 = _rhs112;
      _lhs88.f5 = _rhs113;
      _801_v109 = _rhs114;
      _lhs89.f1 = _rhs115;
      r0 = (_this).f6;
      r1 = _module.D1.create_DC5();
      let _805_v113;
      _805_v113 = _dafny.Map.Empty.slice().updateUnsafe(p0,p1);
      r2 = new BigNumber((_805_v113).length);
      r3 = _798_v106;
      return [r0, r1, r2, r3];
    }
    m4(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = _dafny.ZERO;
      let r3 = new _dafny.CodePoint('D'.codePointAt(0));
      let _806_i0;
      _806_i0 = _dafny.ZERO;
      L2: {
        while (false) {
          C2: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_806_i0)) {
              break L2;
            }
            _806_i0 = (_806_i0).plus(_dafny.ONE);
            r2 = (_dafny.ZERO).minus((_this).f6);
            let _807_v0;
            _807_v0 = false;
            let _808_v1;
            _808_v1 = _dafny.Seq.of(true, _807_v0, _807_v0, _807_v0, _807_v0);
            let _809_v2;
            _809_v2 = _module.D7.create_DC19(((_807_v0) ? (_dafny.MultiSet.fromElements(false, _807_v0)) : ((_dafny.MultiSet.FromArray(_dafny.Seq.update(_808_v1, _module.__default.safeIndex((_this).f6, new BigNumber((_808_v1).length)), _807_v0))).update(_807_v0, _module.__default.abs((_this).f6)))), new BigNumber(415));
            let _source13 = _809_v2;
            if (_source13.is_DC18) {
              let _810___mcc_h0 = (_source13).cf29;
              let _811___mcc_h1 = (_source13).cf30;
              let _812___mcc_h2 = (_source13).cf31;
              let _813___mcc_h3 = (_source13).cf32;
              let _814_cf32 = _813___mcc_h3;
              let _815_cf31 = _812___mcc_h2;
              let _816_cf30 = _811___mcc_h1;
              let _817_cf29 = _810___mcc_h0;
              let _818_v3;
              _818_v3 = _dafny.Set.fromElements(_814_cf32);
              r1 = (new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(454)), ((_819_cf31) => function (_820_i1) {
                return _819_cf31;
              })(_815_cf31))).length)).minus((new BigNumber((_818_v3).length)).minus((_dafny.ZERO).minus((_this).f6)));
              let _821_v4;
              _821_v4 = _dafny.Seq.UnicodeFromString("mqmrskldp");
              let _822_v5;
              let _nw132 = new _module.C2();
              _nw132.__ctor(_821_v4, ((_814_cf32) ? (_816_cf30) : (_816_cf30)));
              _822_v5 = _nw132;
              (globalState).f2 = new BigNumber((_808_v1).length);
              let _823_v6;
              _823_v6 = _module.D3.create_DC7(_dafny.Seq.Concat((_822_v5).f11, (_822_v5).f11));
              let _824_v7;
              let _nw133 = Array((new BigNumber(2)).toNumber()).fill(_dafny.ZERO);
              _824_v7 = _nw133;
              let _825_v8;
              let _init21 = ((_826_v5) => function (_827_i2) {
                return ((_826_v5.f12) ? (new _dafny.CodePoint('g'.codePointAt(0))) : (new _dafny.CodePoint('p'.codePointAt(0))));
              })(_822_v5);
              let _nw134 = Array((new BigNumber(2)).toNumber());
              for (let _i0_21 = 0; _i0_21 < new BigNumber(_nw134.length); _i0_21++) {
                _nw134[_i0_21] = _init21(new BigNumber(_i0_21));
              }
              _825_v8 = _nw134;
              let _828_v9;
              _828_v9 = new _dafny.CodePoint('b'.codePointAt(0));
              let _index118 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_825_v8).length));
              (_825_v8)[_index118] = _828_v9;
              let _829_v10;
              _829_v10 = _dafny.MultiSet.fromElements(_815_cf31);
              let _pat_let_tv28 = _821_v4;
              let _index119 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_825_v8).length));
              let _rhs116 = function (_pat_let28_0) {
                return function (_832_dt__update__tmp_h1) {
                  return function (_pat_let31_0) {
                    return function (_833_dt__update_hcf13_h1) {
                      return _module.D3.create_DC7(_833_dt__update_hcf13_h1);
                    }(_pat_let31_0);
                  }(_dafny.Seq.UnicodeFromString("cwj"));
                }(_pat_let28_0);
              }(function (_pat_let29_0) {
                return function (_830_dt__update__tmp_h0) {
                  return function (_pat_let30_0) {
                    return function (_831_dt__update_hcf13_h0) {
                      return _module.D3.create_DC7(_831_dt__update_hcf13_h0);
                    }(_pat_let30_0);
                  }(_pat_let_tv28);
                }(_pat_let29_0);
              }(_823_v6));
              let _rhs117 = _824_v7;
              let _rhs118 = _828_v9;
              let _rhs119 = _829_v10;
              let _rhs120 = !(_dafny.Seq.contains((_822_v5).f11, _828_v9));
              let _lhs90 = _825_v8;
              let _lhs91 = _module.__default.safeIndex(new BigNumber(628), new BigNumber((_825_v8).length));
              _823_v6 = _rhs116;
              _824_v7 = _rhs117;
              _lhs90[_lhs91] = _rhs118;
              _829_v10 = _rhs119;
              _814_cf32 = _rhs120;
            } else if (_source13.is_DC19) {
              let _834___mcc_h4 = (_source13).cf33;
              let _835___mcc_h5 = (_source13).cf34;
              let _836_cf34 = _835___mcc_h5;
              let _837_cf33 = _834___mcc_h4;
              let _838_v11;
              _838_v11 = new _dafny.CodePoint('i'.codePointAt(0));
              let _839_v12;
              _839_v12 = _dafny.MultiSet.fromElements(_838_v11);
              let _840_v13;
              _840_v13 = _dafny.Map.Empty.slice().updateUnsafe(_836_cf34,_807_v0);
              let _841_v14;
              _841_v14 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm15((((_839_v12).contains(_838_v11)) ? ((_839_v12).get(_838_v11)) : (_836_cf34)), (((_840_v13).contains((_this).f6)) ? ((_840_v13).get((_this).f6)) : (true)), globalState),_807_v0);
              let _842_v15;
              _842_v15 = _module.D13.create_DC32((_this).f6, _807_v0, _807_v0, _807_v0, (_808_v1)[_module.__default.safeIndex(_836_cf34, new BigNumber((_808_v1).length))]);
              let _843_v16;
              _843_v16 = _dafny.Seq.UnicodeFromString("ohdkp");
              let _844_v17;
              _844_v17 = _dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_843_v16).length)));
              let _845_v18;
              _845_v18 = _dafny.Set.fromElements(_807_v0);
              let _846_v19;
              let _init22 = ((_847_v0) => function (_848_i3) {
                return _847_v0;
              })(_807_v0);
              let _nw135 = Array((new BigNumber(4)).toNumber());
              for (let _i0_22 = 0; _i0_22 < new BigNumber(_nw135.length); _i0_22++) {
                _nw135[_i0_22] = _init22(new BigNumber(_i0_22));
              }
              _846_v19 = _nw135;
              let _849_v20;
              _849_v20 = _dafny.Map.Empty.slice().updateUnsafe(_846_v19,_841_v14);
              let _850_v21;
              _850_v21 = _dafny.MultiSet.fromElements(_836_cf34);
              let _851_v22;
              let _nw136 = Array((new BigNumber(28)).toNumber());
              _nw136[(_dafny.ZERO).toNumber()] = _807_v0;
              _nw136[(_dafny.ONE).toNumber()] = _807_v0;
              _nw136[(new BigNumber(2)).toNumber()] = true;
              _nw136[(new BigNumber(3)).toNumber()] = !(_807_v0);
              _nw136[(new BigNumber(4)).toNumber()] = !(_807_v0);
              _nw136[(new BigNumber(5)).toNumber()] = false;
              _nw136[(new BigNumber(6)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(7)).toNumber()] = (_840_v13).equals(((_module.D13.create_DC31(_841_v14)).dtor_cf55).update(_836_cf34, false));
              _nw136[(new BigNumber(8)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(9)).toNumber()] = (_842_v15).dtor_cf59;
              _nw136[(new BigNumber(10)).toNumber()] = ((_this).f6).isLessThan(new BigNumber((_844_v17).length));
              _nw136[(new BigNumber(11)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(12)).toNumber()] = ((_this).f6).isLessThan((_dafny.ZERO).minus(new BigNumber((_845_v18).length)));
              _nw136[(new BigNumber(13)).toNumber()] = !(_807_v0) || (_807_v0);
              _nw136[(new BigNumber(14)).toNumber()] = (_807_v0) && (_807_v0);
              _nw136[(new BigNumber(15)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(16)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(17)).toNumber()] = (_dafny.MultiSet.fromElements(true, _807_v0, _807_v0, _807_v0, (((_840_v13).contains(new BigNumber((_849_v20).length))) ? ((_840_v13).get(new BigNumber((_849_v20).length))) : (_807_v0)))).IsSubsetOf(_837_cf33);
              _nw136[(new BigNumber(18)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(19)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(20)).toNumber()] = !_dafny.areEqual(_844_v17, _844_v17);
              _nw136[(new BigNumber(21)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(22)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(23)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(24)).toNumber()] = _807_v0;
              _nw136[(new BigNumber(25)).toNumber()] = (_836_cf34).isEqualTo((_this).f6);
              _nw136[(new BigNumber(26)).toNumber()] = !((_dafny.MultiSet.FromArray(_844_v17)).IsProperSubsetOf(_850_v21));
              _nw136[(new BigNumber(27)).toNumber()] = !(false) || (_807_v0);
              _851_v22 = _nw136;
              _851_v22 = _846_v19;
              let _852_v23;
              let _nw137 = Array((new BigNumber(14)).toNumber());
              _nw137[(_dafny.ZERO).toNumber()] = _843_v16;
              _nw137[(_dafny.ONE).toNumber()] = _843_v16;
              _nw137[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(896)), ((_853_v11) => function (_854_i4) {
                return _853_v11;
              })(_838_v11));
              _nw137[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("khfcae");
              _nw137[(new BigNumber(4)).toNumber()] = _843_v16;
              _nw137[(new BigNumber(5)).toNumber()] = _843_v16;
              _nw137[(new BigNumber(6)).toNumber()] = _843_v16;
              _nw137[(new BigNumber(7)).toNumber()] = _843_v16;
              _nw137[(new BigNumber(8)).toNumber()] = _843_v16;
              _nw137[(new BigNumber(9)).toNumber()] = _843_v16;
              _nw137[(new BigNumber(10)).toNumber()] = _dafny.Seq.Concat(_843_v16, _843_v16);
              _nw137[(new BigNumber(11)).toNumber()] = _dafny.Seq.UnicodeFromString("yareb");
              _nw137[(new BigNumber(12)).toNumber()] = _843_v16;
              _nw137[(new BigNumber(13)).toNumber()] = _dafny.Seq.UnicodeFromString("aqkxf");
              _852_v23 = _nw137;
              let _index120 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_852_v23).length));
              (_852_v23)[_index120] = _dafny.Seq.Concat(_843_v16, _843_v16);
              let _index121 = _module.__default.safeIndex(new BigNumber(152), new BigNumber((_852_v23).length));
              (_852_v23)[_index121] = _dafny.Seq.update(_dafny.Seq.Concat(_843_v16, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-243)), ((_855_v11) => function (_856_i5) {
                return _855_v11;
              })(_838_v11)), _843_v16)), _module.__default.safeIndex(_836_cf34, new BigNumber((_dafny.Seq.Concat(_843_v16, _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(-243)), ((_857_v11) => function (_858_i5) {
                return _857_v11;
              })(_838_v11)), _843_v16))).length)), _838_v11);
              (globalState).f5 = _807_v0;
              _844_v17 = _844_v17;
            } else {
              let _859___mcc_h6 = (_source13).cf28;
              let _860_cf28 = _859___mcc_h6;
              let _861_v24;
              let _init23 = ((_862_v0) => function (_863_i6) {
                return _862_v0;
              })(_807_v0);
              let _nw138 = Array((new BigNumber(8)).toNumber());
              for (let _i0_23 = 0; _i0_23 < new BigNumber(_nw138.length); _i0_23++) {
                _nw138[_i0_23] = _init23(new BigNumber(_i0_23));
              }
              _861_v24 = _nw138;
              let _864_v25;
              _864_v25 = _dafny.Seq.of(_861_v24, _861_v24);
              _864_v25 = _864_v25;
              (globalState).f1 = _807_v0;
              let _865_v26;
              _865_v26 = new _dafny.CodePoint('o'.codePointAt(0));
              r3 = _865_v26;
              let _866_v27;
              _866_v27 = _dafny.Set.fromElements(_807_v0, _807_v0, _807_v0);
              let _867_v28;
              _867_v28 = _dafny.Map.Empty.slice().updateUnsafe(!(_866_v27).equals(_866_v27),(_this).f6);
              _867_v28 = (_867_v28).update(_807_v0, (_this).f6);
            }
            let _868_v29;
            _868_v29 = _dafny.Seq.UnicodeFromString("lropjewiy");
            let _869_v30;
            _869_v30 = _dafny.Set.fromElements(_807_v0);
            let _870_v31;
            _870_v31 = _dafny.Set.fromElements(new BigNumber((_869_v30).length), (_module.__default.fm15((_this).f6, _807_v0, globalState)).plus((_this).f6), (_this).f6, (_dafny.ZERO).minus((_this).f6));
            let _871_v32;
            let _init24 = ((_872_v0) => function (_873_i7) {
              return _872_v0;
            })(_807_v0);
            let _nw139 = Array((new BigNumber(7)).toNumber());
            for (let _i0_24 = 0; _i0_24 < new BigNumber(_nw139.length); _i0_24++) {
              _nw139[_i0_24] = _init24(new BigNumber(_i0_24));
            }
            _871_v32 = _nw139;
            let _874_v33;
            let _nw140 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
            _874_v33 = _nw140;
            let _index122 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_874_v33).length));
            (_874_v33)[_index122] = (_this).f6;
            let _875_v34;
            _875_v34 = _module.D7.create_DC18(_807_v0, _807_v0, (_dafny.ZERO).minus((_this).f6), _807_v0);
            let _index123 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_874_v33).length));
            let _rhs121 = (_module.D3.create_DC7(_868_v29)).dtor_cf13;
            let _rhs122 = _dafny.Set.fromElements((_this).f6);
            let _rhs123 = _871_v32;
            let _rhs124 = (_this).f6;
            let _rhs125 = (_875_v34).dtor_cf31;
            let _lhs92 = _874_v33;
            let _lhs93 = _module.__default.safeIndex(new BigNumber(381), new BigNumber((_874_v33).length));
            let _lhs94 = globalState;
            _868_v29 = _rhs121;
            _870_v31 = _rhs122;
            _871_v32 = _rhs123;
            _lhs92[_lhs93] = _rhs124;
            _lhs94.f0 = _rhs125;
            let _876_v35;
            _876_v35 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,((_807_v0) ? (new BigNumber((_808_v1).length)) : ((_874_v33)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_874_v33).length))])));
            let _877_v36;
            _877_v36 = _dafny.Seq.of((_874_v33)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_874_v33).length))], (_this).f6);
            let _878_v37;
            _878_v37 = new _dafny.CodePoint('p'.codePointAt(0));
            let _879_v38;
            _879_v38 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.safeModuloInt(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_dafny.Seq.update(_877_v36, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_808_v1).length)), new BigNumber((_877_v36).length)), (_874_v33)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_874_v33).length))])).length),new BigNumber(-615))).length), _module.__default.fm24((_874_v33)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_874_v33).length))], globalState)),_878_v37);
            let _rhs126 = _module.__default.fm15((_874_v33)[_module.__default.safeIndex(new BigNumber(381), new BigNumber((_874_v33).length))], _807_v0, globalState);
            let _rhs127 = (((_879_v38).contains(new BigNumber((function () {
              let _coll46 = new _dafny.Map();
              for (const _compr_46 of (_868_v29).Elements) {
                let _881_v39 = _compr_46;
                if (_dafny.Seq.contains(_868_v29, _881_v39)) {
                  _coll46.push([_881_v39,_807_v0]);
                }
              }
              return _coll46;
            }()).length))) ? ((_879_v38).get(new BigNumber((function () {
              let _coll45 = new _dafny.Map();
              for (const _compr_45 of (_868_v29).Elements) {
                let _880_v39 = _compr_45;
                if (_dafny.Seq.contains(_868_v29, _880_v39)) {
                  _coll45.push([_880_v39,_807_v0]);
                }
              }
              return _coll45;
            }()).length))) : (_878_v37));
            let _rhs128 = _876_v35;
            r1 = _rhs126;
            r3 = _rhs127;
            _876_v35 = _rhs128;
          }
        }
      }
      let _882_v40;
      _882_v40 = true;
      if (_882_v40) {
        (globalState).f5 = _882_v40;
        let _883_v41;
        let _init25 = ((_884_v40) => function (_885_i8) {
          return _884_v40;
        })(_882_v40);
        let _nw141 = Array((new BigNumber(24)).toNumber());
        for (let _i0_25 = 0; _i0_25 < new BigNumber(_nw141.length); _i0_25++) {
          _nw141[_i0_25] = _init25(new BigNumber(_i0_25));
        }
        _883_v41 = _nw141;
        let _886_v42;
        _886_v42 = _dafny.Map.Empty.slice().updateUnsafe(_882_v40,_883_v41);
        let _rhs129 = _886_v42;
        let _rhs130 = _module.__default.safeDivisionInt((_dafny.ZERO).minus((_this).f6), (_this).f6);
        _886_v42 = _rhs129;
        r1 = _rhs130;
        let _887_v43;
        _887_v43 = new _dafny.CodePoint('j'.codePointAt(0));
        let _888_v44;
        _888_v44 = _dafny.Map.Empty.slice().updateUnsafe(_887_v43,(_this).f6);
        let _889_v45;
        _889_v45 = _dafny.Seq.of((((_888_v44).contains(_887_v43)) ? ((_888_v44).get(_887_v43)) : ((_this).f6)));
        _889_v45 = (_module.D1.create_DC4(_889_v45)).dtor_cf11;
        let _rhs131 = (((_886_v42).contains(_882_v40)) ? ((_886_v42).get(_882_v40)) : (_883_v41));
        let _rhs132 = _module.__default.fm24((_this).f6, globalState);
        let _lhs95 = globalState;
        _883_v41 = _rhs131;
        _lhs95.f0 = _rhs132;
        let _890_v46;
        _890_v46 = _dafny.Seq.UnicodeFromString("nemkv");
        let _891_v47;
        _891_v47 = _dafny.Map.Empty.slice().updateUnsafe(_882_v40,_890_v46);
        let _892_v48;
        _892_v48 = _dafny.MultiSet.fromElements(_882_v40);
        let _893_v49;
        let _nw142 = new _module.C3();
        _nw142.__ctor(new BigNumber(186), _889_v45, new BigNumber((_892_v48).cardinality()));
        _893_v49 = _nw142;
        let _894_v50;
        let _nw143 = new _module.C3();
        _nw143.__ctor((_this).f6, _889_v45, new BigNumber((_module.__default.fm2(_890_v46, globalState)).length));
        _894_v50 = _nw143;
        let _895_v51;
        _895_v51 = _dafny.MultiSet.fromElements(_894_v50);
        let _896_v52;
        _896_v52 = _895_v51;
        let _897_v53;
        _897_v53 = _module.D11.create_DC26(_893_v49, _896_v52, new BigNumber((_888_v44).length), _890_v46);
        let _898_v54;
        let _nw144 = Array((new BigNumber(17)).toNumber());
        _nw144[(_dafny.ZERO).toNumber()] = _dafny.Seq.update(_890_v46, _module.__default.safeIndex((_889_v45)[_module.__default.safeIndex((_this).f6, new BigNumber((_889_v45).length))], new BigNumber((_890_v46).length)), _887_v43);
        _nw144[(_dafny.ONE).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(199)), ((_899_v43) => function (_900_i9) {
          return _899_v43;
        })(_887_v43));
        _nw144[(new BigNumber(2)).toNumber()] = _dafny.Seq.Create(_module.__default.abs(new BigNumber(947)), ((_901_v43) => function (_902_i10) {
          return _901_v43;
        })(_887_v43));
        _nw144[(new BigNumber(3)).toNumber()] = _dafny.Seq.UnicodeFromString("npv");
        _nw144[(new BigNumber(4)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("ev"), _890_v46);
        _nw144[(new BigNumber(5)).toNumber()] = _dafny.Seq.Concat(_890_v46, _890_v46);
        _nw144[(new BigNumber(6)).toNumber()] = _890_v46;
        _nw144[(new BigNumber(7)).toNumber()] = _dafny.Seq.Concat(_890_v46, _890_v46);
        _nw144[(new BigNumber(8)).toNumber()] = _890_v46;
        _nw144[(new BigNumber(9)).toNumber()] = _890_v46;
        _nw144[(new BigNumber(10)).toNumber()] = ((_882_v40) ? (_890_v46) : (_dafny.Seq.update(_dafny.Seq.Create(_module.__default.abs(new BigNumber(489)), ((_903_v43) => function (_904_i11) {
          return _903_v43;
        })(_887_v43)), _module.__default.safeIndex((_this).f6, new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(489)), ((_905_v43) => function (_906_i11) {
          return _905_v43;
        })(_887_v43))).length)), new _dafny.CodePoint('p'.codePointAt(0)))));
        _nw144[(new BigNumber(11)).toNumber()] = _890_v46;
        _nw144[(new BigNumber(12)).toNumber()] = _dafny.Seq.Concat(_dafny.Seq.update(_890_v46, _module.__default.safeIndex((_this).f6, new BigNumber((_890_v46).length)), _887_v43), _890_v46);
        _nw144[(new BigNumber(13)).toNumber()] = _dafny.Seq.Concat(_890_v46, (((_891_v47).contains(_882_v40)) ? ((_891_v47).get(_882_v40)) : (_890_v46)));
        _nw144[(new BigNumber(14)).toNumber()] = _890_v46;
        _nw144[(new BigNumber(15)).toNumber()] = ((_882_v40) ? (_890_v46) : (_890_v46));
        _nw144[(new BigNumber(16)).toNumber()] = _dafny.Seq.update((_897_v53).dtor_cf48, _module.__default.safeIndex((_893_v49).f6, new BigNumber(((_897_v53).dtor_cf48).length)), _887_v43);
        _898_v54 = _nw144;
        let _index124 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_898_v54).length));
        (_898_v54)[_index124] = _890_v46;
        let _index125 = _module.__default.safeIndex(new BigNumber(946), new BigNumber((_898_v54).length));
        (_898_v54)[_index125] = _dafny.Seq.UnicodeFromString("nqeesu");
      } else {
        let _907_v55;
        _907_v55 = _dafny.Seq.UnicodeFromString("rgltcl");
        _907_v55 = _module.__default.fm2(_907_v55, globalState);
        let _908_v56;
        let _nw145 = Array((new BigNumber(9)).toNumber()).fill(_dafny.Map.Empty);
        _908_v56 = _nw145;
        let _909_v57;
        _909_v57 = _dafny.Map.Empty.slice().updateUnsafe(_882_v40,new BigNumber((_dafny.Seq.of((_this).f6, (_this).f6)).length));
        let _index126 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_908_v56).length));
        (_908_v56)[_index126] = _909_v57;
        let _910_v58;
        _910_v58 = _dafny.Map.Empty.slice().updateUnsafe(_882_v40,new _dafny.CodePoint('o'.codePointAt(0)));
        let _911_v59;
        _911_v59 = new _dafny.CodePoint('t'.codePointAt(0));
        let _912_v60;
        _912_v60 = _dafny.Seq.of((_dafny.Map.Empty.slice().updateUnsafe(_882_v40,new BigNumber((_dafny.Seq.of((((_910_v58).contains(!(_882_v40))) ? ((_910_v58).get(!(_882_v40))) : (_911_v59)))).length))).Merge(_909_v57), _909_v57);
        let _913_v61;
        _913_v61 = _dafny.MultiSet.fromElements(_882_v40);
        let _index127 = _module.__default.safeIndex(new BigNumber(84), new BigNumber((_908_v56).length));
        (_908_v56)[_index127] = (_912_v60)[_module.__default.safeIndex((new BigNumber((_913_v61).cardinality())).multipliedBy((_this).f6), new BigNumber((_912_v60).length))];
        r0 = new BigNumber(575);
        let _914_v62;
        _914_v62 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f6),_dafny.Set.fromElements(_882_v40));
        let _915_v63;
        _915_v63 = _module.D0.create_DC0(_914_v62);
        let _916_v64;
        let _nw146 = new _module.C1();
        _nw146.__ctor();
        _916_v64 = _nw146;
        let _917_v65;
        _917_v65 = _dafny.Seq.of(_882_v40);
        let _918_v66;
        let _nw147 = Array((new BigNumber(20)).toNumber());
        _nw147[(_dafny.ZERO).toNumber()] = _module.__default.fm1(_915_v63, _882_v40, (_this).f6, _882_v40, globalState);
        _nw147[(_dafny.ONE).toNumber()] = _882_v40;
        _nw147[(new BigNumber(2)).toNumber()] = (_917_v65)[_module.__default.safeIndex(new BigNumber(136), new BigNumber((_917_v65).length))];
        _nw147[(new BigNumber(3)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(4)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(5)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(6)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(7)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(8)).toNumber()] = true;
        _nw147[(new BigNumber(9)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(10)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(11)).toNumber()] = true;
        _nw147[(new BigNumber(12)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(13)).toNumber()] = !(_882_v40);
        _nw147[(new BigNumber(14)).toNumber()] = !(_882_v40);
        _nw147[(new BigNumber(15)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(16)).toNumber()] = false;
        _nw147[(new BigNumber(17)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(18)).toNumber()] = _882_v40;
        _nw147[(new BigNumber(19)).toNumber()] = _882_v40;
        _918_v66 = _nw147;
        let _919_v67;
        _919_v67 = _dafny.MultiSet.fromElements(_918_v66);
        let _920_v68;
        _920_v68 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_916_v64);
        let _921_v69;
        _921_v69 = _dafny.Seq.of(_916_v64);
        r2 = new BigNumber((_dafny.Seq.Concat(((_module.__default.fm1(_915_v63, true, _module.__default.fm24(new BigNumber(731), globalState), _882_v40, globalState)) ? (_dafny.Seq.of(_916_v64)) : (_dafny.Seq.update(_dafny.Seq.of(_916_v64, _916_v64), _module.__default.safeIndex(new BigNumber((_919_v67).cardinality()), new BigNumber((_dafny.Seq.of(_916_v64, _916_v64)).length)), (((_920_v68).contains(new BigNumber(72))) ? ((_920_v68).get(new BigNumber(72))) : (_916_v64))))), _dafny.Seq.Concat(_921_v69, _921_v69))).length);
        let _922_v70;
        _922_v70 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_this).f6);
        r1 = _module.__default.safeDivisionInt(_module.__default.safeModuloInt(new BigNumber((_dafny.Seq.of(_882_v40, _882_v40, (_917_v65)[_module.__default.safeIndex((_this).f6, new BigNumber((_917_v65).length))], _882_v40)).length), (_this).f6), ((((_922_v70).contains((_this).f6)) ? ((_922_v70).get((_this).f6)) : ((_dafny.ZERO).minus(new BigNumber((_907_v55).length))))).multipliedBy((_this).f6));
      }
      let _923_v71;
      let _init26 = function (_924_i13) {
        return _dafny.Seq.IsPrefixOf(_dafny.Seq.UnicodeFromString("igm"), _dafny.Seq.UnicodeFromString("sgxxap"));
      };
      let _nw148 = Array((new BigNumber(24)).toNumber());
      for (let _i0_26 = 0; _i0_26 < new BigNumber(_nw148.length); _i0_26++) {
        _nw148[_i0_26] = _init26(new BigNumber(_i0_26));
      }
      _923_v71 = _nw148;
      for (const _guard_loop_3 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_923_v71).length))) {
        let _925_i12 = _guard_loop_3;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_925_i12)) && ((_925_i12).isLessThan(new BigNumber((_923_v71).length))))) {
          (_923_v71)[(_925_i12)] = !((_dafny.Set.fromElements(_module.D7.create_DC17(_dafny.Seq.of(_882_v40, _882_v40)), _module.D7.create_DC17(_dafny.Seq.of(_882_v40, _882_v40)))).Union(_dafny.Set.fromElements(_module.D7.create_DC17(_dafny.Seq.of(_882_v40))))).equals(_dafny.Set.fromElements(_module.D7.create_DC17(_dafny.Seq.of(_882_v40))));
        }
      }
      let _926_v72;
      _926_v72 = _dafny.Seq.UnicodeFromString("gvljalgu");
      let _927_v73;
      _927_v73 = _module.D3.create_DC7(_926_v72);
      let _source14 = _module.D3.create_DC9(((_882_v40) ? (_module.D3.create_DC9(_927_v73)) : (_927_v73)));
      if (_source14.is_DC8) {
        let _928___mcc_h7 = (_source14).cf14;
        let _929___mcc_h8 = (_source14).cf15;
        let _930___mcc_h9 = (_source14).cf16;
        let _931_cf16 = _930___mcc_h9;
        let _932_cf15 = _929___mcc_h8;
        let _933_cf14 = _928___mcc_h7;
        r1 = _module.__default.safeDivisionInt(_931_cf16, _931_cf16);
        let _934_v74;
        _934_v74 = _923_v71;
        let _rhs133 = (_934_v74);
        let _rhs134 = _932_cf15;
        _923_v71 = _rhs133;
        _932_cf15 = _rhs134;
        let _index128 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_923_v71).length));
        (_923_v71)[_index128] = _882_v40;
        let _index129 = _module.__default.safeIndex(new BigNumber(981), new BigNumber((_923_v71).length));
        (_923_v71)[_index129] = _882_v40;
        let _935_v75;
        let _nw149 = new _module.C1();
        _nw149.__ctor();
        _935_v75 = _nw149;
      } else if (_source14.is_DC7) {
        let _936___mcc_h10 = (_source14).cf13;
        let _937_cf13 = _936___mcc_h10;
        let _938_v76;
        let _nw150 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
        _938_v76 = _nw150;
        let _939_v77;
        _939_v77 = _module.D3.create_DC7(_dafny.Seq.UnicodeFromString("dfgoq"));
        let _index130 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_938_v76).length));
        (_938_v76)[_index130] = new BigNumber(((_939_v77).dtor_cf13).length);
        let _index131 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_938_v76).length));
        (_938_v76)[_index131] = (_this).f6;
        (globalState).f5 = _882_v40;
        (globalState).f5 = (_882_v40) || (_882_v40);
        let _940_v78;
        let _init27 = function (_941_i14) {
          return _dafny.Map.Empty.slice().updateUnsafe(new _dafny.CodePoint('j'.codePointAt(0)),false);
        };
        let _nw151 = Array((new BigNumber(11)).toNumber());
        for (let _i0_27 = 0; _i0_27 < new BigNumber(_nw151.length); _i0_27++) {
          _nw151[_i0_27] = _init27(new BigNumber(_i0_27));
        }
        _940_v78 = _nw151;
        let _942_v79;
        _942_v79 = _module.D14.create_DC35(_940_v78);
        let _943_v80;
        let _nw152 = Array((new BigNumber(16)).toNumber());
        _nw152[(_dafny.ZERO).toNumber()] = _940_v78;
        _nw152[(_dafny.ONE).toNumber()] = _940_v78;
        _nw152[(new BigNumber(2)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(3)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(4)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(5)).toNumber()] = ((_882_v40) ? (_940_v78) : (_940_v78));
        _nw152[(new BigNumber(6)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(7)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(8)).toNumber()] = (_942_v79).dtor_cf67;
        _nw152[(new BigNumber(9)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(10)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(11)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(12)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(13)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(14)).toNumber()] = _940_v78;
        _nw152[(new BigNumber(15)).toNumber()] = ((_882_v40) ? (_940_v78) : (_940_v78));
        _943_v80 = _nw152;
        let _index132 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_943_v80).length));
        (_943_v80)[_index132] = _940_v78;
        let _index133 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_943_v80).length));
        let _index134 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_938_v76).length));
        let _rhs135 = !(_882_v40);
        let _rhs136 = (_938_v76)[_module.__default.safeIndex(new BigNumber(412), new BigNumber((_938_v76).length))];
        let _rhs137 = _940_v78;
        let _rhs138 = (_this).f6;
        let _lhs96 = globalState;
        let _lhs97 = globalState;
        let _lhs98 = _943_v80;
        let _lhs99 = _module.__default.safeIndex(new BigNumber(201), new BigNumber((_943_v80).length));
        let _lhs100 = _938_v76;
        let _lhs101 = _module.__default.safeIndex(new BigNumber(412), new BigNumber((_938_v76).length));
        _lhs96.f1 = _rhs135;
        _lhs97.f2 = _rhs136;
        _lhs98[_lhs99] = _rhs137;
        _lhs100[_lhs101] = _rhs138;
      } else {
        let _944___mcc_h11 = (_source14).cf17;
        let _945_cf17 = _944___mcc_h11;
        let _946_v81;
        _946_v81 = _dafny.Seq.of(_882_v40, !(_882_v40));
        let _947_v82;
        _947_v82 = _module.D12.create_DC28(_dafny.Map.Empty.slice().updateUnsafe(_882_v40,new BigNumber((_946_v81).length)));
        _947_v82 = _947_v82;
        let _948_v83;
        _948_v83 = _dafny.Set.fromElements(_882_v40, _882_v40);
        let _949_v84;
        let _nw153 = new _module.C0();
        _nw153.__ctor(_882_v40);
        _949_v84 = _nw153;
        let _950_v85;
        _950_v85 = _949_v84;
        let _951_v86;
        let _nw154 = Array((new BigNumber(18)).toNumber());
        _nw154[(_dafny.ZERO).toNumber()] = _949_v84;
        _nw154[(_dafny.ONE).toNumber()] = _949_v84;
        _nw154[(new BigNumber(2)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(3)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(4)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(5)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(6)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(7)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(8)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(9)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(10)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(11)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(12)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(13)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(14)).toNumber()] = (_950_v85);
        _nw154[(new BigNumber(15)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(16)).toNumber()] = _949_v84;
        _nw154[(new BigNumber(17)).toNumber()] = _949_v84;
        _951_v86 = _nw154;
        let _952_v87;
        _952_v87 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_948_v83).length),_951_v86);
        _952_v87 = (_952_v87).update((_this).f6, _951_v86);
        let _953_v88;
        _953_v88 = _dafny.Set.fromElements((_this).f6);
        let _954_v89;
        _954_v89 = _dafny.Seq.of((_this).f6, (_dafny.ZERO).minus((_this).f6), (_this).f6, (_this).f6, (_this).f6);
        let _955_v90;
        let _nw155 = new _module.C3();
        _nw155.__ctor(_module.__default.safeModuloInt((_this).f6, new BigNumber((_953_v88).length)), _dafny.Seq.Concat(_dafny.Seq.of((_this).f6), _954_v89), new BigNumber(918));
        _955_v90 = _nw155;
        let _956_v91;
        let _nw156 = new _module.C0();
        _nw156.__ctor(((_882_v40) ? (_882_v40) : (_882_v40)));
        _956_v91 = _nw156;
      }
      let _957_v92;
      _957_v92 = _dafny.Set.fromElements(_882_v40);
      let _958_v93;
      _958_v93 = _module.D6.create_DC15(_957_v92);
      let _959_i15;
      _959_i15 = _dafny.ZERO;
      L3: {
        let _pat_let_tv29 = _882_v40;
        let _pat_let_tv30 = _882_v40;
        while (function (_source15) {
          if (_source15.is_DC16) {
            let _969___mcc_h12 = (_source15).cf27;
            let _970_cf27 = _969___mcc_h12;
            return (new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(_pat_let_tv29,_dafny.Seq.of(new BigNumber(13)))).length)).isLessThan((_this).f6);
          } else {
            let _971___mcc_h13 = (_source15).cf26;
            let _972_cf26 = _971___mcc_h13;
            return _pat_let_tv30;
          }
        }(_958_v93)) {
          C3: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_959_i15)) {
              break L3;
            }
            _959_i15 = (_959_i15).plus(_dafny.ONE);
            let _960_v94;
            _960_v94 = _dafny.Seq.of((_this).f6);
            let _961_v95;
            _961_v95 = _module.D13.create_DC32(new BigNumber((_960_v94).length), _882_v40, _882_v40, _882_v40, _882_v40);
            let _962_v96;
            _962_v96 = _dafny.Seq.of((_961_v95).dtor_cf56);
            let _963_v97;
            let _nw157 = new _module.C3();
            _nw157.__ctor((_this).f6, _960_v94, (_this).f6);
            _963_v97 = _nw157;
            (globalState).f0 = _module.__default.safeDivisionInt((_this).f6, (new BigNumber((_926_v72).length)).multipliedBy((_963_v97).f9));
            let _964_v98;
            let _init28 = ((_965_v92) => function (_966_i16) {
              return _module.__default.safeDivisionInt(_966_i16, new BigNumber((_965_v92).length));
            })(_957_v92);
            let _nw158 = Array((new BigNumber(23)).toNumber());
            for (let _i0_28 = 0; _i0_28 < new BigNumber(_nw158.length); _i0_28++) {
              _nw158[_i0_28] = _init28(new BigNumber(_i0_28));
            }
            _964_v98 = _nw158;
            let _index135 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_964_v98).length));
            (_964_v98)[_index135] = (_this).f6;
            let _967_v99;
            _967_v99 = _dafny.Map.Empty.slice().updateUnsafe(_882_v40,(_this).f6);
            let _968_v100;
            _968_v100 = _module.D12.create_DC28(_967_v99);
            let _index136 = _module.__default.safeIndex(new BigNumber(473), new BigNumber((_964_v98).length));
            (_964_v98)[_index136] = new BigNumber(((_968_v100).dtor_cf51).length);
            (globalState).f1 = _882_v40;
          }
        }
      }
      let _973_i17;
      _973_i17 = _dafny.ZERO;
      L4: {
        while (false) {
          C4: {
            if ((new BigNumber(100)).isLessThanOrEqualTo(_973_i17)) {
              break L4;
            }
            _973_i17 = (_973_i17).plus(_dafny.ONE);
            let _974_v101;
            _974_v101 = new _dafny.CodePoint('j'.codePointAt(0));
            let _975_v102;
            let _nw159 = Array((new BigNumber(19)).toNumber());
            _nw159[(_dafny.ZERO).toNumber()] = _974_v101;
            _nw159[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('l'.codePointAt(0));
            _nw159[(new BigNumber(2)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(3)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(4)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(5)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(6)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(7)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(8)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(9)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(10)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(11)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(12)).toNumber()] = ((false) ? (_974_v101) : (_974_v101));
            _nw159[(new BigNumber(13)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(14)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(15)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(16)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(17)).toNumber()] = _974_v101;
            _nw159[(new BigNumber(18)).toNumber()] = _974_v101;
            _975_v102 = _nw159;
            let _index137 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_975_v102).length));
            (_975_v102)[_index137] = new _dafny.CodePoint('m'.codePointAt(0));
            let _index138 = _module.__default.safeIndex(new BigNumber(37), new BigNumber((_975_v102).length));
            (_975_v102)[_index138] = _974_v101;
            let _976_v103;
            _976_v103 = _dafny.Seq.of(_882_v40, _882_v40);
            let _977_v104;
            _977_v104 = _dafny.MultiSet.fromElements(_882_v40);
            let _978_v105;
            _978_v105 = _dafny.Seq.of((_this).f6);
            let _979_v106;
            _979_v106 = _module.D1.create_DC4(_978_v105);
            let _980_v107;
            let _nw160 = Array((new BigNumber(11)).toNumber()).fill(_dafny.ZERO);
            _980_v107 = _nw160;
            let _981_v108;
            _981_v108 = _module.D3.create_DC8(_979_v106, _980_v107, (_this).f6);
            let _982_v109;
            _982_v109 = _dafny.Map.Empty.slice().updateUnsafe(_882_v40,_882_v40);
            let _983_v110;
            _983_v110 = _dafny.Map.Empty.slice().updateUnsafe(false,new BigNumber((_982_v109).length));
            let _984_v111;
            _984_v111 = _dafny.MultiSet.fromElements(new BigNumber(375), new BigNumber((_982_v109).length));
            let _pat_let_tv31 = _984_v111;
            let _985_v112;
            let _nw161 = Array((new BigNumber(23)).toNumber());
            _nw161[(_dafny.ZERO).toNumber()] = (_this).f6;
            _nw161[(_dafny.ONE).toNumber()] = new BigNumber((_976_v103).length);
            _nw161[(new BigNumber(2)).toNumber()] = (_this).f6;
            _nw161[(new BigNumber(3)).toNumber()] = (_this).f6;
            _nw161[(new BigNumber(4)).toNumber()] = new BigNumber(364);
            _nw161[(new BigNumber(5)).toNumber()] = new BigNumber((_977_v104).cardinality());
            _nw161[(new BigNumber(6)).toNumber()] = (_this).f6;
            _nw161[(new BigNumber(7)).toNumber()] = new BigNumber((_dafny.MultiSet.FromArray(_978_v105)).cardinality());
            _nw161[(new BigNumber(8)).toNumber()] = (_this).f6;
            _nw161[(new BigNumber(9)).toNumber()] = (_this).f6;
            _nw161[(new BigNumber(10)).toNumber()] = (_this).f6;
            _nw161[(new BigNumber(11)).toNumber()] = ((_this).f6).multipliedBy((_this).f6);
            _nw161[(new BigNumber(12)).toNumber()] = _module.__default.fm15((_981_v108).dtor_cf16, _882_v40, globalState);
            _nw161[(new BigNumber(13)).toNumber()] = (_this).f6;
            _nw161[(new BigNumber(14)).toNumber()] = (new BigNumber((_983_v110).length)).minus(new BigNumber((_977_v104).cardinality()));
            _nw161[(new BigNumber(15)).toNumber()] = new BigNumber((_976_v103).length);
            _nw161[(new BigNumber(16)).toNumber()] = _module.__default.safeModuloInt((_this).f6, new BigNumber(((function (_pat_let32_0) {
              return function (_986_dt__update__tmp_h2) {
                return function (_pat_let33_0) {
                  return function (_987_dt__update_hcf73_h0) {
                    return _module.D16.create_DC38(_987_dt__update_hcf73_h0);
                  }(_pat_let33_0);
                }(_pat_let_tv31);
              }(_pat_let32_0);
            }(_module.D16.create_DC38(_dafny.MultiSet.FromArray(_dafny.Seq.of((_this).f6))))).dtor_cf73).cardinality()));
            _nw161[(new BigNumber(17)).toNumber()] = (_this).f6;
            _nw161[(new BigNumber(18)).toNumber()] = (_this).f6;
            _nw161[(new BigNumber(19)).toNumber()] = new BigNumber((_dafny.Seq.Concat(_978_v105, _dafny.Seq.Create(_module.__default.abs(new BigNumber(235)), function (_988_i18) {
              return (_this).f6;
            }))).length);
            _nw161[(new BigNumber(20)).toNumber()] = _module.__default.safeDivisionInt((_this).f6, (_this).f6);
            _nw161[(new BigNumber(21)).toNumber()] = (_this).f6;
            _nw161[(new BigNumber(22)).toNumber()] = (_dafny.ZERO).minus((_this).f6);
            _985_v112 = _nw161;
            let _index139 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length));
            (_985_v112)[_index139] = ((_this).f6).minus((_dafny.ZERO).minus((_this).f6));
            let _index140 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length));
            (_985_v112)[_index140] = new BigNumber((_977_v104).cardinality());
            let _989_v113;
            let _init29 = ((_990_v102, _991_v40) => function (_992_i19) {
              return _dafny.Map.Empty.slice().updateUnsafe((_990_v102)[_module.__default.safeIndex(new BigNumber(37), new BigNumber((_990_v102).length))],_991_v40);
            })(_975_v102, _882_v40);
            let _nw162 = Array((new BigNumber(27)).toNumber());
            for (let _i0_29 = 0; _i0_29 < new BigNumber(_nw162.length); _i0_29++) {
              _nw162[_i0_29] = _init29(new BigNumber(_i0_29));
            }
            _989_v113 = _nw162;
            let _993_v114;
            _993_v114 = _module.D14.create_DC35(_989_v113);
            let _994_v115;
            _994_v115 = _dafny.Seq.of(_993_v114);
            if (!_dafny.areEqual(((false) ? (_994_v115) : (_994_v115)), _dafny.Seq.of(_993_v114))) {
              let _995_v117;
              _995_v117 = _dafny.Map.Empty.slice().updateUnsafe((_975_v102)[_module.__default.safeIndex(new BigNumber(37), new BigNumber((_975_v102).length))],(_975_v102)[_module.__default.safeIndex(new BigNumber(37), new BigNumber((_975_v102).length))]);
              (globalState).f1 = !(new BigNumber(37)).isEqualTo((_dafny.ZERO).minus((new BigNumber((function () {
                let _coll47 = new _dafny.Map();
                for (const _compr_47 of (_995_v117).Keys.Elements) {
                  let _996_v116 = _compr_47;
                  if ((_995_v117).contains(_996_v116)) {
                    _coll47.push([_996_v116,_882_v40]);
                  }
                }
                return _coll47;
              }()).length)).minus(new BigNumber(788))));
              _926_v72 = _926_v72;
              let _997_v118;
              _997_v118 = _dafny.Map.Empty.slice().updateUnsafe((_985_v112)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length))],_976_v103);
              _997_v118 = _997_v118;
              let _index141 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_923_v71).length));
              (_923_v71)[_index141] = _882_v40;
              let _index142 = _module.__default.safeIndex(new BigNumber(529), new BigNumber((_923_v71).length));
              (_923_v71)[_index142] = _882_v40;
              let _998_v120;
              _998_v120 = _module.D0.create_DC0(function () {
  let _coll48 = new _dafny.Map();
  for (const _compr_48 of _dafny.IntegerRange(new BigNumber(934), new BigNumber(861))) {
    let _999_v119 = _compr_48;
    if (((new BigNumber(934)).isLessThanOrEqualTo(_999_v119)) && ((_999_v119).isLessThan(new BigNumber(861)))) {
      _coll48.push([(_999_v119).plus((_985_v112)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length))]),_957_v92]);
    }
  }
  return _coll48;
}());
              let _1000_v121;
              let _init30 = ((_1001_v71) => function (_1002_i20) {
                return (_1001_v71)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_1001_v71).length))];
              })(_923_v71);
              let _nw163 = Array((new BigNumber(14)).toNumber());
              for (let _i0_30 = 0; _i0_30 < new BigNumber(_nw163.length); _i0_30++) {
                _nw163[_i0_30] = _init30(new BigNumber(_i0_30));
              }
              _1000_v121 = _nw163;
              let _1003_v122;
              _1003_v122 = _dafny.Seq.of(_1000_v121);
              let _index143 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length));
              (_985_v112)[_index143] = ((_dafny.ZERO).minus((_985_v112)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length))])).minus((_dafny.ZERO).minus(new BigNumber(((_module.__default.fm35((_985_v112)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length))], globalState)).Merge(_dafny.Map.Empty.slice().updateUnsafe((_985_v112)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length))],_module.__default.fm1(_998_v120, (_923_v71)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_923_v71).length))], new BigNumber((_1003_v122).length), (_923_v71)[_module.__default.safeIndex(new BigNumber(529), new BigNumber((_923_v71).length))], globalState)))).length)));
            } else {
              (globalState).f5 = _882_v40;
              _978_v105 = _978_v105;
              let _index144 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length));
              let _rhs139 = _882_v40;
              let _rhs140 = ((!((_this).f6).isEqualTo(new BigNumber((_926_v72).length))) ? ((_985_v112)[_module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length))]) : (((_this).f6).minus((_this).f6)));
              let _rhs141 = (_957_v92).IsProperSubsetOf((_957_v92).Union(_957_v92));
              let _rhs142 = new BigNumber(((_957_v92).Intersect(_957_v92)).length);
              let _rhs143 = ((_dafny.ZERO).minus((_this).f6)).plus((_this).f6);
              let _lhs102 = _985_v112;
              let _lhs103 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length));
              let _lhs104 = globalState;
              _882_v40 = _rhs139;
              _lhs102[_lhs103] = _rhs140;
              _882_v40 = _rhs141;
              r2 = _rhs142;
              _lhs104.f0 = _rhs143;
              (globalState).f2 = (_this).f6;
              let _index145 = _module.__default.safeIndex(new BigNumber(478), new BigNumber((_985_v112).length));
              (_985_v112)[_index145] = (_this).f6;
            }
            let _1004_v123;
            _1004_v123 = _dafny.Seq.of((_979_v106).dtor_cf11);
            _1004_v123 = _1004_v123;
          }
        }
      }
      r0 = (_this).f6;
      r1 = (_dafny.ZERO).minus((_this).f6);
      r2 = new BigNumber(-773);
      let _1005_v124;
      _1005_v124 = new _dafny.CodePoint('a'.codePointAt(0));
      r3 = _1005_v124;
      return [r0, r1, r2, r3];
    }
  };

  $module.C5 = class C5 {
    constructor () {
      this._tname = "_module.C5";
      this._f7 = false;
      this._f8 = new _dafny.CodePoint('D'.codePointAt(0));
    }
    _parentTraits() {
      return [];
    }
    __ctor(f7, f8) {
      let _this = this;
      (_this)._f7 = f7;
      (_this)._f8 = f8;
      return;
    }
    fm5(globalState) {
      let _this = this;
      return new BigNumber((function (_source16) {
        if (_source16.is_DC1) {
          let _1006___mcc_h0 = (_source16).cf1;
          let _1007___mcc_h1 = (_source16).cf2;
          let _1008___mcc_h2 = (_source16).cf3;
          let _1009___mcc_h3 = (_source16).cf4;
          let _1010_cf4 = _1009___mcc_h3;
          let _1011_cf3 = _1008___mcc_h2;
          let _1012_cf2 = _1007___mcc_h1;
          let _1013_cf1 = _1006___mcc_h0;
          return _dafny.Seq.Concat((_module.D1.create_DC4(_dafny.Seq.of(_1011_cf3, _1013_cf1))).dtor_cf11, _dafny.Seq.Create(_module.__default.abs(new BigNumber(581)), ((_1014_cf3) => function (_1015_i0) {
            return (_dafny.ZERO).minus(_1014_cf3);
          })(_1011_cf3)));
        } else if (_source16.is_DC2) {
          let _1016___mcc_h4 = (_source16).cf5;
          let _1017___mcc_h5 = (_source16).cf6;
          let _1018___mcc_h6 = (_source16).cf7;
          let _1019___mcc_h7 = (_source16).cf8;
          let _1020___mcc_h8 = (_source16).cf9;
          let _1021_cf9 = _1020___mcc_h8;
          let _1022_cf8 = _1019___mcc_h7;
          let _1023_cf7 = _1018___mcc_h6;
          let _1024_cf6 = _1017___mcc_h5;
          let _1025_cf5 = _1016___mcc_h4;
          return _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements((_this).f7)).length), _1025_cf5);
        } else if (_source16.is_DC0) {
          let _1026___mcc_h9 = (_source16).cf0;
          let _1027_cf0 = _1026___mcc_h9;
          return (_module.D1.create_DC4(_dafny.Seq.of(new BigNumber((_dafny.Set.fromElements((_this).f7)).length)))).dtor_cf11;
        } else {
          let _1028___mcc_h10 = (_source16).cf10;
          let _1029_cf10 = _1028___mcc_h10;
          return _dafny.Seq.Concat(_dafny.Seq.of((_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f7)).length)), new BigNumber(943), new BigNumber((_dafny.Map.Empty.slice().updateUnsafe((_this).f7,(_module.D0.create_DC1(new BigNumber((_dafny.Seq.UnicodeFromString("ejdijcrpi")).length), (_this).f7, new BigNumber(757), (_this).f7)).dtor_cf3)).length)), _dafny.Seq.of(new BigNumber(-20), new BigNumber((_dafny.Seq.of((_this).f7)).length)));
        }
      }(_module.D0.create_DC0(function () {
  let _coll49 = new _dafny.Map();
  for (const _compr_49 of (_dafny.Set.fromElements(new BigNumber(714))).Elements) {
    let _1030_v0 = _compr_49;
    if ((_dafny.Set.fromElements(new BigNumber(714))).contains(_1030_v0)) {
      _coll49.push([(_1030_v0).multipliedBy(new BigNumber(322)),_dafny.Set.fromElements((_this).f7)]);
    }
  }
  return _coll49;
}()))).length);
    };
    fm6(p0, p1, p2, globalState) {
      let _this = this;
      return new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(-724)), function (_1031_i0) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("rj"), _dafny.Seq.Create(_module.__default.abs(new BigNumber(846)), function (_1032_i1) {
          return (_this).f8;
        }));
      })).length);
    };
    m1(p0, globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = _dafny.ZERO;
      let r2 = false;
      let r3 = _dafny.ZERO;
      let _1033_v0;
      let _nw164 = Array((new BigNumber(12)).toNumber()).fill(_dafny.ZERO);
      _1033_v0 = _nw164;
      _1033_v0 = _1033_v0;
      let _1034_v1;
      let _nw165 = Array((new BigNumber(17)).toNumber()).fill(false);
      _1034_v1 = _nw165;
      let _1035_v2;
      _1035_v2 = new BigNumber(441);
      let _1036_v3;
      _1036_v3 = _dafny.Map.Empty.slice().updateUnsafe(_1035_v2,_1034_v1);
      let _1037_v4;
      _1037_v4 = _dafny.Set.fromElements(_1035_v2, _1035_v2);
      let _1038_v5;
      _1038_v5 = _dafny.Seq.of(_1035_v2, _1035_v2, new BigNumber((_1037_v4).length));
      let _1039_v6;
      let _nw166 = Array((new BigNumber(18)).toNumber());
      _nw166[(_dafny.ZERO).toNumber()] = _1034_v1;
      _nw166[(_dafny.ONE).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(2)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(3)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(4)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(5)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(6)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(7)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(8)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(9)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(10)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(11)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(12)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(13)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(14)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(15)).toNumber()] = (((_1036_v3).contains((_1038_v5)[_module.__default.safeIndex(_1035_v2, new BigNumber((_1038_v5).length))])) ? ((_1036_v3).get((_1038_v5)[_module.__default.safeIndex(_1035_v2, new BigNumber((_1038_v5).length))])) : (_1034_v1));
      _nw166[(new BigNumber(16)).toNumber()] = _1034_v1;
      _nw166[(new BigNumber(17)).toNumber()] = _1034_v1;
      _1039_v6 = _nw166;
      let _index146 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_1039_v6).length));
      (_1039_v6)[_index146] = _1034_v1;
      let _1040_v7;
      _1040_v7 = _module.D1.create_DC4(_module.__default.fm7(globalState));
      let _index147 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1033_v0).length));
      (_1033_v0)[_index147] = _1035_v2;
      let _1041_v8;
      _1041_v8 = _dafny.Seq.of(false);
      let _1042_v9;
      _1042_v9 = _1034_v1;
      let _1043_v10;
      _1043_v10 = _module.D1.create_DC5();
      let _pat_let_tv32 = _1035_v2;
      let _pat_let_tv33 = _1035_v2;
      let _pat_let_tv34 = _1035_v2;
      let _index148 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_1039_v6).length));
      let _index149 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1033_v0).length));
      let _rhs144 = (_1035_v2).plus(_1035_v2);
      let _rhs145 = ((_this).fm6(_1035_v2, p0, new BigNumber(816), globalState)).plus((_dafny.ZERO).minus(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(true,new BigNumber((_1041_v8).length))).length)));
      let _rhs146 = (_1042_v9);
      let _rhs147 = _module.D1.create_DC4(_1038_v5);
      let _rhs148 = function (_source17) {
        if (_source17.is_DC5) {
          return (_pat_let_tv32).minus(_pat_let_tv33);
        } else {
          let _1044___mcc_h0 = (_source17).cf11;
          let _1045_cf11 = _1044___mcc_h0;
          return _pat_let_tv34;
        }
      }(_1043_v10);
      let _lhs105 = globalState;
      let _lhs106 = _1039_v6;
      let _lhs107 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_1039_v6).length));
      let _lhs108 = _1033_v0;
      let _lhs109 = _module.__default.safeIndex(new BigNumber(425), new BigNumber((_1033_v0).length));
      r3 = _rhs144;
      _lhs105.f0 = _rhs145;
      _lhs106[_lhs107] = _rhs146;
      _1040_v7 = _rhs147;
      _lhs108[_lhs109] = _rhs148;
      let _1046_v11;
      let _nw167 = Array((new BigNumber(14)).toNumber()).fill(_dafny.Map.Empty);
      _1046_v11 = _nw167;
      let _1047_v12;
      _1047_v12 = _dafny.Map.Empty.slice().updateUnsafe(_1046_v11,_1033_v0);
      _1047_v12 = (_1047_v12).update(_1046_v11, _1033_v0);
      let _1048_v13;
      _1048_v13 = _dafny.Set.fromElements(p0);
      _1048_v13 = ((false) ? (_1048_v13) : (_1048_v13));
      let _1049_v14;
      let _nw168 = Array((new BigNumber(11)).toNumber()).fill(_dafny.Map.Empty);
      _1049_v14 = _nw168;
      let _1050_v15;
      _1050_v15 = _module.D0.create_DC1((_this).fm5(globalState), p0, (_1033_v0)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1033_v0).length))], p0);
      let _1051_v16;
      _1051_v16 = _module.D0.create_DC3(_1050_v15);
      let _1052_v17;
      _1052_v17 = _module.D0.create_DC3(_1051_v16);
      let _1053_v18;
      _1053_v18 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Map.Empty.slice().updateUnsafe((_1033_v0)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1033_v0).length))],(_this).f7),_1052_v17);
      let _index150 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_1049_v14).length));
      (_1049_v14)[_index150] = _1053_v18;
      let _1054_v19;
      _1054_v19 = new _dafny.CodePoint('d'.codePointAt(0));
      let _index151 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_1049_v14).length));
      let _rhs149 = _1053_v18;
      let _rhs150 = _module.__default.fm8((_this).f8, ((p0) ? (p0) : ((_this).f7)), (_1033_v0)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1033_v0).length))], (_1033_v0)[_module.__default.safeIndex(new BigNumber(425), new BigNumber((_1033_v0).length))], globalState);
      let _lhs110 = _1049_v14;
      let _lhs111 = _module.__default.safeIndex(new BigNumber(711), new BigNumber((_1049_v14).length));
      _lhs110[_lhs111] = _rhs149;
      _1054_v19 = _rhs150;
      let _index152 = _module.__default.safeIndex(new BigNumber(163), new BigNumber((_1039_v6).length));
      (_1039_v6)[_index152] = (_1039_v6)[_module.__default.safeIndex(new BigNumber(163), new BigNumber((_1039_v6).length))];
      r0 = _1035_v2;
      let _1055_v21;
      _1055_v21 = _dafny.MultiSet.fromElements((_this).f7);
      let _1056_v22;
      _1056_v22 = _dafny.Map.Empty.slice().updateUnsafe(_1055_v21,_1035_v2);
      let _1057_v24;
      _1057_v24 = _dafny.Seq.of((((_this).f7) ? (_module.__default.fm9(globalState)) : (function () {
        let _coll50 = new _dafny.Set();
        for (const _compr_50 of (function () {
          let _coll51 = new _dafny.Map();
          for (const _compr_51 of (_1056_v22).Keys.Elements) {
            let _1058_v20 = _compr_51;
            if ((_1056_v22).contains(_1058_v20)) {
              _coll51.push([_1058_v20,_1041_v8]);
            }
          }
          return _coll51;
        }()).Keys.Elements) {
          let _1059_v23 = _compr_50;
          if ((function () {
            let _coll52 = new _dafny.Map();
            for (const _compr_52 of (_1056_v22).Keys.Elements) {
              let _1058_v20 = _compr_52;
              if ((_1056_v22).contains(_1058_v20)) {
                _coll52.push([_1058_v20,_1041_v8]);
              }
            }
            return _coll52;
          }()).contains(_1059_v23)) {
            _coll50.add(_1059_v23);
          }
        }
        return _coll50;
      }())));
      r1 = new BigNumber(((_1057_v24)[_module.__default.safeIndex(new BigNumber(682), new BigNumber((_1057_v24).length))]).length);
      r2 = p0;
      r3 = _1035_v2;
      return [r0, r1, r2, r3];
    }
    m2(p0, p1, p2, p3, globalState) {
      let _this = this;
      let r0 = false;
      let _1060_v0;
      _1060_v0 = _dafny.Map.Empty.slice().updateUnsafe(p0,_dafny.Set.fromElements((_this).f7));
      let _1061_v1;
      _1061_v1 = _dafny.Map.Empty.slice().updateUnsafe(p1,_1060_v0);
      let _1062_v2;
      _1062_v2 = _module.D0.create_DC0((((_1061_v1).contains(p2)) ? ((_1061_v1).get(p2)) : (_1060_v0)));
      let _pat_let_tv35 = p0;
      let _pat_let_tv36 = p3;
      let _pat_let_tv37 = p0;
      let _pat_let_tv38 = p2;
      let _pat_let_tv39 = p0;
      let _pat_let_tv40 = p0;
      let _pat_let_tv41 = p1;
      let _pat_let_tv42 = p0;
      let _pat_let_tv43 = p0;
      let _source18 = function (_source19) {
        if (_source19.is_DC1) {
          let _1063___mcc_h11 = (_source19).cf1;
          let _1064___mcc_h12 = (_source19).cf2;
          let _1065___mcc_h13 = (_source19).cf3;
          let _1066___mcc_h14 = (_source19).cf4;
          let _1067_cf4 = _1066___mcc_h14;
          let _1068_cf3 = _1065___mcc_h13;
          let _1069_cf2 = _1064___mcc_h12;
          let _1070_cf1 = _1063___mcc_h11;
          return _module.D0.create_DC1(_pat_let_tv35, true, new BigNumber((_pat_let_tv36).length), _1069_cf2);
        } else if (_source19.is_DC2) {
          let _1071___mcc_h15 = (_source19).cf5;
          let _1072___mcc_h16 = (_source19).cf6;
          let _1073___mcc_h17 = (_source19).cf7;
          let _1074___mcc_h18 = (_source19).cf8;
          let _1075___mcc_h19 = (_source19).cf9;
          let _1076_cf9 = _1075___mcc_h19;
          let _1077_cf8 = _1074___mcc_h18;
          let _1078_cf7 = _1073___mcc_h17;
          let _1079_cf6 = _1072___mcc_h16;
          let _1080_cf5 = _1071___mcc_h15;
          return _module.D0.create_DC1(_1077_cf8, (_this).f7, _1077_cf8, (_this).f7);
        } else if (_source19.is_DC0) {
          let _1081___mcc_h20 = (_source19).cf0;
          let _1082_cf0 = _1081___mcc_h20;
          return _module.D0.create_DC1(new BigNumber((_dafny.Seq.of(_pat_let_tv37)).length), _pat_let_tv38, new BigNumber((function () {
  let _coll53 = new _dafny.Map();
  for (const _compr_53 of (_dafny.Seq.of(_pat_let_tv39)).Elements) {
    let _1083_v3 = _compr_53;
    if (_dafny.Seq.contains(_dafny.Seq.of(_pat_let_tv40), _1083_v3)) {
      _coll53.push([(_1083_v3).multipliedBy(new BigNumber((_dafny.Set.fromElements((_this).f7, (_this).f7)).length)),_pat_let_tv41]);
    }
  }
  return _coll53;
}()).length), (_this).f7);
        } else {
          let _1084___mcc_h21 = (_source19).cf10;
          let _1085_cf10 = _1084___mcc_h21;
          return _module.D0.create_DC1(_pat_let_tv42, (_this).f7, _pat_let_tv43, (_this).f7);
        }
      }(_1062_v2);
      if (_source18.is_DC1) {
        let _1086___mcc_h0 = (_source18).cf1;
        let _1087___mcc_h1 = (_source18).cf2;
        let _1088___mcc_h2 = (_source18).cf3;
        let _1089___mcc_h3 = (_source18).cf4;
        let _1090_cf4 = _1089___mcc_h3;
        let _1091_cf3 = _1088___mcc_h2;
        let _1092_cf2 = _1087___mcc_h1;
        let _1093_cf1 = _1086___mcc_h0;
        let _1094_v4;
        _1094_v4 = _dafny.MultiSet.fromElements(_1090_cf4);
        (globalState).f2 = ((_1091_cf3).minus(new BigNumber((_1094_v4).cardinality()))).minus(p0);
        let _1095_v5;
        _1095_v5 = _dafny.Map.Empty.slice().updateUnsafe((((_this).f7) ? (!(p2)) : (_1092_cf2)),p1);
        _1095_v5 = _1095_v5;
        let _1096_v6;
        _1096_v6 = _dafny.Seq.UnicodeFromString("ja");
        _1096_v6 = _dafny.Seq.Concat(_dafny.Seq.Create(_module.__default.abs(new BigNumber(129)), function (_1097_i0) {
          return (_this).f8;
        }), _dafny.Seq.Concat(_1096_v6, _1096_v6));
        let _1098_v7;
        _1098_v7 = _dafny.Set.fromElements(_1093_cf1);
        _1095_v5 = (_1095_v5).update(_1090_cf4, (_1098_v7).IsDisjointFrom(_dafny.Set.fromElements(_1091_cf3)));
      } else if (_source18.is_DC2) {
        let _1099___mcc_h4 = (_source18).cf5;
        let _1100___mcc_h5 = (_source18).cf6;
        let _1101___mcc_h6 = (_source18).cf7;
        let _1102___mcc_h7 = (_source18).cf8;
        let _1103___mcc_h8 = (_source18).cf9;
        let _1104_cf9 = _1103___mcc_h8;
        let _1105_cf8 = _1102___mcc_h7;
        let _1106_cf7 = _1101___mcc_h6;
        let _1107_cf6 = _1100___mcc_h5;
        let _1108_cf5 = _1099___mcc_h4;
        let _1109_v8;
        let _nw169 = Array((new BigNumber(26)).toNumber()).fill(false);
        _1109_v8 = _nw169;
        let _1110_v9;
        _1110_v9 = _1109_v8;
        let _source20 = (((_this).f7) ? (_1110_v9) : (_1109_v8));
        let _1111___mcc_h22 = _source20;
        let _1112_cf12 = _1111___mcc_h22;
        let _1113_v10;
        _1113_v10 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(687),p2);
        let _index153 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_1109_v8).length));
        (_1109_v8)[_index153] = (((_1113_v10).contains(_1104_cf9)) ? ((_1113_v10).get(_1104_cf9)) : (!(p1)));
        let _1114_v11;
        let _nw170 = Array((new BigNumber(5)).toNumber()).fill([]);
        _1114_v11 = _nw170;
        let _index154 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1114_v11).length));
        (_1114_v11)[_index154] = _1112_cf12;
        let _1115_v12;
        _1115_v12 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,p1);
        let _1116_v13;
        _1116_v13 = _dafny.Map.Empty.slice().updateUnsafe((_this).fm5(globalState),_1104_cf9);
        let _index155 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_1109_v8).length));
        let _index156 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1114_v11).length));
        let _rhs151 = new BigNumber((_1115_v12).length);
        let _rhs152 = p2;
        let _rhs153 = p1;
        let _rhs154 = _1112_cf12;
        let _rhs155 = ((((((_1116_v13).contains(new BigNumber(82))) ? ((_1116_v13).get(new BigNumber(82))) : (new BigNumber((_dafny.Seq.UnicodeFromString("qng")).length)))).isLessThan(_1106_cf7)) ? (p1) : ((_module.D0.create_DC1(_1106_cf7, _module.__default.fm1(_1062_v2, !(p2), (_this).fm5(globalState), p2, globalState), _1105_cf8, (_this).f7)).dtor_cf4));
        let _lhs112 = globalState;
        let _lhs113 = _1109_v8;
        let _lhs114 = _module.__default.safeIndex(new BigNumber(161), new BigNumber((_1109_v8).length));
        let _lhs115 = globalState;
        let _lhs116 = _1114_v11;
        let _lhs117 = _module.__default.safeIndex(new BigNumber(205), new BigNumber((_1114_v11).length));
        let _lhs118 = globalState;
        _lhs112.f0 = _rhs151;
        _lhs113[_lhs114] = _rhs152;
        _lhs115.f1 = _rhs153;
        _lhs116[_lhs117] = _rhs154;
        _lhs118.f1 = _rhs155;
        let _1117_v14;
        let _1118_v15;
        let _1119_v16;
        let _1120_v17;
        let _out4;
        let _out5;
        let _out6;
        let _out7;
        let _outcollector1 = (_this).m1(p1, globalState);
        _out4 = _outcollector1[0];
        _out5 = _outcollector1[1];
        _out6 = _outcollector1[2];
        _out7 = _outcollector1[3];
        _1117_v14 = _out4;
        _1118_v15 = _out5;
        _1119_v16 = _out6;
        _1120_v17 = _out7;
        _1118_v15 = _1108_cf5;
        let _1121_v18;
        let _init31 = ((_1122_cf7) => function (_1123_i1) {
          return _module.__default.safeModuloInt(_1123_i1, _1122_cf7);
        })(_1106_cf7);
        let _nw171 = Array((new BigNumber(18)).toNumber());
        for (let _i0_31 = 0; _i0_31 < new BigNumber(_nw171.length); _i0_31++) {
          _nw171[_i0_31] = _init31(new BigNumber(_i0_31));
        }
        _1121_v18 = _nw171;
        let _index157 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1121_v18).length));
        (_1121_v18)[_index157] = p0;
        let _1124_v19;
        _1124_v19 = _dafny.Seq.UnicodeFromString("hjhrdvw");
        let _index158 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1121_v18).length));
        let _rhs156 = _1119_v16;
        let _rhs157 = _1117_v14;
        let _rhs158 = (_dafny.ZERO).minus(new BigNumber((_1124_v19).length));
        let _rhs159 = (_this).f7;
        let _rhs160 = _1104_cf9;
        let _lhs119 = globalState;
        let _lhs120 = _1121_v18;
        let _lhs121 = _module.__default.safeIndex(new BigNumber(690), new BigNumber((_1121_v18).length));
        let _lhs122 = globalState;
        let _lhs123 = globalState;
        _lhs119.f1 = _rhs156;
        _1105_cf8 = _rhs157;
        _lhs120[_lhs121] = _rhs158;
        _lhs122.f5 = _rhs159;
        _lhs123.f0 = _rhs160;
        _1105_cf8 = (_1108_cf5).multipliedBy(new BigNumber(892));
        let _1125_v20;
        _1125_v20 = _dafny.Seq.UnicodeFromString("qtau");
        let _1126_v21;
        _1126_v21 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1125_v20);
        let _1127_v22;
        _1127_v22 = _dafny.MultiSet.fromElements((((_1126_v21).contains((_this).f7)) ? ((_1126_v21).get((_this).f7)) : (_1125_v20)), _1125_v20);
        _1127_v22 = (_1127_v22).Union(_dafny.MultiSet.fromElements(_1125_v20));
        let _index159 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_1109_v8).length));
        (_1109_v8)[_index159] = p2;
        let _index160 = _module.__default.safeIndex(new BigNumber(499), new BigNumber((_1109_v8).length));
        (_1109_v8)[_index160] = (_this).f7;
      } else if (_source18.is_DC0) {
        let _1128___mcc_h9 = (_source18).cf0;
        let _1129_cf0 = _1128___mcc_h9;
        let _1130_v23;
        let _nw172 = Array((new BigNumber(7)).toNumber());
        _nw172[(_dafny.ZERO).toNumber()] = p0;
        _nw172[(_dafny.ONE).toNumber()] = p0;
        _nw172[(new BigNumber(2)).toNumber()] = p0;
        _nw172[(new BigNumber(3)).toNumber()] = p0;
        _nw172[(new BigNumber(4)).toNumber()] = new BigNumber(813);
        _nw172[(new BigNumber(5)).toNumber()] = p0;
        _nw172[(new BigNumber(6)).toNumber()] = p0;
        _1130_v23 = _nw172;
        let _index161 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length));
        (_1130_v23)[_index161] = (_dafny.ZERO).minus(p0);
        let _1131_v24;
        _1131_v24 = _module.D0.create_DC1(p0, p2, p0, (_this).f7);
        let _index162 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length));
        let _rhs161 = (_1131_v24).dtor_cf2;
        let _rhs162 = p0;
        let _lhs124 = globalState;
        let _lhs125 = _1130_v23;
        let _lhs126 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length));
        _lhs124.f5 = _rhs161;
        _lhs125[_lhs126] = _rhs162;
        let _1132_v25;
        let _nw173 = Array((_dafny.ONE).toNumber());
        _1132_v25 = _nw173;
        let _1133_v26;
        _1133_v26 = _dafny.Seq.of(_module.__default.fm1(_module.D0.create_DC0(_1060_v0), p1, new BigNumber(-909), p1, globalState), ((!(p2)) ? (p1) : (false)));
        let _rhs163 = _1132_v25;
        let _rhs164 = _dafny.Seq.of(p2);
        let _rhs165 = p0;
        let _lhs127 = globalState;
        _1132_v25 = _rhs163;
        _1133_v26 = _rhs164;
        _lhs127.f2 = _rhs165;
        let _1134_v27;
        _1134_v27 = _dafny.Set.fromElements(_module.__default.safeModuloInt(p0, (_1130_v23)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length))]));
        _1134_v27 = _1134_v27;
        let _1135_v28;
        _1135_v28 = _dafny.MultiSet.fromElements((_1130_v23)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length))], (_1130_v23)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length))], ((!(p2)) ? ((_1130_v23)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length))]) : (new BigNumber((_dafny.Set.fromElements(_1134_v27, _1134_v27, _1134_v27, _1134_v27, _module.__default.fm10(new BigNumber(802), p1, p2, globalState))).length))), (_this).fm5(globalState));
        let _1136_v29;
        let _nw174 = Array((new BigNumber(26)).toNumber()).fill(false);
        _1136_v29 = _nw174;
        let _index163 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length));
        let _rhs166 = (_1130_v23)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length))];
        let _rhs167 = (_module.__default.fm1(_1062_v2, false, p0, (_this).f7, globalState)) || ((_this).f7);
        let _rhs168 = _dafny.MultiSet.fromElements(p0, (_1130_v23)[_module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length))]);
        let _rhs169 = _1136_v29;
        let _lhs128 = _1130_v23;
        let _lhs129 = _module.__default.safeIndex(new BigNumber(708), new BigNumber((_1130_v23).length));
        let _lhs130 = globalState;
        _lhs128[_lhs129] = _rhs166;
        _lhs130.f5 = _rhs167;
        _1135_v28 = _rhs168;
        _1136_v29 = _rhs169;
      } else {
        let _1137___mcc_h10 = (_source18).cf10;
        let _1138_cf10 = _1137___mcc_h10;
        let _1139_v30;
        _1139_v30 = _dafny.Map.Empty.slice().updateUnsafe(p1,(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(575)), function (_1140_i2) {
          return (_this).f8;
        })).length)).minus(p0));
        _1139_v30 = (_1139_v30).update(!(p1), (_this).fm5(globalState));
        let _1141_v31;
        _1141_v31 = _dafny.Seq.UnicodeFromString("xp");
        let _1142_v32;
        _1142_v32 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber((_1141_v31).length)).multipliedBy(new BigNumber((_1139_v30).length)),!_dafny.areEqual(p3, p3));
        _1142_v32 = (_1142_v32).update(p0, (new BigNumber(723)).isLessThan(p0));
        let _1143_v33;
        _1143_v33 = _dafny.Map.Empty.slice().updateUnsafe(p2,false);
        _1143_v33 = (_1143_v33).update((p0).isEqualTo(new BigNumber((_1141_v31).length)), p1);
        let _1144_v34;
        let _nw175 = Array((new BigNumber(23)).toNumber()).fill([]);
        _1144_v34 = _nw175;
        let _1145_v35;
        let _nw176 = Array((_dafny.ONE).toNumber()).fill(false);
        _1145_v35 = _nw176;
        let _index164 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_1144_v34).length));
        (_1144_v34)[_index164] = _1145_v35;
        let _index165 = _module.__default.safeIndex(new BigNumber(93), new BigNumber((_1144_v34).length));
        (_1144_v34)[_index165] = _1145_v35;
      }
      if ((_this).f7) {
        (globalState).f0 = (new BigNumber(701)).plus(p0);
        let _1146_v36;
        _1146_v36 = _dafny.Seq.of(p3);
        (globalState).f5 = _dafny.Seq.IsPrefixOf(_dafny.Seq.Concat(_1146_v36, _1146_v36), _1146_v36);
        _1146_v36 = _dafny.Seq.Concat(_1146_v36, _dafny.Seq.Create(_module.__default.abs(new BigNumber(966)), ((_1147_p3) => function (_1148_i3) {
          return _1147_p3;
        })(p3)));
        let _1149_v37;
        let _nw177 = Array((new BigNumber(24)).toNumber()).fill(_dafny.ZERO);
        _1149_v37 = _nw177;
        let _index166 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_1149_v37).length));
        (_1149_v37)[_index166] = p0;
        let _1150_v38;
        _1150_v38 = _dafny.Map.Empty.slice().updateUnsafe(p1,p2);
        let _1151_v39;
        _1151_v39 = _module.D0.create_DC1(p0, p1, new BigNumber(883), (((_1150_v38).contains(p2)) ? ((_1150_v38).get(p2)) : (!((_this).f7))));
        let _pat_let_tv44 = p0;
        let _pat_let_tv45 = p0;
        let _index167 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_1149_v37).length));
        let _rhs170 = new BigNumber(721);
        let _rhs171 = (((p1) ? ((((_1150_v38).contains((_this).f7)) ? ((_1150_v38).get((_this).f7)) : (p1))) : ((function (_pat_let34_0) {
          return function (_1152_dt__update__tmp_h0) {
            return function (_pat_let35_0) {
              return function (_1153_dt__update_hcf3_h0) {
                return function (_pat_let36_0) {
                  return function (_1154_dt__update_hcf1_h0) {
                    return function (_pat_let37_0) {
                      return function (_1155_dt__update_hcf2_h0) {
                        return _module.D0.create_DC1(_1154_dt__update_hcf1_h0, _1155_dt__update_hcf2_h0, _1153_dt__update_hcf3_h0, (_1152_dt__update__tmp_h0).dtor_cf4);
                      }(_pat_let37_0);
                    }((_this).f7);
                  }(_pat_let36_0);
                }(_pat_let_tv45);
              }(_pat_let35_0);
            }(_pat_let_tv44);
          }(_pat_let34_0);
        }(_1151_v39)).dtor_cf2))) && ((_this).f7);
        let _lhs131 = _1149_v37;
        let _lhs132 = _module.__default.safeIndex(new BigNumber(612), new BigNumber((_1149_v37).length));
        let _lhs133 = globalState;
        _lhs131[_lhs132] = _rhs170;
        _lhs133.f5 = _rhs171;
        let _1156_v40;
        _1156_v40 = _dafny.Seq.UnicodeFromString("t");
        let _1157_v41;
        _1157_v41 = _dafny.Seq.of(new BigNumber((_dafny.Set.fromElements(new BigNumber((_1156_v40).length))).length));
        let _1158_v42;
        _1158_v42 = _dafny.Set.fromElements(p1, false);
        let _1159_v43;
        _1159_v43 = _dafny.Map.Empty.slice().updateUnsafe(p2,new BigNumber(((_dafny.MultiSet.fromElements(new BigNumber((_1157_v41).length), new BigNumber((_1158_v42).length))).update(new BigNumber(261), _module.__default.abs((_1149_v37)[_module.__default.safeIndex(new BigNumber(612), new BigNumber((_1149_v37).length))]))).cardinality()));
        let _1160_v44;
        let _nw178 = new _module.C4();
        _nw178.__ctor((((_1159_v43).contains(p2)) ? ((_1159_v43).get(p2)) : (p0)));
        _1160_v44 = _nw178;
        let _1161_v45;
        _1161_v45 = _dafny.Map.Empty.slice().updateUnsafe(p2,(_this).f8);
        let _1162_v46;
        _1162_v46 = _dafny.Map.Empty.slice().updateUnsafe(_1160_v44,_1161_v45);
        (globalState).f2 = new BigNumber(((_1162_v46).Merge((_dafny.Map.Empty.slice().updateUnsafe(_1160_v44,_1161_v45)).Merge(_1162_v46))).length);
      } else {
        let _1163_v47;
        _1163_v47 = _module.D13.create_DC32(p0, (_this).f7, p1, p1, p2);
        let _1164_v48;
        _1164_v48 = _module.D13.create_DC34(_1163_v47);
        let _1165_v49;
        _1165_v49 = _dafny.Map.Empty.slice().updateUnsafe((_this).f7,_1164_v48);
        _1165_v49 = _1165_v49;
        let _1166_v50;
        let _init32 = ((_1167_p1) => function (_1168_i4) {
          return _1167_p1;
        })(p1);
        let _nw179 = Array((new BigNumber(12)).toNumber());
        for (let _i0_32 = 0; _i0_32 < new BigNumber(_nw179.length); _i0_32++) {
          _nw179[_i0_32] = _init32(new BigNumber(_i0_32));
        }
        _1166_v50 = _nw179;
        let _1169_v51;
        _1169_v51 = _dafny.Map.Empty.slice().updateUnsafe(p0,p0);
        let _1170_v52;
        _1170_v52 = _dafny.Seq.UnicodeFromString("gomydnwq");
        let _1171_v53;
        let _nw180 = new _module.C3();
        _nw180.__ctor(p0, _dafny.Seq.of(new BigNumber((_1169_v51).length), p0, new BigNumber((_1170_v52).length), p0, p0), new BigNumber((p3).length));
        _1171_v53 = _nw180;
        let _1172_v54;
        _1172_v54 = _module.D11.create_DC25(_1166_v50, _1171_v53, (p0).multipliedBy(p0));
        _1172_v54 = _1172_v54;
        let _1173_v55;
        _1173_v55 = _dafny.Set.fromElements(p1, false, (_this).f7, p1, false);
        let _1174_v56;
        _1174_v56 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber(351),!((_this).f7));
        (globalState).f5 = (_module.D14.create_DC36((_this).f7, new BigNumber((_1173_v55).length), new BigNumber((_1170_v52).length), _1174_v56)).dtor_cf68;
        (globalState).f0 = new BigNumber(909);
        let _1175_v57;
        _1175_v57 = _dafny.Seq.of(p0);
        let _1176_v58;
        let _nw181 = Array((new BigNumber(4)).toNumber()).fill(_dafny.ZERO);
        _1176_v58 = _nw181;
        let _1177_v59;
        _1177_v59 = _dafny.Map.Empty.slice().updateUnsafe(_1175_v57,(_module.D3.create_DC8(_module.D1.create_DC4(_1175_v57), _1176_v58, p0)).dtor_cf15);
        _1177_v59 = _1177_v59;
      }
      let _1178_v60;
      let _nw182 = Array((new BigNumber(4)).toNumber()).fill(_module.D6.Default());
      _1178_v60 = _nw182;
      for (const _guard_loop_4 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1178_v60).length))) {
        let _1179_i5 = _guard_loop_4;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1179_i5)) && ((_1179_i5).isLessThan(new BigNumber((_1178_v60).length))))) {
          (_1178_v60)[(_1179_i5)] = _module.D6.create_DC15(_dafny.Set.fromElements(p2));
        }
      }
      let _1180_v61;
      _1180_v61 = _dafny.Map.Empty.slice().updateUnsafe(p0,(_this).f8);
      _1180_v61 = (_1180_v61).update(p0, (_this).f8);
      let _1181_v62;
      let _init33 = function (_1182_i7) {
        return _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("xbox"), _dafny.Seq.UnicodeFromString("rvl"));
      };
      let _nw183 = Array((new BigNumber(29)).toNumber());
      for (let _i0_33 = 0; _i0_33 < new BigNumber(_nw183.length); _i0_33++) {
        _nw183[_i0_33] = _init33(new BigNumber(_i0_33));
      }
      _1181_v62 = _nw183;
      for (const _guard_loop_5 of _dafny.IntegerRange(_dafny.ZERO, new BigNumber((_1181_v62).length))) {
        let _1183_i6 = _guard_loop_5;
        if ((true) && (((_dafny.ZERO).isLessThanOrEqualTo(_1183_i6)) && ((_1183_i6).isLessThan(new BigNumber((_1181_v62).length))))) {
          (_1181_v62)[(_1183_i6)] = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("jkqigycsa"), _dafny.Seq.UnicodeFromString("ysjgmd"));
        }
      }
      let _1184_v63;
      _1184_v63 = _module.D13.create_DC32(_module.__default.safeModuloInt(new BigNumber(975), p0), p1, (_this).f7, p2, p1);
      let _source21 = _1184_v63;
      if (_source21.is_DC32) {
        let _1185___mcc_h23 = (_source21).cf56;
        let _1186___mcc_h24 = (_source21).cf57;
        let _1187___mcc_h25 = (_source21).cf58;
        let _1188___mcc_h26 = (_source21).cf59;
        let _1189___mcc_h27 = (_source21).cf60;
        let _1190_cf60 = _1189___mcc_h27;
        let _1191_cf59 = _1188___mcc_h26;
        let _1192_cf58 = _1187___mcc_h25;
        let _1193_cf57 = _1186___mcc_h24;
        let _1194_cf56 = _1185___mcc_h23;
        let _1195_v64;
        _1195_v64 = _module.D8.create_DC21(_1191_cf59, _1192_cf58, new BigNumber((_dafny.Seq.update(p3, _module.__default.safeIndex(p0, new BigNumber((p3).length)), _1192_cf58)).length));
        (globalState).f5 = (_1195_v64).dtor_cf36;
        let _1196_v65;
        let _init34 = function (_1197_i8) {
          return _module.__default.safeDivisionInt(_1197_i8, new BigNumber(630));
        };
        let _nw184 = Array((new BigNumber(19)).toNumber());
        for (let _i0_34 = 0; _i0_34 < new BigNumber(_nw184.length); _i0_34++) {
          _nw184[_i0_34] = _init34(new BigNumber(_i0_34));
        }
        _1196_v65 = _nw184;
        let _index168 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_1196_v65).length));
        (_1196_v65)[_index168] = new BigNumber(-844);
        let _1198_v67;
        let _init35 = ((_1199_p2, _1200_cf56) => function (_1201_i9) {
          return _dafny.Seq.Concat(_dafny.Seq.of(new BigNumber((_dafny.Seq.UnicodeFromString("kxvatj")).length), new BigNumber((function () {
            let _coll54 = new _dafny.Set();
            for (const _compr_54 of (_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_1199_p2),_1200_cf56)).length))).Elements) {
              let _1202_v66 = _compr_54;
              if ((_dafny.MultiSet.fromElements(new BigNumber((_dafny.Map.Empty.slice().updateUnsafe(!(_1199_p2),_1200_cf56)).length))).contains(_1202_v66)) {
                _coll54.add((_1202_v66).plus(new BigNumber(372)));
              }
            }
            return _coll54;
          }()).length)), _dafny.Seq.of(_1200_cf56));
        })(p2, _1194_cf56);
        let _nw185 = Array((new BigNumber(6)).toNumber());
        for (let _i0_35 = 0; _i0_35 < new BigNumber(_nw185.length); _i0_35++) {
          _nw185[_i0_35] = _init35(new BigNumber(_i0_35));
        }
        _1198_v67 = _nw185;
        let _1203_v68;
        _1203_v68 = _dafny.Seq.UnicodeFromString("gaq");
        let _index169 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_1196_v65).length));
        let _rhs172 = _1194_cf56;
        let _rhs173 = !(true);
        let _rhs174 = _1198_v67;
        let _rhs175 = _dafny.Seq.Concat(_dafny.Seq.UnicodeFromString("eakxl"), _dafny.Seq.Concat(_1203_v68, _1203_v68));
        let _rhs176 = (_this).f7;
        let _lhs134 = _1196_v65;
        let _lhs135 = _module.__default.safeIndex(new BigNumber(335), new BigNumber((_1196_v65).length));
        let _lhs136 = globalState;
        _lhs134[_lhs135] = _rhs172;
        r0 = _rhs173;
        _1198_v67 = _rhs174;
        _1203_v68 = _rhs175;
        _lhs136.f5 = _rhs176;
        let _1204_v69;
        _1204_v69 = _dafny.Map.Empty.slice().updateUnsafe(_1196_v65,true);
        let _1205_v70;
        let _1206_v71;
        let _1207_v72;
        let _1208_v73;
        let _out8;
        let _out9;
        let _out10;
        let _out11;
        let _outcollector2 = (_this).m1(!(_module.__default.fm1(_1062_v2, (p3)[_module.__default.safeIndex(_1194_cf56, new BigNumber((p3).length))], (_dafny.ZERO).minus(new BigNumber((_1204_v69).length)), false, globalState)), globalState);
        _out8 = _outcollector2[0];
        _out9 = _outcollector2[1];
        _out10 = _outcollector2[2];
        _out11 = _outcollector2[3];
        _1205_v70 = _out8;
        _1206_v71 = _out9;
        _1207_v72 = _out10;
        _1208_v73 = _out11;
        _1196_v65 = _1196_v65;
      } else if (_source21.is_DC33) {
        let _1209___mcc_h28 = (_source21).cf61;
        let _1210___mcc_h29 = (_source21).cf62;
        let _1211___mcc_h30 = (_source21).cf63;
        let _1212___mcc_h31 = (_source21).cf64;
        let _1213___mcc_h32 = (_source21).cf65;
        let _1214_cf65 = _1213___mcc_h32;
        let _1215_cf64 = _1212___mcc_h31;
        let _1216_cf63 = _1211___mcc_h30;
        let _1217_cf62 = _1210___mcc_h29;
        let _1218_cf61 = _1209___mcc_h28;
        let _1219_v74;
        _1219_v74 = new _dafny.CodePoint('u'.codePointAt(0));
        _1219_v74 = _1219_v74;
        _1216_cf63 = _module.__default.safeDivisionInt(_1217_cf62, new BigNumber(566));
        let _1220_v75;
        let _nw186 = Array((new BigNumber(8)).toNumber());
        _1220_v75 = _nw186;
        let _1221_v76;
        _1221_v76 = _dafny.Seq.UnicodeFromString("n");
        let _1222_v77;
        let _nw187 = new _module.C2();
        _nw187.__ctor(_1221_v76, _1218_cf61);
        _1222_v77 = _nw187;
        let _index170 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_1220_v75).length));
        (_1220_v75)[_index170] = _1222_v77;
        let _1223_v78;
        _1223_v78 = _dafny.MultiSet.fromElements((_this).f7);
        let _1224_v79;
        _1224_v79 = _dafny.Seq.of(_1222_v77, _1222_v77);
        let _index171 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_1220_v75).length));
        let _rhs177 = (_this).f8;
        let _rhs178 = _1222_v77;
        let _rhs179 = (_1224_v79)[_module.__default.safeIndex(((_1214_cf65).f9).plus(new BigNumber((_dafny.Set.fromElements((_1214_cf65).f9, p0)).length)), new BigNumber((_1224_v79).length))];
        let _rhs180 = (_1223_v78).Intersect(_1223_v78);
        let _lhs137 = _1220_v75;
        let _lhs138 = _module.__default.safeIndex(new BigNumber(356), new BigNumber((_1220_v75).length));
        _1219_v74 = _rhs177;
        _lhs137[_lhs138] = _rhs178;
        _1222_v77 = _rhs179;
        _1223_v78 = _rhs180;
        if (!(_1215_cf64)) {
          let _1225_v80;
          let _nw188 = new _module.C4();
          _nw188.__ctor((_1214_cf65).f9);
          _1225_v80 = _nw188;
          let _1226_v81;
          _1226_v81 = _dafny.Set.fromElements(_1225_v80, _1225_v80);
          (globalState).f5 = (_1226_v81).IsSubsetOf(_1226_v81);
          let _1227_v82;
          _1227_v82 = _dafny.Map.Empty.slice().updateUnsafe(_1222_v77.f12,(_1222_v77).f11);
          let _1228_v83;
          _1228_v83 = _dafny.MultiSet.fromElements(_1214_cf65);
          let _1229_v84;
          _1229_v84 = _dafny.MultiSet.fromElements(new BigNumber((_1228_v83).cardinality()));
          let _1230_v85;
          _1230_v85 = _dafny.Set.fromElements(new BigNumber((_1229_v84).cardinality()), (_this).fm6(_1216_cf63, _1222_v77.f12, _1216_cf63, globalState), _1216_cf63);
          let _1231_v86;
          let _nw189 = Array((new BigNumber(13)).toNumber());
          _nw189[(_dafny.ZERO).toNumber()] = p0;
          _nw189[(_dafny.ONE).toNumber()] = _1216_cf63;
          _nw189[(new BigNumber(2)).toNumber()] = p0;
          _nw189[(new BigNumber(3)).toNumber()] = (_1214_cf65).f9;
          _nw189[(new BigNumber(4)).toNumber()] = new BigNumber((p3).length);
          _nw189[(new BigNumber(5)).toNumber()] = ((_this).fm5(globalState)).minus(_1216_cf63);
          _nw189[(new BigNumber(6)).toNumber()] = new BigNumber(588);
          _nw189[(new BigNumber(7)).toNumber()] = new BigNumber(882);
          _nw189[(new BigNumber(8)).toNumber()] = p0;
          _nw189[(new BigNumber(9)).toNumber()] = (_1216_cf63).plus(_1217_cf62);
          _nw189[(new BigNumber(10)).toNumber()] = (_1214_cf65).f9;
          _nw189[(new BigNumber(11)).toNumber()] = _module.__default.safeDivisionInt(new BigNumber(((((_1227_v82).contains(_1222_v77.f12)) ? ((_1227_v82).get(_1222_v77.f12)) : ((((_1227_v82).contains(p2)) ? ((_1227_v82).get(p2)) : ((_1222_v77).f11))))).length), _1217_cf62);
          _nw189[(new BigNumber(12)).toNumber()] = (_1214_cf65).fm17(((_1214_cf65).f10)[_module.__default.safeIndex(_1216_cf63, new BigNumber(((_1214_cf65).f10).length))], _1230_v85, globalState);
          _1231_v86 = _nw189;
          let _index172 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_1231_v86).length));
          (_1231_v86)[_index172] = (_1214_cf65).f9;
          let _index173 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_1231_v86).length));
          (_1231_v86)[_index173] = new BigNumber((_1221_v76).length);
          let _1232_v87;
          _1232_v87 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus(_1217_cf62),(_1222_v77).f11);
          _1221_v76 = _dafny.Seq.Concat((_1222_v77).f11, (((_1232_v87).contains((_dafny.ZERO).minus((_1214_cf65).f9))) ? ((_1232_v87).get((_dafny.ZERO).minus((_1214_cf65).f9))) : ((_1222_v77).f11)));
          let _index174 = _module.__default.safeIndex(new BigNumber(300), new BigNumber((_1231_v86).length));
          (_1231_v86)[_index174] = (p0).multipliedBy((_1231_v86)[_module.__default.safeIndex(new BigNumber(300), new BigNumber((_1231_v86).length))]);
          (globalState).f5 = _1218_cf61;
        } else {
          let _1233_v88;
          let _init36 = ((_1234_cf65) => function (_1235_i10) {
            return _module.__default.safeModuloInt(_1235_i10, (_1234_cf65).f9);
          })(_1214_cf65);
          let _nw190 = Array((new BigNumber(19)).toNumber());
          for (let _i0_36 = 0; _i0_36 < new BigNumber(_nw190.length); _i0_36++) {
            _nw190[_i0_36] = _init36(new BigNumber(_i0_36));
          }
          _1233_v88 = _nw190;
          _1233_v88 = _1233_v88;
          (globalState).f5 = _1222_v77.f12;
          let _1236_v89;
          _1236_v89 = _dafny.Set.fromElements((p0).plus((_1214_cf65).f9), p0, _module.__default.fm15(new BigNumber(-984), _1218_cf61, globalState));
          _1236_v89 = _1236_v89;
          let _1237_v90;
          let _nw191 = new _module.C1();
          _nw191.__ctor();
          _1237_v90 = _nw191;
          let _1238_v91;
          let _nw192 = new _module.C3();
          _nw192.__ctor(_1216_cf63, (_1214_cf65).f10, (_1214_cf65).f9);
          _1238_v91 = _nw192;
        }
      } else if (_source21.is_DC31) {
        let _1239___mcc_h33 = (_source21).cf55;
        let _1240_cf55 = _1239___mcc_h33;
        let _1241_v92;
        _1241_v92 = _dafny.Seq.of(_module.__default.fm36((_this).f7, (_this).f7, (_this).f7, new _dafny.CodePoint('x'.codePointAt(0)), globalState));
        (globalState).f1 = !(_dafny.areEqual(_1241_v92, _1241_v92));
        let _1242_v93;
        let _init37 = function (_1243_i12) {
          return _module.__default.safeModuloInt(_1243_i12, new BigNumber(-109));
        };
        let _nw193 = Array((new BigNumber(17)).toNumber());
        for (let _i0_37 = 0; _i0_37 < new BigNumber(_nw193.length); _i0_37++) {
          _nw193[_i0_37] = _init37(new BigNumber(_i0_37));
        }
        _1242_v93 = _nw193;
        let _1244_v94;
        _1244_v94 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(910)), ((_1245_p1) => function (_1246_i11) {
          return new BigNumber((_dafny.Set.fromElements(_1245_p1, (_this).f7)).length);
        })(p1)),_1242_v93);
        let _1247_v95;
        _1247_v95 = _module.D8.create_DC20((_1244_v94).Merge(_1244_v94));
        let _source22 = _1247_v95;
        if (_source22.is_DC21) {
          let _1248___mcc_h35 = (_source22).cf36;
          let _1249___mcc_h36 = (_source22).cf37;
          let _1250___mcc_h37 = (_source22).cf38;
          let _1251_cf38 = _1250___mcc_h37;
          let _1252_cf37 = _1249___mcc_h36;
          let _1253_cf36 = _1248___mcc_h35;
          let _1254_v96;
          _1254_v96 = _dafny.MultiSet.fromElements(_1253_cf36);
          let _1255_v97;
          _1255_v97 = _dafny.Seq.UnicodeFromString("ngnbjrkg");
          let _1256_v98;
          _1256_v98 = _dafny.Map.Empty.slice().updateUnsafe(_1255_v97,_1251_cf38);
          let _1257_v99;
          _1257_v99 = _dafny.Map.Empty.slice().updateUnsafe(_1253_cf36,_1254_v96);
          let _1258_v100;
          _1258_v100 = _dafny.Map.Empty.slice().updateUnsafe(p1,p1);
          let _1259_v101;
          _1259_v101 = _dafny.Seq.of(_1254_v96);
          let _1260_v102;
          let _nw194 = Array((new BigNumber(26)).toNumber());
          _nw194[(_dafny.ZERO).toNumber()] = _1254_v96;
          _nw194[(_dafny.ONE).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(2)).toNumber()] = (_dafny.MultiSet.FromArray(p3)).Union(_1254_v96);
          _nw194[(new BigNumber(3)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(4)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(5)).toNumber()] = (_1254_v96).Union(_1254_v96);
          _nw194[(new BigNumber(6)).toNumber()] = _dafny.MultiSet.fromElements(_1253_cf36, _1252_cf37, _1252_cf37);
          _nw194[(new BigNumber(7)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(8)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(9)).toNumber()] = _module.__default.fm30(_dafny.Seq.UnicodeFromString("ae"), (_dafny.Map.Empty.slice().updateUnsafe(_1255_v97,_1251_cf38)).update(_dafny.Seq.UnicodeFromString("apq"), p0), (_this).f7, globalState);
          _nw194[(new BigNumber(10)).toNumber()] = _module.__default.fm30(_dafny.Seq.update(_1255_v97, _module.__default.safeIndex(_1251_cf38, new BigNumber((_1255_v97).length)), new _dafny.CodePoint('k'.codePointAt(0))), _1256_v98, p2, globalState);
          _nw194[(new BigNumber(11)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(12)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(13)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(14)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(15)).toNumber()] = (_1254_v96).Union((((_1257_v99).contains((((_1258_v100).contains((_this).f7)) ? ((_1258_v100).get((_this).f7)) : (false)))) ? ((_1257_v99).get((((_1258_v100).contains((_this).f7)) ? ((_1258_v100).get((_this).f7)) : (false)))) : (_1254_v96)));
          _nw194[(new BigNumber(16)).toNumber()] = (_1254_v96).Intersect(_dafny.MultiSet.FromArray(p3));
          _nw194[(new BigNumber(17)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(18)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(19)).toNumber()] = _dafny.MultiSet.FromArray(p3);
          _nw194[(new BigNumber(20)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(21)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(22)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(23)).toNumber()] = _1254_v96;
          _nw194[(new BigNumber(24)).toNumber()] = (_1259_v101)[_module.__default.safeIndex(_1251_cf38, new BigNumber((_1259_v101).length))];
          _nw194[(new BigNumber(25)).toNumber()] = _1254_v96;
          _1260_v102 = _nw194;
          let _index175 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1260_v102).length));
          (_1260_v102)[_index175] = _1254_v96;
          let _index176 = _module.__default.safeIndex(new BigNumber(165), new BigNumber((_1260_v102).length));
          (_1260_v102)[_index176] = _1254_v96;
          let _1261_v103;
          let _nw195 = new _module.C2();
          _nw195.__ctor(_dafny.Seq.update(_1255_v97, _module.__default.safeIndex(_1251_cf38, new BigNumber((_1255_v97).length)), (_this).f8), _1253_cf36);
          _1261_v103 = _nw195;
          let _1262_v104;
          _1262_v104 = _dafny.Map.Empty.slice().updateUnsafe(p2,_1261_v103);
          (globalState).f0 = new BigNumber((_1262_v104).length);
          let _1263_v105;
          let _nw196 = new _module.C1();
          _nw196.__ctor();
          _1263_v105 = _nw196;
          let _1264_v106;
          _1264_v106 = _dafny.Seq.of(((p2) ? (_1263_v105) : (_1263_v105)));
          _1263_v105 = (_1264_v106)[_module.__default.safeIndex(new BigNumber(909), new BigNumber((_1264_v106).length))];
          let _1265_v107;
          let _init38 = ((_1266_v103) => function (_1267_i13) {
            return _1266_v103.f12;
          })(_1261_v103);
          let _nw197 = Array((new BigNumber(6)).toNumber());
          for (let _i0_38 = 0; _i0_38 < new BigNumber(_nw197.length); _i0_38++) {
            _nw197[_i0_38] = _init38(new BigNumber(_i0_38));
          }
          _1265_v107 = _nw197;
          _1265_v107 = _1265_v107;
        } else {
          let _1268___mcc_h38 = (_source22).cf35;
          let _1269_cf35 = _1268___mcc_h38;
          let _1270_v108;
          _1270_v108 = _module.D1.create_DC4(_dafny.Seq.Create(_module.__default.abs(new BigNumber(986)), ((_1271_p0) => function (_1272_i14) {
  return _1271_p0;
})(p0)));
          let _1273_v109;
          let _nw198 = new _module.C3();
          _nw198.__ctor(p0, (_1270_v108).dtor_cf11, _module.__default.safeModuloInt(new BigNumber(412), (_dafny.ZERO).minus(new BigNumber(-262))));
          _1273_v109 = _nw198;
          let _index177 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_1242_v93).length));
          (_1242_v93)[_index177] = (_1273_v109).f9;
          let _index178 = _module.__default.safeIndex(new BigNumber(179), new BigNumber((_1242_v93).length));
          (_1242_v93)[_index178] = new BigNumber((_dafny.Set.fromElements(new BigNumber(366), (_1273_v109).f9, ((_1273_v109).f10)[_module.__default.safeIndex((_1273_v109).f9, new BigNumber(((_1273_v109).f10).length))], p0, (_1273_v109).f9)).length);
          let _1274_v110;
          let _nw199 = new _module.C3();
          _nw199.__ctor((_1273_v109).f9, _dafny.Seq.of((_1273_v109).f9), (_1242_v93)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_1242_v93).length))]);
          _1274_v110 = _nw199;
          let _1275_v111;
          _1275_v111 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_1242_v93)[_module.__default.safeIndex(new BigNumber(179), new BigNumber((_1242_v93).length))]),new BigNumber((p3).length));
          _1275_v111 = _1275_v111;
        }
        _1242_v93 = ((p1) ? (_1242_v93) : (_1242_v93));
        (globalState).f5 = p2;
      } else {
        let _1276___mcc_h34 = (_source21).cf66;
        let _1277_cf66 = _1276___mcc_h34;
        let _1278_v112;
        _1278_v112 = _dafny.Seq.of(new BigNumber(77));
        let _1279_v113;
        _1279_v113 = _module.D1.create_DC4(_1278_v112);
        let _1280_v114;
        let _nw200 = Array((new BigNumber(27)).toNumber()).fill(_dafny.ZERO);
        _1280_v114 = _nw200;
        let _1281_v115;
        _1281_v115 = _module.D3.create_DC8(_1279_v113, _1280_v114, p0);
        let _1282_v116;
        _1282_v116 = _dafny.MultiSet.fromElements(p0, new BigNumber((_dafny.Seq.of(p2, !(true), p2)).length));
        let _1283_v117;
        let _nw201 = new _module.C3();
        _nw201.__ctor((_1281_v115).dtor_cf16, _dafny.Seq.of(p0, (_dafny.ZERO).minus(new BigNumber((_1282_v116).cardinality())), p0), p0);
        _1283_v117 = _nw201;
        (globalState).f5 = !((_dafny.ZERO).minus((_1283_v117).f9)).isEqualTo(p0);
        let _1284_v118;
        _1284_v118 = _dafny.Seq.UnicodeFromString("lhbgnik");
        let _rhs181 = _dafny.Seq.Concat(_dafny.Seq.Concat(_1278_v112, (_1283_v117).f10), _dafny.Seq.of((_1283_v117).f9));
        let _rhs182 = p0;
        let _rhs183 = true;
        let _rhs184 = !((_1283_v117).f9).isEqualTo(new BigNumber((_dafny.Seq.Create(_module.__default.abs(new BigNumber(193)), function (_1285_i15) {
          return (_this).f8;
        })).length));
        let _rhs185 = !_dafny.areEqual(_module.__default.fm2(_dafny.Seq.UnicodeFromString("f"), globalState), _dafny.Seq.Concat(_1284_v118, _1284_v118));
        let _lhs139 = globalState;
        let _lhs140 = globalState;
        let _lhs141 = globalState;
        let _lhs142 = globalState;
        _1278_v112 = _rhs181;
        _lhs139.f2 = _rhs182;
        _lhs140.f1 = _rhs183;
        _lhs141.f5 = _rhs184;
        _lhs142.f5 = _rhs185;
        (globalState).f0 = (((_dafny.ZERO).minus(new BigNumber((_1284_v118).length))).plus((_1283_v117).f9)).plus((new BigNumber((_1282_v116).cardinality())).plus(new BigNumber(108)));
      }
      r0 = p2;
      return r0;
    }
    get f7() {
      let _this = this;
      return _this._f7;
    };
    get f8() {
      let _this = this;
      return _this._f8;
    };
  };

  $module.C6 = class C6 {
    constructor () {
      this._tname = "_module.C6";
      this._f6 = _dafny.ZERO;
    }
    _parentTraits() {
      return [_module.T0];
    }
    get f6() {
      let _this = this;
      return _this._f6;
    };
    __ctor(f6) {
      let _this = this;
      (_this)._f6 = f6;
      return;
    }
    fm3(p0, globalState) {
      let _this = this;
      return _module.__default.safeDivisionInt(_module.__default.safeDivisionInt((_this).f6, new BigNumber((_dafny.Seq.of(_dafny.Map.Empty.slice().updateUnsafe(true,false))).length)), ((_this).f6).multipliedBy(new BigNumber((_dafny.Seq.of(new BigNumber(166))).length)));
    };
    m0(globalState) {
      let _this = this;
      let r0 = _dafny.ZERO;
      let r1 = false;
      let r2 = false;
      let _1286_v0;
      _1286_v0 = new _dafny.CodePoint('a'.codePointAt(0));
      let _1287_v1;
      _1287_v1 = _dafny.Seq.of(_1286_v0, _1286_v0);
      let _1288_v2;
      _1288_v2 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,(_dafny.ZERO).minus(new BigNumber((_dafny.Seq.of((_this).f6, new BigNumber((_1287_v1).length))).length)));
      let _1289_v3;
      _1289_v3 = true;
      let _1290_v4;
      _1290_v4 = _dafny.Seq.of(!(_1288_v2).equals(_1288_v2), _1289_v3, _1289_v3, _1289_v3);
      _1290_v4 = _dafny.Seq.update(_1290_v4, _module.__default.safeIndex((_this).f6, new BigNumber((_1290_v4).length)), !((_1290_v4)[_module.__default.safeIndex(new BigNumber((_dafny.Seq.of(_1289_v3)).length), new BigNumber((_1290_v4).length))]));
      let _hi8 = (_this).f6;
      for (let _1291_i0 = (_this).f6; _1291_i0.isLessThan(_hi8); _1291_i0 = _1291_i0.plus(_dafny.ONE)) {
        let _1292_v5;
        _1292_v5 = _module.D0.create_DC2(_1291_i0, _1286_v0, new BigNumber(618), (_this).f6, _1291_i0);
        let _1293_v6;
        _1293_v6 = _dafny.Set.fromElements(_1289_v3, _1289_v3);
        let _1294_v7;
        let _nw202 = Array((new BigNumber(14)).toNumber());
        _nw202[(_dafny.ZERO).toNumber()] = _1286_v0;
        _nw202[(_dafny.ONE).toNumber()] = new _dafny.CodePoint('o'.codePointAt(0));
        _nw202[(new BigNumber(2)).toNumber()] = _1286_v0;
        _nw202[(new BigNumber(3)).toNumber()] = _1286_v0;
        _nw202[(new BigNumber(4)).toNumber()] = _1286_v0;
        _nw202[(new BigNumber(5)).toNumber()] = _1286_v0;
        _nw202[(new BigNumber(6)).toNumber()] = _1286_v0;
        _nw202[(new BigNumber(7)).toNumber()] = _1286_v0;
        _nw202[(new BigNumber(8)).toNumber()] = _1286_v0;
        _nw202[(new BigNumber(9)).toNumber()] = _1286_v0;
        _nw202[(new BigNumber(10)).toNumber()] = ((_1289_v3) ? (_1286_v0) : (_1286_v0));
        _nw202[(new BigNumber(11)).toNumber()] = _1286_v0;
        _nw202[(new BigNumber(12)).toNumber()] = _1286_v0;
        _nw202[(new BigNumber(13)).toNumber()] = _module.__default.fm4(_dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.of((_dafny.ZERO).minus(_1291_i0)),_1292_v5), (_dafny.ZERO).minus(_1291_i0), new BigNumber((_1293_v6).length), globalState);
        _1294_v7 = _nw202;
        let _index179 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1294_v7).length));
        (_1294_v7)[_index179] = _1286_v0;
        let _index180 = _module.__default.safeIndex(new BigNumber(387), new BigNumber((_1294_v7).length));
        (_1294_v7)[_index180] = _1286_v0;
        (globalState).f5 = false;
        if (true) {
          (globalState).f0 = ((_1289_v3) ? ((_this).f6) : ((_this).f6));
          let _1295_v8;
          _1295_v8 = _dafny.Set.fromElements(_dafny.Seq.of(new BigNumber(664), _1291_i0, _1291_i0, (_this).f6), _dafny.Seq.Create(_module.__default.abs(new BigNumber(144)), ((_1296_i0) => function (_1297_i1) {
            return _1296_i0;
          })(_1291_i0)));
          let _1298_v9;
          _1298_v9 = _dafny.Map.Empty.slice().updateUnsafe(!((_1295_v8).IsSubsetOf(_1295_v8)),_1293_v6);
          _1298_v9 = (_1298_v9).Merge(_1298_v9);
          let _1299_v10;
          _1299_v10 = _dafny.Map.Empty.slice().updateUnsafe(_1291_i0,_1289_v3);
          let _1300_v11;
          _1300_v11 = _dafny.Map.Empty.slice().updateUnsafe((new BigNumber(187)).plus(new BigNumber((_1299_v10).length)),_1289_v3);
          _1300_v11 = (_1300_v11).update(_module.__default.safeDivisionInt(_1291_i0, (_this).f6), !(_1289_v3));
          _1299_v10 = (_1299_v10).update((_1292_v5).dtor_cf9, ((_this).f6).isEqualTo((_this).f6));
          let _1301_v12;
          let _nw203 = Array((new BigNumber(26)).toNumber()).fill(_dafny.Seq.UnicodeFromString(""));
          _1301_v12 = _nw203;
          _1301_v12 = _1301_v12;
        } else {
          let _1302_v13;
          let _nw204 = new _module.C1();
          _nw204.__ctor();
          _1302_v13 = _nw204;
          let _1303_v14;
          let _nw205 = new _module.C0();
          _nw205.__ctor(true);
          _1303_v14 = _nw205;
          let _1304_v15;
          _1304_v15 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v3,_1303_v14);
          let _1305_v16;
          _1305_v16 = _dafny.Map.Empty.slice().updateUnsafe(_1304_v15,(_dafny.ZERO).minus(_1291_i0));
          let _1306_v17;
          _1306_v17 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1305_v16).length),(_1303_v14).f13);
          _1306_v17 = (_1306_v17).update((_this).f6, (_1303_v14).f13);
          (globalState).f2 = new BigNumber(((_dafny.Map.Empty.slice().updateUnsafe(!(_1289_v3),new _dafny.CodePoint('y'.codePointAt(0)))).update((_1303_v14).f13, (_1294_v7)[_module.__default.safeIndex(new BigNumber(387), new BigNumber((_1294_v7).length))])).length);
          let _1307_v18;
          _1307_v18 = _dafny.Set.fromElements(_1291_i0);
          let _1308_v19;
          _1308_v19 = _dafny.Seq.of(_1291_i0, new BigNumber((_1307_v18).length));
          let _1309_v20;
          let _nw206 = new _module.C3();
          _nw206.__ctor(_1291_i0, _1308_v19, (_this).f6);
          _1309_v20 = _nw206;
          let _rhs186 = (_module.D13.create_DC33((_1303_v14).f13, (_dafny.ZERO).minus(new BigNumber((_1288_v2).length)), _1291_i0, (_1303_v14).f13, _1309_v20)).dtor_cf64;
          let _rhs187 = (_this).f6;
          let _lhs143 = globalState;
          let _lhs144 = globalState;
          _lhs143.f5 = _rhs186;
          _lhs144.f0 = _rhs187;
          (globalState).f1 = !((_1303_v14).f13);
        }
        _1287_v1 = _1287_v1;
      }
      let _1310_v21;
      _1310_v21 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v3,_dafny.Seq.Create(_module.__default.abs(new BigNumber(707)), ((_1311_v0) => function (_1312_i2) {
        return _1311_v0;
      })(_1286_v0)));
      let _1313_v22;
      _1313_v22 = _dafny.MultiSet.fromElements((((_1310_v21).contains(_1289_v3)) ? ((_1310_v21).get(_1289_v3)) : (_1287_v1)), _dafny.Seq.Concat(_1287_v1, _dafny.Seq.UnicodeFromString("gfyelbune")));
      _1313_v22 = ((_1313_v22).Difference(_dafny.MultiSet.fromElements(_dafny.Seq.Create(_module.__default.abs(new BigNumber(471)), ((_1314_v0) => function (_1315_i3) {
        return _1314_v0;
      })(_1286_v0))))).update(_1287_v1, _module.__default.abs((_this).f6));
      let _1316_v23;
      _1316_v23 = _dafny.MultiSet.fromElements(_1289_v3, _1289_v3, !(_1289_v3), _1289_v3);
      if ((_1316_v23).IsDisjointFrom((_1316_v23).Union(_1316_v23))) {
        let _1317_v24;
        _1317_v24 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v3,_1289_v3);
        let _1318_v25;
        _1318_v25 = _dafny.Seq.of((_this).f6, new BigNumber((_1317_v24).length));
        let _1319_v26;
        _1319_v26 = _dafny.Set.fromElements(new BigNumber((_1318_v25).length), (_this).f6, (_this).f6);
        let _1320_v27;
        let _nw207 = new _module.C1();
        _nw207.__ctor();
        _1320_v27 = _nw207;
        let _1321_v28;
        _1321_v28 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.Set.fromElements((_1318_v25)[_module.__default.safeIndex((_this).f6, new BigNumber((_1318_v25).length))])).IsProperSubsetOf(_1319_v26),_1320_v27);
        _1321_v28 = (_1321_v28).update(_1289_v3, _1320_v27);
        let _1322_v29;
        _1322_v29 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v3,(_this).f6);
        let _1323_v30;
        _1323_v30 = _dafny.Map.Empty.slice().updateUnsafe(_1322_v29,_1289_v3);
        let _rhs188 = ((_this).f6).multipliedBy(new BigNumber((_1323_v30).length));
        let _rhs189 = (_this).f6;
        let _rhs190 = _dafny.Seq.update(_dafny.Seq.UnicodeFromString("npsdnb"), _module.__default.safeIndex(new BigNumber(954), new BigNumber((_dafny.Seq.UnicodeFromString("npsdnb")).length)), _1286_v0);
        let _lhs145 = globalState;
        let _lhs146 = globalState;
        _lhs145.f0 = _rhs188;
        _lhs146.f0 = _rhs189;
        _1287_v1 = _rhs190;
        let _1324_v32;
        _1324_v32 = _module.D0.create_DC2((_this).f6, _1286_v0, (_this).f6, new BigNumber(101), (((_1288_v2).contains(new BigNumber(-905))) ? ((_1288_v2).get(new BigNumber(-905))) : (new BigNumber(-20))));
        let _1325_v33;
        _1325_v33 = _dafny.Set.fromElements(_1286_v0, (_1324_v32).dtor_cf6);
        (globalState).f5 = !((function () {
          let _coll55 = new _dafny.Map();
          for (const _compr_55 of (_1325_v33).Elements) {
            let _1326_v31 = _compr_55;
            if ((_1325_v33).contains(_1326_v31)) {
              _coll55.push([_1326_v31,(_this).f6]);
            }
          }
          return _coll55;
        }()).Merge(_dafny.Map.Empty.slice().updateUnsafe(_1286_v0,(_this).f6))).contains(((!(!(!(_1289_v3)))) ? (_1286_v0) : (_1286_v0)));
        _1317_v24 = (_1317_v24).update(_1289_v3, _1289_v3);
        let _1327_v34;
        let _1328_v35;
        let _1329_v36;
        let _out12;
        let _out13;
        let _out14;
        let _outcollector3 = (_1320_v27).m6(globalState);
        _out12 = _outcollector3[0];
        _out13 = _outcollector3[1];
        _out14 = _outcollector3[2];
        _1327_v34 = _out12;
        _1328_v35 = _out13;
        _1329_v36 = _out14;
      } else {
        r0 = new BigNumber(569);
        if ((_1289_v3) === (_1289_v3)) {
          let _1330_v37;
          _1330_v37 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v3,(_this).f6);
          let _1331_v38;
          _1331_v38 = _dafny.Map.Empty.slice().updateUnsafe(new BigNumber((_1330_v37).length),_1289_v3);
          let _1332_v39;
          _1332_v39 = _dafny.MultiSet.fromElements(_1331_v38, _1331_v38);
          let _1333_v40;
          _1333_v40 = _dafny.Seq.of(_1332_v39);
          (globalState).f5 = ((_1333_v40)[_module.__default.safeIndex((_this).f6, new BigNumber((_1333_v40).length))]).contains(_1331_v38);
          _1330_v37 = (_1330_v37).update(_1289_v3, (_this).f6);
          (globalState).f0 = _module.__default.fm24(new BigNumber((_1287_v1).length), globalState);
          let _1334_v41;
          _1334_v41 = _module.D12.create_DC30(_module.__default.fm37(_1289_v3, (_this).f6, _1316_v23, globalState));
          let _1335_v42;
          _1335_v42 = _module.D12.create_DC28(_1330_v37);
          _1334_v41 = _module.D12.create_DC30(_1335_v42);
          let _1336_v43;
          let _nw208 = Array((new BigNumber(28)).toNumber()).fill([]);
          _1336_v43 = _nw208;
          let _1337_v44;
          let _nw209 = Array((new BigNumber(7)).toNumber()).fill(_dafny.MultiSet.Empty);
          _1337_v44 = _nw209;
          let _index181 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_1336_v43).length));
          (_1336_v43)[_index181] = _1337_v44;
          let _index182 = _module.__default.safeIndex(new BigNumber(324), new BigNumber((_1336_v43).length));
          let _nw210 = Array((new BigNumber(23)).toNumber()).fill(_dafny.MultiSet.Empty);
          (_1336_v43)[_index182] = _nw210;
        } else {
          let _1338_v45;
          _1338_v45 = _dafny.Set.fromElements(_1289_v3, _1289_v3);
          let _1339_v46;
          _1339_v46 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1338_v45);
          let _1340_v47;
          _1340_v47 = _module.D0.create_DC0(_1339_v46);
          let _1341_v48;
          let _nw211 = new _module.C1();
          _nw211.__ctor();
          _1341_v48 = _nw211;
          let _1342_v49;
          _1342_v49 = _dafny.Set.fromElements(_1286_v0);
          let _1343_v50;
          _1343_v50 = _dafny.Map.Empty.slice().updateUnsafe(_1341_v48,new BigNumber((_1342_v49).length));
          let _1344_v51;
          _1344_v51 = _dafny.MultiSet.fromElements(_dafny.Set.fromElements(!(_1289_v3), _1289_v3));
          let _1345_v52;
          _1345_v52 = _dafny.Map.Empty.slice().updateUnsafe(_module.__default.fm1(_1340_v47, _1289_v3, new BigNumber((_1343_v50).length), _1289_v3, globalState),(_1344_v51).Union(_1344_v51));
          _1345_v52 = (_1345_v52).update(_1289_v3, _1344_v51);
          let _1346_v53;
          _1346_v53 = _dafny.Set.fromElements(_dafny.Seq.Concat(_module.__default.fm2(_dafny.Seq.UnicodeFromString("bq"), globalState), _1287_v1));
          let _1347_v54;
          _1347_v54 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.Create(_module.__default.abs(new BigNumber(930)), ((_1348_v0) => function (_1349_i4) {
            return _1348_v0;
          })(_1286_v0)),new BigNumber(-965));
          _1346_v53 = ((_1289_v3) ? (function () {
            let _coll56 = new _dafny.Set();
            for (const _compr_56 of (_1347_v54).Keys.Elements) {
              let _1350_v55 = _compr_56;
              if ((_1347_v54).contains(_1350_v55)) {
                _coll56.add(_1350_v55);
              }
            }
            return _coll56;
          }()) : (_1346_v53));
          (globalState).f1 = _1289_v3;
          let _1351_v56;
          _1351_v56 = _dafny.Seq.of((_this).f6, (_dafny.ZERO).minus(_module.__default.safeModuloInt((_this).f6, (_this).f6)));
          let _1352_v57;
          let _nw212 = Array((new BigNumber(6)).toNumber());
          _nw212[(_dafny.ZERO).toNumber()] = _1286_v0;
          _nw212[(_dafny.ONE).toNumber()] = _1286_v0;
          _nw212[(new BigNumber(2)).toNumber()] = _1286_v0;
          _nw212[(new BigNumber(3)).toNumber()] = _1286_v0;
          _nw212[(new BigNumber(4)).toNumber()] = _1286_v0;
          _nw212[(new BigNumber(5)).toNumber()] = _1286_v0;
          _1352_v57 = _nw212;
          let _1353_v58;
          _1353_v58 = _dafny.MultiSet.fromElements(_1352_v57, _1352_v57);
          let _1354_v59;
          _1354_v59 = _dafny.Map.Empty.slice().updateUnsafe((_1290_v4)[_module.__default.safeIndex((_this).f6, new BigNumber((_1290_v4).length))],_1289_v3);
          let _rhs191 = _dafny.Seq.update(_1351_v56, _module.__default.safeIndex(new BigNumber((_dafny.Seq.Concat(_dafny.Seq.update(_1287_v1, _module.__default.safeIndex((_dafny.ZERO).minus(new BigNumber((_1353_v58).cardinality())), new BigNumber((_1287_v1).length)), _1286_v0), _dafny.Seq.update(_1287_v1, _module.__default.safeIndex(new BigNumber((_1351_v56).length), new BigNumber((_1287_v1).length)), _1286_v0))).length), new BigNumber((_1351_v56).length)), new BigNumber((_1354_v59).length));
          let _rhs192 = ((_1289_v3) ? ((_this).f6) : ((_this).f6));
          let _lhs147 = globalState;
          _1351_v56 = _rhs191;
          _lhs147.f2 = _rhs192;
          let _1355_v60;
          let _nw213 = Array((new BigNumber(19)).toNumber()).fill([]);
          _1355_v60 = _nw213;
          let _1356_v61;
          let _nw214 = Array((_dafny.ONE).toNumber()).fill(_dafny.ZERO);
          _1356_v61 = _nw214;
          let _index183 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_1355_v60).length));
          (_1355_v60)[_index183] = _1356_v61;
          let _index184 = _module.__default.safeIndex(new BigNumber(426), new BigNumber((_1355_v60).length));
          let _nw215 = Array((new BigNumber(10)).toNumber()).fill(_dafny.ZERO);
          (_1355_v60)[_index184] = _nw215;
        }
        let _1357_v62;
        _1357_v62 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v3,_dafny.Seq.of(_1289_v3));
        let _1358_v63;
        _1358_v63 = _dafny.Map.Empty.slice().updateUnsafe(_dafny.Seq.update((((_1357_v62).contains(_1289_v3)) ? ((_1357_v62).get(_1289_v3)) : (_dafny.Seq.of(_1289_v3))), _module.__default.safeIndex((_this).f6, new BigNumber(((((_1357_v62).contains(_1289_v3)) ? ((_1357_v62).get(_1289_v3)) : (_dafny.Seq.of(_1289_v3)))).length)), _1289_v3),_1289_v3);
        _1358_v63 = (_1358_v63).update(_dafny.Seq.of(_1289_v3, (_1290_v4)[_module.__default.safeIndex((_this).f6, new BigNumber((_1290_v4).length))], _1289_v3, _1289_v3, _1289_v3), _1289_v3);
        if (_1289_v3) {
          (globalState).f1 = _1289_v3;
          r2 = _1289_v3;
          let _1359_v64;
          _1359_v64 = _dafny.MultiSet.fromElements((_this).f6, new BigNumber(201));
          r1 = ((((_1359_v64).contains((_dafny.ZERO).minus((_this).f6))) ? ((_1359_v64).get((_dafny.ZERO).minus((_this).f6))) : ((_this).f6))).isLessThan((_module.__default.fm15((_this).f6, _1289_v3, globalState)).multipliedBy(new BigNumber(-22)));
          let _1360_v65;
          let _nw216 = new _module.C1();
          _nw216.__ctor();
          _1360_v65 = _nw216;
          _1360_v65 = ((!(_1289_v3)) ? (_1360_v65) : (_1360_v65));
          let _1361_v66;
          let _nw217 = new _module.C2();
          _nw217.__ctor(_1287_v1, _1289_v3);
          _1361_v66 = _nw217;
        } else {
          (globalState).f5 = _1289_v3;
          let _1362_v67;
          _1362_v67 = _dafny.Map.Empty.slice().updateUnsafe(_1286_v0,(_this).f6);
          let _1363_v68;
          _1363_v68 = _module.D17.create_DC40(_1362_v67);
          let _1364_v69;
          _1364_v69 = _dafny.MultiSet.fromElements(((_1363_v68).dtor_cf76).Merge(_1362_v67), _1362_v67);
          _1364_v69 = (_1364_v69).update(_1362_v67, _module.__default.abs(new BigNumber(152)));
          let _1365_v70;
          _1365_v70 = _dafny.MultiSet.fromElements((_this).f6);
          let _1366_v71;
          _1366_v71 = _dafny.Seq.of(_1365_v70);
          (globalState).f1 = _dafny.Seq.IsProperPrefixOf(_dafny.Seq.of(_1365_v70), _1366_v71);
          let _1367_v72;
          _1367_v72 = _dafny.Map.Empty.slice().updateUnsafe((_dafny.ZERO).minus((_this).f6),true);
          _1367_v72 = _1367_v72;
          let _1368_v73;
          _1368_v73 = _dafny.Set.fromElements((_this).f6, new BigNumber((_1287_v1).length), (_this).f6);
          let _1369_v74;
          _1369_v74 = _dafny.Seq.of((_this).f6);
          let _1370_v75;
          let _nw218 = new _module.C3();
          _nw218.__ctor(new BigNumber((_1368_v73).length), _1369_v74, (_this).f6);
          _1370_v75 = _nw218;
          let _1371_v76;
          _1371_v76 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v3,_1370_v75);
          let _1372_v77;
          _1372_v77 = _dafny.Map.Empty.slice().updateUnsafe(_1371_v76,(_1370_v75).f6);
          let _1373_v78;
          _1373_v78 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v3,_1289_v3);
          (globalState).f1 = ((new BigNumber((_1372_v77).length)).plus((_1370_v75).f6)).isLessThan(new BigNumber((_1373_v78).length));
        }
        _1286_v0 = _1286_v0;
      }
      let _1374_v79;
      _1374_v79 = _dafny.Seq.of((_this).f6);
      let _1375_v80;
      _1375_v80 = _dafny.Set.fromElements((_this).f6, (_this).f6, new BigNumber(((_module.D1.create_DC4(_1374_v79)).dtor_cf11).length), new BigNumber(621), new BigNumber(-963));
      _1289_v3 = !((true) || ((_1375_v80).equals(_dafny.Set.fromElements((_this).f6, (_dafny.ZERO).minus((_this).f6)))));
      let _1376_v81;
      _1376_v81 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v3,_1374_v79);
      let _1377_v82;
      let _nw219 = new _module.C3();
      _nw219.__ctor(_module.__default.fm15((_this).f6, _1289_v3, globalState), (((_1376_v81).contains(_1289_v3)) ? ((_1376_v81).get(_1289_v3)) : (_dafny.Seq.of((_this).f6, (_this).f6))), (_this).f6);
      _1377_v82 = _nw219;
      let _1378_v83;
      _1378_v83 = _dafny.Set.fromElements(_1289_v3);
      let _1379_v84;
      _1379_v84 = _dafny.Map.Empty.slice().updateUnsafe((_this).f6,_1378_v83);
      let _1380_v85;
      _1380_v85 = _module.D0.create_DC0(_1379_v84);
      let _1381_v86;
      _1381_v86 = _dafny.Map.Empty.slice().updateUnsafe(_1289_v3,_1289_v3);
      let _1382_v87;
      let _nw220 = Array((new BigNumber(26)).toNumber());
      _nw220[(_dafny.ZERO).toNumber()] = _1289_v3;
      _nw220[(_dafny.ONE).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(2)).toNumber()] = ((_1289_v3) ? (_1289_v3) : (!(_1289_v3)));
      _nw220[(new BigNumber(3)).toNumber()] = false;
      _nw220[(new BigNumber(4)).toNumber()] = (_module.D13.create_DC33(true, new BigNumber((_1288_v2).length), (_this).f6, _1289_v3, _1377_v82)).dtor_cf61;
      _nw220[(new BigNumber(5)).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(6)).toNumber()] = (_1316_v23).equals(_1316_v23);
      _nw220[(new BigNumber(7)).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(8)).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(9)).toNumber()] = (_1375_v80).IsDisjointFrom(_1375_v80);
      _nw220[(new BigNumber(10)).toNumber()] = (_1289_v3) === (_1289_v3);
      _nw220[(new BigNumber(11)).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(12)).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(13)).toNumber()] = (_1290_v4)[_module.__default.safeIndex(_module.__default.fm24((_this).f6, globalState), new BigNumber((_1290_v4).length))];
      _nw220[(new BigNumber(14)).toNumber()] = (_1377_v82).fm16(_1316_v23, globalState);
      _nw220[(new BigNumber(15)).toNumber()] = ((!(_1289_v3)) ? (_1289_v3) : (_1289_v3));
      _nw220[(new BigNumber(16)).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(17)).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(18)).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(19)).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(20)).toNumber()] = _module.__default.fm1(_1380_v85, _1289_v3, (_1377_v82).f9, _1289_v3, globalState);
      _nw220[(new BigNumber(21)).toNumber()] = !(_1289_v3) || (!(_1289_v3));
      _nw220[(new BigNumber(22)).toNumber()] = ((_module.__default.fm1(_1380_v85, _1289_v3, (_1377_v82).f9, true, globalState)) ? (_1289_v3) : (true));
      _nw220[(new BigNumber(23)).toNumber()] = (_1290_v4)[_module.__default.safeIndex((_this).f6, new BigNumber((_1290_v4).length))];
      _nw220[(new BigNumber(24)).toNumber()] = _1289_v3;
      _nw220[(new BigNumber(25)).toNumber()] = ((_1289_v3) ? ((((_1381_v86).contains(_1289_v3)) ? ((_1381_v86).get(_1289_v3)) : (false))) : (_1289_v3));
      _1382_v87 = _nw220;
      let _1383_v88;
      let _nw221 = Array((new BigNumber(24)).toNumber());
      _1383_v88 = _nw221;
      let _1384_v89;
      _1384_v89 = _dafny.MultiSet.fromElements(_1383_v88, _1383_v88, _1383_v88);
      let _index185 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1382_v87).length));
      (_1382_v87)[_index185] = !((((_1384_v89).contains(_1383_v88)) ? ((_1384_v89).get(_1383_v88)) : ((_1377_v82).f9))).isEqualTo(new BigNumber((_1287_v1).length));
      let _index186 = _module.__default.safeIndex(new BigNumber(235), new BigNumber((_1382_v87).length));
      (_1382_v87)[_index186] = _1289_v3;
      let _1385_v90;
      _1385_v90 = _dafny.MultiSet.fromElements((_1377_v82).fm17((_1377_v82).f9, _1375_v80, globalState));
      let _1386_v91;
      _1386_v91 = _dafny.Map.Empty.slice().updateUnsafe(_module.D16.create_DC38(_1385_v90),_1375_v80);
      r0 = _module.__default.safeModuloInt((_dafny.ZERO).minus((_1377_v82).f9), new BigNumber((_1386_v91).length));
      r1 = _1289_v3;
      r2 = (_1316_v23).IsProperSubsetOf(_1316_v23);
      return [r0, r1, r2];
    }
  };
  return $module;
})(); // end of module _module
_dafny.HandleHaltExceptions(() => _module.__default.Main(_dafny.UnicodeFromMainArguments(require('process').argv)));
